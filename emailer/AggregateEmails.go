package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Reward struct {
	ID              gocql.UUID `json:"id"`               // cassandra uuid
	CardID          gocql.UUID `json:"card_id"`          // cassandra uuid
	Merchant        string     `json:"merchant"`         // cassandra text
	MCC             int        `json:"mcc"`              // cassandra int
	Currency        string     `json:"currency"`         // cassandra text
	Amount          float64    `json:"amount"`           // cassandra double
	SGDAmount       float64    `json:"sgd_amount"`       // cassandra double
	TransactionID   string     `json:"transaction_id"`   // cassandra text
	TransactionDate string     `json:"transaction_date"` // cassandra text
	CardPAN         string     `json:"card_pan"`         // cassandra text
	CardType        string     `json:"card_type"`        // cassandra text
	RewardAmount    float64    `json:"reward_amount"`    // cassandra double
	Remarks         string     `json:"remarks"`          // cassandra text
}

type Card struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	CardPan          string    `json:"card_pan" gorm:"unique;not null" binding:"required,credit_card"`
	UserID           uuid.UUID `json:"user_id" gorm:"type:uuid" binding:"required"` // card belongs to one user
	CardTypeCardType string    `json:"card_type" binding:"required"`                // card belongs to one card type
}

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName    string    `json:"last_name" gorm:"type:varchar(255);not null"`
	Phone       string    `json:"phone" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"-" gorm:"not null"`
	Role        string    `gorm:"type:varchar(255);not null"`
	CreditCards []Card    // one user has many credit cards
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type S3Event struct {
	Records []struct {
		S3 struct {
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

var Cassandra *gocql.Session
var Postgres *gorm.DB
var svc *ses.SES

const YYYYMMDD = "2006-01-02"

func ConnectCassandra(dbHost, dbPort, username, password, keyspace string, useSSL bool) {
	cluster := gocql.NewCluster(dbHost)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.LocalQuorum

	dbPortInt, err := strconv.Atoi(dbPort)
	if err == nil {
		cluster.Port = dbPortInt
	}

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	if useSSL {
		cluster.SslOpts = &gocql.SslOptions{
			CaPath: "./sf-class2-root.crt",
		}
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to Rewards Cassandra DB")
	Cassandra = session
}

func ConnectPostgres(dbConnString string) {
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to Profiler Postgres")
	Postgres = db
}

func CreateSESSession() {
	// Create a new session in the ap-southeast-1 region.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	// Create an SES session.
	svc = ses.New(sess)
}

func GetTodaysRewards() (rewards []Reward, _ error) {
	todaysDate := time.Now().Format(YYYYMMDD)

	todaysDateLiteral := fmt.Sprintf("'%s'", todaysDate)

	//TODO: Replace todaysDate with a date in the csv data from ftp server, since no rewards with todaysdate

	// Equivalent to the following query
	// select * from transactions.rewards where transaction_date == {today's date} and reward_amount > 0 ALLOW FILTERING;
	stmt, _ := qb.Select("transactions.rewards").Where(qb.EqLit("transaction_date", todaysDateLiteral)).Where(qb.GtLit("reward_amount", "0")).AllowFiltering().ToCql()

	err := gocqlx.Select(&rewards, Cassandra.Query(stmt))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rewards, nil
}

func GetUniqueCardIds(rewards []Reward) []uuid.UUID {
	// Create a map to store the unique card IDs
	uniqueCardIds := make(map[uuid.UUID]bool)

	// Iterate over the rewards and add the card IDs to the map
	for _, reward := range rewards {
		card_uuid := uuid.MustParse(reward.CardID.String())
		uniqueCardIds[card_uuid] = true
	}

	// Convert the map keys to an array
	cardIds := make([]uuid.UUID, 0, len(uniqueCardIds))
	for cardId := range uniqueCardIds {
		cardIds = append(cardIds, cardId)
	}

	return cardIds
}

func GetRewardsByEmailAndCardID(rewards []Reward, cardIDs []uuid.UUID) (map[string]map[uuid.UUID][]Reward, error) {
	// Fetch user emails and card IDs associated with the provided card IDs
	type CardUserDetails struct {
		CardID uuid.UUID
		UserID uuid.UUID
		Email  string
	}
	var cardUserDetails []CardUserDetails
	if err := Postgres.Table("cards").
		Select("cards.id AS card_id, users.id AS user_id, users.email").
		Joins("JOIN users ON users.id = cards.user_id").
		Where("cards.id IN (?)", cardIDs).
		Scan(&cardUserDetails).Error; err != nil {
		return nil, err
	}

	emailByCardId := make(map[uuid.UUID]string)

	for _, cardUser := range cardUserDetails {
		emailByCardId[cardUser.CardID] = cardUser.Email
	}

	rewardsByEmailAndCardID := make(map[string]map[uuid.UUID][]Reward)

	// Assemble a map
	// Where the key is a user's email and the value is a map
	// That map has the keys of cardIds associated with that user, and an array of rewards for the day
	// All content that appears on this map are rewards that have been applied today and have a reward_amount > 0
	for _, reward := range rewards {
		rewardCardId := uuid.MustParse(reward.CardID.String())
		email := emailByCardId[rewardCardId]

		if rewardsByEmailAndCardID[email] == nil {
			rewardsByEmailAndCardID[email] = make(map[uuid.UUID][]Reward)
			rewardsByEmailAndCardID[email][rewardCardId] = []Reward{reward}
		} else {
			if rewardsByEmailAndCardID[email][rewardCardId] == nil {
				rewardsByEmailAndCardID[email][rewardCardId] = []Reward{reward}
			} else {
				rewardsByEmailAndCardID[email][rewardCardId] = append(rewardsByEmailAndCardID[email][rewardCardId], reward)
			}
		}
	}

	return rewardsByEmailAndCardID, nil
}

func SendEmail(recipient string, cardRewards map[uuid.UUID][]Reward) error {
	// Assemble the body of the email
	// TODO: Replace with this with a pretty template
	htmlBody := ""

	for cardId, rewards := range cardRewards {
		htmlBody += "Reward For Card ID: " + cardId.String()
		for _, reward := range rewards {
			rewardStr, err := json.Marshal(reward)
			if err != nil {
				fmt.Println("JSON parse reward failed")
			}
			htmlBody += string(rewardStr) + "\n"
		}
		htmlBody += "\n"
	}

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(htmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Your AngelOwl rewards accumulated today"),
			},
		},
		Source: aws.String("noreply@itsag1t2.com"),
		// Uncomment to use a configuration set
		// ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return err
	}

	fmt.Println("Email Sent to address: " + recipient)
	fmt.Println(result)
	return nil
}

func HandleRequest(ctx context.Context, event S3Event) (string, error) {
	os.Getenv("CASSANDRA_CONN_STRING")
	dbHost := os.Getenv("CASSANDRA_CONN_STRING")
	dbPort := os.Getenv("CASSANDRA_PORT")
	dbKeyspace := os.Getenv("CASSANDRA_PORT")
	dbUser := os.Getenv("CASSANDRA_USERNAME")
	dbPass := os.Getenv("CASSANDRA_PASSWORD")
	dbUseSSL := true
	ConnectCassandra(dbHost, dbPort, dbUser, dbPass, dbKeyspace, dbUseSSL)

	dbConnString := os.Getenv("POSTGRES_CONN_STRING")
	ConnectPostgres(dbConnString)

	CreateSESSession()

	rewards, err := GetTodaysRewards()

	cards := GetUniqueCardIds(rewards)

	mailMap, err := GetRewardsByEmailAndCardID(rewards, cards)

	for email, cardRewards := range mailMap {
		err = SendEmail(email, cardRewards)
	}

	return "", err
}

func main() {
	lambda.Start(HandleRequest)
}

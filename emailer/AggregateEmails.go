package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
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
	CreatedAt       string     `json:"created_at"`       // cassandra text
	CardPAN         string     `json:"card_pan"`         // cassandra text
	CardType        string     `json:"card_type"`        // cassandra text
	RewardAmount    float64    `json:"reward_amount"`    // cassandra double
	Remarks         string     `json:"remarks"`          // cassandra text
}

type ParseReward struct {
	Rewards []Reward `json:"data"`
}

type Card struct {
	ID               uuid.UUID `json:"id"`
	CardPan          string    `json:"card_pan"`
	UserID           uuid.UUID `json:"user_id"`   // card belongs to one user
	CardTypeCardType string    `json:"card_type"` // card belongs to one card type
}

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        string
	CreditCards []Card // one user has many credit cards
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// var Cassandra *gocql.Session
var informerUrl string
var Postgres *gorm.DB
var svc *ses.SES

const YYYYMMDD = "2006-01-02"

// If any of the connections fail, lambda won't run
func init() {
	informerUrl = os.Getenv("INFORMER_ENDPOINT")
	if informerUrl == "" {
		log.Fatalln("INFORMER_ENDPOINT environment variable is not set")
	}
	dbConnString := os.Getenv("POSTGRES_CONN_STRING")
	ConnectPostgres(dbConnString)

	CreateSESSession()
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
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	// Create an SES session.
	svc = ses.New(sess)
}

func GetTodaysRewards() ([]Reward, error) {
	resp, err := http.Get(informerUrl + "/reward/today")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch rewards")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)

	var parseReward ParseReward

	err = json.NewDecoder(bodyBuffer).Decode(&parseReward)
	if err != nil {
		return nil, err
	}

	_ = resp.Body.Close()
	return parseReward.Rewards, nil
}

func GetUniqueCardIds(rewards []Reward) []uuid.UUID {
	// Create a map to store the unique card IDs
	uniqueCardIds := make(map[uuid.UUID]bool)

	// Iterate over the rewards and add the card IDs to the map
	for _, reward := range rewards {
		cardUuid := uuid.MustParse(reward.CardID.String())
		uniqueCardIds[cardUuid] = true
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
	htmlBytes, err := os.ReadFile("/var/task/template.html")
	if err != nil {
		fmt.Println("Error reading HTML template file:", err)
		return err
	}

	var allRewards []Reward

	for _, rewards := range cardRewards {
		allRewards = append(allRewards, rewards...)
	}
	tmpl, err := template.New("email").Parse(string(htmlBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}
	data := struct {
		Email   string
		Rewards []Reward
	}{
		Email:   recipient,
		Rewards: allRewards,
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return err
	}

	// Assemble the body of the email
	htmlBody := buf.String()

	emailWhiteList := map[string]bool{
		"omerbaggia123@gmail.com": true,
		"alvinowyong@gmail.com":   true,
		"nicholasongck@gmail.com": true,
		"wenlianggg@gmail.com":    true,
		"lyejianyi@gmail.com":     true,
		"awkhailoong@gmail.com":   true,
		"justin.100600@gmail.com": true,
	}

	if emailWhiteList[recipient] {
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
		}

		// Attempt to send the email.
		_, err = svc.SendEmail(input)

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

	}

	return nil
}

type Event struct{}

func HandleRequest(ctx context.Context, placeholder Event) (string, error) {
	rewards, err := GetTodaysRewards()
	if err != nil {
		log.Println("error getting today's rewards: ", err)
	}

	cards := GetUniqueCardIds(rewards)

	mailMap, err := GetRewardsByEmailAndCardID(rewards, cards)
	if err != nil {
		log.Println("error getting emails: ", err)
	}

	for email, cardRewards := range mailMap {
		err = SendEmail(email, cardRewards)
		if err != nil {
			log.Println("error sending email to "+email, err)
		}
	}

	return "", err
}

func main() {
	lambda.Start(HandleRequest)
}

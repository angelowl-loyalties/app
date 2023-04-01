package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
)

var S3Svc *s3.S3
var profilerUrl string
var userMap = make(map[string]User)
var cardMap = make(map[string]Card)

func init() {
	// Establish an AWS session
	sess := session.Must(session.NewSession())
	S3Svc = s3.New(sess)

	profilerUrl = os.Getenv("PROFILER_ENDPOINT")
	if profilerUrl == "" {
		log.Fatalln("PROFILER_ENDPOINT environment variable is not set")
	}
}

func parse(csvRecord []string) (record CSVRecord, err error) {
	record.ID, err = uuid.Parse(csvRecord[0])
	if err != nil {
		return record, err
	}

	record.FirstName = csvRecord[1]
	record.LastName = csvRecord[2]
	record.Phone = csvRecord[3]
	record.Email = csvRecord[4]
	// 5 created_at
	// 6 updated_at

	record.CardID, err = uuid.Parse(csvRecord[7])
	if err != nil {
		return record, err
	}

	record.CardPAN = strings.ReplaceAll(csvRecord[8], "-", "")
	record.CardType = csvRecord[9]

	return record, nil
}

func postToProfiler() (result OTPEvent, hasError bool) {
	for userID, user := range userMap {
		userJson, err := json.Marshal(user)
		if err != nil {
			hasError = true
			log.Println("failed to marshal JSON for user: ", userID)
			continue
		}

		resp, err := http.Post(profilerUrl+"/user", "application/json", bytes.NewBuffer(userJson))
		if err != nil {
			hasError = true
			log.Println("failed to create user: " + userID + " with error: " + err.Error())
			continue
		}
		if resp.StatusCode != 201 {
			hasError = true
			log.Printf("failed to create user: "+userID+" with error: %d", resp.StatusCode)
			continue
		}

		result.Users = append(result.Users, EmailContent{
			Email:    user.Email,
			Name:     user.FirstName + user.LastName,
			Password: user.Password,
		})
	}

	for cardID, card := range cardMap {
		cardJson, err := json.Marshal(card)
		if err != nil {
			hasError = true
			log.Println("failed to marshal JSON for card: ", cardID)
			continue
		}

		resp, err := http.Post(profilerUrl+"/card", "application/json", bytes.NewBuffer(cardJson))
		if err != nil {
			hasError = true
			log.Println("failed to create card: " + cardID + " with error: " + err.Error())
			continue
		}
		if resp.StatusCode != 201 {
			hasError = true
			log.Printf("failed to create card: "+cardID+" with error: %d", resp.StatusCode)
			//body, _ := ioutil.ReadAll(resp.Body)
			//resp.Body.Close()
			//log.Printf(string(body))
			continue
		}
	}

	return result, hasError
}

func generateTempPassword() string {
	res, _ := password.Generate(12, 4, 4, false, false)
	return res
}

func handleRequest(ctx context.Context, event events.S3Event) (OTPEvent, error) {
	// Define the S3 bucket and file key
	//bucket := "angel-owl-profiler-pet-rock"
	//fileKey := "test.csv"
	bucket := event.Records[0].S3.Bucket.Name
	fileKey := event.Records[0].S3.Object.Key

	// Download the file from S3
	s3Object, err := S3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		log.Fatalln("error getting object from S3: ", err)
	}
	log.Println("processing " + fileKey + " from " + bucket)

	// Read the contents of the file
	reader := csv.NewReader(bufio.NewReader(s3Object.Body))
	_, err = reader.Read()
	if err != nil {
		log.Fatalln("error getting reader: ", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("error reading record from .csv file: ", err)
			continue
		}

		var s3CSV CSVRecord
		s3CSV, err = parse(record)
		if err != nil {
			log.Println("error parsing record: ", err)
			continue
		}

		tempPass := generateTempPassword()
		user := User{
			ID:              s3CSV.ID,
			FirstName:       s3CSV.FirstName,
			LastName:        s3CSV.LastName,
			Phone:           s3CSV.Phone,
			Email:           s3CSV.Email,
			Password:        tempPass,
			ConfirmPassword: tempPass,
		}
		// new user, add to map
		if _, found := userMap[user.ID.String()]; !found {
			userMap[user.ID.String()] = user
		}

		card := Card{
			ID:       s3CSV.CardID,
			CardPAN:  s3CSV.CardPAN,
			UserID:   s3CSV.ID,
			CardType: s3CSV.CardType,
		}
		// new card, add to map
		if _, found := cardMap[card.ID.String()]; !found {
			cardMap[card.ID.String()] = card
		}
	}

	result, hasError := postToProfiler()
	log.Printf("users processed: %d, cards processed: %d, errors: %t\n", len(userMap), len(cardMap), hasError)

	return result, nil
}

func main() {
	lambda.Start(handleRequest)
}

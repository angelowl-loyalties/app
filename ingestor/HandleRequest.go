package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

const (
	DDMMYY   = "02-01-06"
	YYYYMMDD = "2006-01-02"
)

var producer sarama.SyncProducer
var s3Svc *s3.S3

var brokers = []string{"angelowlmsk.aznt6t.c3.kafka.ap-southeast-1.amazonaws.com:9092"}

type Transaction struct {
	ID              uuid.UUID `json:"id"`
	CardID          uuid.UUID `json:"card_id"`
	Merchant        string    `json:"merchant"`
	MCC             int       `json:"mcc"`
	Currency        string    `json:"currency"`
	Amount          float64   `json:"amount"`
	SGDAmount       float64   `json:"sgd_amount"`
	TransactionID   string    `json:"transaction_id"`
	TransactionDate string    `json:"transaction_date"`
	CardPAN         string    `json:"card_pan"`
	CardType        string    `json:"card_type"`
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

func init() {
	// Establish an AWS session
	var err error
	sess := session.Must(session.NewSession())
	s3Svc = s3.New(sess)

	producer, err = newProducer()
	if err != nil {
		fmt.Printf("Failed to create producer: %v\n", err)
		os.Exit(1)
	}
}

type ParseErrorInterface struct {
	message string
}

func (e *ParseErrorInterface) Error() string {
	return e.message
}

func ParseError(message string) error {
	return &ParseErrorInterface{message: message}
}

// Follows FTP ordering
func (transaction *Transaction) Parse(transactionCsv []string) (err error) {
	transaction.ID, err = uuid.Parse(transactionCsv[0])
	if err != nil {
		fmt.Printf("Error parsing uuid: %v\n", err)
		fmt.Println(transactionCsv[0])
		return err
	}

	transaction.TransactionID = transactionCsv[1]
	transaction.Merchant = transactionCsv[2]

	if transactionCsv[3] == "" {
		transactionCsv[3] = "0"
	}

	transaction.MCC, err = strconv.Atoi(transactionCsv[3])
	if err != nil {
		fmt.Printf("Error in ATOI: %v\n", err)
		fmt.Println(transactionCsv[3])
		return err
	}

	transaction.Currency = transactionCsv[4]
	transaction.Amount, err = strconv.ParseFloat(transactionCsv[5], 64)
	if err != nil {
		fmt.Printf("Error in parsing float: %v\n", err)
		fmt.Println(transactionCsv[5])
		return err
	}
	transaction.SGDAmount = 0.0
	transaction.TransactionDate = transactionCsv[6]

	transaction.CardID, err = uuid.Parse(transactionCsv[7])
	if err != nil {
		fmt.Printf("Error parsing uuid: %v\n", err)
		fmt.Println(transactionCsv[7])
		return err
	}

	if len(transactionCsv[8]) < 13 || len(transactionCsv[8]) > 19 {
		return ParseError("Card Pan failed to be validated")
	}
	transaction.CardPAN = transactionCsv[8]

	transaction.CardType = transactionCsv[9]
	return nil
}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	// config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Transaction.Retry.Backoff = 10
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(message []byte) (msg *sarama.ProducerMessage, err error) {
	_, err = sarama.ByteEncoder.Encode(message)

	if err != nil {
		return nil, err
	}

	msg = &sarama.ProducerMessage{
		Topic: "transaction",
		Value: sarama.ByteEncoder(message),
	}

	return msg, nil
}

func HandleRequest(ctx context.Context, event S3Event) (string, error) {
	// Define the S3 bucket and file key
	bucket := "angel-owl-spendfiles"
	fileKey := event.Records[0].S3.Object.Key

	// Download the file from S3
	result, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		fmt.Printf("Error getting object from S3: %v\n", err)
		os.Exit(1)
	}

	// Read the contents of the file
	reader := csv.NewReader(bufio.NewReader(result.Body))

	_, err = reader.Read()

	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			fmt.Printf("Reader EOF: %v", err)
			break
		}
		if err != nil {
			fmt.Printf("Error reading record from .csv file: %v", err)
			continue
		}

		var m Transaction
		err = m.Parse(record)
		if err != nil {
			fmt.Printf("Error parsing transaction: %v", err)
			continue
		}

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v", err)
			continue
		}

		// go func() {
		message, err := prepareMessage(b)

		if err != nil {
			fmt.Printf("Error preparing message to Kafka: %v", err)
		}

		_, _, err = producer.SendMessage(message)
		if err != nil {
			fmt.Printf("Error writing to Producer: %v", err)
		}
		// }()
	}

	return "", err
}

func main() {
	lambda.Start(HandleRequest)
	_ = producer.Close()
}

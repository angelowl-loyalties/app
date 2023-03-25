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

// Follows FTP ordering
func (transaction *Transaction) Parse(transactionCsv []string) error {
	transaction.ID = uuid.MustParse(transactionCsv[0])
	transaction.TransactionID = transactionCsv[1]
	transaction.Merchant = transactionCsv[2]
	transaction.MCC, _ = strconv.Atoi(transactionCsv[3])
	transaction.Currency = transactionCsv[4]
	transaction.Amount, _ = strconv.ParseFloat(transactionCsv[5], 64)
	transaction.SGDAmount = 0.0
	transaction.TransactionDate = transactionCsv[6]
	transaction.CardID = uuid.MustParse(transactionCsv[7])
	transaction.CardPAN = transactionCsv[8]
	transaction.CardType = transactionCsv[9]

	return nil
}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Transaction.Retry.Backoff = 10
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(message []byte) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic: "transaction",
		Value: sarama.ByteEncoder(message),
	}

	return msg
}

func HandleRequest(ctx context.Context, event S3Event) (string, error) {
	// Establish an AWS session
	sess := session.Must(session.NewSession())
	s3Svc := s3.New(sess)

	// Define the S3 bucket and file key
	bucket := "angel-owl-spendfiles"
	fileKey := "3fc97862-a9c7-4c67-bd55-8b83f5832342.csv"

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

	// producer := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers:      []string{"angelowlmsk.aznt6t.c3.kafka.ap-southeast-1.amazonaws.com:9092"},
	// 	Topic:        "transaction",
	// 	Balancer:     &kafka.LeastBytes{},
	// 	BatchSize:    1000,
	// 	BatchTimeout: 100 * time.Millisecond,
	// })
	producer, err := newProducer()
	if err != nil {
		fmt.Printf("Failed to create producer: %v\n", err)
		os.Exit(1)
	}
	defer producer.Close()

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading record from .csv file: %v", err)
		}

		var m Transaction
		err = m.Parse(record)
		if err != nil {
			fmt.Printf("Error parsing transaction: %v", err)
		}

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v", err)
		}

		go func() {
			// err := producer.WriteMessages(ctx, kafka.Message{
			// 	Key:   []byte("transaction6"),
			// 	Value: []byte(b),
			// })
			_, _, err := producer.SendMessage(prepareMessage(b))
			if err != nil {
				fmt.Printf("Error writing to Producer: %v", err)
			}
		}()
	}

	return "", err
}

func main() {
	lambda.Start(HandleRequest)
}

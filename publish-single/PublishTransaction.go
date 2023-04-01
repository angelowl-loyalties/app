package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

const (
	DDMMYY   = "02-01-06"
	YYYYMMDD = "2006-01-02"
)

var producer sarama.SyncProducer

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

type PublishEvent struct {
	ID              string  `json:"id"`
	CardID          string  `json:"card_id"`
	Merchant        string  `json:"merchant"`
	MCC             int     `json:"mcc"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	SGDAmount       float64 `json:"sgd_amount"`
	TransactionID   string  `json:"transaction_id"`
	TransactionDate string  `json:"transaction_date"`
	CardPAN         string  `json:"card_pan"`
	CardType        string  `json:"card_type"`
}

func init() {
	// Create producer
	var err error

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

// Basically marshals the transaction JSON but validates PAN, and MCC
func (transaction *Transaction) Parse(transactionJson PublishEvent) (err error) {
	transaction.ID, err = uuid.Parse(transactionJson.ID)
	if err != nil {
		fmt.Printf("Error parsing uuid: %v\n", err)
		fmt.Println(transactionJson.ID)
		return err
	}

	transaction.TransactionID = transactionJson.TransactionID
	transaction.Merchant = transactionJson.Merchant

	transaction.MCC = transactionJson.MCC

	transaction.Currency = transactionJson.Currency
	transaction.Amount = transactionJson.Amount

	transaction.SGDAmount = 0.0
	transaction.TransactionDate = transactionJson.TransactionDate

	transaction.CardID, err = uuid.Parse(transactionJson.CardID)
	if err != nil {
		fmt.Printf("Error parsing uuid: %v\n", err)
		fmt.Println(transactionJson.CardID)
		return err
	}

	if len(transactionJson.CardPAN) < 13 || len(transactionJson.CardPAN) > 19 {
		return ParseError("Card Pan failed to be validated")
	}
	transaction.CardPAN = transactionJson.CardPAN

	transaction.CardType = transactionJson.CardType
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

func HandleRequest(ctx context.Context, event PublishEvent) (string, error) {
	var m Transaction
	err := m.Parse(event)
	if err != nil {
		fmt.Printf("Error parsing transaction: %v", err)
		return "", err
	}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v", err)
		return "", err
	}

	message, err := prepareMessage(b)
	if err != nil {
		log.Printf("Error preparing message to Kafka: %v", err)
		return "", err
	}

	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Printf("Error writing to Producer: %v", err)
		return "", err
	}

	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
	_ = producer.Close()
}

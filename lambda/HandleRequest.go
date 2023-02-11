package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/segmentio/kafka-go"
)

type S3Event struct {
	Records []struct {
		S3 struct {
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func HandleRequest(ctx context.Context, event S3Event) (string, error) {
	// Establish an AWS session
	sess := session.Must(session.NewSession())
	s3Svc := s3.New(sess)

	// Define the S3 bucket and file key
	bucket := "angel-owl-spendfiles"
	// fileKey := event.Records[0].S3.Object.Key
	fileKey := "01d6779e-a727-41fa-81ca-9a198061ec45.csv"

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

	// Read the header line
	header, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}

	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"angelowlmsk.aznt6t.c3.kafka.ap-southeast-1.amazonaws.com:9092"},
		Topic:   "transaction6",
		Balancer: &kafka.LeastBytes{},
		BatchSize: 1000,
		BatchTimeout: 100 * time.Millisecond,
	})
	defer producer.Close()

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading record from .csv file: %v", err)
		}

		m := map[string]string{}
		for i, h := range header {
			m[h] = record[i]
		}

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v", err)
		}

		go producer.WriteMessages(ctx, kafka.Message{
			Key:   []byte("transaction6"),
			Value: []byte(b),
		})
		if err != nil {
			fmt.Printf("Error writing to Producer: %v", err)
		}
	}

	return "", err
}

func main() {
	lambda.Start(HandleRequest)
}

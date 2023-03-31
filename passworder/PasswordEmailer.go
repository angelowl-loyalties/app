package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type OTPEvent struct {
	Users []struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"users"`
}

var svc *ses.SES

// If SES Session creation fails, lambda wont run
func init() {
	CreateSESSession()
}

func CreateSESSession() {
	// Create a new session in the ap-southeast-1 region.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	// Create an SES session.
	svc = ses.New(sess)
}

func SendEmail(recipient string, name string, password string) error {
	// Read the HTML template file.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}
	htmlBytes, err := ioutil.ReadFile(filepath.Join(dir, "template.html"))
	if err != nil {
		fmt.Println("Error reading HTML template file:", err)
		return err
	}

	// Parse the template.
	tmpl, err := template.New("email").Parse(string(htmlBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	// Define the data to be rendered in the template.
	data := struct {
		Password string
	}{
		Password: password,
	}

	// Render the template with the data.
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return err
	}

	// Assemble the body of the email
	htmlBody := buf.String()

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
				Data:    aws.String("Welcome to AngelOwl, " + name + "!"),
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

	return nil
}

func HandleRequest(ctx context.Context, event OTPEvent) (string, error) {
	for _, detail := range event.Users {
		err := SendEmail(detail.Email, detail.Name, detail.Password)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}

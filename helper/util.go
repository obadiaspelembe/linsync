package helper

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetCredentials() *credentials.Credentials {

	aws_access_key_id := os.Getenv("LINODE_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("LINODE_SECRET_ACCESS_KEY")

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

	fmt.Println("Credentials: Generating credentials")
	return creds
}

func SessionInitialization() *session.Session {

	region := os.Getenv("LINODE_DEFAULT_REGION")

	endpoint := "https://" + region + ".linodeobjects.com"

	creds := GetCredentials()

	session, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: creds,
			Endpoint:    aws.String(endpoint),
		})

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println("Session: Initializing session")

	return session
}

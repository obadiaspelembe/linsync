package helper

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func PullToBucket(uploadFileDir string)  {

	bucletName := os.Getenv("LINODE_BUCKET_NAME")

	session := SessionInitialization()

	log.Println("Pull: Downloading file to bucket")

	resp, err := s3.New(session).GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(bucletName),
			Key:    aws.String(uploadFileDir),
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	log.Println("Pull: Downloaded Successfully")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	log.Println("Pull: Writing file to System directory...")
    f, err := os.Create(uploadFileDir)
    if err != nil {
		log.Print(err)
	}

    defer f.Close()

    _, err = f.Write(body)

	if err != nil {
		log.Fatal(err.Error())
	}

    f.Sync()

    w := bufio.NewWriter(f)

    w.Flush()

	log.Println("Pull: Success.")

}

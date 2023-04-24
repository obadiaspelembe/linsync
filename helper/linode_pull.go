package helper

import (
	"bufio"
	"fmt"
	"io/ioutil" 
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func PullToBucket(uploadFileDir string)  {

	bucletName := os.Getenv("LINODE_BUCKET_NAME")

	session := SessionInitialization()

	fmt.Println("Pull: Downloading file to bucket")

	resp, err := s3.New(session).GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(bucletName),
			Key:    aws.String(uploadFileDir),
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	fmt.Println("Pull: Downloaded Successfully")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Pull: Writing file to System directory")

    f, err := os.Create(uploadFileDir)
    if err != nil {
		fmt.Print(err)
	}

    defer f.Close()

    _, err = f.Write(body)

	if err != nil {
		fmt.Println(err.Error())
	}

    f.Sync()

    w := bufio.NewWriter(f)

    w.Flush()

	fmt.Println("Pull: Success.")

}

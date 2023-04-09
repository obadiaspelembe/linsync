package helper

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)



func PushToBucket(uploadFileDir string) error {

    bucletName := os.Getenv("LINODE_BUCKET_NAME")

    session := SessionInitialization()

    upFile, err := os.Open(uploadFileDir)
    if err != nil {
        return err
    }
    defer upFile.Close()
    
    upFileInfo, _ := upFile.Stat()
    var fileSize int64 = upFileInfo.Size()
    fileBuffer := make([]byte, fileSize)
    upFile.Read(fileBuffer) 

    log.Println("Push: Uploading file to bucket")
    
    _, err = s3.New(session).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(bucletName),
        Key:                  aws.String(uploadFileDir),
        ACL:                  aws.String("private"),
        Body:                 bytes.NewReader(fileBuffer),
        ContentLength:        aws.Int64(fileSize),
        ContentType:          aws.String(http.DetectContentType(fileBuffer)),
        ContentDisposition:   aws.String("attachment"),
        ServerSideEncryption: aws.String("AES256"),
    })

	if err!= nil {
		fmt.Println(err.Error())
    } else {
        log.Println("Push: Successfully pushed")
    }
    return err
}

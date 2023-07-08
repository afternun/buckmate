package upload

import (
	"buckmate/main/common/exception"
	"buckmate/structs"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3 (paths string[], bucket string) {
	sess, err := session.Must(session.NewSession())
	exception.Handle(structs.Exception{Err: err, Message: "Couldn't create AWS session."})

	uploader := s3manager.NewUploader(sess)

	for _, path range paths {
		file, err := os.OpenFile(path)
		exception.Handle(structs.Exception{Err: err, Message: "Couldn't open file for upload."})
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket)
			Key: aws.String(path)
			Body: file
		})
	}


}
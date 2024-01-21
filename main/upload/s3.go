package upload

import (
	"buckmate/main/common/exception"
	"buckmate/structs"
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	LocalDirectory = "build"
)

func S3(bucket string, prefix string) {
	walker := make(fileWalk)
	go func() {
		// Gather the files to upload by walking the path recursively
		if err := filepath.Walk(LocalDirectory, walker.Walk); err != nil {
			log.Fatalln("Walk failed:", err)
		}
		close(walker)
	}()

	cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
	exception.Handle(structs.Exception{Err: cfgErr, Message: "Couldn't get config."})

	// For each file found walking, upload it to Amazon S3
	uploader := manager.NewUploader(s3.NewFromConfig(cfg))
	for path := range walker {
		rel, err := filepath.Rel(LocalDirectory, path)
		exception.Handle(structs.Exception{Err: err, Message: "Couldn't get relative path."})

		file, err := os.Open(path)
		exception.Handle(structs.Exception{Err: err, Message: "Failed to open file"})

		defer file.Close()
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucket,
			Key:    aws.String(filepath.Join(prefix, rel)),
			Body:   file,
		})
		exception.Handle(structs.Exception{Err: err, Message: "Failed to upload"})
		log.Println("Uploaded", path, result.Location)
	}
}

type fileWalk chan string

func (f fileWalk) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}
	return nil
}

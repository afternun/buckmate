package download

import (
	"buckmate/main/common/exception"
	"buckmate/structs"
	"context"
	"fmt"
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
	cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
	exception.Handle(structs.Exception{Err: cfgErr, Message: "Couldn't get config."})

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "eu-central-1"
	})
	manager := manager.NewDownloader(client)

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket:     &bucket,
		Prefix:     &prefix,
		StartAfter: &prefix,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalln("error:", err)
		}
		for _, obj := range page.Contents {
			fmt.Printf("Current obj %s \n", aws.ToString(obj.Key))
			if obj.Size > 0 {
				if err := downloadToFile(manager, LocalDirectory, bucket, aws.ToString(obj.Key)); err != nil {
					log.Fatalln("error:", err)
				}
			}
		}
	}
}

func downloadToFile(downloader *manager.Downloader, targetDirectory, bucket, key string) error {
	// Create the directories in the path
	file := filepath.Join(targetDirectory, key)
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		return err
	}

	// Set up the local file
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	// Download the file using the AWS SDK for Go
	fmt.Printf("Downloading s3://%s/%s to %s...\n", bucket, key, file)
	_, err = downloader.Download(context.TODO(), fd, &s3.GetObjectInput{Bucket: &bucket, Key: &key})
	return err
}

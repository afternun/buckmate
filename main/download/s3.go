package download

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3(bucket string, prefix string, tempDir string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	manager := manager.NewDownloader(client)

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket:     &bucket,
		Prefix:     &prefix,
		StartAfter: &prefix,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, obj := range page.Contents {
			fmt.Printf("Current obj %s \n", aws.ToString(obj.Key))
			if obj.Size > 0 {
				downloadToFile(manager, tempDir, bucket, aws.ToString(obj.Key), prefix)
			}
		}
	}
}

func downloadToFile(downloader *manager.Downloader, targetDirectory, bucket, key string, prefix string) {
	file := filepath.Clean(strings.Replace(filepath.Join(targetDirectory, key), prefix, "", 1))

	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		log.Fatal(err)
	}

	fd, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	fmt.Printf("Downloading s3://%s/%s to %s...\n", bucket, key, file)
	if _, err := downloader.Download(context.TODO(), fd, &s3.GetObjectInput{Bucket: &bucket, Key: &key}); err != nil {
		log.Fatal(err)
	}
}

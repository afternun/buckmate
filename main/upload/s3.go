package upload

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	LocalDirectory = "build"
)

func S3(bucket string, prefix string, version string) {
	walker := make(fileWalk)
	go func() {
		// Gather the files to upload by walking the path recursively
		if err := filepath.Walk(LocalDirectory, walker.Walk); err != nil {
			log.Fatalln("Walk failed:", err)
		}
		close(walker)
	}()

	cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
	if cfgErr != nil {
		log.Fatalln("Could not get AWS config")
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "eu-central-1"
	})

	headObjClient := s3.HeadObjectAPIClient(client)
	removeObjClient := manager.DeleteObjectsAPIClient(client)
	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})

	// For each file found walking, upload it to Amazon S3
	uploader := manager.NewUploader(s3.NewFromConfig(cfg))
	metadata := map[string]string{"buckmate-version": version}
	for path := range walker {
		rel, err := filepath.Rel(LocalDirectory, path)
		if err != nil {
			log.Fatalln("Could not get relative path to file " + path)
		}

		file, err := os.Open(path)
		if err != nil {
			log.Fatalln("Failed to open file " + path)
		}

		defer file.Close()
		fileKey := aws.String(filepath.Join(prefix, rel))
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket:   &bucket,
			Key:      fileKey,
			Body:     file,
			Metadata: metadata,
		})
		if err != nil {
			log.Fatalln("Failed to upload file " + *fileKey)
		}
		log.Println("Uploaded", path, result.Location)
	}
	if err := os.RemoveAll(LocalDirectory); err != nil {
		log.Fatalln("Failed to remove temporary build directory")
	}
	removeVersion(bucket, paginator, headObjClient, removeObjClient, version)
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

func removeVersion(bucket string, paginator *s3.ListObjectsV2Paginator, headObjClient s3.HeadObjectAPIClient, removeObjClient manager.DeleteObjectsAPIClient, version string) {
	var objectsToRemove []types.ObjectIdentifier
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalln("error:", err)
		}
		for _, obj := range page.Contents {
			fmt.Printf("Current obj %s \n", aws.ToString(obj.Key))
			if obj.Size > 0 {
				header, err := headObjClient.HeadObject(context.TODO(), &s3.HeadObjectInput{
					Bucket: &bucket,
					Key:    obj.Key,
				})
				if err != nil {
					log.Fatalln("error:", err)
				}
				if header.Metadata["buckmate-version"] != version {
					objectsToRemove = append(objectsToRemove, types.ObjectIdentifier{Key: obj.Key})
					fmt.Printf("Will Remove: %s\n", *obj.Key)
				}
			}
		}
	}
	fmt.Printf("To Delete %v", *objectsToRemove[0].Key)
	_, err := removeObjClient.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: &bucket,
		Delete: &types.Delete{
			Objects: objectsToRemove,
		},
	})
	if err != nil {
		log.Fatalln("error during rm:", err)
	}
}

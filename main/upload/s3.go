package upload

import (
	"context"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func S3(bucket string, prefix string, version string, tempDir string) {
	walker := make(fileWalk)
	go func() {
		if err := filepath.Walk(tempDir, walker.Walk); err != nil {
			log.Fatal("Walk failed:", err)
		}
		close(walker)
	}()

	cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
	if cfgErr != nil {
		log.Fatal("Could not get AWS config")
	}

	client := s3.NewFromConfig(cfg)

	headObjClient := s3.HeadObjectAPIClient(client)
	removeObjClient := manager.DeleteObjectsAPIClient(client)
	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})

	uploader := manager.NewUploader(s3.NewFromConfig(cfg))
	metadata := map[string]string{"buckmate-version": version}
	for path := range walker {
		rel, err := filepath.Rel(tempDir, path)
		if err != nil {
			log.Fatal("Could not get relative path to file " + path)
		}

		file, err := os.Open(path)
		if err != nil {
			log.Fatal("Failed to open file " + path)
		}

		fileKey := aws.String(filepath.Join(prefix, rel))
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket:   &bucket,
			Key:      fileKey,
			Body:     file,
			Metadata: metadata,
		})
		file.Close()
		if err != nil {
			log.Fatal("Failed to upload file " + *fileKey)
		}
		log.Debug("Uploaded", path, result.Location)
	}
	if err := os.RemoveAll(tempDir); err != nil {
		log.Fatal("Failed to remove temporary build directory")
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
			log.Fatal("Failed to get the next page")
		}
		for _, obj := range page.Contents {
			log.Debug("Heading object for removal: %s \n", aws.ToString(obj.Key))
			if obj.Size > 0 {
				header, err := headObjClient.HeadObject(context.TODO(), &s3.HeadObjectInput{
					Bucket: &bucket,
					Key:    obj.Key,
				})
				if err != nil {
					log.Fatal("Failed to head the object: " + *obj.Key)
				}
				if header.Metadata["buckmate-version"] != version {
					objectsToRemove = append(objectsToRemove, types.ObjectIdentifier{Key: obj.Key})
					log.Debug("Adding to removal list: %s\n", aws.ToString(obj.Key))
				}
			}
		}
	}
	log.Debug("Removing objects: %v", objectsToRemove)
	_, err := removeObjClient.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: &bucket,
		Delete: &types.Delete{
			Objects: objectsToRemove,
		},
	})
	if err != nil {
		log.Fatal("error during rm:", err)
	}
}

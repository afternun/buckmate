package cmd

import (
	"buckmate/main/aws"
	"buckmate/main/common/util"
	"buckmate/main/deploymentConfig"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Applies deployment to the infrastructure",
	Long: `Use:
buckmate apply
	`,
	Run: func(cmd *cobra.Command, args []string) {
		env, err := cmd.Flags().GetString("env")
		if err != nil {
			log.Fatal(err)
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}

		workDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		rootDir := util.Resolve(workDir, path)

		s3Prefix := "s3://"
		tempDir, err := util.RandomDirectory()
		if err != nil {
			log.Fatal(err)
		}
		dConfig, err := deploymentConfig.Load(env, rootDir)
		if err != nil {
			log.Fatal(err)
		}
		buckmateVersion := uuid.New().String()

		client, err := aws.Init()
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(dConfig.Source.Address, s3Prefix) {
			dConfig.Source.Address = strings.Replace(dConfig.Source.Address, s3Prefix, "", 1)
			sourceBucket := aws.NewBucket(client, dConfig.Source)
			downloadOptions := aws.DownloadOptions{
				Prefix:  dConfig.Source.Prefix,
				TempDir: tempDir,
			}
			err := sourceBucket.Download(cmd.Context(), downloadOptions)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			dConfig.Source.Address = util.Resolve(rootDir, dConfig.Source.Address)
			err := util.CopyDirectory(dConfig.Source.Address, tempDir)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = util.CopyDirectory(rootDir+"/files", tempDir)
		if err != nil {
			log.Fatal(err)
		}

		err = util.CopyDirectory(rootDir+"/"+env+"/files/", tempDir)
		if err != nil {
			log.Fatal(err)
		}

		err = util.ReplaceInFiles(tempDir, dConfig.ConfigBoundary, dConfig.ConfigMap)
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(dConfig.Target.Address, s3Prefix) {
			dConfig.Target.Address = strings.Replace(dConfig.Target.Address, s3Prefix, "", 1)
			metadata := map[string]string{deploymentConfig.InternalBuckmateVersionMetadataKey: buckmateVersion}
			if dConfig.FileOptions != nil {
				dConfig.FileOptions[aws.InternalBuckmateFilePrefix] = deploymentConfig.FileOptions{Metadata: metadata}
			} else {
				dConfig.FileOptions = map[string]deploymentConfig.FileOptions{aws.InternalBuckmateFilePrefix: {Metadata: metadata}}
			}

			targetBucket := aws.NewBucket(client, dConfig.Target)

			uploadOptions := aws.UploadOptions{
				Prefix:      dConfig.Target.Prefix,
				FileOptions: dConfig.FileOptions,
				TempDir:     tempDir,
			}
			err := targetBucket.Upload(cmd.Context(), uploadOptions)
			if err != nil {
				log.Fatal(err)
			}

			removeOptions := aws.RemoveOptions{
				CurrentVersion: buckmateVersion,
			}

			err = targetBucket.RemovePreviousVersion(cmd.Context(), removeOptions)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			dConfig.Target.Address = util.Resolve(rootDir, dConfig.Target.Address)
			err = util.CopyDirectory(tempDir, dConfig.Target.Address)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = util.RemoveDirectory(tempDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}

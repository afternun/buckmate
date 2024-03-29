package cmd

import (
	"buckmate/main/common/util"
	"buckmate/main/deployment"
	"buckmate/main/download"
	"buckmate/main/upload"
	"os"

	log "github.com/sirupsen/logrus"

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
		tempDir := util.RandomDirectory()
		deploymentConfig := deployment.Load(env, rootDir)
		buckmateVersion := uuid.New().String()

		if strings.HasPrefix(deploymentConfig.Source.Address, s3Prefix) {
			deploymentConfig.Source.Address = strings.Replace(deploymentConfig.Source.Address, s3Prefix, "", 1)
			download.S3(deploymentConfig.Source.Address, deploymentConfig.Source.Prefix, tempDir)
		} else {
			deploymentConfig.Source.Address = util.Resolve(rootDir, deploymentConfig.Source.Address)
			err := util.CopyDirectory(deploymentConfig.Source.Address, tempDir)
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
		util.ReplaceInFiles(tempDir, deploymentConfig.ConfigBoundary, deploymentConfig.ConfigMap)

		if strings.HasPrefix(deploymentConfig.Target.Address, s3Prefix) {
			deploymentConfig.Target.Address = strings.Replace(deploymentConfig.Target.Address, s3Prefix, "", 1)
			upload.S3(deploymentConfig.Target.Address, deploymentConfig.Target.Prefix, buckmateVersion, tempDir)
		} else {
			deploymentConfig.Target.Address = util.Resolve(rootDir, deploymentConfig.Target.Address)
			util.CopyDirectory(tempDir, deploymentConfig.Target.Address)
		}

		util.RemoveDirectory(tempDir)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}

package cmd

import (
	"buckmate/main/common/util"
	"buckmate/main/deployment"
	"buckmate/main/download"
	"buckmate/main/upload"
	"log"

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
			log.Fatalln("Could not get --env flag")
		}
		deploymentConfig := deployment.Load(env)
		buckmateVersion := uuid.New().String()

		download.S3(deploymentConfig.Source.Bucket, deploymentConfig.Source.Prefix)
		util.ReplaceInFiles("build", deploymentConfig.ConfigBoundary, deploymentConfig.ConfigMap)
		upload.S3(deploymentConfig.Target.Bucket, deploymentConfig.Target.Prefix, buckmateVersion)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}

package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "buckmate",
	Short:   "Deploy to S3 buckets with ease",
	Long:    ``,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		logLevel, err := cmd.Flags().GetString("log")
		if err != nil {
			log.Fatal(err)
		}
		parsedLogLevel, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Fatal(err)
		}
		log.SetLevel(parsedLogLevel)
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("env", "e", "", "Specifies which config to apply - directory name that contains environment specific Config.yaml and files to be copied.")
	rootCmd.PersistentFlags().StringP("path", "p", "", "Specifies path to the directory that contains buckmate directory with Deployment.yaml config.")
	rootCmd.PersistentFlags().StringP("log", "l", "info", "Specifies log level. Options: panic, fatal, error, warn, info, debug trace.")
}

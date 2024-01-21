package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "buckmate",
	Short: "Deploy to S3 buckets with ease",
	Long:  ``,
}

func Execute() {
	logLevel, err := rootCmd.Flags().GetString("log")
	if err != nil {
		log.Fatal(err)
	}
	parsedLogLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(parsedLogLevel)
	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("env", "e", "", "Specifies which config to apply - directory name that contains environment specific Config.yaml and files to be copied.")
	rootCmd.PersistentFlags().StringP("path", "p", "", "Specifies path to the directory that contains buckmate directory with Deployment.yaml config.")
	rootCmd.PersistentFlags().StringP("log", "l", "InfoLevel", "Specifies log level. TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel.")
}

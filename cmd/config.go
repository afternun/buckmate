package cmd

import (
	"buckmate/main/common/util"
	"buckmate/main/deployment"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Apply config to files locally",
	Long: `Use:
buckmate config`,
	Run: func(cmd *cobra.Command, args []string) {
		env, err := cmd.Flags().GetString("env")
		if err != nil {
			log.Fatal(err)
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}
		rootDir := path + "/buckmate"

		tempDir := util.RandomDirectory()
		config := deployment.Load(env, rootDir)
		err = util.CopyDirectory(rootDir+"/files", tempDir)
		if err != nil {
			log.Fatal(err)
		}
		err = util.CopyDirectory(rootDir+"/"+env+"/files", tempDir)
		if err != nil {
			log.Fatal(err)
		}
		util.ReplaceInFiles(tempDir, config.ConfigBoundary, config.ConfigMap)

		log.Info("You can view your configuration in: " + tempDir)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

package cmd

import (
	"buckmate/main/common/util"
	"buckmate/main/deployment"
	"log"

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
			log.Fatalln("Could not get --env flag")
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln("Could not get --path flag")
		}
		rootDir := path + "/buckmate"

		tempDir := util.RandomDirectory()
		config := deployment.Load(env, rootDir)
		util.CopyDirectory(rootDir+"/files", tempDir)
		util.CopyDirectory(rootDir+"/"+env+"/files", tempDir)
		util.ReplaceInFiles(tempDir, config.ConfigBoundary, config.ConfigMap)

		log.Println("You can view your configuration in " + tempDir)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

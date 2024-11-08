package cmd

import (
	"buckmate/main/common/util"
	"buckmate/main/deploymentConfig"
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
			log.Fatal(err)
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}
		rootDir := path + "/buckmate"

		tempDir, err := util.RandomDirectory()
		if err != nil {
			log.Fatal(err)
		}
		config, err := deploymentConfig.Load(env, rootDir)
		if err != nil {
			log.Fatal(err)
		}
		err = util.CopyDirectory(rootDir+"/files", tempDir)
		if err != nil {
			log.Fatal(err)
		}
		err = util.CopyDirectory(rootDir+"/"+env+"/files", tempDir)
		if err != nil {
			log.Fatal(err)
		}
		err = util.ReplaceInFiles(tempDir, config.ConfigBoundary, config.ConfigMap)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("You can view your configuration in: " + tempDir)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

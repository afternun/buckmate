package cmd

import (
	"buckmate/main/common/constants"
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
		config := deployment.Load(env)
		util.ReplaceInFiles(constants.BUILD_DIRECTORY, config.ConfigBoundary, config.ConfigMap)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"buckmate/main/common/exception"
	"buckmate/main/common/util"
	"buckmate/main/config"
	"buckmate/main/deployment"
	"buckmate/main/download"
	"buckmate/main/upload"
	"buckmate/structs"
	"fmt"

	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		env, err := cmd.Flags().GetString("env")
		exception.Handle(structs.Exception{Err: err, Message: "Environment not set."})
		config := config.Load(env)
		deployment := deployment.Load()
		fmt.Printf("%v %v", config, deployment)
		fmt.Printf("Wersja do pobrania %s", config.Version)
		download.S3(deployment.Source.Path, config.Version)
		util.ReplaceInFiles("build", config.ConfigMap)
		upload.S3(deployment.Target.Path, "")
		// list := []string{"test", "test2"}
		// upload.S3(list, "file Key")
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

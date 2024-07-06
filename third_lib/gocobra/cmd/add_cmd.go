package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 创建根命令
var rootCmd = &cobra.Command{
	Use:   "gocobra",
	Short: "My App",
	Long:  "A sample application using Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Cobra!")
	},
}

func init() {
	// 子命令
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(subCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tryCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

var subCmd = &cobra.Command{
	Use:   "subcommand",
	Short: "A subcommand",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running subcommand")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init cmd",
	Long:  `init command something`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init command something")
	},
}
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("RunE tryCmd")
		return nil
	},
}

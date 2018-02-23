package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var filePath *string
var password *string

var rootCmd = &cobra.Command{
	Use:   "tf-secrets",
	Short: "tf-secret is the tool to handle .secrets.tfvars files used by the secrets terraform provider",
	Long:  `tf-secret is the tool to handle .secrets.tfvars files used by the secrets terraform provider`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	filePath = rootCmd.PersistentFlags().StringP("file", "f", "", "")
	password = rootCmd.PersistentFlags().StringP("password", "p", "", "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

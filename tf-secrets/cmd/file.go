package cmd

import (
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}

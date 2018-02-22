package cmd

import (
	"github.com/spf13/cobra"
)

var varFileCmd = &cobra.Command{
	Use:   "var-file",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(varFileCmd)
}

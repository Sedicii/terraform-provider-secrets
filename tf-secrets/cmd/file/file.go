package file

import (
	"github.com/spf13/cobra"
)

var filePath *string
var password *string

var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "Subcommand to manage .secret files",
	Long:  `Subcommand to manage .secret files`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	filePath = FileCmd.PersistentFlags().StringP("file", "f", "", "")
	FileCmd.MarkPersistentFlagFilename("file", "secret")
	FileCmd.MarkPersistentFlagRequired("file")

	password = FileCmd.PersistentFlags().StringP("password", "p", "", "")
	FileCmd.MarkPersistentFlagRequired("password")
}

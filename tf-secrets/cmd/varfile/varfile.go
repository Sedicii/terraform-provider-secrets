package varfile

import (
	"github.com/spf13/cobra"
)

var filePath *string
var password *string

var VarFileCmd = &cobra.Command{
	Use:   "var-file",
	Short: "Subcommand to manage .secret.tfvars files",
	Long:  `Subcommand to manage .secret.tfvars files`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	filePath = VarFileCmd.PersistentFlags().StringP("file", "f", "", "")
	VarFileCmd.MarkPersistentFlagFilename("file", "secret", "secrets.tfvars")
	VarFileCmd.MarkPersistentFlagRequired("file")

	password = VarFileCmd.PersistentFlags().StringP("password", "p", "", "")
	VarFileCmd.MarkPersistentFlagRequired("password")
}

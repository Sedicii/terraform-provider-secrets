package varfile

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"os"
)

var changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Changes the password of a .secrets.tfvars file",
	Long:  `Changes the password of a .secrets.tfvars file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := changePasswordVarFile(*filePath, *password, *varFileNewPassword)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}
var varFileNewPassword *string

func init() {
	varFileNewPassword = changePasswordCmd.PersistentFlags().StringP("new-password", "n", "", "")
	changePasswordCmd.MarkFlagRequired("new-password")

	VarFileCmd.AddCommand(changePasswordCmd)
}

func changePasswordVarFile(filePath string, password string, newPassword string) error {
	encryptedVars, err := lib.ReadHCLEncryptedVarFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing file %s : %s", filePath, err))
	}
	decryptedVars, err := lib.DecryptVars(encryptedVars, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting vars from file %s : %s", filePath, err))
	}

	encryptedVars, err = lib.EncryptVars(&decryptedVars, newPassword)

	if err != nil {
		return errors.New(fmt.Sprintf("Error encrypting vars : %s", err))
	}

	varFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		return errors.New(fmt.Sprintf("Error opening var file %s : %s", filePath, err))
	}

	err = lib.WriteEncryptedVarsAsHCL(varFile, encryptedVars)

	if err != nil {
		return errors.New(fmt.Sprintf("Error writing var file %s : %s", filePath, err))
	}
	return nil
}

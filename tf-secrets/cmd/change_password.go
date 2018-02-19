package cmd

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"os"
)

var newPassword *string

var changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Changes the password a specified .secrets.tfvars file",
	Long:  `Changes the password a specified .secrets.tfvars file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := changePasswordFile(*varFile, *password, *newPassword)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	newPassword = changePasswordCmd.PersistentFlags().StringP("new-password", "n", "", "")
	rootCmd.AddCommand(changePasswordCmd)
}

func changePasswordFile(filePath string, password string, new_password string) error {
	encryptedVars, err := lib.ReadHCLEncryptedVarFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing file %s : %s", filePath, err))
	}
	decryptedVars, err := lib.DecryptVars(encryptedVars, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting vars from file %s : %s", filePath, err))
	}

	encryptedVars, err = lib.EncryptVars(&decryptedVars, *newPassword)

	if err != nil {
		return errors.New(fmt.Sprintf("Error encrypting vars : %s", err))
	}

	varFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return errors.New(fmt.Sprintf("Error opening var file %s : %s", filePath, err))
	}

	err = lib.WriteEncryptedVarsAsHCL(varFile, encryptedVars)

	if err != nil {
		return errors.New(fmt.Sprintf("Error writing var file %s : %s", filePath, err))
	}
	return nil
}

package cmd

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"os"
)

var fileChangePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := changePasswordFile(*filePath, *password, *newPassword)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	newPassword = fileChangePasswordCmd.PersistentFlags().StringP("new-password", "n", "", "")
	fileCmd.AddCommand(fileChangePasswordCmd)
}

func changePasswordFile(filePath string, password string, newPassword string) error {
	encryptedContent, err := lib.ReadEncryptedFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing file %s : %s", filePath, err))
	}
	decryptedContent, err := lib.DecryptData(*encryptedContent, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting file %s : %s", filePath, err))
	}

	*encryptedContent, err = lib.EncryptData(string(decryptedContent), newPassword)

	if err != nil {
		return errors.New(fmt.Sprintf("Error encrypting file : %s", err))
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		return errors.New(fmt.Sprintf("Error opening file %s : %s", filePath, err))
	}

	err = lib.WriteEncryptedFile(file, encryptedContent)

	if err != nil {
		return errors.New(fmt.Sprintf("Error writing file %s : %s", filePath, err))
	}
	return nil
}

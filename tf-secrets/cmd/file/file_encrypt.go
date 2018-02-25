package file

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var fileEncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file to be .secret",
	Long:  `Encrypts a file to be .secret`,
	Run: func(cmd *cobra.Command, args []string) {
		err := encryptFile(*filePath, *encryptDstFilePath, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var encryptDstFilePath *string

func init() {
	encryptDstFilePath = fileEncryptCmd.PersistentFlags().StringP("dst-file", "d", "", "")
	fileEncryptCmd.MarkFlagRequired("dst-file")
	fileEncryptCmd.MarkPersistentFlagFilename("dst-file", "secret")

	FileCmd.AddCommand(fileEncryptCmd)
}

func encryptFile(filePath string, dstFilePath string, password string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading file %s : %s", filePath, err))
	}
	encryptedData, err := lib.EncryptData(string(content), password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error encrypting data of file %s : %s", filePath, err))
	}
	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Error opening file %s : %s", dstFilePath, err))
	}

	err = lib.WriteEncryptedFile(dstFile, &encryptedData)
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing file %s : %s", dstFilePath, err))
	}
	return nil
}

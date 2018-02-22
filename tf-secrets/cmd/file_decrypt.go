package cmd

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var fileDecryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := decryptFile(*filePath, *dstFilePath, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	dstFilePath = fileDecryptCmd.PersistentFlags().StringP("dst-file", "-d", "", "")
	fileCmd.AddCommand(fileDecryptCmd)
}

func decryptFile(filePath string, dstFilePath string, password string) error {
	encryptedContent, err := lib.ReadEncryptedFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading file %s : %s", filePath, err))

		return err
	}
	decryptedData, err := lib.DecryptData(*encryptedContent, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting file %s : %s", filePath, err))
	}
	err = ioutil.WriteFile(dstFilePath, decryptedData, 600)
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing file %s : %s", dstFilePath, err))
	}
	return nil
}

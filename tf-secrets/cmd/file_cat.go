package cmd

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"os"
)

var fileCatCmd = &cobra.Command{
	Use:   "cat",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := catFile(*filePath, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	fileCmd.AddCommand(fileCatCmd)
}

func catFile(filePath string, password string) error {
	encryptedContent, err := lib.ReadEncryptedFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading file %s : %s", filePath, err))

		return err
	}
	decryptedData, err := lib.DecryptData(*encryptedContent, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting file %s : %s", filePath, err))
	}
	_, err = os.Stdout.Write(decryptedData)
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing to stdout : %s", err))
	}
	return nil
}

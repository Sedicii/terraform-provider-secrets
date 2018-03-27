package file

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var fileDecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts a .secret file",
	Long:  `Decrypts a .secret file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := decryptFile(*filePath, *decryptDstFilePath, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var decryptDstFilePath *string

func init() {
	decryptDstFilePath = fileDecryptCmd.PersistentFlags().StringP("dst-file", "d", "", "")
	fileDecryptCmd.MarkFlagRequired("dst-file")
	fileDecryptCmd.MarkPersistentFlagFilename("dst-file")

	FileCmd.AddCommand(fileDecryptCmd)
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
	err = ioutil.WriteFile(dstFilePath, decryptedData, 0600)
	if err != nil {
		fmt.Println(dstFilePath)
		return errors.New(fmt.Sprintf("Error writing file %s : %s", dstFilePath, err))
	}
	return nil
}

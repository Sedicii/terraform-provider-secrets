package varfile

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits a specified .secrets.tfvars file",
	Long:  `Edits a specified .secrets.tfvars file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := editVarFile(*filePath, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	VarFileCmd.AddCommand(editCmd)
}

func editVarFile(filePath string, password string) error {
	encryptedVars, err := lib.ReadHCLEncryptedVarFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing file %s : %s", filePath, err))
	}
	decryptedVars, err := lib.DecryptVars(encryptedVars, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting vars from file %s : %s", filePath, err))
	}

	tmpFile, err := ioutil.TempFile("/tmp", "tf-secrets")

	if err != nil {
		return errors.New(fmt.Sprintf("Error creating temporal file %s : %s", tmpFile.Name(), err))
	}
	defer os.Remove(tmpFile.Name())

	err = lib.WriteDecryptedVarsAsHCL(tmpFile, &decryptedVars)

	if err != nil {
		return errors.New(fmt.Sprintf("Error writing temporal file %s : %s", tmpFile.Name(), err))
	}

	tmpFile.Close()
	tmpFileName := tmpFile.Name()

	return editVarFileAndEncrypt(tmpFileName, password, filePath)
}

func editVarFileAndEncrypt(tmpFileName string, password string, varFilePath string) error {

	lib.EditFileWithEditor(tmpFileName)

	decryptedVars, err := lib.ReadHCLDecryptedVarFile(tmpFileName)

	if err != nil {
		return errors.New(fmt.Sprintf("Error reading var file : %s", err))
	}

	encryptedVars, err := lib.EncryptVars(decryptedVars, password)

	if err != nil {
		return errors.New(fmt.Sprintf("Error encrypting vars : %s", err))
	}

	varFile, err := os.OpenFile(varFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return errors.New(fmt.Sprintf("Error opening var file %s : %s", varFilePath, err))
	}

	err = lib.WriteEncryptedVarsAsHCL(varFile, encryptedVars)

	if err != nil {
		return errors.New(fmt.Sprintf("Error writing var file %s : %s", varFilePath, err))
	}
	return nil
}

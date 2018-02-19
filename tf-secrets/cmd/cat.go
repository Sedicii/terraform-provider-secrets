package cmd

import (
	"errors"
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"github.com/spf13/cobra"
	"os"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "outputs by stdout a specified .secrets.tfvars file",
	Long:  `outputs by stdout a specified .secrets.tfvars file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := catFile(*varFile, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}

func catFile(filePath string, password string) error {
	encryptedVars, err := lib.ReadHCLEncryptedVarFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing file %s : %s", filePath, err))
	}
	decryptedVars, err := lib.DecryptVars(encryptedVars, password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decrypting vars from file %s : %s", filePath, err))
	}

	err = lib.WriteDecryptedVarsAsHCL(os.Stdout, &decryptedVars)

	if err != nil {
		return errors.New(fmt.Sprintf("Error writing to stdout : %s", err))
	}

	return nil
}

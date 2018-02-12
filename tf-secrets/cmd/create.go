package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the new .secrets.tfvars file specified",
	Long:  `Creates the new .secrets.tfvars file specified`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createFile(*varFile, *password)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createFile(filePath string, password string) error {
	tmpFile, err := ioutil.TempFile("/tmp", "tf-secrets")

	if err != nil {
		return errors.New(fmt.Sprintf("Error creating temporal file %s : %s", tmpFile.Name(), err))
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.Close()
	tmpFileName := tmpFile.Name()

	return editFileAndEncrypt(tmpFileName, password, filePath)
}

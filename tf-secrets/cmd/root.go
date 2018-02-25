package cmd

import (
	"fmt"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/cmd/file"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/cmd/varfile"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "tf-secrets",
	Short: "tf-secret is the tool to handle .secrets.tfvars and .secret files used by the secrets terraform provider",
	Long:  `tf-secret is the tool to handle .secrets.tfvars and .secret files used by the secrets terraform provider`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var bashCompletionCmd = &cobra.Command{
	Use:   "bash-completion",
	Short: "Outputs to stdout the bash completion of tf-secrets",
	Long:  `Outputs to stdout the bash completion of tf-secrets`,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletion(os.Stdout)
	},
}

var zshCompletionCmd = &cobra.Command{
	Use:   "zsh-completion",
	Short: "Outputs to stdout the zsh completion of tf-secrets",
	Long:  `Outputs to stdout the zsh completion of tf-secrets`,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(bashCompletionCmd)
	RootCmd.AddCommand(zshCompletionCmd)
	RootCmd.AddCommand(file.FileCmd)
	RootCmd.AddCommand(varfile.VarFileCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

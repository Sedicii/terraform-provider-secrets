package main

import (
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "./")
	if err != nil {
		log.Fatal(err)
	}
}

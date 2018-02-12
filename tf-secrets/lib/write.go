package lib

import (
	"fmt"
	"io"
)

func WriteDecryptedVarsAsHCL(writer io.Writer, decryptedVars *map[string]string) (err error) {
	for key, value := range *decryptedVars {
		variableDeclaration := fmt.Sprintf("%s = \"%s\" \n\n", key, value)
		_, err = writer.Write([]byte(variableDeclaration))
		if err != nil {
			return
		}
	}
	return
}

func WriteEncryptedVarsAsHCL(writer io.Writer, encryptedVars *map[string]map[string]string) (err error) {
	varHCLTemplate := `
%s = {
	data = "%s"
	salt = "%s"
	nonce = "%s"
}

`
	for key, value := range *encryptedVars {
		varHCL := fmt.Sprintf(varHCLTemplate, key, value["data"], value["salt"], value["nonce"])
		_, err = writer.Write([]byte(varHCL))
		if err != nil {
			return
		}
	}
	return
}

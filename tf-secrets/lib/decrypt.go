package lib

import "github.com/sedicii/terraform-provider-secrets/aes"

func DecryptVars(encryptedVars *map[string]map[string]string, password string) (map[string]string, error) {
	decryptedVars := make(map[string]string)
	for key, value := range *encryptedVars {
		decryptedValue, err := decryptVar(value["data"], value["salt"], value["nonce"], password)
		if err != nil {
			return nil, err
		}
		decryptedVars[key] = string(decryptedValue)
	}
	return decryptedVars, nil
}

func decryptVar(data string, salt string, nonce string, password string) ([]byte, error) {
	return aes.HexDecryptAES256(data, salt, nonce, password)
}

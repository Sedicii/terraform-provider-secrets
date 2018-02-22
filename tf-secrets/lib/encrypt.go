package lib

import "github.com/sedicii/terraform-provider-secrets/aes"

func EncryptVars(decryptedVars *map[string]string, password string) (*map[string]map[string]string, error) {
	encryptedVars := make(map[string]map[string]string)
	for key, value := range *decryptedVars {
		encryptedValue, err := EncryptData(value, password)
		if err != nil {
			return nil, err
		}
		encryptedVars[key] = encryptedValue
	}
	return &encryptedVars, nil
}

func EncryptData(data string, password string) (map[string]string, error) {
	encryptedData, salt, nonce, err := aes.HexEncryptAES256(data, password)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"data":  encryptedData,
		"salt":  salt,
		"nonce": nonce,
	}, nil

}

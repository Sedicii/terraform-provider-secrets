package lib

import (
	"github.com/sedicii/terraform-provider-secrets/aes"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func EncryptVars(decryptedVars *map[string]interface{}, password string) (*map[string]interface{}, error) {
	encryptedVars := make(map[string]interface{})
	for key, value := range *decryptedVars {
		switch v :=value.(type) {
		case map[string]interface{}:
			subEncryptedVars, err := EncryptVars(&v, password)
			if err != nil {return nil, err}
			encryptedVars[key] = subEncryptedVars
		case string:
			encryptedVar, err := EncryptData(v, password)
			if err != nil {return nil, err}
			encryptedVars[key] = encryptedVar
		case []interface{}:
			encryptedVarsList, err := EncryptVarsList(&v, password)
			if err != nil {return nil, err}
			encryptedVars[key] = encryptedVarsList
			}
		}
	return &encryptedVars, nil
}

func EncryptVarsList(decryptedVars *[]interface{}, password string) (*[]interface{}, error) {
	encryptedVarsArray := make([]interface{}, len(*decryptedVars))
	for id, value := range *decryptedVars {
		switch v := value.(type) {
		case map[string]interface{}:
			subEncryptedVars, err := EncryptVars(&v, password)
			if err != nil {return nil, err}
			encryptedVarsArray[id] = subEncryptedVars
		case string:
			encryptedVar, err := EncryptData(v, password)
			if err != nil {return nil, err}
			encryptedVarsArray[id] = encryptedVar
		case []interface{}:
			encryptedVarsList, err := EncryptVarsList(&v, password)
			if err != nil {return nil, err}
			encryptedVarsArray[id] = encryptedVarsList
		}

	}
	return &encryptedVarsArray, nil
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

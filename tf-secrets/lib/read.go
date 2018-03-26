package lib

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl"
	"io/ioutil"
)

func ReadHCLEncryptedVarFile(filePath string) (*map[string]map[string]string, error) {
	encryptedVarFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return ParseHCLEncryptedVarFile(string(encryptedVarFile))
}

func ParseHCLEncryptedVarFile(fileContent string) (*map[string]map[string]string, error) {
	astFile, err := hcl.Parse(fileContent)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := hcl.DecodeObject(&result, astFile); err != nil {
		return nil, fmt.Errorf(
			"Error decoding Terraform vars file: %s\n\n"+
				"The vars file should be in the format of `key = \"value\"`.\n"+
				"Decoding errors are usually caused by an invalid format.",
			err)
	}
	encryptedVars := make(map[string]map[string]string)

	err = flattenMultiMaps(result)

	for k, v := range result {
		switch v.(type) {
		case map[string]interface{}:
			strMap, err := mapInterfaceToString(v.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			encryptedVars[k] = strMap
			break
		default:
			fmt.Println(v)
			return nil, fmt.Errorf("wrong type for value of key %s, secrets vars encrypted are maps", k)
		}
	}

	return &encryptedVars, nil
}

func mapInterfaceToString(m map[string]interface{}) (map[string]string, error) {
	m2 := make(map[string]string)
	for k, v := range m {
		switch v.(type) {
		case string:
			m2[k] = v.(string)
			break
		default:
			return nil, fmt.Errorf("wrong type for value of key %s, secrets vars encrypted are maps of strings", k)
		}
	}
	return m2, nil
}

// Variables don't support any type that can be configured via multiple
// declarations of the same HCL map, so any instances of
// []map[string]interface{} are either a single map that can be flattened, or
// are invalid config.
func flattenMultiMaps(m map[string]interface{}) error {
	for k, v := range m {
		switch v := v.(type) {
		case []map[string]interface{}:
			switch {
			case len(v) > 1:
				return fmt.Errorf("multiple map declarations not supported for variables")
			case len(v) == 1:
				m[k] = v[0]
			}
		}
	}
	return nil
}

func ReadHCLDecryptedVarFile(filePath string) (*map[string]interface{}, error) {
	decryptedVarFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	astFile, err := hcl.Parse(string(decryptedVarFile))
	var decryptedVars map[string]interface{}

	if err := hcl.DecodeObject(&decryptedVars, astFile); err != nil {
		return nil, fmt.Errorf(
			"Error decoding Terraform vars file: %s\n\n"+
				"The vars file should be in the format of `key = \"value\"`.\n"+
				"Decoding errors are usually caused by an invalid format.",
			err)
	}

	return &decryptedVars, nil
}

func ReadEncryptedFile(filePath string) (*map[string]string, error) {
	encryptedFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	encryptedFileContent := make(map[string]string)
	err = json.Unmarshal(encryptedFile, &encryptedFileContent)
	if err != nil {
		return nil, err
	}
	return &encryptedFileContent, nil
}

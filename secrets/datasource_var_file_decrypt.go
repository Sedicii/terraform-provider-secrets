package secrets

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
	"io/ioutil"
)

var NoFileError = errors.New("")

func varFileDecrypt() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVarFileDecrypt,

		Schema: map[string]*schema.Schema{
			"var_file": &schema.Schema{
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Default:     "",
				Description: "Encrypted var file content",
			},
			"var_file_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Default:     "",
				Description: "Encrypted var file path",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password used to decrypt the file",
			},
			"values": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Decrypted file content",
			},
		},
	}
}

func dataSourceVarFileDecrypt(d *schema.ResourceData, meta interface{}) error {
	file := d.Get("var_file").(string)
	filePath := d.Get("var_file_path").(string)

	file, err := getFileContent(file, filePath)
	if err == NoFileError {
		d.Set("values", make(map[string]string))
		d.SetId(hash(""))
		return nil
	} else if err != nil {
		return err
	}

	password := d.Get("password").(string)
	values, err := decryptVarFile(file, password)
	if err != nil {
		return err
	}
	d.Set("values", values)
	raw, err := json.Marshal(values)
	if err != nil {
		return err
	}
	d.SetId(hash(string(raw)))
	return nil
}

func getFileContent(file, filePath string) (string, error) {
	if file != "" {
		return file, nil
	} else if filePath != "" {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", NoFileError
		}
		return string(data), nil
	}

	return "", errors.New("file or filePath needed")

}

func decryptVarFile(fileContent string, password string) (map[string]string, error) {
	encryptedVars, err := lib.ParseHCLEncryptedVarFile(fileContent)
	if err != nil {
		return nil, err
	}
	return lib.DecryptVars(encryptedVars, password)
}

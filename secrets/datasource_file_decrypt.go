package secrets

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sedicii/terraform-provider-secrets/aes"
)

func fileDecrypt() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFileDecrypt,

		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Encrypted file content",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password used to decrypt the file",
			},
			"value": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Decrypted file content",
			},
		},
	}
}

func dataSourceFileDecrypt(d *schema.ResourceData, meta interface{}) error {
	file := d.Get("file").(string)
	password := d.Get("password").(string)
	value, err := decryptFile(file, password)
	if err != nil {
		return err
	}
	strValue := string(value)
	d.Set("value", strValue)
	d.SetId(hash(strValue))
	return nil
}

func decryptFile(data string, password string) ([]byte, error) {
	encryptedFileContent := make(map[string]string)
	err := json.Unmarshal([]byte(data), &encryptedFileContent)
	if err != nil {
		return nil, err
	}
	return aes.HexDecryptAES256(encryptedFileContent["data"], encryptedFileContent["salt"], encryptedFileContent["nonce"], password)
}

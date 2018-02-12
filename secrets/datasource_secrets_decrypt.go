package secrets

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sedicii/terraform-provider-secrets/aes"
)

func secretsDecrypt() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretsDecrypt,

		Schema: map[string]*schema.Schema{
			"var": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "Encrypted variable",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password used to decrypt the variable",
			},
			"value": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Decrypted value",
			},
		},
	}
}

func dataSourceSecretsDecrypt(d *schema.ResourceData, meta interface{}) error {
	variable := d.Get("var").(map[string]interface{})
	password := d.Get("password").(string)
	value, err := decryptData(variable["data"].(string), variable["salt"].(string), variable["nonce"].(string), password)
	if err != nil {
		return err
	}
	strValue := string(value)
	d.Set("value", strValue)
	d.SetId(hash(strValue))
	return nil
}

func decryptData(data string, salt string, nonce string, password string) ([]byte, error) {
	return aes.HexDecryptAES256(data, salt, nonce, password)
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

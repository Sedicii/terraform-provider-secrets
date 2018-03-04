package secrets

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
)

func varFileDecrypt() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVarFileDecrypt,

		Schema: map[string]*schema.Schema{
			"var_file": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Encrypted var file content",
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

func decryptVarFile(fileContent string, password string) (map[string]string, error) {
	encryptedVars, err := lib.ParseHCLEncryptedVarFile(fileContent)
	if err != nil {
		return nil, err
	}
	return lib.DecryptVars(encryptedVars, password)
}

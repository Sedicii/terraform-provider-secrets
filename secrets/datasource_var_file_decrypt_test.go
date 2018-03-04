package secrets

import (
	"fmt"
	"testing"

	"bytes"
	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sedicii/terraform-provider-secrets/aes"
	"github.com/sedicii/terraform-provider-secrets/tf-secrets/lib"
)

func TestTemplateRenderingVarFile(t *testing.T) {
	var cases = []struct {
		password string
		want     map[string]string
	}{
		{`test_password2`, map[string]string{`a_secret`: "test"}},
		{`test_password1`, map[string]string{`a_secret`: "test", `another_secret_in_the_same_file`: "test2"}},
	}

	for _, tt := range cases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				r.TestStep{
					Config: testTemplateConfigFile(encryptVarFile(tt.want, tt.password), tt.password),
					Check: func(s *terraform.State) error {
						got := s.RootModule().Outputs["decrypted"]
						for k, _ := range tt.want {
							if tt.want[k] != got.Value.(map[string]string)[k] {
								return fmt.Errorf("secret:\n%s\npassword:\n%s\ngot:\n%s\nwant:\n%s\n", encryptVarFile(tt.want, tt.password), tt.password, got, tt.want)
							}
						}

						return nil
					},
				},
			},
		})
	}
}

func encryptVarFile(vars map[string]string, password string) string {
	encryptedVars, err := lib.EncryptVars(&vars, password)
	if err != nil {
		panic(err)
	}
	encFileContent := ""
	encFileContentW := bytes.NewBufferString(encFileContent)
	err = lib.WriteEncryptedVarsAsHCL(encFileContentW, encryptedVars)
	if err != nil {
		panic(err)
	}
	return encFileContent
}

func testTemplateConfigVarFile(encryptedFile, password string) string {
	str := fmt.Sprintf(`
		data "secrets_var_file_decrypt" "t0" {
			file = %s
			password = "%s"
		}
		output "decrypted" {
				value = "${data.secrets_file_decrypt.t0.value}"
		}`, encryptedFile, password)
	fmt.Println(str)
	return str
}

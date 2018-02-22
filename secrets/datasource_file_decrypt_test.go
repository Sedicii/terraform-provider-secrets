package secrets

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sedicii/terraform-provider-secrets/aes"
)

func TestTemplateRenderingFile(t *testing.T) {
	var cases = []struct {
		password string
		want     string
	}{
		{`test_password1`, `a_secret`},
		{`test_password2`, `another_secret`},
		{`test_password3`, `one_secret_more`},
	}

	for _, tt := range cases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				r.TestStep{
					Config: testTemplateConfigFile(encryptFile(tt.want, tt.password), tt.password),
					Check: func(s *terraform.State) error {
						got := s.RootModule().Outputs["decrypted"]
						if tt.want != got.Value {
							return fmt.Errorf("secret:\n%s\npassword:\n%s\ngot:\n%s\nwant:\n%s\n", encryptFile(tt.want, tt.password), tt.password, got, tt.want)
						}
						return nil
					},
				},
			},
		})
	}
}

func encryptFile(value string, password string) string {
	encData, salt, nonce, err := aes.HexEncryptAES256(value, password)
	if err != nil {
		panic("Error ciphering in test")
	}
	return fmt.Sprintf(`"{\"data\": \"%s\", \"salt\": \"%s\", \"nonce\": \"%s\"}"`, encData, salt, nonce)

}

func testTemplateConfigFile(encryptedFile, password string) string {
	str := fmt.Sprintf(`
		data "secrets_file_decrypt" "t0" {
			file = %s
			password = "%s"
		}
		output "decrypted" {
				value = "${data.secrets_file_decrypt.t0.value}"
		}`, encryptedFile, password)
	fmt.Println(str)
	return str
}

package secrets

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testProviders = map[string]terraform.ResourceProvider{
	"secrets": Provider(),
}

func TestTemplateRendering(t *testing.T) {
	var cases = []struct {
		encryptedVar string
		password     string
		want         string
	}{
		{`{
			data = "1285d9c1eebf1da5910e809a48e92b23219ded2c32111b87"
			salt = "bacc51b4397c2a5ad20759cb3a65b1f7d167d056e0d132626c7b345ab8cd927d"
        	nonce = "759a50abc9ff4a25a38d5f85"}`,
			`test_password1`, `a_secret`},
		{`{
			data = "870bdb56c189f9421dbdaa49e9b9498726fa0c0e97e38188b3e7fa565bf0"
        	salt = "6f3057d7bd5606e54bb8280bec25312010b6bd0e0f83d3eb0d16362053698d35"
        	nonce = "15bb9a3a54df85f4a50d2d68"}`,
			`test_password2`, `another_secret`},
		{`{
 			data = "2b81505aa0cda54dcb864d30877b821aa8fd6bae24e2743071554b81287e15"
        	salt = "fcc0cf6cbac67c0d2096bc9ac68b60b0c5168e8936788185d16093e63d57dd6a"
        	nonce = "408fc39b5dc2213a32903570"}`,
			`test_password3`, `one_secret_more`},
	}

	for _, tt := range cases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				r.TestStep{
					Config: testTemplateConfig(tt.encryptedVar, tt.password),
					Check: func(s *terraform.State) error {
						got := s.RootModule().Outputs["decrypted"]
						if tt.want != got.Value {
							return fmt.Errorf("secret:\n%s\npassword:\n%s\ngot:\n%s\nwant:\n%s\n", tt.encryptedVar, tt.password, got, tt.want)
						}
						return nil
					},
				},
			},
		})
	}
}

func testTemplateConfig(encryptedVar, password string) string {
	return fmt.Sprintf(`
		data "secrets_decrypt" "t0" {
			var = %s
			password = "%s"
		}
		output "decrypted" {
				value = "${data.secrets_decrypt.t0.value}"
		}`, encryptedVar, password)
}

package secrets

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"secrets_decrypt":      secretsDecrypt(),
			"secrets_file_decrypt": fileDecrypt(),
		},
	}
}

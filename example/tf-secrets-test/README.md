# Terraform Provider Secrets Example Project

This project is an example of the features of the secrets provider.

### How to run the example
```bash
# Install secrets provider
curl https://raw.githubusercontent.com/Sedicii/terraform-provider-secrets/master/scripts/install-secrets-tf-plugin.sh | bash
# Init terraform
terraform init
# Apply plan
terraform apply -var-file=./github.secrets.tfvars
```

### How the secrets files where created

This command opens an editor to add the secrets you want to add to the file  
```bash
tf-secrets var-file create -f ./github.secrets.tfvars -p "test_password"
```

This command opens an editor to add the secrets you want to add to the file  
```bash
tf-secrets var-file create -f ./many_secrets.secrets.tfvars -p "test_password"
```

This command encrypts a file to be used by "secrets_file_decrypt" data provider 
```bash
tf-secrets file encrypt -f ./id_rsa -d ./id_rsa.secret -p "test_password"
```


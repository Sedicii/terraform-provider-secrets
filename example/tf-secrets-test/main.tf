
variable "github_oauth_token" {type = "map"}

provider "secrets" {
  version = "~> 0.1.0"
}

data "secrets_decrypt" "github_oauth_token" {
  var = "${var.github_oauth_token}"
  password = "test_password"
}

data "secrets_file_decrypt" "ssh_key" {
  file = "${file("${path.module}/id_rsa.secret")}"
  password = "test_password"
}

data "secrets_var_file_decrypt" "many_secrets" {
  var_file = "${file("${path.module}/many_secrets.secrets.tfvars")}"
  password = "test_password"
}

data "secrets_var_file_decrypt" "no_secrets" {
  var_file_path = "${"${path.module}/non_existing_file.secrets.tfvars"}"
  password = "test_password"
}


output "no_file" {
  value = "${data.secrets_var_file_decrypt.no_secrets.values}"
}


output "github_oauth_token" {
  value = "${data.secrets_decrypt.github_oauth_token.value}"
}

output "ssh_key" {
  value = "${data.secrets_file_decrypt.ssh_key.value}"
}

output "postgres_password" {
value = "${data.secrets_var_file_decrypt.many_secrets.values.postgres_password}"
}

output "mongo_password" {
  value = "${data.secrets_var_file_decrypt.many_secrets.values.mongo_password}"
}

output "redis_password" {
  value = "${data.secrets_var_file_decrypt.many_secrets.values.redis_password}"
}


variable "github_oauth_token" {}

provider "secrets" {
  version = "~> 0.1.0"
}

data "secrets_decrypt" "github_oauth_token" {
  var = "${var.github_oauth_token}"
  password = "test_password"
}

data "secrets_file_decrypt" "cert" {
  file = "${file("${path.module}/cert.secret")}"
  password = "test_password"
}


output "github_oauth_token" {
  value = "${data.secrets_decrypt.github_oauth_token.value}"
}

output "cert" {
  value = "${data.secrets_file_decrypt.cert.value}"
}
# Terraform Secrets Provider

This Terraform provides allows to use secrets in terraform and commit them encrypted in your repository.
The philosophy of this provider is the same as chef's encrypted data bags.
This provider is composed of the provider itself and a cli tool called tf-secrets to manage encrypted terraform vars.

The crypto algorithms used are :
 * PBKDF2 [RFC-2898](https://www.ietf.org/rfc/rfc2898.txt) as key generation algorithm
 * AES256 [RFC-]()

### Maintainers

This provider plugin is maintained by [Sedicii](https://sedicii.com/).

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

### Installation

```bash
curl https://raw.githubusercontent.com/Sedicii/terraform-provider-secrets/master/scripts/install-secrets-tf-pluging.sh | bash
```

### Usage

#### Terraform Code
```

provider "secrets" {
  version = "~> 0.1.0"
}

data "secrets_decrypt" "github_oauth_token" {
  var = "${var.github_oauth_token}"
  password = "${var.secrets_master_password}"
}

output "github_oauth_token" {
  value = "${data.secrets_decrypt.github_oauth_token.value}"
}
```

#### tf-secrets cli
```bash
# To create a new secrets file
tf-secrets create github.secrets.tfvars -p <secrets_master_password>

# To edit a secrets file
tf-secrets edit github.secrets.tfvars -p <secrets_master_password>
```

For a more detailed example look at the example directory

### Building The Provider

Clone repository to: `$GOPATH/src/github.com/sedicii/terraform-provider-secrets`

```sh
$ mkdir -p $GOPATH/src/github.com/sedicii; cd $GOPATH/src/github.com/sedicii
$ git clone git@github.com:sedicii/terraform-provider-secrets
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/sedicii/terraform-provider-secrets
$ make build
```

### Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-secrets
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

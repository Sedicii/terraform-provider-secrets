FROM hashicorp/terraform:light

RUN apk add bash && \
    curl https://raw.githubusercontent.com/Sedicii/terraform-provider-secrets/master/scripts/install-secrets-tf-plugin.sh | bash



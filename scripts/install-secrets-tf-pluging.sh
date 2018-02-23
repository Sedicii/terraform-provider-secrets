#!/usr/bin/env bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
VERSION="v0.1.0"
PROVIDER_DOWNLOAD_URL="https://github.com/Sedicii/terraform-provider-secrets/releases/download/${VERSION}/terraform-provider-secrets_${OS}-amd64_${VERSION}"
CLI_DOWNLOAD_URL="https://github.com/Sedicii/terraform-provider-secrets/releases/download/${VERSION}/tf-secrets_${OS}-amd64_${VERSION}"
TF_PLUGINS_PATH="${HOME}/.terraform.d/plugins/${OS}_amd64/"
PLUGIN_DEST="${TF_PLUGINS_PATH}/terraform-provider-secrets_${VERSION}"
mkdir -p "${TF_PLUGINS_PATH}"

curl -L ${PROVIDER_DOWNLOAD_URL} --output ${PLUGIN_DEST}
chmod +x ${PLUGIN_DEST}

CLI_DEST="/usr/local/bin/tf-secrets"

if [ "${OS}" == "darwin" ]; then
    curl -L ${CLI_DOWNLOAD_URL} --output ${CLI_DEST}
    chmod +x ${CLI_DEST}
else
    sudo curl -L ${CLI_DOWNLOAD_URL} --output ${CLI_DEST}
    sudo chmod 755 ${CLI_DEST}
fi


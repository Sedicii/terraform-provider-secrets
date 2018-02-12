#!/usr/bin/env bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
VERSION="v0.1.0"
PROVIDER_DOWNLOAD_URL="https://github.com/Sedicii/terraform-provider-secrets/releases/download/${VERSION}/terraform-provider-secrets_${OS}-amd64_${VERSION}"
TF_PLUGINS_PATH="${HOME}/.terraform.d/plugins/${OS}_amd64/"

mkdir -p "${TF_PLUGINS_PATH}"
curl ${PROVIDER_DOWNLOAD_URL} --output "${TF_PLUGINS_PATH}/terraform-provider-secrets_${VERSION}"

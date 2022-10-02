#!/usr/bin/env bash
# Copyright 2020 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0

source ./common.sh


header_text "Checking for bin/kubectl ${KUBE_VERSION}"
[[ -f bin/kubectl && `bin/kubectl version -o=yaml 2> /dev/null | bin/yq e '.clientVersion.gitVersion' -` == "v${KUBE_VERSION}" ]] && exit 0
header_text "Installing bin/kubectl ${KUBE_VERSION}"


mkdir -p ./bin
curl -LO  "https://dl.k8s.io/release/v${KUBE_VERSION}/bin/linux/amd64/kubectl"

chmod 777 kubectl
mv kubectl bin

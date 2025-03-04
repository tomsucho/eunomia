#!/usr/bin/env bash

# Copyright 2019 Kohl's Department Stores, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euxo pipefail

TAG_OWNER="gitopsconfig.eunomia.kohls.io/owner"
TAG_APPLIED="gitopsconfig.eunomia.kohls.io/applied"

# this is needed because we want the current namespace to be set as default if a namespace is not specified.
function setContext {
  $kubectl config set-context current --namespace="$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace)"
  $kubectl config use-context current
}

function kube {
  $kubectl \
    -s https://kubernetes.default.svc:443 \
    --token "$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
    --certificate-authority=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt \
    "$@"
}

function deleteResources {
  #first we need to delete the GitOpsConfig resources whose finalizer might not get called otherwise
  for file in $(find "$MANIFEST_DIR" -iregex '.*\.yaml'); do
    cat "$file" | yq 'select(.kind == "GitOpsConfig")' | kube delete -f - --wait=true
  done
  kube delete -R -f "$MANIFEST_DIR"
  # TODO: deleteByOldLabels deleteAll
}

# addLabels OWNER TIMESTAMP - patches the YAML&JSON files in $MANIFEST_DIR,
# adding labels tracking the OWNER and TIMESTAMP. The labels are intended to be
# used later in function deleteByOldLabels.
function addLabels {
  local owner="$1"
  local timestamp="$2"
  local tmpdir="$(mktemp -d)"
  for file in $(find "$MANIFEST_DIR" -iregex '.*\.\(ya?ml\|json\)'); do
    cat "$file" |
      yq -y -s "map(select(.!=null)|setpath([\"metadata\",\"labels\",\"$TAG_OWNER\"]; \"$owner\"))|.[]" |
      yq -y -s "map(select(.!=null)|setpath([\"metadata\",\"labels\",\"$TAG_APPLIED\"]; \"$timestamp\"))|.[]" \
      > "$tmpdir/labeled"
    # We must use a helper file (can't do this in single step), as the file would be truncated if we read & write from it in one pipeline
    cat "$tmpdir/labeled" > "$file"
  done
}

# deleteByOldLabels OWNER TIMESTAMP - deletes all kubernetes resources which have
# the OWNER label as provided, but TIMESTAMP label different than provided.
function deleteByOldLabels {
  local owner="$1"
  local timestamp="$2"
  local allKinds="$(kube api-resources --verbs=list -o name | paste -sd, -)"
  local ownedKinds="$(kube get "$allKinds" --ignore-not-found \
      -l "$TAG_OWNER==$owner" \
      -o custom-columns=kind:.kind \
      --no-headers=true |
    sort -u |
    paste -sd, -)"
  if [ -z "$ownedKinds" ]; then
    return
  fi
  kube delete "$ownedKinds" \
    -l "$TAG_OWNER==$owner,$TAG_APPLIED!=$timestamp"
}

function createUpdateResources {
  # NOTE: Kubernetes currently requires that first *and last* character of
  # label values are alphanumerical - we're adding the ".owned" suffix to
  # ensure that.
  local owner="ns-$NAMESPACE.gitopsconfig-$GITOPSCONFIG_NAME.owned"
  local timestamp="$(date +%s)"
  case "$CREATE_MODE" in
    CreateOrMerge)
      addLabels "$owner" "$timestamp"
      kube apply -R -f "$MANIFEST_DIR"
      deleteByOldLabels "$owner" "$timestamp"
      ;;
    CreateOrUpdate)
      kube create -R -f "$MANIFEST_DIR"
      kube update -R -f "$MANIFEST_DIR"
      ;;
    Patch)
      kube patch -R -f "$MANIFEST_DIR"
      ;;
  esac
}

if [ "$CREATE_MODE" == "None" ] || [ "$DELETE_MODE" == "None" ]; then
  echo "CREATE_MODE and/or DELETE_MODE is set to None; This means that the template processor already applied the resources. Skipping the Manage Resources step."
  exit 0
fi

echo "Managing Resources"
setContext
case "$ACTION" in
  create) createUpdateResources;;
  delete) deleteResources;;
esac


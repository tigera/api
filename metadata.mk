#################################################################################################
# This file contains Makefile configuration parameters and metadata for this branch.
#################################################################################################

# Calico toolchain versions and the calico/base image to use.
GO_BUILD_VER=1.25.5-llvm18.1.8-k8s1.34.2
RUST_BUILD_VER=1.91.1

# Calico Enterprise shipping images now builds on UBI 10. For Calico OSS to Enterprise merges,
# please don't downgrade the base image back to UBI 9.
CALICO_BASE_VER=ubi10-1765220429

# Env var to ACK Ginkgo deprecation warnings, may need updating with go-build.
ACK_GINKGO=ACK_GINKGO_DEPRECATIONS=1.16.5

# Version of Kubernetes to use for tests, rancher/kubectl, and kubectl binary release in
# confd and kube-controllers.
K8S_VERSION=v1.34.2

# Version of various tools used in the build and tests.
COREDNS_VERSION=1.5.2
CRANE_VERSION=v0.20.7
ETCD_VERSION=v3.5.24
GHR_VERSION=v0.17.0
GITHUB_CLI_VERSION=2.76.2
GOTESTSUM_VERSION=v1.12.3
HELM_VERSION=v3.16.4
KINDEST_NODE_VERSION=v1.34.2
KINDEST_NODE_VERSION_DUAL_TOR=v1.24.7
KIND_VERSION=v0.29.0

# This gets embedded into node as the Calico version, the Enterprise release
# is based off of. This should be updated everytime a new opensource Calico
# release is merged into node-private.
CALICO_VERSION=v3.31.0

# The Semaphore calico-private ID, used when making calls to the Semaphore API.
SEMAPHORE_PROJECT_ID=8a309869-f767-49dc-924f-fa927edbf657

# Directory containing the Banzai secrets bundle, which we use to find the license key.
ifneq ("$(wildcard $(HOME)/.banzai/secrets)","")
    SECRETS_PATH ?= $(HOME)/.banzai/secrets
else
    SECRETS_PATH ?= $(HOME)/secrets
endif

# Configuration for Semaphore/Github integration.  This needs to be set
# differently for a forked repo.
ORGANIZATION  ?= tigera
GIT_REPO      ?= calico-private

RELEASE_BRANCH_PREFIX ?=release-calient
DEV_TAG_SUFFIX        ?= calient-0.dev

# Part of the git remote that is common to git and HTTP representations.
# Used to auto-detect the right remote.
GIT_REPO_SLUG ?= $(ORGANIZATION)/$(GIT_REPO)

# Configure git to access repositories using SSH.
GIT_USE_SSH = true

# Configure private repos
EXTRA_DOCKER_ARGS += -e GOPRIVATE=github.com/tigera/*

# The version of BIRD to use for calico/node builds and confd tests.
BIRD_VERSION=v0.3.3-211-g9111ec3c

# DEV_REGISTRIES configures the container image registries which are built from this
# repository.
DEV_REGISTRIES ?= tigera

# RELEASE_REGISTRIES configures the container images registries which are published to
# as part of an official release.
RELEASE_REGISTRIES ?= quay.io/tigera

# Archive bucket
ARTIFACTS_BUCKET ?= tigera-public/ee

# The directory for windows image tarballs
WINDOWS_DIST = dist/windows

# FIXME: Use WINDOWS_HPC_VERSION and remove WINDOWS_VERSIONS when containerd v1.6 is EOL'd
# The Windows HPC container version used as base for Calico Windows images
WINDOWS_HPC_VERSION ?= v1.0.0
# The Windows versions used as base for Calico Windows images
WINDOWS_VERSIONS ?= ltsc2019 ltsc2022

# The CNI plugin and flannel code that will be cloned and rebuilt with this repo's go-build image
# whenever the cni-plugin image is created.
CNI_VERSION=master
FLANNEL_VERSION=main

# The bpftool image to use; this is the output of the https://github.com/projectcalico/bpftool repo.
BPFTOOL_IMAGE=calico/bpftool:v7.5.0

# The operator branch corresponding to this branch.
OPERATOR_BRANCH       ?= master
OPERATOR_ORGANIZATION ?= tigera
OPERATOR_GIT_REPO     ?= operator
# The manager branch corresponding to this branch.
MANAGER_BRANCH ?= master

# The libbpf version to use
LIBBPF_VERSION=v1.4.6

# quay.io expiry time for hashrelease/dev images
QUAY_EXPIRE_DAYS=90

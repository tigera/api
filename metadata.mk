#################################################################################################
# This file contains Makefile configuration parameters and metadata for this branch.
#################################################################################################

# The version of calico/go-build and calico/base to use.
GO_BUILD_VER=1.24.3-llvm18.1.8-k8s1.32.5
CALICO_BASE_VER=ubi8-1744398299
# TODO Remove once CALICO_BASE is updated to UBI9
CALICO_BASE_UBI9_VER=ubi9-1744398299

# Env var to ACK Ginkgo deprecation warnings, may need updating with go-build.
ACK_GINKGO=ACK_GINKGO_DEPRECATIONS=1.16.5

# Version of Kubernetes to use for tests, bitnami/kubectl, and kubectl binary release in
# compliance benchmarker, confd, and kube-controllers.
K8S_VERSION=v1.33.1

# Version of various tools used in the build and tests.
COREDNS_VERSION=1.5.2
ETCD_VERSION=v3.5.6
GHR_VERSION=v0.17.0
GITHUB_CLI_VERSION=2.65.0
HELM_VERSION=v3.16.4
KINDEST_NODE_VERSION=v1.31.4
KINDEST_NODE_VERSION_DUAL_TOR=v1.24.7
KIND_VERSION=v0.25.0

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
ORGANIZATION = tigera
GIT_REPO = calico-private

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

# The suffix added to development tags (and, by association, images)
DEV_TAG_SUFFIX ?= calient-0.dev

# RELEASE_REGISTRIES configures the container images registries which are published to
# as part of an official release.
RELEASE_REGISTRIES = quay.io/tigera

# The directory for windows image tarballs
WINDOWS_DIST = dist/windows

# FIXME: Use WINDOWS_HPC_VERSION and remove WINDOWS_VERSIONS when containerd v1.6 is EOL'd
# The Windows HPC container version used as base for Calico Windows images
WINDOWS_HPC_VERSION ?= v1.0.0
# The Windows versions used as base for Calico Windows images
WINDOWS_VERSIONS ?= 1809 ltsc2022

# The CNI plugin and flannel code that will be cloned and rebuilt with this repo's go-build image
# whenever the cni-plugin image is created.
CNI_VERSION=master
FLANNEL_VERSION=main

# THIRD_PARTY_REGISTRY configures the third-party registry that serves intermediate base image
# for some Calico Enterprise components. They are never released directly to public.
ifeq ($(SEMAPHORE_GIT_REF_TYPE), branch)
    # on master and release-calient branches
    THIRD_PARTY_REGISTRY=gcr.io/unique-caldron-775/cnx/tigera/third-party
else ifeq ($(SEMAPHORE_GIT_REF_TYPE), pull-request)
    # on pull requests
    THIRD_PARTY_REGISTRY=gcr.io/unique-caldron-775/third-party-ci
else
    THIRD_PARTY_REGISTRY=gcr.io/tigera-dev/third-party-ci
endif

# The bpftool image to use; this is the output of the https://github.com/projectcalico/bpftool repo.
BPFTOOL_IMAGE=calico/bpftool:v7.5.0

# The default branch for semaphore
DEFAULT_BRANCH_OVERRIDE ?= release-calient-v3.21
# Default branch prefix for release branches
RELEASE_BRANCH_PREFIX ?= release-calient
# The operator branch corresponding to this branch.
OPERATOR_BRANCH=release-v1.38

# The libbpf version to use
LIBBPF_VERSION=v1.4.6

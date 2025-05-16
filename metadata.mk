#################################################################################################
# This file contains Makefile configuration parameters and metadata for this branch.
#################################################################################################

# The version of github.com/projectcalico/go-build to use.
GO_BUILD_VER=v0.89

# Version of Kubernetes to use for tests, bitnami/kubectl, and kubectl binary release in
# compliance benchmarker, confd, and kube-controllers.
K8S_VERSION=v1.27.8

# Version of various tools used in the build and tests.
COREDNS_VERSION=1.5.2
ELASTIC_VERSION=7.17.14
ETCD_VERSION=v3.5.6
# FIXME upgrading to kindest/node newer than v1.24.7 causes Node/kind-cluster and sig-network conformance
# tests to timeout or fail.
KINDEST_NODE_VERSION=v1.24.7
PROTOC_VER=v0.1
UBI_VERSION=8.9

# Configuration for Semaphore integration.
ORGANIZATION=tigera

# The Semaphore calico-private ID, used when making calls to the Semaphore API.
SEMAPHORE_PROJECT_ID=8a309869-f767-49dc-924f-fa927edbf657

# Directory containing the Banzai secrets bundle, which we use to find the license key.
ifneq ("$(wildcard $(HOME)/.banzai/secrets)","")
    SECRETS_PATH ?= $(HOME)/.banzai/secrets
else
    SECRETS_PATH ?= $(HOME)/secrets
endif

# Configure git to access repositories using SSH.
GIT_USE_SSH = true

# Configure private repos
EXTRA_DOCKER_ARGS += -e GOPRIVATE=github.com/tigera/*

# The version of BIRD to use for calico/node builds and confd tests.
BIRD_VERSION=v0.3.3-202-g7a77fb73

# DEV_REGISTRIES configures the container image registries which are built from this
# repository.
DEV_REGISTRIES ?= tigera

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

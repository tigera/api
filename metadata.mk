#################################################################################################
# This file contains Makefile configuration parameters and metadata for this branch.
#################################################################################################

# The version of github.com/projectcalico/go-build to use.
GO_BUILD_VER = v0.82

# Version of Kubernetes and kindest/node to use for tests.
K8S_VERSION     = v1.24.7
# This is used for bitnami/kubectl and kubectl binary release in compliance benchmarker, confd, and kube-controllers.
KUBECTL_VERSION = v1.25.8

# Version of various tools used in the build and tests.
COREDNS_VERSION=1.5.2
ELASTIC_VERSION=7.17.9
ETCD_VERSION=v3.5.6
PROTOC_VER=v0.1
UBI_VERSION=8.7

# Configuration for Semaphore integration.
ORGANIZATION = tigera

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
EXTRA_DOCKER_ARGS += --init -e GOPRIVATE=github.com/tigera/*

# The version of BIRD to use for calico/node builds and confd tests.
BIRD_VERSION=v0.3.3-200-g245602b0

# DEV_REGISTRIES configures the container image registries which are built from this
# repository.
DEV_REGISTRIES ?= tigera

# RELEASE_REGISTIRES configures the container images registries which are published to 
# as part of an official release.
RELEASE_REGISTRIES = quay.io/tigera

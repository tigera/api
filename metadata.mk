#################################################################################################
# This file contains Makefile configuration parameters and metadata for this branch.
#################################################################################################

# The version of github.com/projectcalico/go-build to use.
GO_BUILD_VER = v0.72

# Version of Kubernetes to use for tests.
K8S_VERSION     = v1.23.3
# This is used for lachlanevenson/k8s-kubectl and kubectl binary release.
KUBECTL_VERSION = v1.23.2

# Version of various tools used in the build and tests.
COREDNS_VERSION=1.5.2
ETCD_VERSION=v3.5.1
PROTOC_VER=v0.1
UBI_VERSION=8.5

# Configuration for Semaphore integration.
ORGANIZATION = tigera

# The Semaphore calico-private ID, used when making calls to the Semaphore API.
SEMAPHORE_PROJECT_ID=8a309869-f767-49dc-924f-fa927edbf657

# Configure git to access repositories using SSH.
GIT_USE_SSH = true

# Conigure private repos
EXTRA_DOCKER_ARGS += --init -e GOPRIVATE=github.com/tigera/*

# The version of BIRD to use for calico/node builds and confd tests.
BIRD_VERSION=v0.3.3-188-g0196eee4

# DEV_REGISTRIES configures the container image registries which are built from this
# repository.
DEV_REGISTRIES ?= tigera

# RELEASE_REGISTIRES configures the container images registries which are published to 
# as part of an official release.
RELEASE_REGISTRIES = quay.io/tigera

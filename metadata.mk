#################################################################################################
# This file contains Makefile configuration parameters and metadata for this branch.
#################################################################################################
# The project Go version
GO_VERSION=1.26.4
# Version of Kubernetes to use for dependencies, tests, registry.k8s.io/kubectl, and kubectl binary release.
K8S_VERSION=v1.36.2
# The version of LLVM to use for go-build and calico/base images.
LLVM_VERSION=21.1.8

# Calico toolchain versions and the calico/base image to use.
GO_BUILD_VER=$(GO_VERSION)-llvm$(LLVM_VERSION)-k8s$(K8S_VERSION:v%=%)
RUST_BUILD_VER=1.96.0

# Calico Enterprise shipping images now builds on UBI 10. For Calico OSS to Enterprise merges,
# please don't downgrade the base image back to UBI 9.
CALICO_BASE_VER=ubi10-1781722225

# Version of various tools used in the build and tests.
COREDNS_VERSION=1.5.2
CRANE_VERSION=v0.21.6
ETCD_VERSION=v3.5.31
GHR_VERSION=v0.18.3
GITHUB_CLI_VERSION=2.94.0
GOTESTSUM_VERSION=v1.13.0
HELM_VERSION=v3.21.2
# KINDEST_NODE_VERSION is the Kubernetes version of the KIND cluster used in
# tests, and is deliberately held one minor behind K8S_VERSION: the KubeVirt
# live-migration tests deploy KubeVirt (tigera/kubevirt mockvirt-v1.8.1, i.e.
# KubeVirt 1.8), which only supports Kubernetes 1.33-1.35. On a 1.36 node image
# VMIs never leave the "Scheduled" phase and the suite times out. v1.35.5 is the
# 1.35.x node image shipped with KIND_VERSION below. Bump this only once a
# KubeVirt/mockvirt release that supports the target Kubernetes minor exists.
KINDEST_NODE_VERSION=v1.35.5
KINDEST_NODE_VERSION_DUAL_TOR=v1.24.7
KIND_VERSION=v0.32.0

# This gets embedded into node as the Calico version, the Enterprise release
# is based off of. This should be updated everytime a new opensource Calico
# release is merged into node-private.
CALICO_VERSION=v3.33.0

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
GIT_REMOTE    ?= origin

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

# ALLOWED_DEV_REGISTRIES is the prefix allowlist enforced on every DEV_REGISTRIES
# entry at push time (see validate-dev-registries in lib.Makefile). Entries in
# DEV_REGISTRIES must start with one of these prefixes or the push fails. The
# intent is to keep every dev and CD push inside Tigera-controlled registries
# even when DEV_REGISTRIES is overridden ad-hoc, so that a misconfigured CI job
# or a slipped command line cannot publish Enterprise images to a public OSS
# registry path.
ALLOWED_DEV_REGISTRIES ?= gcr.io/unique-caldron-775 quay.io/tigeradev

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

# The CNI plugin and flannel code that will be cloned and rebuilt with this repo's go-build image.
# Pinned so the content-addressed third-party-cni-plugins image hash changes when these move.
# CNI_VERSION is a commit SHA because the fork has no release tag at the toolchain we build with;
# bump it to pick up upstream changes.
CNI_VERSION=9ffe547cb3b66f80dd32a00fc69a6d0082b55321
FLANNEL_VERSION=v1.2.0-flannel2-go1.22.7

# The libbpf version to use
LIBBPF_VERSION=v1.6.2

# The bpftool image to use; this is the output of the https://github.com/projectcalico/bpftool repo.
BPFTOOL_IMAGE=calico/bpftool:v7.5.0

# Patched nftables + libnftnl RPMs shipped in calico/node and the istio CNI
# install image, pulled as calico/nftables-rpms:$(NFT_RPMS_TAG)-<arch>. Built
# and published by the OSS calico repo from hack/rpms/nftables/, which is also
# where the patched (GPL) source lives - we deliberately carry no copy of it.
# To bump: land the change in OSS first, then update this pin to the output of
# `make -C hack/rpms/nftables print-tag` in an OSS checkout.
NFT_RPMS_TAG=6eb1d57c0512

# The operator branch corresponding to this branch.
OPERATOR_BRANCH ?= master
OPERATOR_ORGANIZATION ?= tigera
OPERATOR_GIT_REPO     ?= operator
# The manager branch corresponding to this branch.
MANAGER_BRANCH ?= master

# quay.io expiry time for hashrelease/dev images
QUAY_EXPIRE_DAYS=90

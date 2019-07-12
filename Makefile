MY_UID:=$(shell id -u)
ifdef SSH_AUTH_SOCK
  EXTRA_DOCKER_ARGS += -v $(SSH_AUTH_SOCK):/ssh-agent --env SSH_AUTH_SOCK=/ssh-agent
endif
PACKAGE_NAME?=github.com/tigera/api

GO_BUILD_VER?=v0.20
# For building, we use the go-build image for the *host* architecture, even if the target is different
# the one for the host should contain all the necessary cross-compilation tools
# we do not need to use the arch since go-build:v0.15 now is multi-arch manifest
CALICO_BUILD=calico/go-build:$(GO_BUILD_VER)

BINDIR        ?= bin
BUILD_DIR     ?= build

DOCKER_RUN := mkdir -p .go-pkg-cache && \
			  docker run --rm \
				 --net=host \
				 $(EXTRA_DOCKER_ARGS) \
				 -e LOCAL_USER_ID=$(MY_UID) \
				 -v $${PWD}:/go/src/github.com/tigera/api:rw \
				 -v $${PWD}/.go-pkg-cache:/go/pkg:rw \
				 -v $${PWD}/hack/boilerplate:/go/src/k8s.io/kubernetes/hack/boilerplate:rw \
				 -w /go/src/github.com/tigera/api \
				 -e GOARCH=$(ARCH) 

# Update the vendored dependencies with the latest upstream versions matching
# our glide.yaml.  If there area any changes, this updates glide.lock
# as a side effect.  Unless you're adding/updating a dependency, you probably
# want to use the vendor target to install the versions from glide.lock.
.PHONY: update-vendor
update-vendor:
	mkdir -p $$HOME/.glide
	$(DOCKER_RUN) $(CALICO_BUILD) glide up --strip-vendor
	touch vendor/.up-to-date

# vendor is a shortcut for force rebuilding the go vendor directory.
.PHONY: vendor
vendor vendor/.up-to-date: glide.lock
	mkdir -p $$HOME/.glide
	$(DOCKER_RUN) $(CALICO_BUILD) glide install --strip-vendor
	touch vendor/.up-to-date

###############################################################################
## Set the default upstream repo branch to the current repo's branch,
## e.g. "master" or "release-vX.Y", but allow it to be overridden.
PIN_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)

# This section contains the code generation stuff
#################################################
.generate_exes: $(BINDIR)/deepcopy-gen \
                $(BINDIR)/client-gen \
                $(BINDIR)/lister-gen \
                $(BINDIR)/informer-gen
	touch $@

$(BINDIR)/deepcopy-gen: vendor/.up-to-date
	$(DOCKER_RUN) $(CALICO_BUILD) \
	    sh -c 'go build -o $@ $(PACKAGE_NAME)/vendor/k8s.io/code-generator/cmd/deepcopy-gen'

$(BINDIR)/client-gen: vendor/.up-to-date
	$(DOCKER_RUN) $(CALICO_BUILD) \
	    sh -c 'go build -o $@ $(PACKAGE_NAME)/vendor/k8s.io/code-generator/cmd/client-gen'

$(BINDIR)/lister-gen: vendor/.up-to-date
	$(DOCKER_RUN) $(CALICO_BUILD) \
	    sh -c 'go build -o $@ $(PACKAGE_NAME)/vendor/k8s.io/code-generator/cmd/lister-gen'

$(BINDIR)/informer-gen: vendor/.up-to-date
	$(DOCKER_RUN) $(CALICO_BUILD) \
	    sh -c 'go build -o $@ $(PACKAGE_NAME)/vendor/k8s.io/code-generator/cmd/informer-gen'

# Regenerate all files if the gen exes changed or any "types.go" files changed
.generate_files: .generate_exes $(TYPES_FILES)
	# Generate deep copies
	$(DOCKER_RUN) $(CALICO_BUILD) \
	   sh -c '$(BINDIR)/deepcopy-gen \
		--v 1 --logtostderr \
		--go-header-file "/go/src/$(PACKAGE_NAME)/hack/boilerplate/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/projectcalico" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/projectcalico/v3" \
		--bounding-dirs "github.com/tigera/api" \
		--output-file-base zz_generated.deepcopy'
	# Generate all pkg/client contents
	$(DOCKER_RUN) $(CALICO_BUILD) \
	   sh -c '$(BUILD_DIR)/update-client-gen.sh'
	touch $@

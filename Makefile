#!make
SHELL = /bin/bash

# Standard colors
BLACK := $(shell tput -Txterm setaf 0)
RED := $(shell tput -Txterm setaf 1)
GREEN := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
LIGHTPURPLE := $(shell tput -Txterm setaf 4)
PURPLE := $(shell tput -Txterm setaf 5)
BLUE := $(shell tput -Txterm setaf 6)
WHITE := $(shell tput -Txterm setaf 7)

RESET := $(shell tput -Txterm sgr0)

# Prints the output in colored text
#   $1: the color from the above list
#   $2: the text value
define colored
	@echo "$1$2${RESET}"
endef

define cleanup
	make cleanup
endef

# Retrieve git repo name
GIT_REPO_NAME := $(shell basename `git rev-parse --show-toplevel`)

# Retrieve git branch
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

# Retrieve the total number of commits in the git branch
COMMIT_NUMBER := $(shell git rev-list --count $(GIT_BRANCH))

# Define new version based on number of commits
# E.g: Num of commits = 27, then the new version = v1.0.0-build.27
MAJOR := 1
MINOR := 0
PATCH := 0
NEW_TAG_VERSION := v$(MAJOR).$(MINOR).$(PATCH)-build.$(COMMIT_NUMBER)

# Get the latest tag
CURRENT_TAG_VERSION := $(shell git describe --abbrev=0 --tags 2>/dev/null)

# Get the latest commit id
LATEST_GIT_COMMIT := $(shell git rev-parse HEAD)

# Initialize location to install binaries
# Binaries will be installed at location "./autobots/bin/"
INSTALL_BASE_DIR := ./autobots
BINARIES_BASE_PATH := $(INSTALL_BASE_DIR)/bin
GORELEASER_BINARY := $(BINARIES_BASE_PATH)/goreleaser

# Changelog details
CHANGELOG_DIR := ./
CHANGELOG_FILE_NAME := CHANGELOG.md

local: install-go-releaser
	@echo "Running local build... "
	$(GORELEASER_BINARY) --snapshot --skip-publish --rm-dist
	$(call cleanup)
	$(call colored,"${GREEN}","Successfully build pkg locally")


all: tag github-release artifactory-release cleanup success


tag:
# If the new version and current version are not equal, then create and push the new tag version
ifneq ($(strip $(NEW_TAG_VERSION)),$(strip $(CURRENT_TAG_VERSION)))
ifeq ($(strip $(CURRENT_TAG_VERSION)),)
	$(call colored,"${BLUE}","Creating first tag with version $(NEW_TAG_VERSION)")
else
	$(call colored,"${BLUE}","Updating the version from $(CURRENT_TAG_VERSION) to $(NEW_TAG_VERSION)")
endif
# Create and push the new tag
	git tag -a $(NEW_TAG_VERSION) -m ''
	git push --tag
	$(call colored,"${BLUE}","Tagged New Version: $(NEW_TAG_VERSION) with commit id: $(LATEST_GIT_COMMIT)")
else
	$(call colored,"${YELLOW}","No new commits to create a new tag version... ")
	$(call colored,"${YELLOW}","Current Tag Version: $(CURRENT_TAG_VERSION) with commit id: $(LATEST_GIT_COMMIT)")
endif


build-env:
# Generate required directories to install packages
	@echo "Creating required directories to install packages... "
	$(shell mkdir -p $(BINARIES_BASE_PATH))


install-go-releaser: build-env
# Install goreleaser cli
	@echo "Installing goreleaser cli... "
	cd $(INSTALL_BASE_DIR) && curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
	$(GORELEASER_BINARY) --version


github-release-checks: install-go-releaser
	@echo "Initiating the pre build and release checks... "
# Check env var GITHUB_TOKEN. Required by goreleaser
ifeq ($(strip $(GITHUB_TOKEN)),)
	$(call colored,"${RED}","Env var GITHUB_TOKEN is missing or empty. Export using command: export GITHUB_TOKEN=\<git_token\>")
	exit 1
endif
# Check .goreleaser.yml file is valid
	$(GORELEASER_BINARY) check


github-release: github-release-checks
# Create release using GoReleaser and upload to GitHub
	@echo "Initiating the build and release process... "
	$(GORELEASER_BINARY) release --release-notes $(CHANGELOG_DIR)$(CHANGELOG_FILE_NAME) --rm-dist


cleanup:
	@echo "Initiating the cleanup process... "
	rm -rf $(INSTALL_BASE_DIR)


success:
	$(call colored,"${GREEN}","Successfully build and released New Version: $(NEW_TAG_VERSION) with commit id: $(LATEST_GIT_COMMIT)")

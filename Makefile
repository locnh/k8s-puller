.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

HELM         := helm
HELM_VERSION := $(shell $(HELM) version --short)
HELM_SCRIPTS := $(CURDIR)/scripts

.PHONY: release
release: test package index

.PHONY: index
index:
	@sh -c "'$(HELM_SCRIPTS)/helm-index.sh'"

.PHONY: package
package:
	@sh -c "'$(HELM_SCRIPTS)/helm-package.sh'"

.PHONY: test
test:
	@sh -c "'$(HELM_SCRIPTS)/helm-test.sh'"

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

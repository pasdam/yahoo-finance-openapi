OPENAPI_VERSION := 5.2.0

.PHONY: clean
clean: | clean-go

.PHONY: clean-go
clean-go:
	@rm -rf ./gen/go/yq*/*.go \
					./gen/go/yq*/*.md \
					./gen/go/yq*/.openapi-generator \
					./gen/go/yq*/api \
					./gen/go/yq*/docs \
					./gen/go/yq*/go.*

.PHONY: generate-go-sdk
generate-go-sdk: | clean
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v${OPENAPI_VERSION} \
		generate -i /workspace/query1.yml -g go -o /workspace/gen/go/yq1 -c /workspace/query1-config-go.yml --git-repo-id yahoo-finance-openapi/gen/go --git-user-id pasdam
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v${OPENAPI_VERSION} \
		generate -i /workspace/query2.yml -g go -o /workspace/gen/go/yq2 -c /workspace/query2-config-go.yml --git-repo-id yahoo-finance-openapi/gen/go --git-user-id pasdam

.PHONY: version
version:
	@cat query1.yml | grep -e "^  version: " | sed "s/  version: //"

.PHONY: check-version-bump
check-version-bump:
	@$(eval VER1 := $(shell cat query1.yml | grep -e "^  version: " | sed "s/  version: //"))
	@$(eval VER2 := $(shell cat query2.yml | grep -e "^  version: " | sed "s/  version: //"))
	@if [ "${VER1}" != "${VER2}" ]; then\
		1>&2 echo "Version mismatch, please make sure API version in query1.yml matches the one in query2.yml";\
		exit 1;\
	fi
	@$(eval LATEST_TAG := $(shell git fetch && git describe --tags $(shell git rev-list --tags --max-count=1) 2> /dev/null | sed "s/^v//"))
	@if [ "${VER1}" = "${LATEST_TAG}" ]; then\
		1>&2 echo "Please bump the API version in the api.yml OpenApi spec";\
	 	exit 2;\
	fi

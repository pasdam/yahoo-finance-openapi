.PHONY: clean
clean: | clean-go

.PHONY: clean-go
clean-go:
	@rm -rf ./gen/go/yf/*.go
	@rm -rf ./gen/go/yf/*.md

.PHONY: generate-go-client
generate-go-client: | clean
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v5.1.1 \
		generate -i /workspace/api.yml -g go -o /workspace/gen/go/yf -c /workspace/config-go.yml --git-repo-id yahoo-finance-openapi/gen/go --git-user-id pasdam

.PHONY: generate-go-client
generate-go-client:
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v5.1.1 \
		generate -i /workspace/api.yml -g go -o /workspace/gen/go/yf -c /workspace/config-go.yml

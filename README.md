# yahoo-finance-openapi

This repository contains the unofficial [OpenAPI specification](./api.yml) for the Yahoo Finance APIs.

From the manifest the following clients are automatically generated using [OpenAPI-Generator](https://openapi-generator.tech/):

* [go](./gen/go/yf).

## Contribute

### Add API

To add an API simply update the manifest with the new endpoint(s) and related data models.

### Add client

To add a new client for a new language simply add a job in the [CI workflow](./.github/workflows/generate-clients.yaml) to:

1. generate the code in the required language;
2. build the library artifact;
3. publish the artifact.

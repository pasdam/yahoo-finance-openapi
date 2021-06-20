module github.com/pasdam/yahoo-finance-openapi/go/test

replace github.com/pasdam/yahoo-finance-openapi => ../../

go 1.13

require (
	github.com/pasdam/yahoo-finance-openapi/gen/go/yf v0.0.0-20210619154829-1c47d252b851
	github.com/stretchr/testify v1.7.0
)

/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SparkApiService SparkApi service
type SparkApiService service

type ApiSparkRequest struct {
	ctx _context.Context
	ApiService *SparkApiService
	interval *Interval
	range_ *Range
	symbols *string
	lang *string
	includePrePost *bool
	includeTimestamps *bool
	indicators *string
	corsDomain *string
}

func (r ApiSparkRequest) Interval(interval Interval) ApiSparkRequest {
	r.interval = &interval
	return r
}
func (r ApiSparkRequest) Range_(range_ Range) ApiSparkRequest {
	r.range_ = &range_
	return r
}
func (r ApiSparkRequest) Symbols(symbols string) ApiSparkRequest {
	r.symbols = &symbols
	return r
}
func (r ApiSparkRequest) Lang(lang string) ApiSparkRequest {
	r.lang = &lang
	return r
}
func (r ApiSparkRequest) IncludePrePost(includePrePost bool) ApiSparkRequest {
	r.includePrePost = &includePrePost
	return r
}
func (r ApiSparkRequest) IncludeTimestamps(includeTimestamps bool) ApiSparkRequest {
	r.includeTimestamps = &includeTimestamps
	return r
}
func (r ApiSparkRequest) Indicators(indicators string) ApiSparkRequest {
	r.indicators = &indicators
	return r
}
func (r ApiSparkRequest) CorsDomain(corsDomain string) ApiSparkRequest {
	r.corsDomain = &corsDomain
	return r
}

func (r ApiSparkRequest) Execute() (QuoteResponse, *_nethttp.Response, error) {
	return r.ApiService.SparkExecute(r)
}

/*
 * Spark Method for Spark
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSparkRequest
 */
func (a *SparkApiService) Spark(ctx _context.Context) ApiSparkRequest {
	return ApiSparkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return QuoteResponse
 */
func (a *SparkApiService) SparkExecute(r ApiSparkRequest) (QuoteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SparkApiService.Spark")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v7/finance/spark"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.range_ == nil {
		return localVarReturnValue, nil, reportError("range_ is required and must be specified")
	}
	if r.symbols == nil {
		return localVarReturnValue, nil, reportError("symbols is required and must be specified")
	}

	localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	localVarQueryParams.Add("range", parameterToString(*r.range_, ""))
	if r.lang != nil {
		localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	}
	if r.includePrePost != nil {
		localVarQueryParams.Add("includePrePost", parameterToString(*r.includePrePost, ""))
	}
	if r.includeTimestamps != nil {
		localVarQueryParams.Add("includeTimestamps", parameterToString(*r.includeTimestamps, ""))
	}
	if r.indicators != nil {
		localVarQueryParams.Add("indicators", parameterToString(*r.indicators, ""))
	}
	if r.corsDomain != nil {
		localVarQueryParams.Add("corsDomain", parameterToString(*r.corsDomain, ""))
	}
	localVarQueryParams.Add("symbols", parameterToString(*r.symbols, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

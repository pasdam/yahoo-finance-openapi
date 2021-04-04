/*
 * Yahoo Finance
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiV7FinanceQuoteGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	symbols *string
	formatted *bool
	region *string
	lang *string
	includePrePost *bool
	fields *string
	corsDomain *string
}

func (r ApiV7FinanceQuoteGetRequest) Symbols(symbols string) ApiV7FinanceQuoteGetRequest {
	r.symbols = &symbols
	return r
}
func (r ApiV7FinanceQuoteGetRequest) Formatted(formatted bool) ApiV7FinanceQuoteGetRequest {
	r.formatted = &formatted
	return r
}
func (r ApiV7FinanceQuoteGetRequest) Region(region string) ApiV7FinanceQuoteGetRequest {
	r.region = &region
	return r
}
func (r ApiV7FinanceQuoteGetRequest) Lang(lang string) ApiV7FinanceQuoteGetRequest {
	r.lang = &lang
	return r
}
func (r ApiV7FinanceQuoteGetRequest) IncludePrePost(includePrePost bool) ApiV7FinanceQuoteGetRequest {
	r.includePrePost = &includePrePost
	return r
}
func (r ApiV7FinanceQuoteGetRequest) Fields(fields string) ApiV7FinanceQuoteGetRequest {
	r.fields = &fields
	return r
}
func (r ApiV7FinanceQuoteGetRequest) CorsDomain(corsDomain string) ApiV7FinanceQuoteGetRequest {
	r.corsDomain = &corsDomain
	return r
}

func (r ApiV7FinanceQuoteGetRequest) Execute() (QuoteResponse, *_nethttp.Response, error) {
	return r.ApiService.V7FinanceQuoteGetExecute(r)
}

/*
 * V7FinanceQuoteGet Returns a list of users.
 * Optional extended description in CommonMark or HTML.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiV7FinanceQuoteGetRequest
 */
func (a *DefaultApiService) V7FinanceQuoteGet(ctx _context.Context) ApiV7FinanceQuoteGetRequest {
	return ApiV7FinanceQuoteGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return QuoteResponse
 */
func (a *DefaultApiService) V7FinanceQuoteGetExecute(r ApiV7FinanceQuoteGetRequest) (QuoteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.V7FinanceQuoteGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v7/finance/quote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.symbols == nil {
		return localVarReturnValue, nil, reportError("symbols is required and must be specified")
	}

	if r.formatted != nil {
		localVarQueryParams.Add("formatted", parameterToString(*r.formatted, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.lang != nil {
		localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	}
	if r.includePrePost != nil {
		localVarQueryParams.Add("includePrePost", parameterToString(*r.includePrePost, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type ApiV8FinanceChartSymbolGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	symbol string
	interval *Interval
	range_ *Range
	region *string
	includePrePost *bool
	lang *string
	useYfid *bool
	corsDomain *string
	tsrc *string
}

func (r ApiV8FinanceChartSymbolGetRequest) Interval(interval Interval) ApiV8FinanceChartSymbolGetRequest {
	r.interval = &interval
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) Range_(range_ Range) ApiV8FinanceChartSymbolGetRequest {
	r.range_ = &range_
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) Region(region string) ApiV8FinanceChartSymbolGetRequest {
	r.region = &region
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) IncludePrePost(includePrePost bool) ApiV8FinanceChartSymbolGetRequest {
	r.includePrePost = &includePrePost
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) Lang(lang string) ApiV8FinanceChartSymbolGetRequest {
	r.lang = &lang
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) UseYfid(useYfid bool) ApiV8FinanceChartSymbolGetRequest {
	r.useYfid = &useYfid
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) CorsDomain(corsDomain string) ApiV8FinanceChartSymbolGetRequest {
	r.corsDomain = &corsDomain
	return r
}
func (r ApiV8FinanceChartSymbolGetRequest) Tsrc(tsrc string) ApiV8FinanceChartSymbolGetRequest {
	r.tsrc = &tsrc
	return r
}

func (r ApiV8FinanceChartSymbolGetRequest) Execute() (Chart, *_nethttp.Response, error) {
	return r.ApiService.V8FinanceChartSymbolGetExecute(r)
}

/*
 * V8FinanceChartSymbolGet Returns a list of users.
 * Optional extended description in CommonMark or HTML.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol Pet object that needs to be added to the store
 * @return ApiV8FinanceChartSymbolGetRequest
 */
func (a *DefaultApiService) V8FinanceChartSymbolGet(ctx _context.Context, symbol string) ApiV8FinanceChartSymbolGetRequest {
	return ApiV8FinanceChartSymbolGetRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

/*
 * Execute executes the request
 * @return Chart
 */
func (a *DefaultApiService) V8FinanceChartSymbolGetExecute(r ApiV8FinanceChartSymbolGetRequest) (Chart, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Chart
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.V8FinanceChartSymbolGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v8/finance/chart/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", _neturl.PathEscape(parameterToString(r.symbol, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.range_ == nil {
		return localVarReturnValue, nil, reportError("range_ is required and must be specified")
	}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.includePrePost != nil {
		localVarQueryParams.Add("includePrePost", parameterToString(*r.includePrePost, ""))
	}
	if r.lang != nil {
		localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	}
	localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	if r.useYfid != nil {
		localVarQueryParams.Add("useYfid", parameterToString(*r.useYfid, ""))
	}
	localVarQueryParams.Add("range", parameterToString(*r.range_, ""))
	if r.corsDomain != nil {
		localVarQueryParams.Add("corsDomain", parameterToString(*r.corsDomain, ""))
	}
	if r.tsrc != nil {
		localVarQueryParams.Add(".tsrc", parameterToString(*r.tsrc, ""))
	}
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

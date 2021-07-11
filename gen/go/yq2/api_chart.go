/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq2

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

// ChartApiService ChartApi service
type ChartApiService service

type ApiGetChartRequest struct {
	ctx _context.Context
	ApiService *ChartApiService
	symbol string
	formatted *bool
	lang *string
	region *string
	modules *[]string
	corsDomain *string
	crumb *string
}

func (r ApiGetChartRequest) Formatted(formatted bool) ApiGetChartRequest {
	r.formatted = &formatted
	return r
}
func (r ApiGetChartRequest) Lang(lang string) ApiGetChartRequest {
	r.lang = &lang
	return r
}
func (r ApiGetChartRequest) Region(region string) ApiGetChartRequest {
	r.region = &region
	return r
}
func (r ApiGetChartRequest) Modules(modules []string) ApiGetChartRequest {
	r.modules = &modules
	return r
}
func (r ApiGetChartRequest) CorsDomain(corsDomain string) ApiGetChartRequest {
	r.corsDomain = &corsDomain
	return r
}
func (r ApiGetChartRequest) Crumb(crumb string) ApiGetChartRequest {
	r.crumb = &crumb
	return r
}

func (r ApiGetChartRequest) Execute() (QuoteSummary, *_nethttp.Response, error) {
	return r.ApiService.GetChartExecute(r)
}

/*
 * GetChart Method for GetChart
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @return ApiGetChartRequest
 */
func (a *ChartApiService) GetChart(ctx _context.Context, symbol string) ApiGetChartRequest {
	return ApiGetChartRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

/*
 * Execute executes the request
 * @return QuoteSummary
 */
func (a *ChartApiService) GetChartExecute(r ApiGetChartRequest) (QuoteSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuoteSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartApiService.GetChart")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v10/finance/quoteSummary/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", _neturl.PathEscape(parameterToString(r.symbol, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.formatted == nil {
		return localVarReturnValue, nil, reportError("formatted is required and must be specified")
	}
	if r.lang == nil {
		return localVarReturnValue, nil, reportError("lang is required and must be specified")
	}
	if r.region == nil {
		return localVarReturnValue, nil, reportError("region is required and must be specified")
	}
	if r.modules == nil {
		return localVarReturnValue, nil, reportError("modules is required and must be specified")
	}
	if r.corsDomain == nil {
		return localVarReturnValue, nil, reportError("corsDomain is required and must be specified")
	}

	localVarQueryParams.Add("formatted", parameterToString(*r.formatted, ""))
	if r.crumb != nil {
		localVarQueryParams.Add("crumb", parameterToString(*r.crumb, ""))
	}
	localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	localVarQueryParams.Add("modules", parameterToString(*r.modules, "csv"))
	localVarQueryParams.Add("corsDomain", parameterToString(*r.corsDomain, ""))
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

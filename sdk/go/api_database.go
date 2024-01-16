/*
Vearch Database API

API for sending dynamic records to the Vearch database.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vearch_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DatabaseAPIService DatabaseAPI service
type DatabaseAPIService service

type ApiCreateDatabaseRequest struct {
	ctx context.Context
	ApiService *DatabaseAPIService
	createDatabaseRequest *CreateDatabaseRequest
}

func (r ApiCreateDatabaseRequest) CreateDatabaseRequest(createDatabaseRequest CreateDatabaseRequest) ApiCreateDatabaseRequest {
	r.createDatabaseRequest = &createDatabaseRequest
	return r
}

func (r ApiCreateDatabaseRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateDatabaseExecute(r)
}

/*
CreateDatabase Create a new database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDatabaseRequest
*/
func (a *DatabaseAPIService) CreateDatabase(ctx context.Context) ApiCreateDatabaseRequest {
	return ApiCreateDatabaseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DatabaseAPIService) CreateDatabaseExecute(r ApiCreateDatabaseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseAPIService.CreateDatabase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/db/_create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createDatabaseRequest == nil {
		return nil, reportError("createDatabaseRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createDatabaseRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteDBRequest struct {
	ctx context.Context
	ApiService *DatabaseAPIService
	dBNAME string
}

func (r ApiDeleteDBRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDBExecute(r)
}

/*
DeleteDB Delete a specific db

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dBNAME The name of the database
 @return ApiDeleteDBRequest
*/
func (a *DatabaseAPIService) DeleteDB(ctx context.Context, dBNAME string) ApiDeleteDBRequest {
	return ApiDeleteDBRequest{
		ApiService: a,
		ctx: ctx,
		dBNAME: dBNAME,
	}
}

// Execute executes the request
func (a *DatabaseAPIService) DeleteDBExecute(r ApiDeleteDBRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseAPIService.DeleteDB")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/db/{DB_NAME}"
	localVarPath = strings.Replace(localVarPath, "{"+"DB_NAME"+"}", url.PathEscape(parameterValueToString(r.dBNAME, "dBNAME")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

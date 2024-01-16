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


// SpaceAPIService SpaceAPI service
type SpaceAPIService service

type ApiCreateSpaceRequest struct {
	ctx context.Context
	ApiService *SpaceAPIService
	dBNAME string
	createSpaceRequest *CreateSpaceRequest
}

func (r ApiCreateSpaceRequest) CreateSpaceRequest(createSpaceRequest CreateSpaceRequest) ApiCreateSpaceRequest {
	r.createSpaceRequest = &createSpaceRequest
	return r
}

func (r ApiCreateSpaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateSpaceExecute(r)
}

/*
CreateSpace Create a new space

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dBNAME The name of the database where the space will be created.
 @return ApiCreateSpaceRequest
*/
func (a *SpaceAPIService) CreateSpace(ctx context.Context, dBNAME string) ApiCreateSpaceRequest {
	return ApiCreateSpaceRequest{
		ApiService: a,
		ctx: ctx,
		dBNAME: dBNAME,
	}
}

// Execute executes the request
func (a *SpaceAPIService) CreateSpaceExecute(r ApiCreateSpaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpaceAPIService.CreateSpace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/space/{DB_NAME}/_create"
	localVarPath = strings.Replace(localVarPath, "{"+"DB_NAME"+"}", url.PathEscape(parameterValueToString(r.dBNAME, "dBNAME")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSpaceRequest == nil {
		return nil, reportError("createSpaceRequest is required and must be specified")
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
	localVarPostBody = r.createSpaceRequest
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

type ApiDeleteSpaceRequest struct {
	ctx context.Context
	ApiService *SpaceAPIService
	dBNAME string
	sPACENAME string
}

func (r ApiDeleteSpaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSpaceExecute(r)
}

/*
DeleteSpace Delete a specific space

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dBNAME The name of the database
 @param sPACENAME The name of the space to delete
 @return ApiDeleteSpaceRequest
*/
func (a *SpaceAPIService) DeleteSpace(ctx context.Context, dBNAME string, sPACENAME string) ApiDeleteSpaceRequest {
	return ApiDeleteSpaceRequest{
		ApiService: a,
		ctx: ctx,
		dBNAME: dBNAME,
		sPACENAME: sPACENAME,
	}
}

// Execute executes the request
func (a *SpaceAPIService) DeleteSpaceExecute(r ApiDeleteSpaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpaceAPIService.DeleteSpace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/space/{DB_NAME}/{SPACE_NAME}"
	localVarPath = strings.Replace(localVarPath, "{"+"DB_NAME"+"}", url.PathEscape(parameterValueToString(r.dBNAME, "dBNAME")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"SPACE_NAME"+"}", url.PathEscape(parameterValueToString(r.sPACENAME, "sPACENAME")), -1)

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

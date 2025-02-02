// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EmailTemplateClient contains the methods for the EmailTemplate group.
// Don't use this type directly, use NewEmailTemplateClient() instead.
type EmailTemplateClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewEmailTemplateClient creates a new instance of EmailTemplateClient with the specified values.
func NewEmailTemplateClient(con *armcore.Connection, subscriptionID string) *EmailTemplateClient {
	return &EmailTemplateClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Updates an Email Template.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmailTemplateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters, options *EmailTemplateCreateOrUpdateOptions) (EmailTemplateCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, templateName, parameters, options)
	if err != nil {
		return EmailTemplateCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return EmailTemplateCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EmailTemplateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters, options *EmailTemplateCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EmailTemplateClient) createOrUpdateHandleResponse(resp *azcore.Response) (EmailTemplateCreateOrUpdateResponse, error) {
	result := EmailTemplateCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.EmailTemplateContract); err != nil {
		return EmailTemplateCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *EmailTemplateClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Reset the Email Template to default template provided by the API Management service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmailTemplateClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, options *EmailTemplateDeleteOptions) (EmailTemplateDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, templateName, ifMatch, options)
	if err != nil {
		return EmailTemplateDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return EmailTemplateDeleteResponse{}, client.deleteHandleError(resp)
	}
	return EmailTemplateDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EmailTemplateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, options *EmailTemplateDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *EmailTemplateClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the details of the email template specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmailTemplateClient) Get(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateGetOptions) (EmailTemplateGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, templateName, options)
	if err != nil {
		return EmailTemplateGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EmailTemplateGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EmailTemplateClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmailTemplateClient) getHandleResponse(resp *azcore.Response) (EmailTemplateGetResponse, error) {
	result := EmailTemplateGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.EmailTemplateContract); err != nil {
		return EmailTemplateGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EmailTemplateClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Gets the entity state (Etag) version of the email template specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmailTemplateClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateGetEntityTagOptions) (EmailTemplateGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, templateName, options)
	if err != nil {
		return EmailTemplateGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *EmailTemplateClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *EmailTemplateClient) getEntityTagHandleResponse(resp *azcore.Response) (EmailTemplateGetEntityTagResponse, error) {
	result := EmailTemplateGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Gets all email templates
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmailTemplateClient) ListByService(resourceGroupName string, serviceName string, options *EmailTemplateListByServiceOptions) EmailTemplateListByServicePager {
	return &emailTemplateListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp EmailTemplateListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EmailTemplateCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *EmailTemplateClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EmailTemplateListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *EmailTemplateClient) listByServiceHandleResponse(resp *azcore.Response) (EmailTemplateListByServiceResponse, error) {
	result := EmailTemplateListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.EmailTemplateCollection); err != nil {
		return EmailTemplateListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *EmailTemplateClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Updates API Management email template
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmailTemplateClient) Update(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, parameters EmailTemplateUpdateParameters, options *EmailTemplateUpdateOptions) (EmailTemplateUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, templateName, ifMatch, parameters, options)
	if err != nil {
		return EmailTemplateUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EmailTemplateUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EmailTemplateClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, parameters EmailTemplateUpdateParameters, options *EmailTemplateUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *EmailTemplateClient) updateHandleResponse(resp *azcore.Response) (EmailTemplateUpdateResponse, error) {
	result := EmailTemplateUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.EmailTemplateContract); err != nil {
		return EmailTemplateUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *EmailTemplateClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

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

// IssueClient contains the methods for the Issue group.
// Don't use this type directly, use NewIssueClient() instead.
type IssueClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewIssueClient creates a new instance of IssueClient with the specified values.
func NewIssueClient(con *armcore.Connection, subscriptionID string) *IssueClient {
	return &IssueClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets API Management issue details
// If the operation fails it returns the *ErrorResponse error type.
func (client *IssueClient) Get(ctx context.Context, resourceGroupName string, serviceName string, issueID string, options *IssueGetOptions) (IssueGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, issueID, options)
	if err != nil {
		return IssueGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IssueGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IssueGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IssueClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, issueID string, options *IssueGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues/{issueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
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
func (client *IssueClient) getHandleResponse(resp *azcore.Response) (IssueGetResponse, error) {
	result := IssueGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.IssueContract); err != nil {
		return IssueGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IssueClient) getHandleError(resp *azcore.Response) error {
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

// ListByService - Lists a collection of issues in the specified service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IssueClient) ListByService(resourceGroupName string, serviceName string, options *IssueListByServiceOptions) IssueListByServicePager {
	return &issueListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp IssueListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.IssueCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *IssueClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *IssueListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues"
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
func (client *IssueClient) listByServiceHandleResponse(resp *azcore.Response) (IssueListByServiceResponse, error) {
	result := IssueListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.IssueCollection); err != nil {
		return IssueListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *IssueClient) listByServiceHandleError(resp *azcore.Response) error {
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
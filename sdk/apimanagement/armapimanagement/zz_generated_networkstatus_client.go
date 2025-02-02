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
	"strings"
)

// NetworkStatusClient contains the methods for the NetworkStatus group.
// Don't use this type directly, use NewNetworkStatusClient() instead.
type NetworkStatusClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkStatusClient creates a new instance of NetworkStatusClient with the specified values.
func NewNetworkStatusClient(con *armcore.Connection, subscriptionID string) *NetworkStatusClient {
	return &NetworkStatusClient{con: con, subscriptionID: subscriptionID}
}

// ListByLocation - Gets the Connectivity Status to the external resources on which the Api Management service depends from inside the Cloud Service. This
// also returns the DNS Servers as visible to the CloudService.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NetworkStatusClient) ListByLocation(ctx context.Context, resourceGroupName string, serviceName string, locationName string, options *NetworkStatusListByLocationOptions) (NetworkStatusListByLocationResponse, error) {
	req, err := client.listByLocationCreateRequest(ctx, resourceGroupName, serviceName, locationName, options)
	if err != nil {
		return NetworkStatusListByLocationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NetworkStatusListByLocationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NetworkStatusListByLocationResponse{}, client.listByLocationHandleError(resp)
	}
	return client.listByLocationHandleResponse(resp)
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *NetworkStatusClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, locationName string, options *NetworkStatusListByLocationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/locations/{locationName}/networkstatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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

// listByLocationHandleResponse handles the ListByLocation response.
func (client *NetworkStatusClient) listByLocationHandleResponse(resp *azcore.Response) (NetworkStatusListByLocationResponse, error) {
	result := NetworkStatusListByLocationResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NetworkStatusContract); err != nil {
		return NetworkStatusListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *NetworkStatusClient) listByLocationHandleError(resp *azcore.Response) error {
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

// ListByService - Gets the Connectivity Status to the external resources on which the Api Management service depends from inside the Cloud Service. This
// also returns the DNS Servers as visible to the CloudService.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NetworkStatusClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *NetworkStatusListByServiceOptions) (NetworkStatusListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return NetworkStatusListByServiceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NetworkStatusListByServiceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NetworkStatusListByServiceResponse{}, client.listByServiceHandleError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NetworkStatusClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *NetworkStatusListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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

// listByServiceHandleResponse handles the ListByService response.
func (client *NetworkStatusClient) listByServiceHandleResponse(resp *azcore.Response) (NetworkStatusListByServiceResponse, error) {
	result := NetworkStatusListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NetworkStatusContractByLocationArray); err != nil {
		return NetworkStatusListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *NetworkStatusClient) listByServiceHandleError(resp *azcore.Response) error {
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

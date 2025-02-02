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

// ReportsClient contains the methods for the Reports group.
// Don't use this type directly, use NewReportsClient() instead.
type ReportsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewReportsClient creates a new instance of ReportsClient with the specified values.
func NewReportsClient(con *armcore.Connection, subscriptionID string) *ReportsClient {
	return &ReportsClient{con: con, subscriptionID: subscriptionID}
}

// ListByAPI - Lists report records by API.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByAPI(resourceGroupName string, serviceName string, filter string, options *ReportsListByAPIOptions) ReportsListByAPIPager {
	return &reportsListByAPIPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByAPICreateRequest(ctx, resourceGroupName, serviceName, filter, options)
		},
		advancer: func(ctx context.Context, resp ReportsListByAPIResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listByAPICreateRequest creates the ListByAPI request.
func (client *ReportsClient) listByAPICreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByAPIOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byApi"
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByAPIHandleResponse handles the ListByAPI response.
func (client *ReportsClient) listByAPIHandleResponse(resp *azcore.Response) (ReportsListByAPIResponse, error) {
	result := ReportsListByAPIResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListByAPIResponse{}, err
	}
	return result, nil
}

// listByAPIHandleError handles the ListByAPI error response.
func (client *ReportsClient) listByAPIHandleError(resp *azcore.Response) error {
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

// ListByGeo - Lists report records by geography.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByGeo(resourceGroupName string, serviceName string, filter string, options *ReportsListByGeoOptions) ReportsListByGeoPager {
	return &reportsListByGeoPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByGeoCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
		},
		advancer: func(ctx context.Context, resp ReportsListByGeoResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listByGeoCreateRequest creates the ListByGeo request.
func (client *ReportsClient) listByGeoCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByGeoOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byGeo"
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
	reqQP.Set("$filter", filter)
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

// listByGeoHandleResponse handles the ListByGeo response.
func (client *ReportsClient) listByGeoHandleResponse(resp *azcore.Response) (ReportsListByGeoResponse, error) {
	result := ReportsListByGeoResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListByGeoResponse{}, err
	}
	return result, nil
}

// listByGeoHandleError handles the ListByGeo error response.
func (client *ReportsClient) listByGeoHandleError(resp *azcore.Response) error {
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

// ListByOperation - Lists report records by API Operations.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByOperation(resourceGroupName string, serviceName string, filter string, options *ReportsListByOperationOptions) ReportsListByOperationPager {
	return &reportsListByOperationPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByOperationCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
		},
		advancer: func(ctx context.Context, resp ReportsListByOperationResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listByOperationCreateRequest creates the ListByOperation request.
func (client *ReportsClient) listByOperationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByOperationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byOperation"
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByOperationHandleResponse handles the ListByOperation response.
func (client *ReportsClient) listByOperationHandleResponse(resp *azcore.Response) (ReportsListByOperationResponse, error) {
	result := ReportsListByOperationResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListByOperationResponse{}, err
	}
	return result, nil
}

// listByOperationHandleError handles the ListByOperation error response.
func (client *ReportsClient) listByOperationHandleError(resp *azcore.Response) error {
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

// ListByProduct - Lists report records by Product.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByProduct(resourceGroupName string, serviceName string, filter string, options *ReportsListByProductOptions) ReportsListByProductPager {
	return &reportsListByProductPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
		},
		advancer: func(ctx context.Context, resp ReportsListByProductResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ReportsClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByProductOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byProduct"
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *ReportsClient) listByProductHandleResponse(resp *azcore.Response) (ReportsListByProductResponse, error) {
	result := ReportsListByProductResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListByProductResponse{}, err
	}
	return result, nil
}

// listByProductHandleError handles the ListByProduct error response.
func (client *ReportsClient) listByProductHandleError(resp *azcore.Response) error {
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

// ListByRequest - Lists report records by Request.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByRequestOptions) (ReportsListByRequestResponse, error) {
	req, err := client.listByRequestCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
	if err != nil {
		return ReportsListByRequestResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ReportsListByRequestResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ReportsListByRequestResponse{}, client.listByRequestHandleError(resp)
	}
	return client.listByRequestHandleResponse(resp)
}

// listByRequestCreateRequest creates the ListByRequest request.
func (client *ReportsClient) listByRequestCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByRequestOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byRequest"
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
	reqQP.Set("$filter", filter)
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

// listByRequestHandleResponse handles the ListByRequest response.
func (client *ReportsClient) listByRequestHandleResponse(resp *azcore.Response) (ReportsListByRequestResponse, error) {
	result := ReportsListByRequestResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RequestReportCollection); err != nil {
		return ReportsListByRequestResponse{}, err
	}
	return result, nil
}

// listByRequestHandleError handles the ListByRequest error response.
func (client *ReportsClient) listByRequestHandleError(resp *azcore.Response) error {
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

// ListBySubscription - Lists report records by subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListBySubscription(resourceGroupName string, serviceName string, filter string, options *ReportsListBySubscriptionOptions) ReportsListBySubscriptionPager {
	return &reportsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
		},
		advancer: func(ctx context.Context, resp ReportsListBySubscriptionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ReportsClient) listBySubscriptionCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/bySubscription"
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ReportsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (ReportsListBySubscriptionResponse, error) {
	result := ReportsListBySubscriptionResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ReportsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
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

// ListByTime - Lists report records by Time.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByTime(resourceGroupName string, serviceName string, filter string, interval string, options *ReportsListByTimeOptions) ReportsListByTimePager {
	return &reportsListByTimePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByTimeCreateRequest(ctx, resourceGroupName, serviceName, filter, interval, options)
		},
		advancer: func(ctx context.Context, resp ReportsListByTimeResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listByTimeCreateRequest creates the ListByTime request.
func (client *ReportsClient) listByTimeCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, interval string, options *ReportsListByTimeOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byTime"
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("interval", interval)
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByTimeHandleResponse handles the ListByTime response.
func (client *ReportsClient) listByTimeHandleResponse(resp *azcore.Response) (ReportsListByTimeResponse, error) {
	result := ReportsListByTimeResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListByTimeResponse{}, err
	}
	return result, nil
}

// listByTimeHandleError handles the ListByTime error response.
func (client *ReportsClient) listByTimeHandleError(resp *azcore.Response) error {
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

// ListByUser - Lists report records by User.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReportsClient) ListByUser(resourceGroupName string, serviceName string, filter string, options *ReportsListByUserOptions) ReportsListByUserPager {
	return &reportsListByUserPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByUserCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
		},
		advancer: func(ctx context.Context, resp ReportsListByUserResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReportCollection.NextLink)
		},
	}
}

// listByUserCreateRequest creates the ListByUser request.
func (client *ReportsClient) listByUserCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsListByUserOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byUser"
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByUserHandleResponse handles the ListByUser response.
func (client *ReportsClient) listByUserHandleResponse(resp *azcore.Response) (ReportsListByUserResponse, error) {
	result := ReportsListByUserResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReportCollection); err != nil {
		return ReportsListByUserResponse{}, err
	}
	return result, nil
}

// listByUserHandleError handles the ListByUser error response.
func (client *ReportsClient) listByUserHandleError(resp *azcore.Response) error {
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

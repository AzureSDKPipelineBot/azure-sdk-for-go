//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FarmBeatsModelsClient contains the methods for the FarmBeatsModels group.
// Don't use this type directly, use NewFarmBeatsModelsClient() instead.
type FarmBeatsModelsClient struct {
	ep string
	pl runtime.Pipeline
	subscriptionID string
}

// NewFarmBeatsModelsClient creates a new instance of FarmBeatsModelsClient with the specified values.
func NewFarmBeatsModelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FarmBeatsModelsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &FarmBeatsModelsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) CreateOrUpdate(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeats, options *FarmBeatsModelsCreateOrUpdateOptions) (FarmBeatsModelsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, farmBeatsResourceName, resourceGroupName, body, options)
	if err != nil {
		return FarmBeatsModelsCreateOrUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return FarmBeatsModelsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FarmBeatsModelsClient) createOrUpdateCreateRequest(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeats, options *FarmBeatsModelsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FarmBeatsModelsClient) createOrUpdateHandleResponse(resp *http.Response) (FarmBeatsModelsCreateOrUpdateResponse, error) {
	result := FarmBeatsModelsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeats); err != nil {
		return FarmBeatsModelsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FarmBeatsModelsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete a FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) Delete(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsDeleteOptions) (FarmBeatsModelsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
	if err != nil {
		return FarmBeatsModelsDeleteResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FarmBeatsModelsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return FarmBeatsModelsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FarmBeatsModelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *FarmBeatsModelsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) Get(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsGetOptions) (FarmBeatsModelsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
	if err != nil {
		return FarmBeatsModelsGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FarmBeatsModelsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FarmBeatsModelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FarmBeatsModelsClient) getHandleResponse(resp *http.Response) (FarmBeatsModelsGetResponse, error) {
	result := FarmBeatsModelsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeats); err != nil {
		return FarmBeatsModelsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *FarmBeatsModelsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists the FarmBeats instances for a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) ListByResourceGroup(resourceGroupName string, options *FarmBeatsModelsListByResourceGroupOptions) (*FarmBeatsModelsListByResourceGroupPager) {
	return &FarmBeatsModelsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp FarmBeatsModelsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FarmBeatsListResponse.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FarmBeatsModelsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FarmBeatsModelsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FarmBeatsModelsClient) listByResourceGroupHandleResponse(resp *http.Response) (FarmBeatsModelsListByResourceGroupResponse, error) {
	result := FarmBeatsModelsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeatsListResponse); err != nil {
		return FarmBeatsModelsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *FarmBeatsModelsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists the FarmBeats instances for a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) ListBySubscription(options *FarmBeatsModelsListBySubscriptionOptions) (*FarmBeatsModelsListBySubscriptionPager) {
	return &FarmBeatsModelsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp FarmBeatsModelsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FarmBeatsListResponse.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *FarmBeatsModelsClient) listBySubscriptionCreateRequest(ctx context.Context, options *FarmBeatsModelsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *FarmBeatsModelsClient) listBySubscriptionHandleResponse(resp *http.Response) (FarmBeatsModelsListBySubscriptionResponse, error) {
	result := FarmBeatsModelsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeatsListResponse); err != nil {
		return FarmBeatsModelsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *FarmBeatsModelsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update a FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) Update(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeatsUpdateRequestModel, options *FarmBeatsModelsUpdateOptions) (FarmBeatsModelsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, farmBeatsResourceName, resourceGroupName, body, options)
	if err != nil {
		return FarmBeatsModelsUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return FarmBeatsModelsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FarmBeatsModelsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FarmBeatsModelsClient) updateCreateRequest(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeatsUpdateRequestModel, options *FarmBeatsModelsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *FarmBeatsModelsClient) updateHandleResponse(resp *http.Response) (FarmBeatsModelsUpdateResponse, error) {
	result := FarmBeatsModelsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FarmBeats); err != nil {
		return FarmBeatsModelsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *FarmBeatsModelsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

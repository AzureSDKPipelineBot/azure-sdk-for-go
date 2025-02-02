// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

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

// DisasterRecoveryConfigsClient contains the methods for the DisasterRecoveryConfigs group.
// Don't use this type directly, use NewDisasterRecoveryConfigsClient() instead.
type DisasterRecoveryConfigsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDisasterRecoveryConfigsClient creates a new instance of DisasterRecoveryConfigsClient with the specified values.
func NewDisasterRecoveryConfigsClient(con *armcore.Connection, subscriptionID string) *DisasterRecoveryConfigsClient {
	return &DisasterRecoveryConfigsClient{con: con, subscriptionID: subscriptionID}
}

// BreakPairing - This operation disables the Disaster Recovery and stops replicating changes from primary to secondary namespaces
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) BreakPairing(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsBreakPairingOptions) (DisasterRecoveryConfigsBreakPairingResponse, error) {
	req, err := client.breakPairingCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsBreakPairingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsBreakPairingResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsBreakPairingResponse{}, client.breakPairingHandleError(resp)
	}
	return DisasterRecoveryConfigsBreakPairingResponse{RawResponse: resp.Response}, nil
}

// breakPairingCreateRequest creates the BreakPairing request.
func (client *DisasterRecoveryConfigsClient) breakPairingCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsBreakPairingOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/breakPairing"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// breakPairingHandleError handles the BreakPairing error response.
func (client *DisasterRecoveryConfigsClient) breakPairingHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CheckNameAvailability - Check the give Namespace name availability.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailabilityParameter, options *DisasterRecoveryConfigsCheckNameAvailabilityOptions) (DisasterRecoveryConfigsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, namespaceName, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailabilityParameter, options *DisasterRecoveryConfigsCheckNameAvailabilityOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsCheckNameAvailabilityResponse, error) {
	result := DisasterRecoveryConfigsCheckNameAvailabilityResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.CheckNameAvailabilityResult); err != nil {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CreateOrUpdate - Creates or updates a new Alias(Disaster Recovery configuration)
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsCreateOrUpdateOptions) (DisasterRecoveryConfigsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, alias, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisasterRecoveryConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DisasterRecoveryConfigsClient) createOrUpdateHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsCreateOrUpdateResponse, error) {
	result := DisasterRecoveryConfigsCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DisasterRecoveryConfigsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes an Alias(Disaster Recovery configuration)
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsDeleteOptions) (DisasterRecoveryConfigsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DisasterRecoveryConfigsDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DisasterRecoveryConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DisasterRecoveryConfigsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// FailOver - Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) FailOver(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsFailOverOptions) (DisasterRecoveryConfigsFailOverResponse, error) {
	req, err := client.failOverCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsFailOverResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsFailOverResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsFailOverResponse{}, client.failOverHandleError(resp)
	}
	return DisasterRecoveryConfigsFailOverResponse{RawResponse: resp.Response}, nil
}

// failOverCreateRequest creates the FailOver request.
func (client *DisasterRecoveryConfigsClient) failOverCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsFailOverOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// failOverHandleError handles the FailOver error response.
func (client *DisasterRecoveryConfigsClient) failOverHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsGetOptions) (DisasterRecoveryConfigsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DisasterRecoveryConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisasterRecoveryConfigsClient) getHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsGetResponse, error) {
	result := DisasterRecoveryConfigsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DisasterRecoveryConfigsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetAuthorizationRule - Gets an AuthorizationRule for a Namespace by rule name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsGetAuthorizationRuleOptions) (DisasterRecoveryConfigsGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, client.getAuthorizationRuleHandleError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsGetAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
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
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsGetAuthorizationRuleResponse, error) {
	result := DisasterRecoveryConfigsGetAuthorizationRuleResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AuthorizationRule); err != nil {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// getAuthorizationRuleHandleError handles the GetAuthorizationRule error response.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all Alias(Disaster Recovery configurations)
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) List(resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsListOptions) DisasterRecoveryConfigsListPager {
	return &disasterRecoveryConfigsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp DisasterRecoveryConfigsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ArmDisasterRecoveryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DisasterRecoveryConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
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
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisasterRecoveryConfigsClient) listHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsListResponse, error) {
	result := DisasterRecoveryConfigsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ArmDisasterRecoveryListResult); err != nil {
		return DisasterRecoveryConfigsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DisasterRecoveryConfigsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListAuthorizationRules - Gets a list of authorization rules for a Namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsListAuthorizationRulesOptions) DisasterRecoveryConfigsListAuthorizationRulesPager {
	return &disasterRecoveryConfigsListAuthorizationRulesPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
		},
		advancer: func(ctx context.Context, resp DisasterRecoveryConfigsListAuthorizationRulesResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AuthorizationRuleListResult.NextLink)
		},
	}
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsListAuthorizationRulesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
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
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsListAuthorizationRulesResponse, error) {
	result := DisasterRecoveryConfigsListAuthorizationRulesResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AuthorizationRuleListResult); err != nil {
		return DisasterRecoveryConfigsListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// listAuthorizationRulesHandleError handles the ListAuthorizationRules error response.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListKeys - Gets the primary and secondary connection strings for the Namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsListKeysOptions) (DisasterRecoveryConfigsListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsListKeysResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsListKeysResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DisasterRecoveryConfigsListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *DisasterRecoveryConfigsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsListKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules/{authorizationRuleName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *DisasterRecoveryConfigsClient) listKeysHandleResponse(resp *azcore.Response) (DisasterRecoveryConfigsListKeysResponse, error) {
	result := DisasterRecoveryConfigsListKeysResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AccessKeys); err != nil {
		return DisasterRecoveryConfigsListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *DisasterRecoveryConfigsClient) listKeysHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

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

// CheckNameAvailabilityClient contains the methods for the CheckNameAvailability group.
// Don't use this type directly, use NewCheckNameAvailabilityClient() instead.
type CheckNameAvailabilityClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient with the specified values.
func NewCheckNameAvailabilityClient(con *armcore.Connection, subscriptionID string) *CheckNameAvailabilityClient {
	return &CheckNameAvailabilityClient{con: con, subscriptionID: subscriptionID}
}

// Execute - Check the availability of name for resource
// If the operation fails it returns the *CloudError error type.
func (client *CheckNameAvailabilityClient) Execute(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *CheckNameAvailabilityExecuteOptions) (CheckNameAvailabilityExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return CheckNameAvailabilityExecuteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return CheckNameAvailabilityExecuteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CheckNameAvailabilityExecuteResponse{}, client.executeHandleError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *CheckNameAvailabilityClient) executeCreateRequest(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *CheckNameAvailabilityExecuteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/checkNameAvailability"
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
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(nameAvailabilityRequest)
}

// executeHandleResponse handles the Execute response.
func (client *CheckNameAvailabilityClient) executeHandleResponse(resp *azcore.Response) (CheckNameAvailabilityExecuteResponse, error) {
	result := CheckNameAvailabilityExecuteResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NameAvailability); err != nil {
		return CheckNameAvailabilityExecuteResponse{}, err
	}
	return result, nil
}

// executeHandleError handles the Execute error response.
func (client *CheckNameAvailabilityClient) executeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

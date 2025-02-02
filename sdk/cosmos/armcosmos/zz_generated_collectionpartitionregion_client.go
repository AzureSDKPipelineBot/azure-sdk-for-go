// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// CollectionPartitionRegionClient contains the methods for the CollectionPartitionRegion group.
// Don't use this type directly, use NewCollectionPartitionRegionClient() instead.
type CollectionPartitionRegionClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewCollectionPartitionRegionClient creates a new instance of CollectionPartitionRegionClient with the specified values.
func NewCollectionPartitionRegionClient(con *armcore.Connection, subscriptionID string) *CollectionPartitionRegionClient {
	return &CollectionPartitionRegionClient{con: con, subscriptionID: subscriptionID}
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given collection and region, split by partition.
// If the operation fails it returns a generic error.
func (client *CollectionPartitionRegionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionRegionListMetricsOptions) (CollectionPartitionRegionListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, region, databaseRid, collectionRid, filter, options)
	if err != nil {
		return CollectionPartitionRegionListMetricsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return CollectionPartitionRegionListMetricsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CollectionPartitionRegionListMetricsResponse{}, client.listMetricsHandleError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionPartitionRegionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionRegionListMetricsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if region == "" {
		return nil, errors.New("parameter region cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{region}", url.PathEscape(region))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	reqQP.Set("$filter", filter)
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionPartitionRegionClient) listMetricsHandleResponse(resp *azcore.Response) (CollectionPartitionRegionListMetricsResponse, error) {
	result := CollectionPartitionRegionListMetricsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PartitionMetricListResult); err != nil {
		return CollectionPartitionRegionListMetricsResponse{}, err
	}
	return result, nil
}

// listMetricsHandleError handles the ListMetrics error response.
func (client *CollectionPartitionRegionClient) listMetricsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

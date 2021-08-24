// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ManagementGroupsClient contains the methods for the ManagementGroups group.
// Don't use this type directly, use NewManagementGroupsClient() instead.
type ManagementGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewManagementGroupsClient creates a new instance of ManagementGroupsClient with the specified values.
func NewManagementGroupsClient(con *armcore.Connection, subscriptionID string) *ManagementGroupsClient {
	return &ManagementGroupsClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets a list of management groups connected to a workspace.
// If the operation fails it returns a generic error.
func (client *ManagementGroupsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagementGroupsListOptions) (ManagementGroupsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return ManagementGroupsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagementGroupsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ManagementGroupsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ManagementGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagementGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/managementGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagementGroupsClient) listHandleResponse(resp *azcore.Response) (ManagementGroupsListResponse, error) {
	result := ManagementGroupsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.WorkspaceListManagementGroupsResult); err != nil {
		return ManagementGroupsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ManagementGroupsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
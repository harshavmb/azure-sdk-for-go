//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AvailableServiceTiersClient contains the methods for the AvailableServiceTiers group.
// Don't use this type directly, use NewAvailableServiceTiersClient() instead.
type AvailableServiceTiersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableServiceTiersClient creates a new instance of AvailableServiceTiersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableServiceTiersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableServiceTiersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableServiceTiersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListByWorkspace - Gets the available service tiers for the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - AvailableServiceTiersClientListByWorkspaceOptions contains the optional parameters for the AvailableServiceTiersClient.ListByWorkspace
// method.
func (client *AvailableServiceTiersClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *AvailableServiceTiersClientListByWorkspaceOptions) (AvailableServiceTiersClientListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return AvailableServiceTiersClientListByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailableServiceTiersClientListByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailableServiceTiersClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *AvailableServiceTiersClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AvailableServiceTiersClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/availableServiceTiers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *AvailableServiceTiersClient) listByWorkspaceHandleResponse(resp *http.Response) (AvailableServiceTiersClientListByWorkspaceResponse, error) {
	result := AvailableServiceTiersClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableServiceTierArray); err != nil {
		return AvailableServiceTiersClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

package azurestackhci

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ArcSettingsClient is the azure Stack HCI management service
type ArcSettingsClient struct {
	BaseClient
}

// NewArcSettingsClient creates an instance of the ArcSettingsClient client.
func NewArcSettingsClient(subscriptionID string) ArcSettingsClient {
	return NewArcSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewArcSettingsClientWithBaseURI creates an instance of the ArcSettingsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewArcSettingsClientWithBaseURI(baseURI string, subscriptionID string) ArcSettingsClient {
	return ArcSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create ArcSetting for HCI cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
// arcSettingName - the name of the proxy resource holding details of HCI ArcSetting information.
// arcSetting - parameters supplied to the Create ArcSetting resource for this HCI cluster.
func (client ArcSettingsClient) Create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSetting) (result ArcSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, arcSettingName, arcSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ArcSettingsClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"arcSettingName":    autorest.Encode("path", arcSettingName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	arcSetting.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}", pathParameters),
		autorest.WithJSON(arcSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) CreateResponder(resp *http.Response) (result ArcSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateIdentity create Aad identity for arc settings.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
// arcSettingName - the name of the proxy resource holding details of HCI ArcSetting information.
func (client ArcSettingsClient) CreateIdentity(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result ArcSettingsCreateIdentityFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.CreateIdentity")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "CreateIdentity", err.Error())
	}

	req, err := client.CreateIdentityPreparer(ctx, resourceGroupName, clusterName, arcSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "CreateIdentity", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateIdentitySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "CreateIdentity", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateIdentityPreparer prepares the CreateIdentity request.
func (client ArcSettingsClient) CreateIdentityPreparer(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"arcSettingName":    autorest.Encode("path", arcSettingName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/createArcIdentity", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateIdentitySender sends the CreateIdentity request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) CreateIdentitySender(req *http.Request) (future ArcSettingsCreateIdentityFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateIdentityResponder handles the response to the CreateIdentity request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) CreateIdentityResponder(resp *http.Response) (result ArcIdentityResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete ArcSetting resource details of HCI Cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
// arcSettingName - the name of the proxy resource holding details of HCI ArcSetting information.
func (client ArcSettingsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result ArcSettingsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, arcSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ArcSettingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"arcSettingName":    autorest.Encode("path", arcSettingName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) DeleteSender(req *http.Request) (future ArcSettingsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GeneratePassword generate password for arc settings.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
// arcSettingName - the name of the proxy resource holding details of HCI ArcSetting information.
func (client ArcSettingsClient) GeneratePassword(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result PasswordCredential, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.GeneratePassword")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "GeneratePassword", err.Error())
	}

	req, err := client.GeneratePasswordPreparer(ctx, resourceGroupName, clusterName, arcSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "GeneratePassword", nil, "Failure preparing request")
		return
	}

	resp, err := client.GeneratePasswordSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "GeneratePassword", resp, "Failure sending request")
		return
	}

	result, err = client.GeneratePasswordResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "GeneratePassword", resp, "Failure responding to request")
		return
	}

	return
}

// GeneratePasswordPreparer prepares the GeneratePassword request.
func (client ArcSettingsClient) GeneratePasswordPreparer(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"arcSettingName":    autorest.Encode("path", arcSettingName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/generatePassword", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GeneratePasswordSender sends the GeneratePassword request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) GeneratePasswordSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GeneratePasswordResponder handles the response to the GeneratePassword request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) GeneratePasswordResponder(resp *http.Response) (result PasswordCredential, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get ArcSetting resource details of HCI Cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
// arcSettingName - the name of the proxy resource holding details of HCI ArcSetting information.
func (client ArcSettingsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result ArcSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, arcSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ArcSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"arcSettingName":    autorest.Encode("path", arcSettingName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) GetResponder(resp *http.Response) (result ArcSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByCluster get ArcSetting resources of HCI Cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
func (client ArcSettingsClient) ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result ArcSettingListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.ListByCluster")
		defer func() {
			sc := -1
			if result.asl.Response.Response != nil {
				sc = result.asl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "ListByCluster", err.Error())
	}

	result.fn = client.listByClusterNextResults
	req, err := client.ListByClusterPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "ListByCluster", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByClusterSender(req)
	if err != nil {
		result.asl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "ListByCluster", resp, "Failure sending request")
		return
	}

	result.asl, err = client.ListByClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "ListByCluster", resp, "Failure responding to request")
		return
	}
	if result.asl.hasNextLink() && result.asl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByClusterPreparer prepares the ListByCluster request.
func (client ArcSettingsClient) ListByClusterPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByClusterSender sends the ListByCluster request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) ListByClusterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByClusterResponder handles the response to the ListByCluster request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) ListByClusterResponder(resp *http.Response) (result ArcSettingList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByClusterNextResults retrieves the next set of results, if any.
func (client ArcSettingsClient) listByClusterNextResults(ctx context.Context, lastResults ArcSettingList) (result ArcSettingList, err error) {
	req, err := lastResults.arcSettingListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "listByClusterNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByClusterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "listByClusterNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "listByClusterNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByClusterComplete enumerates all values, automatically crossing page boundaries as required.
func (client ArcSettingsClient) ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result ArcSettingListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.ListByCluster")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCluster(ctx, resourceGroupName, clusterName)
	return
}

// Update update ArcSettings for HCI cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the cluster.
// arcSettingName - the name of the proxy resource holding details of HCI ArcSetting information.
// arcSetting - arcSettings parameters that needs to be updated
func (client ArcSettingsClient) Update(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSettingsPatch) (result ArcSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArcSettingsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.ArcSettingsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, arcSettingName, arcSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.ArcSettingsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ArcSettingsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSettingsPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"arcSettingName":    autorest.Encode("path", arcSettingName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}", pathParameters),
		autorest.WithJSON(arcSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ArcSettingsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ArcSettingsClient) UpdateResponder(resp *http.Response) (result ArcSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

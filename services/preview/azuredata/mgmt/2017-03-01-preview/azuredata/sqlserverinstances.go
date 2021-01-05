package azuredata

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// SQLServerInstancesClient is the the AzureData management API provides a RESTful set of web APIs to manage Azure Data
// Resources.
type SQLServerInstancesClient struct {
	BaseClient
}

// NewSQLServerInstancesClient creates an instance of the SQLServerInstancesClient client.
func NewSQLServerInstancesClient(subscriptionID string, subscriptionID1 string) SQLServerInstancesClient {
	return NewSQLServerInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewSQLServerInstancesClientWithBaseURI creates an instance of the SQLServerInstancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSQLServerInstancesClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) SQLServerInstancesClient {
	return SQLServerInstancesClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// Create creates or replaces a SQL Server Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLServerInstanceName - the name of SQL Server Instance
// parameters - the SQL Server Instance to be created or updated.
func (client SQLServerInstancesClient) Create(ctx context.Context, resourceGroupName string, SQLServerInstanceName string, parameters SQLServerInstance) (result SQLServerInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SQLServerInstanceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.SQLServerInstanceProperties.Version", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.SQLServerInstanceProperties.Edition", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.SQLServerInstanceProperties.ContainerResourceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.SQLServerInstanceProperties.VCore", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.SQLServerInstanceProperties.Status", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("azuredata.SQLServerInstancesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, SQLServerInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SQLServerInstancesClient) CreatePreparer(ctx context.Context, resourceGroupName string, SQLServerInstanceName string, parameters SQLServerInstance) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlServerInstanceName": autorest.Encode("path", SQLServerInstanceName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerInstances/{sqlServerInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerInstancesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SQLServerInstancesClient) CreateResponder(resp *http.Response) (result SQLServerInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a SQL Server Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLServerInstanceName - the name of SQL Server Instance
func (client SQLServerInstancesClient) Delete(ctx context.Context, resourceGroupName string, SQLServerInstanceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, SQLServerInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLServerInstancesClient) DeletePreparer(ctx context.Context, resourceGroupName string, SQLServerInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlServerInstanceName": autorest.Encode("path", SQLServerInstanceName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerInstances/{sqlServerInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerInstancesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLServerInstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves a SQL Server Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLServerInstanceName - name of SQL Server Instance
func (client SQLServerInstancesClient) Get(ctx context.Context, resourceGroupName string, SQLServerInstanceName string) (result SQLServerInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, SQLServerInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLServerInstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, SQLServerInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlServerInstanceName": autorest.Encode("path", SQLServerInstanceName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerInstances/{sqlServerInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLServerInstancesClient) GetResponder(resp *http.Response) (result SQLServerInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client SQLServerInstancesClient) List(ctx context.Context) (result SQLServerInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.List")
		defer func() {
			sc := -1
			if result.ssilr.Response.Response != nil {
				sc = result.ssilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ssilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.ssilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ssilr.hasNextLink() && result.ssilr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client SQLServerInstancesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlServerInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SQLServerInstancesClient) ListResponder(resp *http.Response) (result SQLServerInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SQLServerInstancesClient) listNextResults(ctx context.Context, lastResults SQLServerInstanceListResult) (result SQLServerInstanceListResult, err error) {
	req, err := lastResults.sQLServerInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLServerInstancesClient) ListComplete(ctx context.Context) (result SQLServerInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets all sqlServerInstances in a resource group.
// Parameters:
// resourceGroupName - the name of the Azure resource group
func (client SQLServerInstancesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result SQLServerInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ssilr.Response.Response != nil {
				sc = result.ssilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ssilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ssilr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ssilr.hasNextLink() && result.ssilr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SQLServerInstancesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerInstancesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SQLServerInstancesClient) ListByResourceGroupResponder(resp *http.Response) (result SQLServerInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SQLServerInstancesClient) listByResourceGroupNextResults(ctx context.Context, lastResults SQLServerInstanceListResult) (result SQLServerInstanceListResult, err error) {
	req, err := lastResults.sQLServerInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLServerInstancesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result SQLServerInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates a SQL Server Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLServerInstanceName - name of sqlServerInstance
// parameters - the SQL Server Instance.
func (client SQLServerInstancesClient) Update(ctx context.Context, resourceGroupName string, SQLServerInstanceName string, parameters SQLServerInstanceUpdate) (result SQLServerInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerInstancesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, SQLServerInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerInstancesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SQLServerInstancesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, SQLServerInstanceName string, parameters SQLServerInstanceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlServerInstanceName": autorest.Encode("path", SQLServerInstanceName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerInstances/{sqlServerInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerInstancesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SQLServerInstancesClient) UpdateResponder(resp *http.Response) (result SQLServerInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
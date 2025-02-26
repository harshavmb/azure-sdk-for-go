package datamigration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ResourceSkusClient is the data Migration Client
type ResourceSkusClient struct {
	BaseClient
}

// NewResourceSkusClient creates an instance of the ResourceSkusClient client.
func NewResourceSkusClient(subscriptionID string) ResourceSkusClient {
	return NewResourceSkusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceSkusClientWithBaseURI creates an instance of the ResourceSkusClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewResourceSkusClientWithBaseURI(baseURI string, subscriptionID string) ResourceSkusClient {
	return ResourceSkusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListSkus the skus action returns the list of SKUs that DMS supports.
func (client ResourceSkusClient) ListSkus(ctx context.Context) (result ResourceSkusResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceSkusClient.ListSkus")
		defer func() {
			sc := -1
			if result.rsr.Response.Response != nil {
				sc = result.rsr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listSkusNextResults
	req, err := client.ListSkusPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ResourceSkusClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.rsr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.ResourceSkusClient", "ListSkus", resp, "Failure sending request")
		return
	}

	result.rsr, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ResourceSkusClient", "ListSkus", resp, "Failure responding to request")
		return
	}
	if result.rsr.hasNextLink() && result.rsr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListSkusPreparer prepares the ListSkus request.
func (client ResourceSkusClient) ListSkusPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSkusSender sends the ListSkus request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceSkusClient) ListSkusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client ResourceSkusClient) ListSkusResponder(resp *http.Response) (result ResourceSkusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listSkusNextResults retrieves the next set of results, if any.
func (client ResourceSkusClient) listSkusNextResults(ctx context.Context, lastResults ResourceSkusResult) (result ResourceSkusResult, err error) {
	req, err := lastResults.resourceSkusResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datamigration.ResourceSkusClient", "listSkusNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datamigration.ResourceSkusClient", "listSkusNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ResourceSkusClient", "listSkusNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListSkusComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceSkusClient) ListSkusComplete(ctx context.Context) (result ResourceSkusResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceSkusClient.ListSkus")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListSkus(ctx)
	return
}

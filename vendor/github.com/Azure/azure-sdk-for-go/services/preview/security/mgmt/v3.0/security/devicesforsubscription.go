package security

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

// DevicesForSubscriptionClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type DevicesForSubscriptionClient struct {
	BaseClient
}

// NewDevicesForSubscriptionClient creates an instance of the DevicesForSubscriptionClient client.
func NewDevicesForSubscriptionClient(subscriptionID string, ascLocation string) DevicesForSubscriptionClient {
	return NewDevicesForSubscriptionClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewDevicesForSubscriptionClientWithBaseURI creates an instance of the DevicesForSubscriptionClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewDevicesForSubscriptionClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) DevicesForSubscriptionClient {
	return DevicesForSubscriptionClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// List get list of the devices by their subscription.
// Parameters:
// limit - limit the number of items returned in a single page
// skipToken - skip token used for pagination
// deviceManagementType - get devices only from specific type, Managed or Unmanaged.
func (client DevicesForSubscriptionClient) List(ctx context.Context, limit *int32, skipToken string, deviceManagementType ManagementState) (result DeviceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DevicesForSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.dl.Response.Response != nil {
				sc = result.dl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.DevicesForSubscriptionClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, limit, skipToken, deviceManagementType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DevicesForSubscriptionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.DevicesForSubscriptionClient", "List", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DevicesForSubscriptionClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dl.hasNextLink() && result.dl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DevicesForSubscriptionClient) ListPreparer(ctx context.Context, limit *int32, skipToken string, deviceManagementType ManagementState) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if limit != nil {
		queryParameters["$limit"] = autorest.Encode("query", *limit)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}
	if len(string(deviceManagementType)) > 0 {
		queryParameters["deviceManagementType"] = autorest.Encode("query", deviceManagementType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/devices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DevicesForSubscriptionClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DevicesForSubscriptionClient) ListResponder(resp *http.Response) (result DeviceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DevicesForSubscriptionClient) listNextResults(ctx context.Context, lastResults DeviceList) (result DeviceList, err error) {
	req, err := lastResults.deviceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.DevicesForSubscriptionClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.DevicesForSubscriptionClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.DevicesForSubscriptionClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DevicesForSubscriptionClient) ListComplete(ctx context.Context, limit *int32, skipToken string, deviceManagementType ManagementState) (result DeviceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DevicesForSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, limit, skipToken, deviceManagementType)
	return
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationGatewayWafDynamicManifestsClient contains the methods for the ApplicationGatewayWafDynamicManifests group.
// Don't use this type directly, use NewApplicationGatewayWafDynamicManifestsClient() instead.
type ApplicationGatewayWafDynamicManifestsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationGatewayWafDynamicManifestsClient creates a new instance of ApplicationGatewayWafDynamicManifestsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationGatewayWafDynamicManifestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationGatewayWafDynamicManifestsClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationGatewayWafDynamicManifestsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationGatewayWafDynamicManifestsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewGetPager - Gets the regional application gateway waf manifest.
//
// Generated from API version 2022-11-01
//   - location - The region where the nrp are located at.
//   - options - ApplicationGatewayWafDynamicManifestsClientGetOptions contains the optional parameters for the ApplicationGatewayWafDynamicManifestsClient.NewGetPager
//     method.
func (client *ApplicationGatewayWafDynamicManifestsClient) NewGetPager(location string, options *ApplicationGatewayWafDynamicManifestsClientGetOptions) *runtime.Pager[ApplicationGatewayWafDynamicManifestsClientGetResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationGatewayWafDynamicManifestsClientGetResponse]{
		More: func(page ApplicationGatewayWafDynamicManifestsClientGetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationGatewayWafDynamicManifestsClientGetResponse) (ApplicationGatewayWafDynamicManifestsClientGetResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationGatewayWafDynamicManifestsClientGetResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationGatewayWafDynamicManifestsClientGetResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationGatewayWafDynamicManifestsClientGetResponse{}, runtime.NewResponseError(resp)
			}
			return client.getHandleResponse(resp)
		},
	})
}

// getCreateRequest creates the Get request.
func (client *ApplicationGatewayWafDynamicManifestsClient) getCreateRequest(ctx context.Context, location string, options *ApplicationGatewayWafDynamicManifestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/applicationGatewayWafDynamicManifests"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGatewayWafDynamicManifestsClient) getHandleResponse(resp *http.Response) (ApplicationGatewayWafDynamicManifestsClientGetResponse, error) {
	result := ApplicationGatewayWafDynamicManifestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGatewayWafDynamicManifestResultList); err != nil {
		return ApplicationGatewayWafDynamicManifestsClientGetResponse{}, err
	}
	return result, nil
}

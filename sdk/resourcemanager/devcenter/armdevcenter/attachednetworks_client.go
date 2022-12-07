//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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
	"strconv"
	"strings"
)

// AttachedNetworksClient contains the methods for the AttachedNetworks group.
// Don't use this type directly, use NewAttachedNetworksClient() instead.
type AttachedNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAttachedNetworksClient creates a new instance of AttachedNetworksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAttachedNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AttachedNetworksClient, error) {
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
	client := &AttachedNetworksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an attached NetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// attachedNetworkConnectionName - The name of the attached NetworkConnection.
// body - Represents an attached NetworkConnection.
// options - AttachedNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the AttachedNetworksClient.BeginCreateOrUpdate
// method.
func (client *AttachedNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, body AttachedNetworkConnection, options *AttachedNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[AttachedNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, devCenterName, attachedNetworkConnectionName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AttachedNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AttachedNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an attached NetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
func (client *AttachedNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, body AttachedNetworkConnection, options *AttachedNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, devCenterName, attachedNetworkConnectionName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AttachedNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, body AttachedNetworkConnection, options *AttachedNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/attachednetworks/{attachedNetworkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if attachedNetworkConnectionName == "" {
		return nil, errors.New("parameter attachedNetworkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedNetworkConnectionName}", url.PathEscape(attachedNetworkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Un-attach a NetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// attachedNetworkConnectionName - The name of the attached NetworkConnection.
// options - AttachedNetworksClientBeginDeleteOptions contains the optional parameters for the AttachedNetworksClient.BeginDelete
// method.
func (client *AttachedNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, options *AttachedNetworksClientBeginDeleteOptions) (*runtime.Poller[AttachedNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, devCenterName, attachedNetworkConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AttachedNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AttachedNetworksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Un-attach a NetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
func (client *AttachedNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, options *AttachedNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, devCenterName, attachedNetworkConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AttachedNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, options *AttachedNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/attachednetworks/{attachedNetworkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if attachedNetworkConnectionName == "" {
		return nil, errors.New("parameter attachedNetworkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedNetworkConnectionName}", url.PathEscape(attachedNetworkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetByDevCenter - Gets an attached NetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// attachedNetworkConnectionName - The name of the attached NetworkConnection.
// options - AttachedNetworksClientGetByDevCenterOptions contains the optional parameters for the AttachedNetworksClient.GetByDevCenter
// method.
func (client *AttachedNetworksClient) GetByDevCenter(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, options *AttachedNetworksClientGetByDevCenterOptions) (AttachedNetworksClientGetByDevCenterResponse, error) {
	req, err := client.getByDevCenterCreateRequest(ctx, resourceGroupName, devCenterName, attachedNetworkConnectionName, options)
	if err != nil {
		return AttachedNetworksClientGetByDevCenterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttachedNetworksClientGetByDevCenterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttachedNetworksClientGetByDevCenterResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByDevCenterHandleResponse(resp)
}

// getByDevCenterCreateRequest creates the GetByDevCenter request.
func (client *AttachedNetworksClient) getByDevCenterCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, attachedNetworkConnectionName string, options *AttachedNetworksClientGetByDevCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/attachednetworks/{attachedNetworkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if attachedNetworkConnectionName == "" {
		return nil, errors.New("parameter attachedNetworkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedNetworkConnectionName}", url.PathEscape(attachedNetworkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByDevCenterHandleResponse handles the GetByDevCenter response.
func (client *AttachedNetworksClient) getByDevCenterHandleResponse(resp *http.Response) (AttachedNetworksClientGetByDevCenterResponse, error) {
	result := AttachedNetworksClientGetByDevCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedNetworkConnection); err != nil {
		return AttachedNetworksClientGetByDevCenterResponse{}, err
	}
	return result, nil
}

// GetByProject - Gets an attached NetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// attachedNetworkConnectionName - The name of the attached NetworkConnection.
// options - AttachedNetworksClientGetByProjectOptions contains the optional parameters for the AttachedNetworksClient.GetByProject
// method.
func (client *AttachedNetworksClient) GetByProject(ctx context.Context, resourceGroupName string, projectName string, attachedNetworkConnectionName string, options *AttachedNetworksClientGetByProjectOptions) (AttachedNetworksClientGetByProjectResponse, error) {
	req, err := client.getByProjectCreateRequest(ctx, resourceGroupName, projectName, attachedNetworkConnectionName, options)
	if err != nil {
		return AttachedNetworksClientGetByProjectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttachedNetworksClientGetByProjectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttachedNetworksClientGetByProjectResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByProjectHandleResponse(resp)
}

// getByProjectCreateRequest creates the GetByProject request.
func (client *AttachedNetworksClient) getByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, attachedNetworkConnectionName string, options *AttachedNetworksClientGetByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/attachednetworks/{attachedNetworkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if attachedNetworkConnectionName == "" {
		return nil, errors.New("parameter attachedNetworkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedNetworkConnectionName}", url.PathEscape(attachedNetworkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByProjectHandleResponse handles the GetByProject response.
func (client *AttachedNetworksClient) getByProjectHandleResponse(resp *http.Response) (AttachedNetworksClientGetByProjectResponse, error) {
	result := AttachedNetworksClientGetByProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedNetworkConnection); err != nil {
		return AttachedNetworksClientGetByProjectResponse{}, err
	}
	return result, nil
}

// NewListByDevCenterPager - Lists the attached NetworkConnections for a DevCenter.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// options - AttachedNetworksClientListByDevCenterOptions contains the optional parameters for the AttachedNetworksClient.ListByDevCenter
// method.
func (client *AttachedNetworksClient) NewListByDevCenterPager(resourceGroupName string, devCenterName string, options *AttachedNetworksClientListByDevCenterOptions) *runtime.Pager[AttachedNetworksClientListByDevCenterResponse] {
	return runtime.NewPager(runtime.PagingHandler[AttachedNetworksClientListByDevCenterResponse]{
		More: func(page AttachedNetworksClientListByDevCenterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttachedNetworksClientListByDevCenterResponse) (AttachedNetworksClientListByDevCenterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDevCenterCreateRequest(ctx, resourceGroupName, devCenterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AttachedNetworksClientListByDevCenterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttachedNetworksClientListByDevCenterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttachedNetworksClientListByDevCenterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDevCenterHandleResponse(resp)
		},
	})
}

// listByDevCenterCreateRequest creates the ListByDevCenter request.
func (client *AttachedNetworksClient) listByDevCenterCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *AttachedNetworksClientListByDevCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/attachednetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDevCenterHandleResponse handles the ListByDevCenter response.
func (client *AttachedNetworksClient) listByDevCenterHandleResponse(resp *http.Response) (AttachedNetworksClientListByDevCenterResponse, error) {
	result := AttachedNetworksClientListByDevCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedNetworkListResult); err != nil {
		return AttachedNetworksClientListByDevCenterResponse{}, err
	}
	return result, nil
}

// NewListByProjectPager - Lists the attached NetworkConnections for a Project.
// Generated from API version 2022-10-12-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// options - AttachedNetworksClientListByProjectOptions contains the optional parameters for the AttachedNetworksClient.ListByProject
// method.
func (client *AttachedNetworksClient) NewListByProjectPager(resourceGroupName string, projectName string, options *AttachedNetworksClientListByProjectOptions) *runtime.Pager[AttachedNetworksClientListByProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[AttachedNetworksClientListByProjectResponse]{
		More: func(page AttachedNetworksClientListByProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttachedNetworksClientListByProjectResponse) (AttachedNetworksClientListByProjectResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AttachedNetworksClientListByProjectResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttachedNetworksClientListByProjectResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttachedNetworksClientListByProjectResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProjectHandleResponse(resp)
		},
	})
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *AttachedNetworksClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *AttachedNetworksClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/attachednetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-12-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *AttachedNetworksClient) listByProjectHandleResponse(resp *http.Response) (AttachedNetworksClientListByProjectResponse, error) {
	result := AttachedNetworksClientListByProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedNetworkListResult); err != nil {
		return AttachedNetworksClientListByProjectResponse{}, err
	}
	return result, nil
}

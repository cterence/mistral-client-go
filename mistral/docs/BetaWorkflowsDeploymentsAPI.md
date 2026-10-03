# \BetaWorkflowsDeploymentsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeploymentV1WorkflowsDeploymentsNameGet**](BetaWorkflowsDeploymentsAPI.md#GetDeploymentV1WorkflowsDeploymentsNameGet) | **Get** /v1/workflows/deployments/{name} | Get Deployment
[**ListDeploymentsV1WorkflowsDeploymentsGet**](BetaWorkflowsDeploymentsAPI.md#ListDeploymentsV1WorkflowsDeploymentsGet) | **Get** /v1/workflows/deployments | List Deployments



## GetDeploymentV1WorkflowsDeploymentsNameGet

> DeploymentDetailResponse GetDeploymentV1WorkflowsDeploymentsNameGet(ctx, name).Execute()

Get Deployment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsDeploymentsAPI.GetDeploymentV1WorkflowsDeploymentsNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsDeploymentsAPI.GetDeploymentV1WorkflowsDeploymentsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentV1WorkflowsDeploymentsNameGet`: DeploymentDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsDeploymentsAPI.GetDeploymentV1WorkflowsDeploymentsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentV1WorkflowsDeploymentsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentDetailResponse**](DeploymentDetailResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeploymentsV1WorkflowsDeploymentsGet

> DeploymentListResponse ListDeploymentsV1WorkflowsDeploymentsGet(ctx).ActiveOnly(activeOnly).WorkflowName(workflowName).Execute()

List Deployments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	activeOnly := true // bool |  (optional) (default to true)
	workflowName := "workflowName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsDeploymentsAPI.ListDeploymentsV1WorkflowsDeploymentsGet(context.Background()).ActiveOnly(activeOnly).WorkflowName(workflowName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsDeploymentsAPI.ListDeploymentsV1WorkflowsDeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeploymentsV1WorkflowsDeploymentsGet`: DeploymentListResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsDeploymentsAPI.ListDeploymentsV1WorkflowsDeploymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentsV1WorkflowsDeploymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeOnly** | **bool** |  | [default to true]
 **workflowName** | **string** |  | 

### Return type

[**DeploymentListResponse**](DeploymentListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


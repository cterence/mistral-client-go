# \WorkflowsDeploymentsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentV1WorkflowsDeploymentsPost**](WorkflowsDeploymentsAPI.md#CreateDeploymentV1WorkflowsDeploymentsPost) | **Post** /v1/workflows/deployments | Create Deployment
[**DeleteDeploymentV1WorkflowsDeploymentsNameDelete**](WorkflowsDeploymentsAPI.md#DeleteDeploymentV1WorkflowsDeploymentsNameDelete) | **Delete** /v1/workflows/deployments/{name} | Delete Deployment
[**GetDeploymentLogs**](WorkflowsDeploymentsAPI.md#GetDeploymentLogs) | **Get** /v1/workflows/deployments/{name}/logs | Get Deployment Logs
[**GetDeploymentV1WorkflowsDeploymentsNameGet**](WorkflowsDeploymentsAPI.md#GetDeploymentV1WorkflowsDeploymentsNameGet) | **Get** /v1/workflows/deployments/{name} | Get Deployment
[**ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet**](WorkflowsDeploymentsAPI.md#ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet) | **Get** /v1/workflows/deployments/{name}/workers | List Deployment Workers
[**ListDeploymentsV1WorkflowsDeploymentsGet**](WorkflowsDeploymentsAPI.md#ListDeploymentsV1WorkflowsDeploymentsGet) | **Get** /v1/workflows/deployments | List Deployments
[**RestartDeploymentV1WorkflowsDeploymentsNameRestartPost**](WorkflowsDeploymentsAPI.md#RestartDeploymentV1WorkflowsDeploymentsNameRestartPost) | **Post** /v1/workflows/deployments/{name}/restart | Restart Deployment
[**StartDeploymentV1WorkflowsDeploymentsNameStartPost**](WorkflowsDeploymentsAPI.md#StartDeploymentV1WorkflowsDeploymentsNameStartPost) | **Post** /v1/workflows/deployments/{name}/start | Start Deployment
[**StopDeploymentV1WorkflowsDeploymentsNameStopPost**](WorkflowsDeploymentsAPI.md#StopDeploymentV1WorkflowsDeploymentsNameStopPost) | **Post** /v1/workflows/deployments/{name}/stop | Stop Deployment
[**StreamDeploymentLogs**](WorkflowsDeploymentsAPI.md#StreamDeploymentLogs) | **Get** /v1/workflows/deployments/{name}/logs/stream | Stream Deployment Logs
[**UnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPost**](WorkflowsDeploymentsAPI.md#UnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPost) | **Post** /v1/workflows/deployments/{deployment_id}/unharden | Unharden Deployment
[**UpdateDeploymentV1WorkflowsDeploymentsNamePatch**](WorkflowsDeploymentsAPI.md#UpdateDeploymentV1WorkflowsDeploymentsNamePatch) | **Patch** /v1/workflows/deployments/{name} | Update Deployment



## CreateDeploymentV1WorkflowsDeploymentsPost

> ManagedDeploymentResponse CreateDeploymentV1WorkflowsDeploymentsPost(ctx).CreateDeploymentRequest(createDeploymentRequest).Execute()

Create Deployment



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
	createDeploymentRequest := *openapiclient.NewCreateDeploymentRequest("Name_example", *openapiclient.NewDeploymentWorkerSpecInput("GithubUrl_example")) // CreateDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.CreateDeploymentV1WorkflowsDeploymentsPost(context.Background()).CreateDeploymentRequest(createDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.CreateDeploymentV1WorkflowsDeploymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentV1WorkflowsDeploymentsPost`: ManagedDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.CreateDeploymentV1WorkflowsDeploymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentV1WorkflowsDeploymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeploymentRequest** | [**CreateDeploymentRequest**](CreateDeploymentRequest.md) |  | 

### Return type

[**ManagedDeploymentResponse**](ManagedDeploymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentV1WorkflowsDeploymentsNameDelete

> ManagedDeploymentResponse DeleteDeploymentV1WorkflowsDeploymentsNameDelete(ctx, name).Execute()

Delete Deployment



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
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.DeleteDeploymentV1WorkflowsDeploymentsNameDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.DeleteDeploymentV1WorkflowsDeploymentsNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeploymentV1WorkflowsDeploymentsNameDelete`: ManagedDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.DeleteDeploymentV1WorkflowsDeploymentsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentV1WorkflowsDeploymentsNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedDeploymentResponse**](ManagedDeploymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentLogs

> DeploymentLogSearchResponse GetDeploymentLogs(ctx, name).LogType(logType).WorkerName(workerName).WorkflowName(workflowName).After(after).Before(before).Order(order).Cursor(cursor).Limit(limit).Execute()

Get Deployment Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	name := "name_example" // string | 
	logType := openapiclient.DeploymentLogType("build") // DeploymentLogType | Which stream to read. 'runtime' is the workers' own logs; 'build' reads the image build output and requires a managed deployment. (optional) (default to "runtime")
	workerName := "workerName_example" // string | Filter logs by worker name (optional)
	workflowName := "workflowName_example" // string | Filter logs by workflow name (optional)
	after := time.Now() // time.Time | Only return logs at or after this timestamp (optional)
	before := time.Now() // time.Time | Only return logs before this timestamp (optional)
	order := "order_example" // string | First-page sort order: 'asc' (oldest first) or 'desc'. Ignored when `cursor` is set. (optional) (default to "asc")
	cursor := "cursor_example" // string | Pagination cursor from a previous response's `next_cursor`; carries the window and order (optional)
	limit := int32(56) // int32 | Maximum number of logs to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.GetDeploymentLogs(context.Background(), name).LogType(logType).WorkerName(workerName).WorkflowName(workflowName).After(after).Before(before).Order(order).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.GetDeploymentLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentLogs`: DeploymentLogSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.GetDeploymentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logType** | [**DeploymentLogType**](DeploymentLogType.md) | Which stream to read. &#39;runtime&#39; is the workers&#39; own logs; &#39;build&#39; reads the image build output and requires a managed deployment. | [default to &quot;runtime&quot;]
 **workerName** | **string** | Filter logs by worker name | 
 **workflowName** | **string** | Filter logs by workflow name | 
 **after** | **time.Time** | Only return logs at or after this timestamp | 
 **before** | **time.Time** | Only return logs before this timestamp | 
 **order** | **string** | First-page sort order: &#39;asc&#39; (oldest first) or &#39;desc&#39;. Ignored when &#x60;cursor&#x60; is set. | [default to &quot;asc&quot;]
 **cursor** | **string** | Pagination cursor from a previous response&#39;s &#x60;next_cursor&#x60;; carries the window and order | 
 **limit** | **int32** | Maximum number of logs to return | [default to 50]

### Return type

[**DeploymentLogSearchResponse**](DeploymentLogSearchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentV1WorkflowsDeploymentsNameGet

> DeploymentDetailResponse GetDeploymentV1WorkflowsDeploymentsNameGet(ctx, name).WorkflowName(workflowName).Execute()

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
	workflowName := "workflowName_example" // string | Scope serving status to this workflow (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.GetDeploymentV1WorkflowsDeploymentsNameGet(context.Background(), name).WorkflowName(workflowName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.GetDeploymentV1WorkflowsDeploymentsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentV1WorkflowsDeploymentsNameGet`: DeploymentDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.GetDeploymentV1WorkflowsDeploymentsNameGet`: %v\n", resp)
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

 **workflowName** | **string** | Scope serving status to this workflow | 

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


## ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet

> DeploymentWorkerListResponse ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet(ctx, name).WorkerStatus(workerStatus).Limit(limit).Cursor(cursor).Execute()

List Deployment Workers



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
	workerStatus := "workerStatus_example" // string | Filter by worker activity. active=only active, inactive=only inactive, None=no filter (optional)
	limit := int32(56) // int32 | Maximum number of workers to return (optional) (default to 50)
	cursor := "cursor_example" // string | Cursor from a previous response's `next_cursor`. Resend `worker_status` unchanged alongside it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet(context.Background(), name).WorkerStatus(workerStatus).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet`: DeploymentWorkerListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.ListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentWorkersV1WorkflowsDeploymentsNameWorkersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerStatus** | **string** | Filter by worker activity. active&#x3D;only active, inactive&#x3D;only inactive, None&#x3D;no filter | 
 **limit** | **int32** | Maximum number of workers to return | [default to 50]
 **cursor** | **string** | Cursor from a previous response&#39;s &#x60;next_cursor&#x60;. Resend &#x60;worker_status&#x60; unchanged alongside it. | 

### Return type

[**DeploymentWorkerListResponse**](DeploymentWorkerListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeploymentsV1WorkflowsDeploymentsGet

> DeploymentListResponse ListDeploymentsV1WorkflowsDeploymentsGet(ctx).ActiveOnly(activeOnly).IsHardened(isHardened).WorkflowName(workflowName).CreatedBy(createdBy).Owner(owner).LocationTypes(locationTypes).Search(search).OrderBy(orderBy).Order(order).Limit(limit).Cursor(cursor).WorkspaceId(workspaceId).Execute()

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
	isHardened := true // bool | Filter deployments by hardened status (optional)
	workflowName := "workflowName_example" // string |  (optional)
	createdBy := "createdBy_example" // string | Filter deployments by creator's user id (optional)
	owner := "owner_example" // string | Filter deployments by owner's user id (optional)
	locationTypes := []openapiclient.LocationType{openapiclient.LocationType("local")} // []LocationType | Filter deployments with at least one worker on any of these location types (OR) (optional)
	search := "search_example" // string | Filter deployments by name or ID prefix (optional)
	orderBy := "orderBy_example" // string | Field to sort by. When omitted, active and managed deployments are grouped first, then sorted by created_at. When set, results are sorted purely by the specified field with no grouping. (optional)
	order := "order_example" // string | Sort direction. Applied to order_by when set, or within each activity group when omitted. (optional) (default to "desc")
	limit := int32(56) // int32 | Maximum number of deployments to return (optional)
	cursor := "cursor_example" // string | Cursor from a previous response for pagination (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID to scope the request to. Defaults to the caller's context. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.ListDeploymentsV1WorkflowsDeploymentsGet(context.Background()).ActiveOnly(activeOnly).IsHardened(isHardened).WorkflowName(workflowName).CreatedBy(createdBy).Owner(owner).LocationTypes(locationTypes).Search(search).OrderBy(orderBy).Order(order).Limit(limit).Cursor(cursor).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.ListDeploymentsV1WorkflowsDeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeploymentsV1WorkflowsDeploymentsGet`: DeploymentListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.ListDeploymentsV1WorkflowsDeploymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentsV1WorkflowsDeploymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeOnly** | **bool** |  | [default to true]
 **isHardened** | **bool** | Filter deployments by hardened status | 
 **workflowName** | **string** |  | 
 **createdBy** | **string** | Filter deployments by creator&#39;s user id | 
 **owner** | **string** | Filter deployments by owner&#39;s user id | 
 **locationTypes** | [**[]LocationType**](LocationType.md) | Filter deployments with at least one worker on any of these location types (OR) | 
 **search** | **string** | Filter deployments by name or ID prefix | 
 **orderBy** | **string** | Field to sort by. When omitted, active and managed deployments are grouped first, then sorted by created_at. When set, results are sorted purely by the specified field with no grouping. | 
 **order** | **string** | Sort direction. Applied to order_by when set, or within each activity group when omitted. | [default to &quot;desc&quot;]
 **limit** | **int32** | Maximum number of deployments to return | 
 **cursor** | **string** | Cursor from a previous response for pagination | 
 **workspaceId** | **string** | Workspace ID to scope the request to. Defaults to the caller&#39;s context. | 

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


## RestartDeploymentV1WorkflowsDeploymentsNameRestartPost

> ManagedDeploymentResponse RestartDeploymentV1WorkflowsDeploymentsNameRestartPost(ctx, name).Execute()

Restart Deployment



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
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.RestartDeploymentV1WorkflowsDeploymentsNameRestartPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.RestartDeploymentV1WorkflowsDeploymentsNameRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartDeploymentV1WorkflowsDeploymentsNameRestartPost`: ManagedDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.RestartDeploymentV1WorkflowsDeploymentsNameRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartDeploymentV1WorkflowsDeploymentsNameRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedDeploymentResponse**](ManagedDeploymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartDeploymentV1WorkflowsDeploymentsNameStartPost

> ManagedDeploymentResponse StartDeploymentV1WorkflowsDeploymentsNameStartPost(ctx, name).Execute()

Start Deployment



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
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.StartDeploymentV1WorkflowsDeploymentsNameStartPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.StartDeploymentV1WorkflowsDeploymentsNameStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartDeploymentV1WorkflowsDeploymentsNameStartPost`: ManagedDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.StartDeploymentV1WorkflowsDeploymentsNameStartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartDeploymentV1WorkflowsDeploymentsNameStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedDeploymentResponse**](ManagedDeploymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopDeploymentV1WorkflowsDeploymentsNameStopPost

> ManagedDeploymentResponse StopDeploymentV1WorkflowsDeploymentsNameStopPost(ctx, name).Execute()

Stop Deployment



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
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.StopDeploymentV1WorkflowsDeploymentsNameStopPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.StopDeploymentV1WorkflowsDeploymentsNameStopPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopDeploymentV1WorkflowsDeploymentsNameStopPost`: ManagedDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.StopDeploymentV1WorkflowsDeploymentsNameStopPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopDeploymentV1WorkflowsDeploymentsNameStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedDeploymentResponse**](ManagedDeploymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamDeploymentLogs

> StreamDeploymentLogs200Response StreamDeploymentLogs(ctx, name).WorkerName(workerName).WorkflowName(workflowName).After(after).LastEventId(lastEventId).LastEventID(lastEventID).Execute()

Stream Deployment Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	name := "name_example" // string | 
	workerName := "workerName_example" // string | Filter logs by worker name (optional)
	workflowName := "workflowName_example" // string | Filter logs by workflow name (optional)
	after := time.Now() // time.Time | Start a fresh stream at this timestamp (ignored when resuming via last_event_id) (optional)
	lastEventId := "lastEventId_example" // string | Resume from this cursor (a prior response's SSE id) (optional)
	lastEventID := "lastEventID_example" // string | Resume from this cursor (a prior response's SSE id). Takes precedence over the query parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.StreamDeploymentLogs(context.Background(), name).WorkerName(workerName).WorkflowName(workflowName).After(after).LastEventId(lastEventId).LastEventID(lastEventID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.StreamDeploymentLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamDeploymentLogs`: StreamDeploymentLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.StreamDeploymentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDeploymentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerName** | **string** | Filter logs by worker name | 
 **workflowName** | **string** | Filter logs by workflow name | 
 **after** | **time.Time** | Start a fresh stream at this timestamp (ignored when resuming via last_event_id) | 
 **lastEventId** | **string** | Resume from this cursor (a prior response&#39;s SSE id) | 
 **lastEventID** | **string** | Resume from this cursor (a prior response&#39;s SSE id). Takes precedence over the query parameter. | 

### Return type

[**StreamDeploymentLogs200Response**](StreamDeploymentLogs200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPost

> UnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPost(ctx, deploymentId).WorkspaceId(workspaceId).Execute()

Unharden Deployment



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID to scope the request to. Defaults to the caller's context. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsDeploymentsAPI.UnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPost(context.Background(), deploymentId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.UnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnhardenDeploymentV1WorkflowsDeploymentsDeploymentIdUnhardenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** | Workspace ID to scope the request to. Defaults to the caller&#39;s context. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentV1WorkflowsDeploymentsNamePatch

> ManagedDeploymentResponse UpdateDeploymentV1WorkflowsDeploymentsNamePatch(ctx, name).UpdateDeploymentRequest(updateDeploymentRequest).Execute()

Update Deployment



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
	updateDeploymentRequest := *openapiclient.NewUpdateDeploymentRequest() // UpdateDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsDeploymentsAPI.UpdateDeploymentV1WorkflowsDeploymentsNamePatch(context.Background(), name).UpdateDeploymentRequest(updateDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsDeploymentsAPI.UpdateDeploymentV1WorkflowsDeploymentsNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentV1WorkflowsDeploymentsNamePatch`: ManagedDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsDeploymentsAPI.UpdateDeploymentV1WorkflowsDeploymentsNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentV1WorkflowsDeploymentsNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeploymentRequest** | [**UpdateDeploymentRequest**](UpdateDeploymentRequest.md) |  | 

### Return type

[**ManagedDeploymentResponse**](ManagedDeploymentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


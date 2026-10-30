# \WorkflowsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut**](WorkflowsAPI.md#ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut) | **Put** /v1/workflows/{workflow_identifier}/archive | Archive Workflow
[**BulkArchiveWorkflowsV1WorkflowsArchivePut**](WorkflowsAPI.md#BulkArchiveWorkflowsV1WorkflowsArchivePut) | **Put** /v1/workflows/archive | Bulk Archive Workflows
[**BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut**](WorkflowsAPI.md#BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut) | **Put** /v1/workflows/unarchive | Bulk Unarchive Workflows
[**ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost**](WorkflowsAPI.md#ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) | **Post** /v1/workflows/registrations/{workflow_registration_id}/execute | Execute Workflow Registration
[**ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost**](WorkflowsAPI.md#ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) | **Post** /v1/workflows/{workflow_identifier}/execute | Execute Workflow
[**GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet**](WorkflowsAPI.md#GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet) | **Get** /v1/workflows/registrations/{workflow_registration_id} | Get Workflow Registration
[**GetWorkflowRegistrationsV1WorkflowsRegistrationsGet**](WorkflowsAPI.md#GetWorkflowRegistrationsV1WorkflowsRegistrationsGet) | **Get** /v1/workflows/registrations | Get Workflow Registrations
[**GetWorkflowV1WorkflowsWorkflowIdentifierGet**](WorkflowsAPI.md#GetWorkflowV1WorkflowsWorkflowIdentifierGet) | **Get** /v1/workflows/{workflow_identifier} | Get Workflow
[**GetWorkflowsV1WorkflowsGet**](WorkflowsAPI.md#GetWorkflowsV1WorkflowsGet) | **Get** /v1/workflows | Get Workflows
[**UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut**](WorkflowsAPI.md#UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut) | **Put** /v1/workflows/{workflow_identifier}/unarchive | Unarchive Workflow
[**UpdateWorkflowV1WorkflowsWorkflowIdentifierPut**](WorkflowsAPI.md#UpdateWorkflowV1WorkflowsWorkflowIdentifierPut) | **Put** /v1/workflows/{workflow_identifier} | Update Workflow



## ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut

> WorkflowArchiveResponse ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut(ctx, workflowIdentifier).Execute()

Archive Workflow



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
	workflowIdentifier := "workflowIdentifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut(context.Background(), workflowIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut`: WorkflowArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowArchiveResponse**](WorkflowArchiveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkArchiveWorkflowsV1WorkflowsArchivePut

> WorkflowBulkArchiveResponse BulkArchiveWorkflowsV1WorkflowsArchivePut(ctx).WorkflowBulkArchiveRequest(workflowBulkArchiveRequest).Execute()

Bulk Archive Workflows



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
	workflowBulkArchiveRequest := *openapiclient.NewWorkflowBulkArchiveRequest([]string{"WorkflowIds_example"}) // WorkflowBulkArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.BulkArchiveWorkflowsV1WorkflowsArchivePut(context.Background()).WorkflowBulkArchiveRequest(workflowBulkArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.BulkArchiveWorkflowsV1WorkflowsArchivePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkArchiveWorkflowsV1WorkflowsArchivePut`: WorkflowBulkArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.BulkArchiveWorkflowsV1WorkflowsArchivePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkArchiveWorkflowsV1WorkflowsArchivePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowBulkArchiveRequest** | [**WorkflowBulkArchiveRequest**](WorkflowBulkArchiveRequest.md) |  | 

### Return type

[**WorkflowBulkArchiveResponse**](WorkflowBulkArchiveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut

> WorkflowBulkUnarchiveResponse BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut(ctx).WorkflowBulkUnarchiveRequest(workflowBulkUnarchiveRequest).Execute()

Bulk Unarchive Workflows



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
	workflowBulkUnarchiveRequest := *openapiclient.NewWorkflowBulkUnarchiveRequest([]string{"WorkflowIds_example"}) // WorkflowBulkUnarchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut(context.Background()).WorkflowBulkUnarchiveRequest(workflowBulkUnarchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut`: WorkflowBulkUnarchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.BulkUnarchiveWorkflowsV1WorkflowsUnarchivePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUnarchiveWorkflowsV1WorkflowsUnarchivePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowBulkUnarchiveRequest** | [**WorkflowBulkUnarchiveRequest**](WorkflowBulkUnarchiveRequest.md) |  | 

### Return type

[**WorkflowBulkUnarchiveResponse**](WorkflowBulkUnarchiveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost

> ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost(ctx, workflowRegistrationId).WorkflowExecutionRequest(workflowExecutionRequest).Execute()

Execute Workflow Registration



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
	workflowRegistrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workflowExecutionRequest := *openapiclient.NewWorkflowExecutionRequest() // WorkflowExecutionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost(context.Background(), workflowRegistrationId).WorkflowExecutionRequest(workflowExecutionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost`: ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRegistrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowExecutionRequest** | [**WorkflowExecutionRequest**](WorkflowExecutionRequest.md) |  | 

### Return type

[**ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost**](ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost

> ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost(ctx, workflowIdentifier).WorkflowExecutionRequest(workflowExecutionRequest).Execute()

Execute Workflow



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
	workflowIdentifier := "workflowIdentifier_example" // string | 
	workflowExecutionRequest := *openapiclient.NewWorkflowExecutionRequest() // WorkflowExecutionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost(context.Background(), workflowIdentifier).WorkflowExecutionRequest(workflowExecutionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost`: ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowExecutionRequest** | [**WorkflowExecutionRequest**](WorkflowExecutionRequest.md) |  | 

### Return type

[**ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost**](ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet

> WorkflowRegistrationGetResponse GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet(ctx, workflowRegistrationId).WithWorkflow(withWorkflow).IncludeShared(includeShared).Execute()

Get Workflow Registration



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
	workflowRegistrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	withWorkflow := true // bool | Whether to include the workflow definition (optional) (default to false)
	includeShared := true // bool | Whether to include shared workflow versions (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet(context.Background(), workflowRegistrationId).WithWorkflow(withWorkflow).IncludeShared(includeShared).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet`: WorkflowRegistrationGetResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRegistrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withWorkflow** | **bool** | Whether to include the workflow definition | [default to false]
 **includeShared** | **bool** | Whether to include shared workflow versions | [default to true]

### Return type

[**WorkflowRegistrationGetResponse**](WorkflowRegistrationGetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowRegistrationsV1WorkflowsRegistrationsGet

> WorkflowRegistrationListResponse GetWorkflowRegistrationsV1WorkflowsRegistrationsGet(ctx).WorkflowId(workflowId).TaskQueue(taskQueue).ActiveOnly(activeOnly).IncludeShared(includeShared).WorkflowName(workflowName).WorkflowSearch(workflowSearch).Archived(archived).WithWorkflow(withWorkflow).AvailableInChatAssistant(availableInChatAssistant).Limit(limit).Cursor(cursor).Execute()

Get Workflow Registrations



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The workflow ID to filter by (optional)
	taskQueue := "taskQueue_example" // string | The task queue to filter by (optional)
	activeOnly := true // bool | Whether to only return active workflows versions (optional) (default to false)
	includeShared := true // bool | Whether to include shared workflow versions (optional) (default to true)
	workflowName := "workflowName_example" // string | Exact workflow name to filter by (optional)
	workflowSearch := "workflowSearch_example" // string | The workflow name to filter by (optional)
	archived := true // bool | Filter by archived state. False=exclude archived, True=only archived, None=include all (optional)
	withWorkflow := true // bool | Whether to include the workflow definition (optional) (default to false)
	availableInChatAssistant := true // bool | Whether to only return workflows available in chat assistant (optional)
	limit := int32(56) // int32 | The maximum number of workflows versions to return (optional) (default to 50)
	cursor := "cursor_example" // string | The cursor for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowRegistrationsV1WorkflowsRegistrationsGet(context.Background()).WorkflowId(workflowId).TaskQueue(taskQueue).ActiveOnly(activeOnly).IncludeShared(includeShared).WorkflowName(workflowName).WorkflowSearch(workflowSearch).Archived(archived).WithWorkflow(withWorkflow).AvailableInChatAssistant(availableInChatAssistant).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowRegistrationsV1WorkflowsRegistrationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowRegistrationsV1WorkflowsRegistrationsGet`: WorkflowRegistrationListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowRegistrationsV1WorkflowsRegistrationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRegistrationsV1WorkflowsRegistrationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** | The workflow ID to filter by | 
 **taskQueue** | **string** | The task queue to filter by | 
 **activeOnly** | **bool** | Whether to only return active workflows versions | [default to false]
 **includeShared** | **bool** | Whether to include shared workflow versions | [default to true]
 **workflowName** | **string** | Exact workflow name to filter by | 
 **workflowSearch** | **string** | The workflow name to filter by | 
 **archived** | **bool** | Filter by archived state. False&#x3D;exclude archived, True&#x3D;only archived, None&#x3D;include all | 
 **withWorkflow** | **bool** | Whether to include the workflow definition | [default to false]
 **availableInChatAssistant** | **bool** | Whether to only return workflows available in chat assistant | 
 **limit** | **int32** | The maximum number of workflows versions to return | [default to 50]
 **cursor** | **string** | The cursor for pagination | 

### Return type

[**WorkflowRegistrationListResponse**](WorkflowRegistrationListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowV1WorkflowsWorkflowIdentifierGet

> WorkflowGetResponse GetWorkflowV1WorkflowsWorkflowIdentifierGet(ctx, workflowIdentifier).Execute()

Get Workflow



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
	workflowIdentifier := "workflowIdentifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowV1WorkflowsWorkflowIdentifierGet(context.Background(), workflowIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowV1WorkflowsWorkflowIdentifierGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowV1WorkflowsWorkflowIdentifierGet`: WorkflowGetResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowV1WorkflowsWorkflowIdentifierGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowV1WorkflowsWorkflowIdentifierGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowGetResponse**](WorkflowGetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowsV1WorkflowsGet

> WorkflowListResponse GetWorkflowsV1WorkflowsGet(ctx).Status(status).IncludeShared(includeShared).AvailableInChatAssistant(availableInChatAssistant).DeploymentName(deploymentName).DeploymentStatus(deploymentStatus).Archived(archived).Tags(tags).SortBy(sortBy).Order(order).Cursor(cursor).Limit(limit).ActiveOnly(activeOnly).Search(search).Execute()

Get Workflows



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
	status := *openapiclient.NewStatus() // Status | Filter by workflow status (optional)
	includeShared := true // bool | Whether to include shared workflows (optional) (default to true)
	availableInChatAssistant := true // bool | Whether to only return workflows available in chat assistant (optional)
	deploymentName := []string{"Inner_example"} // []string | Filter by deployment name(s) (optional)
	deploymentStatus := "deploymentStatus_example" // string | Filter by deployment activity. active=only active, inactive=only inactive, None=no filter (optional)
	archived := true // bool | Filter by archived state. False=exclude archived, True=only archived, None=include all (optional)
	tags := []string{"Inner_example"} // []string | Filter to workflows tagged with all listed tags (AND). (optional)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	order := "order_example" // string | Sort direction (optional) (default to "asc")
	cursor := "cursor_example" // string | The cursor for pagination (optional)
	limit := int32(56) // int32 | The maximum number of workflows to return (optional) (default to 50)
	activeOnly := true // bool | Deprecated: use deployment_status instead (optional) (default to false)
	search := "search_example" // string | Fuzzy search query for workflow name, display name, description, or ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowsV1WorkflowsGet(context.Background()).Status(status).IncludeShared(includeShared).AvailableInChatAssistant(availableInChatAssistant).DeploymentName(deploymentName).DeploymentStatus(deploymentStatus).Archived(archived).Tags(tags).SortBy(sortBy).Order(order).Cursor(cursor).Limit(limit).ActiveOnly(activeOnly).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowsV1WorkflowsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowsV1WorkflowsGet`: WorkflowListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowsV1WorkflowsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsV1WorkflowsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**Status**](Status.md) | Filter by workflow status | 
 **includeShared** | **bool** | Whether to include shared workflows | [default to true]
 **availableInChatAssistant** | **bool** | Whether to only return workflows available in chat assistant | 
 **deploymentName** | **[]string** | Filter by deployment name(s) | 
 **deploymentStatus** | **string** | Filter by deployment activity. active&#x3D;only active, inactive&#x3D;only inactive, None&#x3D;no filter | 
 **archived** | **bool** | Filter by archived state. False&#x3D;exclude archived, True&#x3D;only archived, None&#x3D;include all | 
 **tags** | **[]string** | Filter to workflows tagged with all listed tags (AND). | 
 **sortBy** | **string** | Field to sort by | 
 **order** | **string** | Sort direction | [default to &quot;asc&quot;]
 **cursor** | **string** | The cursor for pagination | 
 **limit** | **int32** | The maximum number of workflows to return | [default to 50]
 **activeOnly** | **bool** | Deprecated: use deployment_status instead | [default to false]
 **search** | **string** | Fuzzy search query for workflow name, display name, description, or ID | 

### Return type

[**WorkflowListResponse**](WorkflowListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut

> WorkflowUnarchiveResponse UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut(ctx, workflowIdentifier).Execute()

Unarchive Workflow



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
	workflowIdentifier := "workflowIdentifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut(context.Background(), workflowIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut`: WorkflowUnarchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowUnarchiveResponse**](WorkflowUnarchiveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowV1WorkflowsWorkflowIdentifierPut

> WorkflowUpdateResponse UpdateWorkflowV1WorkflowsWorkflowIdentifierPut(ctx, workflowIdentifier).WorkflowUpdateRequest(workflowUpdateRequest).Execute()

Update Workflow



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
	workflowIdentifier := "workflowIdentifier_example" // string | 
	workflowUpdateRequest := *openapiclient.NewWorkflowUpdateRequest() // WorkflowUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.UpdateWorkflowV1WorkflowsWorkflowIdentifierPut(context.Background(), workflowIdentifier).WorkflowUpdateRequest(workflowUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.UpdateWorkflowV1WorkflowsWorkflowIdentifierPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowV1WorkflowsWorkflowIdentifierPut`: WorkflowUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.UpdateWorkflowV1WorkflowsWorkflowIdentifierPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowV1WorkflowsWorkflowIdentifierPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowUpdateRequest** | [**WorkflowUpdateRequest**](WorkflowUpdateRequest.md) |  | 

### Return type

[**WorkflowUpdateResponse**](WorkflowUpdateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


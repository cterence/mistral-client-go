# \BetaWorkflowsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut**](BetaWorkflowsAPI.md#ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut) | **Put** /v1/workflows/{workflow_identifier}/archive | Archive Workflow
[**ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost**](BetaWorkflowsAPI.md#ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) | **Post** /v1/workflows/registrations/{workflow_registration_id}/execute | Execute Workflow Registration
[**ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost**](BetaWorkflowsAPI.md#ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) | **Post** /v1/workflows/{workflow_identifier}/execute | Execute Workflow
[**GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet**](BetaWorkflowsAPI.md#GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet) | **Get** /v1/workflows/registrations/{workflow_registration_id} | Get Workflow Registration
[**GetWorkflowRegistrationsV1WorkflowsRegistrationsGet**](BetaWorkflowsAPI.md#GetWorkflowRegistrationsV1WorkflowsRegistrationsGet) | **Get** /v1/workflows/registrations | Get Workflow Registrations
[**GetWorkflowV1WorkflowsWorkflowIdentifierGet**](BetaWorkflowsAPI.md#GetWorkflowV1WorkflowsWorkflowIdentifierGet) | **Get** /v1/workflows/{workflow_identifier} | Get Workflow
[**UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut**](BetaWorkflowsAPI.md#UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut) | **Put** /v1/workflows/{workflow_identifier}/unarchive | Unarchive Workflow
[**UpdateWorkflowV1WorkflowsWorkflowIdentifierPut**](BetaWorkflowsAPI.md#UpdateWorkflowV1WorkflowsWorkflowIdentifierPut) | **Put** /v1/workflows/{workflow_identifier} | Update Workflow



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
	resp, r, err := apiClient.BetaWorkflowsAPI.ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut(context.Background(), workflowIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut`: WorkflowArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.ArchiveWorkflowV1WorkflowsWorkflowIdentifierArchivePut`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsAPI.ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost(context.Background(), workflowRegistrationId).WorkflowExecutionRequest(workflowExecutionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost`: ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.ExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsAPI.ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost(context.Background(), workflowIdentifier).WorkflowExecutionRequest(workflowExecutionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost`: ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.ExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsAPI.GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet(context.Background(), workflowRegistrationId).WithWorkflow(withWorkflow).IncludeShared(includeShared).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet`: WorkflowRegistrationGetResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.GetWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdGet`: %v\n", resp)
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

> WorkflowRegistrationListResponse GetWorkflowRegistrationsV1WorkflowsRegistrationsGet(ctx).WorkflowId(workflowId).TaskQueue(taskQueue).ActiveOnly(activeOnly).IncludeShared(includeShared).WorkflowSearch(workflowSearch).Archived(archived).WithWorkflow(withWorkflow).AvailableInChatAssistant(availableInChatAssistant).Limit(limit).Cursor(cursor).Execute()

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
	workflowSearch := "workflowSearch_example" // string | The workflow name to filter by (optional)
	archived := true // bool | Filter by archived state. False=exclude archived, True=only archived, None=include all (optional)
	withWorkflow := true // bool | Whether to include the workflow definition (optional) (default to false)
	availableInChatAssistant := true // bool | Whether to only return workflows compatible with chat assistant (optional)
	limit := int32(56) // int32 | The maximum number of workflows versions to return (optional) (default to 50)
	cursor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The cursor for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsAPI.GetWorkflowRegistrationsV1WorkflowsRegistrationsGet(context.Background()).WorkflowId(workflowId).TaskQueue(taskQueue).ActiveOnly(activeOnly).IncludeShared(includeShared).WorkflowSearch(workflowSearch).Archived(archived).WithWorkflow(withWorkflow).AvailableInChatAssistant(availableInChatAssistant).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.GetWorkflowRegistrationsV1WorkflowsRegistrationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowRegistrationsV1WorkflowsRegistrationsGet`: WorkflowRegistrationListResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.GetWorkflowRegistrationsV1WorkflowsRegistrationsGet`: %v\n", resp)
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
 **workflowSearch** | **string** | The workflow name to filter by | 
 **archived** | **bool** | Filter by archived state. False&#x3D;exclude archived, True&#x3D;only archived, None&#x3D;include all | 
 **withWorkflow** | **bool** | Whether to include the workflow definition | [default to false]
 **availableInChatAssistant** | **bool** | Whether to only return workflows compatible with chat assistant | 
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
	resp, r, err := apiClient.BetaWorkflowsAPI.GetWorkflowV1WorkflowsWorkflowIdentifierGet(context.Background(), workflowIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.GetWorkflowV1WorkflowsWorkflowIdentifierGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowV1WorkflowsWorkflowIdentifierGet`: WorkflowGetResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.GetWorkflowV1WorkflowsWorkflowIdentifierGet`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsAPI.UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut(context.Background(), workflowIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut`: WorkflowUnarchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.UnarchiveWorkflowV1WorkflowsWorkflowIdentifierUnarchivePut`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsAPI.UpdateWorkflowV1WorkflowsWorkflowIdentifierPut(context.Background(), workflowIdentifier).WorkflowUpdateRequest(workflowUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsAPI.UpdateWorkflowV1WorkflowsWorkflowIdentifierPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowV1WorkflowsWorkflowIdentifierPut`: WorkflowUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsAPI.UpdateWorkflowV1WorkflowsWorkflowIdentifierPut`: %v\n", resp)
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


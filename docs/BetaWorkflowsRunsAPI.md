# \BetaWorkflowsRunsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRunHistoryV1WorkflowsRunsRunIdHistoryGet**](BetaWorkflowsRunsAPI.md#GetRunHistoryV1WorkflowsRunsRunIdHistoryGet) | **Get** /v1/workflows/runs/{run_id}/history | Get Run History
[**GetRunV1WorkflowsRunsRunIdGet**](BetaWorkflowsRunsAPI.md#GetRunV1WorkflowsRunsRunIdGet) | **Get** /v1/workflows/runs/{run_id} | Get Run
[**ListRunsV1WorkflowsRunsGet**](BetaWorkflowsRunsAPI.md#ListRunsV1WorkflowsRunsGet) | **Get** /v1/workflows/runs | List Runs



## GetRunHistoryV1WorkflowsRunsRunIdHistoryGet

> interface{} GetRunHistoryV1WorkflowsRunsRunIdHistoryGet(ctx, runId).DecodePayloads(decodePayloads).Execute()

Get Run History

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	decodePayloads := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsRunsAPI.GetRunHistoryV1WorkflowsRunsRunIdHistoryGet(context.Background(), runId).DecodePayloads(decodePayloads).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsRunsAPI.GetRunHistoryV1WorkflowsRunsRunIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunHistoryV1WorkflowsRunsRunIdHistoryGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsRunsAPI.GetRunHistoryV1WorkflowsRunsRunIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunHistoryV1WorkflowsRunsRunIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decodePayloads** | **bool** |  | [default to false]

### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunV1WorkflowsRunsRunIdGet

> WorkflowExecutionResponse GetRunV1WorkflowsRunsRunIdGet(ctx, runId).Execute()

Get Run

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsRunsAPI.GetRunV1WorkflowsRunsRunIdGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsRunsAPI.GetRunV1WorkflowsRunsRunIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunV1WorkflowsRunsRunIdGet`: WorkflowExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsRunsAPI.GetRunV1WorkflowsRunsRunIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunV1WorkflowsRunsRunIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowExecutionResponse**](WorkflowExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRunsV1WorkflowsRunsGet

> WorkflowExecutionListResponse ListRunsV1WorkflowsRunsGet(ctx).WorkflowIdentifier(workflowIdentifier).Search(search).Status(status).PageSize(pageSize).NextPageToken(nextPageToken).Execute()

List Runs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	workflowIdentifier := "workflowIdentifier_example" // string | Filter by workflow name or id (optional)
	search := "search_example" // string | Search by workflow name, display name or id (optional)
	status := *openapiclient.NewStatus() // Status | Filter by workflow status (optional)
	pageSize := int32(56) // int32 | Number of items per page (optional) (default to 50)
	nextPageToken := "nextPageToken_example" // string | Token for the next page of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsRunsAPI.ListRunsV1WorkflowsRunsGet(context.Background()).WorkflowIdentifier(workflowIdentifier).Search(search).Status(status).PageSize(pageSize).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsRunsAPI.ListRunsV1WorkflowsRunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRunsV1WorkflowsRunsGet`: WorkflowExecutionListResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsRunsAPI.ListRunsV1WorkflowsRunsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRunsV1WorkflowsRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowIdentifier** | **string** | Filter by workflow name or id | 
 **search** | **string** | Search by workflow name, display name or id | 
 **status** | [**Status**](Status.md) | Filter by workflow status | 
 **pageSize** | **int32** | Number of items per page | [default to 50]
 **nextPageToken** | **string** | Token for the next page of results | 

### Return type

[**WorkflowExecutionListResponse**](WorkflowExecutionListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


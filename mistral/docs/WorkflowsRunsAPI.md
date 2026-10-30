# \WorkflowsRunsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRunHistoryV1WorkflowsRunsRunIdHistoryGet**](WorkflowsRunsAPI.md#GetRunHistoryV1WorkflowsRunsRunIdHistoryGet) | **Get** /v1/workflows/runs/{run_id}/history | Get Run History
[**GetRunV1WorkflowsRunsRunIdGet**](WorkflowsRunsAPI.md#GetRunV1WorkflowsRunsRunIdGet) | **Get** /v1/workflows/runs/{run_id} | Get Run
[**ListRunsV1WorkflowsRunsGet**](WorkflowsRunsAPI.md#ListRunsV1WorkflowsRunsGet) | **Get** /v1/workflows/runs | List Runs



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
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	decodePayloads := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsRunsAPI.GetRunHistoryV1WorkflowsRunsRunIdHistoryGet(context.Background(), runId).DecodePayloads(decodePayloads).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsRunsAPI.GetRunHistoryV1WorkflowsRunsRunIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunHistoryV1WorkflowsRunsRunIdHistoryGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsRunsAPI.GetRunHistoryV1WorkflowsRunsRunIdHistoryGet`: %v\n", resp)
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
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsRunsAPI.GetRunV1WorkflowsRunsRunIdGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsRunsAPI.GetRunV1WorkflowsRunsRunIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunV1WorkflowsRunsRunIdGet`: WorkflowExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsRunsAPI.GetRunV1WorkflowsRunsRunIdGet`: %v\n", resp)
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

> WorkflowExecutionListResponse ListRunsV1WorkflowsRunsGet(ctx).WorkflowIdentifier(workflowIdentifier).RootExecutionId(rootExecutionId).Search(search).Status(status).DeploymentName(deploymentName).SortBy(sortBy).Order(order).StartTimeAfter(startTimeAfter).StartTimeBefore(startTimeBefore).EndTimeAfter(endTimeAfter).EndTimeBefore(endTimeBefore).UserId(userId).WorkflowTags(workflowTags).IncludeInternal(includeInternal).PageSize(pageSize).NextPageToken(nextPageToken).SearchKey(searchKey).Execute()

List Runs



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
	workflowIdentifier := "workflowIdentifier_example" // string | Filter by workflow name or id (optional)
	rootExecutionId := "rootExecutionId_example" // string | Filter by root execution id; returns the whole execution tree (the root and all its descendant sub-workflows). (optional)
	search := "search_example" // string | Search by workflow name, display name, or ID (optional)
	status := *openapiclient.NewStatus() // Status | Filter by workflow status (optional)
	deploymentName := "deploymentName_example" // string | Filter by deployment name (optional)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	order := "order_example" // string | Sort direction (optional) (default to "desc")
	startTimeAfter := time.Now() // time.Time | Include runs with start_time >= value (optional)
	startTimeBefore := time.Now() // time.Time | Include runs with start_time <= value (optional)
	endTimeAfter := time.Now() // time.Time | Include runs with end_time >= value. Running executions (no end_time) are excluded; use the status filter to include them. (optional)
	endTimeBefore := time.Now() // time.Time | Include runs with end_time <= value. Running executions (no end_time) are excluded; use the status filter to include them. (optional)
	userId := "userId_example" // string | Filter by user id. Use 'current' to filter by the authenticated user (optional)
	workflowTags := []*string{"Inner_example"} // []*string | Filter to runs of workflows tagged with all listed tags (AND). (optional)
	includeInternal := true // bool | Include runs of internal/technical workflows (e.g. parallel-execution) (optional) (default to true)
	pageSize := int32(56) // int32 | Number of items per page (optional) (default to 50)
	nextPageToken := "nextPageToken_example" // string | Token for the next page of results (optional)
	searchKey := []string{"Inner_example"} // []string | Filter executions by search key as repeated entries (max 3, AND'd together). 'key:value' matches a similar value, 'key==:value' matches the exact value, a bare 'key' matches executions that have the key set at all, and 'key==' matches executions whose value for the key is null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsRunsAPI.ListRunsV1WorkflowsRunsGet(context.Background()).WorkflowIdentifier(workflowIdentifier).RootExecutionId(rootExecutionId).Search(search).Status(status).DeploymentName(deploymentName).SortBy(sortBy).Order(order).StartTimeAfter(startTimeAfter).StartTimeBefore(startTimeBefore).EndTimeAfter(endTimeAfter).EndTimeBefore(endTimeBefore).UserId(userId).WorkflowTags(workflowTags).IncludeInternal(includeInternal).PageSize(pageSize).NextPageToken(nextPageToken).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsRunsAPI.ListRunsV1WorkflowsRunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRunsV1WorkflowsRunsGet`: WorkflowExecutionListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsRunsAPI.ListRunsV1WorkflowsRunsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRunsV1WorkflowsRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowIdentifier** | **string** | Filter by workflow name or id | 
 **rootExecutionId** | **string** | Filter by root execution id; returns the whole execution tree (the root and all its descendant sub-workflows). | 
 **search** | **string** | Search by workflow name, display name, or ID | 
 **status** | [**Status**](Status.md) | Filter by workflow status | 
 **deploymentName** | **string** | Filter by deployment name | 
 **sortBy** | **string** | Field to sort by | 
 **order** | **string** | Sort direction | [default to &quot;desc&quot;]
 **startTimeAfter** | **time.Time** | Include runs with start_time &gt;&#x3D; value | 
 **startTimeBefore** | **time.Time** | Include runs with start_time &lt;&#x3D; value | 
 **endTimeAfter** | **time.Time** | Include runs with end_time &gt;&#x3D; value. Running executions (no end_time) are excluded; use the status filter to include them. | 
 **endTimeBefore** | **time.Time** | Include runs with end_time &lt;&#x3D; value. Running executions (no end_time) are excluded; use the status filter to include them. | 
 **userId** | **string** | Filter by user id. Use &#39;current&#39; to filter by the authenticated user | 
 **workflowTags** | **[]string** | Filter to runs of workflows tagged with all listed tags (AND). | 
 **includeInternal** | **bool** | Include runs of internal/technical workflows (e.g. parallel-execution) | [default to true]
 **pageSize** | **int32** | Number of items per page | [default to 50]
 **nextPageToken** | **string** | Token for the next page of results | 
 **searchKey** | **[]string** | Filter executions by search key as repeated entries (max 3, AND&#39;d together). &#39;key:value&#39; matches a similar value, &#39;key&#x3D;&#x3D;:value&#39; matches the exact value, a bare &#39;key&#39; matches executions that have the key set at all, and &#39;key&#x3D;&#x3D;&#39; matches executions whose value for the key is null. | 

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


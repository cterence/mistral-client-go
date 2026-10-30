# \WorkflowsExecutionsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost**](WorkflowsExecutionsAPI.md#BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost) | **Post** /v1/workflows/executions/cancel | Batch Cancel Workflow Executions
[**BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost**](WorkflowsExecutionsAPI.md#BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost) | **Post** /v1/workflows/executions/terminate | Batch Terminate Workflow Executions
[**CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost**](WorkflowsExecutionsAPI.md#CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost) | **Post** /v1/workflows/executions/{execution_id}/cancel | Cancel Workflow Execution
[**GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet) | **Get** /v1/workflows/executions/{execution_id}/history | Get Workflow Execution History
[**GetWorkflowExecutionLogs**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionLogs) | **Get** /v1/workflows/executions/{execution_id}/logs | Get Workflow Execution Logs
[**GetWorkflowExecutionTraceEvents**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceEvents) | **Get** /v1/workflows/executions/{execution_id}/trace/events | Get Workflow Execution Trace Events
[**GetWorkflowExecutionTraceInfo**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceInfo) | **Get** /v1/workflows/executions/{execution_id}/trace/info | Get Workflow Execution Trace Info
[**GetWorkflowExecutionTraceOtel**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceOtel) | **Get** /v1/workflows/executions/{execution_id}/trace/otel | Get Workflow Execution Trace Otel
[**GetWorkflowExecutionTraceSummary**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceSummary) | **Get** /v1/workflows/executions/{execution_id}/trace/summary | Get Workflow Execution Trace Summary
[**GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet**](WorkflowsExecutionsAPI.md#GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet) | **Get** /v1/workflows/executions/{execution_id} | Get Workflow Execution
[**QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost**](WorkflowsExecutionsAPI.md#QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost) | **Post** /v1/workflows/executions/{execution_id}/queries | Query Workflow Execution
[**ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost**](WorkflowsExecutionsAPI.md#ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost) | **Post** /v1/workflows/executions/{execution_id}/reset | Reset Workflow
[**SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost**](WorkflowsExecutionsAPI.md#SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost) | **Post** /v1/workflows/executions/{execution_id}/signals | Signal Workflow Execution
[**StreamV1WorkflowsExecutionsExecutionIdStreamGet**](WorkflowsExecutionsAPI.md#StreamV1WorkflowsExecutionsExecutionIdStreamGet) | **Get** /v1/workflows/executions/{execution_id}/stream | Stream
[**StreamWorkflowExecutionLogs**](WorkflowsExecutionsAPI.md#StreamWorkflowExecutionLogs) | **Get** /v1/workflows/executions/{execution_id}/logs/stream | Stream Workflow Execution Logs
[**TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost**](WorkflowsExecutionsAPI.md#TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost) | **Post** /v1/workflows/executions/{execution_id}/terminate | Terminate Workflow Execution
[**UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost**](WorkflowsExecutionsAPI.md#UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost) | **Post** /v1/workflows/executions/{execution_id}/updates | Update Workflow Execution



## BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost

> BatchExecutionResponse BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost(ctx).BatchExecutionBody(batchExecutionBody).Execute()

Batch Cancel Workflow Executions



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
	batchExecutionBody := *openapiclient.NewBatchExecutionBody([]string{"ExecutionIds_example"}) // BatchExecutionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost(context.Background()).BatchExecutionBody(batchExecutionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost`: BatchExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchExecutionBody** | [**BatchExecutionBody**](BatchExecutionBody.md) |  | 

### Return type

[**BatchExecutionResponse**](BatchExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost

> BatchExecutionResponse BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost(ctx).BatchExecutionBody(batchExecutionBody).Execute()

Batch Terminate Workflow Executions



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
	batchExecutionBody := *openapiclient.NewBatchExecutionBody([]string{"ExecutionIds_example"}) // BatchExecutionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost(context.Background()).BatchExecutionBody(batchExecutionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost`: BatchExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchExecutionBody** | [**BatchExecutionBody**](BatchExecutionBody.md) |  | 

### Return type

[**BatchExecutionResponse**](BatchExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost

> CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost(ctx, executionId).Execute()

Cancel Workflow Execution



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
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsExecutionsAPI.CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet

> interface{} GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet(ctx, executionId).DecodePayloads(decodePayloads).Execute()

Get Workflow Execution History



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
	executionId := "executionId_example" // string | 
	decodePayloads := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet(context.Background(), executionId).DecodePayloads(decodePayloads).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGetRequest struct via the builder pattern


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


## GetWorkflowExecutionLogs

> ExecutionLogSearchResponse GetWorkflowExecutionLogs(ctx, executionId).RunId(runId).ActivityId(activityId).After(after).Before(before).Order(order).Cursor(cursor).Limit(limit).Execute()

Get Workflow Execution Logs



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
	executionId := "executionId_example" // string | 
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter logs by workflow run ID (optional)
	activityId := "activityId_example" // string | Filter logs by activity ID (optional)
	after := time.Now() // time.Time | Only return logs at or after this timestamp (optional)
	before := time.Now() // time.Time | Only return logs before this timestamp (optional)
	order := "order_example" // string | First-page sort order: 'asc' (oldest first) or 'desc'. Ignored when `cursor` is set. (optional) (default to "asc")
	cursor := "cursor_example" // string | Pagination cursor from a previous response's `next_cursor`; carries the window and order (optional)
	limit := int32(56) // int32 | Maximum number of logs to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionLogs(context.Background(), executionId).RunId(runId).ActivityId(activityId).After(after).Before(before).Order(order).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionLogs`: ExecutionLogSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runId** | **string** | Filter logs by workflow run ID | 
 **activityId** | **string** | Filter logs by activity ID | 
 **after** | **time.Time** | Only return logs at or after this timestamp | 
 **before** | **time.Time** | Only return logs before this timestamp | 
 **order** | **string** | First-page sort order: &#39;asc&#39; (oldest first) or &#39;desc&#39;. Ignored when &#x60;cursor&#x60; is set. | [default to &quot;asc&quot;]
 **cursor** | **string** | Pagination cursor from a previous response&#39;s &#x60;next_cursor&#x60;; carries the window and order | 
 **limit** | **int32** | Maximum number of logs to return | [default to 50]

### Return type

[**ExecutionLogSearchResponse**](ExecutionLogSearchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionTraceEvents

> WorkflowExecutionTraceEventsResponse GetWorkflowExecutionTraceEvents(ctx, executionId).MergeSameIdEvents(mergeSameIdEvents).IncludeInternalEvents(includeInternalEvents).Execute()

Get Workflow Execution Trace Events



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
	executionId := "executionId_example" // string | 
	mergeSameIdEvents := true // bool |  (optional) (default to false)
	includeInternalEvents := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionTraceEvents(context.Background(), executionId).MergeSameIdEvents(mergeSameIdEvents).IncludeInternalEvents(includeInternalEvents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceEvents`: WorkflowExecutionTraceEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionTraceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mergeSameIdEvents** | **bool** |  | [default to false]
 **includeInternalEvents** | **bool** |  | [default to false]

### Return type

[**WorkflowExecutionTraceEventsResponse**](WorkflowExecutionTraceEventsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionTraceInfo

> ExecutionTraceInfoResponse GetWorkflowExecutionTraceInfo(ctx, executionId).Execute()

Get Workflow Execution Trace Info



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
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionTraceInfo(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceInfo`: ExecutionTraceInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionTraceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExecutionTraceInfoResponse**](ExecutionTraceInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionTraceOtel

> WorkflowExecutionTraceOTelResponse GetWorkflowExecutionTraceOtel(ctx, executionId).Execute()

Get Workflow Execution Trace Otel



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
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionTraceOtel(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceOtel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceOtel`: WorkflowExecutionTraceOTelResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceOtel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionTraceOtelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowExecutionTraceOTelResponse**](WorkflowExecutionTraceOTelResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionTraceSummary

> WorkflowExecutionTraceSummaryResponse GetWorkflowExecutionTraceSummary(ctx, executionId).Execute()

Get Workflow Execution Trace Summary



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
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionTraceSummary(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceSummary`: WorkflowExecutionTraceSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionTraceSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionTraceSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowExecutionTraceSummaryResponse**](WorkflowExecutionTraceSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet

> WorkflowExecutionResponse GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet(ctx, executionId).IncludeSearchKeys(includeSearchKeys).Execute()

Get Workflow Execution



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
	executionId := "executionId_example" // string | 
	includeSearchKeys := true // bool | Include the execution's search keys (metadata) in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet(context.Background(), executionId).IncludeSearchKeys(includeSearchKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet`: WorkflowExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSearchKeys** | **bool** | Include the execution&#39;s search keys (metadata) in the response. | [default to false]

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


## QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost

> QueryWorkflowResponse QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost(ctx, executionId).QueryInvocationBody(queryInvocationBody).Execute()

Query Workflow Execution



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
	executionId := "executionId_example" // string | 
	queryInvocationBody := *openapiclient.NewQueryInvocationBody("Name_example") // QueryInvocationBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost(context.Background(), executionId).QueryInvocationBody(queryInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost`: QueryWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryInvocationBody** | [**QueryInvocationBody**](QueryInvocationBody.md) |  | 

### Return type

[**QueryWorkflowResponse**](QueryWorkflowResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost

> ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost(ctx, executionId).ResetInvocationBody(resetInvocationBody).Execute()

Reset Workflow



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
	executionId := "executionId_example" // string | 
	resetInvocationBody := *openapiclient.NewResetInvocationBody(int32(123)) // ResetInvocationBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsExecutionsAPI.ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost(context.Background(), executionId).ResetInvocationBody(resetInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetWorkflowV1WorkflowsExecutionsExecutionIdResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetInvocationBody** | [**ResetInvocationBody**](ResetInvocationBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost

> SignalWorkflowResponse SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost(ctx, executionId).SignalInvocationBody(signalInvocationBody).Execute()

Signal Workflow Execution



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
	executionId := "executionId_example" // string | 
	signalInvocationBody := *openapiclient.NewSignalInvocationBody("Name_example") // SignalInvocationBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost(context.Background(), executionId).SignalInvocationBody(signalInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost`: SignalWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signalInvocationBody** | [**SignalInvocationBody**](SignalInvocationBody.md) |  | 

### Return type

[**SignalWorkflowResponse**](SignalWorkflowResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamV1WorkflowsExecutionsExecutionIdStreamGet

> StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response StreamV1WorkflowsExecutionsExecutionIdStreamGet(ctx, executionId).EventSource(eventSource).LastEventId(lastEventId).Execute()

Stream



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
	executionId := "executionId_example" // string | 
	eventSource := openapiclient.EventSource("DATABASE") // EventSource |  (optional)
	lastEventId := "lastEventId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.StreamV1WorkflowsExecutionsExecutionIdStreamGet(context.Background(), executionId).EventSource(eventSource).LastEventId(lastEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.StreamV1WorkflowsExecutionsExecutionIdStreamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamV1WorkflowsExecutionsExecutionIdStreamGet`: StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.StreamV1WorkflowsExecutionsExecutionIdStreamGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamV1WorkflowsExecutionsExecutionIdStreamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventSource** | [**EventSource**](EventSource.md) |  | 
 **lastEventId** | **string** |  | 

### Return type

[**StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response**](StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamWorkflowExecutionLogs

> StreamWorkflowExecutionLogs200Response StreamWorkflowExecutionLogs(ctx, executionId).RunId(runId).ActivityId(activityId).After(after).LastEventId(lastEventId).LastEventID(lastEventID).Execute()

Stream Workflow Execution Logs



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
	executionId := "executionId_example" // string | 
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter logs by workflow run ID (optional)
	activityId := "activityId_example" // string | Filter logs by activity ID (optional)
	after := time.Now() // time.Time | Start a fresh stream at this timestamp (ignored when resuming via last_event_id) (optional)
	lastEventId := "lastEventId_example" // string | Resume from this cursor (a prior response's SSE id) (optional)
	lastEventID := "lastEventID_example" // string | Resume from this cursor (a prior response's SSE id). Takes precedence over the query parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.StreamWorkflowExecutionLogs(context.Background(), executionId).RunId(runId).ActivityId(activityId).After(after).LastEventId(lastEventId).LastEventID(lastEventID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.StreamWorkflowExecutionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamWorkflowExecutionLogs`: StreamWorkflowExecutionLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.StreamWorkflowExecutionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamWorkflowExecutionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runId** | **string** | Filter logs by workflow run ID | 
 **activityId** | **string** | Filter logs by activity ID | 
 **after** | **time.Time** | Start a fresh stream at this timestamp (ignored when resuming via last_event_id) | 
 **lastEventId** | **string** | Resume from this cursor (a prior response&#39;s SSE id) | 
 **lastEventID** | **string** | Resume from this cursor (a prior response&#39;s SSE id). Takes precedence over the query parameter. | 

### Return type

[**StreamWorkflowExecutionLogs200Response**](StreamWorkflowExecutionLogs200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost

> TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost(ctx, executionId).Execute()

Terminate Workflow Execution



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
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsExecutionsAPI.TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost

> UpdateWorkflowResponse UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost(ctx, executionId).UpdateInvocationBody(updateInvocationBody).Execute()

Update Workflow Execution



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
	executionId := "executionId_example" // string | 
	updateInvocationBody := *openapiclient.NewUpdateInvocationBody("Name_example") // UpdateInvocationBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsExecutionsAPI.UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost(context.Background(), executionId).UpdateInvocationBody(updateInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsExecutionsAPI.UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost`: UpdateWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsExecutionsAPI.UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInvocationBody** | [**UpdateInvocationBody**](UpdateInvocationBody.md) |  | 

### Return type

[**UpdateWorkflowResponse**](UpdateWorkflowResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


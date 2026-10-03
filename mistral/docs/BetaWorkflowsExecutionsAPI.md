# \BetaWorkflowsExecutionsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost**](BetaWorkflowsExecutionsAPI.md#BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost) | **Post** /v1/workflows/executions/cancel | Batch Cancel Workflow Executions
[**BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost**](BetaWorkflowsExecutionsAPI.md#BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost) | **Post** /v1/workflows/executions/terminate | Batch Terminate Workflow Executions
[**CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost**](BetaWorkflowsExecutionsAPI.md#CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost) | **Post** /v1/workflows/executions/{execution_id}/cancel | Cancel Workflow Execution
[**GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet**](BetaWorkflowsExecutionsAPI.md#GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet) | **Get** /v1/workflows/executions/{execution_id}/history | Get Workflow Execution History
[**GetWorkflowExecutionTraceEvents**](BetaWorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceEvents) | **Get** /v1/workflows/executions/{execution_id}/trace/events | Get Workflow Execution Trace Events
[**GetWorkflowExecutionTraceOtel**](BetaWorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceOtel) | **Get** /v1/workflows/executions/{execution_id}/trace/otel | Get Workflow Execution Trace Otel
[**GetWorkflowExecutionTraceSummary**](BetaWorkflowsExecutionsAPI.md#GetWorkflowExecutionTraceSummary) | **Get** /v1/workflows/executions/{execution_id}/trace/summary | Get Workflow Execution Trace Summary
[**GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet**](BetaWorkflowsExecutionsAPI.md#GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet) | **Get** /v1/workflows/executions/{execution_id} | Get Workflow Execution
[**QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost**](BetaWorkflowsExecutionsAPI.md#QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost) | **Post** /v1/workflows/executions/{execution_id}/queries | Query Workflow Execution
[**ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost**](BetaWorkflowsExecutionsAPI.md#ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost) | **Post** /v1/workflows/executions/{execution_id}/reset | Reset Workflow
[**SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost**](BetaWorkflowsExecutionsAPI.md#SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost) | **Post** /v1/workflows/executions/{execution_id}/signals | Signal Workflow Execution
[**StreamV1WorkflowsExecutionsExecutionIdStreamGet**](BetaWorkflowsExecutionsAPI.md#StreamV1WorkflowsExecutionsExecutionIdStreamGet) | **Get** /v1/workflows/executions/{execution_id}/stream | Stream
[**TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost**](BetaWorkflowsExecutionsAPI.md#TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost) | **Post** /v1/workflows/executions/{execution_id}/terminate | Terminate Workflow Execution
[**UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost**](BetaWorkflowsExecutionsAPI.md#UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost) | **Post** /v1/workflows/executions/{execution_id}/updates | Update Workflow Execution



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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost(context.Background()).BatchExecutionBody(batchExecutionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost`: BatchExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.BatchCancelWorkflowExecutionsV1WorkflowsExecutionsCancelPost`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost(context.Background()).BatchExecutionBody(batchExecutionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost`: BatchExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.BatchTerminateWorkflowExecutionsV1WorkflowsExecutionsTerminatePost`: %v\n", resp)
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
	r, err := apiClient.BetaWorkflowsExecutionsAPI.CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.CancelWorkflowExecutionV1WorkflowsExecutionsExecutionIdCancelPost``: %v\n", err)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet(context.Background(), executionId).DecodePayloads(decodePayloads).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionHistoryV1WorkflowsExecutionsExecutionIdHistoryGet`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceEvents(context.Background(), executionId).MergeSameIdEvents(mergeSameIdEvents).IncludeInternalEvents(includeInternalEvents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceEvents`: WorkflowExecutionTraceEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceEvents`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceOtel(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceOtel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceOtel`: WorkflowExecutionTraceOTelResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceOtel`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceSummary(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionTraceSummary`: WorkflowExecutionTraceSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionTraceSummary`: %v\n", resp)
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

> WorkflowExecutionResponse GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet(ctx, executionId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet`: WorkflowExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.GetWorkflowExecutionV1WorkflowsExecutionsExecutionIdGet`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost(context.Background(), executionId).QueryInvocationBody(queryInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost`: QueryWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.QueryWorkflowExecutionV1WorkflowsExecutionsExecutionIdQueriesPost`: %v\n", resp)
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
	r, err := apiClient.BetaWorkflowsExecutionsAPI.ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost(context.Background(), executionId).ResetInvocationBody(resetInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.ResetWorkflowV1WorkflowsExecutionsExecutionIdResetPost``: %v\n", err)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost(context.Background(), executionId).SignalInvocationBody(signalInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost`: SignalWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.SignalWorkflowExecutionV1WorkflowsExecutionsExecutionIdSignalsPost`: %v\n", resp)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.StreamV1WorkflowsExecutionsExecutionIdStreamGet(context.Background(), executionId).EventSource(eventSource).LastEventId(lastEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.StreamV1WorkflowsExecutionsExecutionIdStreamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamV1WorkflowsExecutionsExecutionIdStreamGet`: StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.StreamV1WorkflowsExecutionsExecutionIdStreamGet`: %v\n", resp)
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
	r, err := apiClient.BetaWorkflowsExecutionsAPI.TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost(context.Background(), executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.TerminateWorkflowExecutionV1WorkflowsExecutionsExecutionIdTerminatePost``: %v\n", err)
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
	resp, r, err := apiClient.BetaWorkflowsExecutionsAPI.UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost(context.Background(), executionId).UpdateInvocationBody(updateInvocationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsExecutionsAPI.UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost`: UpdateWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsExecutionsAPI.UpdateWorkflowExecutionV1WorkflowsExecutionsExecutionIdUpdatesPost`: %v\n", resp)
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


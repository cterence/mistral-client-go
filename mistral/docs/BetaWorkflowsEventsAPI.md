# \BetaWorkflowsEventsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStreamEventsV1WorkflowsEventsStreamGet**](BetaWorkflowsEventsAPI.md#GetStreamEventsV1WorkflowsEventsStreamGet) | **Get** /v1/workflows/events/stream | Get Stream Events
[**GetWorkflowEventsV1WorkflowsEventsListGet**](BetaWorkflowsEventsAPI.md#GetWorkflowEventsV1WorkflowsEventsListGet) | **Get** /v1/workflows/events/list | Get Workflow Events



## GetStreamEventsV1WorkflowsEventsStreamGet

> StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response GetStreamEventsV1WorkflowsEventsStreamGet(ctx).Scope(scope).ActivityName(activityName).ActivityId(activityId).WorkflowName(workflowName).WorkflowExecId(workflowExecId).RootWorkflowExecId(rootWorkflowExecId).ParentWorkflowExecId(parentWorkflowExecId).Stream(stream).StartSeq(startSeq).MetadataFilters(metadataFilters).WorkflowEventTypes(workflowEventTypes).LastEventId(lastEventId).Execute()

Get Stream Events

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
	scope := "scope_example" // string |  (optional) (default to "*")
	activityName := "activityName_example" // string |  (optional) (default to "*")
	activityId := "activityId_example" // string |  (optional) (default to "*")
	workflowName := "workflowName_example" // string |  (optional) (default to "*")
	workflowExecId := "workflowExecId_example" // string |  (optional) (default to "*")
	rootWorkflowExecId := "rootWorkflowExecId_example" // string |  (optional) (default to "*")
	parentWorkflowExecId := "parentWorkflowExecId_example" // string |  (optional) (default to "*")
	stream := "stream_example" // string |  (optional) (default to "*")
	startSeq := int32(56) // int32 |  (optional) (default to 0)
	metadataFilters := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)
	workflowEventTypes := []openapiclient.WorkflowEventType{openapiclient.WorkflowEventType("WORKFLOW_EXECUTION_STARTED")} // []WorkflowEventType |  (optional)
	lastEventId := "lastEventId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsEventsAPI.GetStreamEventsV1WorkflowsEventsStreamGet(context.Background()).Scope(scope).ActivityName(activityName).ActivityId(activityId).WorkflowName(workflowName).WorkflowExecId(workflowExecId).RootWorkflowExecId(rootWorkflowExecId).ParentWorkflowExecId(parentWorkflowExecId).Stream(stream).StartSeq(startSeq).MetadataFilters(metadataFilters).WorkflowEventTypes(workflowEventTypes).LastEventId(lastEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsEventsAPI.GetStreamEventsV1WorkflowsEventsStreamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamEventsV1WorkflowsEventsStreamGet`: StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsEventsAPI.GetStreamEventsV1WorkflowsEventsStreamGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamEventsV1WorkflowsEventsStreamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** |  | [default to &quot;*&quot;]
 **activityName** | **string** |  | [default to &quot;*&quot;]
 **activityId** | **string** |  | [default to &quot;*&quot;]
 **workflowName** | **string** |  | [default to &quot;*&quot;]
 **workflowExecId** | **string** |  | [default to &quot;*&quot;]
 **rootWorkflowExecId** | **string** |  | [default to &quot;*&quot;]
 **parentWorkflowExecId** | **string** |  | [default to &quot;*&quot;]
 **stream** | **string** |  | [default to &quot;*&quot;]
 **startSeq** | **int32** |  | [default to 0]
 **metadataFilters** | **map[string]interface{}** |  | 
 **workflowEventTypes** | [**[]WorkflowEventType**](WorkflowEventType.md) |  | 
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


## GetWorkflowEventsV1WorkflowsEventsListGet

> ListWorkflowEventResponse GetWorkflowEventsV1WorkflowsEventsListGet(ctx).RootWorkflowExecId(rootWorkflowExecId).WorkflowExecId(workflowExecId).WorkflowRunId(workflowRunId).Limit(limit).Cursor(cursor).Execute()

Get Workflow Events

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
	rootWorkflowExecId := "rootWorkflowExecId_example" // string | Execution ID of the root workflow that initiated this execution chain. (optional)
	workflowExecId := "workflowExecId_example" // string | Execution ID of the workflow that emitted this event. (optional)
	workflowRunId := "workflowRunId_example" // string | Run ID of the workflow that emitted this event. (optional)
	limit := int32(56) // int32 | Maximum number of events to return. (optional) (default to 100)
	cursor := "cursor_example" // string | Cursor for pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsEventsAPI.GetWorkflowEventsV1WorkflowsEventsListGet(context.Background()).RootWorkflowExecId(rootWorkflowExecId).WorkflowExecId(workflowExecId).WorkflowRunId(workflowRunId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsEventsAPI.GetWorkflowEventsV1WorkflowsEventsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowEventsV1WorkflowsEventsListGet`: ListWorkflowEventResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsEventsAPI.GetWorkflowEventsV1WorkflowsEventsListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowEventsV1WorkflowsEventsListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootWorkflowExecId** | **string** | Execution ID of the root workflow that initiated this execution chain. | 
 **workflowExecId** | **string** | Execution ID of the workflow that emitted this event. | 
 **workflowRunId** | **string** | Run ID of the workflow that emitted this event. | 
 **limit** | **int32** | Maximum number of events to return. | [default to 100]
 **cursor** | **string** | Cursor for pagination. | 

### Return type

[**ListWorkflowEventResponse**](ListWorkflowEventResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \BetaWorkflowsSchedulesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchedulesV1WorkflowsSchedulesGet**](BetaWorkflowsSchedulesAPI.md#GetSchedulesV1WorkflowsSchedulesGet) | **Get** /v1/workflows/schedules | Get Schedules
[**ScheduleWorkflowV1WorkflowsSchedulesPost**](BetaWorkflowsSchedulesAPI.md#ScheduleWorkflowV1WorkflowsSchedulesPost) | **Post** /v1/workflows/schedules | Schedule Workflow
[**UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete**](BetaWorkflowsSchedulesAPI.md#UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete) | **Delete** /v1/workflows/schedules/{schedule_id} | Unschedule Workflow



## GetSchedulesV1WorkflowsSchedulesGet

> WorkflowScheduleListResponse GetSchedulesV1WorkflowsSchedulesGet(ctx).Execute()

Get Schedules

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsSchedulesAPI.GetSchedulesV1WorkflowsSchedulesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsSchedulesAPI.GetSchedulesV1WorkflowsSchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulesV1WorkflowsSchedulesGet`: WorkflowScheduleListResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsSchedulesAPI.GetSchedulesV1WorkflowsSchedulesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulesV1WorkflowsSchedulesGetRequest struct via the builder pattern


### Return type

[**WorkflowScheduleListResponse**](WorkflowScheduleListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleWorkflowV1WorkflowsSchedulesPost

> WorkflowScheduleResponse ScheduleWorkflowV1WorkflowsSchedulesPost(ctx).WorkflowScheduleRequest(workflowScheduleRequest).Execute()

Schedule Workflow

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
	workflowScheduleRequest := *openapiclient.NewWorkflowScheduleRequest(*openapiclient.NewScheduleDefinition(interface{}(123))) // WorkflowScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsSchedulesAPI.ScheduleWorkflowV1WorkflowsSchedulesPost(context.Background()).WorkflowScheduleRequest(workflowScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsSchedulesAPI.ScheduleWorkflowV1WorkflowsSchedulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleWorkflowV1WorkflowsSchedulesPost`: WorkflowScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsSchedulesAPI.ScheduleWorkflowV1WorkflowsSchedulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleWorkflowV1WorkflowsSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowScheduleRequest** | [**WorkflowScheduleRequest**](WorkflowScheduleRequest.md) |  | 

### Return type

[**WorkflowScheduleResponse**](WorkflowScheduleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete

> UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete(ctx, scheduleId).Execute()

Unschedule Workflow

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
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaWorkflowsSchedulesAPI.UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsSchedulesAPI.UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDeleteRequest struct via the builder pattern


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


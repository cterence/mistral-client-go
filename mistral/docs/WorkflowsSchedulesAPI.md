# \WorkflowsSchedulesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScheduleV1WorkflowsSchedulesScheduleIdGet**](WorkflowsSchedulesAPI.md#GetScheduleV1WorkflowsSchedulesScheduleIdGet) | **Get** /v1/workflows/schedules/{schedule_id} | Get Schedule
[**GetSchedulesV1WorkflowsSchedulesGet**](WorkflowsSchedulesAPI.md#GetSchedulesV1WorkflowsSchedulesGet) | **Get** /v1/workflows/schedules | Get Schedules
[**PauseScheduleV1WorkflowsSchedulesScheduleIdPausePost**](WorkflowsSchedulesAPI.md#PauseScheduleV1WorkflowsSchedulesScheduleIdPausePost) | **Post** /v1/workflows/schedules/{schedule_id}/pause | Pause Schedule
[**ResumeScheduleV1WorkflowsSchedulesScheduleIdResumePost**](WorkflowsSchedulesAPI.md#ResumeScheduleV1WorkflowsSchedulesScheduleIdResumePost) | **Post** /v1/workflows/schedules/{schedule_id}/resume | Resume Schedule
[**ScheduleWorkflowV1WorkflowsSchedulesPost**](WorkflowsSchedulesAPI.md#ScheduleWorkflowV1WorkflowsSchedulesPost) | **Post** /v1/workflows/schedules | Schedule Workflow
[**TriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPost**](WorkflowsSchedulesAPI.md#TriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPost) | **Post** /v1/workflows/schedules/{schedule_id}/trigger | Trigger Schedule
[**UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete**](WorkflowsSchedulesAPI.md#UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete) | **Delete** /v1/workflows/schedules/{schedule_id} | Unschedule Workflow
[**UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch**](WorkflowsSchedulesAPI.md#UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch) | **Patch** /v1/workflows/schedules/{schedule_id} | Update Schedule



## GetScheduleV1WorkflowsSchedulesScheduleIdGet

> ScheduleDefinitionOutput GetScheduleV1WorkflowsSchedulesScheduleIdGet(ctx, scheduleId).Execute()

Get Schedule



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
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsSchedulesAPI.GetScheduleV1WorkflowsSchedulesScheduleIdGet(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.GetScheduleV1WorkflowsSchedulesScheduleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduleV1WorkflowsSchedulesScheduleIdGet`: ScheduleDefinitionOutput
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsSchedulesAPI.GetScheduleV1WorkflowsSchedulesScheduleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleV1WorkflowsSchedulesScheduleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduleDefinitionOutput**](ScheduleDefinitionOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedulesV1WorkflowsSchedulesGet

> WorkflowScheduleListResponse GetSchedulesV1WorkflowsSchedulesGet(ctx).WorkflowName(workflowName).UserId(userId).Status(status).Search(search).PageSize(pageSize).NextPageToken(nextPageToken).Execute()

Get Schedules



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
	workflowName := "workflowName_example" // string | Filter by exact workflow name (optional)
	userId := "userId_example" // string | Filter by user ID. Pass 'current' to resolve to the authenticated user's ID. (optional)
	status := "status_example" // string | Filter by schedule status: 'active' or 'paused' (optional)
	search := "search_example" // string | Prefix search query for workflow name (optional)
	pageSize := int32(56) // int32 | Number of items per page. Omitting this parameter fetches all results at once (deprecated — pass page_size to use pagination). (optional)
	nextPageToken := "nextPageToken_example" // string | Token for the next page of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsSchedulesAPI.GetSchedulesV1WorkflowsSchedulesGet(context.Background()).WorkflowName(workflowName).UserId(userId).Status(status).Search(search).PageSize(pageSize).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.GetSchedulesV1WorkflowsSchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulesV1WorkflowsSchedulesGet`: WorkflowScheduleListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsSchedulesAPI.GetSchedulesV1WorkflowsSchedulesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulesV1WorkflowsSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowName** | **string** | Filter by exact workflow name | 
 **userId** | **string** | Filter by user ID. Pass &#39;current&#39; to resolve to the authenticated user&#39;s ID. | 
 **status** | **string** | Filter by schedule status: &#39;active&#39; or &#39;paused&#39; | 
 **search** | **string** | Prefix search query for workflow name | 
 **pageSize** | **int32** | Number of items per page. Omitting this parameter fetches all results at once (deprecated — pass page_size to use pagination). | 
 **nextPageToken** | **string** | Token for the next page of results | 

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


## PauseScheduleV1WorkflowsSchedulesScheduleIdPausePost

> PauseScheduleV1WorkflowsSchedulesScheduleIdPausePost(ctx, scheduleId).WorkflowSchedulePauseRequest(workflowSchedulePauseRequest).Execute()

Pause Schedule



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
	scheduleId := "scheduleId_example" // string | 
	workflowSchedulePauseRequest := *openapiclient.NewWorkflowSchedulePauseRequest() // WorkflowSchedulePauseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsSchedulesAPI.PauseScheduleV1WorkflowsSchedulesScheduleIdPausePost(context.Background(), scheduleId).WorkflowSchedulePauseRequest(workflowSchedulePauseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.PauseScheduleV1WorkflowsSchedulesScheduleIdPausePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPauseScheduleV1WorkflowsSchedulesScheduleIdPausePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowSchedulePauseRequest** | [**WorkflowSchedulePauseRequest**](WorkflowSchedulePauseRequest.md) |  | 

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


## ResumeScheduleV1WorkflowsSchedulesScheduleIdResumePost

> ResumeScheduleV1WorkflowsSchedulesScheduleIdResumePost(ctx, scheduleId).WorkflowSchedulePauseRequest(workflowSchedulePauseRequest).Execute()

Resume Schedule



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
	scheduleId := "scheduleId_example" // string | 
	workflowSchedulePauseRequest := *openapiclient.NewWorkflowSchedulePauseRequest() // WorkflowSchedulePauseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsSchedulesAPI.ResumeScheduleV1WorkflowsSchedulesScheduleIdResumePost(context.Background(), scheduleId).WorkflowSchedulePauseRequest(workflowSchedulePauseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.ResumeScheduleV1WorkflowsSchedulesScheduleIdResumePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResumeScheduleV1WorkflowsSchedulesScheduleIdResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowSchedulePauseRequest** | [**WorkflowSchedulePauseRequest**](WorkflowSchedulePauseRequest.md) |  | 

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
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	workflowScheduleRequest := *openapiclient.NewWorkflowScheduleRequest(*openapiclient.NewScheduleDefinition(interface{}(123))) // WorkflowScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsSchedulesAPI.ScheduleWorkflowV1WorkflowsSchedulesPost(context.Background()).WorkflowScheduleRequest(workflowScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.ScheduleWorkflowV1WorkflowsSchedulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleWorkflowV1WorkflowsSchedulesPost`: WorkflowScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsSchedulesAPI.ScheduleWorkflowV1WorkflowsSchedulesPost`: %v\n", resp)
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


## TriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPost

> TriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPost(ctx, scheduleId).WorkflowScheduleTriggerRequest(workflowScheduleTriggerRequest).Execute()

Trigger Schedule



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
	scheduleId := "scheduleId_example" // string | 
	workflowScheduleTriggerRequest := *openapiclient.NewWorkflowScheduleTriggerRequest() // WorkflowScheduleTriggerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsSchedulesAPI.TriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPost(context.Background(), scheduleId).WorkflowScheduleTriggerRequest(workflowScheduleTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.TriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTriggerScheduleV1WorkflowsSchedulesScheduleIdTriggerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowScheduleTriggerRequest** | [**WorkflowScheduleTriggerRequest**](WorkflowScheduleTriggerRequest.md) |  | 

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
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	scheduleId := "scheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsSchedulesAPI.UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.UnscheduleWorkflowV1WorkflowsSchedulesScheduleIdDelete``: %v\n", err)
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


## UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch

> WorkflowScheduleResponse UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch(ctx, scheduleId).WorkflowScheduleUpdateRequest(workflowScheduleUpdateRequest).Execute()

Update Schedule



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
	scheduleId := "scheduleId_example" // string | 
	workflowScheduleUpdateRequest := *openapiclient.NewWorkflowScheduleUpdateRequest(*openapiclient.NewPartialScheduleDefinition()) // WorkflowScheduleUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsSchedulesAPI.UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch(context.Background(), scheduleId).WorkflowScheduleUpdateRequest(workflowScheduleUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsSchedulesAPI.UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch`: WorkflowScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsSchedulesAPI.UpdateScheduleV1WorkflowsSchedulesScheduleIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduleV1WorkflowsSchedulesScheduleIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowScheduleUpdateRequest** | [**WorkflowScheduleUpdateRequest**](WorkflowScheduleUpdateRequest.md) |  | 

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


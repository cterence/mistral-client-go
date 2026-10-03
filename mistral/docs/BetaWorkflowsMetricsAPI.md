# \BetaWorkflowsMetricsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet**](BetaWorkflowsMetricsAPI.md#GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet) | **Get** /v1/workflows/{workflow_name}/metrics | Get Workflow Metrics



## GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet

> WorkflowMetrics GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet(ctx, workflowName).StartTime(startTime).EndTime(endTime).Execute()

Get Workflow Metrics



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
	workflowName := "workflowName_example" // string | 
	startTime := time.Now() // time.Time | Filter workflows started after this time (ISO 8601) (optional)
	endTime := time.Now() // time.Time | Filter workflows started before this time (ISO 8601) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaWorkflowsMetricsAPI.GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet(context.Background(), workflowName).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsMetricsAPI.GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet`: WorkflowMetrics
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsMetricsAPI.GetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowMetricsV1WorkflowsWorkflowNameMetricsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | Filter workflows started after this time (ISO 8601) | 
 **endTime** | **time.Time** | Filter workflows started before this time (ISO 8601) | 

### Return type

[**WorkflowMetrics**](WorkflowMetrics.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


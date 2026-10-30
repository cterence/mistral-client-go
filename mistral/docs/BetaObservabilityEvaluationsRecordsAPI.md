# \BetaObservabilityEvaluationsRecordsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet**](BetaObservabilityEvaluationsRecordsAPI.md#GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet) | **Get** /v1/observability/evaluation-runs_v2/{run_id}/records/{record_id} | Get a specific record from an evaluation run
[**SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost**](BetaObservabilityEvaluationsRecordsAPI.md#SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost) | **Post** /v1/observability/evaluation-runs_v2/{run_id}/records/search | Search records for an evaluation run with content filters



## GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet

> EvaluationRunRecordV2Response GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet(ctx, runId, recordId).Execute()

Get a specific record from an evaluation run



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
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsRecordsAPI.GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet(context.Background(), runId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRecordsAPI.GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet`: EvaluationRunRecordV2Response
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRecordsAPI.GetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluationRunRecordV2V1ObservabilityEvaluationRunsV2RunIdRecordsRecordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EvaluationRunRecordV2Response**](EvaluationRunRecordV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost

> ListEvaluationRunRecordsResponse SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost(ctx, runId).PageSize(pageSize).Page(page).EvaluationRunRecordFilterInput(evaluationRunRecordFilterInput).Execute()

Search records for an evaluation run with content filters



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
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	evaluationRunRecordFilterInput := *openapiclient.NewEvaluationRunRecordFilterInput() // EvaluationRunRecordFilterInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsRecordsAPI.SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost(context.Background(), runId).PageSize(pageSize).Page(page).EvaluationRunRecordFilterInput(evaluationRunRecordFilterInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRecordsAPI.SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost`: ListEvaluationRunRecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRecordsAPI.SearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEvaluationRunRecordsV2V1ObservabilityEvaluationRunsV2RunIdRecordsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **evaluationRunRecordFilterInput** | [**EvaluationRunRecordFilterInput**](EvaluationRunRecordFilterInput.md) |  | 

### Return type

[**ListEvaluationRunRecordsResponse**](ListEvaluationRunRecordsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


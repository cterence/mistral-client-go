# \BetaObservabilityEvaluationsRunsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost**](BetaObservabilityEvaluationsRunsAPI.md#ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost) | **Post** /v1/observability/evaluation-runs_v2/{run_id}/statistics/compute | Compute statistics for a specific evaluation run with optional filters
[**GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet**](BetaObservabilityEvaluationsRunsAPI.md#GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet) | **Get** /v1/observability/evaluation-runs_v2/{run_id} | Get a specific evaluation run by ID
[**GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet**](BetaObservabilityEvaluationsRunsAPI.md#GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet) | **Get** /v1/observability/evaluation-runs_v2/{run_id}/statistics | Get statistics for a specific evaluation run
[**SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost**](BetaObservabilityEvaluationsRunsAPI.md#SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost) | **Post** /v1/observability/evaluation-runs_v2/search | Search evaluation runs across all evaluations
[**SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost**](BetaObservabilityEvaluationsRunsAPI.md#SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost) | **Post** /v1/observability/evaluations_v2/{evaluation_slug}/runs/search | Search evaluation runs with filters



## ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost

> RunStatisticsResponse ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost(ctx, runId).EvaluationRunRecordFilterInput(evaluationRunRecordFilterInput).Execute()

Compute statistics for a specific evaluation run with optional filters



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
	evaluationRunRecordFilterInput := *openapiclient.NewEvaluationRunRecordFilterInput() // EvaluationRunRecordFilterInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsRunsAPI.ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost(context.Background(), runId).EvaluationRunRecordFilterInput(evaluationRunRecordFilterInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRunsAPI.ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost`: RunStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRunsAPI.ComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeFilteredRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evaluationRunRecordFilterInput** | [**EvaluationRunRecordFilterInput**](EvaluationRunRecordFilterInput.md) |  | 

### Return type

[**RunStatisticsResponse**](RunStatisticsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet

> EvaluationRunV2Response GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet(ctx, runId).Execute()

Get a specific evaluation run by ID



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
	resp, r, err := apiClient.BetaObservabilityEvaluationsRunsAPI.GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRunsAPI.GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet`: EvaluationRunV2Response
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRunsAPI.GetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluationRunV2V1ObservabilityEvaluationRunsV2RunIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EvaluationRunV2Response**](EvaluationRunV2Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet

> RunStatisticsResponse GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet(ctx, runId).Execute()

Get statistics for a specific evaluation run



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
	resp, r, err := apiClient.BetaObservabilityEvaluationsRunsAPI.GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRunsAPI.GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet`: RunStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRunsAPI.GetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleEvaluationRunStatisticsV2V1ObservabilityEvaluationRunsV2RunIdStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunStatisticsResponse**](RunStatisticsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost

> SearchAllEvaluationRunsResponse SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost(ctx).PageSize(pageSize).Page(page).EvaluationRunFilterInput(evaluationRunFilterInput).Execute()

Search evaluation runs across all evaluations



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
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	evaluationRunFilterInput := *openapiclient.NewEvaluationRunFilterInput() // EvaluationRunFilterInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsRunsAPI.SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost(context.Background()).PageSize(pageSize).Page(page).EvaluationRunFilterInput(evaluationRunFilterInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRunsAPI.SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost`: SearchAllEvaluationRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRunsAPI.SearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAllEvaluationRunsV2V1ObservabilityEvaluationRunsV2SearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **evaluationRunFilterInput** | [**EvaluationRunFilterInput**](EvaluationRunFilterInput.md) |  | 

### Return type

[**SearchAllEvaluationRunsResponse**](SearchAllEvaluationRunsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost

> ListEvaluationRunsResponse SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost(ctx, evaluationSlug).PageSize(pageSize).Page(page).EvaluationRunFilterInput(evaluationRunFilterInput).Execute()

Search evaluation runs with filters



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
	evaluationSlug := "evaluationSlug_example" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	evaluationRunFilterInput := *openapiclient.NewEvaluationRunFilterInput() // EvaluationRunFilterInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsRunsAPI.SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost(context.Background(), evaluationSlug).PageSize(pageSize).Page(page).EvaluationRunFilterInput(evaluationRunFilterInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsRunsAPI.SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost`: ListEvaluationRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsRunsAPI.SearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**evaluationSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEvaluationRunsV2V1ObservabilityEvaluationsV2EvaluationSlugRunsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **evaluationRunFilterInput** | [**EvaluationRunFilterInput**](EvaluationRunFilterInput.md) |  | 

### Return type

[**ListEvaluationRunsResponse**](ListEvaluationRunsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


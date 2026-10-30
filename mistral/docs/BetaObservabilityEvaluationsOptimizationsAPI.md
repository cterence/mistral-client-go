# \BetaObservabilityEvaluationsOptimizationsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet**](BetaObservabilityEvaluationsOptimizationsAPI.md#GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet) | **Get** /v1/observability/optimizations/{optimization_id} | Get an optimization by id (with its anchored evaluation + project)
[**ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet**](BetaObservabilityEvaluationsOptimizationsAPI.md#ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet) | **Get** /v1/observability/optimizations/{optimization_id}/trials | List an optimization&#39;s trials with their observation runs (paginated)
[**ListOptimizationsV1ObservabilityOptimizationsGet**](BetaObservabilityEvaluationsOptimizationsAPI.md#ListOptimizationsV1ObservabilityOptimizationsGet) | **Get** /v1/observability/optimizations | List optimizations with optional filters



## GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet

> OptimizationDetail GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet(ctx, optimizationId).Execute()

Get an optimization by id (with its anchored evaluation + project)



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
	optimizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsOptimizationsAPI.GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet(context.Background(), optimizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsOptimizationsAPI.GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet`: OptimizationDetail
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsOptimizationsAPI.GetOptimizationV1ObservabilityOptimizationsOptimizationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optimizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptimizationV1ObservabilityOptimizationsOptimizationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OptimizationDetail**](OptimizationDetail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet

> ListOptimizationTrialsResponse ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet(ctx, optimizationId).PageSize(pageSize).Page(page).Execute()

List an optimization's trials with their observation runs (paginated)



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
	optimizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsOptimizationsAPI.ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet(context.Background(), optimizationId).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsOptimizationsAPI.ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet`: ListOptimizationTrialsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsOptimizationsAPI.ListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optimizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOptimizationTrialsV1ObservabilityOptimizationsOptimizationIdTrialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**ListOptimizationTrialsResponse**](ListOptimizationTrialsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptimizationsV1ObservabilityOptimizationsGet

> ListOptimizationsResponse ListOptimizationsV1ObservabilityOptimizationsGet(ctx).Project(project).Evaluation(evaluation).Status(status).Algorithm(algorithm).Tags(tags).PageSize(pageSize).Page(page).Execute()

List optimizations with optional filters



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
	project := "project_example" // string | Filter by project slug (optional)
	evaluation := "evaluation_example" // string | Filter by evaluation slug (optional)
	status := "status_example" // string | Filter by lifecycle status (optional)
	algorithm := "algorithm_example" // string | Filter by algorithm (optional)
	tags := []string{"Inner_example"} // []string | Filter by tags (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsOptimizationsAPI.ListOptimizationsV1ObservabilityOptimizationsGet(context.Background()).Project(project).Evaluation(evaluation).Status(status).Algorithm(algorithm).Tags(tags).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsOptimizationsAPI.ListOptimizationsV1ObservabilityOptimizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptimizationsV1ObservabilityOptimizationsGet`: ListOptimizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsOptimizationsAPI.ListOptimizationsV1ObservabilityOptimizationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptimizationsV1ObservabilityOptimizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Filter by project slug | 
 **evaluation** | **string** | Filter by evaluation slug | 
 **status** | **string** | Filter by lifecycle status | 
 **algorithm** | **string** | Filter by algorithm | 
 **tags** | **[]string** | Filter by tags | 
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**ListOptimizationsResponse**](ListOptimizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


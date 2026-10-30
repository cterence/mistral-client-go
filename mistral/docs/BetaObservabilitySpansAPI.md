# \BetaObservabilitySpansAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost**](BetaObservabilitySpansAPI.md#AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost) | **Post** /v1/observability/spans/evaluations/aggregate | Aggregate span evaluations
[**AggregateSpansV1ObservabilitySpansAggregatePost**](BetaObservabilitySpansAPI.md#AggregateSpansV1ObservabilitySpansAggregatePost) | **Post** /v1/observability/spans/aggregate | Aggregate spans
[**GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet**](BetaObservabilitySpansAPI.md#GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet) | **Get** /v1/observability/spans/evaluations/fields/{field_name}/options | Get options for a span evaluation field
[**GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet**](BetaObservabilitySpansAPI.md#GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet) | **Get** /v1/observability/spans/evaluations/fields | Get span evaluation field definitions
[**GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet**](BetaObservabilitySpansAPI.md#GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet) | **Get** /v1/observability/spans/fields/{field_name}/options | Get options for a span field
[**GetSpanFieldsV1ObservabilitySpansFieldsGet**](BetaObservabilitySpansAPI.md#GetSpanFieldsV1ObservabilitySpansFieldsGet) | **Get** /v1/observability/spans/fields | Get span field definitions
[**SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost**](BetaObservabilitySpansAPI.md#SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost) | **Post** /v1/observability/spans/evaluations/search/latest | Search latest span evaluations
[**SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost**](BetaObservabilitySpansAPI.md#SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost) | **Post** /v1/observability/pipeline-results/search | Search pipeline results
[**SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost**](BetaObservabilitySpansAPI.md#SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost) | **Post** /v1/observability/spans/evaluations/search | Search span evaluations
[**SearchSpansV1ObservabilitySpansSearchPost**](BetaObservabilitySpansAPI.md#SearchSpansV1ObservabilitySpansSearchPost) | **Post** /v1/observability/spans/search | Search spans



## AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost

> Aggregation AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost(ctx).AggregationRequest(aggregationRequest).From(from).To(to).Execute()

Aggregate span evaluations



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
	aggregationRequest := *openapiclient.NewAggregationRequest(*openapiclient.NewMetricDefinition("Measure_example", openapiclient.MetricAggregation("count"))) // AggregationRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost(context.Background()).AggregationRequest(aggregationRequest).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.AggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateSpanEvaluationsV1ObservabilitySpansEvaluationsAggregatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregationRequest** | [**AggregationRequest**](AggregationRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**Aggregation**](Aggregation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateSpansV1ObservabilitySpansAggregatePost

> Aggregation AggregateSpansV1ObservabilitySpansAggregatePost(ctx).AggregationRequest(aggregationRequest).From(from).To(to).Execute()

Aggregate spans



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
	aggregationRequest := *openapiclient.NewAggregationRequest(*openapiclient.NewMetricDefinition("Measure_example", openapiclient.MetricAggregation("count"))) // AggregationRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.AggregateSpansV1ObservabilitySpansAggregatePost(context.Background()).AggregationRequest(aggregationRequest).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.AggregateSpansV1ObservabilitySpansAggregatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateSpansV1ObservabilitySpansAggregatePost`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.AggregateSpansV1ObservabilitySpansAggregatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateSpansV1ObservabilitySpansAggregatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregationRequest** | [**AggregationRequest**](AggregationRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**Aggregation**](Aggregation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet

> GetSpanEvaluationFieldOptions GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet(ctx, fieldName).From(from).To(to).Execute()

Get options for a span evaluation field



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
	fieldName := "fieldName_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet(context.Background(), fieldName).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet`: GetSpanEvaluationFieldOptions
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.GetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanEvaluationFieldOptionsV1ObservabilitySpansEvaluationsFieldsFieldNameOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**GetSpanEvaluationFieldOptions**](GetSpanEvaluationFieldOptions.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet

> GetSpanEvaluationFields GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet(ctx).Execute()

Get span evaluation field definitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet`: GetSpanEvaluationFields
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.GetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanEvaluationFieldsV1ObservabilitySpansEvaluationsFieldsGetRequest struct via the builder pattern


### Return type

[**GetSpanEvaluationFields**](GetSpanEvaluationFields.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet

> GetSpanFieldOptions GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet(ctx, fieldName).From(from).To(to).Execute()

Get options for a span field



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
	fieldName := "fieldName_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet(context.Background(), fieldName).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet`: GetSpanFieldOptions
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.GetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanFieldOptionsV1ObservabilitySpansFieldsFieldNameOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**GetSpanFieldOptions**](GetSpanFieldOptions.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanFieldsV1ObservabilitySpansFieldsGet

> GetSpanFields GetSpanFieldsV1ObservabilitySpansFieldsGet(ctx).Execute()

Get span field definitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.GetSpanFieldsV1ObservabilitySpansFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.GetSpanFieldsV1ObservabilitySpansFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpanFieldsV1ObservabilitySpansFieldsGet`: GetSpanFields
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.GetSpanFieldsV1ObservabilitySpansFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanFieldsV1ObservabilitySpansFieldsGetRequest struct via the builder pattern


### Return type

[**GetSpanFields**](GetSpanFields.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost

> GetSpanEvaluations SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost(ctx).SpanEvaluationsRequest(spanEvaluationsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Search latest span evaluations



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
	spanEvaluationsRequest := *openapiclient.NewSpanEvaluationsRequest() // SpanEvaluationsRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost(context.Background()).SpanEvaluationsRequest(spanEvaluationsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost`: GetSpanEvaluations
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.SearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLatestSpanEvaluationsV1ObservabilitySpansEvaluationsSearchLatestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spanEvaluationsRequest** | [**SpanEvaluationsRequest**](SpanEvaluationsRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetSpanEvaluations**](GetSpanEvaluations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost

> GetPipelineResults SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost(ctx).PipelineResultsRequest(pipelineResultsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Search pipeline results



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
	pipelineResultsRequest := *openapiclient.NewPipelineResultsRequest([]string{"PipelineConfigSlugs_example"}) // PipelineResultsRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost(context.Background()).PipelineResultsRequest(pipelineResultsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost`: GetPipelineResults
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.SearchPipelineResultsV1ObservabilityPipelineResultsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPipelineResultsV1ObservabilityPipelineResultsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineResultsRequest** | [**PipelineResultsRequest**](PipelineResultsRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetPipelineResults**](GetPipelineResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost

> GetSpanEvaluations SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost(ctx).SpanEvaluationsRequest(spanEvaluationsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Search span evaluations



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
	spanEvaluationsRequest := *openapiclient.NewSpanEvaluationsRequest() // SpanEvaluationsRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost(context.Background()).SpanEvaluationsRequest(spanEvaluationsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost`: GetSpanEvaluations
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.SearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpanEvaluationsV1ObservabilitySpansEvaluationsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spanEvaluationsRequest** | [**SpanEvaluationsRequest**](SpanEvaluationsRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetSpanEvaluations**](GetSpanEvaluations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpansV1ObservabilitySpansSearchPost

> GetSpans SearchSpansV1ObservabilitySpansSearchPost(ctx).SpansRequest(spansRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Search spans



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
	spansRequest := *openapiclient.NewSpansRequest() // SpansRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilitySpansAPI.SearchSpansV1ObservabilitySpansSearchPost(context.Background()).SpansRequest(spansRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilitySpansAPI.SearchSpansV1ObservabilitySpansSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSpansV1ObservabilitySpansSearchPost`: GetSpans
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilitySpansAPI.SearchSpansV1ObservabilitySpansSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpansV1ObservabilitySpansSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spansRequest** | [**SpansRequest**](SpansRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetSpans**](GetSpans.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


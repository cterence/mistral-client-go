# \BetaObservabilityTracesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateTracesV1ObservabilityTracesAggregatePost**](BetaObservabilityTracesAPI.md#AggregateTracesV1ObservabilityTracesAggregatePost) | **Post** /v1/observability/traces/aggregate | Aggregate traces
[**GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet**](BetaObservabilityTracesAPI.md#GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet) | **Get** /v1/observability/traces/{trace_id}/spans/{span_id} | Get span by id
[**GetTraceByIdV1ObservabilityTracesTraceIdGet**](BetaObservabilityTracesAPI.md#GetTraceByIdV1ObservabilityTracesTraceIdGet) | **Get** /v1/observability/traces/{trace_id} | Get trace by id
[**GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet**](BetaObservabilityTracesAPI.md#GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet) | **Get** /v1/observability/traces/fields/{field_name}/options | Get options for a trace field
[**GetTraceFieldsV1ObservabilityTracesFieldsGet**](BetaObservabilityTracesAPI.md#GetTraceFieldsV1ObservabilityTracesFieldsGet) | **Get** /v1/observability/traces/fields | Get trace field definitions
[**GetTraceSpansV1ObservabilityTracesTraceIdSpansGet**](BetaObservabilityTracesAPI.md#GetTraceSpansV1ObservabilityTracesTraceIdSpansGet) | **Get** /v1/observability/traces/{trace_id}/spans | Get trace spans
[**SearchTracesV1ObservabilityTracesSearchPost**](BetaObservabilityTracesAPI.md#SearchTracesV1ObservabilityTracesSearchPost) | **Post** /v1/observability/traces/search | Search traces



## AggregateTracesV1ObservabilityTracesAggregatePost

> Aggregation AggregateTracesV1ObservabilityTracesAggregatePost(ctx).AggregationRequest(aggregationRequest).From(from).To(to).Execute()

Aggregate traces



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
	resp, r, err := apiClient.BetaObservabilityTracesAPI.AggregateTracesV1ObservabilityTracesAggregatePost(context.Background()).AggregationRequest(aggregationRequest).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.AggregateTracesV1ObservabilityTracesAggregatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateTracesV1ObservabilityTracesAggregatePost`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.AggregateTracesV1ObservabilityTracesAggregatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateTracesV1ObservabilityTracesAggregatePostRequest struct via the builder pattern


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


## GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet

> GetSpan GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet(ctx, traceId, spanId).From(from).To(to).Execute()

Get span by id



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
	traceId := "traceId_example" // string | 
	spanId := "spanId_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityTracesAPI.GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet(context.Background(), traceId, spanId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet`: GetSpan
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.GetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 
**spanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanByIdV1ObservabilityTracesTraceIdSpansSpanIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**GetSpan**](GetSpan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceByIdV1ObservabilityTracesTraceIdGet

> GetTrace GetTraceByIdV1ObservabilityTracesTraceIdGet(ctx, traceId).Execute()

Get trace by id



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
	traceId := "traceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityTracesAPI.GetTraceByIdV1ObservabilityTracesTraceIdGet(context.Background(), traceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.GetTraceByIdV1ObservabilityTracesTraceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceByIdV1ObservabilityTracesTraceIdGet`: GetTrace
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.GetTraceByIdV1ObservabilityTracesTraceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceByIdV1ObservabilityTracesTraceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTrace**](GetTrace.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet

> GetTraceFieldOptions GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet(ctx, fieldName).From(from).To(to).Execute()

Get options for a trace field



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
	resp, r, err := apiClient.BetaObservabilityTracesAPI.GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet(context.Background(), fieldName).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet`: GetTraceFieldOptions
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.GetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFieldOptionsV1ObservabilityTracesFieldsFieldNameOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**GetTraceFieldOptions**](GetTraceFieldOptions.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFieldsV1ObservabilityTracesFieldsGet

> GetTraceFields GetTraceFieldsV1ObservabilityTracesFieldsGet(ctx).Execute()

Get trace field definitions



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
	resp, r, err := apiClient.BetaObservabilityTracesAPI.GetTraceFieldsV1ObservabilityTracesFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.GetTraceFieldsV1ObservabilityTracesFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFieldsV1ObservabilityTracesFieldsGet`: GetTraceFields
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.GetTraceFieldsV1ObservabilityTracesFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFieldsV1ObservabilityTracesFieldsGetRequest struct via the builder pattern


### Return type

[**GetTraceFields**](GetTraceFields.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceSpansV1ObservabilityTracesTraceIdSpansGet

> GetSpans GetTraceSpansV1ObservabilityTracesTraceIdSpansGet(ctx, traceId).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Get trace spans



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
	traceId := "traceId_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityTracesAPI.GetTraceSpansV1ObservabilityTracesTraceIdSpansGet(context.Background(), traceId).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.GetTraceSpansV1ObservabilityTracesTraceIdSpansGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceSpansV1ObservabilityTracesTraceIdSpansGet`: GetSpans
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.GetTraceSpansV1ObservabilityTracesTraceIdSpansGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceSpansV1ObservabilityTracesTraceIdSpansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetSpans**](GetSpans.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTracesV1ObservabilityTracesSearchPost

> GetTraces SearchTracesV1ObservabilityTracesSearchPost(ctx).TracesRequest(tracesRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Search traces



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
	tracesRequest := *openapiclient.NewTracesRequest() // TracesRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityTracesAPI.SearchTracesV1ObservabilityTracesSearchPost(context.Background()).TracesRequest(tracesRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityTracesAPI.SearchTracesV1ObservabilityTracesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTracesV1ObservabilityTracesSearchPost`: GetTraces
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityTracesAPI.SearchTracesV1ObservabilityTracesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTracesV1ObservabilityTracesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tracesRequest** | [**TracesRequest**](TracesRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetTraces**](GetTraces.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


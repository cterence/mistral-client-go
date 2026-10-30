# \BetaObservabilityLogsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet**](BetaObservabilityLogsAPI.md#GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet) | **Get** /v1/observability/logs/fields/{field_name}/options | Get options for a log field
[**GetLogFieldsV1ObservabilityLogsFieldsGet**](BetaObservabilityLogsAPI.md#GetLogFieldsV1ObservabilityLogsFieldsGet) | **Get** /v1/observability/logs/fields | Get log field definitions
[**SearchLogsV1ObservabilityLogsSearchPost**](BetaObservabilityLogsAPI.md#SearchLogsV1ObservabilityLogsSearchPost) | **Post** /v1/observability/logs/search | Search logs



## GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet

> GetLogFieldOptions GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet(ctx, fieldName).From(from).To(to).Execute()

Get options for a log field



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
	resp, r, err := apiClient.BetaObservabilityLogsAPI.GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet(context.Background(), fieldName).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityLogsAPI.GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet`: GetLogFieldOptions
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityLogsAPI.GetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFieldOptionsV1ObservabilityLogsFieldsFieldNameOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

[**GetLogFieldOptions**](GetLogFieldOptions.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogFieldsV1ObservabilityLogsFieldsGet

> GetLogFields GetLogFieldsV1ObservabilityLogsFieldsGet(ctx).Execute()

Get log field definitions



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
	resp, r, err := apiClient.BetaObservabilityLogsAPI.GetLogFieldsV1ObservabilityLogsFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityLogsAPI.GetLogFieldsV1ObservabilityLogsFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogFieldsV1ObservabilityLogsFieldsGet`: GetLogFields
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityLogsAPI.GetLogFieldsV1ObservabilityLogsFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFieldsV1ObservabilityLogsFieldsGetRequest struct via the builder pattern


### Return type

[**GetLogFields**](GetLogFields.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLogsV1ObservabilityLogsSearchPost

> GetLogs SearchLogsV1ObservabilityLogsSearchPost(ctx).LogsRequest(logsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()

Search logs



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
	logsRequest := *openapiclient.NewLogsRequest() // LogsRequest | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityLogsAPI.SearchLogsV1ObservabilityLogsSearchPost(context.Background()).LogsRequest(logsRequest).From(from).To(to).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityLogsAPI.SearchLogsV1ObservabilityLogsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLogsV1ObservabilityLogsSearchPost`: GetLogs
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityLogsAPI.SearchLogsV1ObservabilityLogsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLogsV1ObservabilityLogsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logsRequest** | [**LogsRequest**](LogsRequest.md) |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**GetLogs**](GetLogs.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


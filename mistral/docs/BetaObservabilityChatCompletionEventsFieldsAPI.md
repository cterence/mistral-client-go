# \BetaObservabilityChatCompletionEventsFieldsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost**](BetaObservabilityChatCompletionEventsFieldsAPI.md#GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost) | **Post** /v1/observability/chat-completion-fields/{field_name}/options-counts | Get Chat Completion Field Options Counts
[**GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet**](BetaObservabilityChatCompletionEventsFieldsAPI.md#GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet) | **Get** /v1/observability/chat-completion-fields/{field_name}/options | Get Chat Completion Field Options
[**GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet**](BetaObservabilityChatCompletionEventsFieldsAPI.md#GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet) | **Get** /v1/observability/chat-completion-fields | Get Chat Completion Fields



## GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost

> FieldOptionCounts GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost(ctx, fieldName).FieldOptionCountsInSchema(fieldOptionCountsInSchema).Execute()

Get Chat Completion Field Options Counts

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
	fieldName := "fieldName_example" // string | 
	fieldOptionCountsInSchema := *openapiclient.NewFieldOptionCountsInSchema() // FieldOptionCountsInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost(context.Background(), fieldName).FieldOptionCountsInSchema(fieldOptionCountsInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost`: FieldOptionCounts
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionFieldOptionsCountsV1ObservabilityChatCompletionFieldsFieldNameOptionsCountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldOptionCountsInSchema** | [**FieldOptionCountsInSchema**](FieldOptionCountsInSchema.md) |  | 

### Return type

[**FieldOptionCounts**](FieldOptionCounts.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet

> ChatCompletionFieldOptions GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet(ctx, fieldName).Operator(operator).Execute()

Get Chat Completion Field Options

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
	fieldName := "fieldName_example" // string | 
	operator := "operator_example" // string | The operator to use for filtering options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet(context.Background(), fieldName).Operator(operator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet`: ChatCompletionFieldOptions
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionFieldOptionsV1ObservabilityChatCompletionFieldsFieldNameOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operator** | **string** | The operator to use for filtering options | 

### Return type

[**ChatCompletionFieldOptions**](ChatCompletionFieldOptions.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet

> ChatCompletionFields GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet(ctx).Execute()

Get Chat Completion Fields

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
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet`: ChatCompletionFields
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsFieldsAPI.GetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionFieldsV1ObservabilityChatCompletionFieldsGetRequest struct via the builder pattern


### Return type

[**ChatCompletionFields**](ChatCompletionFields.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


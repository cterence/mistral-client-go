# \BetaObservabilityChatCompletionEventsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost**](BetaObservabilityChatCompletionEventsAPI.md#GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost) | **Post** /v1/observability/chat-completion-events/search-ids | Alternative to /search that returns only the IDs and that can return many IDs at once
[**GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet**](BetaObservabilityChatCompletionEventsAPI.md#GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet) | **Get** /v1/observability/chat-completion-events/{event_id} | Get Chat Completion Event
[**GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost**](BetaObservabilityChatCompletionEventsAPI.md#GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost) | **Post** /v1/observability/chat-completion-events/search | Get Chat Completion Events
[**GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet**](BetaObservabilityChatCompletionEventsAPI.md#GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet) | **Get** /v1/observability/chat-completion-events/{event_id}/similar-events | Get Similar Chat Completion Events
[**JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost**](BetaObservabilityChatCompletionEventsAPI.md#JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost) | **Post** /v1/observability/chat-completion-events/{event_id}/live-judging | Run Judge on an event based on the given options



## GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost

> ChatCompletionEventIds GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost(ctx).GetChatCompletionEventIdsInSchema(getChatCompletionEventIdsInSchema).Execute()

Alternative to /search that returns only the IDs and that can return many IDs at once

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
	getChatCompletionEventIdsInSchema := *openapiclient.NewGetChatCompletionEventIdsInSchema(*openapiclient.NewFilterPayload("TODO")) // GetChatCompletionEventIdsInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost(context.Background()).GetChatCompletionEventIdsInSchema(getChatCompletionEventIdsInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost`: ChatCompletionEventIds
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionEventIdsV1ObservabilityChatCompletionEventsSearchIdsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getChatCompletionEventIdsInSchema** | [**GetChatCompletionEventIdsInSchema**](GetChatCompletionEventIdsInSchema.md) |  | 

### Return type

[**ChatCompletionEventIds**](ChatCompletionEventIds.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet

> ChatCompletionEvent GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet(ctx, eventId).Execute()

Get Chat Completion Event

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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet`: ChatCompletionEvent
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionEventV1ObservabilityChatCompletionEventsEventIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatCompletionEvent**](ChatCompletionEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost

> ChatCompletionEvents GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost(ctx).GetChatCompletionEventsInSchema(getChatCompletionEventsInSchema).PageSize(pageSize).Cursor(cursor).Execute()

Get Chat Completion Events

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
	getChatCompletionEventsInSchema := *openapiclient.NewGetChatCompletionEventsInSchema(*openapiclient.NewFilterPayload("TODO")) // GetChatCompletionEventsInSchema | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost(context.Background()).GetChatCompletionEventsInSchema(getChatCompletionEventsInSchema).PageSize(pageSize).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost`: ChatCompletionEvents
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsAPI.GetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChatCompletionEventsV1ObservabilityChatCompletionEventsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getChatCompletionEventsInSchema** | [**GetChatCompletionEventsInSchema**](GetChatCompletionEventsInSchema.md) |  | 
 **pageSize** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ChatCompletionEvents**](ChatCompletionEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet

> ChatCompletionEvents GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet(ctx, eventId).Execute()

Get Similar Chat Completion Events

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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsAPI.GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsAPI.GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet`: ChatCompletionEvents
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsAPI.GetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarChatCompletionEventsV1ObservabilityChatCompletionEventsEventIdSimilarEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatCompletionEvents**](ChatCompletionEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost

> JudgeOutput JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost(ctx, eventId).PostChatCompletionEventJudgingInSchema(postChatCompletionEventJudgingInSchema).Execute()

Run Judge on an event based on the given options

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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postChatCompletionEventJudgingInSchema := *openapiclient.NewPostChatCompletionEventJudgingInSchema(*openapiclient.NewPostJudgeInSchema("Name_example", "Description_example", "ModelName_example", openapiclient.Output{JudgeClassificationOutput: openapiclient.NewJudgeClassificationOutput([]openapiclient.JudgeClassificationOutputOption{*openapiclient.NewJudgeClassificationOutputOption("Value_example", "Description_example")})}, "Instructions_example", []string{"Tools_example"})) // PostChatCompletionEventJudgingInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityChatCompletionEventsAPI.JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost(context.Background(), eventId).PostChatCompletionEventJudgingInSchema(postChatCompletionEventJudgingInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityChatCompletionEventsAPI.JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost`: JudgeOutput
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityChatCompletionEventsAPI.JudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJudgeChatCompletionEventV1ObservabilityChatCompletionEventsEventIdLiveJudgingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postChatCompletionEventJudgingInSchema** | [**PostChatCompletionEventJudgingInSchema**](PostChatCompletionEventJudgingInSchema.md) |  | 

### Return type

[**JudgeOutput**](JudgeOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


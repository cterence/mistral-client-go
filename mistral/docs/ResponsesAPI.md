# \ResponsesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountInputTokensV1ResponsesInputTokensPost**](ResponsesAPI.md#CountInputTokensV1ResponsesInputTokensPost) | **Post** /v1/responses/input_tokens | Count Input Tokens
[**CreateResponseV1ResponsesPost**](ResponsesAPI.md#CreateResponseV1ResponsesPost) | **Post** /v1/responses | Create Response



## CountInputTokensV1ResponsesInputTokensPost

> CountResponseInputTokensResponse CountInputTokensV1ResponsesInputTokensPost(ctx).CountResponseInputTokensRequest(countResponseInputTokensRequest).Execute()

Count Input Tokens



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
	countResponseInputTokensRequest := *openapiclient.NewCountResponseInputTokensRequest("Model_example", *openapiclient.NewInput6()) // CountResponseInputTokensRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResponsesAPI.CountInputTokensV1ResponsesInputTokensPost(context.Background()).CountResponseInputTokensRequest(countResponseInputTokensRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResponsesAPI.CountInputTokensV1ResponsesInputTokensPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountInputTokensV1ResponsesInputTokensPost`: CountResponseInputTokensResponse
	fmt.Fprintf(os.Stdout, "Response from `ResponsesAPI.CountInputTokensV1ResponsesInputTokensPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountInputTokensV1ResponsesInputTokensPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countResponseInputTokensRequest** | [**CountResponseInputTokensRequest**](CountResponseInputTokensRequest.md) |  | 

### Return type

[**CountResponseInputTokensResponse**](CountResponseInputTokensResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResponseV1ResponsesPost

> CreateResponseResponse CreateResponseV1ResponsesPost(ctx).CreateResponseRequest(createResponseRequest).Execute()

Create Response



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
	createResponseRequest := *openapiclient.NewCreateResponseRequest("Model_example", *openapiclient.NewInput6()) // CreateResponseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResponsesAPI.CreateResponseV1ResponsesPost(context.Background()).CreateResponseRequest(createResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResponsesAPI.CreateResponseV1ResponsesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResponseV1ResponsesPost`: CreateResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResponsesAPI.CreateResponseV1ResponsesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResponseV1ResponsesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResponseRequest** | [**CreateResponseRequest**](CreateResponseRequest.md) |  | 

### Return type

[**CreateResponseResponse**](CreateResponseResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


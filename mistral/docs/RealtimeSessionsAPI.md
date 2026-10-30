# \RealtimeSessionsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientSessionV1ClientSessionsPost**](RealtimeSessionsAPI.md#CreateClientSessionV1ClientSessionsPost) | **Post** /v1/client/sessions | Create Client Session



## CreateClientSessionV1ClientSessionsPost

> CreateRealtimeSessionResponse CreateClientSessionV1ClientSessionsPost(ctx).CreateRealtimeSessionRequest(createRealtimeSessionRequest).Execute()

Create Client Session



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
	createRealtimeSessionRequest := *openapiclient.NewCreateRealtimeSessionRequest("Model_example") // CreateRealtimeSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealtimeSessionsAPI.CreateClientSessionV1ClientSessionsPost(context.Background()).CreateRealtimeSessionRequest(createRealtimeSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealtimeSessionsAPI.CreateClientSessionV1ClientSessionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClientSessionV1ClientSessionsPost`: CreateRealtimeSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `RealtimeSessionsAPI.CreateClientSessionV1ClientSessionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientSessionV1ClientSessionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRealtimeSessionRequest** | [**CreateRealtimeSessionRequest**](CreateRealtimeSessionRequest.md) |  | 

### Return type

[**CreateRealtimeSessionResponse**](CreateRealtimeSessionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


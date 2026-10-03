# \AudioSpeechAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpeechV1AudioSpeechPost**](AudioSpeechAPI.md#SpeechV1AudioSpeechPost) | **Post** /v1/audio/speech | Speech



## SpeechV1AudioSpeechPost

> SpeechResponse SpeechV1AudioSpeechPost(ctx).SpeechRequest(speechRequest).Execute()

Speech



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
	speechRequest := *openapiclient.NewSpeechRequest("Input_example") // SpeechRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioSpeechAPI.SpeechV1AudioSpeechPost(context.Background()).SpeechRequest(speechRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioSpeechAPI.SpeechV1AudioSpeechPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeechV1AudioSpeechPost`: SpeechResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioSpeechAPI.SpeechV1AudioSpeechPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpeechV1AudioSpeechPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speechRequest** | [**SpeechRequest**](SpeechRequest.md) |  | 

### Return type

[**SpeechResponse**](SpeechResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


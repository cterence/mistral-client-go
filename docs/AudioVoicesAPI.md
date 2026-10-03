# \AudioVoicesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoiceV1AudioVoicesPost**](AudioVoicesAPI.md#CreateVoiceV1AudioVoicesPost) | **Post** /v1/audio/voices | Create a new voice
[**DeleteVoiceV1AudioVoicesVoiceIdDelete**](AudioVoicesAPI.md#DeleteVoiceV1AudioVoicesVoiceIdDelete) | **Delete** /v1/audio/voices/{voice_id} | Delete a custom voice
[**GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet**](AudioVoicesAPI.md#GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet) | **Get** /v1/audio/voices/{voice_id}/sample | Get voice sample audio
[**GetVoiceV1AudioVoicesVoiceIdGet**](AudioVoicesAPI.md#GetVoiceV1AudioVoicesVoiceIdGet) | **Get** /v1/audio/voices/{voice_id} | Get voice details
[**ListVoicesV1AudioVoicesGet**](AudioVoicesAPI.md#ListVoicesV1AudioVoicesGet) | **Get** /v1/audio/voices | List all voices
[**UpdateVoiceV1AudioVoicesVoiceIdPatch**](AudioVoicesAPI.md#UpdateVoiceV1AudioVoicesVoiceIdPatch) | **Patch** /v1/audio/voices/{voice_id} | Update voice metadata



## CreateVoiceV1AudioVoicesPost

> VoiceResponse CreateVoiceV1AudioVoicesPost(ctx).VoiceCreateRequest(voiceCreateRequest).Execute()

Create a new voice



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
	voiceCreateRequest := *openapiclient.NewVoiceCreateRequest("Name_example", "SampleAudio_example") // VoiceCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioVoicesAPI.CreateVoiceV1AudioVoicesPost(context.Background()).VoiceCreateRequest(voiceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioVoicesAPI.CreateVoiceV1AudioVoicesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoiceV1AudioVoicesPost`: VoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioVoicesAPI.CreateVoiceV1AudioVoicesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoiceV1AudioVoicesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceCreateRequest** | [**VoiceCreateRequest**](VoiceCreateRequest.md) |  | 

### Return type

[**VoiceResponse**](VoiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVoiceV1AudioVoicesVoiceIdDelete

> VoiceResponse DeleteVoiceV1AudioVoicesVoiceIdDelete(ctx, voiceId).Execute()

Delete a custom voice



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
	voiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioVoicesAPI.DeleteVoiceV1AudioVoicesVoiceIdDelete(context.Background(), voiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioVoicesAPI.DeleteVoiceV1AudioVoicesVoiceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVoiceV1AudioVoicesVoiceIdDelete`: VoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioVoicesAPI.DeleteVoiceV1AudioVoicesVoiceIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoiceV1AudioVoicesVoiceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VoiceResponse**](VoiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet

> string GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet(ctx, voiceId).Execute()

Get voice sample audio



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
	voiceId := "voiceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioVoicesAPI.GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet(context.Background(), voiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioVoicesAPI.GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet`: string
	fmt.Fprintf(os.Stdout, "Response from `AudioVoicesAPI.GetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoiceSampleAudioV1AudioVoicesVoiceIdSampleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoiceV1AudioVoicesVoiceIdGet

> VoiceResponse GetVoiceV1AudioVoicesVoiceIdGet(ctx, voiceId).Execute()

Get voice details



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
	voiceId := "voiceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioVoicesAPI.GetVoiceV1AudioVoicesVoiceIdGet(context.Background(), voiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioVoicesAPI.GetVoiceV1AudioVoicesVoiceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoiceV1AudioVoicesVoiceIdGet`: VoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioVoicesAPI.GetVoiceV1AudioVoicesVoiceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoiceV1AudioVoicesVoiceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VoiceResponse**](VoiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVoicesV1AudioVoicesGet

> VoiceListResponse ListVoicesV1AudioVoicesGet(ctx).Limit(limit).Offset(offset).Execute()

List all voices



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
	limit := int32(56) // int32 | Maximum number of voices to return (optional) (default to 10)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioVoicesAPI.ListVoicesV1AudioVoicesGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioVoicesAPI.ListVoicesV1AudioVoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVoicesV1AudioVoicesGet`: VoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioVoicesAPI.ListVoicesV1AudioVoicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVoicesV1AudioVoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of voices to return | [default to 10]
 **offset** | **int32** | Offset for pagination | [default to 0]

### Return type

[**VoiceListResponse**](VoiceListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVoiceV1AudioVoicesVoiceIdPatch

> VoiceResponse UpdateVoiceV1AudioVoicesVoiceIdPatch(ctx, voiceId).VoiceUpdateRequest(voiceUpdateRequest).Execute()

Update voice metadata



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
	voiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	voiceUpdateRequest := *openapiclient.NewVoiceUpdateRequest() // VoiceUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioVoicesAPI.UpdateVoiceV1AudioVoicesVoiceIdPatch(context.Background(), voiceId).VoiceUpdateRequest(voiceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioVoicesAPI.UpdateVoiceV1AudioVoicesVoiceIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVoiceV1AudioVoicesVoiceIdPatch`: VoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioVoicesAPI.UpdateVoiceV1AudioVoicesVoiceIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoiceV1AudioVoicesVoiceIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voiceUpdateRequest** | [**VoiceUpdateRequest**](VoiceUpdateRequest.md) |  | 

### Return type

[**VoiceResponse**](VoiceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


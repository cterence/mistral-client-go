# \AudioTranscriptionsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AudioApiV1TranscriptionsPost**](AudioTranscriptionsAPI.md#AudioApiV1TranscriptionsPost) | **Post** /v1/audio/transcriptions | Create Transcription
[**AudioApiV1TranscriptionsPostStream**](AudioTranscriptionsAPI.md#AudioApiV1TranscriptionsPostStream) | **Post** /v1/audio/transcriptions#stream | Create Streaming Transcription (SSE)



## AudioApiV1TranscriptionsPost

> TranscriptionResponse AudioApiV1TranscriptionsPost(ctx).Model(model).File(file).FileUrl(fileUrl).FileId(fileId).Language(language).Temperature(temperature).Stream(stream).Diarize(diarize).ContextBias(contextBias).TimestampGranularities(timestampGranularities).Execute()

Create Transcription

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
	model := "model_example" // string | ID of the model to be used.
	file := TODO // *os.File | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  ```bash  file=@path/to/your/file.jsonl;filename=custom_name.jsonl  ```  Otherwise, you can just keep the original file name:  ```bash  file=@path/to/your/file.jsonl  ``` (optional)
	fileUrl := "fileUrl_example" // string | Url of a file to be transcribed (optional)
	fileId := "fileId_example" // string | ID of a file uploaded to /v1/files (optional)
	language := "language_example" // string | Language of the audio, e.g. 'en'. Providing the language can boost accuracy. (optional)
	temperature := float32(8.14) // float32 |  (optional)
	stream := true // bool |  (optional) (default to false)
	diarize := true // bool |  (optional) (default to false)
	contextBias := []string{"Inner_example"} // []string |  (optional)
	timestampGranularities := []openapiclient.TimestampGranularity{openapiclient.TimestampGranularity("segment")} // []TimestampGranularity | Granularities of timestamps to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioTranscriptionsAPI.AudioApiV1TranscriptionsPost(context.Background()).Model(model).File(file).FileUrl(fileUrl).FileId(fileId).Language(language).Temperature(temperature).Stream(stream).Diarize(diarize).ContextBias(contextBias).TimestampGranularities(timestampGranularities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioTranscriptionsAPI.AudioApiV1TranscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AudioApiV1TranscriptionsPost`: TranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioTranscriptionsAPI.AudioApiV1TranscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAudioApiV1TranscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | ID of the model to be used. | 
 **file** | [***os.File**](*os.File.md) | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl;filename&#x3D;custom_name.jsonl  &#x60;&#x60;&#x60;  Otherwise, you can just keep the original file name:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl  &#x60;&#x60;&#x60; | 
 **fileUrl** | **string** | Url of a file to be transcribed | 
 **fileId** | **string** | ID of a file uploaded to /v1/files | 
 **language** | **string** | Language of the audio, e.g. &#39;en&#39;. Providing the language can boost accuracy. | 
 **temperature** | **float32** |  | 
 **stream** | **bool** |  | [default to false]
 **diarize** | **bool** |  | [default to false]
 **contextBias** | **[]string** |  | 
 **timestampGranularities** | [**[]TimestampGranularity**](TimestampGranularity.md) | Granularities of timestamps to include in the response. | 

### Return type

[**TranscriptionResponse**](TranscriptionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AudioApiV1TranscriptionsPostStream

> TranscriptionStreamEvents AudioApiV1TranscriptionsPostStream(ctx).Model(model).File(file).FileUrl(fileUrl).FileId(fileId).Language(language).Temperature(temperature).Stream(stream).Diarize(diarize).ContextBias(contextBias).TimestampGranularities(timestampGranularities).Execute()

Create Streaming Transcription (SSE)

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
	model := "model_example" // string | 
	file := TODO // *os.File | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  ```bash  file=@path/to/your/file.jsonl;filename=custom_name.jsonl  ```  Otherwise, you can just keep the original file name:  ```bash  file=@path/to/your/file.jsonl  ``` (optional)
	fileUrl := "fileUrl_example" // string | Url of a file to be transcribed (optional)
	fileId := "fileId_example" // string | ID of a file uploaded to /v1/files (optional)
	language := "language_example" // string | Language of the audio, e.g. 'en'. Providing the language can boost accuracy. (optional)
	temperature := float32(8.14) // float32 |  (optional)
	stream := true // bool |  (optional) (default to true)
	diarize := true // bool |  (optional) (default to false)
	contextBias := []string{"Inner_example"} // []string |  (optional)
	timestampGranularities := []openapiclient.TimestampGranularity{openapiclient.TimestampGranularity("segment")} // []TimestampGranularity | Granularities of timestamps to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioTranscriptionsAPI.AudioApiV1TranscriptionsPostStream(context.Background()).Model(model).File(file).FileUrl(fileUrl).FileId(fileId).Language(language).Temperature(temperature).Stream(stream).Diarize(diarize).ContextBias(contextBias).TimestampGranularities(timestampGranularities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioTranscriptionsAPI.AudioApiV1TranscriptionsPostStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AudioApiV1TranscriptionsPostStream`: TranscriptionStreamEvents
	fmt.Fprintf(os.Stdout, "Response from `AudioTranscriptionsAPI.AudioApiV1TranscriptionsPostStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAudioApiV1TranscriptionsPostStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** |  | 
 **file** | [***os.File**](*os.File.md) | The File object (not file name) to be uploaded.  To upload a file and specify a custom file name you should format your request as such:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl;filename&#x3D;custom_name.jsonl  &#x60;&#x60;&#x60;  Otherwise, you can just keep the original file name:  &#x60;&#x60;&#x60;bash  file&#x3D;@path/to/your/file.jsonl  &#x60;&#x60;&#x60; | 
 **fileUrl** | **string** | Url of a file to be transcribed | 
 **fileId** | **string** | ID of a file uploaded to /v1/files | 
 **language** | **string** | Language of the audio, e.g. &#39;en&#39;. Providing the language can boost accuracy. | 
 **temperature** | **float32** |  | 
 **stream** | **bool** |  | [default to true]
 **diarize** | **bool** |  | [default to false]
 **contextBias** | **[]string** |  | 
 **timestampGranularities** | [**[]TimestampGranularity**](TimestampGranularity.md) | Granularities of timestamps to include in the response. | 

### Return type

[**TranscriptionStreamEvents**](TranscriptionStreamEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


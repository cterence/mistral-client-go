# \BetaObservabilityJudgesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJudgeV1ObservabilityJudgesPost**](BetaObservabilityJudgesAPI.md#CreateJudgeV1ObservabilityJudgesPost) | **Post** /v1/observability/judges | Create a new judge
[**DeleteJudgeV1ObservabilityJudgesJudgeIdDelete**](BetaObservabilityJudgesAPI.md#DeleteJudgeV1ObservabilityJudgesJudgeIdDelete) | **Delete** /v1/observability/judges/{judge_id} | Delete a judge
[**GetJudgeByIdV1ObservabilityJudgesJudgeIdGet**](BetaObservabilityJudgesAPI.md#GetJudgeByIdV1ObservabilityJudgesJudgeIdGet) | **Get** /v1/observability/judges/{judge_id} | Get judge by id
[**GetJudgesV1ObservabilityJudgesGet**](BetaObservabilityJudgesAPI.md#GetJudgesV1ObservabilityJudgesGet) | **Get** /v1/observability/judges | Get judges with optional filtering and search
[**JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost**](BetaObservabilityJudgesAPI.md#JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost) | **Post** /v1/observability/judges/{judge_id}/live-judging | Run a saved judge on a conversation
[**UpdateJudgeV1ObservabilityJudgesJudgeIdPut**](BetaObservabilityJudgesAPI.md#UpdateJudgeV1ObservabilityJudgesJudgeIdPut) | **Put** /v1/observability/judges/{judge_id} | Update a judge



## CreateJudgeV1ObservabilityJudgesPost

> Judge CreateJudgeV1ObservabilityJudgesPost(ctx).CreateJudgeRequest(createJudgeRequest).Execute()

Create a new judge



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
	createJudgeRequest := *openapiclient.NewCreateJudgeRequest("Name_example", "Description_example", "ModelName_example", openapiclient.Output{JudgeClassificationOutput: openapiclient.NewJudgeClassificationOutput([]openapiclient.JudgeClassificationOutputOption{*openapiclient.NewJudgeClassificationOutputOption("Value_example", "Description_example")})}, "Instructions_example", []string{"Tools_example"}) // CreateJudgeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityJudgesAPI.CreateJudgeV1ObservabilityJudgesPost(context.Background()).CreateJudgeRequest(createJudgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityJudgesAPI.CreateJudgeV1ObservabilityJudgesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJudgeV1ObservabilityJudgesPost`: Judge
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityJudgesAPI.CreateJudgeV1ObservabilityJudgesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJudgeV1ObservabilityJudgesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJudgeRequest** | [**CreateJudgeRequest**](CreateJudgeRequest.md) |  | 

### Return type

[**Judge**](Judge.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJudgeV1ObservabilityJudgesJudgeIdDelete

> DeleteJudgeV1ObservabilityJudgesJudgeIdDelete(ctx, judgeId).Execute()

Delete a judge



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
	judgeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityJudgesAPI.DeleteJudgeV1ObservabilityJudgesJudgeIdDelete(context.Background(), judgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityJudgesAPI.DeleteJudgeV1ObservabilityJudgesJudgeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**judgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJudgeV1ObservabilityJudgesJudgeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJudgeByIdV1ObservabilityJudgesJudgeIdGet

> Judge GetJudgeByIdV1ObservabilityJudgesJudgeIdGet(ctx, judgeId).Execute()

Get judge by id



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
	judgeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityJudgesAPI.GetJudgeByIdV1ObservabilityJudgesJudgeIdGet(context.Background(), judgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityJudgesAPI.GetJudgeByIdV1ObservabilityJudgesJudgeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJudgeByIdV1ObservabilityJudgesJudgeIdGet`: Judge
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityJudgesAPI.GetJudgeByIdV1ObservabilityJudgesJudgeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**judgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJudgeByIdV1ObservabilityJudgesJudgeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Judge**](Judge.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJudgesV1ObservabilityJudgesGet

> ListJudgesResponse GetJudgesV1ObservabilityJudgesGet(ctx).TypeFilter(typeFilter).ModelFilter(modelFilter).PageSize(pageSize).Page(page).Q(q).Execute()

Get judges with optional filtering and search



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
	typeFilter := []openapiclient.JudgeOutputType{openapiclient.JudgeOutputType("REGRESSION")} // []JudgeOutputType | Filter by judge output types (optional)
	modelFilter := []string{"Inner_example"} // []string | Filter by model names (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	q := "q_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityJudgesAPI.GetJudgesV1ObservabilityJudgesGet(context.Background()).TypeFilter(typeFilter).ModelFilter(modelFilter).PageSize(pageSize).Page(page).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityJudgesAPI.GetJudgesV1ObservabilityJudgesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJudgesV1ObservabilityJudgesGet`: ListJudgesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityJudgesAPI.GetJudgesV1ObservabilityJudgesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJudgesV1ObservabilityJudgesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeFilter** | [**[]JudgeOutputType**](JudgeOutputType.md) | Filter by judge output types | 
 **modelFilter** | **[]string** | Filter by model names | 
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **q** | **string** |  | 

### Return type

[**ListJudgesResponse**](ListJudgesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost

> JudgeOutput JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost(ctx, judgeId).JudgeConversationRequest(judgeConversationRequest).Execute()

Run a saved judge on a conversation



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
	judgeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	judgeConversationRequest := *openapiclient.NewJudgeConversationRequest([]*map[string]interface{}{nil}) // JudgeConversationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityJudgesAPI.JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost(context.Background(), judgeId).JudgeConversationRequest(judgeConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityJudgesAPI.JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost`: JudgeOutput
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityJudgesAPI.JudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**judgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJudgeConversationV1ObservabilityJudgesJudgeIdLiveJudgingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **judgeConversationRequest** | [**JudgeConversationRequest**](JudgeConversationRequest.md) |  | 

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


## UpdateJudgeV1ObservabilityJudgesJudgeIdPut

> UpdateJudgeV1ObservabilityJudgesJudgeIdPut(ctx, judgeId).UpdateJudgeRequest(updateJudgeRequest).Execute()

Update a judge



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
	judgeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateJudgeRequest := *openapiclient.NewUpdateJudgeRequest("Name_example", "Description_example", "ModelName_example", openapiclient.Output{JudgeClassificationOutput: openapiclient.NewJudgeClassificationOutput([]openapiclient.JudgeClassificationOutputOption{*openapiclient.NewJudgeClassificationOutputOption("Value_example", "Description_example")})}, "Instructions_example", []string{"Tools_example"}) // UpdateJudgeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityJudgesAPI.UpdateJudgeV1ObservabilityJudgesJudgeIdPut(context.Background(), judgeId).UpdateJudgeRequest(updateJudgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityJudgesAPI.UpdateJudgeV1ObservabilityJudgesJudgeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**judgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJudgeV1ObservabilityJudgesJudgeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateJudgeRequest** | [**UpdateJudgeRequest**](UpdateJudgeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


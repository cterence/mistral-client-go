# \BetaObservabilityEvaluationsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePipelineV1ObservabilityPipelinesPost**](BetaObservabilityEvaluationsAPI.md#CreatePipelineV1ObservabilityPipelinesPost) | **Post** /v1/observability/pipelines | Create a pipeline
[**DeletePipelineV1ObservabilityPipelinesPipelineIdDelete**](BetaObservabilityEvaluationsAPI.md#DeletePipelineV1ObservabilityPipelinesPipelineIdDelete) | **Delete** /v1/observability/pipelines/{pipeline_id} | Delete a pipeline
[**GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet**](BetaObservabilityEvaluationsAPI.md#GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet) | **Get** /v1/observability/evaluations_v2/{evaluation_slug} | Get evaluation by slug
[**GetPipelineV1ObservabilityPipelinesPipelineIdGet**](BetaObservabilityEvaluationsAPI.md#GetPipelineV1ObservabilityPipelinesPipelineIdGet) | **Get** /v1/observability/pipelines/{pipeline_id} | Get a pipeline
[**GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get**](BetaObservabilityEvaluationsAPI.md#GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get) | **Get** /v1/observability/projects/{project_slug}/evaluations_v2 | List evaluations for a project
[**ListPipelinesV1ObservabilityPipelinesGet**](BetaObservabilityEvaluationsAPI.md#ListPipelinesV1ObservabilityPipelinesGet) | **Get** /v1/observability/pipelines | List pipelines
[**UpdatePipelineV1ObservabilityPipelinesPipelineIdPut**](BetaObservabilityEvaluationsAPI.md#UpdatePipelineV1ObservabilityPipelinesPipelineIdPut) | **Put** /v1/observability/pipelines/{pipeline_id} | Update a pipeline



## CreatePipelineV1ObservabilityPipelinesPost

> Pipeline CreatePipelineV1ObservabilityPipelinesPost(ctx).CreatePipelineRequest(createPipelineRequest).Execute()

Create a pipeline



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
	createPipelineRequest := *openapiclient.NewCreatePipelineRequest("Name_example", []openapiclient.PipelineConfigSelector{*openapiclient.NewPipelineConfigSelector(openapiclient.SourceKind("span"))}, []openapiclient.LlmJudgeReference{*openapiclient.NewLlmJudgeReference("Slug_example")}) // CreatePipelineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsAPI.CreatePipelineV1ObservabilityPipelinesPost(context.Background()).CreatePipelineRequest(createPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.CreatePipelineV1ObservabilityPipelinesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePipelineV1ObservabilityPipelinesPost`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsAPI.CreatePipelineV1ObservabilityPipelinesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineV1ObservabilityPipelinesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPipelineRequest** | [**CreatePipelineRequest**](CreatePipelineRequest.md) |  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineV1ObservabilityPipelinesPipelineIdDelete

> DeletePipelineV1ObservabilityPipelinesPipelineIdDelete(ctx, pipelineId).Execute()

Delete a pipeline



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
	pipelineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityEvaluationsAPI.DeletePipelineV1ObservabilityPipelinesPipelineIdDelete(context.Background(), pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.DeletePipelineV1ObservabilityPipelinesPipelineIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineV1ObservabilityPipelinesPipelineIdDeleteRequest struct via the builder pattern


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


## GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet

> EvaluationV2 GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet(ctx, evaluationSlug).Execute()

Get evaluation by slug



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
	evaluationSlug := "evaluationSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsAPI.GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet(context.Background(), evaluationSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet`: EvaluationV2
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsAPI.GetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**evaluationSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluationV2V1ObservabilityEvaluationsV2EvaluationSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EvaluationV2**](EvaluationV2.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineV1ObservabilityPipelinesPipelineIdGet

> Pipeline GetPipelineV1ObservabilityPipelinesPipelineIdGet(ctx, pipelineId).Execute()

Get a pipeline



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
	pipelineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsAPI.GetPipelineV1ObservabilityPipelinesPipelineIdGet(context.Background(), pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.GetPipelineV1ObservabilityPipelinesPipelineIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineV1ObservabilityPipelinesPipelineIdGet`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsAPI.GetPipelineV1ObservabilityPipelinesPipelineIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineV1ObservabilityPipelinesPipelineIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get

> EvaluationsV2 GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get(ctx, projectSlug).PageSize(pageSize).Page(page).Execute()

List evaluations for a project



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
	projectSlug := "projectSlug_example" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsAPI.GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get(context.Background(), projectSlug).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get`: EvaluationsV2
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsAPI.GetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEvaluationsV2V1ObservabilityProjectsProjectSlugEvaluationsV2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**EvaluationsV2**](EvaluationsV2.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelinesV1ObservabilityPipelinesGet

> PipelinesResponse ListPipelinesV1ObservabilityPipelinesGet(ctx).Enabled(enabled).PageSize(pageSize).Page(page).Q(q).Execute()

List pipelines



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
	enabled := true // bool |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	q := "q_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsAPI.ListPipelinesV1ObservabilityPipelinesGet(context.Background()).Enabled(enabled).PageSize(pageSize).Page(page).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.ListPipelinesV1ObservabilityPipelinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPipelinesV1ObservabilityPipelinesGet`: PipelinesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsAPI.ListPipelinesV1ObservabilityPipelinesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPipelinesV1ObservabilityPipelinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **q** | **string** |  | 

### Return type

[**PipelinesResponse**](PipelinesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineV1ObservabilityPipelinesPipelineIdPut

> Pipeline UpdatePipelineV1ObservabilityPipelinesPipelineIdPut(ctx, pipelineId).UpdatePipelineRequest(updatePipelineRequest).Execute()

Update a pipeline



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
	pipelineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updatePipelineRequest := *openapiclient.NewUpdatePipelineRequest() // UpdatePipelineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsAPI.UpdatePipelineV1ObservabilityPipelinesPipelineIdPut(context.Background(), pipelineId).UpdatePipelineRequest(updatePipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsAPI.UpdatePipelineV1ObservabilityPipelinesPipelineIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePipelineV1ObservabilityPipelinesPipelineIdPut`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsAPI.UpdatePipelineV1ObservabilityPipelinesPipelineIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineV1ObservabilityPipelinesPipelineIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePipelineRequest** | [**UpdatePipelineRequest**](UpdatePipelineRequest.md) |  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \BetaObservabilityEvaluationsProjectsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectV1ObservabilityProjectsProjectSlugGet**](BetaObservabilityEvaluationsProjectsAPI.md#GetProjectV1ObservabilityProjectsProjectSlugGet) | **Get** /v1/observability/projects/{project_slug} | Get project by slug
[**GetProjectsV1ObservabilityProjectsGet**](BetaObservabilityEvaluationsProjectsAPI.md#GetProjectsV1ObservabilityProjectsGet) | **Get** /v1/observability/projects | List projects



## GetProjectV1ObservabilityProjectsProjectSlugGet

> Project GetProjectV1ObservabilityProjectsProjectSlugGet(ctx, projectSlug).Execute()

Get project by slug



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsProjectsAPI.GetProjectV1ObservabilityProjectsProjectSlugGet(context.Background(), projectSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsProjectsAPI.GetProjectV1ObservabilityProjectsProjectSlugGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectV1ObservabilityProjectsProjectSlugGet`: Project
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsProjectsAPI.GetProjectV1ObservabilityProjectsProjectSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectV1ObservabilityProjectsProjectSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsV1ObservabilityProjectsGet

> Projects GetProjectsV1ObservabilityProjectsGet(ctx).Q(q).PageSize(pageSize).Page(page).Execute()

List projects



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
	q := "q_example" // string | Case-insensitive filter on project name/slug. (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityEvaluationsProjectsAPI.GetProjectsV1ObservabilityProjectsGet(context.Background()).Q(q).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityEvaluationsProjectsAPI.GetProjectsV1ObservabilityProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsV1ObservabilityProjectsGet`: Projects
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityEvaluationsProjectsAPI.GetProjectsV1ObservabilityProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsV1ObservabilityProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Case-insensitive filter on project name/slug. | 
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**Projects**](Projects.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


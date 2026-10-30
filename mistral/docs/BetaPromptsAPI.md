# \BetaPromptsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PromptsCreate**](BetaPromptsAPI.md#PromptsCreate) | **Post** /v2/prompts | CreatePrompt
[**PromptsCreateVersion**](BetaPromptsAPI.md#PromptsCreateVersion) | **Post** /v2/prompts/{prompt_id}/versions | CreatePromptVersion
[**PromptsDelete**](BetaPromptsAPI.md#PromptsDelete) | **Delete** /v2/prompts/{prompt_id} | DeletePrompt
[**PromptsGet**](BetaPromptsAPI.md#PromptsGet) | **Get** /v2/prompts/{prompt_id} | GetPrompt
[**PromptsGetVersion**](BetaPromptsAPI.md#PromptsGetVersion) | **Get** /v2/prompts/{prompt_id}/versions/{version} | GetPromptVersion
[**PromptsList**](BetaPromptsAPI.md#PromptsList) | **Get** /v2/prompts | ListPrompts
[**PromptsListVersions**](BetaPromptsAPI.md#PromptsListVersions) | **Get** /v2/prompts/{prompt_id}/versions | ListPromptVersions
[**PromptsUpdate**](BetaPromptsAPI.md#PromptsUpdate) | **Patch** /v2/prompts/{prompt_id} | UpdatePrompt
[**PromptsUpdateVersionMetadata**](BetaPromptsAPI.md#PromptsUpdateVersionMetadata) | **Patch** /v2/prompts/{prompt_id}/versions/{version} | UpdatePromptVersionMetadata



## PromptsCreate

> Prompt PromptsCreate(ctx).CreatePromptRequest(createPromptRequest).Execute()

CreatePrompt



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
	createPromptRequest := *openapiclient.NewCreatePromptRequest("Name_example", *openapiclient.NewPromptDefinition("Content_example")) // CreatePromptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsCreate(context.Background()).CreatePromptRequest(createPromptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsCreate`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromptsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPromptRequest** | [**CreatePromptRequest**](CreatePromptRequest.md) |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsCreateVersion

> CreatePromptVersionResponse PromptsCreateVersion(ctx, promptId).CreatePromptVersionRequest(createPromptVersionRequest).Execute()

CreatePromptVersion



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
	promptId := "promptId_example" // string | 
	createPromptVersionRequest := *openapiclient.NewCreatePromptVersionRequest(*openapiclient.NewPromptDefinition("Content_example")) // CreatePromptVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsCreateVersion(context.Background(), promptId).CreatePromptVersionRequest(createPromptVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsCreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsCreateVersion`: CreatePromptVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsCreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPromptVersionRequest** | [**CreatePromptVersionRequest**](CreatePromptVersionRequest.md) |  | 

### Return type

[**CreatePromptVersionResponse**](CreatePromptVersionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsDelete

> map[string]interface{} PromptsDelete(ctx, promptId).Execute()

DeletePrompt



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
	promptId := "promptId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsDelete(context.Background(), promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsGet

> Prompt PromptsGet(ctx, promptId).Version(version).Alias(alias).Fields(fields).Execute()

GetPrompt



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
	promptId := "promptId_example" // string | 
	version := int32(1) // int32 |  (optional)
	alias := "alias_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsGet(context.Background(), promptId).Version(version).Alias(alias).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsGet`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 
 **alias** | **string** |  | 
 **fields** | **[]string** |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsGetVersion

> Prompt PromptsGetVersion(ctx, promptId, version).Fields(fields).Execute()

GetPromptVersion



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
	promptId := "promptId_example" // string | 
	version := int32(1) // int32 | 
	fields := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsGetVersion(context.Background(), promptId, version).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsGetVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsGetVersion`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsGetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsList

> ListPromptsResponse PromptsList(ctx).PageSize(pageSize).PageToken(pageToken).Alias(alias).Fields(fields).SortField(sortField).SortDirection(sortDirection).SortBy(sortBy).SortDirection2(sortDirection2).Execute()

ListPrompts



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
	pageSize := int32(56) // int32 |  (optional)
	pageToken := "pageToken_example" // string |  (optional)
	alias := "alias_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string |  (optional)
	sortField := openapiclient.ListSortField("list_sort_field_unspecified") // ListSortField | Defaults to created_at when omitted. (optional)
	sortDirection := openapiclient.ListSortDirection("list_sort_direction_unspecified") // ListSortDirection | Defaults to descending for timestamp fields and ascending for text fields. (optional)
	sortBy := "sortBy_example" // string | REST-friendly alias for sort.field. Supported values: created_at, last_modified_at, name, title. (optional)
	sortDirection2 := "sortDirection_example" // string | REST-friendly alias for sort.direction. Supported values: asc, desc. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsList(context.Background()).PageSize(pageSize).PageToken(pageToken).Alias(alias).Fields(fields).SortField(sortField).SortDirection(sortDirection).SortBy(sortBy).SortDirection2(sortDirection2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsList`: ListPromptsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromptsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | 
 **pageToken** | **string** |  | 
 **alias** | **string** |  | 
 **fields** | **[]string** |  | 
 **sortField** | [**ListSortField**](ListSortField.md) | Defaults to created_at when omitted. | 
 **sortDirection** | [**ListSortDirection**](ListSortDirection.md) | Defaults to descending for timestamp fields and ascending for text fields. | 
 **sortBy** | **string** | REST-friendly alias for sort.field. Supported values: created_at, last_modified_at, name, title. | 
 **sortDirection2** | **string** | REST-friendly alias for sort.direction. Supported values: asc, desc. | 

### Return type

[**ListPromptsResponse**](ListPromptsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsListVersions

> ListPromptVersionsResponse PromptsListVersions(ctx, promptId).Execute()

ListPromptVersions



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
	promptId := "promptId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsListVersions(context.Background(), promptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsListVersions`: ListPromptVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPromptVersionsResponse**](ListPromptVersionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsUpdate

> Prompt PromptsUpdate(ctx, promptId).UpdatePromptRequest(updatePromptRequest).Execute()

UpdatePrompt



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
	promptId := "promptId_example" // string | 
	updatePromptRequest := *openapiclient.NewUpdatePromptRequest() // UpdatePromptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsUpdate(context.Background(), promptId).UpdatePromptRequest(updatePromptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsUpdate`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePromptRequest** | [**UpdatePromptRequest**](UpdatePromptRequest.md) |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsUpdateVersionMetadata

> Prompt PromptsUpdateVersionMetadata(ctx, promptId, version).UpdatePromptVersionRequest(updatePromptVersionRequest).Execute()

UpdatePromptVersionMetadata



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
	promptId := "promptId_example" // string | 
	version := int32(1) // int32 | 
	updatePromptVersionRequest := *openapiclient.NewUpdatePromptVersionRequest() // UpdatePromptVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaPromptsAPI.PromptsUpdateVersionMetadata(context.Background(), promptId, version).UpdatePromptVersionRequest(updatePromptVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaPromptsAPI.PromptsUpdateVersionMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsUpdateVersionMetadata`: Prompt
	fmt.Fprintf(os.Stdout, "Response from `BetaPromptsAPI.PromptsUpdateVersionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsUpdateVersionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePromptVersionRequest** | [**UpdatePromptVersionRequest**](UpdatePromptVersionRequest.md) |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


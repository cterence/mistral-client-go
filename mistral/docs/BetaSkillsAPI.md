# \BetaSkillsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SkillsCreate**](BetaSkillsAPI.md#SkillsCreate) | **Post** /v2/skills | CreateSkill
[**SkillsCreateVersion**](BetaSkillsAPI.md#SkillsCreateVersion) | **Post** /v2/skills/{skill_id}/versions | CreateSkillVersion
[**SkillsDelete**](BetaSkillsAPI.md#SkillsDelete) | **Delete** /v2/skills/{skill_id} | DeleteSkill
[**SkillsGet**](BetaSkillsAPI.md#SkillsGet) | **Get** /v2/skills/{skill_id} | GetSkill
[**SkillsGetVersion**](BetaSkillsAPI.md#SkillsGetVersion) | **Get** /v2/skills/{skill_id}/versions/{version} | GetSkillVersion
[**SkillsList**](BetaSkillsAPI.md#SkillsList) | **Get** /v2/skills | ListSkills
[**SkillsListVersions**](BetaSkillsAPI.md#SkillsListVersions) | **Get** /v2/skills/{skill_id}/versions | ListSkillVersions
[**SkillsUpdate**](BetaSkillsAPI.md#SkillsUpdate) | **Patch** /v2/skills/{skill_id} | UpdateSkill
[**SkillsUpdateVersionMetadata**](BetaSkillsAPI.md#SkillsUpdateVersionMetadata) | **Patch** /v2/skills/{skill_id}/versions/{version} | UpdateSkillVersionMetadata



## SkillsCreate

> Skill SkillsCreate(ctx).CreateSkillRequest(createSkillRequest).Execute()

CreateSkill



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
	createSkillRequest := *openapiclient.NewCreateSkillRequest("Name_example", *openapiclient.NewSkillDefinition()) // CreateSkillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsCreate(context.Background()).CreateSkillRequest(createSkillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsCreate`: Skill
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSkillsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSkillRequest** | [**CreateSkillRequest**](CreateSkillRequest.md) |  | 

### Return type

[**Skill**](Skill.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsCreateVersion

> CreateSkillVersionResponse SkillsCreateVersion(ctx, skillId).CreateSkillVersionRequest(createSkillVersionRequest).Execute()

CreateSkillVersion



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
	skillId := "skillId_example" // string | 
	createSkillVersionRequest := *openapiclient.NewCreateSkillVersionRequest(*openapiclient.NewSkillDefinition()) // CreateSkillVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsCreateVersion(context.Background(), skillId).CreateSkillVersionRequest(createSkillVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsCreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsCreateVersion`: CreateSkillVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsCreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSkillVersionRequest** | [**CreateSkillVersionRequest**](CreateSkillVersionRequest.md) |  | 

### Return type

[**CreateSkillVersionResponse**](CreateSkillVersionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsDelete

> map[string]interface{} SkillsDelete(ctx, skillId).Execute()

DeleteSkill



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
	skillId := "skillId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsDelete(context.Background(), skillId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsDeleteRequest struct via the builder pattern


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


## SkillsGet

> Skill SkillsGet(ctx, skillId).Version(version).Alias(alias).Fields(fields).Execute()

GetSkill



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
	skillId := "skillId_example" // string | 
	version := int32(1) // int32 |  (optional)
	alias := "alias_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsGet(context.Background(), skillId).Version(version).Alias(alias).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsGet`: Skill
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 
 **alias** | **string** |  | 
 **fields** | **[]string** |  | 

### Return type

[**Skill**](Skill.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsGetVersion

> Skill SkillsGetVersion(ctx, skillId, version).Fields(fields).Execute()

GetSkillVersion



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
	skillId := "skillId_example" // string | 
	version := int32(1) // int32 | 
	fields := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsGetVersion(context.Background(), skillId, version).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsGetVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsGetVersion`: Skill
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsGetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** |  | 

### Return type

[**Skill**](Skill.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsList

> ListSkillsResponse SkillsList(ctx).PageSize(pageSize).PageToken(pageToken).Alias(alias).Fields(fields).SortField(sortField).SortDirection(sortDirection).SortBy(sortBy).SortDirection2(sortDirection2).Execute()

ListSkills



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
	resp, r, err := apiClient.BetaSkillsAPI.SkillsList(context.Background()).PageSize(pageSize).PageToken(pageToken).Alias(alias).Fields(fields).SortField(sortField).SortDirection(sortDirection).SortBy(sortBy).SortDirection2(sortDirection2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsList`: ListSkillsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSkillsListRequest struct via the builder pattern


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

[**ListSkillsResponse**](ListSkillsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsListVersions

> ListSkillVersionsResponse SkillsListVersions(ctx, skillId).Execute()

ListSkillVersions



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
	skillId := "skillId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsListVersions(context.Background(), skillId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsListVersions`: ListSkillVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSkillVersionsResponse**](ListSkillVersionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsUpdate

> Skill SkillsUpdate(ctx, skillId).UpdateSkillRequest(updateSkillRequest).Execute()

UpdateSkill



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
	skillId := "skillId_example" // string | 
	updateSkillRequest := *openapiclient.NewUpdateSkillRequest() // UpdateSkillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsUpdate(context.Background(), skillId).UpdateSkillRequest(updateSkillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsUpdate`: Skill
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSkillRequest** | [**UpdateSkillRequest**](UpdateSkillRequest.md) |  | 

### Return type

[**Skill**](Skill.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsUpdateVersionMetadata

> Skill SkillsUpdateVersionMetadata(ctx, skillId, version).UpdateSkillVersionRequest(updateSkillVersionRequest).Execute()

UpdateSkillVersionMetadata



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
	skillId := "skillId_example" // string | 
	version := int32(1) // int32 | 
	updateSkillVersionRequest := *openapiclient.NewUpdateSkillVersionRequest() // UpdateSkillVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaSkillsAPI.SkillsUpdateVersionMetadata(context.Background(), skillId, version).UpdateSkillVersionRequest(updateSkillVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaSkillsAPI.SkillsUpdateVersionMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkillsUpdateVersionMetadata`: Skill
	fmt.Fprintf(os.Stdout, "Response from `BetaSkillsAPI.SkillsUpdateVersionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkillsUpdateVersionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSkillVersionRequest** | [**UpdateSkillVersionRequest**](UpdateSkillVersionRequest.md) |  | 

### Return type

[**Skill**](Skill.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


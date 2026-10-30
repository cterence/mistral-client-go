# \BetaLibrariesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibrariesCreateV1**](BetaLibrariesAPI.md#LibrariesCreateV1) | **Post** /v1/libraries | Create a new Library.
[**LibrariesDeleteV1**](BetaLibrariesAPI.md#LibrariesDeleteV1) | **Delete** /v1/libraries/{library_id} | Delete a library and all of it&#39;s document.
[**LibrariesGetV1**](BetaLibrariesAPI.md#LibrariesGetV1) | **Get** /v1/libraries/{library_id} | Detailed information about a specific Library.
[**LibrariesListV1**](BetaLibrariesAPI.md#LibrariesListV1) | **Get** /v1/libraries | List all libraries you have access to.
[**LibrariesPatchV1**](BetaLibrariesAPI.md#LibrariesPatchV1) | **Patch** /v1/libraries/{library_id} | Update a library.
[**LibrariesUpdateV1**](BetaLibrariesAPI.md#LibrariesUpdateV1) | **Put** /v1/libraries/{library_id} | Update a library.



## LibrariesCreateV1

> Library LibrariesCreateV1(ctx).CreateLibraryRequest(createLibraryRequest).Execute()

Create a new Library.



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
	createLibraryRequest := *openapiclient.NewCreateLibraryRequest("Name_example") // CreateLibraryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesCreateV1(context.Background()).CreateLibraryRequest(createLibraryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesCreateV1`: Library
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLibraryRequest** | [**CreateLibraryRequest**](CreateLibraryRequest.md) |  | 

### Return type

[**Library**](Library.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesDeleteV1

> Library LibrariesDeleteV1(ctx, libraryId).Execute()

Delete a library and all of it's document.



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesDeleteV1(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesDeleteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesDeleteV1`: Library
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesDeleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesDeleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Library**](Library.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesGetV1

> Library LibrariesGetV1(ctx, libraryId).Execute()

Detailed information about a specific Library.



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesGetV1(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesGetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesGetV1`: Library
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesGetV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesGetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Library**](Library.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesListV1

> ListLibrariesResponse LibrariesListV1(ctx).PageSize(pageSize).PageToken(pageToken).Page(page).Search(search).FilterOwnedByMe(filterOwnedByMe).Execute()

List all libraries you have access to.



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
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	pageToken := "pageToken_example" // string | Continuation token from a previous response's next_page_token. Preferred over `page`. (optional)
	page := int32(56) // int32 | Deprecated: use page_token. Offset paging re-scans earlier pages and is being phased out. (optional) (default to 0)
	search := "search_example" // string | Case-insensitive search on the library name. (optional)
	filterOwnedByMe := true // bool | Deprecated: this parameter will be removed in a future version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesListV1(context.Background()).PageSize(pageSize).PageToken(pageToken).Page(page).Search(search).FilterOwnedByMe(filterOwnedByMe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesListV1`: ListLibrariesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | [default to 100]
 **pageToken** | **string** | Continuation token from a previous response&#39;s next_page_token. Preferred over &#x60;page&#x60;. | 
 **page** | **int32** | Deprecated: use page_token. Offset paging re-scans earlier pages and is being phased out. | [default to 0]
 **search** | **string** | Case-insensitive search on the library name. | 
 **filterOwnedByMe** | **bool** | Deprecated: this parameter will be removed in a future version. | 

### Return type

[**ListLibrariesResponse**](ListLibrariesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesPatchV1

> Library LibrariesPatchV1(ctx, libraryId).UpdateLibraryRequest(updateLibraryRequest).Execute()

Update a library.



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateLibraryRequest := *openapiclient.NewUpdateLibraryRequest() // UpdateLibraryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesPatchV1(context.Background(), libraryId).UpdateLibraryRequest(updateLibraryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesPatchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesPatchV1`: Library
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesPatchV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesPatchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLibraryRequest** | [**UpdateLibraryRequest**](UpdateLibraryRequest.md) |  | 

### Return type

[**Library**](Library.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesUpdateV1

> Library LibrariesUpdateV1(ctx, libraryId).UpdateLibraryRequest(updateLibraryRequest).Execute()

Update a library.



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateLibraryRequest := *openapiclient.NewUpdateLibraryRequest() // UpdateLibraryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAPI.LibrariesUpdateV1(context.Background(), libraryId).UpdateLibraryRequest(updateLibraryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAPI.LibrariesUpdateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesUpdateV1`: Library
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAPI.LibrariesUpdateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesUpdateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLibraryRequest** | [**UpdateLibraryRequest**](UpdateLibraryRequest.md) |  | 

### Return type

[**Library**](Library.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


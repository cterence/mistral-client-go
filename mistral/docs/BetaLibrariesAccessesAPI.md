# \BetaLibrariesAccessesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibrariesShareCreateV1**](BetaLibrariesAccessesAPI.md#LibrariesShareCreateV1) | **Put** /v1/libraries/{library_id}/share | Create or update an access level.
[**LibrariesShareDeleteV1**](BetaLibrariesAccessesAPI.md#LibrariesShareDeleteV1) | **Delete** /v1/libraries/{library_id}/share | Delete an access level.
[**LibrariesShareListV1**](BetaLibrariesAccessesAPI.md#LibrariesShareListV1) | **Get** /v1/libraries/{library_id}/share | List all of the access to this library.



## LibrariesShareCreateV1

> Sharing LibrariesShareCreateV1(ctx, libraryId).SharingRequest(sharingRequest).Execute()

Create or update an access level.



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
	sharingRequest := *openapiclient.NewSharingRequest(openapiclient.ShareEnum("Viewer"), "ShareWithUuid_example", openapiclient.EntityType("User")) // SharingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAccessesAPI.LibrariesShareCreateV1(context.Background(), libraryId).SharingRequest(sharingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAccessesAPI.LibrariesShareCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesShareCreateV1`: Sharing
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAccessesAPI.LibrariesShareCreateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesShareCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharingRequest** | [**SharingRequest**](SharingRequest.md) |  | 

### Return type

[**Sharing**](Sharing.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesShareDeleteV1

> Sharing LibrariesShareDeleteV1(ctx, libraryId).SharingDelete(sharingDelete).Execute()

Delete an access level.



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
	sharingDelete := *openapiclient.NewSharingDelete("ShareWithUuid_example", openapiclient.EntityType("User")) // SharingDelete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaLibrariesAccessesAPI.LibrariesShareDeleteV1(context.Background(), libraryId).SharingDelete(sharingDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAccessesAPI.LibrariesShareDeleteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesShareDeleteV1`: Sharing
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAccessesAPI.LibrariesShareDeleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesShareDeleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharingDelete** | [**SharingDelete**](SharingDelete.md) |  | 

### Return type

[**Sharing**](Sharing.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibrariesShareListV1

> ListSharingResponse LibrariesShareListV1(ctx, libraryId).Execute()

List all of the access to this library.



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
	resp, r, err := apiClient.BetaLibrariesAccessesAPI.LibrariesShareListV1(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaLibrariesAccessesAPI.LibrariesShareListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibrariesShareListV1`: ListSharingResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaLibrariesAccessesAPI.LibrariesShareListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibrariesShareListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSharingResponse**](ListSharingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


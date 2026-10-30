# \BetaAdminApiKeysAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiAdminApiKeysCreateApiKey**](BetaAdminApiKeysAPI.md#UsersApiAdminApiKeysCreateApiKey) | **Post** /v1/admin/api-keys | Create Api Key
[**UsersApiAdminApiKeysDeleteApiKey**](BetaAdminApiKeysAPI.md#UsersApiAdminApiKeysDeleteApiKey) | **Delete** /v1/admin/api-keys/{key_id} | Delete Api Key
[**UsersApiAdminApiKeysGetApiKeys**](BetaAdminApiKeysAPI.md#UsersApiAdminApiKeysGetApiKeys) | **Get** /v1/admin/api-keys | Get Api Keys



## UsersApiAdminApiKeysCreateApiKey

> APIKeyOUT UsersApiAdminApiKeysCreateApiKey(ctx).AdminCreateAPIKeyIN(adminCreateAPIKeyIN).Execute()

Create Api Key



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
	adminCreateAPIKeyIN := *openapiclient.NewAdminCreateAPIKeyIN("WorkspaceUuid_example", "UserId_example") // AdminCreateAPIKeyIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminApiKeysAPI.UsersApiAdminApiKeysCreateApiKey(context.Background()).AdminCreateAPIKeyIN(adminCreateAPIKeyIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminApiKeysAPI.UsersApiAdminApiKeysCreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminApiKeysCreateApiKey`: APIKeyOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminApiKeysAPI.UsersApiAdminApiKeysCreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminApiKeysCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminCreateAPIKeyIN** | [**AdminCreateAPIKeyIN**](AdminCreateAPIKeyIN.md) |  | 

### Return type

[**APIKeyOUT**](APIKeyOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminApiKeysDeleteApiKey

> DeleteAPIKeyOUT UsersApiAdminApiKeysDeleteApiKey(ctx, keyId).Execute()

Delete Api Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | API key ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminApiKeysAPI.UsersApiAdminApiKeysDeleteApiKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminApiKeysAPI.UsersApiAdminApiKeysDeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminApiKeysDeleteApiKey`: DeleteAPIKeyOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminApiKeysAPI.UsersApiAdminApiKeysDeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | API key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminApiKeysDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAPIKeyOUT**](DeleteAPIKeyOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminApiKeysGetApiKeys

> APIKeysExtendedOUT UsersApiAdminApiKeysGetApiKeys(ctx).Limit(limit).Offset(offset).Name(name).Execute()

Get Api Keys



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
	limit := int32(56) // int32 | Maximum number of results to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of results to skip before returning results. (optional) (default to 0)
	name := "name_example" // string | Filter API keys by name substring or the last 4 characters of the key. Matching is case-insensitive and accent-insensitive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminApiKeysAPI.UsersApiAdminApiKeysGetApiKeys(context.Background()).Limit(limit).Offset(offset).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminApiKeysAPI.UsersApiAdminApiKeysGetApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminApiKeysGetApiKeys`: APIKeysExtendedOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminApiKeysAPI.UsersApiAdminApiKeysGetApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminApiKeysGetApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of results to return. | [default to 100]
 **offset** | **int32** | Number of results to skip before returning results. | [default to 0]
 **name** | **string** | Filter API keys by name substring or the last 4 characters of the key. Matching is case-insensitive and accent-insensitive. | 

### Return type

[**APIKeysExtendedOUT**](APIKeysExtendedOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


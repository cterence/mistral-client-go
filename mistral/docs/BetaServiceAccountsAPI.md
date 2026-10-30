# \BetaServiceAccountsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccountV1ServiceAccountsPost**](BetaServiceAccountsAPI.md#CreateServiceAccountV1ServiceAccountsPost) | **Post** /v1/service-accounts | Create Service Account
[**DeleteServiceAccountV1ServiceAccountsServiceAccountIdDelete**](BetaServiceAccountsAPI.md#DeleteServiceAccountV1ServiceAccountsServiceAccountIdDelete) | **Delete** /v1/service-accounts/{service_account_id} | Delete Service Account
[**GetServiceAccountV1ServiceAccountsServiceAccountIdGet**](BetaServiceAccountsAPI.md#GetServiceAccountV1ServiceAccountsServiceAccountIdGet) | **Get** /v1/service-accounts/{service_account_id} | Get Service Account
[**ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet**](BetaServiceAccountsAPI.md#ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet) | **Get** /v1/service-accounts/assignable-roles | List Assignable Service Account Roles
[**ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet**](BetaServiceAccountsAPI.md#ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet) | **Get** /v1/service-accounts/{service_account_id}/roles | List Service Account Roles
[**ListServiceAccountsV1ServiceAccountsGet**](BetaServiceAccountsAPI.md#ListServiceAccountsV1ServiceAccountsGet) | **Get** /v1/service-accounts | List Service Accounts
[**SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut**](BetaServiceAccountsAPI.md#SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut) | **Put** /v1/service-accounts/{service_account_id}/roles | Set Service Account Roles
[**UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch**](BetaServiceAccountsAPI.md#UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch) | **Patch** /v1/service-accounts/{service_account_id} | Update Service Account



## CreateServiceAccountV1ServiceAccountsPost

> ServiceAccount CreateServiceAccountV1ServiceAccountsPost(ctx).CreateServiceAccountRequest(createServiceAccountRequest).Execute()

Create Service Account



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
	createServiceAccountRequest := *openapiclient.NewCreateServiceAccountRequest("Name_example", "WorkspaceId_example") // CreateServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.CreateServiceAccountV1ServiceAccountsPost(context.Background()).CreateServiceAccountRequest(createServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.CreateServiceAccountV1ServiceAccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccountV1ServiceAccountsPost`: ServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.CreateServiceAccountV1ServiceAccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountV1ServiceAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceAccountRequest** | [**CreateServiceAccountRequest**](CreateServiceAccountRequest.md) |  | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccountV1ServiceAccountsServiceAccountIdDelete

> DeleteServiceAccountV1ServiceAccountsServiceAccountIdDelete(ctx, serviceAccountId).Execute()

Delete Service Account



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
	serviceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaServiceAccountsAPI.DeleteServiceAccountV1ServiceAccountsServiceAccountIdDelete(context.Background(), serviceAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.DeleteServiceAccountV1ServiceAccountsServiceAccountIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountV1ServiceAccountsServiceAccountIdDeleteRequest struct via the builder pattern


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


## GetServiceAccountV1ServiceAccountsServiceAccountIdGet

> ServiceAccount GetServiceAccountV1ServiceAccountsServiceAccountIdGet(ctx, serviceAccountId).Execute()

Get Service Account



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
	serviceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.GetServiceAccountV1ServiceAccountsServiceAccountIdGet(context.Background(), serviceAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.GetServiceAccountV1ServiceAccountsServiceAccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountV1ServiceAccountsServiceAccountIdGet`: ServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.GetServiceAccountV1ServiceAccountsServiceAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountV1ServiceAccountsServiceAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet

> ListAssignableServiceAccountRolesResponse ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet(ctx).WorkspaceId(workspaceId).Execute()

List Assignable Service Account Roles



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet`: ListAssignableServiceAccountRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.ListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssignableServiceAccountRolesV1ServiceAccountsAssignableRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 

### Return type

[**ListAssignableServiceAccountRolesResponse**](ListAssignableServiceAccountRolesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet

> ListServiceAccountRolesResponse ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet(ctx, serviceAccountId).Execute()

List Service Account Roles



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
	serviceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet(context.Background(), serviceAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet`: ListServiceAccountRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.ListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListServiceAccountRolesResponse**](ListServiceAccountRolesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccountsV1ServiceAccountsGet

> ListServiceAccountsResponse ListServiceAccountsV1ServiceAccountsGet(ctx).Offset(offset).Limit(limit).WorkspaceId(workspaceId).IncludeDeleted(includeDeleted).Q(q).Order(order).Execute()

List Service Accounts



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
	offset := int32(56) // int32 | 
	limit := int32(56) // int32 | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	includeDeleted := true // bool |  (optional) (default to false)
	q := "q_example" // string |  (optional)
	order := openapiclient.SortOrder("asc") // SortOrder |  (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.ListServiceAccountsV1ServiceAccountsGet(context.Background()).Offset(offset).Limit(limit).WorkspaceId(workspaceId).IncludeDeleted(includeDeleted).Q(q).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.ListServiceAccountsV1ServiceAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccountsV1ServiceAccountsGet`: ListServiceAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.ListServiceAccountsV1ServiceAccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsV1ServiceAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **workspaceId** | **string** |  | 
 **includeDeleted** | **bool** |  | [default to false]
 **q** | **string** |  | 
 **order** | [**SortOrder**](SortOrder.md) |  | [default to &quot;asc&quot;]

### Return type

[**ListServiceAccountsResponse**](ListServiceAccountsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut

> ListServiceAccountRolesResponse SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut(ctx, serviceAccountId).SetServiceAccountRolesRequest(setServiceAccountRolesRequest).Execute()

Set Service Account Roles



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
	serviceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	setServiceAccountRolesRequest := *openapiclient.NewSetServiceAccountRolesRequest([]string{"RoleIds_example"}) // SetServiceAccountRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut(context.Background(), serviceAccountId).SetServiceAccountRolesRequest(setServiceAccountRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut`: ListServiceAccountRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.SetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServiceAccountRolesV1ServiceAccountsServiceAccountIdRolesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setServiceAccountRolesRequest** | [**SetServiceAccountRolesRequest**](SetServiceAccountRolesRequest.md) |  | 

### Return type

[**ListServiceAccountRolesResponse**](ListServiceAccountRolesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch

> ServiceAccount UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch(ctx, serviceAccountId).UpdateServiceAccountRequest(updateServiceAccountRequest).Execute()

Update Service Account



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
	serviceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateServiceAccountRequest := *openapiclient.NewUpdateServiceAccountRequest() // UpdateServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaServiceAccountsAPI.UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch(context.Background(), serviceAccountId).UpdateServiceAccountRequest(updateServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaServiceAccountsAPI.UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch`: ServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `BetaServiceAccountsAPI.UpdateServiceAccountV1ServiceAccountsServiceAccountIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountV1ServiceAccountsServiceAccountIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServiceAccountRequest** | [**UpdateServiceAccountRequest**](UpdateServiceAccountRequest.md) |  | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


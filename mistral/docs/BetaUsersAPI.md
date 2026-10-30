# \BetaUsersAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiGetIdentity**](BetaUsersAPI.md#UsersApiGetIdentity) | **Get** /v1/users/me | Get Identity
[**UsersApiListOrganizations**](BetaUsersAPI.md#UsersApiListOrganizations) | **Get** /v1/users/me/organizations | List Organizations
[**UsersApiListWorkspaces**](BetaUsersAPI.md#UsersApiListWorkspaces) | **Get** /v1/users/me/workspaces | List Workspaces
[**UsersApiWhoami**](BetaUsersAPI.md#UsersApiWhoami) | **Get** /v1/whoami | Whoami



## UsersApiGetIdentity

> UserIdentity UsersApiGetIdentity(ctx).Execute()

Get Identity



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaUsersAPI.UsersApiGetIdentity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaUsersAPI.UsersApiGetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiGetIdentity`: UserIdentity
	fmt.Fprintf(os.Stdout, "Response from `BetaUsersAPI.UsersApiGetIdentity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiGetIdentityRequest struct via the builder pattern


### Return type

[**UserIdentity**](UserIdentity.md)

### Authorization

[DashboardUserContextAuth](../README.md#DashboardUserContextAuth), [BearerUserContextAuth](../README.md#BearerUserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiListOrganizations

> ListOrganizationsResponse UsersApiListOrganizations(ctx).Offset(offset).Limit(limit).Execute()

List Organizations



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
	offset := int32(56) // int32 | Number of organizations to skip before returning results. (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of organizations to return. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaUsersAPI.UsersApiListOrganizations(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaUsersAPI.UsersApiListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiListOrganizations`: ListOrganizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaUsersAPI.UsersApiListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of organizations to skip before returning results. | [default to 0]
 **limit** | **int32** | Maximum number of organizations to return. | [default to 100]

### Return type

[**ListOrganizationsResponse**](ListOrganizationsResponse.md)

### Authorization

[DashboardUserContextAuth](../README.md#DashboardUserContextAuth), [BearerUserContextAuth](../README.md#BearerUserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiListWorkspaces

> ListWorkspacesResponse UsersApiListWorkspaces(ctx).OrganizationId(organizationId).Offset(offset).Limit(limit).Execute()

List Workspaces



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Return only workspaces belonging to this organization. (optional)
	offset := int32(56) // int32 | Number of workspaces to skip before returning results. (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of workspaces to return. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaUsersAPI.UsersApiListWorkspaces(context.Background()).OrganizationId(organizationId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaUsersAPI.UsersApiListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiListWorkspaces`: ListWorkspacesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaUsersAPI.UsersApiListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | Return only workspaces belonging to this organization. | 
 **offset** | **int32** | Number of workspaces to skip before returning results. | [default to 0]
 **limit** | **int32** | Maximum number of workspaces to return. | [default to 100]

### Return type

[**ListWorkspacesResponse**](ListWorkspacesResponse.md)

### Authorization

[DashboardUserContextAuth](../README.md#DashboardUserContextAuth), [BearerUserContextAuth](../README.md#BearerUserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiWhoami

> WhoamiResponse UsersApiWhoami(ctx).Execute()

Whoami



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaUsersAPI.UsersApiWhoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaUsersAPI.UsersApiWhoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiWhoami`: WhoamiResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaUsersAPI.UsersApiWhoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiWhoamiRequest struct via the builder pattern


### Return type

[**WhoamiResponse**](WhoamiResponse.md)

### Authorization

[DashboardUserOrServiceAccountContextAuth](../README.md#DashboardUserOrServiceAccountContextAuth), [BearerUserContextAuth](../README.md#BearerUserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


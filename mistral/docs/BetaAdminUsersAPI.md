# \BetaAdminUsersAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiAdminRolesGetRoles**](BetaAdminUsersAPI.md#UsersApiAdminRolesGetRoles) | **Get** /v1/admin/roles | Get Roles
[**UsersApiAdminUsersCreateUsers**](BetaAdminUsersAPI.md#UsersApiAdminUsersCreateUsers) | **Post** /v1/admin/users | Create Users
[**UsersApiAdminUsersDeactivateUser**](BetaAdminUsersAPI.md#UsersApiAdminUsersDeactivateUser) | **Post** /v1/admin/users/{user_id}/deactivate | Deactivate User
[**UsersApiAdminUsersDeleteInvite**](BetaAdminUsersAPI.md#UsersApiAdminUsersDeleteInvite) | **Delete** /v1/admin/users-invite/{invite_uuid} | Delete Invite
[**UsersApiAdminUsersDeleteUser**](BetaAdminUsersAPI.md#UsersApiAdminUsersDeleteUser) | **Delete** /v1/admin/users/{user_id} | Delete User
[**UsersApiAdminUsersGetInvite**](BetaAdminUsersAPI.md#UsersApiAdminUsersGetInvite) | **Get** /v1/admin/users-invite | Get Invite
[**UsersApiAdminUsersGetUser**](BetaAdminUsersAPI.md#UsersApiAdminUsersGetUser) | **Get** /v1/admin/users/{user_id} | Get User
[**UsersApiAdminUsersGetUsers**](BetaAdminUsersAPI.md#UsersApiAdminUsersGetUsers) | **Get** /v1/admin/users | Get Users
[**UsersApiAdminUsersInviteUsers**](BetaAdminUsersAPI.md#UsersApiAdminUsersInviteUsers) | **Post** /v1/admin/users-invite | Invite Users
[**UsersApiAdminUsersUpdateUser**](BetaAdminUsersAPI.md#UsersApiAdminUsersUpdateUser) | **Patch** /v1/admin/users/{user_id} | Update User



## UsersApiAdminRolesGetRoles

> RolesOut UsersApiAdminRolesGetRoles(ctx).Execute()

Get Roles



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
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminRolesGetRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminRolesGetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminRolesGetRoles`: RolesOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminRolesGetRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminRolesGetRolesRequest struct via the builder pattern


### Return type

[**RolesOut**](RolesOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersCreateUsers

> OrganizationUsersCreateOUT UsersApiAdminUsersCreateUsers(ctx).OrganizationMemberCreate(organizationMemberCreate).Execute()

Create Users



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
	organizationMemberCreate := []openapiclient.OrganizationMemberCreate{*openapiclient.NewOrganizationMemberCreate("Email_example", "FirstName_example", "LastName_example")} // []OrganizationMemberCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersCreateUsers(context.Background()).OrganizationMemberCreate(organizationMemberCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersCreateUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersCreateUsers`: OrganizationUsersCreateOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersCreateUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersCreateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationMemberCreate** | [**[]OrganizationMemberCreate**](OrganizationMemberCreate.md) |  | 

### Return type

[**OrganizationUsersCreateOUT**](OrganizationUsersCreateOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersDeactivateUser

> DeactivateUserOUT UsersApiAdminUsersDeactivateUser(ctx, userId).Execute()

Deactivate User



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersDeactivateUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersDeactivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersDeactivateUser`: DeactivateUserOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersDeactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersDeactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeactivateUserOUT**](DeactivateUserOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersDeleteInvite

> DeleteOUT UsersApiAdminUsersDeleteInvite(ctx, inviteUuid).Execute()

Delete Invite



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
	inviteUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization invitation ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersDeleteInvite(context.Background(), inviteUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersDeleteInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersDeleteInvite`: DeleteOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersDeleteInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUuid** | **string** | Organization invitation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersDeleteInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOUT**](DeleteOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersDeleteUser

> DeleteOUT UsersApiAdminUsersDeleteUser(ctx, userId).Execute()

Delete User



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersDeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersDeleteUser`: DeleteOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOUT**](DeleteOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersGetInvite

> []OrganizationUserInviteOUT UsersApiAdminUsersGetInvite(ctx).Execute()

Get Invite



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
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersGetInvite(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersGetInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersGetInvite`: []OrganizationUserInviteOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersGetInvite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersGetInviteRequest struct via the builder pattern


### Return type

[**[]OrganizationUserInviteOUT**](OrganizationUserInviteOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersGetUser

> AdminUserOUT UsersApiAdminUsersGetUser(ctx, userId).Execute()

Get User



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersGetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersGetUser`: AdminUserOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminUserOUT**](AdminUserOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersGetUsers

> OrganizationAdminUsersOUT UsersApiAdminUsersGetUsers(ctx).Page(page).PageSize(pageSize).Email(email).Execute()

Get Users



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
	page := int32(56) // int32 | Page number to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | Maximum number of results per page. (optional) (default to 25)
	email := "email_example" // string | Email address to filter users and invitations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersGetUsers(context.Background()).Page(page).PageSize(pageSize).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersGetUsers`: OrganizationAdminUsersOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to return. | [default to 1]
 **pageSize** | **int32** | Maximum number of results per page. | [default to 25]
 **email** | **string** | Email address to filter users and invitations. | 

### Return type

[**OrganizationAdminUsersOUT**](OrganizationAdminUsersOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersInviteUsers

> OrganizationInvitesCreateOUT UsersApiAdminUsersInviteUsers(ctx).OrganizationInviteIN(organizationInviteIN).Execute()

Invite Users



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
	organizationInviteIN := *openapiclient.NewOrganizationInviteIN("Email_example") // OrganizationInviteIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersInviteUsers(context.Background()).OrganizationInviteIN(organizationInviteIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersInviteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersInviteUsers`: OrganizationInvitesCreateOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersInviteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersInviteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationInviteIN** | [**OrganizationInviteIN**](OrganizationInviteIN.md) |  | 

### Return type

[**OrganizationInvitesCreateOUT**](OrganizationInvitesCreateOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsersUpdateUser

> AdminOrganizationMemberOUT UsersApiAdminUsersUpdateUser(ctx, userId).AdminOrganizationMemberUpdate(adminOrganizationMemberUpdate).Execute()

Update User



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User ID.
	adminOrganizationMemberUpdate := *openapiclient.NewAdminOrganizationMemberUpdate() // AdminOrganizationMemberUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUsersAPI.UsersApiAdminUsersUpdateUser(context.Background(), userId).AdminOrganizationMemberUpdate(adminOrganizationMemberUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUsersAPI.UsersApiAdminUsersUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsersUpdateUser`: AdminOrganizationMemberOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUsersAPI.UsersApiAdminUsersUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsersUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminOrganizationMemberUpdate** | [**AdminOrganizationMemberUpdate**](AdminOrganizationMemberUpdate.md) |  | 

### Return type

[**AdminOrganizationMemberOUT**](AdminOrganizationMemberOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


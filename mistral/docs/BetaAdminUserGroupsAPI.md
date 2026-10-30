# \BetaAdminUserGroupsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersAdminUserGroupsGetGroupWorkspaceAssignments**](BetaAdminUserGroupsAPI.md#UsersAdminUserGroupsGetGroupWorkspaceAssignments) | **Get** /v1/admin/user-groups/{group_uuid}/workspaces | Get Group Workspace Assignments
[**UsersAdminUserGroupsRemoveGroupFromWorkspace**](BetaAdminUserGroupsAPI.md#UsersAdminUserGroupsRemoveGroupFromWorkspace) | **Delete** /v1/admin/user-groups/{group_uuid}/workspaces/{workspace_uuid} | Remove Group From Workspace
[**UsersAdminUserGroupsUpdateGroupWorkspaceAssignment**](BetaAdminUserGroupsAPI.md#UsersAdminUserGroupsUpdateGroupWorkspaceAssignment) | **Patch** /v1/admin/user-groups/{group_uuid}/workspaces/{workspace_uuid} | Update Group Workspace Assignment
[**UsersAdminUserGroupsUpdateUserGroupOrganizationRole**](BetaAdminUserGroupsAPI.md#UsersAdminUserGroupsUpdateUserGroupOrganizationRole) | **Patch** /v1/admin/user-groups/{group_uuid}/organization-role | Update User Group Organization Role
[**UsersApiAdminUserGroupsAssignGroupToWorkspace**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsAssignGroupToWorkspace) | **Post** /v1/admin/user-groups/{group_uuid}/workspaces | Assign Group To Workspace
[**UsersApiAdminUserGroupsAssignUsersToGroup**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsAssignUsersToGroup) | **Post** /v1/admin/user-groups/{group_uuid}/members | Assign Users To Group
[**UsersApiAdminUserGroupsCreateUserGroup**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsCreateUserGroup) | **Post** /v1/admin/user-groups | Create User Group
[**UsersApiAdminUserGroupsDeleteUserGroup**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsDeleteUserGroup) | **Delete** /v1/admin/user-groups/{group_uuid} | Delete User Group
[**UsersApiAdminUserGroupsGetNestedGroups**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsGetNestedGroups) | **Get** /v1/admin/user-groups/{group_uuid}/nested | Get Nested Groups Admin
[**UsersApiAdminUserGroupsGetUserGroup**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsGetUserGroup) | **Get** /v1/admin/user-groups/{group_uuid} | Get User Group
[**UsersApiAdminUserGroupsGetUserGroupMembers**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsGetUserGroupMembers) | **Get** /v1/admin/user-groups/{group_uuid}/members | Get User Group Members
[**UsersApiAdminUserGroupsGetUserGroups**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsGetUserGroups) | **Get** /v1/admin/user-groups | Get User Groups
[**UsersApiAdminUserGroupsProvisionGroupToWorkspace**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsProvisionGroupToWorkspace) | **Post** /v1/admin/user-groups/provision-workspace | Provision Group To Workspace
[**UsersApiAdminUserGroupsRemoveUsersFromGroup**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsRemoveUsersFromGroup) | **Delete** /v1/admin/user-groups/{group_uuid}/members | Remove Users From Group
[**UsersApiAdminUserGroupsSetNestedGroups**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsSetNestedGroups) | **Patch** /v1/admin/user-groups/{group_uuid}/nested | Set Nested Groups Admin
[**UsersApiAdminUserGroupsUpdateUserGroup**](BetaAdminUserGroupsAPI.md#UsersApiAdminUserGroupsUpdateUserGroup) | **Patch** /v1/admin/user-groups/{group_uuid} | Update User Group



## UsersAdminUserGroupsGetGroupWorkspaceAssignments

> GroupWorkspaceAssignmentsOut UsersAdminUserGroupsGetGroupWorkspaceAssignments(ctx, groupUuid).Page(page).PageSize(pageSize).Execute()

Get Group Workspace Assignments



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	page := int32(56) // int32 | Page number to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | Maximum number of results per page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersAdminUserGroupsGetGroupWorkspaceAssignments(context.Background(), groupUuid).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersAdminUserGroupsGetGroupWorkspaceAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersAdminUserGroupsGetGroupWorkspaceAssignments`: GroupWorkspaceAssignmentsOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersAdminUserGroupsGetGroupWorkspaceAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAdminUserGroupsGetGroupWorkspaceAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to return. | [default to 1]
 **pageSize** | **int32** | Maximum number of results per page. | [default to 25]

### Return type

[**GroupWorkspaceAssignmentsOut**](GroupWorkspaceAssignmentsOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersAdminUserGroupsRemoveGroupFromWorkspace

> UsersAdminUserGroupsRemoveGroupFromWorkspace(ctx, groupUuid, workspaceUuid).Execute()

Remove Group From Workspace



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersAdminUserGroupsRemoveGroupFromWorkspace(context.Background(), groupUuid, workspaceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersAdminUserGroupsRemoveGroupFromWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAdminUserGroupsRemoveGroupFromWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersAdminUserGroupsUpdateGroupWorkspaceAssignment

> UsersAdminUserGroupsUpdateGroupWorkspaceAssignment(ctx, groupUuid, workspaceUuid).AdminUpdateGroupWorkspaceAssignmentIn(adminUpdateGroupWorkspaceAssignmentIn).Execute()

Update Group Workspace Assignment



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.
	adminUpdateGroupWorkspaceAssignmentIn := *openapiclient.NewAdminUpdateGroupWorkspaceAssignmentIn() // AdminUpdateGroupWorkspaceAssignmentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersAdminUserGroupsUpdateGroupWorkspaceAssignment(context.Background(), groupUuid, workspaceUuid).AdminUpdateGroupWorkspaceAssignmentIn(adminUpdateGroupWorkspaceAssignmentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersAdminUserGroupsUpdateGroupWorkspaceAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAdminUserGroupsUpdateGroupWorkspaceAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adminUpdateGroupWorkspaceAssignmentIn** | [**AdminUpdateGroupWorkspaceAssignmentIn**](AdminUpdateGroupWorkspaceAssignmentIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersAdminUserGroupsUpdateUserGroupOrganizationRole

> UsersAdminUserGroupsUpdateUserGroupOrganizationRole(ctx, groupUuid).AdminUpdateUserGroupOrganizationRoleIn(adminUpdateUserGroupOrganizationRoleIn).Execute()

Update User Group Organization Role



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	adminUpdateUserGroupOrganizationRoleIn := *openapiclient.NewAdminUpdateUserGroupOrganizationRoleIn("OrganizationRole_example") // AdminUpdateUserGroupOrganizationRoleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersAdminUserGroupsUpdateUserGroupOrganizationRole(context.Background(), groupUuid).AdminUpdateUserGroupOrganizationRoleIn(adminUpdateUserGroupOrganizationRoleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersAdminUserGroupsUpdateUserGroupOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAdminUserGroupsUpdateUserGroupOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminUpdateUserGroupOrganizationRoleIn** | [**AdminUpdateUserGroupOrganizationRoleIn**](AdminUpdateUserGroupOrganizationRoleIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsAssignGroupToWorkspace

> UsersApiAdminUserGroupsAssignGroupToWorkspace(ctx, groupUuid).AdminAssignGroupToWorkspaceIn(adminAssignGroupToWorkspaceIn).Execute()

Assign Group To Workspace



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	adminAssignGroupToWorkspaceIn := *openapiclient.NewAdminAssignGroupToWorkspaceIn("WorkspaceUuid_example") // AdminAssignGroupToWorkspaceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsAssignGroupToWorkspace(context.Background(), groupUuid).AdminAssignGroupToWorkspaceIn(adminAssignGroupToWorkspaceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsAssignGroupToWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsAssignGroupToWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminAssignGroupToWorkspaceIn** | [**AdminAssignGroupToWorkspaceIn**](AdminAssignGroupToWorkspaceIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsAssignUsersToGroup

> UsersApiAdminUserGroupsAssignUsersToGroup(ctx, groupUuid).AdminAssignUsersToGroupIn(adminAssignUsersToGroupIn).Execute()

Assign Users To Group



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	adminAssignUsersToGroupIn := *openapiclient.NewAdminAssignUsersToGroupIn([]string{"UserUuids_example"}) // AdminAssignUsersToGroupIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsAssignUsersToGroup(context.Background(), groupUuid).AdminAssignUsersToGroupIn(adminAssignUsersToGroupIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsAssignUsersToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsAssignUsersToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminAssignUsersToGroupIn** | [**AdminAssignUsersToGroupIn**](AdminAssignUsersToGroupIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsCreateUserGroup

> AdminUserGroupOut UsersApiAdminUserGroupsCreateUserGroup(ctx).AdminUserGroupIn(adminUserGroupIn).Execute()

Create User Group



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
	adminUserGroupIn := *openapiclient.NewAdminUserGroupIn("Name_example") // AdminUserGroupIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsCreateUserGroup(context.Background()).AdminUserGroupIn(adminUserGroupIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsCreateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUserGroupsCreateUserGroup`: AdminUserGroupOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsCreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUserGroupIn** | [**AdminUserGroupIn**](AdminUserGroupIn.md) |  | 

### Return type

[**AdminUserGroupOut**](AdminUserGroupOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsDeleteUserGroup

> UsersApiAdminUserGroupsDeleteUserGroup(ctx, groupUuid).Execute()

Delete User Group



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsDeleteUserGroup(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsDeleteUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsDeleteUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsGetNestedGroups

> NestedGroupsOut UsersApiAdminUserGroupsGetNestedGroups(ctx, groupUuid).Execute()

Get Nested Groups Admin



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetNestedGroups(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetNestedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUserGroupsGetNestedGroups`: NestedGroupsOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetNestedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsGetNestedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NestedGroupsOut**](NestedGroupsOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsGetUserGroup

> AdminUserGroupOut UsersApiAdminUserGroupsGetUserGroup(ctx, groupUuid).Execute()

Get User Group



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroup(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUserGroupsGetUserGroup`: AdminUserGroupOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsGetUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminUserGroupOut**](AdminUserGroupOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsGetUserGroupMembers

> AdminUserGroupMembersOut UsersApiAdminUserGroupsGetUserGroupMembers(ctx, groupUuid).Page(page).PageSize(pageSize).Execute()

Get User Group Members



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	page := int32(56) // int32 | Page number to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | Maximum number of results per page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroupMembers(context.Background(), groupUuid).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUserGroupsGetUserGroupMembers`: AdminUserGroupMembersOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsGetUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to return. | [default to 1]
 **pageSize** | **int32** | Maximum number of results per page. | [default to 25]

### Return type

[**AdminUserGroupMembersOut**](AdminUserGroupMembersOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsGetUserGroups

> AdminUserGroupsOut UsersApiAdminUserGroupsGetUserGroups(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

Get User Groups



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
	search := "search_example" // string | Search term used to filter user groups by name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroups(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUserGroupsGetUserGroups`: AdminUserGroupsOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsGetUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to return. | [default to 1]
 **pageSize** | **int32** | Maximum number of results per page. | [default to 25]
 **search** | **string** | Search term used to filter user groups by name. | 

### Return type

[**AdminUserGroupsOut**](AdminUserGroupsOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsProvisionGroupToWorkspace

> UsersApiAdminUserGroupsProvisionGroupToWorkspace(ctx).AdminProvisionGroupToWorkspaceIn(adminProvisionGroupToWorkspaceIn).Execute()

Provision Group To Workspace



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
	adminProvisionGroupToWorkspaceIn := *openapiclient.NewAdminProvisionGroupToWorkspaceIn("UserGroupUuid_example", "WorkspaceUuid_example") // AdminProvisionGroupToWorkspaceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsProvisionGroupToWorkspace(context.Background()).AdminProvisionGroupToWorkspaceIn(adminProvisionGroupToWorkspaceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsProvisionGroupToWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsProvisionGroupToWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminProvisionGroupToWorkspaceIn** | [**AdminProvisionGroupToWorkspaceIn**](AdminProvisionGroupToWorkspaceIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsRemoveUsersFromGroup

> UsersApiAdminUserGroupsRemoveUsersFromGroup(ctx, groupUuid).AdminAssignUsersToGroupIn(adminAssignUsersToGroupIn).Execute()

Remove Users From Group



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	adminAssignUsersToGroupIn := *openapiclient.NewAdminAssignUsersToGroupIn([]string{"UserUuids_example"}) // AdminAssignUsersToGroupIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsRemoveUsersFromGroup(context.Background(), groupUuid).AdminAssignUsersToGroupIn(adminAssignUsersToGroupIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsRemoveUsersFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsRemoveUsersFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminAssignUsersToGroupIn** | [**AdminAssignUsersToGroupIn**](AdminAssignUsersToGroupIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsSetNestedGroups

> UsersApiAdminUserGroupsSetNestedGroups(ctx, groupUuid).SetNestedGroupsIn(setNestedGroupsIn).Execute()

Set Nested Groups Admin



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User group ID.
	setNestedGroupsIn := *openapiclient.NewSetNestedGroupsIn([]string{"ChildGroupUuids_example"}) // SetNestedGroupsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsSetNestedGroups(context.Background(), groupUuid).SetNestedGroupsIn(setNestedGroupsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsSetNestedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | User group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsSetNestedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNestedGroupsIn** | [**SetNestedGroupsIn**](SetNestedGroupsIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUserGroupsUpdateUserGroup

> AdminUserGroupOut UsersApiAdminUserGroupsUpdateUserGroup(ctx, groupUuid).AdminUpdateUserGroupIn(adminUpdateUserGroupIn).Execute()

Update User Group



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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	adminUpdateUserGroupIn := *openapiclient.NewAdminUpdateUserGroupIn() // AdminUpdateUserGroupIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsUpdateUserGroup(context.Background(), groupUuid).AdminUpdateUserGroupIn(adminUpdateUserGroupIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsUpdateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUserGroupsUpdateUserGroup`: AdminUserGroupOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminUserGroupsAPI.UsersApiAdminUserGroupsUpdateUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUserGroupsUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminUpdateUserGroupIn** | [**AdminUpdateUserGroupIn**](AdminUpdateUserGroupIn.md) |  | 

### Return type

[**AdminUserGroupOut**](AdminUserGroupOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


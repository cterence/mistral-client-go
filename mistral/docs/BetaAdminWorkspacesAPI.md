# \BetaAdminWorkspacesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces) | **Patch** /v1/admin/workspaces/{workspace_uuid}/users | Add Or Update Users Workspaces
[**UsersApiAdminWorkspacesAddUsersWorkspaces**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesAddUsersWorkspaces) | **Post** /v1/admin/workspaces/{workspace_uuid}/add-users | Add Users Workspaces
[**UsersApiAdminWorkspacesCreateWorkspace**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesCreateWorkspace) | **Post** /v1/admin/workspaces | Create Workspace
[**UsersApiAdminWorkspacesDeleteWorkspaces**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesDeleteWorkspaces) | **Delete** /v1/admin/workspaces/{workspace_uuid} | Delete Workspaces
[**UsersApiAdminWorkspacesGetWorkspaces**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesGetWorkspaces) | **Get** /v1/admin/workspaces | Get Workspaces
[**UsersApiAdminWorkspacesRemoveUsersWorkspaces**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesRemoveUsersWorkspaces) | **Delete** /v1/admin/workspaces/{workspace_uuid}/remove-users | Remove Users Workspaces
[**UsersApiAdminWorkspacesUpdateWorkspaces**](BetaAdminWorkspacesAPI.md#UsersApiAdminWorkspacesUpdateWorkspaces) | **Patch** /v1/admin/workspaces/{workspace_uuid} | Update Workspaces



## UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces

> AddOrUpdateUsersToWorkspaceOUT UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces(ctx, workspaceUuid).AdminWorkspaceMemberIN(adminWorkspaceMemberIN).Execute()

Add Or Update Users Workspaces



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
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.
	adminWorkspaceMemberIN := *openapiclient.NewAdminWorkspaceMemberIN() // AdminWorkspaceMemberIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces(context.Background(), workspaceUuid).AdminWorkspaceMemberIN(adminWorkspaceMemberIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces`: AddOrUpdateUsersToWorkspaceOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesAddOrUpdateUsersWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesAddOrUpdateUsersWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminWorkspaceMemberIN** | [**AdminWorkspaceMemberIN**](AdminWorkspaceMemberIN.md) |  | 

### Return type

[**AddOrUpdateUsersToWorkspaceOUT**](AddOrUpdateUsersToWorkspaceOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminWorkspacesAddUsersWorkspaces

> AddUsersToWorkspaceOUT UsersApiAdminWorkspacesAddUsersWorkspaces(ctx, workspaceUuid).AdminWorkspaceMemberIN(adminWorkspaceMemberIN).Execute()

Add Users Workspaces



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
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.
	adminWorkspaceMemberIN := *openapiclient.NewAdminWorkspaceMemberIN() // AdminWorkspaceMemberIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesAddUsersWorkspaces(context.Background(), workspaceUuid).AdminWorkspaceMemberIN(adminWorkspaceMemberIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesAddUsersWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminWorkspacesAddUsersWorkspaces`: AddUsersToWorkspaceOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesAddUsersWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesAddUsersWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminWorkspaceMemberIN** | [**AdminWorkspaceMemberIN**](AdminWorkspaceMemberIN.md) |  | 

### Return type

[**AddUsersToWorkspaceOUT**](AddUsersToWorkspaceOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminWorkspacesCreateWorkspace

> WorkspaceEnrichedOUT UsersApiAdminWorkspacesCreateWorkspace(ctx).AdminWorkspaceIn(adminWorkspaceIn).Execute()

Create Workspace



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
	adminWorkspaceIn := *openapiclient.NewAdminWorkspaceIn("Name_example", "AdminUserId_example") // AdminWorkspaceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesCreateWorkspace(context.Background()).AdminWorkspaceIn(adminWorkspaceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesCreateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminWorkspacesCreateWorkspace`: WorkspaceEnrichedOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesCreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminWorkspaceIn** | [**AdminWorkspaceIn**](AdminWorkspaceIn.md) |  | 

### Return type

[**WorkspaceEnrichedOUT**](WorkspaceEnrichedOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminWorkspacesDeleteWorkspaces

> UsersApiAdminWorkspacesDeleteWorkspaces(ctx, workspaceUuid).Execute()

Delete Workspaces



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
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesDeleteWorkspaces(context.Background(), workspaceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesDeleteWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesDeleteWorkspacesRequest struct via the builder pattern


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


## UsersApiAdminWorkspacesGetWorkspaces

> WorkspacesOut UsersApiAdminWorkspacesGetWorkspaces(ctx).IsArchived(isArchived).Page(page).PageSize(pageSize).Search(search).Execute()

Get Workspaces



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
	isArchived := true // bool | Whether to include archived Workspaces. (optional) (default to false)
	page := int32(56) // int32 | Page number to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | Maximum number of results per page. (optional) (default to 25)
	search := "search_example" // string | Search term to filter Workspaces by name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesGetWorkspaces(context.Background()).IsArchived(isArchived).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesGetWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminWorkspacesGetWorkspaces`: WorkspacesOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesGetWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesGetWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isArchived** | **bool** | Whether to include archived Workspaces. | [default to false]
 **page** | **int32** | Page number to return. | [default to 1]
 **pageSize** | **int32** | Maximum number of results per page. | [default to 25]
 **search** | **string** | Search term to filter Workspaces by name. | 

### Return type

[**WorkspacesOut**](WorkspacesOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminWorkspacesRemoveUsersWorkspaces

> RemoveWorkspaceMembersOUT UsersApiAdminWorkspacesRemoveUsersWorkspaces(ctx, workspaceUuid).RemoveWorkspaceMembersIN(removeWorkspaceMembersIN).Execute()

Remove Users Workspaces



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
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.
	removeWorkspaceMembersIN := *openapiclient.NewRemoveWorkspaceMembersIN([]openapiclient.BaseWorkspaceMemberIN{*openapiclient.NewBaseWorkspaceMemberIN("UserUuid_example")}) // RemoveWorkspaceMembersIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesRemoveUsersWorkspaces(context.Background(), workspaceUuid).RemoveWorkspaceMembersIN(removeWorkspaceMembersIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesRemoveUsersWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminWorkspacesRemoveUsersWorkspaces`: RemoveWorkspaceMembersOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesRemoveUsersWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesRemoveUsersWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeWorkspaceMembersIN** | [**RemoveWorkspaceMembersIN**](RemoveWorkspaceMembersIN.md) |  | 

### Return type

[**RemoveWorkspaceMembersOUT**](RemoveWorkspaceMembersOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminWorkspacesUpdateWorkspaces

> WorkspaceOUT UsersApiAdminWorkspacesUpdateWorkspaces(ctx, workspaceUuid).UpdateWorkspaceIN(updateWorkspaceIN).Execute()

Update Workspaces



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
	workspaceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID.
	updateWorkspaceIN := *openapiclient.NewUpdateWorkspaceIN() // UpdateWorkspaceIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesUpdateWorkspaces(context.Background(), workspaceUuid).UpdateWorkspaceIN(updateWorkspaceIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesUpdateWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminWorkspacesUpdateWorkspaces`: WorkspaceOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminWorkspacesAPI.UsersApiAdminWorkspacesUpdateWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceUuid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminWorkspacesUpdateWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkspaceIN** | [**UpdateWorkspaceIN**](UpdateWorkspaceIN.md) |  | 

### Return type

[**WorkspaceOUT**](WorkspaceOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


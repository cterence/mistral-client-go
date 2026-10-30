# \BetaAdminScimAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiAdminScimSyncGetScimSyncRun**](BetaAdminScimAPI.md#UsersApiAdminScimSyncGetScimSyncRun) | **Get** /v1/admin/scim/sync/{run_id} | Get Scim Sync Run
[**UsersApiAdminScimSyncTriggerScimSync**](BetaAdminScimAPI.md#UsersApiAdminScimSyncTriggerScimSync) | **Post** /v1/admin/scim/sync | Trigger Scim Sync



## UsersApiAdminScimSyncGetScimSyncRun

> AdminScimSyncRunOUT UsersApiAdminScimSyncGetScimSyncRun(ctx, runId).Execute()

Get Scim Sync Run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifier of the SCIM synchronization run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminScimAPI.UsersApiAdminScimSyncGetScimSyncRun(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminScimAPI.UsersApiAdminScimSyncGetScimSyncRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminScimSyncGetScimSyncRun`: AdminScimSyncRunOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminScimAPI.UsersApiAdminScimSyncGetScimSyncRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Identifier of the SCIM synchronization run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminScimSyncGetScimSyncRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminScimSyncRunOUT**](AdminScimSyncRunOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminScimSyncTriggerScimSync

> AdminScimSyncTriggerOUT UsersApiAdminScimSyncTriggerScimSync(ctx).AdminScimSyncTriggerIN(adminScimSyncTriggerIN).Execute()

Trigger Scim Sync



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
	adminScimSyncTriggerIN := *openapiclient.NewAdminScimSyncTriggerIN() // AdminScimSyncTriggerIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminScimAPI.UsersApiAdminScimSyncTriggerScimSync(context.Background()).AdminScimSyncTriggerIN(adminScimSyncTriggerIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminScimAPI.UsersApiAdminScimSyncTriggerScimSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminScimSyncTriggerScimSync`: AdminScimSyncTriggerOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminScimAPI.UsersApiAdminScimSyncTriggerScimSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminScimSyncTriggerScimSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminScimSyncTriggerIN** | [**AdminScimSyncTriggerIN**](AdminScimSyncTriggerIN.md) |  | 

### Return type

[**AdminScimSyncTriggerOUT**](AdminScimSyncTriggerOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


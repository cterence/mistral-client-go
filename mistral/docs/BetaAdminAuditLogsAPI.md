# \BetaAdminAuditLogsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiAdminAuditLogsGetAuditLogs**](BetaAdminAuditLogsAPI.md#UsersApiAdminAuditLogsGetAuditLogs) | **Get** /v1/admin/audit-logs | Get Audit Logs



## UsersApiAdminAuditLogsGetAuditLogs

> []AuditLogOut UsersApiAdminAuditLogsGetAuditLogs(ctx).ActorType(actorType).EventType(eventType).TargetType(targetType).ActorUserUuid(actorUserUuid).Sort(sort).After(after).Before(before).Limit(limit).Execute()

Get Audit Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	actorType := []openapiclient.ActorType{openapiclient.ActorType("HUMAN")} // []ActorType | Actor types to include in the audit log results. (optional)
	eventType := []openapiclient.AuditLogEventType{openapiclient.AuditLogEventType("user.create")} // []AuditLogEventType | Event types to include in the audit log results. (optional)
	targetType := []openapiclient.TargetType{openapiclient.TargetType("USER")} // []TargetType | Target resource types to include in the audit log results. (optional)
	actorUserUuid := "actorUserUuid_example" // string | Filter logs by the UUID of the user who performed the action. (optional)
	sort := TODO // string | Sort order for audit log entries. (optional) (default to "descending")
	after := time.Now() // time.Time | Return audit log entries after this time. (optional)
	before := time.Now() // time.Time | Return audit log entries before this time. (optional)
	limit := int32(56) // int32 | Maximum number of results to return. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminAuditLogsAPI.UsersApiAdminAuditLogsGetAuditLogs(context.Background()).ActorType(actorType).EventType(eventType).TargetType(targetType).ActorUserUuid(actorUserUuid).Sort(sort).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminAuditLogsAPI.UsersApiAdminAuditLogsGetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminAuditLogsGetAuditLogs`: []AuditLogOut
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminAuditLogsAPI.UsersApiAdminAuditLogsGetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminAuditLogsGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorType** | [**[]ActorType**](ActorType.md) | Actor types to include in the audit log results. | 
 **eventType** | [**[]AuditLogEventType**](AuditLogEventType.md) | Event types to include in the audit log results. | 
 **targetType** | [**[]TargetType**](TargetType.md) | Target resource types to include in the audit log results. | 
 **actorUserUuid** | **string** | Filter logs by the UUID of the user who performed the action. | 
 **sort** | [**string**](string.md) | Sort order for audit log entries. | [default to &quot;descending&quot;]
 **after** | **time.Time** | Return audit log entries after this time. | 
 **before** | **time.Time** | Return audit log entries before this time. | 
 **limit** | **int32** | Maximum number of results to return. | [default to 20]

### Return type

[**[]AuditLogOut**](AuditLogOut.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


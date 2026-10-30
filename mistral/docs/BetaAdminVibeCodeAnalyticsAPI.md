# \BetaAdminVibeCodeAnalyticsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization**](BetaAdminVibeCodeAnalyticsAPI.md#GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization) | **Get** /v1/admin/analytics/vibe/code/usage/by_organization | Usage by organization
[**GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace**](BetaAdminVibeCodeAnalyticsAPI.md#GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace) | **Get** /v1/admin/analytics/vibe/code/usage/by_workspace | Usage by workspace



## GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization

> VibeOrganizationStatsOUT GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization(ctx).StartTime(startTime).EndTime(endTime).Execute()

Usage by organization



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
	startTime := int32(56) // int32 | Start of the queried window, as a Unix timestamp in seconds.
	endTime := int32(56) // int32 | End of the queried window, as a Unix timestamp in seconds.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminVibeCodeAnalyticsAPI.GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminVibeCodeAnalyticsAPI.GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization`: VibeOrganizationStatsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminVibeCodeAnalyticsAPI.GetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationStatsV1AdminAnalyticsVibeCodeUsageByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
 **endTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 

### Return type

[**VibeOrganizationStatsOUT**](VibeOrganizationStatsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace

> VibeWorkspaceStatsOUT GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace(ctx).StartTime(startTime).EndTime(endTime).WorkspaceId(workspaceId).Execute()

Usage by workspace



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
	startTime := int32(56) // int32 | Start of the queried window, as a Unix timestamp in seconds.
	endTime := int32(56) // int32 | End of the queried window, as a Unix timestamp in seconds.
	workspaceId := "workspaceId_example" // string | Workspace ID to filter results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminVibeCodeAnalyticsAPI.GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace(context.Background()).StartTime(startTime).EndTime(endTime).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminVibeCodeAnalyticsAPI.GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace`: VibeWorkspaceStatsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminVibeCodeAnalyticsAPI.GetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceStatsV1AdminAnalyticsVibeCodeUsageByWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
 **endTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 
 **workspaceId** | **string** | Workspace ID to filter results. | 

### Return type

[**VibeWorkspaceStatsOUT**](VibeWorkspaceStatsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


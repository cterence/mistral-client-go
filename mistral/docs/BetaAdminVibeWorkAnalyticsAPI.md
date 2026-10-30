# \BetaAdminVibeWorkAnalyticsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats**](BetaAdminVibeWorkAnalyticsAPI.md#GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats) | **Get** /v1/admin/analytics/vibe/work/usage/by_agent_stats | Usage by agent
[**GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats**](BetaAdminVibeWorkAnalyticsAPI.md#GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats) | **Get** /v1/admin/analytics/vibe/work/usage/by_time_stats | Usage over time
[**GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats**](BetaAdminVibeWorkAnalyticsAPI.md#GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats) | **Get** /v1/admin/analytics/vibe/work/usage/by_user_stats | Usage by user



## GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats

> VibeWorkByAgentStatsOUT GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats(ctx).StartTime(startTime).EndTime(endTime).Execute()

Usage by agent



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
	resp, r, err := apiClient.BetaAdminVibeWorkAnalyticsAPI.GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminVibeWorkAnalyticsAPI.GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats`: VibeWorkByAgentStatsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminVibeWorkAnalyticsAPI.GetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByAgentStatsV1AdminAnalyticsVibeWorkUsageByAgentStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
 **endTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 

### Return type

[**VibeWorkByAgentStatsOUT**](VibeWorkByAgentStatsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats

> VibeWorkByTimeStatsOUT GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats(ctx).StartTime(startTime).EndTime(endTime).Granularity(granularity).Execute()

Usage over time



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
	granularity := "granularity_example" // string | Time interval used to group usage results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminVibeWorkAnalyticsAPI.GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats(context.Background()).StartTime(startTime).EndTime(endTime).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminVibeWorkAnalyticsAPI.GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats`: VibeWorkByTimeStatsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminVibeWorkAnalyticsAPI.GetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByTimeStatsV1AdminAnalyticsVibeWorkUsageByTimeStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
 **endTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 
 **granularity** | **string** | Time interval used to group usage results. | 

### Return type

[**VibeWorkByTimeStatsOUT**](VibeWorkByTimeStatsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats

> VibeWorkByUserStatsOUT GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats(ctx).StartTime(startTime).EndTime(endTime).Execute()

Usage by user



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
	resp, r, err := apiClient.BetaAdminVibeWorkAnalyticsAPI.GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminVibeWorkAnalyticsAPI.GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats`: VibeWorkByUserStatsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminVibeWorkAnalyticsAPI.GetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByUserStatsV1AdminAnalyticsVibeWorkUsageByUserStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
 **endTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 

### Return type

[**VibeWorkByUserStatsOUT**](VibeWorkByUserStatsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


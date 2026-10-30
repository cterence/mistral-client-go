# \BetaAdminBillingAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiAdminRateLimitsGetRateLimits**](BetaAdminBillingAPI.md#UsersApiAdminRateLimitsGetRateLimits) | **Get** /v1/admin/rate-limit | Get Rate Limits
[**UsersApiAdminSpendLimitsGetSpendLimits**](BetaAdminBillingAPI.md#UsersApiAdminSpendLimitsGetSpendLimits) | **Get** /v1/admin/spend-limit | Get Spend Limits
[**UsersApiAdminSpendLimitsUpdateSpendLimits**](BetaAdminBillingAPI.md#UsersApiAdminSpendLimitsUpdateSpendLimits) | **Post** /v1/admin/spend-limit | Update Spend Limits
[**UsersApiAdminUsageGetUsage**](BetaAdminBillingAPI.md#UsersApiAdminUsageGetUsage) | **Get** /v1/admin/usage | Monthly Usage and Pricing
[**UsersApiAdminUsageQueryUsage**](BetaAdminBillingAPI.md#UsersApiAdminUsageQueryUsage) | **Post** /v1/admin/usage/query | Detailed Usage Query



## UsersApiAdminRateLimitsGetRateLimits

> RateLimitsOUT UsersApiAdminRateLimitsGetRateLimits(ctx).Execute()

Get Rate Limits



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
	resp, r, err := apiClient.BetaAdminBillingAPI.UsersApiAdminRateLimitsGetRateLimits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminBillingAPI.UsersApiAdminRateLimitsGetRateLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminRateLimitsGetRateLimits`: RateLimitsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminBillingAPI.UsersApiAdminRateLimitsGetRateLimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminRateLimitsGetRateLimitsRequest struct via the builder pattern


### Return type

[**RateLimitsOUT**](RateLimitsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminSpendLimitsGetSpendLimits

> LimitsOUT UsersApiAdminSpendLimitsGetSpendLimits(ctx).Execute()

Get Spend Limits



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
	resp, r, err := apiClient.BetaAdminBillingAPI.UsersApiAdminSpendLimitsGetSpendLimits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminBillingAPI.UsersApiAdminSpendLimitsGetSpendLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminSpendLimitsGetSpendLimits`: LimitsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminBillingAPI.UsersApiAdminSpendLimitsGetSpendLimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminSpendLimitsGetSpendLimitsRequest struct via the builder pattern


### Return type

[**LimitsOUT**](LimitsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminSpendLimitsUpdateSpendLimits

> LimitsOUT UsersApiAdminSpendLimitsUpdateSpendLimits(ctx).NewUsageLimitIN(newUsageLimitIN).Execute()

Update Spend Limits



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
	newUsageLimitIN := *openapiclient.NewNewUsageLimitIN(int32(123)) // NewUsageLimitIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminBillingAPI.UsersApiAdminSpendLimitsUpdateSpendLimits(context.Background()).NewUsageLimitIN(newUsageLimitIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminBillingAPI.UsersApiAdminSpendLimitsUpdateSpendLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminSpendLimitsUpdateSpendLimits`: LimitsOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminBillingAPI.UsersApiAdminSpendLimitsUpdateSpendLimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminSpendLimitsUpdateSpendLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newUsageLimitIN** | [**NewUsageLimitIN**](NewUsageLimitIN.md) |  | 

### Return type

[**LimitsOUT**](LimitsOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsageGetUsage

> UsageOUTJSON UsersApiAdminUsageGetUsage(ctx).Month(month).Year(year).WorkspaceId(workspaceId).ApiZone(apiZone).Execute()

Monthly Usage and Pricing



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
	month := "month_example" // string | Month to return usage for. (optional)
	year := "year_example" // string | Year to return usage for. (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace ID to filter results. (optional)
	apiZone := openapiclient.ApiZone("global") // ApiZone | Regional inference zone to filter results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminBillingAPI.UsersApiAdminUsageGetUsage(context.Background()).Month(month).Year(year).WorkspaceId(workspaceId).ApiZone(apiZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminBillingAPI.UsersApiAdminUsageGetUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsageGetUsage`: UsageOUTJSON
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminBillingAPI.UsersApiAdminUsageGetUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsageGetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | Month to return usage for. | 
 **year** | **string** | Year to return usage for. | 
 **workspaceId** | **string** | Workspace ID to filter results. | 
 **apiZone** | [**ApiZone**](ApiZone.md) | Regional inference zone to filter results. | 

### Return type

[**UsageOUTJSON**](UsageOUTJSON.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiAdminUsageQueryUsage

> UsageQueryOUT UsersApiAdminUsageQueryUsage(ctx).UsageQueryIN(usageQueryIN).Execute()

Detailed Usage Query



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
	usageQueryIN := *openapiclient.NewUsageQueryIN(time.Now(), time.Now()) // UsageQueryIN | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAdminBillingAPI.UsersApiAdminUsageQueryUsage(context.Background()).UsageQueryIN(usageQueryIN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAdminBillingAPI.UsersApiAdminUsageQueryUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiAdminUsageQueryUsage`: UsageQueryOUT
	fmt.Fprintf(os.Stdout, "Response from `BetaAdminBillingAPI.UsersApiAdminUsageQueryUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiAdminUsageQueryUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usageQueryIN** | [**UsageQueryIN**](UsageQueryIN.md) |  | 

### Return type

[**UsageQueryOUT**](UsageQueryOUT.md)

### Authorization

[AdminApiKey](../README.md#AdminApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


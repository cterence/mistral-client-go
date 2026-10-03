# \BetaObservabilityCampaignsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaignV1ObservabilityCampaignsPost**](BetaObservabilityCampaignsAPI.md#CreateCampaignV1ObservabilityCampaignsPost) | **Post** /v1/observability/campaigns | Create and start a new campaign
[**DeleteCampaignV1ObservabilityCampaignsCampaignIdDelete**](BetaObservabilityCampaignsAPI.md#DeleteCampaignV1ObservabilityCampaignsCampaignIdDelete) | **Delete** /v1/observability/campaigns/{campaign_id} | Delete a campaign
[**GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet**](BetaObservabilityCampaignsAPI.md#GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet) | **Get** /v1/observability/campaigns/{campaign_id} | Get campaign by id
[**GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet**](BetaObservabilityCampaignsAPI.md#GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet) | **Get** /v1/observability/campaigns/{campaign_id}/selected-events | Get event ids that were selected by the given campaign
[**GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet**](BetaObservabilityCampaignsAPI.md#GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet) | **Get** /v1/observability/campaigns/{campaign_id}/status | Get campaign status by campaign id
[**GetCampaignsV1ObservabilityCampaignsGet**](BetaObservabilityCampaignsAPI.md#GetCampaignsV1ObservabilityCampaignsGet) | **Get** /v1/observability/campaigns | Get all campaigns



## CreateCampaignV1ObservabilityCampaignsPost

> CampaignPreview CreateCampaignV1ObservabilityCampaignsPost(ctx).PostCampaignInSchema(postCampaignInSchema).Execute()

Create and start a new campaign

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
	postCampaignInSchema := *openapiclient.NewPostCampaignInSchema(*openapiclient.NewFilterPayload("TODO"), "JudgeId_example", "Name_example", "Description_example", int32(123)) // PostCampaignInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityCampaignsAPI.CreateCampaignV1ObservabilityCampaignsPost(context.Background()).PostCampaignInSchema(postCampaignInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityCampaignsAPI.CreateCampaignV1ObservabilityCampaignsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignV1ObservabilityCampaignsPost`: CampaignPreview
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityCampaignsAPI.CreateCampaignV1ObservabilityCampaignsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignV1ObservabilityCampaignsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCampaignInSchema** | [**PostCampaignInSchema**](PostCampaignInSchema.md) |  | 

### Return type

[**CampaignPreview**](CampaignPreview.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignV1ObservabilityCampaignsCampaignIdDelete

> DeleteCampaignV1ObservabilityCampaignsCampaignIdDelete(ctx, campaignId).Execute()

Delete a campaign

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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityCampaignsAPI.DeleteCampaignV1ObservabilityCampaignsCampaignIdDelete(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityCampaignsAPI.DeleteCampaignV1ObservabilityCampaignsCampaignIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignV1ObservabilityCampaignsCampaignIdDeleteRequest struct via the builder pattern


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


## GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet

> CampaignPreview GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet(ctx, campaignId).Execute()

Get campaign by id

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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityCampaignsAPI.GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityCampaignsAPI.GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet`: CampaignPreview
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityCampaignsAPI.GetCampaignByIdV1ObservabilityCampaignsCampaignIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdV1ObservabilityCampaignsCampaignIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignPreview**](CampaignPreview.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet

> CampaignSelectedEvents GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet(ctx, campaignId).PageSize(pageSize).Page(page).Execute()

Get event ids that were selected by the given campaign

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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityCampaignsAPI.GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet(context.Background(), campaignId).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityCampaignsAPI.GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet`: CampaignSelectedEvents
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityCampaignsAPI.GetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignSelectedEventsV1ObservabilityCampaignsCampaignIdSelectedEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**CampaignSelectedEvents**](CampaignSelectedEvents.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet

> CampaignStatus GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet(ctx, campaignId).Execute()

Get campaign status by campaign id

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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityCampaignsAPI.GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityCampaignsAPI.GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet`: CampaignStatus
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityCampaignsAPI.GetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignStatusByIdV1ObservabilityCampaignsCampaignIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignStatus**](CampaignStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsV1ObservabilityCampaignsGet

> CampaignPreviews GetCampaignsV1ObservabilityCampaignsGet(ctx).PageSize(pageSize).Page(page).Q(q).Execute()

Get all campaigns

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
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	q := "q_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityCampaignsAPI.GetCampaignsV1ObservabilityCampaignsGet(context.Background()).PageSize(pageSize).Page(page).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityCampaignsAPI.GetCampaignsV1ObservabilityCampaignsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignsV1ObservabilityCampaignsGet`: CampaignPreviews
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityCampaignsAPI.GetCampaignsV1ObservabilityCampaignsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsV1ObservabilityCampaignsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **q** | **string** |  | 

### Return type

[**CampaignPreviews**](CampaignPreviews.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


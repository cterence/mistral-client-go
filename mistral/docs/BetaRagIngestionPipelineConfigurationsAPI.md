# \BetaRagIngestionPipelineConfigurationsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigsV1RagIngestionPipelineConfigurationsGet**](BetaRagIngestionPipelineConfigurationsAPI.md#GetConfigsV1RagIngestionPipelineConfigurationsGet) | **Get** /v1/rag/ingestion_pipeline_configurations | List ingestion pipeline configurations
[**RegisterConfigV1RagIngestionPipelineConfigurationsPut**](BetaRagIngestionPipelineConfigurationsAPI.md#RegisterConfigV1RagIngestionPipelineConfigurationsPut) | **Put** /v1/rag/ingestion_pipeline_configurations | Register Config
[**UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut**](BetaRagIngestionPipelineConfigurationsAPI.md#UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut) | **Put** /v1/rag/ingestion_pipeline_configurations/{id}/run_info | Update Run Info



## GetConfigsV1RagIngestionPipelineConfigurationsGet

> []IngestionPipelineConfiguration GetConfigsV1RagIngestionPipelineConfigurationsGet(ctx).Execute()

List ingestion pipeline configurations



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
	resp, r, err := apiClient.BetaRagIngestionPipelineConfigurationsAPI.GetConfigsV1RagIngestionPipelineConfigurationsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagIngestionPipelineConfigurationsAPI.GetConfigsV1RagIngestionPipelineConfigurationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigsV1RagIngestionPipelineConfigurationsGet`: []IngestionPipelineConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BetaRagIngestionPipelineConfigurationsAPI.GetConfigsV1RagIngestionPipelineConfigurationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigsV1RagIngestionPipelineConfigurationsGetRequest struct via the builder pattern


### Return type

[**[]IngestionPipelineConfiguration**](IngestionPipelineConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterConfigV1RagIngestionPipelineConfigurationsPut

> IngestionPipelineConfiguration RegisterConfigV1RagIngestionPipelineConfigurationsPut(ctx).CreateIngestionPipelineConfigurationRequest(createIngestionPipelineConfigurationRequest).Execute()

Register Config



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
	createIngestionPipelineConfigurationRequest := *openapiclient.NewCreateIngestionPipelineConfigurationRequest("Name_example") // CreateIngestionPipelineConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagIngestionPipelineConfigurationsAPI.RegisterConfigV1RagIngestionPipelineConfigurationsPut(context.Background()).CreateIngestionPipelineConfigurationRequest(createIngestionPipelineConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagIngestionPipelineConfigurationsAPI.RegisterConfigV1RagIngestionPipelineConfigurationsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterConfigV1RagIngestionPipelineConfigurationsPut`: IngestionPipelineConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BetaRagIngestionPipelineConfigurationsAPI.RegisterConfigV1RagIngestionPipelineConfigurationsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterConfigV1RagIngestionPipelineConfigurationsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIngestionPipelineConfigurationRequest** | [**CreateIngestionPipelineConfigurationRequest**](CreateIngestionPipelineConfigurationRequest.md) |  | 

### Return type

[**IngestionPipelineConfiguration**](IngestionPipelineConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut

> IngestionPipelineConfiguration UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut(ctx, id).UpdateRunInfo(updateRunInfo).Execute()

Update Run Info



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateRunInfo := *openapiclient.NewUpdateRunInfo(time.Now(), int32(123)) // UpdateRunInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagIngestionPipelineConfigurationsAPI.UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut(context.Background(), id).UpdateRunInfo(updateRunInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagIngestionPipelineConfigurationsAPI.UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut`: IngestionPipelineConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BetaRagIngestionPipelineConfigurationsAPI.UpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRunInfoV1RagIngestionPipelineConfigurationsIdRunInfoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRunInfo** | [**UpdateRunInfo**](UpdateRunInfo.md) |  | 

### Return type

[**IngestionPipelineConfiguration**](IngestionPipelineConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


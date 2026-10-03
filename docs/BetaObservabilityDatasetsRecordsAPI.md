# \BetaObservabilityDatasetsRecordsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete**](BetaObservabilityDatasetsRecordsAPI.md#DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete) | **Delete** /v1/observability/dataset-records/{dataset_record_id} | Delete a record from a dataset
[**DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost**](BetaObservabilityDatasetsRecordsAPI.md#DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost) | **Post** /v1/observability/dataset-records/bulk-delete | Delete multiple records from datasets
[**GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet**](BetaObservabilityDatasetsRecordsAPI.md#GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet) | **Get** /v1/observability/dataset-records/{dataset_record_id} | Get the content of a given conversation from a dataset
[**JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost**](BetaObservabilityDatasetsRecordsAPI.md#JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost) | **Post** /v1/observability/dataset-records/{dataset_record_id}/live-judging | Run Judge on a dataset record based on the given options
[**UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut**](BetaObservabilityDatasetsRecordsAPI.md#UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut) | **Put** /v1/observability/dataset-records/{dataset_record_id}/payload | Update a dataset record conversation payload
[**UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut**](BetaObservabilityDatasetsRecordsAPI.md#UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut) | **Put** /v1/observability/dataset-records/{dataset_record_id}/properties | Update conversation properties



## DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete

> DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete(ctx, datasetRecordId).Execute()

Delete a record from a dataset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete(context.Background(), datasetRecordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsRecordsAPI.DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDeleteRequest struct via the builder pattern


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


## DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost

> DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost(ctx).DeleteDatasetRecordsInSchema(deleteDatasetRecordsInSchema).Execute()

Delete multiple records from datasets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	deleteDatasetRecordsInSchema := *openapiclient.NewDeleteDatasetRecordsInSchema([]string{"DatasetRecordIds_example"}) // DeleteDatasetRecordsInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost(context.Background()).DeleteDatasetRecordsInSchema(deleteDatasetRecordsInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsRecordsAPI.DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDatasetRecordsInSchema** | [**DeleteDatasetRecordsInSchema**](DeleteDatasetRecordsInSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet

> DatasetRecord GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet(ctx, datasetRecordId).Execute()

Get the content of a given conversation from a dataset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet(context.Background(), datasetRecordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsRecordsAPI.GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet`: DatasetRecord
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsRecordsAPI.GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetRecord**](DatasetRecord.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost

> JudgeOutput JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost(ctx, datasetRecordId).PostDatasetRecordJudgingInSchema(postDatasetRecordJudgingInSchema).Execute()

Run Judge on a dataset record based on the given options

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetRecordJudgingInSchema := *openapiclient.NewPostDatasetRecordJudgingInSchema(*openapiclient.NewPostJudgeInSchema("Name_example", "Description_example", "ModelName_example", openapiclient.Output{JudgeClassificationOutput: openapiclient.NewJudgeClassificationOutput([]openapiclient.JudgeClassificationOutputOption{*openapiclient.NewJudgeClassificationOutputOption("Value_example", "Description_example")})}, "Instructions_example", []string{"Tools_example"})) // PostDatasetRecordJudgingInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost(context.Background(), datasetRecordId).PostDatasetRecordJudgingInSchema(postDatasetRecordJudgingInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsRecordsAPI.JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost`: JudgeOutput
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsRecordsAPI.JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetRecordJudgingInSchema** | [**PostDatasetRecordJudgingInSchema**](PostDatasetRecordJudgingInSchema.md) |  | 

### Return type

[**JudgeOutput**](JudgeOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut

> UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut(ctx, datasetRecordId).PutDatasetRecordPayloadInSchema(putDatasetRecordPayloadInSchema).Execute()

Update a dataset record conversation payload

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	putDatasetRecordPayloadInSchema := *openapiclient.NewPutDatasetRecordPayloadInSchema(*openapiclient.NewConversationPayload([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}})) // PutDatasetRecordPayloadInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut(context.Background(), datasetRecordId).PutDatasetRecordPayloadInSchema(putDatasetRecordPayloadInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsRecordsAPI.UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putDatasetRecordPayloadInSchema** | [**PutDatasetRecordPayloadInSchema**](PutDatasetRecordPayloadInSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut

> UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut(ctx, datasetRecordId).PutDatasetRecordPropertiesInSchema(putDatasetRecordPropertiesInSchema).Execute()

Update conversation properties

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	putDatasetRecordPropertiesInSchema := *openapiclient.NewPutDatasetRecordPropertiesInSchema(map[string]interface{}{"key": interface{}(123)}) // PutDatasetRecordPropertiesInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut(context.Background(), datasetRecordId).PutDatasetRecordPropertiesInSchema(putDatasetRecordPropertiesInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsRecordsAPI.UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putDatasetRecordPropertiesInSchema** | [**PutDatasetRecordPropertiesInSchema**](PutDatasetRecordPropertiesInSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


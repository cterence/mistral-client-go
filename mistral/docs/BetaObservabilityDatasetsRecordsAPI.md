# \BetaObservabilityDatasetsRecordsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete**](BetaObservabilityDatasetsRecordsAPI.md#DeleteDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdDelete) | **Delete** /v1/observability/dataset-records/{dataset_record_id} | Delete a record from a dataset
[**DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost**](BetaObservabilityDatasetsRecordsAPI.md#DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost) | **Post** /v1/observability/dataset-records/bulk-delete | Delete multiple records from datasets
[**GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet**](BetaObservabilityDatasetsRecordsAPI.md#GetDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdGet) | **Get** /v1/observability/dataset-records/{dataset_record_id} | Get the content of a given dataset record
[**JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost**](BetaObservabilityDatasetsRecordsAPI.md#JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost) | **Post** /v1/observability/dataset-records/{dataset_record_id}/live-judging | Run Judge on a dataset record based on the given options
[**UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut**](BetaObservabilityDatasetsRecordsAPI.md#UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut) | **Put** /v1/observability/dataset-records/{dataset_record_id}/payload | Update a dataset record payload
[**UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut**](BetaObservabilityDatasetsRecordsAPI.md#UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut) | **Put** /v1/observability/dataset-records/{dataset_record_id}/properties | Update dataset record properties



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
	openapiclient "github.com/cterence/mistral-client-go"
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

> DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost(ctx).DeleteDatasetRecordsRequest(deleteDatasetRecordsRequest).Execute()

Delete multiple records from datasets



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
	deleteDatasetRecordsRequest := *openapiclient.NewDeleteDatasetRecordsRequest([]string{"DatasetRecordIds_example"}) // DeleteDatasetRecordsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.DeleteDatasetRecordsV1ObservabilityDatasetRecordsBulkDeletePost(context.Background()).DeleteDatasetRecordsRequest(deleteDatasetRecordsRequest).Execute()
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
 **deleteDatasetRecordsRequest** | [**DeleteDatasetRecordsRequest**](DeleteDatasetRecordsRequest.md) |  | 

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

Get the content of a given dataset record



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

> JudgeOutput JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost(ctx, datasetRecordId).JudgeDatasetRecordRequest(judgeDatasetRecordRequest).Execute()

Run Judge on a dataset record based on the given options



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
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	judgeDatasetRecordRequest := *openapiclient.NewJudgeDatasetRecordRequest(*openapiclient.NewCreateJudgeRequest("Name_example", "Description_example", "ModelName_example", openapiclient.Output{JudgeClassificationOutput: openapiclient.NewJudgeClassificationOutput([]openapiclient.JudgeClassificationOutputOption{*openapiclient.NewJudgeClassificationOutputOption("Value_example", "Description_example")})}, "Instructions_example", []string{"Tools_example"})) // JudgeDatasetRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.JudgeDatasetRecordV1ObservabilityDatasetRecordsDatasetRecordIdLiveJudgingPost(context.Background(), datasetRecordId).JudgeDatasetRecordRequest(judgeDatasetRecordRequest).Execute()
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

 **judgeDatasetRecordRequest** | [**JudgeDatasetRecordRequest**](JudgeDatasetRecordRequest.md) |  | 

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

> UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut(ctx, datasetRecordId).UpdateDatasetRecordPayloadRequest(updateDatasetRecordPayloadRequest).Execute()

Update a dataset record payload



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
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateDatasetRecordPayloadRequest := *openapiclient.NewUpdateDatasetRecordPayloadRequest(map[string]interface{}{"key": interface{}(123)}) // UpdateDatasetRecordPayloadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.UpdateDatasetRecordPayloadV1ObservabilityDatasetRecordsDatasetRecordIdPayloadPut(context.Background(), datasetRecordId).UpdateDatasetRecordPayloadRequest(updateDatasetRecordPayloadRequest).Execute()
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

 **updateDatasetRecordPayloadRequest** | [**UpdateDatasetRecordPayloadRequest**](UpdateDatasetRecordPayloadRequest.md) |  | 

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

> UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut(ctx, datasetRecordId).UpdateDatasetRecordPropertiesRequest(updateDatasetRecordPropertiesRequest).Execute()

Update dataset record properties



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
	datasetRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateDatasetRecordPropertiesRequest := *openapiclient.NewUpdateDatasetRecordPropertiesRequest(map[string]interface{}{"key": interface{}(123)}) // UpdateDatasetRecordPropertiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsRecordsAPI.UpdateDatasetRecordPropertiesV1ObservabilityDatasetRecordsDatasetRecordIdPropertiesPut(context.Background(), datasetRecordId).UpdateDatasetRecordPropertiesRequest(updateDatasetRecordPropertiesRequest).Execute()
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

 **updateDatasetRecordPropertiesRequest** | [**UpdateDatasetRecordPropertiesRequest**](UpdateDatasetRecordPropertiesRequest.md) |  | 

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


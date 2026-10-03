# \BetaObservabilityDatasetsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost**](BetaObservabilityDatasetsAPI.md#CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost) | **Post** /v1/observability/datasets/{dataset_id}/records | Add a conversation to the dataset
[**CreateDatasetV1ObservabilityDatasetsPost**](BetaObservabilityDatasetsAPI.md#CreateDatasetV1ObservabilityDatasetsPost) | **Post** /v1/observability/datasets | Create a new empty dataset
[**DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete**](BetaObservabilityDatasetsAPI.md#DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete) | **Delete** /v1/observability/datasets/{dataset_id} | Delete a dataset
[**ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet**](BetaObservabilityDatasetsAPI.md#ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet) | **Get** /v1/observability/datasets/{dataset_id}/exports/to-jsonl | Export to the Files API and retrieve presigned URL to download the resulting JSONL file
[**GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet**](BetaObservabilityDatasetsAPI.md#GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet) | **Get** /v1/observability/datasets/{dataset_id} | Get dataset by id
[**GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet**](BetaObservabilityDatasetsAPI.md#GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet) | **Get** /v1/observability/datasets/{dataset_id}/tasks/{task_id} | Get status of a dataset import task
[**GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet**](BetaObservabilityDatasetsAPI.md#GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet) | **Get** /v1/observability/datasets/{dataset_id}/tasks | List import tasks for the given dataset
[**GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet**](BetaObservabilityDatasetsAPI.md#GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet) | **Get** /v1/observability/datasets/{dataset_id}/records | List existing records in the dataset
[**GetDatasetsV1ObservabilityDatasetsGet**](BetaObservabilityDatasetsAPI.md#GetDatasetsV1ObservabilityDatasetsGet) | **Get** /v1/observability/datasets | List existing datasets
[**PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-campaign | Populate the dataset with a campaign
[**PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-dataset | Populate the dataset with samples from another dataset
[**PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-explorer | Populate the dataset with samples from the explorer
[**PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-file | Populate the dataset with samples from an uploaded file
[**PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-playground | Populate the dataset with samples from the playground
[**UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch**](BetaObservabilityDatasetsAPI.md#UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch) | **Patch** /v1/observability/datasets/{dataset_id} | Patch dataset



## CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost

> DatasetRecord CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost(ctx, datasetId).PostDatasetRecordInSchema(postDatasetRecordInSchema).Execute()

Add a conversation to the dataset

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetRecordInSchema := *openapiclient.NewPostDatasetRecordInSchema(*openapiclient.NewConversationPayload([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), map[string]interface{}{"key": interface{}(123)}) // PostDatasetRecordInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost(context.Background(), datasetId).PostDatasetRecordInSchema(postDatasetRecordInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost`: DatasetRecord
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetRecordInSchema** | [**PostDatasetRecordInSchema**](PostDatasetRecordInSchema.md) |  | 

### Return type

[**DatasetRecord**](DatasetRecord.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatasetV1ObservabilityDatasetsPost

> Dataset CreateDatasetV1ObservabilityDatasetsPost(ctx).PostDatasetInSchema(postDatasetInSchema).Execute()

Create a new empty dataset

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
	postDatasetInSchema := *openapiclient.NewPostDatasetInSchema("Name_example", "Description_example") // PostDatasetInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.CreateDatasetV1ObservabilityDatasetsPost(context.Background()).PostDatasetInSchema(postDatasetInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.CreateDatasetV1ObservabilityDatasetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatasetV1ObservabilityDatasetsPost`: Dataset
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.CreateDatasetV1ObservabilityDatasetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatasetV1ObservabilityDatasetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postDatasetInSchema** | [**PostDatasetInSchema**](PostDatasetInSchema.md) |  | 

### Return type

[**Dataset**](Dataset.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete

> DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete(ctx, datasetId).Execute()

Delete a dataset

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaObservabilityDatasetsAPI.DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete(context.Background(), datasetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatasetV1ObservabilityDatasetsDatasetIdDeleteRequest struct via the builder pattern


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


## ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet

> DatasetExport ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet(ctx, datasetId).Execute()

Export to the Files API and retrieve presigned URL to download the resulting JSONL file

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet(context.Background(), datasetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet`: DatasetExport
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetExport**](DatasetExport.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet

> DatasetPreview GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet(ctx, datasetId).Execute()

Get dataset by id

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet(context.Background(), datasetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet`: DatasetPreview
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetByIdV1ObservabilityDatasetsDatasetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetPreview**](DatasetPreview.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet

> DatasetImportTask GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet(ctx, datasetId, taskId).Execute()

Get status of a dataset import task

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet(context.Background(), datasetId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DatasetImportTask**](DatasetImportTask.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet

> DatasetImportTasks GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet(ctx, datasetId).PageSize(pageSize).Page(page).Execute()

List import tasks for the given dataset

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet(context.Background(), datasetId).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet`: DatasetImportTasks
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**DatasetImportTasks**](DatasetImportTasks.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet

> DatasetRecords GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet(ctx, datasetId).PageSize(pageSize).Page(page).Execute()

List existing records in the dataset

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet(context.Background(), datasetId).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet`: DatasetRecords
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**DatasetRecords**](DatasetRecords.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetsV1ObservabilityDatasetsGet

> DatasetPreviews GetDatasetsV1ObservabilityDatasetsGet(ctx).PageSize(pageSize).Page(page).Q(q).Execute()

List existing datasets

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
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	q := "q_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetsV1ObservabilityDatasetsGet(context.Background()).PageSize(pageSize).Page(page).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetsV1ObservabilityDatasetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetsV1ObservabilityDatasetsGet`: DatasetPreviews
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetsV1ObservabilityDatasetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetsV1ObservabilityDatasetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **q** | **string** |  | 

### Return type

[**DatasetPreviews**](DatasetPreviews.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost

> DatasetImportTask PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost(ctx, datasetId).PostDatasetImportFromCampaignInSchema(postDatasetImportFromCampaignInSchema).Execute()

Populate the dataset with a campaign

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetImportFromCampaignInSchema := *openapiclient.NewPostDatasetImportFromCampaignInSchema("CampaignId_example") // PostDatasetImportFromCampaignInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost(context.Background(), datasetId).PostDatasetImportFromCampaignInSchema(postDatasetImportFromCampaignInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromCampaignV1ObservabilityDatasetsDatasetIdImportsFromCampaignPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetImportFromCampaignInSchema** | [**PostDatasetImportFromCampaignInSchema**](PostDatasetImportFromCampaignInSchema.md) |  | 

### Return type

[**DatasetImportTask**](DatasetImportTask.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost

> DatasetImportTask PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost(ctx, datasetId).PostDatasetImportFromDatasetInSchema(postDatasetImportFromDatasetInSchema).Execute()

Populate the dataset with samples from another dataset

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetImportFromDatasetInSchema := *openapiclient.NewPostDatasetImportFromDatasetInSchema([]string{"DatasetRecordIds_example"}) // PostDatasetImportFromDatasetInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost(context.Background(), datasetId).PostDatasetImportFromDatasetInSchema(postDatasetImportFromDatasetInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetImportFromDatasetInSchema** | [**PostDatasetImportFromDatasetInSchema**](PostDatasetImportFromDatasetInSchema.md) |  | 

### Return type

[**DatasetImportTask**](DatasetImportTask.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost

> DatasetImportTask PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost(ctx, datasetId).PostDatasetImportFromExplorerInSchema(postDatasetImportFromExplorerInSchema).Execute()

Populate the dataset with samples from the explorer

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetImportFromExplorerInSchema := *openapiclient.NewPostDatasetImportFromExplorerInSchema([]string{"CompletionEventIds_example"}) // PostDatasetImportFromExplorerInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost(context.Background(), datasetId).PostDatasetImportFromExplorerInSchema(postDatasetImportFromExplorerInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromExplorerV1ObservabilityDatasetsDatasetIdImportsFromExplorerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetImportFromExplorerInSchema** | [**PostDatasetImportFromExplorerInSchema**](PostDatasetImportFromExplorerInSchema.md) |  | 

### Return type

[**DatasetImportTask**](DatasetImportTask.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost

> DatasetImportTask PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost(ctx, datasetId).PostDatasetImportFromFileInSchema(postDatasetImportFromFileInSchema).Execute()

Populate the dataset with samples from an uploaded file

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetImportFromFileInSchema := *openapiclient.NewPostDatasetImportFromFileInSchema("FileId_example") // PostDatasetImportFromFileInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost(context.Background(), datasetId).PostDatasetImportFromFileInSchema(postDatasetImportFromFileInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetImportFromFileInSchema** | [**PostDatasetImportFromFileInSchema**](PostDatasetImportFromFileInSchema.md) |  | 

### Return type

[**DatasetImportTask**](DatasetImportTask.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost

> DatasetImportTask PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost(ctx, datasetId).PostDatasetImportFromPlaygroundInSchema(postDatasetImportFromPlaygroundInSchema).Execute()

Populate the dataset with samples from the playground

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postDatasetImportFromPlaygroundInSchema := *openapiclient.NewPostDatasetImportFromPlaygroundInSchema([]string{"ConversationIds_example"}) // PostDatasetImportFromPlaygroundInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost(context.Background(), datasetId).PostDatasetImportFromPlaygroundInSchema(postDatasetImportFromPlaygroundInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatasetImportFromPlaygroundInSchema** | [**PostDatasetImportFromPlaygroundInSchema**](PostDatasetImportFromPlaygroundInSchema.md) |  | 

### Return type

[**DatasetImportTask**](DatasetImportTask.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch

> DatasetPreview UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch(ctx, datasetId).PatchDatasetInSchema(patchDatasetInSchema).Execute()

Patch dataset

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
	datasetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchDatasetInSchema := *openapiclient.NewPatchDatasetInSchema() // PatchDatasetInSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch(context.Background(), datasetId).PatchDatasetInSchema(patchDatasetInSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch`: DatasetPreview
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatasetV1ObservabilityDatasetsDatasetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDatasetInSchema** | [**PatchDatasetInSchema**](PatchDatasetInSchema.md) |  | 

### Return type

[**DatasetPreview**](DatasetPreview.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


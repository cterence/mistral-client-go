# \BetaObservabilityDatasetsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost**](BetaObservabilityDatasetsAPI.md#CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost) | **Post** /v1/observability/datasets/{dataset_id}/records | Add a record to the dataset
[**CreateDatasetV1ObservabilityDatasetsPost**](BetaObservabilityDatasetsAPI.md#CreateDatasetV1ObservabilityDatasetsPost) | **Post** /v1/observability/datasets | Create a new empty dataset
[**DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete**](BetaObservabilityDatasetsAPI.md#DeleteDatasetV1ObservabilityDatasetsDatasetIdDelete) | **Delete** /v1/observability/datasets/{dataset_id} | Delete a dataset
[**ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet**](BetaObservabilityDatasetsAPI.md#ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet) | **Get** /v1/observability/datasets/{dataset_id}/exports/to-jsonl | Export to the Files API and retrieve presigned URL to download the resulting JSONL file
[**GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet**](BetaObservabilityDatasetsAPI.md#GetDatasetByIdV1ObservabilityDatasetsDatasetIdGet) | **Get** /v1/observability/datasets/{dataset_id} | Get dataset by ID or slug
[**GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet**](BetaObservabilityDatasetsAPI.md#GetDatasetImportTaskV1ObservabilityDatasetsDatasetIdTasksTaskIdGet) | **Get** /v1/observability/datasets/{dataset_id}/tasks/{task_id} | Get status of a dataset import task
[**GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet**](BetaObservabilityDatasetsAPI.md#GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet) | **Get** /v1/observability/datasets/{dataset_id}/tasks | List import tasks for the given dataset
[**GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet**](BetaObservabilityDatasetsAPI.md#GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet) | **Get** /v1/observability/datasets/{dataset_id}/records | List existing records in the dataset
[**GetDatasetsV1ObservabilityDatasetsGet**](BetaObservabilityDatasetsAPI.md#GetDatasetsV1ObservabilityDatasetsGet) | **Get** /v1/observability/datasets | List existing datasets
[**PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-dataset | Populate the dataset with records from another dataset
[**PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-file | Populate the dataset with records from an uploaded file
[**PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-playground | Populate the dataset with records from playground conversations
[**PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost**](BetaObservabilityDatasetsAPI.md#PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost) | **Post** /v1/observability/datasets/{dataset_id}/imports/from-spans | Populate the dataset with records mapped from telemetry spans
[**UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch**](BetaObservabilityDatasetsAPI.md#UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch) | **Patch** /v1/observability/datasets/{dataset_id} | Patch dataset



## CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost

> DatasetRecord CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost(ctx, datasetId).CreateDatasetRecordRequest(createDatasetRecordRequest).Execute()

Add a record to the dataset



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	createDatasetRecordRequest := *openapiclient.NewCreateDatasetRecordRequest(map[string]interface{}{"key": interface{}(123)}) // CreateDatasetRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.CreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPost(context.Background(), datasetId).CreateDatasetRecordRequest(createDatasetRecordRequest).Execute()
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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatasetRecordV1ObservabilityDatasetsDatasetIdRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDatasetRecordRequest** | [**CreateDatasetRecordRequest**](CreateDatasetRecordRequest.md) |  | 

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

> Dataset CreateDatasetV1ObservabilityDatasetsPost(ctx).CreateDatasetRequest(createDatasetRequest).Execute()

Create a new empty dataset



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
	createDatasetRequest := *openapiclient.NewCreateDatasetRequest("Name_example", "Description_example") // CreateDatasetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.CreateDatasetV1ObservabilityDatasetsPost(context.Background()).CreateDatasetRequest(createDatasetRequest).Execute()
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
 **createDatasetRequest** | [**CreateDatasetRequest**](CreateDatasetRequest.md) |  | 

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
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.

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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

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

> ExportDatasetResponse ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet(ctx, datasetId).Execute()

Export to the Files API and retrieve presigned URL to download the resulting JSONL file



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet(context.Background(), datasetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet`: ExportDatasetResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.ExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDatasetToJsonlV1ObservabilityDatasetsDatasetIdExportsToJsonlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportDatasetResponse**](ExportDatasetResponse.md)

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

Get dataset by ID or slug



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.

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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

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
	openapiclient "github.com/cterence/mistral-client-go"
)

func main() {
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 
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

> ListDatasetImportTasksResponse GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet(ctx, datasetId).PageSize(pageSize).Page(page).Execute()

List import tasks for the given dataset



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet(context.Background(), datasetId).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet`: ListDatasetImportTasksResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetImportTasksV1ObservabilityDatasetsDatasetIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**ListDatasetImportTasksResponse**](ListDatasetImportTasksResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet

> ListDatasetRecordsResponse GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet(ctx, datasetId).PageSize(pageSize).Page(page).Execute()

List existing records in the dataset



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet(context.Background(), datasetId).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet`: ListDatasetRecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.GetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetRecordsV1ObservabilityDatasetsDatasetIdRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]

### Return type

[**ListDatasetRecordsResponse**](ListDatasetRecordsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetsV1ObservabilityDatasetsGet

> ListDatasetsResponse GetDatasetsV1ObservabilityDatasetsGet(ctx).PageSize(pageSize).Page(page).Q(q).Execute()

List existing datasets



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
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.GetDatasetsV1ObservabilityDatasetsGet(context.Background()).PageSize(pageSize).Page(page).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.GetDatasetsV1ObservabilityDatasetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetsV1ObservabilityDatasetsGet`: ListDatasetsResponse
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

[**ListDatasetsResponse**](ListDatasetsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost

> DatasetImportTask PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost(ctx, datasetId).ImportDatasetFromDatasetRequest(importDatasetFromDatasetRequest).Execute()

Populate the dataset with records from another dataset



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	importDatasetFromDatasetRequest := *openapiclient.NewImportDatasetFromDatasetRequest([]string{"DatasetRecordIds_example"}) // ImportDatasetFromDatasetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPost(context.Background(), datasetId).ImportDatasetFromDatasetRequest(importDatasetFromDatasetRequest).Execute()
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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromDatasetV1ObservabilityDatasetsDatasetIdImportsFromDatasetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importDatasetFromDatasetRequest** | [**ImportDatasetFromDatasetRequest**](ImportDatasetFromDatasetRequest.md) |  | 

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

> DatasetImportTask PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost(ctx, datasetId).ImportDatasetFromFileRequest(importDatasetFromFileRequest).Execute()

Populate the dataset with records from an uploaded file



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	importDatasetFromFileRequest := *openapiclient.NewImportDatasetFromFileRequest("FileId_example") // ImportDatasetFromFileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePost(context.Background(), datasetId).ImportDatasetFromFileRequest(importDatasetFromFileRequest).Execute()
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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromFileV1ObservabilityDatasetsDatasetIdImportsFromFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importDatasetFromFileRequest** | [**ImportDatasetFromFileRequest**](ImportDatasetFromFileRequest.md) |  | 

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

> DatasetImportTask PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost(ctx, datasetId).ImportDatasetFromPlaygroundRequest(importDatasetFromPlaygroundRequest).Execute()

Populate the dataset with records from playground conversations



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	importDatasetFromPlaygroundRequest := *openapiclient.NewImportDatasetFromPlaygroundRequest([]string{"ConversationIds_example"}) // ImportDatasetFromPlaygroundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPost(context.Background(), datasetId).ImportDatasetFromPlaygroundRequest(importDatasetFromPlaygroundRequest).Execute()
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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromPlaygroundV1ObservabilityDatasetsDatasetIdImportsFromPlaygroundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importDatasetFromPlaygroundRequest** | [**ImportDatasetFromPlaygroundRequest**](ImportDatasetFromPlaygroundRequest.md) |  | 

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


## PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost

> DatasetImportTask PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost(ctx, datasetId).ImportDatasetFromSpansRequest(importDatasetFromSpansRequest).Execute()

Populate the dataset with records mapped from telemetry spans



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	importDatasetFromSpansRequest := *openapiclient.NewImportDatasetFromSpansRequest([]openapiclient.TelemetrySpanReference{*openapiclient.NewTelemetrySpanReference("TraceId_example", "SpanId_example")}, *openapiclient.NewSpanDatasetMappingContract(int32(123), []openapiclient.SpanDatasetMapping{*openapiclient.NewSpanDatasetMapping("TargetField_example", *openapiclient.NewSpanDatasetSourceReference(openapiclient.SpanDatasetSourceNamespace("span_attributes"), "Key_example"))})) // ImportDatasetFromSpansRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost(context.Background(), datasetId).ImportDatasetFromSpansRequest(importDatasetFromSpansRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost`: DatasetImportTask
	fmt.Fprintf(os.Stdout, "Response from `BetaObservabilityDatasetsAPI.PostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatasetRecordsFromSpansV1ObservabilityDatasetsDatasetIdImportsFromSpansPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importDatasetFromSpansRequest** | [**ImportDatasetFromSpansRequest**](ImportDatasetFromSpansRequest.md) |  | 

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

> DatasetPreview UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch(ctx, datasetId).UpdateDatasetRequest(updateDatasetRequest).Execute()

Patch dataset



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
	datasetId := "datasetId_example" // string | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs.
	updateDatasetRequest := *openapiclient.NewUpdateDatasetRequest() // UpdateDatasetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaObservabilityDatasetsAPI.UpdateDatasetV1ObservabilityDatasetsDatasetIdPatch(context.Background(), datasetId).UpdateDatasetRequest(updateDatasetRequest).Execute()
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
**datasetId** | **string** | Dataset UUID or workspace-scoped slug. UUID-shaped values are always interpreted as dataset IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatasetV1ObservabilityDatasetsDatasetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatasetRequest** | [**UpdateDatasetRequest**](UpdateDatasetRequest.md) |  | 

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


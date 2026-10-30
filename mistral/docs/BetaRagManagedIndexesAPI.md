# \BetaRagManagedIndexesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndexV1RagManagedIndexesPost**](BetaRagManagedIndexesAPI.md#CreateIndexV1RagManagedIndexesPost) | **Post** /v1/rag/managed_indexes | Create a managed index
[**DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete**](BetaRagManagedIndexesAPI.md#DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete) | **Delete** /v1/rag/managed_indexes/{index_name}/documents | Delete documents from a managed index
[**DeleteIndexV1RagManagedIndexesIndexNameDelete**](BetaRagManagedIndexesAPI.md#DeleteIndexV1RagManagedIndexesIndexNameDelete) | **Delete** /v1/rag/managed_indexes/{index_name} | Delete a managed index
[**GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet**](BetaRagManagedIndexesAPI.md#GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet) | **Get** /v1/rag/managed_indexes/{index_name}/chunks/{chunk_id} | Get a single chunk by id
[**GetIndexV1RagManagedIndexesIndexNameGet**](BetaRagManagedIndexesAPI.md#GetIndexV1RagManagedIndexesIndexNameGet) | **Get** /v1/rag/managed_indexes/{index_name} | Get a managed index
[**GrepIndexV1RagManagedIndexesIndexNameGrepPost**](BetaRagManagedIndexesAPI.md#GrepIndexV1RagManagedIndexesIndexNameGrepPost) | **Post** /v1/rag/managed_indexes/{index_name}/grep | Grep for a pattern within a source
[**IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost**](BetaRagManagedIndexesAPI.md#IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost) | **Post** /v1/rag/managed_indexes/{index_name}/documents | Ingest documents into a managed index
[**ListIndexesV1RagManagedIndexesGet**](BetaRagManagedIndexesAPI.md#ListIndexesV1RagManagedIndexesGet) | **Get** /v1/rag/managed_indexes | List managed indexes
[**NavigateIndexV1RagManagedIndexesIndexNameNavigatePost**](BetaRagManagedIndexesAPI.md#NavigateIndexV1RagManagedIndexesIndexNameNavigatePost) | **Post** /v1/rag/managed_indexes/{index_name}/navigate | Navigate to adjacent chunks
[**ReadIndexV1RagManagedIndexesIndexNameReadPost**](BetaRagManagedIndexesAPI.md#ReadIndexV1RagManagedIndexesIndexNameReadPost) | **Post** /v1/rag/managed_indexes/{index_name}/read | Read chunks within a span
[**SearchIndexV1RagManagedIndexesIndexNameSearchPost**](BetaRagManagedIndexesAPI.md#SearchIndexV1RagManagedIndexesIndexNameSearchPost) | **Post** /v1/rag/managed_indexes/{index_name}/search | Search a managed index
[**UpdateIndexV1RagManagedIndexesIndexNamePut**](BetaRagManagedIndexesAPI.md#UpdateIndexV1RagManagedIndexesIndexNamePut) | **Put** /v1/rag/managed_indexes/{index_name} | Update a managed index schema



## CreateIndexV1RagManagedIndexesPost

> ManagedIndexResponse CreateIndexV1RagManagedIndexesPost(ctx).CreateManagedIndexRequest(createManagedIndexRequest).Execute()

Create a managed index



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
	createManagedIndexRequest := *openapiclient.NewCreateManagedIndexRequest("Name_example", *openapiclient.NewManagedIndexConfig(openapiclient.Embedding{CustomEmbeddingModel: openapiclient.NewCustomEmbeddingModel("Name_example", int32(123))})) // CreateManagedIndexRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.CreateIndexV1RagManagedIndexesPost(context.Background()).CreateManagedIndexRequest(createManagedIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.CreateIndexV1RagManagedIndexesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIndexV1RagManagedIndexesPost`: ManagedIndexResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.CreateIndexV1RagManagedIndexesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexV1RagManagedIndexesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createManagedIndexRequest** | [**CreateManagedIndexRequest**](CreateManagedIndexRequest.md) |  | 

### Return type

[**ManagedIndexResponse**](ManagedIndexResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete

> DeleteDocumentsResponse DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete(ctx, indexName).DeleteDocumentsRequest(deleteDocumentsRequest).Execute()

Delete documents from a managed index



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
	indexName := "indexName_example" // string | 
	deleteDocumentsRequest := *openapiclient.NewDeleteDocumentsRequest([]string{"DocumentIds_example"}) // DeleteDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete(context.Background(), indexName).DeleteDocumentsRequest(deleteDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete`: DeleteDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.DeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentsV1RagManagedIndexesIndexNameDocumentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteDocumentsRequest** | [**DeleteDocumentsRequest**](DeleteDocumentsRequest.md) |  | 

### Return type

[**DeleteDocumentsResponse**](DeleteDocumentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndexV1RagManagedIndexesIndexNameDelete

> DeleteManagedIndexResponse DeleteIndexV1RagManagedIndexesIndexNameDelete(ctx, indexName).Execute()

Delete a managed index



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
	indexName := "indexName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.DeleteIndexV1RagManagedIndexesIndexNameDelete(context.Background(), indexName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.DeleteIndexV1RagManagedIndexesIndexNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIndexV1RagManagedIndexesIndexNameDelete`: DeleteManagedIndexResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.DeleteIndexV1RagManagedIndexesIndexNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexV1RagManagedIndexesIndexNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteManagedIndexResponse**](DeleteManagedIndexResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet

> NavigationResponse GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet(ctx, indexName, chunkId).Execute()

Get a single chunk by id



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
	indexName := "indexName_example" // string | 
	chunkId := "chunkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet(context.Background(), indexName, chunkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet`: NavigationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.GetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 
**chunkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChunkIndexV1RagManagedIndexesIndexNameChunksChunkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NavigationResponse**](NavigationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexV1RagManagedIndexesIndexNameGet

> ManagedIndexResponse GetIndexV1RagManagedIndexesIndexNameGet(ctx, indexName).Execute()

Get a managed index



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
	indexName := "indexName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.GetIndexV1RagManagedIndexesIndexNameGet(context.Background(), indexName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.GetIndexV1RagManagedIndexesIndexNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexV1RagManagedIndexesIndexNameGet`: ManagedIndexResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.GetIndexV1RagManagedIndexesIndexNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexV1RagManagedIndexesIndexNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedIndexResponse**](ManagedIndexResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrepIndexV1RagManagedIndexesIndexNameGrepPost

> NavigationResponse GrepIndexV1RagManagedIndexesIndexNameGrepPost(ctx, indexName).GrepRequest(grepRequest).Execute()

Grep for a pattern within a source



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
	indexName := "indexName_example" // string | 
	grepRequest := *openapiclient.NewGrepRequest("SourceId_example", "Pattern_example") // GrepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.GrepIndexV1RagManagedIndexesIndexNameGrepPost(context.Background(), indexName).GrepRequest(grepRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.GrepIndexV1RagManagedIndexesIndexNameGrepPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrepIndexV1RagManagedIndexesIndexNameGrepPost`: NavigationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.GrepIndexV1RagManagedIndexesIndexNameGrepPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrepIndexV1RagManagedIndexesIndexNameGrepPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grepRequest** | [**GrepRequest**](GrepRequest.md) |  | 

### Return type

[**NavigationResponse**](NavigationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost

> IngestDocumentsResponse IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost(ctx, indexName).IngestDocumentsRequest(ingestDocumentsRequest).Execute()

Ingest documents into a managed index



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
	indexName := "indexName_example" // string | 
	ingestDocumentsRequest := *openapiclient.NewIngestDocumentsRequest([]*map[string]interface{}{nil}) // IngestDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost(context.Background(), indexName).IngestDocumentsRequest(ingestDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost`: IngestDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.IngestDocumentsV1RagManagedIndexesIndexNameDocumentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIngestDocumentsV1RagManagedIndexesIndexNameDocumentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingestDocumentsRequest** | [**IngestDocumentsRequest**](IngestDocumentsRequest.md) |  | 

### Return type

[**IngestDocumentsResponse**](IngestDocumentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexesV1RagManagedIndexesGet

> ListManagedIndexesResponse ListIndexesV1RagManagedIndexesGet(ctx).PageSize(pageSize).PageToken(pageToken).Name(name).Status(status).CreatorId(creatorId).Execute()

List managed indexes



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
	pageSize := int32(56) // int32 | Maximum number of indexes to return (optional) (default to 20)
	pageToken := "pageToken_example" // string | Cursor returned as next_page_token by the previous page (optional)
	name := "name_example" // string | Case-insensitive substring to match against index names (optional)
	status := openapiclient.ManagedIndexStatus("provisioning") // ManagedIndexStatus | Status to match (optional)
	creatorId := "creatorId_example" // string | Creator ID to match (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.ListIndexesV1RagManagedIndexesGet(context.Background()).PageSize(pageSize).PageToken(pageToken).Name(name).Status(status).CreatorId(creatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.ListIndexesV1RagManagedIndexesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndexesV1RagManagedIndexesGet`: ListManagedIndexesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.ListIndexesV1RagManagedIndexesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIndexesV1RagManagedIndexesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Maximum number of indexes to return | [default to 20]
 **pageToken** | **string** | Cursor returned as next_page_token by the previous page | 
 **name** | **string** | Case-insensitive substring to match against index names | 
 **status** | [**ManagedIndexStatus**](ManagedIndexStatus.md) | Status to match | 
 **creatorId** | **string** | Creator ID to match | 

### Return type

[**ListManagedIndexesResponse**](ListManagedIndexesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NavigateIndexV1RagManagedIndexesIndexNameNavigatePost

> NavigationResponse NavigateIndexV1RagManagedIndexesIndexNameNavigatePost(ctx, indexName).NavigateRequest(navigateRequest).Execute()

Navigate to adjacent chunks



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
	indexName := "indexName_example" // string | 
	navigateRequest := *openapiclient.NewNavigateRequest("SourceId_example", int32(123), int32(123), openapiclient.NavigationDirection("next")) // NavigateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.NavigateIndexV1RagManagedIndexesIndexNameNavigatePost(context.Background(), indexName).NavigateRequest(navigateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.NavigateIndexV1RagManagedIndexesIndexNameNavigatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NavigateIndexV1RagManagedIndexesIndexNameNavigatePost`: NavigationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.NavigateIndexV1RagManagedIndexesIndexNameNavigatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNavigateIndexV1RagManagedIndexesIndexNameNavigatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **navigateRequest** | [**NavigateRequest**](NavigateRequest.md) |  | 

### Return type

[**NavigationResponse**](NavigationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIndexV1RagManagedIndexesIndexNameReadPost

> NavigationResponse ReadIndexV1RagManagedIndexesIndexNameReadPost(ctx, indexName).ReadRequest(readRequest).Execute()

Read chunks within a span



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
	indexName := "indexName_example" // string | 
	readRequest := *openapiclient.NewReadRequest("SourceId_example") // ReadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.ReadIndexV1RagManagedIndexesIndexNameReadPost(context.Background(), indexName).ReadRequest(readRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.ReadIndexV1RagManagedIndexesIndexNameReadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadIndexV1RagManagedIndexesIndexNameReadPost`: NavigationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.ReadIndexV1RagManagedIndexesIndexNameReadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndexV1RagManagedIndexesIndexNameReadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **readRequest** | [**ReadRequest**](ReadRequest.md) |  | 

### Return type

[**NavigationResponse**](NavigationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIndexV1RagManagedIndexesIndexNameSearchPost

> SearchResponse SearchIndexV1RagManagedIndexesIndexNameSearchPost(ctx, indexName).SearchRequest(searchRequest).Execute()

Search a managed index



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
	indexName := "indexName_example" // string | 
	searchRequest := *openapiclient.NewSearchRequest(openapiclient.Retriever{KeywordRetriever: openapiclient.NewKeywordRetriever("Query_example")}) // SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.SearchIndexV1RagManagedIndexesIndexNameSearchPost(context.Background(), indexName).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.SearchIndexV1RagManagedIndexesIndexNameSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIndexV1RagManagedIndexesIndexNameSearchPost`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.SearchIndexV1RagManagedIndexesIndexNameSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchIndexV1RagManagedIndexesIndexNameSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndexV1RagManagedIndexesIndexNamePut

> ManagedIndexResponse UpdateIndexV1RagManagedIndexesIndexNamePut(ctx, indexName).UpdateManagedIndexRequest(updateManagedIndexRequest).Execute()

Update a managed index schema



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
	indexName := "indexName_example" // string | 
	updateManagedIndexRequest := *openapiclient.NewUpdateManagedIndexRequest(*openapiclient.NewManagedIndexFields()) // UpdateManagedIndexRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagManagedIndexesAPI.UpdateIndexV1RagManagedIndexesIndexNamePut(context.Background(), indexName).UpdateManagedIndexRequest(updateManagedIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagManagedIndexesAPI.UpdateIndexV1RagManagedIndexesIndexNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIndexV1RagManagedIndexesIndexNamePut`: ManagedIndexResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagManagedIndexesAPI.UpdateIndexV1RagManagedIndexesIndexNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexV1RagManagedIndexesIndexNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateManagedIndexRequest** | [**UpdateManagedIndexRequest**](UpdateManagedIndexRequest.md) |  | 

### Return type

[**ManagedIndexResponse**](ManagedIndexResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


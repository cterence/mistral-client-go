# \BetaRagSearchIndexesAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeploymentSummariesV1RagDeploymentsGet**](BetaRagSearchIndexesAPI.md#GetDeploymentSummariesV1RagDeploymentsGet) | **Get** /v1/rag/deployments | Get Deployment Summaries
[**RegisterDeploymentV1RagDeploymentsPut**](BetaRagSearchIndexesAPI.md#RegisterDeploymentV1RagDeploymentsPut) | **Put** /v1/rag/deployments | Register (or re-register) a search index
[**UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete**](BetaRagSearchIndexesAPI.md#UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete) | **Delete** /v1/rag/deployments/{deployment_id} | Unregister Deployment
[**UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut**](BetaRagSearchIndexesAPI.md#UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut) | **Put** /v1/rag/deployments/{deployment_id}/metrics | Update Index Metrics



## GetDeploymentSummariesV1RagDeploymentsGet

> GetDeploymentSummariesResponse GetDeploymentSummariesV1RagDeploymentsGet(ctx).Execute()

Get Deployment Summaries



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
	resp, r, err := apiClient.BetaRagSearchIndexesAPI.GetDeploymentSummariesV1RagDeploymentsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagSearchIndexesAPI.GetDeploymentSummariesV1RagDeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentSummariesV1RagDeploymentsGet`: GetDeploymentSummariesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaRagSearchIndexesAPI.GetDeploymentSummariesV1RagDeploymentsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentSummariesV1RagDeploymentsGetRequest struct via the builder pattern


### Return type

[**GetDeploymentSummariesResponse**](GetDeploymentSummariesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDeploymentV1RagDeploymentsPut

> RegisterSearchIndexResponseIndex RegisterDeploymentV1RagDeploymentsPut(ctx).RegisterDeploymentRequestDeployment(registerDeploymentRequestDeployment).Execute()

Register (or re-register) a search index



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
	registerDeploymentRequestDeployment := *openapiclient.NewRegisterDeploymentRequestDeployment("Name_example", *openapiclient.NewRegisterDeploymentRequestVespaDeployment("VespaVersion_example", []openapiclient.RegisterDeploymentRequestVespaIndex{*openapiclient.NewRegisterDeploymentRequestVespaIndex("Name_example", []openapiclient.RegisterDeploymentRequestVespaField{*openapiclient.NewRegisterDeploymentRequestVespaField("Name_example", openapiclient.SchemaFieldDataType("int"), openapiclient.SchemaFieldStorage("in_memory"), openapiclient.SchemaFieldRankingType("count"), "TODO", false)}, "Sd_example")}, "QueryUrl_example")) // RegisterDeploymentRequestDeployment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagSearchIndexesAPI.RegisterDeploymentV1RagDeploymentsPut(context.Background()).RegisterDeploymentRequestDeployment(registerDeploymentRequestDeployment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagSearchIndexesAPI.RegisterDeploymentV1RagDeploymentsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterDeploymentV1RagDeploymentsPut`: RegisterSearchIndexResponseIndex
	fmt.Fprintf(os.Stdout, "Response from `BetaRagSearchIndexesAPI.RegisterDeploymentV1RagDeploymentsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDeploymentV1RagDeploymentsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerDeploymentRequestDeployment** | [**RegisterDeploymentRequestDeployment**](RegisterDeploymentRequestDeployment.md) |  | 

### Return type

[**RegisterSearchIndexResponseIndex**](RegisterSearchIndexResponseIndex.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete

> interface{} UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete(ctx, deploymentId).Execute()

Unregister Deployment



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagSearchIndexesAPI.UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagSearchIndexesAPI.UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BetaRagSearchIndexesAPI.UnregisterDeploymentV1RagDeploymentsDeploymentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterDeploymentV1RagDeploymentsDeploymentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut

> interface{} UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut(ctx, deploymentId).MetricsData(metricsData).Execute()

Update Index Metrics



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metricsData := *openapiclient.NewMetricsData("Status_example", int32(123), []openapiclient.UpdateMetricsRequestIndexMetrics{*openapiclient.NewUpdateMetricsRequestIndexMetrics("Name_example", int32(123))}) // MetricsData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaRagSearchIndexesAPI.UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut(context.Background(), deploymentId).MetricsData(metricsData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaRagSearchIndexesAPI.UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BetaRagSearchIndexesAPI.UpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexMetricsV1RagDeploymentsDeploymentIdMetricsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsData** | [**MetricsData**](MetricsData.md) |  | 

### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


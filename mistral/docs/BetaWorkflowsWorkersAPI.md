# \BetaWorkflowsWorkersAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkerInfoV1WorkflowsWorkersWhoamiGet**](BetaWorkflowsWorkersAPI.md#GetWorkerInfoV1WorkflowsWorkersWhoamiGet) | **Get** /v1/workflows/workers/whoami | Get Worker Info



## GetWorkerInfoV1WorkflowsWorkersWhoamiGet

> WorkerInfo GetWorkerInfoV1WorkflowsWorkersWhoamiGet(ctx).Execute()

Get Worker Info

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
	resp, r, err := apiClient.BetaWorkflowsWorkersAPI.GetWorkerInfoV1WorkflowsWorkersWhoamiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaWorkflowsWorkersAPI.GetWorkerInfoV1WorkflowsWorkersWhoamiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkerInfoV1WorkflowsWorkersWhoamiGet`: WorkerInfo
	fmt.Fprintf(os.Stdout, "Response from `BetaWorkflowsWorkersAPI.GetWorkerInfoV1WorkflowsWorkersWhoamiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkerInfoV1WorkflowsWorkersWhoamiGetRequest struct via the builder pattern


### Return type

[**WorkerInfo**](WorkerInfo.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


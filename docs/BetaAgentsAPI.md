# \BetaAgentsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsApiV1AgentsCreate**](BetaAgentsAPI.md#AgentsApiV1AgentsCreate) | **Post** /v1/agents | Create a agent that can be used within a conversation.
[**AgentsApiV1AgentsCreateOrUpdateAlias**](BetaAgentsAPI.md#AgentsApiV1AgentsCreateOrUpdateAlias) | **Put** /v1/agents/{agent_id}/aliases | Create or update an agent version alias.
[**AgentsApiV1AgentsDelete**](BetaAgentsAPI.md#AgentsApiV1AgentsDelete) | **Delete** /v1/agents/{agent_id} | Delete an agent entity.
[**AgentsApiV1AgentsDeleteAlias**](BetaAgentsAPI.md#AgentsApiV1AgentsDeleteAlias) | **Delete** /v1/agents/{agent_id}/aliases | Delete an agent version alias.
[**AgentsApiV1AgentsGet**](BetaAgentsAPI.md#AgentsApiV1AgentsGet) | **Get** /v1/agents/{agent_id} | Retrieve an agent entity.
[**AgentsApiV1AgentsGetVersion**](BetaAgentsAPI.md#AgentsApiV1AgentsGetVersion) | **Get** /v1/agents/{agent_id}/versions/{version} | Retrieve a specific version of an agent.
[**AgentsApiV1AgentsList**](BetaAgentsAPI.md#AgentsApiV1AgentsList) | **Get** /v1/agents | List agent entities.
[**AgentsApiV1AgentsListPages**](BetaAgentsAPI.md#AgentsApiV1AgentsListPages) | **Get** /v1/agents/pages | List agent entities, cursor-paginated.
[**AgentsApiV1AgentsListVersionAliases**](BetaAgentsAPI.md#AgentsApiV1AgentsListVersionAliases) | **Get** /v1/agents/{agent_id}/aliases | List all aliases for an agent.
[**AgentsApiV1AgentsListVersions**](BetaAgentsAPI.md#AgentsApiV1AgentsListVersions) | **Get** /v1/agents/{agent_id}/versions | List all versions of an agent.
[**AgentsApiV1AgentsUpdate**](BetaAgentsAPI.md#AgentsApiV1AgentsUpdate) | **Patch** /v1/agents/{agent_id} | Update an agent entity.
[**AgentsApiV1AgentsUpdateVersion**](BetaAgentsAPI.md#AgentsApiV1AgentsUpdateVersion) | **Patch** /v1/agents/{agent_id}/version | Update an agent version.



## AgentsApiV1AgentsCreate

> Agent AgentsApiV1AgentsCreate(ctx).AgentCreationRequest(agentCreationRequest).Execute()

Create a agent that can be used within a conversation.



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
	agentCreationRequest := *openapiclient.NewAgentCreationRequest("Model_example", "Name_example") // AgentCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsCreate(context.Background()).AgentCreationRequest(agentCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsCreate`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentCreationRequest** | [**AgentCreationRequest**](AgentCreationRequest.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsCreateOrUpdateAlias

> AgentAliasResponse AgentsApiV1AgentsCreateOrUpdateAlias(ctx, agentId).Alias(alias).Version(version).Execute()

Create or update an agent version alias.



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
	agentId := "agentId_example" // string | 
	alias := "alias_example" // string | 
	version := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsCreateOrUpdateAlias(context.Background(), agentId).Alias(alias).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsCreateOrUpdateAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsCreateOrUpdateAlias`: AgentAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsCreateOrUpdateAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsCreateOrUpdateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alias** | **string** |  | 
 **version** | **int32** |  | 

### Return type

[**AgentAliasResponse**](AgentAliasResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsDelete

> AgentsApiV1AgentsDelete(ctx, agentId).Execute()

Delete an agent entity.

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
	agentId := "agentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsDelete(context.Background(), agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsDeleteRequest struct via the builder pattern


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


## AgentsApiV1AgentsDeleteAlias

> AgentsApiV1AgentsDeleteAlias(ctx, agentId).Alias(alias).Execute()

Delete an agent version alias.



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
	agentId := "agentId_example" // string | 
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsDeleteAlias(context.Background(), agentId).Alias(alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsDeleteAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsDeleteAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alias** | **string** |  | 

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


## AgentsApiV1AgentsGet

> Agent AgentsApiV1AgentsGet(ctx, agentId).AgentVersion(agentVersion).Execute()

Retrieve an agent entity.



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
	agentId := "agentId_example" // string | 
	agentVersion := *openapiclient.NewAgentVersion() // AgentVersion |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsGet(context.Background(), agentId).AgentVersion(agentVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsGet`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentVersion** | [**AgentVersion**](AgentVersion.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsGetVersion

> Agent AgentsApiV1AgentsGetVersion(ctx, agentId, version).Execute()

Retrieve a specific version of an agent.



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
	agentId := "agentId_example" // string | 
	version := "version_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsGetVersion(context.Background(), agentId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsGetVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsGetVersion`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsGetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsList

> []Agent AgentsApiV1AgentsList(ctx).Page(page).PageSize(pageSize).DeploymentChat(deploymentChat).Sources(sources).Name(name).Search(search).Id(id).Metadata(metadata).Execute()

List agent entities.



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
	page := int32(56) // int32 | Page number (0-indexed) (optional) (default to 0)
	pageSize := int32(56) // int32 | Number of agents per page (optional) (default to 20)
	deploymentChat := true // bool |  (optional)
	sources := []openapiclient.RequestSource{openapiclient.RequestSource("api")} // []RequestSource |  (optional)
	name := "name_example" // string | Filter by agent name (optional)
	search := "search_example" // string | Search agents by name or ID (optional)
	id := "id_example" // string |  (optional)
	metadata := TODO // AnyOfmapnull |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsList(context.Background()).Page(page).PageSize(pageSize).DeploymentChat(deploymentChat).Sources(sources).Name(name).Search(search).Id(id).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsList`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (0-indexed) | [default to 0]
 **pageSize** | **int32** | Number of agents per page | [default to 20]
 **deploymentChat** | **bool** |  | 
 **sources** | [**[]RequestSource**](RequestSource.md) |  | 
 **name** | **string** | Filter by agent name | 
 **search** | **string** | Search agents by name or ID | 
 **id** | **string** |  | 
 **metadata** | [**AnyOfmapnull**](AnyOfmapnull.md) |  | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsListPages

> AgentListPage AgentsApiV1AgentsListPages(ctx).PageSize(pageSize).DeploymentChat(deploymentChat).Sources(sources).Name(name).Search(search).Id(id).Metadata(metadata).PageToken(pageToken).Execute()

List agent entities, cursor-paginated.



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
	pageSize := int32(56) // int32 | Number of agents per page (optional) (default to 20)
	deploymentChat := true // bool |  (optional)
	sources := []openapiclient.RequestSource{openapiclient.RequestSource("api")} // []RequestSource |  (optional)
	name := "name_example" // string | Filter by agent name (optional)
	search := "search_example" // string | Search agents by name or ID (optional)
	id := "id_example" // string |  (optional)
	metadata := TODO // AnyOfmapnull |  (optional)
	pageToken := "pageToken_example" // string | Opaque cursor from a previous response's next_page_token. When set, results page forward from the cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsListPages(context.Background()).PageSize(pageSize).DeploymentChat(deploymentChat).Sources(sources).Name(name).Search(search).Id(id).Metadata(metadata).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsListPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsListPages`: AgentListPage
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsListPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsListPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of agents per page | [default to 20]
 **deploymentChat** | **bool** |  | 
 **sources** | [**[]RequestSource**](RequestSource.md) |  | 
 **name** | **string** | Filter by agent name | 
 **search** | **string** | Search agents by name or ID | 
 **id** | **string** |  | 
 **metadata** | [**AnyOfmapnull**](AnyOfmapnull.md) |  | 
 **pageToken** | **string** | Opaque cursor from a previous response&#39;s next_page_token. When set, results page forward from the cursor. | 

### Return type

[**AgentListPage**](AgentListPage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsListVersionAliases

> []AgentAliasResponse AgentsApiV1AgentsListVersionAliases(ctx, agentId).Execute()

List all aliases for an agent.



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
	agentId := "agentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsListVersionAliases(context.Background(), agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsListVersionAliases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsListVersionAliases`: []AgentAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsListVersionAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsListVersionAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AgentAliasResponse**](AgentAliasResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsListVersions

> []Agent AgentsApiV1AgentsListVersions(ctx, agentId).Page(page).PageSize(pageSize).Execute()

List all versions of an agent.



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
	agentId := "agentId_example" // string | 
	page := int32(56) // int32 | Page number (0-indexed) (optional) (default to 0)
	pageSize := int32(56) // int32 | Number of versions per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsListVersions(context.Background(), agentId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsListVersions`: []Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (0-indexed) | [default to 0]
 **pageSize** | **int32** | Number of versions per page | [default to 20]

### Return type

[**[]Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsUpdate

> Agent AgentsApiV1AgentsUpdate(ctx, agentId).AgentUpdateRequest(agentUpdateRequest).Execute()

Update an agent entity.



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
	agentId := "agentId_example" // string | 
	agentUpdateRequest := *openapiclient.NewAgentUpdateRequest() // AgentUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsUpdate(context.Background(), agentId).AgentUpdateRequest(agentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsUpdate`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentUpdateRequest** | [**AgentUpdateRequest**](AgentUpdateRequest.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsApiV1AgentsUpdateVersion

> Agent AgentsApiV1AgentsUpdateVersion(ctx, agentId).Version(version).Execute()

Update an agent version.



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
	agentId := "agentId_example" // string | 
	version := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAgentsAPI.AgentsApiV1AgentsUpdateVersion(context.Background(), agentId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAgentsAPI.AgentsApiV1AgentsUpdateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsApiV1AgentsUpdateVersion`: Agent
	fmt.Fprintf(os.Stdout, "Response from `BetaAgentsAPI.AgentsApiV1AgentsUpdateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsApiV1AgentsUpdateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


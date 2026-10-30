# \BetaConnectorsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectorActivateForConsumerV1**](BetaConnectorsAPI.md#ConnectorActivateForConsumerV1) | **Post** /v1/connectors/{connector_id}/{consumer_scope}/activate | Activate a connector for the given consumer (organization, workspace, user).
[**ConnectorCallToolV1**](BetaConnectorsAPI.md#ConnectorCallToolV1) | **Post** /v1/connectors/{connector_id_or_name}/tools/{tool_name}/call | Call Connector Tool
[**ConnectorCreateCredentialsV1**](BetaConnectorsAPI.md#ConnectorCreateCredentialsV1) | **Post** /v1/connectors/{connector_id_or_name}/{consumer_scope}/credentials | Create consumer credentials for a connector.
[**ConnectorCreateV1**](BetaConnectorsAPI.md#ConnectorCreateV1) | **Post** /v1/connectors | Create a new connector.
[**ConnectorDeactivateForConsumerV1**](BetaConnectorsAPI.md#ConnectorDeactivateForConsumerV1) | **Post** /v1/connectors/{connector_id}/{consumer_scope}/deactivate | Deactivate a connector for the current consumer (at organization, workspace or user level).
[**ConnectorDeleteAllUserCredentialsV1**](BetaConnectorsAPI.md#ConnectorDeleteAllUserCredentialsV1) | **Delete** /v1/connectors/{connector_id_or_name}/user/credentials | Delete all user credentials for a connector.
[**ConnectorDeleteCredentials**](BetaConnectorsAPI.md#ConnectorDeleteCredentials) | **Delete** /v1/connectors/{connector_id_or_name}/{consumer_scope}/credentials/{credentials_name} | Delete credentials for a consumer.
[**ConnectorDeleteV1**](BetaConnectorsAPI.md#ConnectorDeleteV1) | **Delete** /v1/connectors/{connector_id}#id | Delete a connector.
[**ConnectorGetAuthUrlV1**](BetaConnectorsAPI.md#ConnectorGetAuthUrlV1) | **Get** /v1/connectors/{connector_id_or_name}/auth_url | Get the auth URL for a connector.
[**ConnectorGetAuthenticationMethodsV1**](BetaConnectorsAPI.md#ConnectorGetAuthenticationMethodsV1) | **Get** /v1/connectors/{connector_id_or_name}/authentication_methods | Get authentication methods for a connector.
[**ConnectorGetV1**](BetaConnectorsAPI.md#ConnectorGetV1) | **Get** /v1/connectors/{connector_id_or_name}#idOrName | Get a connector.
[**ConnectorListOrganizationCredentialsV1**](BetaConnectorsAPI.md#ConnectorListOrganizationCredentialsV1) | **Get** /v1/connectors/{connector_id_or_name}/organization/credentials | List organization credentials for a connector.
[**ConnectorListToolsV1**](BetaConnectorsAPI.md#ConnectorListToolsV1) | **Get** /v1/connectors/{connector_id_or_name}/tools | List tools for a connector.
[**ConnectorListUserCredentialsV1**](BetaConnectorsAPI.md#ConnectorListUserCredentialsV1) | **Get** /v1/connectors/{connector_id_or_name}/user/credentials | List user credentials for a connector.
[**ConnectorListV1**](BetaConnectorsAPI.md#ConnectorListV1) | **Get** /v1/connectors | List all connectors.
[**ConnectorListWorkspaceCredentialsV1**](BetaConnectorsAPI.md#ConnectorListWorkspaceCredentialsV1) | **Get** /v1/connectors/{connector_id_or_name}/workspace/credentials | List workspace credentials for a connector.
[**ConnectorShareToOrganizationV1**](BetaConnectorsAPI.md#ConnectorShareToOrganizationV1) | **Put** /v1/connectors/{connector_id}/organization/share | Share a connector to the current organization.
[**ConnectorShareV1**](BetaConnectorsAPI.md#ConnectorShareV1) | **Put** /v1/connectors/{connector_id}/share | Share a private connector to the current workspace.
[**ConnectorUnshareFromOrganizationV1**](BetaConnectorsAPI.md#ConnectorUnshareFromOrganizationV1) | **Delete** /v1/connectors/{connector_id}/organization/share | Unshare a connector from the current organization.
[**ConnectorUnshareV1**](BetaConnectorsAPI.md#ConnectorUnshareV1) | **Delete** /v1/connectors/{connector_id}/share | Unshare a connector from the current workspace.
[**ConnectorUpdateCredentials**](BetaConnectorsAPI.md#ConnectorUpdateCredentials) | **Patch** /v1/connectors/{connector_id_or_name}/{consumer_scope}/credentials | Create or update consumer credentials for a connector.
[**ConnectorUpdateV1**](BetaConnectorsAPI.md#ConnectorUpdateV1) | **Patch** /v1/connectors/{connector_id}#id | Update a connector.



## ConnectorActivateForConsumerV1

> MessageResponse ConnectorActivateForConsumerV1(ctx, connectorId, consumerScope).Execute()

Activate a connector for the given consumer (organization, workspace, user).



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	consumerScope := "consumerScope_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorActivateForConsumerV1(context.Background(), connectorId, consumerScope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorActivateForConsumerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorActivateForConsumerV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorActivateForConsumerV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 
**consumerScope** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorActivateForConsumerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorCallToolV1

> ConnectorToolCallResponse ConnectorCallToolV1(ctx, toolName, connectorIdOrName).ConnectorCallToolRequest(connectorCallToolRequest).CredentialsName(credentialsName).Execute()

Call Connector Tool



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
	toolName := "toolName_example" // string | 
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	connectorCallToolRequest := *openapiclient.NewConnectorCallToolRequest() // ConnectorCallToolRequest | 
	credentialsName := "credentialsName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCallToolV1(context.Background(), toolName, connectorIdOrName).ConnectorCallToolRequest(connectorCallToolRequest).CredentialsName(credentialsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCallToolV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCallToolV1`: ConnectorToolCallResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCallToolV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolName** | **string** |  | 
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCallToolV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectorCallToolRequest** | [**ConnectorCallToolRequest**](ConnectorCallToolRequest.md) |  | 
 **credentialsName** | **string** |  | 

### Return type

[**ConnectorToolCallResponse**](ConnectorToolCallResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorCreateCredentialsV1

> MessageResponse ConnectorCreateCredentialsV1(ctx, connectorIdOrName, consumerScope).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()

Create consumer credentials for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	consumerScope := "consumerScope_example" // string | 
	credentialsCreateOrUpdate := *openapiclient.NewCredentialsCreateOrUpdate("Name_example") // CredentialsCreateOrUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCreateCredentialsV1(context.Background(), connectorIdOrName, consumerScope).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCreateCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCreateCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCreateCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 
**consumerScope** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCreateCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialsCreateOrUpdate** | [**CredentialsCreateOrUpdate**](CredentialsCreateOrUpdate.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorCreateV1

> Connector ConnectorCreateV1(ctx).Payload(payload).Execute()

Create a new connector.



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
	payload := openapiclient.Payload{CreateConnectorRequest: openapiclient.NewCreateConnectorRequest("Name_example", "Description_example", "Server_example")} // Payload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCreateV1(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCreateV1`: Connector
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**Payload**](Payload.md) |  | 

### Return type

[**Connector**](Connector.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorDeactivateForConsumerV1

> MessageResponse ConnectorDeactivateForConsumerV1(ctx, connectorId, consumerScope).Execute()

Deactivate a connector for the current consumer (at organization, workspace or user level).



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	consumerScope := "consumerScope_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeactivateForConsumerV1(context.Background(), connectorId, consumerScope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeactivateForConsumerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeactivateForConsumerV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeactivateForConsumerV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 
**consumerScope** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeactivateForConsumerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorDeleteAllUserCredentialsV1

> MessageResponse ConnectorDeleteAllUserCredentialsV1(ctx, connectorIdOrName).Execute()

Delete all user credentials for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeleteAllUserCredentialsV1(context.Background(), connectorIdOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeleteAllUserCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeleteAllUserCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeleteAllUserCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeleteAllUserCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorDeleteCredentials

> MessageResponse ConnectorDeleteCredentials(ctx, credentialsName, connectorIdOrName, consumerScope).Execute()

Delete credentials for a consumer.



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
	credentialsName := "credentialsName_example" // string | 
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	consumerScope := "consumerScope_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeleteCredentials(context.Background(), credentialsName, connectorIdOrName, consumerScope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeleteCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeleteCredentials`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeleteCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsName** | **string** |  | 
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 
**consumerScope** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeleteCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorDeleteV1

> MessageResponse ConnectorDeleteV1(ctx, connectorId).Execute()

Delete a connector.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeleteV1(context.Background(), connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeleteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeleteV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorGetAuthUrlV1

> AuthUrlResponse ConnectorGetAuthUrlV1(ctx, connectorIdOrName).AppReturnUrl(appReturnUrl).MethodType(methodType).CredentialsName(credentialsName).CredentialsTitle(credentialsTitle).GithubInstallationLink(githubInstallationLink).Execute()

Get the auth URL for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName() // ConnectorIdOrName | 
	appReturnUrl := "appReturnUrl_example" // string |  (optional)
	methodType := openapiclient.AuthenticationType("oauth2") // AuthenticationType | Auth method type to use for the authorization URL. Required when the connector supports multiple interactive auth methods; otherwise the sole method is selected automatically. Use this to pick a specific method (e.g. 'oauth2' vs 'github_app'). (optional) (default to "oauth2")
	credentialsName := "credentialsName_example" // string |  (optional)
	credentialsTitle := "credentialsTitle_example" // string |  (optional)
	githubInstallationLink := true // bool | Only valid with method_type=oauth2. When true, returns a GitHub App installation URL (https://github.com/apps/<slug>/installations/new) if the connector has the proper configuration The Github application needs to have 'Request user authorization (OAuth) during installation' enabled to perform the proper auth loop. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorGetAuthUrlV1(context.Background(), connectorIdOrName).AppReturnUrl(appReturnUrl).MethodType(methodType).CredentialsName(credentialsName).CredentialsTitle(credentialsTitle).GithubInstallationLink(githubInstallationLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorGetAuthUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorGetAuthUrlV1`: AuthUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorGetAuthUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorGetAuthUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appReturnUrl** | **string** |  | 
 **methodType** | [**AuthenticationType**](AuthenticationType.md) | Auth method type to use for the authorization URL. Required when the connector supports multiple interactive auth methods; otherwise the sole method is selected automatically. Use this to pick a specific method (e.g. &#39;oauth2&#39; vs &#39;github_app&#39;). | [default to &quot;oauth2&quot;]
 **credentialsName** | **string** |  | 
 **credentialsTitle** | **string** |  | 
 **githubInstallationLink** | **bool** | Only valid with method_type&#x3D;oauth2. When true, returns a GitHub App installation URL (https://github.com/apps/&lt;slug&gt;/installations/new) if the connector has the proper configuration The Github application needs to have &#39;Request user authorization (OAuth) during installation&#39; enabled to perform the proper auth loop. | [default to false]

### Return type

[**AuthUrlResponse**](AuthUrlResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorGetAuthenticationMethodsV1

> []PublicAuthenticationMethod ConnectorGetAuthenticationMethodsV1(ctx, connectorIdOrName).Execute()

Get authentication methods for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorGetAuthenticationMethodsV1(context.Background(), connectorIdOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorGetAuthenticationMethodsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorGetAuthenticationMethodsV1`: []PublicAuthenticationMethod
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorGetAuthenticationMethodsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorGetAuthenticationMethodsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PublicAuthenticationMethod**](PublicAuthenticationMethod.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorGetV1

> Connector ConnectorGetV1(ctx, connectorIdOrName).FetchUserData(fetchUserData).Execute()

Get a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	fetchUserData := true // bool | Fetch the user-level data associated with the connector (e.g. connection credentials). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorGetV1(context.Background(), connectorIdOrName).FetchUserData(fetchUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorGetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorGetV1`: Connector
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorGetV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorGetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchUserData** | **bool** | Fetch the user-level data associated with the connector (e.g. connection credentials). | [default to false]

### Return type

[**Connector**](Connector.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorListOrganizationCredentialsV1

> CredentialsResponse ConnectorListOrganizationCredentialsV1(ctx, connectorIdOrName).AuthType(authType).FetchDefault(fetchDefault).Execute()

List organization credentials for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	authType := openapiclient.AuthenticationType("oauth2") // AuthenticationType |  (optional)
	fetchDefault := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorListOrganizationCredentialsV1(context.Background(), connectorIdOrName).AuthType(authType).FetchDefault(fetchDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorListOrganizationCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorListOrganizationCredentialsV1`: CredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorListOrganizationCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListOrganizationCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | [**AuthenticationType**](AuthenticationType.md) |  | 
 **fetchDefault** | **bool** |  | [default to false]

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorListToolsV1

> ResponseConnectorListToolsV1 ConnectorListToolsV1(ctx, connectorIdOrName).PageSize(pageSize).Refresh(refresh).Pretty(pretty).CredentialsName(credentialsName).Execute()

List tools for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	refresh := true // bool |  (optional) (default to false)
	pretty := true // bool | Return a simplified payload with only name, description, annotations, and a compact inputSchema. (optional) (default to false)
	credentialsName := "credentialsName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorListToolsV1(context.Background(), connectorIdOrName).PageSize(pageSize).Refresh(refresh).Pretty(pretty).CredentialsName(credentialsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorListToolsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorListToolsV1`: ResponseConnectorListToolsV1
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorListToolsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListToolsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | [default to 100]
 **refresh** | **bool** |  | [default to false]
 **pretty** | **bool** | Return a simplified payload with only name, description, annotations, and a compact inputSchema. | [default to false]
 **credentialsName** | **string** |  | 

### Return type

[**ResponseConnectorListToolsV1**](ResponseConnectorListToolsV1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorListUserCredentialsV1

> CredentialsResponse ConnectorListUserCredentialsV1(ctx, connectorIdOrName).AuthType(authType).FetchDefault(fetchDefault).Execute()

List user credentials for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	authType := openapiclient.AuthenticationType("oauth2") // AuthenticationType |  (optional)
	fetchDefault := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorListUserCredentialsV1(context.Background(), connectorIdOrName).AuthType(authType).FetchDefault(fetchDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorListUserCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorListUserCredentialsV1`: CredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorListUserCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListUserCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | [**AuthenticationType**](AuthenticationType.md) |  | 
 **fetchDefault** | **bool** |  | [default to false]

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorListV1

> PaginatedConnectors ConnectorListV1(ctx).QueryFilters(queryFilters).Cursor(cursor).PageSize(pageSize).Execute()

List all connectors.



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
	queryFilters := *openapiclient.NewConnectorsQueryFilters() // ConnectorsQueryFilters |  (optional) (default to {"fetch_user_data":false,"check_credentials_status":false})
	cursor := "cursor_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorListV1(context.Background()).QueryFilters(queryFilters).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorListV1`: PaginatedConnectors
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryFilters** | [**ConnectorsQueryFilters**](ConnectorsQueryFilters.md) |  | [default to {&quot;fetch_user_data&quot;:false,&quot;check_credentials_status&quot;:false}]
 **cursor** | **string** |  | 
 **pageSize** | **int32** |  | [default to 100]

### Return type

[**PaginatedConnectors**](PaginatedConnectors.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorListWorkspaceCredentialsV1

> CredentialsResponse ConnectorListWorkspaceCredentialsV1(ctx, connectorIdOrName).AuthType(authType).FetchDefault(fetchDefault).Execute()

List workspace credentials for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	authType := openapiclient.AuthenticationType("oauth2") // AuthenticationType |  (optional)
	fetchDefault := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorListWorkspaceCredentialsV1(context.Background(), connectorIdOrName).AuthType(authType).FetchDefault(fetchDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorListWorkspaceCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorListWorkspaceCredentialsV1`: CredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorListWorkspaceCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListWorkspaceCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | [**AuthenticationType**](AuthenticationType.md) |  | 
 **fetchDefault** | **bool** |  | [default to false]

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorShareToOrganizationV1

> MessageResponse ConnectorShareToOrganizationV1(ctx, connectorId).Execute()

Share a connector to the current organization.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorShareToOrganizationV1(context.Background(), connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorShareToOrganizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorShareToOrganizationV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorShareToOrganizationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorShareToOrganizationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorShareV1

> MessageResponse ConnectorShareV1(ctx, connectorId).Execute()

Share a private connector to the current workspace.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorShareV1(context.Background(), connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorShareV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorShareV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorShareV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorShareV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorUnshareFromOrganizationV1

> MessageResponse ConnectorUnshareFromOrganizationV1(ctx, connectorId).Execute()

Unshare a connector from the current organization.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorUnshareFromOrganizationV1(context.Background(), connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorUnshareFromOrganizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorUnshareFromOrganizationV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorUnshareFromOrganizationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorUnshareFromOrganizationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorUnshareV1

> MessageResponse ConnectorUnshareV1(ctx, connectorId).Execute()

Unshare a connector from the current workspace.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorUnshareV1(context.Background(), connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorUnshareV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorUnshareV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorUnshareV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorUnshareV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorUpdateCredentials

> MessageResponse ConnectorUpdateCredentials(ctx, connectorIdOrName, consumerScope).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()

Create or update consumer credentials for a connector.



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
	connectorIdOrName := *openapiclient.NewConnectorIdOrName1() // ConnectorIdOrName1 | 
	consumerScope := "consumerScope_example" // string | 
	credentialsCreateOrUpdate := *openapiclient.NewCredentialsCreateOrUpdate("Name_example") // CredentialsCreateOrUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorUpdateCredentials(context.Background(), connectorIdOrName, consumerScope).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorUpdateCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorUpdateCredentials`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorUpdateCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | [**ConnectorIdOrName1**](.md) |  | 
**consumerScope** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorUpdateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialsCreateOrUpdate** | [**CredentialsCreateOrUpdate**](CredentialsCreateOrUpdate.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorUpdateV1

> Connector ConnectorUpdateV1(ctx, connectorId).Payload1(payload1).Execute()

Update a connector.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	payload1 := openapiclient.Payload_1{ConnectorMCPPublicUpdate: openapiclient.NewConnectorMCPPublicUpdate()} // Payload1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorUpdateV1(context.Background(), connectorId).Payload1(payload1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorUpdateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorUpdateV1`: Connector
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorUpdateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorUpdateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload1** | [**Payload1**](Payload1.md) |  | 

### Return type

[**Connector**](Connector.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


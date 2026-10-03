# \BetaConnectorsAPI

All URIs are relative to *https://api.mistral.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectorCallToolV1**](BetaConnectorsAPI.md#ConnectorCallToolV1) | **Post** /v1/connectors/{connector_id_or_name}/tools/{tool_name}/call | Call Connector Tool
[**ConnectorCreateOrUpdateOrganizationCredentialsV1**](BetaConnectorsAPI.md#ConnectorCreateOrUpdateOrganizationCredentialsV1) | **Post** /v1/connectors/{connector_id_or_name}/organization/credentials | Create or update organization credentials for a connector.
[**ConnectorCreateOrUpdateUserCredentialsV1**](BetaConnectorsAPI.md#ConnectorCreateOrUpdateUserCredentialsV1) | **Post** /v1/connectors/{connector_id_or_name}/user/credentials | Create or update user credentials for a connector.
[**ConnectorCreateOrUpdateWorkspaceCredentialsV1**](BetaConnectorsAPI.md#ConnectorCreateOrUpdateWorkspaceCredentialsV1) | **Post** /v1/connectors/{connector_id_or_name}/workspace/credentials | Create or update workspace credentials for a connector.
[**ConnectorCreateV1**](BetaConnectorsAPI.md#ConnectorCreateV1) | **Post** /v1/connectors | Create a new connector.
[**ConnectorDeleteOrganizationCredentialsV1**](BetaConnectorsAPI.md#ConnectorDeleteOrganizationCredentialsV1) | **Delete** /v1/connectors/{connector_id_or_name}/organization/credentials/{credentials_name} | Delete organization credentials for a connector.
[**ConnectorDeleteUserCredentialsV1**](BetaConnectorsAPI.md#ConnectorDeleteUserCredentialsV1) | **Delete** /v1/connectors/{connector_id_or_name}/user/credentials/{credentials_name} | Delete user credentials for a connector.
[**ConnectorDeleteV1**](BetaConnectorsAPI.md#ConnectorDeleteV1) | **Delete** /v1/connectors/{connector_id}#id | Delete a connector.
[**ConnectorDeleteWorkspaceCredentialsV1**](BetaConnectorsAPI.md#ConnectorDeleteWorkspaceCredentialsV1) | **Delete** /v1/connectors/{connector_id_or_name}/workspace/credentials/{credentials_name} | Delete workspace credentials for a connector.
[**ConnectorGetAuthUrlV1**](BetaConnectorsAPI.md#ConnectorGetAuthUrlV1) | **Get** /v1/connectors/{connector_id_or_name}/auth_url | Get the auth URL for a connector.
[**ConnectorGetAuthenticationMethodsV1**](BetaConnectorsAPI.md#ConnectorGetAuthenticationMethodsV1) | **Get** /v1/connectors/{connector_id_or_name}/authentication_methods | Get authentication methods for a connector.
[**ConnectorGetV1**](BetaConnectorsAPI.md#ConnectorGetV1) | **Get** /v1/connectors/{connector_id_or_name}#idOrName | Get a connector.
[**ConnectorListOrganizationCredentialsV1**](BetaConnectorsAPI.md#ConnectorListOrganizationCredentialsV1) | **Get** /v1/connectors/{connector_id_or_name}/organization/credentials | List organization credentials for a connector.
[**ConnectorListToolsV1**](BetaConnectorsAPI.md#ConnectorListToolsV1) | **Get** /v1/connectors/{connector_id_or_name}/tools | List tools for a connector.
[**ConnectorListUserCredentialsV1**](BetaConnectorsAPI.md#ConnectorListUserCredentialsV1) | **Get** /v1/connectors/{connector_id_or_name}/user/credentials | List user credentials for a connector.
[**ConnectorListV1**](BetaConnectorsAPI.md#ConnectorListV1) | **Get** /v1/connectors | List all connectors.
[**ConnectorListWorkspaceCredentialsV1**](BetaConnectorsAPI.md#ConnectorListWorkspaceCredentialsV1) | **Get** /v1/connectors/{connector_id_or_name}/workspace/credentials | List workspace credentials for a connector.
[**ConnectorUpdateV1**](BetaConnectorsAPI.md#ConnectorUpdateV1) | **Patch** /v1/connectors/{connector_id}#id | Update a connector.



## ConnectorCallToolV1

> MCPToolCallResponse ConnectorCallToolV1(ctx, toolName, connectorIdOrName).MCPToolCallRequest(mCPToolCallRequest).CredentialsName(credentialsName).Execute()

Call Connector Tool



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
	toolName := "toolName_example" // string | 
	connectorIdOrName := "connectorIdOrName_example" // string | 
	mCPToolCallRequest := *openapiclient.NewMCPToolCallRequest() // MCPToolCallRequest | 
	credentialsName := "credentialsName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCallToolV1(context.Background(), toolName, connectorIdOrName).MCPToolCallRequest(mCPToolCallRequest).CredentialsName(credentialsName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCallToolV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCallToolV1`: MCPToolCallResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCallToolV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolName** | **string** |  | 
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCallToolV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mCPToolCallRequest** | [**MCPToolCallRequest**](MCPToolCallRequest.md) |  | 
 **credentialsName** | **string** |  | 

### Return type

[**MCPToolCallResponse**](MCPToolCallResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorCreateOrUpdateOrganizationCredentialsV1

> MessageResponse ConnectorCreateOrUpdateOrganizationCredentialsV1(ctx, connectorIdOrName).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()

Create or update organization credentials for a connector.



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
	connectorIdOrName := "connectorIdOrName_example" // string | 
	credentialsCreateOrUpdate := *openapiclient.NewCredentialsCreateOrUpdate("Name_example") // CredentialsCreateOrUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCreateOrUpdateOrganizationCredentialsV1(context.Background(), connectorIdOrName).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCreateOrUpdateOrganizationCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCreateOrUpdateOrganizationCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCreateOrUpdateOrganizationCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCreateOrUpdateOrganizationCredentialsV1Request struct via the builder pattern


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


## ConnectorCreateOrUpdateUserCredentialsV1

> MessageResponse ConnectorCreateOrUpdateUserCredentialsV1(ctx, connectorIdOrName).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()

Create or update user credentials for a connector.



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
	connectorIdOrName := "connectorIdOrName_example" // string | 
	credentialsCreateOrUpdate := *openapiclient.NewCredentialsCreateOrUpdate("Name_example") // CredentialsCreateOrUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCreateOrUpdateUserCredentialsV1(context.Background(), connectorIdOrName).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCreateOrUpdateUserCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCreateOrUpdateUserCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCreateOrUpdateUserCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCreateOrUpdateUserCredentialsV1Request struct via the builder pattern


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


## ConnectorCreateOrUpdateWorkspaceCredentialsV1

> MessageResponse ConnectorCreateOrUpdateWorkspaceCredentialsV1(ctx, connectorIdOrName).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()

Create or update workspace credentials for a connector.



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
	connectorIdOrName := "connectorIdOrName_example" // string | 
	credentialsCreateOrUpdate := *openapiclient.NewCredentialsCreateOrUpdate("Name_example") // CredentialsCreateOrUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCreateOrUpdateWorkspaceCredentialsV1(context.Background(), connectorIdOrName).CredentialsCreateOrUpdate(credentialsCreateOrUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorCreateOrUpdateWorkspaceCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorCreateOrUpdateWorkspaceCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorCreateOrUpdateWorkspaceCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorCreateOrUpdateWorkspaceCredentialsV1Request struct via the builder pattern


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

> Connector ConnectorCreateV1(ctx).ConnectorMCPCreate(connectorMCPCreate).Execute()

Create a new connector.



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
	connectorMCPCreate := *openapiclient.NewConnectorMCPCreate("Name_example", "Description_example", "Server_example") // ConnectorMCPCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorCreateV1(context.Background()).ConnectorMCPCreate(connectorMCPCreate).Execute()
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
 **connectorMCPCreate** | [**ConnectorMCPCreate**](ConnectorMCPCreate.md) |  | 

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


## ConnectorDeleteOrganizationCredentialsV1

> MessageResponse ConnectorDeleteOrganizationCredentialsV1(ctx, credentialsName, connectorIdOrName).Execute()

Delete organization credentials for a connector.



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
	credentialsName := "credentialsName_example" // string | 
	connectorIdOrName := "connectorIdOrName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeleteOrganizationCredentialsV1(context.Background(), credentialsName, connectorIdOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeleteOrganizationCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeleteOrganizationCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeleteOrganizationCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsName** | **string** |  | 
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeleteOrganizationCredentialsV1Request struct via the builder pattern


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


## ConnectorDeleteUserCredentialsV1

> MessageResponse ConnectorDeleteUserCredentialsV1(ctx, credentialsName, connectorIdOrName).Execute()

Delete user credentials for a connector.



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
	credentialsName := "credentialsName_example" // string | 
	connectorIdOrName := "connectorIdOrName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeleteUserCredentialsV1(context.Background(), credentialsName, connectorIdOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeleteUserCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeleteUserCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeleteUserCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsName** | **string** |  | 
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeleteUserCredentialsV1Request struct via the builder pattern


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
	openapiclient "github.com/cterence/mistral-client-golang"
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


## ConnectorDeleteWorkspaceCredentialsV1

> MessageResponse ConnectorDeleteWorkspaceCredentialsV1(ctx, credentialsName, connectorIdOrName).Execute()

Delete workspace credentials for a connector.



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
	credentialsName := "credentialsName_example" // string | 
	connectorIdOrName := "connectorIdOrName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorDeleteWorkspaceCredentialsV1(context.Background(), credentialsName, connectorIdOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaConnectorsAPI.ConnectorDeleteWorkspaceCredentialsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorDeleteWorkspaceCredentialsV1`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaConnectorsAPI.ConnectorDeleteWorkspaceCredentialsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsName** | **string** |  | 
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorDeleteWorkspaceCredentialsV1Request struct via the builder pattern


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

> AuthUrlResponse ConnectorGetAuthUrlV1(ctx, connectorIdOrName).AppReturnUrl(appReturnUrl).CredentialsName(credentialsName).Execute()

Get the auth URL for a connector.



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
	connectorIdOrName := "connectorIdOrName_example" // string | 
	appReturnUrl := "appReturnUrl_example" // string |  (optional)
	credentialsName := "credentialsName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorGetAuthUrlV1(context.Background(), connectorIdOrName).AppReturnUrl(appReturnUrl).CredentialsName(credentialsName).Execute()
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
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorGetAuthUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appReturnUrl** | **string** |  | 
 **credentialsName** | **string** |  | 

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
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	connectorIdOrName := "connectorIdOrName_example" // string | 

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
**connectorIdOrName** | **string** |  | 

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

> Connector ConnectorGetV1(ctx, connectorIdOrName).FetchCustomerData(fetchCustomerData).FetchConnectionSecrets(fetchConnectionSecrets).Execute()

Get a connector.



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
	connectorIdOrName := "connectorIdOrName_example" // string | 
	fetchCustomerData := true // bool | Fetch the customer data associated with the connector (e.g. customer secrets / config). (optional) (default to false)
	fetchConnectionSecrets := true // bool | Fetch the general connection secrets associated with the connector. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorGetV1(context.Background(), connectorIdOrName).FetchCustomerData(fetchCustomerData).FetchConnectionSecrets(fetchConnectionSecrets).Execute()
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
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorGetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchCustomerData** | **bool** | Fetch the customer data associated with the connector (e.g. customer secrets / config). | [default to false]
 **fetchConnectionSecrets** | **bool** | Fetch the general connection secrets associated with the connector. | [default to false]

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
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	connectorIdOrName := "connectorIdOrName_example" // string | 
	authType := openapiclient.OutboundAuthenticationType("oauth2") // OutboundAuthenticationType |  (optional)
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
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListOrganizationCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | [**OutboundAuthenticationType**](OutboundAuthenticationType.md) |  | 
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

> ResponseConnectorListToolsV1 ConnectorListToolsV1(ctx, connectorIdOrName).Page(page).PageSize(pageSize).Refresh(refresh).Pretty(pretty).CredentialsName(credentialsName).Execute()

List tools for a connector.



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
	connectorIdOrName := "connectorIdOrName_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	refresh := true // bool |  (optional) (default to false)
	pretty := true // bool | Return a simplified payload with only name, description, annotations, and a compact inputSchema. (optional) (default to false)
	credentialsName := "credentialsName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorListToolsV1(context.Background(), connectorIdOrName).Page(page).PageSize(pageSize).Refresh(refresh).Pretty(pretty).CredentialsName(credentialsName).Execute()
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
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListToolsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
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
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	connectorIdOrName := "connectorIdOrName_example" // string | 
	authType := openapiclient.OutboundAuthenticationType("oauth2") // OutboundAuthenticationType |  (optional)
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
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListUserCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | [**OutboundAuthenticationType**](OutboundAuthenticationType.md) |  | 
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
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	queryFilters := *openapiclient.NewConnectorsQueryFilters() // ConnectorsQueryFilters |  (optional) (default to {"fetch_user_data":false,"fetch_customer_data":false,"fetch_connection_secrets":false,"fetch_execution_data":false})
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
 **queryFilters** | [**ConnectorsQueryFilters**](ConnectorsQueryFilters.md) |  | [default to {&quot;fetch_user_data&quot;:false,&quot;fetch_customer_data&quot;:false,&quot;fetch_connection_secrets&quot;:false,&quot;fetch_execution_data&quot;:false}]
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
	openapiclient "github.com/cterence/mistral-client-golang"
)

func main() {
	connectorIdOrName := "connectorIdOrName_example" // string | 
	authType := openapiclient.OutboundAuthenticationType("oauth2") // OutboundAuthenticationType |  (optional)
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
**connectorIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorListWorkspaceCredentialsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | [**OutboundAuthenticationType**](OutboundAuthenticationType.md) |  | 
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


## ConnectorUpdateV1

> Connector ConnectorUpdateV1(ctx, connectorId).ConnectorMCPUpdate(connectorMCPUpdate).Execute()

Update a connector.



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
	connectorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	connectorMCPUpdate := *openapiclient.NewConnectorMCPUpdate() // ConnectorMCPUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaConnectorsAPI.ConnectorUpdateV1(context.Background(), connectorId).ConnectorMCPUpdate(connectorMCPUpdate).Execute()
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

 **connectorMCPUpdate** | [**ConnectorMCPUpdate**](ConnectorMCPUpdate.md) |  | 

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


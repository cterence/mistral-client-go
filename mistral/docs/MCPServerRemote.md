# MCPServerRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Transport type | 
**Url** | **string** | Transport endpoint URL | 
**SupportedProtocolVersions** | Pointer to **[]string** |  | [optional] 
**Headers** | Pointer to [**[]MCPServerRemoteHeader**](MCPServerRemoteHeader.md) |  | [optional] 
**Authentication** | Pointer to [**NullableMCPServerAuthenticationRequirement**](MCPServerAuthenticationRequirement.md) |  | [optional] 

## Methods

### NewMCPServerRemote

`func NewMCPServerRemote(type_ string, url string, ) *MCPServerRemote`

NewMCPServerRemote instantiates a new MCPServerRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPServerRemoteWithDefaults

`func NewMCPServerRemoteWithDefaults() *MCPServerRemote`

NewMCPServerRemoteWithDefaults instantiates a new MCPServerRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MCPServerRemote) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MCPServerRemote) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MCPServerRemote) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *MCPServerRemote) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MCPServerRemote) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MCPServerRemote) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSupportedProtocolVersions

`func (o *MCPServerRemote) GetSupportedProtocolVersions() []string`

GetSupportedProtocolVersions returns the SupportedProtocolVersions field if non-nil, zero value otherwise.

### GetSupportedProtocolVersionsOk

`func (o *MCPServerRemote) GetSupportedProtocolVersionsOk() (*[]string, bool)`

GetSupportedProtocolVersionsOk returns a tuple with the SupportedProtocolVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocolVersions

`func (o *MCPServerRemote) SetSupportedProtocolVersions(v []string)`

SetSupportedProtocolVersions sets SupportedProtocolVersions field to given value.

### HasSupportedProtocolVersions

`func (o *MCPServerRemote) HasSupportedProtocolVersions() bool`

HasSupportedProtocolVersions returns a boolean if a field has been set.

### SetSupportedProtocolVersionsNil

`func (o *MCPServerRemote) SetSupportedProtocolVersionsNil(b bool)`

 SetSupportedProtocolVersionsNil sets the value for SupportedProtocolVersions to be an explicit nil

### UnsetSupportedProtocolVersions
`func (o *MCPServerRemote) UnsetSupportedProtocolVersions()`

UnsetSupportedProtocolVersions ensures that no value is present for SupportedProtocolVersions, not even an explicit nil
### GetHeaders

`func (o *MCPServerRemote) GetHeaders() []MCPServerRemoteHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MCPServerRemote) GetHeadersOk() (*[]MCPServerRemoteHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MCPServerRemote) SetHeaders(v []MCPServerRemoteHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MCPServerRemote) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *MCPServerRemote) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *MCPServerRemote) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetAuthentication

`func (o *MCPServerRemote) GetAuthentication() MCPServerAuthenticationRequirement`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *MCPServerRemote) GetAuthenticationOk() (*MCPServerAuthenticationRequirement, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *MCPServerRemote) SetAuthentication(v MCPServerAuthenticationRequirement)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *MCPServerRemote) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *MCPServerRemote) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *MCPServerRemote) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



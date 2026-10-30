# UpdateHTTPConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Optional human-readable title for the connector. | [optional] 
**Name** | Pointer to **NullableString** | The name of the connector. | [optional] 
**Description** | Pointer to **NullableString** | The description of the connector. | [optional] 
**IconUrl** | Pointer to **NullableString** | The optional url of the icon you want to associate to the connector. | [optional] 
**Server** | Pointer to **NullableString** | New connector server URL. | [optional] 
**AuthMethods** | Pointer to [**[]ConnectorMCPPublicUpdateAuthMethodsInner**](ConnectorMCPPublicUpdateAuthMethodsInner.md) | Authentication methods supported by the connector. | [optional] 
**Protocol** | **string** |  | 

## Methods

### NewUpdateHTTPConnectorRequest

`func NewUpdateHTTPConnectorRequest(protocol string, ) *UpdateHTTPConnectorRequest`

NewUpdateHTTPConnectorRequest instantiates a new UpdateHTTPConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHTTPConnectorRequestWithDefaults

`func NewUpdateHTTPConnectorRequestWithDefaults() *UpdateHTTPConnectorRequest`

NewUpdateHTTPConnectorRequestWithDefaults instantiates a new UpdateHTTPConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateHTTPConnectorRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateHTTPConnectorRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateHTTPConnectorRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateHTTPConnectorRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateHTTPConnectorRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateHTTPConnectorRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetName

`func (o *UpdateHTTPConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHTTPConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHTTPConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateHTTPConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateHTTPConnectorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateHTTPConnectorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateHTTPConnectorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHTTPConnectorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHTTPConnectorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHTTPConnectorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateHTTPConnectorRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateHTTPConnectorRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *UpdateHTTPConnectorRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *UpdateHTTPConnectorRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *UpdateHTTPConnectorRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *UpdateHTTPConnectorRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *UpdateHTTPConnectorRequest) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *UpdateHTTPConnectorRequest) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetServer

`func (o *UpdateHTTPConnectorRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateHTTPConnectorRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateHTTPConnectorRequest) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateHTTPConnectorRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *UpdateHTTPConnectorRequest) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *UpdateHTTPConnectorRequest) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetAuthMethods

`func (o *UpdateHTTPConnectorRequest) GetAuthMethods() []ConnectorMCPPublicUpdateAuthMethodsInner`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *UpdateHTTPConnectorRequest) GetAuthMethodsOk() (*[]ConnectorMCPPublicUpdateAuthMethodsInner, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *UpdateHTTPConnectorRequest) SetAuthMethods(v []ConnectorMCPPublicUpdateAuthMethodsInner)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *UpdateHTTPConnectorRequest) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### SetAuthMethodsNil

`func (o *UpdateHTTPConnectorRequest) SetAuthMethodsNil(b bool)`

 SetAuthMethodsNil sets the value for AuthMethods to be an explicit nil

### UnsetAuthMethods
`func (o *UpdateHTTPConnectorRequest) UnsetAuthMethods()`

UnsetAuthMethods ensures that no value is present for AuthMethods, not even an explicit nil
### GetProtocol

`func (o *UpdateHTTPConnectorRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateHTTPConnectorRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateHTTPConnectorRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



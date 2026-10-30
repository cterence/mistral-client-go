# Payload1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Optional human-readable title for the connector. | [optional] 
**Name** | Pointer to **string** | The name of the connector. | [optional] 
**Description** | Pointer to **string** | The description of the connector. | [optional] 
**IconUrl** | Pointer to **string** | The optional url of the icon you want to associate to the connector. | [optional] 
**Server** | Pointer to **string** | New connector server URL. | [optional] 
**AuthMethods** | Pointer to [**[]ConnectorMCPPublicUpdateAuthMethodsInner**](ConnectorMCPPublicUpdateAuthMethodsInner.md) | Authentication methods supported by the connector. | [optional] 
**Protocol** | **string** |  | 
**SystemPrompt** | Pointer to **string** | Optional system prompt for the connector. | [optional] 

## Methods

### NewPayload1

`func NewPayload1(protocol string, ) *Payload1`

NewPayload1 instantiates a new Payload1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayload1WithDefaults

`func NewPayload1WithDefaults() *Payload1`

NewPayload1WithDefaults instantiates a new Payload1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Payload1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Payload1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Payload1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Payload1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *Payload1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Payload1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Payload1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Payload1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Payload1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Payload1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Payload1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Payload1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *Payload1) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Payload1) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Payload1) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Payload1) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetServer

`func (o *Payload1) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Payload1) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Payload1) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *Payload1) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAuthMethods

`func (o *Payload1) GetAuthMethods() []ConnectorMCPPublicUpdateAuthMethodsInner`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *Payload1) GetAuthMethodsOk() (*[]ConnectorMCPPublicUpdateAuthMethodsInner, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *Payload1) SetAuthMethods(v []ConnectorMCPPublicUpdateAuthMethodsInner)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *Payload1) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### GetProtocol

`func (o *Payload1) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Payload1) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Payload1) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSystemPrompt

`func (o *Payload1) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *Payload1) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *Payload1) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *Payload1) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



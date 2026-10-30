# PublicConnectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ConnectionConfigType**](ConnectionConfigType.md) |  | [optional] [default to CONNECTIONCONFIGTYPE_MCP]
**BaseUrl** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Signed** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPublicConnectionConfig

`func NewPublicConnectionConfig() *PublicConnectionConfig`

NewPublicConnectionConfig instantiates a new PublicConnectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConnectionConfigWithDefaults

`func NewPublicConnectionConfigWithDefaults() *PublicConnectionConfig`

NewPublicConnectionConfigWithDefaults instantiates a new PublicConnectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicConnectionConfig) GetType() ConnectionConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicConnectionConfig) GetTypeOk() (*ConnectionConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicConnectionConfig) SetType(v ConnectionConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *PublicConnectionConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseUrl

`func (o *PublicConnectionConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PublicConnectionConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PublicConnectionConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *PublicConnectionConfig) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *PublicConnectionConfig) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *PublicConnectionConfig) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetHeaders

`func (o *PublicConnectionConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PublicConnectionConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PublicConnectionConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PublicConnectionConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *PublicConnectionConfig) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *PublicConnectionConfig) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetSigned

`func (o *PublicConnectionConfig) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *PublicConnectionConfig) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *PublicConnectionConfig) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *PublicConnectionConfig) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### SetSignedNil

`func (o *PublicConnectionConfig) SetSignedNil(b bool)`

 SetSignedNil sets the value for Signed to be an explicit nil

### UnsetSigned
`func (o *PublicConnectionConfig) UnsetSigned()`

UnsetSigned ensures that no value is present for Signed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



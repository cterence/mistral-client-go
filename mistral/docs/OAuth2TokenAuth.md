# OAuth2TokenAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "oauth2-token"]
**Value** | **string** |  | 

## Methods

### NewOAuth2TokenAuth

`func NewOAuth2TokenAuth(value string, ) *OAuth2TokenAuth`

NewOAuth2TokenAuth instantiates a new OAuth2TokenAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2TokenAuthWithDefaults

`func NewOAuth2TokenAuthWithDefaults() *OAuth2TokenAuth`

NewOAuth2TokenAuthWithDefaults instantiates a new OAuth2TokenAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OAuth2TokenAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuth2TokenAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuth2TokenAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuth2TokenAuth) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *OAuth2TokenAuth) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OAuth2TokenAuth) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OAuth2TokenAuth) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConnectionCredentialsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oauth** | Pointer to [**NullableOAuth2ClientCredentialsInput**](OAuth2ClientCredentialsInput.md) |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**BearerToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConnectionCredentialsInput

`func NewConnectionCredentialsInput() *ConnectionCredentialsInput`

NewConnectionCredentialsInput instantiates a new ConnectionCredentialsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCredentialsInputWithDefaults

`func NewConnectionCredentialsInputWithDefaults() *ConnectionCredentialsInput`

NewConnectionCredentialsInputWithDefaults instantiates a new ConnectionCredentialsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauth

`func (o *ConnectionCredentialsInput) GetOauth() OAuth2ClientCredentialsInput`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *ConnectionCredentialsInput) GetOauthOk() (*OAuth2ClientCredentialsInput, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *ConnectionCredentialsInput) SetOauth(v OAuth2ClientCredentialsInput)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *ConnectionCredentialsInput) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### SetOauthNil

`func (o *ConnectionCredentialsInput) SetOauthNil(b bool)`

 SetOauthNil sets the value for Oauth to be an explicit nil

### UnsetOauth
`func (o *ConnectionCredentialsInput) UnsetOauth()`

UnsetOauth ensures that no value is present for Oauth, not even an explicit nil
### GetHeaders

`func (o *ConnectionCredentialsInput) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectionCredentialsInput) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectionCredentialsInput) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectionCredentialsInput) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *ConnectionCredentialsInput) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *ConnectionCredentialsInput) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetBearerToken

`func (o *ConnectionCredentialsInput) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *ConnectionCredentialsInput) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *ConnectionCredentialsInput) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *ConnectionCredentialsInput) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### SetBearerTokenNil

`func (o *ConnectionCredentialsInput) SetBearerTokenNil(b bool)`

 SetBearerTokenNil sets the value for BearerToken to be an explicit nil

### UnsetBearerToken
`func (o *ConnectionCredentialsInput) UnsetBearerToken()`

UnsetBearerToken ensures that no value is present for BearerToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



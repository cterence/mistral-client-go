# AuthenticationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**AuthenticationType** | [**AuthenticationType**](AuthenticationType.md) |  | 
**Scope** | [**ConsumerType**](ConsumerType.md) |  | 
**Status** | Pointer to [**NullableCredentialsStatus**](CredentialsStatus.md) |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]
**CreatorId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuthenticationConfiguration

`func NewAuthenticationConfiguration(name string, authenticationType AuthenticationType, scope ConsumerType, ) *AuthenticationConfiguration`

NewAuthenticationConfiguration instantiates a new AuthenticationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationConfigurationWithDefaults

`func NewAuthenticationConfigurationWithDefaults() *AuthenticationConfiguration`

NewAuthenticationConfigurationWithDefaults instantiates a new AuthenticationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticationConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AuthenticationConfiguration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthenticationConfiguration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthenticationConfiguration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthenticationConfiguration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AuthenticationConfiguration) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AuthenticationConfiguration) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthenticationType

`func (o *AuthenticationConfiguration) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AuthenticationConfiguration) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AuthenticationConfiguration) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetScope

`func (o *AuthenticationConfiguration) GetScope() ConsumerType`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AuthenticationConfiguration) GetScopeOk() (*ConsumerType, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AuthenticationConfiguration) SetScope(v ConsumerType)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *AuthenticationConfiguration) GetStatus() CredentialsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticationConfiguration) GetStatusOk() (*CredentialsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticationConfiguration) SetStatus(v CredentialsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticationConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AuthenticationConfiguration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AuthenticationConfiguration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsDefault

`func (o *AuthenticationConfiguration) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AuthenticationConfiguration) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AuthenticationConfiguration) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AuthenticationConfiguration) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetCreatorId

`func (o *AuthenticationConfiguration) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AuthenticationConfiguration) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AuthenticationConfiguration) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AuthenticationConfiguration) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *AuthenticationConfiguration) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *AuthenticationConfiguration) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



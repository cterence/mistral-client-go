# AuthenticationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationType** | [**OutboundAuthenticationType**](OutboundAuthenticationType.md) |  | 
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthenticationConfiguration

`func NewAuthenticationConfiguration(name string, authenticationType OutboundAuthenticationType, ) *AuthenticationConfiguration`

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


### GetAuthenticationType

`func (o *AuthenticationConfiguration) GetAuthenticationType() OutboundAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AuthenticationConfiguration) GetAuthenticationTypeOk() (*OutboundAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AuthenticationConfiguration) SetAuthenticationType(v OutboundAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



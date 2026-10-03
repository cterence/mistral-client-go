# CredentialsCreateOrUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the credentials. Use this name to access or modify your credentials. | 
**IsDefault** | Pointer to **NullableBool** | Controls whether this credential is the default for its auth method. On creation: if no credential exists yet for this auth method, the credential is automatically set as default when is_default is true or omitted; setting is_default to false is rejected because a default must exist. If other credentials already exist, setting is_default to true promotes this credential (demoting the previous default); false or omitted creates it as non-default. On update: true promotes this credential, false is rejected if it is currently the default (promote another credential first), omitted leaves the default status unchanged. | [optional] 
**Credentials** | Pointer to [**NullableConnectionCredentials**](ConnectionCredentials.md) | The credential data (headers, bearer_token). | [optional] 

## Methods

### NewCredentialsCreateOrUpdate

`func NewCredentialsCreateOrUpdate(name string, ) *CredentialsCreateOrUpdate`

NewCredentialsCreateOrUpdate instantiates a new CredentialsCreateOrUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsCreateOrUpdateWithDefaults

`func NewCredentialsCreateOrUpdateWithDefaults() *CredentialsCreateOrUpdate`

NewCredentialsCreateOrUpdateWithDefaults instantiates a new CredentialsCreateOrUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CredentialsCreateOrUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialsCreateOrUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialsCreateOrUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *CredentialsCreateOrUpdate) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CredentialsCreateOrUpdate) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CredentialsCreateOrUpdate) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CredentialsCreateOrUpdate) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *CredentialsCreateOrUpdate) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *CredentialsCreateOrUpdate) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetCredentials

`func (o *CredentialsCreateOrUpdate) GetCredentials() ConnectionCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialsCreateOrUpdate) GetCredentialsOk() (*ConnectionCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialsCreateOrUpdate) SetCredentials(v ConnectionCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CredentialsCreateOrUpdate) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *CredentialsCreateOrUpdate) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *CredentialsCreateOrUpdate) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



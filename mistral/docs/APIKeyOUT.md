# APIKeyOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | API key ID. | 
**Key** | **NullableString** | Plaintext API key value. Only returned at creation time. | 
**Name** | **NullableString** | Name of the API key. | 
**HiddenKey** | **string** | Masked API key value for display. | 
**CreatedAt** | **NullableTime** | Time when the API key was created. | 
**ExpirationDate** | **NullableString** | Date when the API key expires. | 
**Actions** | [**APIKeyActions**](APIKeyActions.md) | Per-action availability for this API key (e.g. whether it can be rotated or deleted). | 

## Methods

### NewAPIKeyOUT

`func NewAPIKeyOUT(keyId string, key NullableString, name NullableString, hiddenKey string, createdAt NullableTime, expirationDate NullableString, actions APIKeyActions, ) *APIKeyOUT`

NewAPIKeyOUT instantiates a new APIKeyOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyOUTWithDefaults

`func NewAPIKeyOUTWithDefaults() *APIKeyOUT`

NewAPIKeyOUTWithDefaults instantiates a new APIKeyOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *APIKeyOUT) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *APIKeyOUT) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *APIKeyOUT) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKey

`func (o *APIKeyOUT) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *APIKeyOUT) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *APIKeyOUT) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *APIKeyOUT) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *APIKeyOUT) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *APIKeyOUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIKeyOUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIKeyOUT) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *APIKeyOUT) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *APIKeyOUT) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHiddenKey

`func (o *APIKeyOUT) GetHiddenKey() string`

GetHiddenKey returns the HiddenKey field if non-nil, zero value otherwise.

### GetHiddenKeyOk

`func (o *APIKeyOUT) GetHiddenKeyOk() (*string, bool)`

GetHiddenKeyOk returns a tuple with the HiddenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenKey

`func (o *APIKeyOUT) SetHiddenKey(v string)`

SetHiddenKey sets HiddenKey field to given value.


### GetCreatedAt

`func (o *APIKeyOUT) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIKeyOUT) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIKeyOUT) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *APIKeyOUT) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *APIKeyOUT) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpirationDate

`func (o *APIKeyOUT) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *APIKeyOUT) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *APIKeyOUT) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *APIKeyOUT) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *APIKeyOUT) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetActions

`func (o *APIKeyOUT) GetActions() APIKeyActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *APIKeyOUT) GetActionsOk() (*APIKeyActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *APIKeyOUT) SetActions(v APIKeyActions)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



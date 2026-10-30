# WhoamiServiceAccountPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalType** | Pointer to **string** | Discriminator identifying a service-account principal. | [optional] [default to "service_account"]
**ServiceAccountId** | **string** | The service account&#39;s unique identifier. | 
**Name** | Pointer to **NullableString** | The service account&#39;s display name, when known. | [optional] 

## Methods

### NewWhoamiServiceAccountPrincipal

`func NewWhoamiServiceAccountPrincipal(serviceAccountId string, ) *WhoamiServiceAccountPrincipal`

NewWhoamiServiceAccountPrincipal instantiates a new WhoamiServiceAccountPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiServiceAccountPrincipalWithDefaults

`func NewWhoamiServiceAccountPrincipalWithDefaults() *WhoamiServiceAccountPrincipal`

NewWhoamiServiceAccountPrincipalWithDefaults instantiates a new WhoamiServiceAccountPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalType

`func (o *WhoamiServiceAccountPrincipal) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *WhoamiServiceAccountPrincipal) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *WhoamiServiceAccountPrincipal) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *WhoamiServiceAccountPrincipal) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *WhoamiServiceAccountPrincipal) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *WhoamiServiceAccountPrincipal) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *WhoamiServiceAccountPrincipal) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetName

`func (o *WhoamiServiceAccountPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhoamiServiceAccountPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhoamiServiceAccountPrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhoamiServiceAccountPrincipal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WhoamiServiceAccountPrincipal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WhoamiServiceAccountPrincipal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



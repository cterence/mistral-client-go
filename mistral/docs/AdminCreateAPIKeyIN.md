# AdminCreateAPIKeyIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Optional name for the API key. | [optional] 
**Expiration** | Pointer to **NullableString** | Date when the API key should expire. | [optional] 
**WorkspaceUuid** | **string** | Workspace ID for the API key. | 
**UserId** | **string** | ID of the user the API key is created for. | 

## Methods

### NewAdminCreateAPIKeyIN

`func NewAdminCreateAPIKeyIN(workspaceUuid string, userId string, ) *AdminCreateAPIKeyIN`

NewAdminCreateAPIKeyIN instantiates a new AdminCreateAPIKeyIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCreateAPIKeyINWithDefaults

`func NewAdminCreateAPIKeyINWithDefaults() *AdminCreateAPIKeyIN`

NewAdminCreateAPIKeyINWithDefaults instantiates a new AdminCreateAPIKeyIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminCreateAPIKeyIN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminCreateAPIKeyIN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminCreateAPIKeyIN) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminCreateAPIKeyIN) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminCreateAPIKeyIN) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminCreateAPIKeyIN) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExpiration

`func (o *AdminCreateAPIKeyIN) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AdminCreateAPIKeyIN) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AdminCreateAPIKeyIN) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AdminCreateAPIKeyIN) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *AdminCreateAPIKeyIN) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *AdminCreateAPIKeyIN) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetWorkspaceUuid

`func (o *AdminCreateAPIKeyIN) GetWorkspaceUuid() string`

GetWorkspaceUuid returns the WorkspaceUuid field if non-nil, zero value otherwise.

### GetWorkspaceUuidOk

`func (o *AdminCreateAPIKeyIN) GetWorkspaceUuidOk() (*string, bool)`

GetWorkspaceUuidOk returns a tuple with the WorkspaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuid

`func (o *AdminCreateAPIKeyIN) SetWorkspaceUuid(v string)`

SetWorkspaceUuid sets WorkspaceUuid field to given value.


### GetUserId

`func (o *AdminCreateAPIKeyIN) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AdminCreateAPIKeyIN) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AdminCreateAPIKeyIN) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



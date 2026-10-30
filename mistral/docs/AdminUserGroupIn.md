# AdminUserGroupIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the user group. | 
**Description** | Pointer to **NullableString** | Optional description of the user group. | [optional] 
**ParentGroupIds** | Pointer to **[]string** | UUIDs of existing groups to nest this new group inside. Omit or pass an empty list to create a top-level group. | [optional] 

## Methods

### NewAdminUserGroupIn

`func NewAdminUserGroupIn(name string, ) *AdminUserGroupIn`

NewAdminUserGroupIn instantiates a new AdminUserGroupIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserGroupInWithDefaults

`func NewAdminUserGroupInWithDefaults() *AdminUserGroupIn`

NewAdminUserGroupInWithDefaults instantiates a new AdminUserGroupIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminUserGroupIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminUserGroupIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminUserGroupIn) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdminUserGroupIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminUserGroupIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminUserGroupIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminUserGroupIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdminUserGroupIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdminUserGroupIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentGroupIds

`func (o *AdminUserGroupIn) GetParentGroupIds() []string`

GetParentGroupIds returns the ParentGroupIds field if non-nil, zero value otherwise.

### GetParentGroupIdsOk

`func (o *AdminUserGroupIn) GetParentGroupIdsOk() (*[]string, bool)`

GetParentGroupIdsOk returns a tuple with the ParentGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupIds

`func (o *AdminUserGroupIn) SetParentGroupIds(v []string)`

SetParentGroupIds sets ParentGroupIds field to given value.

### HasParentGroupIds

`func (o *AdminUserGroupIn) HasParentGroupIds() bool`

HasParentGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RoleOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Role name. | 
**Uuid** | **string** | Role UUID. | 
**Description** | **string** | Short description of the role. | 
**IsCustomRole** | **bool** | Whether this role is custom-defined. | 

## Methods

### NewRoleOut

`func NewRoleOut(name string, uuid string, description string, isCustomRole bool, ) *RoleOut`

NewRoleOut instantiates a new RoleOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleOutWithDefaults

`func NewRoleOutWithDefaults() *RoleOut`

NewRoleOutWithDefaults instantiates a new RoleOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleOut) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *RoleOut) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoleOut) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoleOut) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetDescription

`func (o *RoleOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsCustomRole

`func (o *RoleOut) GetIsCustomRole() bool`

GetIsCustomRole returns the IsCustomRole field if non-nil, zero value otherwise.

### GetIsCustomRoleOk

`func (o *RoleOut) GetIsCustomRoleOk() (*bool, bool)`

GetIsCustomRoleOk returns a tuple with the IsCustomRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomRole

`func (o *RoleOut) SetIsCustomRole(v bool)`

SetIsCustomRole sets IsCustomRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



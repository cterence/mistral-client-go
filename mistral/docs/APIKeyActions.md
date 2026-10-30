# APIKeyActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rotate** | Pointer to [**NullableAPIKeyActionsRotate**](APIKeyActionsRotate.md) |  | [optional] 
**Delete** | Pointer to [**NullableAPIKeyActionsDelete**](APIKeyActionsDelete.md) |  | [optional] 

## Methods

### NewAPIKeyActions

`func NewAPIKeyActions() *APIKeyActions`

NewAPIKeyActions instantiates a new APIKeyActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyActionsWithDefaults

`func NewAPIKeyActionsWithDefaults() *APIKeyActions`

NewAPIKeyActionsWithDefaults instantiates a new APIKeyActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRotate

`func (o *APIKeyActions) GetRotate() APIKeyActionsRotate`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *APIKeyActions) GetRotateOk() (*APIKeyActionsRotate, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *APIKeyActions) SetRotate(v APIKeyActionsRotate)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *APIKeyActions) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### SetRotateNil

`func (o *APIKeyActions) SetRotateNil(b bool)`

 SetRotateNil sets the value for Rotate to be an explicit nil

### UnsetRotate
`func (o *APIKeyActions) UnsetRotate()`

UnsetRotate ensures that no value is present for Rotate, not even an explicit nil
### GetDelete

`func (o *APIKeyActions) GetDelete() APIKeyActionsDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *APIKeyActions) GetDeleteOk() (*APIKeyActionsDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *APIKeyActions) SetDelete(v APIKeyActionsDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *APIKeyActions) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *APIKeyActions) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *APIKeyActions) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



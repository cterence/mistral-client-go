# BaseFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Type** | **string** |  | 
**Group** | Pointer to **NullableString** |  | [optional] 
**SupportedOperators** | **[]string** |  | [readonly] 

## Methods

### NewBaseFieldDefinition

`func NewBaseFieldDefinition(name string, label string, type_ string, supportedOperators []string, ) *BaseFieldDefinition`

NewBaseFieldDefinition instantiates a new BaseFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFieldDefinitionWithDefaults

`func NewBaseFieldDefinitionWithDefaults() *BaseFieldDefinition`

NewBaseFieldDefinitionWithDefaults instantiates a new BaseFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BaseFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BaseFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BaseFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *BaseFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseFieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetGroup

`func (o *BaseFieldDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BaseFieldDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BaseFieldDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BaseFieldDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BaseFieldDefinition) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BaseFieldDefinition) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSupportedOperators

`func (o *BaseFieldDefinition) GetSupportedOperators() []string`

GetSupportedOperators returns the SupportedOperators field if non-nil, zero value otherwise.

### GetSupportedOperatorsOk

`func (o *BaseFieldDefinition) GetSupportedOperatorsOk() (*[]string, bool)`

GetSupportedOperatorsOk returns a tuple with the SupportedOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperators

`func (o *BaseFieldDefinition) SetSupportedOperators(v []string)`

SetSupportedOperators sets SupportedOperators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



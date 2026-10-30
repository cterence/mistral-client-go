# Tools1Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**Description** | **string** |  | 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Tools** | [**[]Tools1Inner1**](Tools1Inner1.md) |  | 

## Methods

### NewTools1Inner

`func NewTools1Inner(name string, type_ string, description string, tools []Tools1Inner1, ) *Tools1Inner`

NewTools1Inner instantiates a new Tools1Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTools1InnerWithDefaults

`func NewTools1InnerWithDefaults() *Tools1Inner`

NewTools1InnerWithDefaults instantiates a new Tools1Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tools1Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tools1Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tools1Inner) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *Tools1Inner) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Tools1Inner) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Tools1Inner) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Tools1Inner) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *Tools1Inner) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *Tools1Inner) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetStrict

`func (o *Tools1Inner) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *Tools1Inner) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *Tools1Inner) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *Tools1Inner) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *Tools1Inner) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *Tools1Inner) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil
### GetType

`func (o *Tools1Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tools1Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tools1Inner) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *Tools1Inner) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *Tools1Inner) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *Tools1Inner) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *Tools1Inner) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *Tools1Inner) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *Tools1Inner) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetDeferLoading

`func (o *Tools1Inner) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *Tools1Inner) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *Tools1Inner) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *Tools1Inner) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *Tools1Inner) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *Tools1Inner) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetDescription

`func (o *Tools1Inner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tools1Inner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tools1Inner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOutputSchema

`func (o *Tools1Inner) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *Tools1Inner) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *Tools1Inner) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *Tools1Inner) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *Tools1Inner) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *Tools1Inner) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetTools

`func (o *Tools1Inner) GetTools() []Tools1Inner1`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Tools1Inner) GetToolsOk() (*[]Tools1Inner1, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Tools1Inner) SetTools(v []Tools1Inner1)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



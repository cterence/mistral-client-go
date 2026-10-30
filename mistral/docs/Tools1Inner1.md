# Tools1Inner1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 
**Format** | Pointer to [**NullableFormat**](Format.md) |  | [optional] 

## Methods

### NewTools1Inner1

`func NewTools1Inner1(name string, type_ string, ) *Tools1Inner1`

NewTools1Inner1 instantiates a new Tools1Inner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTools1Inner1WithDefaults

`func NewTools1Inner1WithDefaults() *Tools1Inner1`

NewTools1Inner1WithDefaults instantiates a new Tools1Inner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tools1Inner1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tools1Inner1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tools1Inner1) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Tools1Inner1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tools1Inner1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tools1Inner1) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *Tools1Inner1) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *Tools1Inner1) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *Tools1Inner1) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *Tools1Inner1) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *Tools1Inner1) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *Tools1Inner1) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetDeferLoading

`func (o *Tools1Inner1) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *Tools1Inner1) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *Tools1Inner1) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *Tools1Inner1) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *Tools1Inner1) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *Tools1Inner1) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetDescription

`func (o *Tools1Inner1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tools1Inner1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tools1Inner1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tools1Inner1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Tools1Inner1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tools1Inner1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOutputSchema

`func (o *Tools1Inner1) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *Tools1Inner1) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *Tools1Inner1) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *Tools1Inner1) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *Tools1Inner1) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *Tools1Inner1) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetParameters

`func (o *Tools1Inner1) GetParameters() AnyOf`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Tools1Inner1) GetParametersOk() (*AnyOf, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Tools1Inner1) SetParameters(v AnyOf)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Tools1Inner1) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *Tools1Inner1) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *Tools1Inner1) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetStrict

`func (o *Tools1Inner1) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *Tools1Inner1) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *Tools1Inner1) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *Tools1Inner1) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *Tools1Inner1) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *Tools1Inner1) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil
### GetFormat

`func (o *Tools1Inner1) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Tools1Inner1) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Tools1Inner1) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Tools1Inner1) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *Tools1Inner1) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *Tools1Inner1) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



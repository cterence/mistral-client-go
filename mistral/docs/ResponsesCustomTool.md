# ResponsesCustomTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to [**NullableFormat**](Format.md) |  | [optional] 

## Methods

### NewResponsesCustomTool

`func NewResponsesCustomTool(name string, type_ string, ) *ResponsesCustomTool`

NewResponsesCustomTool instantiates a new ResponsesCustomTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesCustomToolWithDefaults

`func NewResponsesCustomToolWithDefaults() *ResponsesCustomTool`

NewResponsesCustomToolWithDefaults instantiates a new ResponsesCustomTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponsesCustomTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesCustomTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesCustomTool) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResponsesCustomTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesCustomTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesCustomTool) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *ResponsesCustomTool) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *ResponsesCustomTool) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *ResponsesCustomTool) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *ResponsesCustomTool) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *ResponsesCustomTool) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *ResponsesCustomTool) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetDeferLoading

`func (o *ResponsesCustomTool) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *ResponsesCustomTool) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *ResponsesCustomTool) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *ResponsesCustomTool) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *ResponsesCustomTool) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *ResponsesCustomTool) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetDescription

`func (o *ResponsesCustomTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponsesCustomTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponsesCustomTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponsesCustomTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResponsesCustomTool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResponsesCustomTool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFormat

`func (o *ResponsesCustomTool) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ResponsesCustomTool) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ResponsesCustomTool) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ResponsesCustomTool) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *ResponsesCustomTool) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *ResponsesCustomTool) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



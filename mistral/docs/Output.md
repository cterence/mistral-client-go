# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "CLASSIFICATION"]
**Options** | [**[]JudgeClassificationOutputOption**](JudgeClassificationOutputOption.md) |  | 
**Min** | Pointer to **float32** |  | [optional] [default to 0]
**MinDescription** | **string** |  | 
**Max** | Pointer to **float32** |  | [optional] [default to 1]
**MaxDescription** | **string** |  | 

## Methods

### NewOutput

`func NewOutput(options []JudgeClassificationOutputOption, minDescription string, maxDescription string, ) *Output`

NewOutput instantiates a new Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWithDefaults

`func NewOutputWithDefaults() *Output`

NewOutputWithDefaults instantiates a new Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Output) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Output) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Output) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Output) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *Output) GetOptions() []JudgeClassificationOutputOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Output) GetOptionsOk() (*[]JudgeClassificationOutputOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Output) SetOptions(v []JudgeClassificationOutputOption)`

SetOptions sets Options field to given value.


### GetMin

`func (o *Output) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Output) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Output) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *Output) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMinDescription

`func (o *Output) GetMinDescription() string`

GetMinDescription returns the MinDescription field if non-nil, zero value otherwise.

### GetMinDescriptionOk

`func (o *Output) GetMinDescriptionOk() (*string, bool)`

GetMinDescriptionOk returns a tuple with the MinDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDescription

`func (o *Output) SetMinDescription(v string)`

SetMinDescription sets MinDescription field to given value.


### GetMax

`func (o *Output) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Output) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Output) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *Output) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxDescription

`func (o *Output) GetMaxDescription() string`

GetMaxDescription returns the MaxDescription field if non-nil, zero value otherwise.

### GetMaxDescriptionOk

`func (o *Output) GetMaxDescriptionOk() (*string, bool)`

GetMaxDescriptionOk returns a tuple with the MaxDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDescription

`func (o *Output) SetMaxDescription(v string)`

SetMaxDescription sets MaxDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



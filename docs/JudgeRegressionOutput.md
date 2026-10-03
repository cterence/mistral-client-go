# JudgeRegressionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "REGRESSION"]
**Min** | Pointer to **float32** |  | [optional] [default to 0]
**MinDescription** | **string** |  | 
**Max** | Pointer to **float32** |  | [optional] [default to 1]
**MaxDescription** | **string** |  | 

## Methods

### NewJudgeRegressionOutput

`func NewJudgeRegressionOutput(minDescription string, maxDescription string, ) *JudgeRegressionOutput`

NewJudgeRegressionOutput instantiates a new JudgeRegressionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeRegressionOutputWithDefaults

`func NewJudgeRegressionOutputWithDefaults() *JudgeRegressionOutput`

NewJudgeRegressionOutputWithDefaults instantiates a new JudgeRegressionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JudgeRegressionOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JudgeRegressionOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JudgeRegressionOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JudgeRegressionOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMin

`func (o *JudgeRegressionOutput) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *JudgeRegressionOutput) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *JudgeRegressionOutput) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *JudgeRegressionOutput) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMinDescription

`func (o *JudgeRegressionOutput) GetMinDescription() string`

GetMinDescription returns the MinDescription field if non-nil, zero value otherwise.

### GetMinDescriptionOk

`func (o *JudgeRegressionOutput) GetMinDescriptionOk() (*string, bool)`

GetMinDescriptionOk returns a tuple with the MinDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDescription

`func (o *JudgeRegressionOutput) SetMinDescription(v string)`

SetMinDescription sets MinDescription field to given value.


### GetMax

`func (o *JudgeRegressionOutput) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *JudgeRegressionOutput) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *JudgeRegressionOutput) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *JudgeRegressionOutput) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxDescription

`func (o *JudgeRegressionOutput) GetMaxDescription() string`

GetMaxDescription returns the MaxDescription field if non-nil, zero value otherwise.

### GetMaxDescriptionOk

`func (o *JudgeRegressionOutput) GetMaxDescriptionOk() (*string, bool)`

GetMaxDescriptionOk returns a tuple with the MaxDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDescription

`func (o *JudgeRegressionOutput) SetMaxDescription(v string)`

SetMaxDescription sets MaxDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



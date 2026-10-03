# JudgeOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | **string** |  | 
**Answer** | [**Answer**](Answer.md) |  | 

## Methods

### NewJudgeOutput

`func NewJudgeOutput(analysis string, answer Answer, ) *JudgeOutput`

NewJudgeOutput instantiates a new JudgeOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeOutputWithDefaults

`func NewJudgeOutputWithDefaults() *JudgeOutput`

NewJudgeOutputWithDefaults instantiates a new JudgeOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysis

`func (o *JudgeOutput) GetAnalysis() string`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *JudgeOutput) GetAnalysisOk() (*string, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *JudgeOutput) SetAnalysis(v string)`

SetAnalysis sets Analysis field to given value.


### GetAnswer

`func (o *JudgeOutput) GetAnswer() Answer`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *JudgeOutput) GetAnswerOk() (*Answer, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *JudgeOutput) SetAnswer(v Answer)`

SetAnswer sets Answer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



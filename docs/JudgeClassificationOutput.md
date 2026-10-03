# JudgeClassificationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "CLASSIFICATION"]
**Options** | [**[]JudgeClassificationOutputOption**](JudgeClassificationOutputOption.md) |  | 

## Methods

### NewJudgeClassificationOutput

`func NewJudgeClassificationOutput(options []JudgeClassificationOutputOption, ) *JudgeClassificationOutput`

NewJudgeClassificationOutput instantiates a new JudgeClassificationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeClassificationOutputWithDefaults

`func NewJudgeClassificationOutputWithDefaults() *JudgeClassificationOutput`

NewJudgeClassificationOutputWithDefaults instantiates a new JudgeClassificationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JudgeClassificationOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JudgeClassificationOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JudgeClassificationOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JudgeClassificationOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *JudgeClassificationOutput) GetOptions() []JudgeClassificationOutputOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JudgeClassificationOutput) GetOptionsOk() (*[]JudgeClassificationOutputOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JudgeClassificationOutput) SetOptions(v []JudgeClassificationOutputOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



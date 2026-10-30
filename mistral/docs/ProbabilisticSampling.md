# ProbabilisticSampling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "probabilistic"]
**Percentage** | **float32** |  | 

## Methods

### NewProbabilisticSampling

`func NewProbabilisticSampling(percentage float32, ) *ProbabilisticSampling`

NewProbabilisticSampling instantiates a new ProbabilisticSampling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbabilisticSamplingWithDefaults

`func NewProbabilisticSamplingWithDefaults() *ProbabilisticSampling`

NewProbabilisticSamplingWithDefaults instantiates a new ProbabilisticSampling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProbabilisticSampling) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProbabilisticSampling) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProbabilisticSampling) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProbabilisticSampling) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPercentage

`func (o *ProbabilisticSampling) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ProbabilisticSampling) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ProbabilisticSampling) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



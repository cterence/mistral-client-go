# ClassifierTargetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Labels** | **[]string** |  | 
**Weight** | **float32** |  | 
**LossFunction** | [**FTClassifierLossFunction**](FTClassifierLossFunction.md) |  | 

## Methods

### NewClassifierTargetResult

`func NewClassifierTargetResult(name string, labels []string, weight float32, lossFunction FTClassifierLossFunction, ) *ClassifierTargetResult`

NewClassifierTargetResult instantiates a new ClassifierTargetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierTargetResultWithDefaults

`func NewClassifierTargetResultWithDefaults() *ClassifierTargetResult`

NewClassifierTargetResultWithDefaults instantiates a new ClassifierTargetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClassifierTargetResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClassifierTargetResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClassifierTargetResult) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *ClassifierTargetResult) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClassifierTargetResult) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClassifierTargetResult) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetWeight

`func (o *ClassifierTargetResult) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ClassifierTargetResult) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ClassifierTargetResult) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetLossFunction

`func (o *ClassifierTargetResult) GetLossFunction() FTClassifierLossFunction`

GetLossFunction returns the LossFunction field if non-nil, zero value otherwise.

### GetLossFunctionOk

`func (o *ClassifierTargetResult) GetLossFunctionOk() (*FTClassifierLossFunction, bool)`

GetLossFunctionOk returns a tuple with the LossFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossFunction

`func (o *ClassifierTargetResult) SetLossFunction(v FTClassifierLossFunction)`

SetLossFunction sets LossFunction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



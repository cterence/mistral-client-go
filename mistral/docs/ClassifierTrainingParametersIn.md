# ClassifierTrainingParametersIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainingSteps** | Pointer to **NullableInt32** | The number of training steps to perform. A training step refers to a single update of the model weights during the fine-tuning process. This update is typically calculated using a batch of samples from the training dataset. | [optional] 
**LearningRate** | Pointer to **float32** | A parameter describing how much to adjust the pre-trained model&#39;s weights in response to the estimated error each time the weights are updated during the fine-tuning process. | [optional] [default to 1.0E-4]
**WeightDecay** | Pointer to **NullableFloat32** | (Advanced Usage) Weight decay adds a term to the loss function that is proportional to the sum of the squared weights. This term reduces the magnitude of the weights and prevents them from growing too large. | [optional] 
**WarmupFraction** | Pointer to **NullableFloat32** | (Advanced Usage) A parameter that specifies the percentage of the total training steps at which the learning rate warm-up phase ends. During this phase, the learning rate gradually increases from a small value to the initial learning rate, helping to stabilize the training process and improve convergence. Similar to &#x60;pct_start&#x60; in [mistral-finetune](https://github.com/mistralai/mistral-finetune) | [optional] 
**Epochs** | Pointer to **NullableFloat32** |  | [optional] 
**SeqLen** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewClassifierTrainingParametersIn

`func NewClassifierTrainingParametersIn() *ClassifierTrainingParametersIn`

NewClassifierTrainingParametersIn instantiates a new ClassifierTrainingParametersIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierTrainingParametersInWithDefaults

`func NewClassifierTrainingParametersInWithDefaults() *ClassifierTrainingParametersIn`

NewClassifierTrainingParametersInWithDefaults instantiates a new ClassifierTrainingParametersIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainingSteps

`func (o *ClassifierTrainingParametersIn) GetTrainingSteps() int32`

GetTrainingSteps returns the TrainingSteps field if non-nil, zero value otherwise.

### GetTrainingStepsOk

`func (o *ClassifierTrainingParametersIn) GetTrainingStepsOk() (*int32, bool)`

GetTrainingStepsOk returns a tuple with the TrainingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingSteps

`func (o *ClassifierTrainingParametersIn) SetTrainingSteps(v int32)`

SetTrainingSteps sets TrainingSteps field to given value.

### HasTrainingSteps

`func (o *ClassifierTrainingParametersIn) HasTrainingSteps() bool`

HasTrainingSteps returns a boolean if a field has been set.

### SetTrainingStepsNil

`func (o *ClassifierTrainingParametersIn) SetTrainingStepsNil(b bool)`

 SetTrainingStepsNil sets the value for TrainingSteps to be an explicit nil

### UnsetTrainingSteps
`func (o *ClassifierTrainingParametersIn) UnsetTrainingSteps()`

UnsetTrainingSteps ensures that no value is present for TrainingSteps, not even an explicit nil
### GetLearningRate

`func (o *ClassifierTrainingParametersIn) GetLearningRate() float32`

GetLearningRate returns the LearningRate field if non-nil, zero value otherwise.

### GetLearningRateOk

`func (o *ClassifierTrainingParametersIn) GetLearningRateOk() (*float32, bool)`

GetLearningRateOk returns a tuple with the LearningRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearningRate

`func (o *ClassifierTrainingParametersIn) SetLearningRate(v float32)`

SetLearningRate sets LearningRate field to given value.

### HasLearningRate

`func (o *ClassifierTrainingParametersIn) HasLearningRate() bool`

HasLearningRate returns a boolean if a field has been set.

### GetWeightDecay

`func (o *ClassifierTrainingParametersIn) GetWeightDecay() float32`

GetWeightDecay returns the WeightDecay field if non-nil, zero value otherwise.

### GetWeightDecayOk

`func (o *ClassifierTrainingParametersIn) GetWeightDecayOk() (*float32, bool)`

GetWeightDecayOk returns a tuple with the WeightDecay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightDecay

`func (o *ClassifierTrainingParametersIn) SetWeightDecay(v float32)`

SetWeightDecay sets WeightDecay field to given value.

### HasWeightDecay

`func (o *ClassifierTrainingParametersIn) HasWeightDecay() bool`

HasWeightDecay returns a boolean if a field has been set.

### SetWeightDecayNil

`func (o *ClassifierTrainingParametersIn) SetWeightDecayNil(b bool)`

 SetWeightDecayNil sets the value for WeightDecay to be an explicit nil

### UnsetWeightDecay
`func (o *ClassifierTrainingParametersIn) UnsetWeightDecay()`

UnsetWeightDecay ensures that no value is present for WeightDecay, not even an explicit nil
### GetWarmupFraction

`func (o *ClassifierTrainingParametersIn) GetWarmupFraction() float32`

GetWarmupFraction returns the WarmupFraction field if non-nil, zero value otherwise.

### GetWarmupFractionOk

`func (o *ClassifierTrainingParametersIn) GetWarmupFractionOk() (*float32, bool)`

GetWarmupFractionOk returns a tuple with the WarmupFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupFraction

`func (o *ClassifierTrainingParametersIn) SetWarmupFraction(v float32)`

SetWarmupFraction sets WarmupFraction field to given value.

### HasWarmupFraction

`func (o *ClassifierTrainingParametersIn) HasWarmupFraction() bool`

HasWarmupFraction returns a boolean if a field has been set.

### SetWarmupFractionNil

`func (o *ClassifierTrainingParametersIn) SetWarmupFractionNil(b bool)`

 SetWarmupFractionNil sets the value for WarmupFraction to be an explicit nil

### UnsetWarmupFraction
`func (o *ClassifierTrainingParametersIn) UnsetWarmupFraction()`

UnsetWarmupFraction ensures that no value is present for WarmupFraction, not even an explicit nil
### GetEpochs

`func (o *ClassifierTrainingParametersIn) GetEpochs() float32`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *ClassifierTrainingParametersIn) GetEpochsOk() (*float32, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *ClassifierTrainingParametersIn) SetEpochs(v float32)`

SetEpochs sets Epochs field to given value.

### HasEpochs

`func (o *ClassifierTrainingParametersIn) HasEpochs() bool`

HasEpochs returns a boolean if a field has been set.

### SetEpochsNil

`func (o *ClassifierTrainingParametersIn) SetEpochsNil(b bool)`

 SetEpochsNil sets the value for Epochs to be an explicit nil

### UnsetEpochs
`func (o *ClassifierTrainingParametersIn) UnsetEpochs()`

UnsetEpochs ensures that no value is present for Epochs, not even an explicit nil
### GetSeqLen

`func (o *ClassifierTrainingParametersIn) GetSeqLen() int32`

GetSeqLen returns the SeqLen field if non-nil, zero value otherwise.

### GetSeqLenOk

`func (o *ClassifierTrainingParametersIn) GetSeqLenOk() (*int32, bool)`

GetSeqLenOk returns a tuple with the SeqLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqLen

`func (o *ClassifierTrainingParametersIn) SetSeqLen(v int32)`

SetSeqLen sets SeqLen field to given value.

### HasSeqLen

`func (o *ClassifierTrainingParametersIn) HasSeqLen() bool`

HasSeqLen returns a boolean if a field has been set.

### SetSeqLenNil

`func (o *ClassifierTrainingParametersIn) SetSeqLenNil(b bool)`

 SetSeqLenNil sets the value for SeqLen to be an explicit nil

### UnsetSeqLen
`func (o *ClassifierTrainingParametersIn) UnsetSeqLen()`

UnsetSeqLen ensures that no value is present for SeqLen, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



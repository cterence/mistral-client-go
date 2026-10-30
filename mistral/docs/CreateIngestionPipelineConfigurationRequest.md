# CreateIngestionPipelineConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PipelineComposition** | Pointer to **map[string]string** |  | [optional] 
**TargetIndexes** | Pointer to [**[]IngestionPipelineTargetIndexRef**](IngestionPipelineTargetIndexRef.md) |  | [optional] 

## Methods

### NewCreateIngestionPipelineConfigurationRequest

`func NewCreateIngestionPipelineConfigurationRequest(name string, ) *CreateIngestionPipelineConfigurationRequest`

NewCreateIngestionPipelineConfigurationRequest instantiates a new CreateIngestionPipelineConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIngestionPipelineConfigurationRequestWithDefaults

`func NewCreateIngestionPipelineConfigurationRequestWithDefaults() *CreateIngestionPipelineConfigurationRequest`

NewCreateIngestionPipelineConfigurationRequestWithDefaults instantiates a new CreateIngestionPipelineConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIngestionPipelineConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIngestionPipelineConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIngestionPipelineConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPipelineComposition

`func (o *CreateIngestionPipelineConfigurationRequest) GetPipelineComposition() map[string]string`

GetPipelineComposition returns the PipelineComposition field if non-nil, zero value otherwise.

### GetPipelineCompositionOk

`func (o *CreateIngestionPipelineConfigurationRequest) GetPipelineCompositionOk() (*map[string]string, bool)`

GetPipelineCompositionOk returns a tuple with the PipelineComposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineComposition

`func (o *CreateIngestionPipelineConfigurationRequest) SetPipelineComposition(v map[string]string)`

SetPipelineComposition sets PipelineComposition field to given value.

### HasPipelineComposition

`func (o *CreateIngestionPipelineConfigurationRequest) HasPipelineComposition() bool`

HasPipelineComposition returns a boolean if a field has been set.

### SetPipelineCompositionNil

`func (o *CreateIngestionPipelineConfigurationRequest) SetPipelineCompositionNil(b bool)`

 SetPipelineCompositionNil sets the value for PipelineComposition to be an explicit nil

### UnsetPipelineComposition
`func (o *CreateIngestionPipelineConfigurationRequest) UnsetPipelineComposition()`

UnsetPipelineComposition ensures that no value is present for PipelineComposition, not even an explicit nil
### GetTargetIndexes

`func (o *CreateIngestionPipelineConfigurationRequest) GetTargetIndexes() []IngestionPipelineTargetIndexRef`

GetTargetIndexes returns the TargetIndexes field if non-nil, zero value otherwise.

### GetTargetIndexesOk

`func (o *CreateIngestionPipelineConfigurationRequest) GetTargetIndexesOk() (*[]IngestionPipelineTargetIndexRef, bool)`

GetTargetIndexesOk returns a tuple with the TargetIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIndexes

`func (o *CreateIngestionPipelineConfigurationRequest) SetTargetIndexes(v []IngestionPipelineTargetIndexRef)`

SetTargetIndexes sets TargetIndexes field to given value.

### HasTargetIndexes

`func (o *CreateIngestionPipelineConfigurationRequest) HasTargetIndexes() bool`

HasTargetIndexes returns a boolean if a field has been set.

### SetTargetIndexesNil

`func (o *CreateIngestionPipelineConfigurationRequest) SetTargetIndexesNil(b bool)`

 SetTargetIndexesNil sets the value for TargetIndexes to be an explicit nil

### UnsetTargetIndexes
`func (o *CreateIngestionPipelineConfigurationRequest) UnsetTargetIndexes()`

UnsetTargetIndexes ensures that no value is present for TargetIndexes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



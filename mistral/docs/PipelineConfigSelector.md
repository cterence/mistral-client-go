# PipelineConfigSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceKind** | [**SourceKind**](SourceKind.md) |  | 
**Filter** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPipelineConfigSelector

`func NewPipelineConfigSelector(sourceKind SourceKind, ) *PipelineConfigSelector`

NewPipelineConfigSelector instantiates a new PipelineConfigSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineConfigSelectorWithDefaults

`func NewPipelineConfigSelectorWithDefaults() *PipelineConfigSelector`

NewPipelineConfigSelectorWithDefaults instantiates a new PipelineConfigSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceKind

`func (o *PipelineConfigSelector) GetSourceKind() SourceKind`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *PipelineConfigSelector) GetSourceKindOk() (*SourceKind, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *PipelineConfigSelector) SetSourceKind(v SourceKind)`

SetSourceKind sets SourceKind field to given value.


### GetFilter

`func (o *PipelineConfigSelector) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PipelineConfigSelector) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PipelineConfigSelector) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PipelineConfigSelector) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *PipelineConfigSelector) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *PipelineConfigSelector) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



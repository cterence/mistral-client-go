# PipelineResultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineConfigSlugs** | **[]string** |  | 
**SearchExpression** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPipelineResultsRequest

`func NewPipelineResultsRequest(pipelineConfigSlugs []string, ) *PipelineResultsRequest`

NewPipelineResultsRequest instantiates a new PipelineResultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineResultsRequestWithDefaults

`func NewPipelineResultsRequestWithDefaults() *PipelineResultsRequest`

NewPipelineResultsRequestWithDefaults instantiates a new PipelineResultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineConfigSlugs

`func (o *PipelineResultsRequest) GetPipelineConfigSlugs() []string`

GetPipelineConfigSlugs returns the PipelineConfigSlugs field if non-nil, zero value otherwise.

### GetPipelineConfigSlugsOk

`func (o *PipelineResultsRequest) GetPipelineConfigSlugsOk() (*[]string, bool)`

GetPipelineConfigSlugsOk returns a tuple with the PipelineConfigSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineConfigSlugs

`func (o *PipelineResultsRequest) SetPipelineConfigSlugs(v []string)`

SetPipelineConfigSlugs sets PipelineConfigSlugs field to given value.


### GetSearchExpression

`func (o *PipelineResultsRequest) GetSearchExpression() string`

GetSearchExpression returns the SearchExpression field if non-nil, zero value otherwise.

### GetSearchExpressionOk

`func (o *PipelineResultsRequest) GetSearchExpressionOk() (*string, bool)`

GetSearchExpressionOk returns a tuple with the SearchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchExpression

`func (o *PipelineResultsRequest) SetSearchExpression(v string)`

SetSearchExpression sets SearchExpression field to given value.

### HasSearchExpression

`func (o *PipelineResultsRequest) HasSearchExpression() bool`

HasSearchExpression returns a boolean if a field has been set.

### SetSearchExpressionNil

`func (o *PipelineResultsRequest) SetSearchExpressionNil(b bool)`

 SetSearchExpressionNil sets the value for SearchExpression to be an explicit nil

### UnsetSearchExpression
`func (o *PipelineResultsRequest) UnsetSearchExpression()`

UnsetSearchExpression ensures that no value is present for SearchExpression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



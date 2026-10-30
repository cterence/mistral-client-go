# DetectionDefinitionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttributes** | **[]string** |  | 
**Patterns** | [**[]DetectionPattern**](DetectionPattern.md) |  | 

## Methods

### NewDetectionDefinitionV1

`func NewDetectionDefinitionV1(targetAttributes []string, patterns []DetectionPattern, ) *DetectionDefinitionV1`

NewDetectionDefinitionV1 instantiates a new DetectionDefinitionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectionDefinitionV1WithDefaults

`func NewDetectionDefinitionV1WithDefaults() *DetectionDefinitionV1`

NewDetectionDefinitionV1WithDefaults instantiates a new DetectionDefinitionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttributes

`func (o *DetectionDefinitionV1) GetTargetAttributes() []string`

GetTargetAttributes returns the TargetAttributes field if non-nil, zero value otherwise.

### GetTargetAttributesOk

`func (o *DetectionDefinitionV1) GetTargetAttributesOk() (*[]string, bool)`

GetTargetAttributesOk returns a tuple with the TargetAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributes

`func (o *DetectionDefinitionV1) SetTargetAttributes(v []string)`

SetTargetAttributes sets TargetAttributes field to given value.


### GetPatterns

`func (o *DetectionDefinitionV1) GetPatterns() []DetectionPattern`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *DetectionDefinitionV1) GetPatternsOk() (*[]DetectionPattern, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *DetectionDefinitionV1) SetPatterns(v []DetectionPattern)`

SetPatterns sets Patterns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



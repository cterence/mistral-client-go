# SpanDatasetMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetField** | **string** | Literal top-level key to create in the dataset record payload. | 
**SourceField** | [**SpanDatasetSourceReference**](SpanDatasetSourceReference.md) |  | 

## Methods

### NewSpanDatasetMapping

`func NewSpanDatasetMapping(targetField string, sourceField SpanDatasetSourceReference, ) *SpanDatasetMapping`

NewSpanDatasetMapping instantiates a new SpanDatasetMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanDatasetMappingWithDefaults

`func NewSpanDatasetMappingWithDefaults() *SpanDatasetMapping`

NewSpanDatasetMappingWithDefaults instantiates a new SpanDatasetMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetField

`func (o *SpanDatasetMapping) GetTargetField() string`

GetTargetField returns the TargetField field if non-nil, zero value otherwise.

### GetTargetFieldOk

`func (o *SpanDatasetMapping) GetTargetFieldOk() (*string, bool)`

GetTargetFieldOk returns a tuple with the TargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetField

`func (o *SpanDatasetMapping) SetTargetField(v string)`

SetTargetField sets TargetField field to given value.


### GetSourceField

`func (o *SpanDatasetMapping) GetSourceField() SpanDatasetSourceReference`

GetSourceField returns the SourceField field if non-nil, zero value otherwise.

### GetSourceFieldOk

`func (o *SpanDatasetMapping) GetSourceFieldOk() (*SpanDatasetSourceReference, bool)`

GetSourceFieldOk returns a tuple with the SourceField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceField

`func (o *SpanDatasetMapping) SetSourceField(v SpanDatasetSourceReference)`

SetSourceField sets SourceField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



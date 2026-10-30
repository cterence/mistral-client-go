# SpanDatasetMappingContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Mappings** | [**[]SpanDatasetMapping**](SpanDatasetMapping.md) |  | 

## Methods

### NewSpanDatasetMappingContract

`func NewSpanDatasetMappingContract(version int32, mappings []SpanDatasetMapping, ) *SpanDatasetMappingContract`

NewSpanDatasetMappingContract instantiates a new SpanDatasetMappingContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanDatasetMappingContractWithDefaults

`func NewSpanDatasetMappingContractWithDefaults() *SpanDatasetMappingContract`

NewSpanDatasetMappingContractWithDefaults instantiates a new SpanDatasetMappingContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SpanDatasetMappingContract) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SpanDatasetMappingContract) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SpanDatasetMappingContract) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetMappings

`func (o *SpanDatasetMappingContract) GetMappings() []SpanDatasetMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *SpanDatasetMappingContract) GetMappingsOk() (*[]SpanDatasetMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *SpanDatasetMappingContract) SetMappings(v []SpanDatasetMapping)`

SetMappings sets Mappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



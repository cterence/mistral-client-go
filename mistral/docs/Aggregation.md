# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AggregationRow**](AggregationRow.md) |  | 
**Meta** | [**AggregationMeta**](AggregationMeta.md) |  | 

## Methods

### NewAggregation

`func NewAggregation(data []AggregationRow, meta AggregationMeta, ) *Aggregation`

NewAggregation instantiates a new Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationWithDefaults

`func NewAggregationWithDefaults() *Aggregation`

NewAggregationWithDefaults instantiates a new Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Aggregation) GetData() []AggregationRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Aggregation) GetDataOk() (*[]AggregationRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Aggregation) SetData(v []AggregationRow)`

SetData sets Data field to given value.


### GetMeta

`func (o *Aggregation) GetMeta() AggregationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Aggregation) GetMetaOk() (*AggregationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Aggregation) SetMeta(v AggregationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



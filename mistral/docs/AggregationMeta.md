# AggregationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTimestamp** | **time.Time** |  | 
**ToTimestamp** | **time.Time** |  | 
**GranularitySeconds** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAggregationMeta

`func NewAggregationMeta(fromTimestamp time.Time, toTimestamp time.Time, ) *AggregationMeta`

NewAggregationMeta instantiates a new AggregationMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationMetaWithDefaults

`func NewAggregationMetaWithDefaults() *AggregationMeta`

NewAggregationMetaWithDefaults instantiates a new AggregationMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTimestamp

`func (o *AggregationMeta) GetFromTimestamp() time.Time`

GetFromTimestamp returns the FromTimestamp field if non-nil, zero value otherwise.

### GetFromTimestampOk

`func (o *AggregationMeta) GetFromTimestampOk() (*time.Time, bool)`

GetFromTimestampOk returns a tuple with the FromTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimestamp

`func (o *AggregationMeta) SetFromTimestamp(v time.Time)`

SetFromTimestamp sets FromTimestamp field to given value.


### GetToTimestamp

`func (o *AggregationMeta) GetToTimestamp() time.Time`

GetToTimestamp returns the ToTimestamp field if non-nil, zero value otherwise.

### GetToTimestampOk

`func (o *AggregationMeta) GetToTimestampOk() (*time.Time, bool)`

GetToTimestampOk returns a tuple with the ToTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimestamp

`func (o *AggregationMeta) SetToTimestamp(v time.Time)`

SetToTimestamp sets ToTimestamp field to given value.


### GetGranularitySeconds

`func (o *AggregationMeta) GetGranularitySeconds() int32`

GetGranularitySeconds returns the GranularitySeconds field if non-nil, zero value otherwise.

### GetGranularitySecondsOk

`func (o *AggregationMeta) GetGranularitySecondsOk() (*int32, bool)`

GetGranularitySecondsOk returns a tuple with the GranularitySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularitySeconds

`func (o *AggregationMeta) SetGranularitySeconds(v int32)`

SetGranularitySeconds sets GranularitySeconds field to given value.

### HasGranularitySeconds

`func (o *AggregationMeta) HasGranularitySeconds() bool`

HasGranularitySeconds returns a boolean if a field has been set.

### SetGranularitySecondsNil

`func (o *AggregationMeta) SetGranularitySecondsNil(b bool)`

 SetGranularitySecondsNil sets the value for GranularitySeconds to be an explicit nil

### UnsetGranularitySeconds
`func (o *AggregationMeta) UnsetGranularitySeconds()`

UnsetGranularitySeconds ensures that no value is present for GranularitySeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



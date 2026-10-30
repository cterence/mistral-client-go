# AggregationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeBucket** | Pointer to **NullableTime** |  | [optional] 
**Dimensions** | Pointer to **map[string]interface{}** |  | [optional] 
**MetricName** | **string** |  | 
**MetricValue** | Pointer to [**NullableMetricValue**](MetricValue.md) |  | [optional] 

## Methods

### NewAggregationRow

`func NewAggregationRow(metricName string, ) *AggregationRow`

NewAggregationRow instantiates a new AggregationRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationRowWithDefaults

`func NewAggregationRowWithDefaults() *AggregationRow`

NewAggregationRowWithDefaults instantiates a new AggregationRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeBucket

`func (o *AggregationRow) GetTimeBucket() time.Time`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *AggregationRow) GetTimeBucketOk() (*time.Time, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *AggregationRow) SetTimeBucket(v time.Time)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *AggregationRow) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### SetTimeBucketNil

`func (o *AggregationRow) SetTimeBucketNil(b bool)`

 SetTimeBucketNil sets the value for TimeBucket to be an explicit nil

### UnsetTimeBucket
`func (o *AggregationRow) UnsetTimeBucket()`

UnsetTimeBucket ensures that no value is present for TimeBucket, not even an explicit nil
### GetDimensions

`func (o *AggregationRow) GetDimensions() map[string]interface{}`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AggregationRow) GetDimensionsOk() (*map[string]interface{}, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AggregationRow) SetDimensions(v map[string]interface{})`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *AggregationRow) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetMetricName

`func (o *AggregationRow) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AggregationRow) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AggregationRow) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricValue

`func (o *AggregationRow) GetMetricValue() MetricValue`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *AggregationRow) GetMetricValueOk() (*MetricValue, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *AggregationRow) SetMetricValue(v MetricValue)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *AggregationRow) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.

### SetMetricValueNil

`func (o *AggregationRow) SetMetricValueNil(b bool)`

 SetMetricValueNil sets the value for MetricValue to be an explicit nil

### UnsetMetricValue
`func (o *AggregationRow) UnsetMetricValue()`

UnsetMetricValue ensures that no value is present for MetricValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



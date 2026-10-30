# AggregationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | [**MetricDefinition**](MetricDefinition.md) |  | 
**Dimensions** | Pointer to **[]string** |  | [optional] [default to {}]
**TimeDimension** | Pointer to [**NullableTimeDimension**](TimeDimension.md) |  | [optional] 
**SearchExpression** | Pointer to **NullableString** |  | [optional] 
**OrderBy** | Pointer to [**[]OrderByClause**](OrderByClause.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 1000]

## Methods

### NewAggregationRequest

`func NewAggregationRequest(metric MetricDefinition, ) *AggregationRequest`

NewAggregationRequest instantiates a new AggregationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationRequestWithDefaults

`func NewAggregationRequestWithDefaults() *AggregationRequest`

NewAggregationRequestWithDefaults instantiates a new AggregationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AggregationRequest) GetMetric() MetricDefinition`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AggregationRequest) GetMetricOk() (*MetricDefinition, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AggregationRequest) SetMetric(v MetricDefinition)`

SetMetric sets Metric field to given value.


### GetDimensions

`func (o *AggregationRequest) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AggregationRequest) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AggregationRequest) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *AggregationRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetTimeDimension

`func (o *AggregationRequest) GetTimeDimension() TimeDimension`

GetTimeDimension returns the TimeDimension field if non-nil, zero value otherwise.

### GetTimeDimensionOk

`func (o *AggregationRequest) GetTimeDimensionOk() (*TimeDimension, bool)`

GetTimeDimensionOk returns a tuple with the TimeDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDimension

`func (o *AggregationRequest) SetTimeDimension(v TimeDimension)`

SetTimeDimension sets TimeDimension field to given value.

### HasTimeDimension

`func (o *AggregationRequest) HasTimeDimension() bool`

HasTimeDimension returns a boolean if a field has been set.

### SetTimeDimensionNil

`func (o *AggregationRequest) SetTimeDimensionNil(b bool)`

 SetTimeDimensionNil sets the value for TimeDimension to be an explicit nil

### UnsetTimeDimension
`func (o *AggregationRequest) UnsetTimeDimension()`

UnsetTimeDimension ensures that no value is present for TimeDimension, not even an explicit nil
### GetSearchExpression

`func (o *AggregationRequest) GetSearchExpression() string`

GetSearchExpression returns the SearchExpression field if non-nil, zero value otherwise.

### GetSearchExpressionOk

`func (o *AggregationRequest) GetSearchExpressionOk() (*string, bool)`

GetSearchExpressionOk returns a tuple with the SearchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchExpression

`func (o *AggregationRequest) SetSearchExpression(v string)`

SetSearchExpression sets SearchExpression field to given value.

### HasSearchExpression

`func (o *AggregationRequest) HasSearchExpression() bool`

HasSearchExpression returns a boolean if a field has been set.

### SetSearchExpressionNil

`func (o *AggregationRequest) SetSearchExpressionNil(b bool)`

 SetSearchExpressionNil sets the value for SearchExpression to be an explicit nil

### UnsetSearchExpression
`func (o *AggregationRequest) UnsetSearchExpression()`

UnsetSearchExpression ensures that no value is present for SearchExpression, not even an explicit nil
### GetOrderBy

`func (o *AggregationRequest) GetOrderBy() []OrderByClause`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *AggregationRequest) GetOrderByOk() (*[]OrderByClause, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *AggregationRequest) SetOrderBy(v []OrderByClause)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *AggregationRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *AggregationRequest) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *AggregationRequest) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetLimit

`func (o *AggregationRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AggregationRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AggregationRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AggregationRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



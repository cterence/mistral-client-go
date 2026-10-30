# MetricDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measure** | **string** |  | 
**Aggregation** | [**MetricAggregation**](MetricAggregation.md) |  | 

## Methods

### NewMetricDefinition

`func NewMetricDefinition(measure string, aggregation MetricAggregation, ) *MetricDefinition`

NewMetricDefinition instantiates a new MetricDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDefinitionWithDefaults

`func NewMetricDefinitionWithDefaults() *MetricDefinition`

NewMetricDefinitionWithDefaults instantiates a new MetricDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasure

`func (o *MetricDefinition) GetMeasure() string`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *MetricDefinition) GetMeasureOk() (*string, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *MetricDefinition) SetMeasure(v string)`

SetMeasure sets Measure field to given value.


### GetAggregation

`func (o *MetricDefinition) GetAggregation() MetricAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricDefinition) GetAggregationOk() (*MetricAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricDefinition) SetAggregation(v MetricAggregation)`

SetAggregation sets Aggregation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



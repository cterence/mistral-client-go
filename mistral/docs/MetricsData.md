# MetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**DocumentCount** | **int32** |  | 
**IndexMetrics** | [**[]UpdateMetricsRequestIndexMetrics**](UpdateMetricsRequestIndexMetrics.md) |  | 
**ClearMetrics** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMetricsData

`func NewMetricsData(status string, documentCount int32, indexMetrics []UpdateMetricsRequestIndexMetrics, ) *MetricsData`

NewMetricsData instantiates a new MetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDataWithDefaults

`func NewMetricsDataWithDefaults() *MetricsData`

NewMetricsDataWithDefaults instantiates a new MetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MetricsData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricsData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricsData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDocumentCount

`func (o *MetricsData) GetDocumentCount() int32`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *MetricsData) GetDocumentCountOk() (*int32, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *MetricsData) SetDocumentCount(v int32)`

SetDocumentCount sets DocumentCount field to given value.


### GetIndexMetrics

`func (o *MetricsData) GetIndexMetrics() []UpdateMetricsRequestIndexMetrics`

GetIndexMetrics returns the IndexMetrics field if non-nil, zero value otherwise.

### GetIndexMetricsOk

`func (o *MetricsData) GetIndexMetricsOk() (*[]UpdateMetricsRequestIndexMetrics, bool)`

GetIndexMetricsOk returns a tuple with the IndexMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexMetrics

`func (o *MetricsData) SetIndexMetrics(v []UpdateMetricsRequestIndexMetrics)`

SetIndexMetrics sets IndexMetrics field to given value.


### GetClearMetrics

`func (o *MetricsData) GetClearMetrics() bool`

GetClearMetrics returns the ClearMetrics field if non-nil, zero value otherwise.

### GetClearMetricsOk

`func (o *MetricsData) GetClearMetricsOk() (*bool, bool)`

GetClearMetricsOk returns a tuple with the ClearMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearMetrics

`func (o *MetricsData) SetClearMetrics(v bool)`

SetClearMetrics sets ClearMetrics field to given value.

### HasClearMetrics

`func (o *MetricsData) HasClearMetrics() bool`

HasClearMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateMetricsRequestDeploymentMetricsOnline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**DocumentCount** | **int32** |  | 
**IndexMetrics** | [**[]UpdateMetricsRequestIndexMetrics**](UpdateMetricsRequestIndexMetrics.md) |  | 

## Methods

### NewUpdateMetricsRequestDeploymentMetricsOnline

`func NewUpdateMetricsRequestDeploymentMetricsOnline(status string, documentCount int32, indexMetrics []UpdateMetricsRequestIndexMetrics, ) *UpdateMetricsRequestDeploymentMetricsOnline`

NewUpdateMetricsRequestDeploymentMetricsOnline instantiates a new UpdateMetricsRequestDeploymentMetricsOnline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMetricsRequestDeploymentMetricsOnlineWithDefaults

`func NewUpdateMetricsRequestDeploymentMetricsOnlineWithDefaults() *UpdateMetricsRequestDeploymentMetricsOnline`

NewUpdateMetricsRequestDeploymentMetricsOnlineWithDefaults instantiates a new UpdateMetricsRequestDeploymentMetricsOnline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDocumentCount

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) GetDocumentCount() int32`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) GetDocumentCountOk() (*int32, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) SetDocumentCount(v int32)`

SetDocumentCount sets DocumentCount field to given value.


### GetIndexMetrics

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) GetIndexMetrics() []UpdateMetricsRequestIndexMetrics`

GetIndexMetrics returns the IndexMetrics field if non-nil, zero value otherwise.

### GetIndexMetricsOk

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) GetIndexMetricsOk() (*[]UpdateMetricsRequestIndexMetrics, bool)`

GetIndexMetricsOk returns a tuple with the IndexMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexMetrics

`func (o *UpdateMetricsRequestDeploymentMetricsOnline) SetIndexMetrics(v []UpdateMetricsRequestIndexMetrics)`

SetIndexMetrics sets IndexMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



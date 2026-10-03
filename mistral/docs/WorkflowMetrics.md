# WorkflowMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionCount** | [**ScalarMetric**](ScalarMetric.md) |  | 
**SuccessCount** | [**ScalarMetric**](ScalarMetric.md) |  | 
**ErrorCount** | [**ScalarMetric**](ScalarMetric.md) |  | 
**AverageLatencyMs** | [**ScalarMetric**](ScalarMetric.md) |  | 
**LatencyOverTime** | [**TimeSeriesMetric**](TimeSeriesMetric.md) |  | 
**RetryRate** | [**ScalarMetric**](ScalarMetric.md) |  | 

## Methods

### NewWorkflowMetrics

`func NewWorkflowMetrics(executionCount ScalarMetric, successCount ScalarMetric, errorCount ScalarMetric, averageLatencyMs ScalarMetric, latencyOverTime TimeSeriesMetric, retryRate ScalarMetric, ) *WorkflowMetrics`

NewWorkflowMetrics instantiates a new WorkflowMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMetricsWithDefaults

`func NewWorkflowMetricsWithDefaults() *WorkflowMetrics`

NewWorkflowMetricsWithDefaults instantiates a new WorkflowMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionCount

`func (o *WorkflowMetrics) GetExecutionCount() ScalarMetric`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *WorkflowMetrics) GetExecutionCountOk() (*ScalarMetric, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *WorkflowMetrics) SetExecutionCount(v ScalarMetric)`

SetExecutionCount sets ExecutionCount field to given value.


### GetSuccessCount

`func (o *WorkflowMetrics) GetSuccessCount() ScalarMetric`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *WorkflowMetrics) GetSuccessCountOk() (*ScalarMetric, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *WorkflowMetrics) SetSuccessCount(v ScalarMetric)`

SetSuccessCount sets SuccessCount field to given value.


### GetErrorCount

`func (o *WorkflowMetrics) GetErrorCount() ScalarMetric`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *WorkflowMetrics) GetErrorCountOk() (*ScalarMetric, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *WorkflowMetrics) SetErrorCount(v ScalarMetric)`

SetErrorCount sets ErrorCount field to given value.


### GetAverageLatencyMs

`func (o *WorkflowMetrics) GetAverageLatencyMs() ScalarMetric`

GetAverageLatencyMs returns the AverageLatencyMs field if non-nil, zero value otherwise.

### GetAverageLatencyMsOk

`func (o *WorkflowMetrics) GetAverageLatencyMsOk() (*ScalarMetric, bool)`

GetAverageLatencyMsOk returns a tuple with the AverageLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLatencyMs

`func (o *WorkflowMetrics) SetAverageLatencyMs(v ScalarMetric)`

SetAverageLatencyMs sets AverageLatencyMs field to given value.


### GetLatencyOverTime

`func (o *WorkflowMetrics) GetLatencyOverTime() TimeSeriesMetric`

GetLatencyOverTime returns the LatencyOverTime field if non-nil, zero value otherwise.

### GetLatencyOverTimeOk

`func (o *WorkflowMetrics) GetLatencyOverTimeOk() (*TimeSeriesMetric, bool)`

GetLatencyOverTimeOk returns a tuple with the LatencyOverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyOverTime

`func (o *WorkflowMetrics) SetLatencyOverTime(v TimeSeriesMetric)`

SetLatencyOverTime sets LatencyOverTime field to given value.


### GetRetryRate

`func (o *WorkflowMetrics) GetRetryRate() ScalarMetric`

GetRetryRate returns the RetryRate field if non-nil, zero value otherwise.

### GetRetryRateOk

`func (o *WorkflowMetrics) GetRetryRateOk() (*ScalarMetric, bool)`

GetRetryRateOk returns a tuple with the RetryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryRate

`func (o *WorkflowMetrics) SetRetryRate(v ScalarMetric)`

SetRetryRate sets RetryRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



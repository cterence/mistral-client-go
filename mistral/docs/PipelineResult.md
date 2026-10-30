# PipelineResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**SpanName** | **string** |  | 
**DurationNs** | **int32** |  | 
**EvaluatedAt** | **time.Time** |  | 
**JudgeScores** | **map[string]float32** |  | 
**JudgeLabels** | **map[string]string** |  | 
**JudgeExplanations** | **map[string]string** |  | 

## Methods

### NewPipelineResult

`func NewPipelineResult(traceId string, spanId string, spanName string, durationNs int32, evaluatedAt time.Time, judgeScores map[string]float32, judgeLabels map[string]string, judgeExplanations map[string]string, ) *PipelineResult`

NewPipelineResult instantiates a new PipelineResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineResultWithDefaults

`func NewPipelineResultWithDefaults() *PipelineResult`

NewPipelineResultWithDefaults instantiates a new PipelineResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *PipelineResult) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *PipelineResult) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *PipelineResult) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *PipelineResult) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *PipelineResult) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *PipelineResult) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetSpanName

`func (o *PipelineResult) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *PipelineResult) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *PipelineResult) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.


### GetDurationNs

`func (o *PipelineResult) GetDurationNs() int32`

GetDurationNs returns the DurationNs field if non-nil, zero value otherwise.

### GetDurationNsOk

`func (o *PipelineResult) GetDurationNsOk() (*int32, bool)`

GetDurationNsOk returns a tuple with the DurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNs

`func (o *PipelineResult) SetDurationNs(v int32)`

SetDurationNs sets DurationNs field to given value.


### GetEvaluatedAt

`func (o *PipelineResult) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *PipelineResult) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *PipelineResult) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### GetJudgeScores

`func (o *PipelineResult) GetJudgeScores() map[string]float32`

GetJudgeScores returns the JudgeScores field if non-nil, zero value otherwise.

### GetJudgeScoresOk

`func (o *PipelineResult) GetJudgeScoresOk() (*map[string]float32, bool)`

GetJudgeScoresOk returns a tuple with the JudgeScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeScores

`func (o *PipelineResult) SetJudgeScores(v map[string]float32)`

SetJudgeScores sets JudgeScores field to given value.


### GetJudgeLabels

`func (o *PipelineResult) GetJudgeLabels() map[string]string`

GetJudgeLabels returns the JudgeLabels field if non-nil, zero value otherwise.

### GetJudgeLabelsOk

`func (o *PipelineResult) GetJudgeLabelsOk() (*map[string]string, bool)`

GetJudgeLabelsOk returns a tuple with the JudgeLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeLabels

`func (o *PipelineResult) SetJudgeLabels(v map[string]string)`

SetJudgeLabels sets JudgeLabels field to given value.


### GetJudgeExplanations

`func (o *PipelineResult) GetJudgeExplanations() map[string]string`

GetJudgeExplanations returns the JudgeExplanations field if non-nil, zero value otherwise.

### GetJudgeExplanationsOk

`func (o *PipelineResult) GetJudgeExplanationsOk() (*map[string]string, bool)`

GetJudgeExplanationsOk returns a tuple with the JudgeExplanations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeExplanations

`func (o *PipelineResult) SetJudgeExplanations(v map[string]string)`

SetJudgeExplanations sets JudgeExplanations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



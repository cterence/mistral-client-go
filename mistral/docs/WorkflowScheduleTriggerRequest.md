# WorkflowScheduleTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overlap** | Pointer to [**NullableScheduleOverlapPolicy**](ScheduleOverlapPolicy.md) | Optional overlap policy override to use for the immediate trigger. | [optional] 

## Methods

### NewWorkflowScheduleTriggerRequest

`func NewWorkflowScheduleTriggerRequest() *WorkflowScheduleTriggerRequest`

NewWorkflowScheduleTriggerRequest instantiates a new WorkflowScheduleTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleTriggerRequestWithDefaults

`func NewWorkflowScheduleTriggerRequestWithDefaults() *WorkflowScheduleTriggerRequest`

NewWorkflowScheduleTriggerRequestWithDefaults instantiates a new WorkflowScheduleTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverlap

`func (o *WorkflowScheduleTriggerRequest) GetOverlap() ScheduleOverlapPolicy`

GetOverlap returns the Overlap field if non-nil, zero value otherwise.

### GetOverlapOk

`func (o *WorkflowScheduleTriggerRequest) GetOverlapOk() (*ScheduleOverlapPolicy, bool)`

GetOverlapOk returns a tuple with the Overlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlap

`func (o *WorkflowScheduleTriggerRequest) SetOverlap(v ScheduleOverlapPolicy)`

SetOverlap sets Overlap field to given value.

### HasOverlap

`func (o *WorkflowScheduleTriggerRequest) HasOverlap() bool`

HasOverlap returns a boolean if a field has been set.

### SetOverlapNil

`func (o *WorkflowScheduleTriggerRequest) SetOverlapNil(b bool)`

 SetOverlapNil sets the value for Overlap to be an explicit nil

### UnsetOverlap
`func (o *WorkflowScheduleTriggerRequest) UnsetOverlap()`

UnsetOverlap ensures that no value is present for Overlap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



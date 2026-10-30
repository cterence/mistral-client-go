# WorkflowScheduleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**PartialScheduleDefinition**](PartialScheduleDefinition.md) | Partial schedule definition to update. Unset fields preserve existing values. | 

## Methods

### NewWorkflowScheduleUpdateRequest

`func NewWorkflowScheduleUpdateRequest(schedule PartialScheduleDefinition, ) *WorkflowScheduleUpdateRequest`

NewWorkflowScheduleUpdateRequest instantiates a new WorkflowScheduleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleUpdateRequestWithDefaults

`func NewWorkflowScheduleUpdateRequestWithDefaults() *WorkflowScheduleUpdateRequest`

NewWorkflowScheduleUpdateRequestWithDefaults instantiates a new WorkflowScheduleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *WorkflowScheduleUpdateRequest) GetSchedule() PartialScheduleDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WorkflowScheduleUpdateRequest) GetScheduleOk() (*PartialScheduleDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WorkflowScheduleUpdateRequest) SetSchedule(v PartialScheduleDefinition)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



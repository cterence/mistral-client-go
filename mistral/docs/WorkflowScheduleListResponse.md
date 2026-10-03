# WorkflowScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | [**[]ScheduleDefinitionOutput**](ScheduleDefinitionOutput.md) | A list of workflow schedules | 

## Methods

### NewWorkflowScheduleListResponse

`func NewWorkflowScheduleListResponse(schedules []ScheduleDefinitionOutput, ) *WorkflowScheduleListResponse`

NewWorkflowScheduleListResponse instantiates a new WorkflowScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleListResponseWithDefaults

`func NewWorkflowScheduleListResponseWithDefaults() *WorkflowScheduleListResponse`

NewWorkflowScheduleListResponseWithDefaults instantiates a new WorkflowScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *WorkflowScheduleListResponse) GetSchedules() []ScheduleDefinitionOutput`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *WorkflowScheduleListResponse) GetSchedulesOk() (*[]ScheduleDefinitionOutput, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *WorkflowScheduleListResponse) SetSchedules(v []ScheduleDefinitionOutput)`

SetSchedules sets Schedules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



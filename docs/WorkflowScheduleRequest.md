# WorkflowScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**ScheduleDefinition**](ScheduleDefinition.md) | The schedule definition | 
**WorkflowRegistrationId** | Pointer to **NullableString** | The ID of the workflow registration to schedule | [optional] 
**WorkflowVersionId** | Pointer to **NullableString** | Deprecated: use workflow_registration_id | [optional] 
**WorkflowIdentifier** | Pointer to **NullableString** | The name or ID of the workflow to schedule | [optional] 
**WorkflowTaskQueue** | Pointer to **NullableString** | Deprecated. Use deployment_name instead. | [optional] 
**ScheduleId** | Pointer to **NullableString** | Allows you to specify a custom schedule ID. If not provided, a random ID will be generated. | [optional] 
**DeploymentName** | Pointer to **NullableString** | Name of the deployment to route this schedule to | [optional] 

## Methods

### NewWorkflowScheduleRequest

`func NewWorkflowScheduleRequest(schedule ScheduleDefinition, ) *WorkflowScheduleRequest`

NewWorkflowScheduleRequest instantiates a new WorkflowScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleRequestWithDefaults

`func NewWorkflowScheduleRequestWithDefaults() *WorkflowScheduleRequest`

NewWorkflowScheduleRequestWithDefaults instantiates a new WorkflowScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *WorkflowScheduleRequest) GetSchedule() ScheduleDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WorkflowScheduleRequest) GetScheduleOk() (*ScheduleDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WorkflowScheduleRequest) SetSchedule(v ScheduleDefinition)`

SetSchedule sets Schedule field to given value.


### GetWorkflowRegistrationId

`func (o *WorkflowScheduleRequest) GetWorkflowRegistrationId() string`

GetWorkflowRegistrationId returns the WorkflowRegistrationId field if non-nil, zero value otherwise.

### GetWorkflowRegistrationIdOk

`func (o *WorkflowScheduleRequest) GetWorkflowRegistrationIdOk() (*string, bool)`

GetWorkflowRegistrationIdOk returns a tuple with the WorkflowRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRegistrationId

`func (o *WorkflowScheduleRequest) SetWorkflowRegistrationId(v string)`

SetWorkflowRegistrationId sets WorkflowRegistrationId field to given value.

### HasWorkflowRegistrationId

`func (o *WorkflowScheduleRequest) HasWorkflowRegistrationId() bool`

HasWorkflowRegistrationId returns a boolean if a field has been set.

### SetWorkflowRegistrationIdNil

`func (o *WorkflowScheduleRequest) SetWorkflowRegistrationIdNil(b bool)`

 SetWorkflowRegistrationIdNil sets the value for WorkflowRegistrationId to be an explicit nil

### UnsetWorkflowRegistrationId
`func (o *WorkflowScheduleRequest) UnsetWorkflowRegistrationId()`

UnsetWorkflowRegistrationId ensures that no value is present for WorkflowRegistrationId, not even an explicit nil
### GetWorkflowVersionId

`func (o *WorkflowScheduleRequest) GetWorkflowVersionId() string`

GetWorkflowVersionId returns the WorkflowVersionId field if non-nil, zero value otherwise.

### GetWorkflowVersionIdOk

`func (o *WorkflowScheduleRequest) GetWorkflowVersionIdOk() (*string, bool)`

GetWorkflowVersionIdOk returns a tuple with the WorkflowVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersionId

`func (o *WorkflowScheduleRequest) SetWorkflowVersionId(v string)`

SetWorkflowVersionId sets WorkflowVersionId field to given value.

### HasWorkflowVersionId

`func (o *WorkflowScheduleRequest) HasWorkflowVersionId() bool`

HasWorkflowVersionId returns a boolean if a field has been set.

### SetWorkflowVersionIdNil

`func (o *WorkflowScheduleRequest) SetWorkflowVersionIdNil(b bool)`

 SetWorkflowVersionIdNil sets the value for WorkflowVersionId to be an explicit nil

### UnsetWorkflowVersionId
`func (o *WorkflowScheduleRequest) UnsetWorkflowVersionId()`

UnsetWorkflowVersionId ensures that no value is present for WorkflowVersionId, not even an explicit nil
### GetWorkflowIdentifier

`func (o *WorkflowScheduleRequest) GetWorkflowIdentifier() string`

GetWorkflowIdentifier returns the WorkflowIdentifier field if non-nil, zero value otherwise.

### GetWorkflowIdentifierOk

`func (o *WorkflowScheduleRequest) GetWorkflowIdentifierOk() (*string, bool)`

GetWorkflowIdentifierOk returns a tuple with the WorkflowIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIdentifier

`func (o *WorkflowScheduleRequest) SetWorkflowIdentifier(v string)`

SetWorkflowIdentifier sets WorkflowIdentifier field to given value.

### HasWorkflowIdentifier

`func (o *WorkflowScheduleRequest) HasWorkflowIdentifier() bool`

HasWorkflowIdentifier returns a boolean if a field has been set.

### SetWorkflowIdentifierNil

`func (o *WorkflowScheduleRequest) SetWorkflowIdentifierNil(b bool)`

 SetWorkflowIdentifierNil sets the value for WorkflowIdentifier to be an explicit nil

### UnsetWorkflowIdentifier
`func (o *WorkflowScheduleRequest) UnsetWorkflowIdentifier()`

UnsetWorkflowIdentifier ensures that no value is present for WorkflowIdentifier, not even an explicit nil
### GetWorkflowTaskQueue

`func (o *WorkflowScheduleRequest) GetWorkflowTaskQueue() string`

GetWorkflowTaskQueue returns the WorkflowTaskQueue field if non-nil, zero value otherwise.

### GetWorkflowTaskQueueOk

`func (o *WorkflowScheduleRequest) GetWorkflowTaskQueueOk() (*string, bool)`

GetWorkflowTaskQueueOk returns a tuple with the WorkflowTaskQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskQueue

`func (o *WorkflowScheduleRequest) SetWorkflowTaskQueue(v string)`

SetWorkflowTaskQueue sets WorkflowTaskQueue field to given value.

### HasWorkflowTaskQueue

`func (o *WorkflowScheduleRequest) HasWorkflowTaskQueue() bool`

HasWorkflowTaskQueue returns a boolean if a field has been set.

### SetWorkflowTaskQueueNil

`func (o *WorkflowScheduleRequest) SetWorkflowTaskQueueNil(b bool)`

 SetWorkflowTaskQueueNil sets the value for WorkflowTaskQueue to be an explicit nil

### UnsetWorkflowTaskQueue
`func (o *WorkflowScheduleRequest) UnsetWorkflowTaskQueue()`

UnsetWorkflowTaskQueue ensures that no value is present for WorkflowTaskQueue, not even an explicit nil
### GetScheduleId

`func (o *WorkflowScheduleRequest) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *WorkflowScheduleRequest) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *WorkflowScheduleRequest) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *WorkflowScheduleRequest) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *WorkflowScheduleRequest) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *WorkflowScheduleRequest) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDeploymentName

`func (o *WorkflowScheduleRequest) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *WorkflowScheduleRequest) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *WorkflowScheduleRequest) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *WorkflowScheduleRequest) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### SetDeploymentNameNil

`func (o *WorkflowScheduleRequest) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *WorkflowScheduleRequest) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowRegistrationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowRegistrations** | [**[]WorkflowRegistration**](WorkflowRegistration.md) | A list of workflow registrations | 
**NextCursor** | **NullableString** |  | 
**WorkflowVersions** | [**[]WorkflowRegistration**](WorkflowRegistration.md) | Deprecated: use workflow_registrations | [readonly] 

## Methods

### NewWorkflowRegistrationListResponse

`func NewWorkflowRegistrationListResponse(workflowRegistrations []WorkflowRegistration, nextCursor NullableString, workflowVersions []WorkflowRegistration, ) *WorkflowRegistrationListResponse`

NewWorkflowRegistrationListResponse instantiates a new WorkflowRegistrationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRegistrationListResponseWithDefaults

`func NewWorkflowRegistrationListResponseWithDefaults() *WorkflowRegistrationListResponse`

NewWorkflowRegistrationListResponseWithDefaults instantiates a new WorkflowRegistrationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowRegistrations

`func (o *WorkflowRegistrationListResponse) GetWorkflowRegistrations() []WorkflowRegistration`

GetWorkflowRegistrations returns the WorkflowRegistrations field if non-nil, zero value otherwise.

### GetWorkflowRegistrationsOk

`func (o *WorkflowRegistrationListResponse) GetWorkflowRegistrationsOk() (*[]WorkflowRegistration, bool)`

GetWorkflowRegistrationsOk returns a tuple with the WorkflowRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRegistrations

`func (o *WorkflowRegistrationListResponse) SetWorkflowRegistrations(v []WorkflowRegistration)`

SetWorkflowRegistrations sets WorkflowRegistrations field to given value.


### GetNextCursor

`func (o *WorkflowRegistrationListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *WorkflowRegistrationListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *WorkflowRegistrationListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *WorkflowRegistrationListResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *WorkflowRegistrationListResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetWorkflowVersions

`func (o *WorkflowRegistrationListResponse) GetWorkflowVersions() []WorkflowRegistration`

GetWorkflowVersions returns the WorkflowVersions field if non-nil, zero value otherwise.

### GetWorkflowVersionsOk

`func (o *WorkflowRegistrationListResponse) GetWorkflowVersionsOk() (*[]WorkflowRegistration, bool)`

GetWorkflowVersionsOk returns a tuple with the WorkflowVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersions

`func (o *WorkflowRegistrationListResponse) SetWorkflowVersions(v []WorkflowRegistration)`

SetWorkflowVersions sets WorkflowVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



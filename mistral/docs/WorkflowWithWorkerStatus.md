# WorkflowWithWorkerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the workflow | 
**Name** | **string** | Name of the workflow | 
**DisplayName** | **string** | Display name of the workflow | 
**Type** | [**WorkflowType**](WorkflowType.md) | Type of the workflow | 
**Description** | Pointer to **NullableString** | Description of the workflow | [optional] 
**CustomerId** | **string** | Customer ID of the workflow | 
**WorkspaceId** | **string** | Workspace ID of the workflow | 
**SharedNamespace** | Pointer to **NullableString** | Reserved namespace for shared workflows (e.g., &#39;shared:my-shared-workflow&#39;) | [optional] 
**AvailableInChatAssistant** | Pointer to **bool** | Whether the workflow is available in chat assistant | [optional] [default to false]
**IsTechnical** | Pointer to **bool** | Whether the workflow is technical (e.g. SDK-managed) | [optional] [default to false]
**Archived** | Pointer to **bool** | Whether the workflow is archived | [optional] [default to false]
**Active** | **bool** | Whether the workflow is active | 

## Methods

### NewWorkflowWithWorkerStatus

`func NewWorkflowWithWorkerStatus(id string, name string, displayName string, type_ WorkflowType, customerId string, workspaceId string, active bool, ) *WorkflowWithWorkerStatus`

NewWorkflowWithWorkerStatus instantiates a new WorkflowWithWorkerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithWorkerStatusWithDefaults

`func NewWorkflowWithWorkerStatusWithDefaults() *WorkflowWithWorkerStatus`

NewWorkflowWithWorkerStatusWithDefaults instantiates a new WorkflowWithWorkerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowWithWorkerStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowWithWorkerStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowWithWorkerStatus) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowWithWorkerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWithWorkerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWithWorkerStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *WorkflowWithWorkerStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowWithWorkerStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowWithWorkerStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetType

`func (o *WorkflowWithWorkerStatus) GetType() WorkflowType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWithWorkerStatus) GetTypeOk() (*WorkflowType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWithWorkerStatus) SetType(v WorkflowType)`

SetType sets Type field to given value.


### GetDescription

`func (o *WorkflowWithWorkerStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWithWorkerStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWithWorkerStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWithWorkerStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowWithWorkerStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowWithWorkerStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomerId

`func (o *WorkflowWithWorkerStatus) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *WorkflowWithWorkerStatus) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *WorkflowWithWorkerStatus) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetWorkspaceId

`func (o *WorkflowWithWorkerStatus) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkflowWithWorkerStatus) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkflowWithWorkerStatus) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSharedNamespace

`func (o *WorkflowWithWorkerStatus) GetSharedNamespace() string`

GetSharedNamespace returns the SharedNamespace field if non-nil, zero value otherwise.

### GetSharedNamespaceOk

`func (o *WorkflowWithWorkerStatus) GetSharedNamespaceOk() (*string, bool)`

GetSharedNamespaceOk returns a tuple with the SharedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedNamespace

`func (o *WorkflowWithWorkerStatus) SetSharedNamespace(v string)`

SetSharedNamespace sets SharedNamespace field to given value.

### HasSharedNamespace

`func (o *WorkflowWithWorkerStatus) HasSharedNamespace() bool`

HasSharedNamespace returns a boolean if a field has been set.

### SetSharedNamespaceNil

`func (o *WorkflowWithWorkerStatus) SetSharedNamespaceNil(b bool)`

 SetSharedNamespaceNil sets the value for SharedNamespace to be an explicit nil

### UnsetSharedNamespace
`func (o *WorkflowWithWorkerStatus) UnsetSharedNamespace()`

UnsetSharedNamespace ensures that no value is present for SharedNamespace, not even an explicit nil
### GetAvailableInChatAssistant

`func (o *WorkflowWithWorkerStatus) GetAvailableInChatAssistant() bool`

GetAvailableInChatAssistant returns the AvailableInChatAssistant field if non-nil, zero value otherwise.

### GetAvailableInChatAssistantOk

`func (o *WorkflowWithWorkerStatus) GetAvailableInChatAssistantOk() (*bool, bool)`

GetAvailableInChatAssistantOk returns a tuple with the AvailableInChatAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableInChatAssistant

`func (o *WorkflowWithWorkerStatus) SetAvailableInChatAssistant(v bool)`

SetAvailableInChatAssistant sets AvailableInChatAssistant field to given value.

### HasAvailableInChatAssistant

`func (o *WorkflowWithWorkerStatus) HasAvailableInChatAssistant() bool`

HasAvailableInChatAssistant returns a boolean if a field has been set.

### GetIsTechnical

`func (o *WorkflowWithWorkerStatus) GetIsTechnical() bool`

GetIsTechnical returns the IsTechnical field if non-nil, zero value otherwise.

### GetIsTechnicalOk

`func (o *WorkflowWithWorkerStatus) GetIsTechnicalOk() (*bool, bool)`

GetIsTechnicalOk returns a tuple with the IsTechnical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnical

`func (o *WorkflowWithWorkerStatus) SetIsTechnical(v bool)`

SetIsTechnical sets IsTechnical field to given value.

### HasIsTechnical

`func (o *WorkflowWithWorkerStatus) HasIsTechnical() bool`

HasIsTechnical returns a boolean if a field has been set.

### GetArchived

`func (o *WorkflowWithWorkerStatus) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkflowWithWorkerStatus) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkflowWithWorkerStatus) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *WorkflowWithWorkerStatus) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetActive

`func (o *WorkflowWithWorkerStatus) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkflowWithWorkerStatus) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkflowWithWorkerStatus) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



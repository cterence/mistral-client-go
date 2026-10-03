# WorkflowBasicDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The name of the workflow | 
**DisplayName** | **string** | The display name of the workflow | 
**Description** | Pointer to **NullableString** | A description of the workflow | [optional] 
**Metadata** | Pointer to [**WorkflowMetadata**](WorkflowMetadata.md) | Workflow metadata | [optional] 
**Archived** | **bool** | Whether the workflow is archived | 

## Methods

### NewWorkflowBasicDefinition

`func NewWorkflowBasicDefinition(id string, name string, displayName string, archived bool, ) *WorkflowBasicDefinition`

NewWorkflowBasicDefinition instantiates a new WorkflowBasicDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBasicDefinitionWithDefaults

`func NewWorkflowBasicDefinitionWithDefaults() *WorkflowBasicDefinition`

NewWorkflowBasicDefinitionWithDefaults instantiates a new WorkflowBasicDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowBasicDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowBasicDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowBasicDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowBasicDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBasicDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBasicDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *WorkflowBasicDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowBasicDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowBasicDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *WorkflowBasicDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBasicDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBasicDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBasicDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowBasicDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowBasicDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *WorkflowBasicDefinition) GetMetadata() WorkflowMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkflowBasicDefinition) GetMetadataOk() (*WorkflowMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkflowBasicDefinition) SetMetadata(v WorkflowMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WorkflowBasicDefinition) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetArchived

`func (o *WorkflowBasicDefinition) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkflowBasicDefinition) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkflowBasicDefinition) SetArchived(v bool)`

SetArchived sets Archived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



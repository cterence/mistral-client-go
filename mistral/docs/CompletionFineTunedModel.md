# CompletionFineTunedModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | **int32** |  | 
**OwnedBy** | **string** |  | 
**WorkspaceId** | **string** |  | 
**Root** | **string** |  | 
**RootVersion** | **string** |  | 
**Archived** | **bool** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Capabilities** | [**FineTunedModelCapabilities**](FineTunedModelCapabilities.md) |  | 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to {}]
**Job** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **string** |  | [optional] [default to "completion"]

## Methods

### NewCompletionFineTunedModel

`func NewCompletionFineTunedModel(id string, created int32, ownedBy string, workspaceId string, root string, rootVersion string, archived bool, capabilities FineTunedModelCapabilities, ) *CompletionFineTunedModel`

NewCompletionFineTunedModel instantiates a new CompletionFineTunedModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionFineTunedModelWithDefaults

`func NewCompletionFineTunedModelWithDefaults() *CompletionFineTunedModel`

NewCompletionFineTunedModelWithDefaults instantiates a new CompletionFineTunedModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompletionFineTunedModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompletionFineTunedModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompletionFineTunedModel) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CompletionFineTunedModel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CompletionFineTunedModel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CompletionFineTunedModel) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CompletionFineTunedModel) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *CompletionFineTunedModel) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CompletionFineTunedModel) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CompletionFineTunedModel) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *CompletionFineTunedModel) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *CompletionFineTunedModel) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *CompletionFineTunedModel) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetWorkspaceId

`func (o *CompletionFineTunedModel) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *CompletionFineTunedModel) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *CompletionFineTunedModel) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRoot

`func (o *CompletionFineTunedModel) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *CompletionFineTunedModel) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *CompletionFineTunedModel) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetRootVersion

`func (o *CompletionFineTunedModel) GetRootVersion() string`

GetRootVersion returns the RootVersion field if non-nil, zero value otherwise.

### GetRootVersionOk

`func (o *CompletionFineTunedModel) GetRootVersionOk() (*string, bool)`

GetRootVersionOk returns a tuple with the RootVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVersion

`func (o *CompletionFineTunedModel) SetRootVersion(v string)`

SetRootVersion sets RootVersion field to given value.


### GetArchived

`func (o *CompletionFineTunedModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CompletionFineTunedModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CompletionFineTunedModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *CompletionFineTunedModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompletionFineTunedModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompletionFineTunedModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompletionFineTunedModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CompletionFineTunedModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompletionFineTunedModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CompletionFineTunedModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompletionFineTunedModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompletionFineTunedModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompletionFineTunedModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CompletionFineTunedModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CompletionFineTunedModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapabilities

`func (o *CompletionFineTunedModel) GetCapabilities() FineTunedModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CompletionFineTunedModel) GetCapabilitiesOk() (*FineTunedModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CompletionFineTunedModel) SetCapabilities(v FineTunedModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *CompletionFineTunedModel) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *CompletionFineTunedModel) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *CompletionFineTunedModel) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *CompletionFineTunedModel) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *CompletionFineTunedModel) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CompletionFineTunedModel) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CompletionFineTunedModel) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CompletionFineTunedModel) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *CompletionFineTunedModel) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CompletionFineTunedModel) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CompletionFineTunedModel) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *CompletionFineTunedModel) HasJob() bool`

HasJob returns a boolean if a field has been set.

### SetJobNil

`func (o *CompletionFineTunedModel) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *CompletionFineTunedModel) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetModelType

`func (o *CompletionFineTunedModel) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *CompletionFineTunedModel) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *CompletionFineTunedModel) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *CompletionFineTunedModel) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



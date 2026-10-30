# ClassifierFineTunedModel

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
**ClassifierTargets** | [**[]ClassifierTargetResult**](ClassifierTargetResult.md) |  | 
**ModelType** | Pointer to **string** |  | [optional] [default to "classifier"]

## Methods

### NewClassifierFineTunedModel

`func NewClassifierFineTunedModel(id string, created int32, ownedBy string, workspaceId string, root string, rootVersion string, archived bool, capabilities FineTunedModelCapabilities, classifierTargets []ClassifierTargetResult, ) *ClassifierFineTunedModel`

NewClassifierFineTunedModel instantiates a new ClassifierFineTunedModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifierFineTunedModelWithDefaults

`func NewClassifierFineTunedModelWithDefaults() *ClassifierFineTunedModel`

NewClassifierFineTunedModelWithDefaults instantiates a new ClassifierFineTunedModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClassifierFineTunedModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClassifierFineTunedModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClassifierFineTunedModel) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ClassifierFineTunedModel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ClassifierFineTunedModel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ClassifierFineTunedModel) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ClassifierFineTunedModel) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *ClassifierFineTunedModel) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClassifierFineTunedModel) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClassifierFineTunedModel) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *ClassifierFineTunedModel) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ClassifierFineTunedModel) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ClassifierFineTunedModel) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetWorkspaceId

`func (o *ClassifierFineTunedModel) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ClassifierFineTunedModel) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ClassifierFineTunedModel) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRoot

`func (o *ClassifierFineTunedModel) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ClassifierFineTunedModel) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ClassifierFineTunedModel) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetRootVersion

`func (o *ClassifierFineTunedModel) GetRootVersion() string`

GetRootVersion returns the RootVersion field if non-nil, zero value otherwise.

### GetRootVersionOk

`func (o *ClassifierFineTunedModel) GetRootVersionOk() (*string, bool)`

GetRootVersionOk returns a tuple with the RootVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVersion

`func (o *ClassifierFineTunedModel) SetRootVersion(v string)`

SetRootVersion sets RootVersion field to given value.


### GetArchived

`func (o *ClassifierFineTunedModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ClassifierFineTunedModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ClassifierFineTunedModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *ClassifierFineTunedModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClassifierFineTunedModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClassifierFineTunedModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClassifierFineTunedModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClassifierFineTunedModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClassifierFineTunedModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ClassifierFineTunedModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClassifierFineTunedModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClassifierFineTunedModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClassifierFineTunedModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClassifierFineTunedModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClassifierFineTunedModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapabilities

`func (o *ClassifierFineTunedModel) GetCapabilities() FineTunedModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ClassifierFineTunedModel) GetCapabilitiesOk() (*FineTunedModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ClassifierFineTunedModel) SetCapabilities(v FineTunedModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *ClassifierFineTunedModel) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *ClassifierFineTunedModel) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *ClassifierFineTunedModel) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *ClassifierFineTunedModel) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *ClassifierFineTunedModel) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ClassifierFineTunedModel) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ClassifierFineTunedModel) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ClassifierFineTunedModel) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *ClassifierFineTunedModel) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ClassifierFineTunedModel) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ClassifierFineTunedModel) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *ClassifierFineTunedModel) HasJob() bool`

HasJob returns a boolean if a field has been set.

### SetJobNil

`func (o *ClassifierFineTunedModel) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *ClassifierFineTunedModel) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetClassifierTargets

`func (o *ClassifierFineTunedModel) GetClassifierTargets() []ClassifierTargetResult`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *ClassifierFineTunedModel) GetClassifierTargetsOk() (*[]ClassifierTargetResult, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *ClassifierFineTunedModel) SetClassifierTargets(v []ClassifierTargetResult)`

SetClassifierTargets sets ClassifierTargets field to given value.


### GetModelType

`func (o *ClassifierFineTunedModel) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ClassifierFineTunedModel) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ClassifierFineTunedModel) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ClassifierFineTunedModel) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Judge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**OwnerId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**ModelName** | **string** |  | 
**Output** | [**Output**](Output.md) |  | 
**Instructions** | **string** |  | 
**Tools** | **[]string** |  | 
**UpRevision** | Pointer to **NullableString** |  | [optional] 
**DownRevision** | Pointer to **NullableString** |  | [optional] 
**BaseRevision** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJudge

`func NewJudge(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, ownerId string, workspaceId string, name string, description string, modelName string, output Output, instructions string, tools []string, ) *Judge`

NewJudge instantiates a new Judge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeWithDefaults

`func NewJudgeWithDefaults() *Judge`

NewJudgeWithDefaults instantiates a new Judge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Judge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Judge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Judge) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Judge) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Judge) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Judge) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Judge) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Judge) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Judge) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Judge) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Judge) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Judge) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *Judge) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Judge) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOwnerId

`func (o *Judge) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Judge) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Judge) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetWorkspaceId

`func (o *Judge) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Judge) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Judge) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *Judge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Judge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Judge) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Judge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Judge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Judge) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelName

`func (o *Judge) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Judge) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Judge) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOutput

`func (o *Judge) GetOutput() Output`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Judge) GetOutputOk() (*Output, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Judge) SetOutput(v Output)`

SetOutput sets Output field to given value.


### GetInstructions

`func (o *Judge) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Judge) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Judge) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetTools

`func (o *Judge) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Judge) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Judge) SetTools(v []string)`

SetTools sets Tools field to given value.


### GetUpRevision

`func (o *Judge) GetUpRevision() string`

GetUpRevision returns the UpRevision field if non-nil, zero value otherwise.

### GetUpRevisionOk

`func (o *Judge) GetUpRevisionOk() (*string, bool)`

GetUpRevisionOk returns a tuple with the UpRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRevision

`func (o *Judge) SetUpRevision(v string)`

SetUpRevision sets UpRevision field to given value.

### HasUpRevision

`func (o *Judge) HasUpRevision() bool`

HasUpRevision returns a boolean if a field has been set.

### SetUpRevisionNil

`func (o *Judge) SetUpRevisionNil(b bool)`

 SetUpRevisionNil sets the value for UpRevision to be an explicit nil

### UnsetUpRevision
`func (o *Judge) UnsetUpRevision()`

UnsetUpRevision ensures that no value is present for UpRevision, not even an explicit nil
### GetDownRevision

`func (o *Judge) GetDownRevision() string`

GetDownRevision returns the DownRevision field if non-nil, zero value otherwise.

### GetDownRevisionOk

`func (o *Judge) GetDownRevisionOk() (*string, bool)`

GetDownRevisionOk returns a tuple with the DownRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRevision

`func (o *Judge) SetDownRevision(v string)`

SetDownRevision sets DownRevision field to given value.

### HasDownRevision

`func (o *Judge) HasDownRevision() bool`

HasDownRevision returns a boolean if a field has been set.

### SetDownRevisionNil

`func (o *Judge) SetDownRevisionNil(b bool)`

 SetDownRevisionNil sets the value for DownRevision to be an explicit nil

### UnsetDownRevision
`func (o *Judge) UnsetDownRevision()`

UnsetDownRevision ensures that no value is present for DownRevision, not even an explicit nil
### GetBaseRevision

`func (o *Judge) GetBaseRevision() string`

GetBaseRevision returns the BaseRevision field if non-nil, zero value otherwise.

### GetBaseRevisionOk

`func (o *Judge) GetBaseRevisionOk() (*string, bool)`

GetBaseRevisionOk returns a tuple with the BaseRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRevision

`func (o *Judge) SetBaseRevision(v string)`

SetBaseRevision sets BaseRevision field to given value.

### HasBaseRevision

`func (o *Judge) HasBaseRevision() bool`

HasBaseRevision returns a boolean if a field has been set.

### SetBaseRevisionNil

`func (o *Judge) SetBaseRevisionNil(b bool)`

 SetBaseRevisionNil sets the value for BaseRevision to be an explicit nil

### UnsetBaseRevision
`func (o *Judge) UnsetBaseRevision()`

UnsetBaseRevision ensures that no value is present for BaseRevision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



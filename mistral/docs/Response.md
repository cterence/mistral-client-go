# Response

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
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Capabilities** | [**FineTunedModelCapabilities**](FineTunedModelCapabilities.md) |  | 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to {}]
**Job** | Pointer to **string** |  | [optional] 
**ModelType** | Pointer to **string** |  | [optional] [default to "classifier"]
**ClassifierTargets** | [**[]ClassifierTargetResult**](ClassifierTargetResult.md) |  | 

## Methods

### NewResponse

`func NewResponse(id string, created int32, ownedBy string, workspaceId string, root string, rootVersion string, archived bool, capabilities FineTunedModelCapabilities, classifierTargets []ClassifierTargetResult, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Response) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Response) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Response) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Response) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Response) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *Response) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Response) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Response) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *Response) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *Response) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *Response) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetWorkspaceId

`func (o *Response) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Response) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Response) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRoot

`func (o *Response) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Response) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *Response) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetRootVersion

`func (o *Response) GetRootVersion() string`

GetRootVersion returns the RootVersion field if non-nil, zero value otherwise.

### GetRootVersionOk

`func (o *Response) GetRootVersionOk() (*string, bool)`

GetRootVersionOk returns a tuple with the RootVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVersion

`func (o *Response) SetRootVersion(v string)`

SetRootVersion sets RootVersion field to given value.


### GetArchived

`func (o *Response) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Response) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Response) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCapabilities

`func (o *Response) GetCapabilities() FineTunedModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Response) GetCapabilitiesOk() (*FineTunedModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Response) SetCapabilities(v FineTunedModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetMaxContextLength

`func (o *Response) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *Response) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *Response) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *Response) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *Response) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Response) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Response) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Response) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetJob

`func (o *Response) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Response) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Response) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetModelType

`func (o *Response) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *Response) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *Response) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *Response) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetClassifierTargets

`func (o *Response) GetClassifierTargets() []ClassifierTargetResult`

GetClassifierTargets returns the ClassifierTargets field if non-nil, zero value otherwise.

### GetClassifierTargetsOk

`func (o *Response) GetClassifierTargetsOk() (*[]ClassifierTargetResult, bool)`

GetClassifierTargetsOk returns a tuple with the ClassifierTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifierTargets

`func (o *Response) SetClassifierTargets(v []ClassifierTargetResult)`

SetClassifierTargets sets ClassifierTargets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



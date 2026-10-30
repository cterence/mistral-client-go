# CreateLibraryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ChunkSize** | Pointer to **NullableInt32** | The size of the chunks (in characters) to split document text into. Must be between 256 and 32768. | [optional] 
**OwnerType** | Pointer to **NullableString** | Determines who owns the created library. &#39;User&#39; creates a private library accessible only to its owner. &#39;Workspace&#39; creates a library shared with the workspace. Defaults to &#39;Workspace&#39; for API key sessions. Only API keys with the &#39;Private and shared connectors&#39; connector access scope can create private, user-owned libraries. | [optional] 

## Methods

### NewCreateLibraryRequest

`func NewCreateLibraryRequest(name string, ) *CreateLibraryRequest`

NewCreateLibraryRequest instantiates a new CreateLibraryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLibraryRequestWithDefaults

`func NewCreateLibraryRequestWithDefaults() *CreateLibraryRequest`

NewCreateLibraryRequestWithDefaults instantiates a new CreateLibraryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLibraryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLibraryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLibraryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateLibraryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLibraryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLibraryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLibraryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLibraryRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLibraryRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChunkSize

`func (o *CreateLibraryRequest) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *CreateLibraryRequest) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *CreateLibraryRequest) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *CreateLibraryRequest) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSizeNil

`func (o *CreateLibraryRequest) SetChunkSizeNil(b bool)`

 SetChunkSizeNil sets the value for ChunkSize to be an explicit nil

### UnsetChunkSize
`func (o *CreateLibraryRequest) UnsetChunkSize()`

UnsetChunkSize ensures that no value is present for ChunkSize, not even an explicit nil
### GetOwnerType

`func (o *CreateLibraryRequest) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *CreateLibraryRequest) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *CreateLibraryRequest) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *CreateLibraryRequest) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### SetOwnerTypeNil

`func (o *CreateLibraryRequest) SetOwnerTypeNil(b bool)`

 SetOwnerTypeNil sets the value for OwnerType to be an explicit nil

### UnsetOwnerType
`func (o *CreateLibraryRequest) UnsetOwnerType()`

UnsetOwnerType ensures that no value is present for OwnerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



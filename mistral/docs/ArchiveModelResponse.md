# ArchiveModelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Archived** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewArchiveModelResponse

`func NewArchiveModelResponse(id string, ) *ArchiveModelResponse`

NewArchiveModelResponse instantiates a new ArchiveModelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveModelResponseWithDefaults

`func NewArchiveModelResponseWithDefaults() *ArchiveModelResponse`

NewArchiveModelResponseWithDefaults instantiates a new ArchiveModelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveModelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveModelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveModelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ArchiveModelResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ArchiveModelResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ArchiveModelResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ArchiveModelResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetArchived

`func (o *ArchiveModelResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ArchiveModelResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ArchiveModelResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ArchiveModelResponse) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DatasetImportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**CreatorId** | **string** |  | 
**DatasetId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**Status** | [**BaseTaskStatus**](BaseTaskStatus.md) |  | 
**Progress** | Pointer to **NullableInt32** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatasetImportTask

`func NewDatasetImportTask(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, creatorId string, datasetId string, workspaceId string, status BaseTaskStatus, ) *DatasetImportTask`

NewDatasetImportTask instantiates a new DatasetImportTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetImportTaskWithDefaults

`func NewDatasetImportTaskWithDefaults() *DatasetImportTask`

NewDatasetImportTaskWithDefaults instantiates a new DatasetImportTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetImportTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetImportTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetImportTask) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DatasetImportTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetImportTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetImportTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DatasetImportTask) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatasetImportTask) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatasetImportTask) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DatasetImportTask) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DatasetImportTask) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DatasetImportTask) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *DatasetImportTask) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DatasetImportTask) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatorId

`func (o *DatasetImportTask) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *DatasetImportTask) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *DatasetImportTask) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetDatasetId

`func (o *DatasetImportTask) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *DatasetImportTask) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *DatasetImportTask) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetWorkspaceId

`func (o *DatasetImportTask) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DatasetImportTask) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DatasetImportTask) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetStatus

`func (o *DatasetImportTask) GetStatus() BaseTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasetImportTask) GetStatusOk() (*BaseTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasetImportTask) SetStatus(v BaseTaskStatus)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *DatasetImportTask) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DatasetImportTask) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DatasetImportTask) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DatasetImportTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *DatasetImportTask) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *DatasetImportTask) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetMessage

`func (o *DatasetImportTask) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DatasetImportTask) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DatasetImportTask) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DatasetImportTask) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DatasetImportTask) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DatasetImportTask) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



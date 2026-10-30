# IngestionPipelineConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AuthorId** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**LastRunTime** | **NullableTime** |  | 
**LastRunChunksCount** | **int32** |  | 
**TotalChunksCount** | **int32** |  | 
**PipelineComposition** | **map[string]string** |  | 

## Methods

### NewIngestionPipelineConfiguration

`func NewIngestionPipelineConfiguration(id string, authorId string, name string, createdAt time.Time, modifiedAt time.Time, lastRunTime NullableTime, lastRunChunksCount int32, totalChunksCount int32, pipelineComposition map[string]string, ) *IngestionPipelineConfiguration`

NewIngestionPipelineConfiguration instantiates a new IngestionPipelineConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionPipelineConfigurationWithDefaults

`func NewIngestionPipelineConfigurationWithDefaults() *IngestionPipelineConfiguration`

NewIngestionPipelineConfigurationWithDefaults instantiates a new IngestionPipelineConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestionPipelineConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionPipelineConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionPipelineConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorId

`func (o *IngestionPipelineConfiguration) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *IngestionPipelineConfiguration) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *IngestionPipelineConfiguration) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetName

`func (o *IngestionPipelineConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestionPipelineConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestionPipelineConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *IngestionPipelineConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngestionPipelineConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngestionPipelineConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *IngestionPipelineConfiguration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *IngestionPipelineConfiguration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *IngestionPipelineConfiguration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetLastRunTime

`func (o *IngestionPipelineConfiguration) GetLastRunTime() time.Time`

GetLastRunTime returns the LastRunTime field if non-nil, zero value otherwise.

### GetLastRunTimeOk

`func (o *IngestionPipelineConfiguration) GetLastRunTimeOk() (*time.Time, bool)`

GetLastRunTimeOk returns a tuple with the LastRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunTime

`func (o *IngestionPipelineConfiguration) SetLastRunTime(v time.Time)`

SetLastRunTime sets LastRunTime field to given value.


### SetLastRunTimeNil

`func (o *IngestionPipelineConfiguration) SetLastRunTimeNil(b bool)`

 SetLastRunTimeNil sets the value for LastRunTime to be an explicit nil

### UnsetLastRunTime
`func (o *IngestionPipelineConfiguration) UnsetLastRunTime()`

UnsetLastRunTime ensures that no value is present for LastRunTime, not even an explicit nil
### GetLastRunChunksCount

`func (o *IngestionPipelineConfiguration) GetLastRunChunksCount() int32`

GetLastRunChunksCount returns the LastRunChunksCount field if non-nil, zero value otherwise.

### GetLastRunChunksCountOk

`func (o *IngestionPipelineConfiguration) GetLastRunChunksCountOk() (*int32, bool)`

GetLastRunChunksCountOk returns a tuple with the LastRunChunksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunChunksCount

`func (o *IngestionPipelineConfiguration) SetLastRunChunksCount(v int32)`

SetLastRunChunksCount sets LastRunChunksCount field to given value.


### GetTotalChunksCount

`func (o *IngestionPipelineConfiguration) GetTotalChunksCount() int32`

GetTotalChunksCount returns the TotalChunksCount field if non-nil, zero value otherwise.

### GetTotalChunksCountOk

`func (o *IngestionPipelineConfiguration) GetTotalChunksCountOk() (*int32, bool)`

GetTotalChunksCountOk returns a tuple with the TotalChunksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChunksCount

`func (o *IngestionPipelineConfiguration) SetTotalChunksCount(v int32)`

SetTotalChunksCount sets TotalChunksCount field to given value.


### GetPipelineComposition

`func (o *IngestionPipelineConfiguration) GetPipelineComposition() map[string]string`

GetPipelineComposition returns the PipelineComposition field if non-nil, zero value otherwise.

### GetPipelineCompositionOk

`func (o *IngestionPipelineConfiguration) GetPipelineCompositionOk() (*map[string]string, bool)`

GetPipelineCompositionOk returns a tuple with the PipelineComposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineComposition

`func (o *IngestionPipelineConfiguration) SetPipelineComposition(v map[string]string)`

SetPipelineComposition sets PipelineComposition field to given value.


### SetPipelineCompositionNil

`func (o *IngestionPipelineConfiguration) SetPipelineCompositionNil(b bool)`

 SetPipelineCompositionNil sets the value for PipelineComposition to be an explicit nil

### UnsetPipelineComposition
`func (o *IngestionPipelineConfiguration) UnsetPipelineComposition()`

UnsetPipelineComposition ensures that no value is present for PipelineComposition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



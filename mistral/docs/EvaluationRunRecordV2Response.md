# EvaluationRunRecordV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Input** | **interface{}** |  | 
**DatasetRecordId** | Pointer to **NullableString** |  | [optional] 
**OutputRecord** | Pointer to [**NullableEvaluationRunOutputRecordV2**](EvaluationRunOutputRecordV2.md) |  | [optional] 

## Methods

### NewEvaluationRunRecordV2Response

`func NewEvaluationRunRecordV2Response(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, input interface{}, ) *EvaluationRunRecordV2Response`

NewEvaluationRunRecordV2Response instantiates a new EvaluationRunRecordV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunRecordV2ResponseWithDefaults

`func NewEvaluationRunRecordV2ResponseWithDefaults() *EvaluationRunRecordV2Response`

NewEvaluationRunRecordV2ResponseWithDefaults instantiates a new EvaluationRunRecordV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluationRunRecordV2Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluationRunRecordV2Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluationRunRecordV2Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EvaluationRunRecordV2Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvaluationRunRecordV2Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvaluationRunRecordV2Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EvaluationRunRecordV2Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvaluationRunRecordV2Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvaluationRunRecordV2Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EvaluationRunRecordV2Response) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EvaluationRunRecordV2Response) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EvaluationRunRecordV2Response) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *EvaluationRunRecordV2Response) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EvaluationRunRecordV2Response) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetInput

`func (o *EvaluationRunRecordV2Response) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EvaluationRunRecordV2Response) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EvaluationRunRecordV2Response) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *EvaluationRunRecordV2Response) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *EvaluationRunRecordV2Response) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetDatasetRecordId

`func (o *EvaluationRunRecordV2Response) GetDatasetRecordId() string`

GetDatasetRecordId returns the DatasetRecordId field if non-nil, zero value otherwise.

### GetDatasetRecordIdOk

`func (o *EvaluationRunRecordV2Response) GetDatasetRecordIdOk() (*string, bool)`

GetDatasetRecordIdOk returns a tuple with the DatasetRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetRecordId

`func (o *EvaluationRunRecordV2Response) SetDatasetRecordId(v string)`

SetDatasetRecordId sets DatasetRecordId field to given value.

### HasDatasetRecordId

`func (o *EvaluationRunRecordV2Response) HasDatasetRecordId() bool`

HasDatasetRecordId returns a boolean if a field has been set.

### SetDatasetRecordIdNil

`func (o *EvaluationRunRecordV2Response) SetDatasetRecordIdNil(b bool)`

 SetDatasetRecordIdNil sets the value for DatasetRecordId to be an explicit nil

### UnsetDatasetRecordId
`func (o *EvaluationRunRecordV2Response) UnsetDatasetRecordId()`

UnsetDatasetRecordId ensures that no value is present for DatasetRecordId, not even an explicit nil
### GetOutputRecord

`func (o *EvaluationRunRecordV2Response) GetOutputRecord() EvaluationRunOutputRecordV2`

GetOutputRecord returns the OutputRecord field if non-nil, zero value otherwise.

### GetOutputRecordOk

`func (o *EvaluationRunRecordV2Response) GetOutputRecordOk() (*EvaluationRunOutputRecordV2, bool)`

GetOutputRecordOk returns a tuple with the OutputRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRecord

`func (o *EvaluationRunRecordV2Response) SetOutputRecord(v EvaluationRunOutputRecordV2)`

SetOutputRecord sets OutputRecord field to given value.

### HasOutputRecord

`func (o *EvaluationRunRecordV2Response) HasOutputRecord() bool`

HasOutputRecord returns a boolean if a field has been set.

### SetOutputRecordNil

`func (o *EvaluationRunRecordV2Response) SetOutputRecordNil(b bool)`

 SetOutputRecordNil sets the value for OutputRecord to be an explicit nil

### UnsetOutputRecord
`func (o *EvaluationRunRecordV2Response) UnsetOutputRecord()`

UnsetOutputRecord ensures that no value is present for OutputRecord, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



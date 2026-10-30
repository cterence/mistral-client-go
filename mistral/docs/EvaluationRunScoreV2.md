# EvaluationRunScoreV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Value** | [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | 
**Rationale** | Pointer to **NullableString** |  | [optional] 
**Metadata** | **map[string]interface{}** |  | 
**Status** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEvaluationRunScoreV2

`func NewEvaluationRunScoreV2(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, value NullableAnyOf, metadata map[string]*interface{}, status string, ) *EvaluationRunScoreV2`

NewEvaluationRunScoreV2 instantiates a new EvaluationRunScoreV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunScoreV2WithDefaults

`func NewEvaluationRunScoreV2WithDefaults() *EvaluationRunScoreV2`

NewEvaluationRunScoreV2WithDefaults instantiates a new EvaluationRunScoreV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluationRunScoreV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluationRunScoreV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluationRunScoreV2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EvaluationRunScoreV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvaluationRunScoreV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvaluationRunScoreV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EvaluationRunScoreV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvaluationRunScoreV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvaluationRunScoreV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EvaluationRunScoreV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EvaluationRunScoreV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EvaluationRunScoreV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *EvaluationRunScoreV2) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EvaluationRunScoreV2) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetValue

`func (o *EvaluationRunScoreV2) GetValue() AnyOf`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EvaluationRunScoreV2) GetValueOk() (*AnyOf, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EvaluationRunScoreV2) SetValue(v AnyOf)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *EvaluationRunScoreV2) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EvaluationRunScoreV2) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRationale

`func (o *EvaluationRunScoreV2) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *EvaluationRunScoreV2) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *EvaluationRunScoreV2) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *EvaluationRunScoreV2) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### SetRationaleNil

`func (o *EvaluationRunScoreV2) SetRationaleNil(b bool)`

 SetRationaleNil sets the value for Rationale to be an explicit nil

### UnsetRationale
`func (o *EvaluationRunScoreV2) UnsetRationale()`

UnsetRationale ensures that no value is present for Rationale, not even an explicit nil
### GetMetadata

`func (o *EvaluationRunScoreV2) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvaluationRunScoreV2) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvaluationRunScoreV2) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *EvaluationRunScoreV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationRunScoreV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationRunScoreV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *EvaluationRunScoreV2) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EvaluationRunScoreV2) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EvaluationRunScoreV2) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EvaluationRunScoreV2) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EvaluationRunScoreV2) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EvaluationRunScoreV2) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



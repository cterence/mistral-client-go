# ManagedIndexFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentFields** | Pointer to [**map[string]DocumentFieldsValue**](DocumentFieldsValue.md) |  | [optional] [default to {}]
**ChunkFields** | Pointer to [**map[string]DocumentFieldsValue**](DocumentFieldsValue.md) |  | [optional] [default to {}]

## Methods

### NewManagedIndexFields

`func NewManagedIndexFields() *ManagedIndexFields`

NewManagedIndexFields instantiates a new ManagedIndexFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedIndexFieldsWithDefaults

`func NewManagedIndexFieldsWithDefaults() *ManagedIndexFields`

NewManagedIndexFieldsWithDefaults instantiates a new ManagedIndexFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentFields

`func (o *ManagedIndexFields) GetDocumentFields() map[string]DocumentFieldsValue`

GetDocumentFields returns the DocumentFields field if non-nil, zero value otherwise.

### GetDocumentFieldsOk

`func (o *ManagedIndexFields) GetDocumentFieldsOk() (*map[string]DocumentFieldsValue, bool)`

GetDocumentFieldsOk returns a tuple with the DocumentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFields

`func (o *ManagedIndexFields) SetDocumentFields(v map[string]DocumentFieldsValue)`

SetDocumentFields sets DocumentFields field to given value.

### HasDocumentFields

`func (o *ManagedIndexFields) HasDocumentFields() bool`

HasDocumentFields returns a boolean if a field has been set.

### GetChunkFields

`func (o *ManagedIndexFields) GetChunkFields() map[string]DocumentFieldsValue`

GetChunkFields returns the ChunkFields field if non-nil, zero value otherwise.

### GetChunkFieldsOk

`func (o *ManagedIndexFields) GetChunkFieldsOk() (*map[string]DocumentFieldsValue, bool)`

GetChunkFieldsOk returns a tuple with the ChunkFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkFields

`func (o *ManagedIndexFields) SetChunkFields(v map[string]DocumentFieldsValue)`

SetChunkFields sets ChunkFields field to given value.

### HasChunkFields

`func (o *ManagedIndexFields) HasChunkFields() bool`

HasChunkFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



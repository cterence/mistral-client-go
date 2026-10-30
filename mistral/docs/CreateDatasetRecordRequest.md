# CreateDatasetRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **map[string]interface{}** | Caller-authored input object stored on a dataset record. | 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to **string** | Caller-declared channel that initiated record creation. This value does not certify that the payload is an unmodified copy of its source. | [optional] [default to "DIRECT_INPUT"]

## Methods

### NewCreateDatasetRecordRequest

`func NewCreateDatasetRecordRequest(payload map[string]interface{}, ) *CreateDatasetRecordRequest`

NewCreateDatasetRecordRequest instantiates a new CreateDatasetRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatasetRecordRequestWithDefaults

`func NewCreateDatasetRecordRequestWithDefaults() *CreateDatasetRecordRequest`

NewCreateDatasetRecordRequestWithDefaults instantiates a new CreateDatasetRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *CreateDatasetRecordRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateDatasetRecordRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateDatasetRecordRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetProperties

`func (o *CreateDatasetRecordRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateDatasetRecordRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateDatasetRecordRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateDatasetRecordRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSource

`func (o *CreateDatasetRecordRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateDatasetRecordRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateDatasetRecordRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateDatasetRecordRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomId** | Pointer to **NullableString** |  | [optional] 
**Body** | **map[string]interface{}** |  | 

## Methods

### NewBatchRequest

`func NewBatchRequest(body map[string]interface{}, ) *BatchRequest`

NewBatchRequest instantiates a new BatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestWithDefaults

`func NewBatchRequestWithDefaults() *BatchRequest`

NewBatchRequestWithDefaults instantiates a new BatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomId

`func (o *BatchRequest) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *BatchRequest) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *BatchRequest) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *BatchRequest) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### SetCustomIdNil

`func (o *BatchRequest) SetCustomIdNil(b bool)`

 SetCustomIdNil sets the value for CustomId to be an explicit nil

### UnsetCustomId
`func (o *BatchRequest) UnsetCustomId()`

UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
### GetBody

`func (o *BatchRequest) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchRequest) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchRequest) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



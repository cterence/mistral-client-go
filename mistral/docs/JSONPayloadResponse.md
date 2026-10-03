# JSONPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator indicating this is a raw JSON payload. | [default to "json"]
**Value** | **interface{}** | The JSON-serializable payload value. | 

## Methods

### NewJSONPayloadResponse

`func NewJSONPayloadResponse(type_ string, value interface{}, ) *JSONPayloadResponse`

NewJSONPayloadResponse instantiates a new JSONPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPayloadResponseWithDefaults

`func NewJSONPayloadResponseWithDefaults() *JSONPayloadResponse`

NewJSONPayloadResponseWithDefaults instantiates a new JSONPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JSONPayloadResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JSONPayloadResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JSONPayloadResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *JSONPayloadResponse) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPayloadResponse) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPayloadResponse) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JSONPayloadResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JSONPayloadResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



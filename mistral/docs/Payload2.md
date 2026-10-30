# Payload2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator indicating this is a raw JSON payload. | [default to "json"]
**Value** | [**JSONPatchPayloadValueResponse**](JSONPatchPayloadValueResponse.md) | The list of JSON Patch operations. When encrypted, contains base64-encoded data. | 
**EncodingOptions** | Pointer to [**[]EncodedPayloadOptions**](EncodedPayloadOptions.md) | Encoding options applied to the payload. | [optional] 

## Methods

### NewPayload2

`func NewPayload2(type_ string, value JSONPatchPayloadValueResponse, ) *Payload2`

NewPayload2 instantiates a new Payload2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayload2WithDefaults

`func NewPayload2WithDefaults() *Payload2`

NewPayload2WithDefaults instantiates a new Payload2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Payload2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payload2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payload2) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Payload2) GetValue() JSONPatchPayloadValueResponse`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Payload2) GetValueOk() (*JSONPatchPayloadValueResponse, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Payload2) SetValue(v JSONPatchPayloadValueResponse)`

SetValue sets Value field to given value.


### GetEncodingOptions

`func (o *Payload2) GetEncodingOptions() []EncodedPayloadOptions`

GetEncodingOptions returns the EncodingOptions field if non-nil, zero value otherwise.

### GetEncodingOptionsOk

`func (o *Payload2) GetEncodingOptionsOk() (*[]EncodedPayloadOptions, bool)`

GetEncodingOptionsOk returns a tuple with the EncodingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingOptions

`func (o *Payload2) SetEncodingOptions(v []EncodedPayloadOptions)`

SetEncodingOptions sets EncodingOptions field to given value.

### HasEncodingOptions

`func (o *Payload2) HasEncodingOptions() bool`

HasEncodingOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



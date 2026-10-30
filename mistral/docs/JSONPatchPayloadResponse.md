# JSONPatchPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator indicating this is a JSON Patch payload. | [default to "json_patch"]
**Value** | [**JSONPatchPayloadValueResponse**](JSONPatchPayloadValueResponse.md) | The list of JSON Patch operations. When encrypted, contains base64-encoded data. | 
**EncodingOptions** | Pointer to [**[]EncodedPayloadOptions**](EncodedPayloadOptions.md) | Encoding options applied to the payload. | [optional] 

## Methods

### NewJSONPatchPayloadResponse

`func NewJSONPatchPayloadResponse(type_ string, value JSONPatchPayloadValueResponse, ) *JSONPatchPayloadResponse`

NewJSONPatchPayloadResponse instantiates a new JSONPatchPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchPayloadResponseWithDefaults

`func NewJSONPatchPayloadResponseWithDefaults() *JSONPatchPayloadResponse`

NewJSONPatchPayloadResponseWithDefaults instantiates a new JSONPatchPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JSONPatchPayloadResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JSONPatchPayloadResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JSONPatchPayloadResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *JSONPatchPayloadResponse) GetValue() JSONPatchPayloadValueResponse`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPatchPayloadResponse) GetValueOk() (*JSONPatchPayloadValueResponse, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPatchPayloadResponse) SetValue(v JSONPatchPayloadValueResponse)`

SetValue sets Value field to given value.


### GetEncodingOptions

`func (o *JSONPatchPayloadResponse) GetEncodingOptions() []EncodedPayloadOptions`

GetEncodingOptions returns the EncodingOptions field if non-nil, zero value otherwise.

### GetEncodingOptionsOk

`func (o *JSONPatchPayloadResponse) GetEncodingOptionsOk() (*[]EncodedPayloadOptions, bool)`

GetEncodingOptionsOk returns a tuple with the EncodingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingOptions

`func (o *JSONPatchPayloadResponse) SetEncodingOptions(v []EncodedPayloadOptions)`

SetEncodingOptions sets EncodingOptions field to given value.

### HasEncodingOptions

`func (o *JSONPatchPayloadResponse) HasEncodingOptions() bool`

HasEncodingOptions returns a boolean if a field has been set.

### SetEncodingOptionsNil

`func (o *JSONPatchPayloadResponse) SetEncodingOptionsNil(b bool)`

 SetEncodingOptionsNil sets the value for EncodingOptions to be an explicit nil

### UnsetEncodingOptions
`func (o *JSONPatchPayloadResponse) UnsetEncodingOptions()`

UnsetEncodingOptions ensures that no value is present for EncodingOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Input4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**B64payload** | **string** | The encoded payload | 
**EncodingOptions** | Pointer to [**[]EncodedPayloadOptions**](EncodedPayloadOptions.md) | The encoding of the payload | [optional] [default to {}]
**Empty** | Pointer to **bool** | Whether the payload is empty | [optional] [default to false]

## Methods

### NewInput4

`func NewInput4(b64payload string, ) *Input4`

NewInput4 instantiates a new Input4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInput4WithDefaults

`func NewInput4WithDefaults() *Input4`

NewInput4WithDefaults instantiates a new Input4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetB64payload

`func (o *Input4) GetB64payload() string`

GetB64payload returns the B64payload field if non-nil, zero value otherwise.

### GetB64payloadOk

`func (o *Input4) GetB64payloadOk() (*string, bool)`

GetB64payloadOk returns a tuple with the B64payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB64payload

`func (o *Input4) SetB64payload(v string)`

SetB64payload sets B64payload field to given value.


### GetEncodingOptions

`func (o *Input4) GetEncodingOptions() []EncodedPayloadOptions`

GetEncodingOptions returns the EncodingOptions field if non-nil, zero value otherwise.

### GetEncodingOptionsOk

`func (o *Input4) GetEncodingOptionsOk() (*[]EncodedPayloadOptions, bool)`

GetEncodingOptionsOk returns a tuple with the EncodingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingOptions

`func (o *Input4) SetEncodingOptions(v []EncodedPayloadOptions)`

SetEncodingOptions sets EncodingOptions field to given value.

### HasEncodingOptions

`func (o *Input4) HasEncodingOptions() bool`

HasEncodingOptions returns a boolean if a field has been set.

### GetEmpty

`func (o *Input4) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *Input4) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *Input4) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *Input4) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# JSONPatchPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator indicating this is a JSON Patch payload. | [default to "json_patch"]
**Value** | [**[]ValueInner**](ValueInner.md) | The list of JSON Patch operations to apply in order. | 

## Methods

### NewJSONPatchPayloadResponse

`func NewJSONPatchPayloadResponse(type_ string, value []ValueInner, ) *JSONPatchPayloadResponse`

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

`func (o *JSONPatchPayloadResponse) GetValue() []ValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPatchPayloadResponse) GetValueOk() (*[]ValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPatchPayloadResponse) SetValue(v []ValueInner)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



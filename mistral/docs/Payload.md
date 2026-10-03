# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator indicating this is a raw JSON payload. | [default to "json"]
**Value** | [**[]ValueInner**](ValueInner.md) | The list of JSON Patch operations to apply in order. | 

## Methods

### NewPayload

`func NewPayload(type_ string, value []ValueInner, ) *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Payload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payload) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Payload) GetValue() []ValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Payload) GetValueOk() (*[]ValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Payload) SetValue(v []ValueInner)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Content7Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | [**[]Annotations1Inner**](Annotations1Inner.md) |  | 
**Text** | **string** |  | 
**Type** | **string** |  | 
**Logprobs** | Pointer to [**[]ResponsesOutputTextLogprob**](ResponsesOutputTextLogprob.md) |  | [optional] 
**Refusal** | **string** |  | 

## Methods

### NewContent7Inner

`func NewContent7Inner(annotations []Annotations1Inner, text string, type_ string, refusal string, ) *Content7Inner`

NewContent7Inner instantiates a new Content7Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent7InnerWithDefaults

`func NewContent7InnerWithDefaults() *Content7Inner`

NewContent7InnerWithDefaults instantiates a new Content7Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *Content7Inner) GetAnnotations() []Annotations1Inner`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Content7Inner) GetAnnotationsOk() (*[]Annotations1Inner, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Content7Inner) SetAnnotations(v []Annotations1Inner)`

SetAnnotations sets Annotations field to given value.


### GetText

`func (o *Content7Inner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Content7Inner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Content7Inner) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *Content7Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content7Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content7Inner) SetType(v string)`

SetType sets Type field to given value.


### GetLogprobs

`func (o *Content7Inner) GetLogprobs() []ResponsesOutputTextLogprob`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *Content7Inner) GetLogprobsOk() (*[]ResponsesOutputTextLogprob, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *Content7Inner) SetLogprobs(v []ResponsesOutputTextLogprob)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *Content7Inner) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *Content7Inner) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *Content7Inner) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetRefusal

`func (o *Content7Inner) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *Content7Inner) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *Content7Inner) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



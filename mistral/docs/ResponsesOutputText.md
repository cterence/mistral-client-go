# ResponsesOutputText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | [**[]Annotations1Inner**](Annotations1Inner.md) |  | 
**Text** | **string** |  | 
**Type** | **string** |  | 
**Logprobs** | Pointer to [**[]ResponsesOutputTextLogprob**](ResponsesOutputTextLogprob.md) |  | [optional] 

## Methods

### NewResponsesOutputText

`func NewResponsesOutputText(annotations []Annotations1Inner, text string, type_ string, ) *ResponsesOutputText`

NewResponsesOutputText instantiates a new ResponsesOutputText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesOutputTextWithDefaults

`func NewResponsesOutputTextWithDefaults() *ResponsesOutputText`

NewResponsesOutputTextWithDefaults instantiates a new ResponsesOutputText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *ResponsesOutputText) GetAnnotations() []Annotations1Inner`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ResponsesOutputText) GetAnnotationsOk() (*[]Annotations1Inner, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ResponsesOutputText) SetAnnotations(v []Annotations1Inner)`

SetAnnotations sets Annotations field to given value.


### GetText

`func (o *ResponsesOutputText) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ResponsesOutputText) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ResponsesOutputText) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ResponsesOutputText) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesOutputText) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesOutputText) SetType(v string)`

SetType sets Type field to given value.


### GetLogprobs

`func (o *ResponsesOutputText) GetLogprobs() []ResponsesOutputTextLogprob`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ResponsesOutputText) GetLogprobsOk() (*[]ResponsesOutputTextLogprob, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ResponsesOutputText) SetLogprobs(v []ResponsesOutputTextLogprob)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *ResponsesOutputText) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *ResponsesOutputText) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *ResponsesOutputText) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



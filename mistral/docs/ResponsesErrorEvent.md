# ResponsesErrorEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 
**Param** | Pointer to **NullableString** |  | [optional] 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesErrorEvent

`func NewResponsesErrorEvent(message string, sequenceNumber int32, type_ string, ) *ResponsesErrorEvent`

NewResponsesErrorEvent instantiates a new ResponsesErrorEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesErrorEventWithDefaults

`func NewResponsesErrorEventWithDefaults() *ResponsesErrorEvent`

NewResponsesErrorEventWithDefaults instantiates a new ResponsesErrorEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResponsesErrorEvent) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponsesErrorEvent) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponsesErrorEvent) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResponsesErrorEvent) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ResponsesErrorEvent) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ResponsesErrorEvent) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMessage

`func (o *ResponsesErrorEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponsesErrorEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponsesErrorEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *ResponsesErrorEvent) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ResponsesErrorEvent) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ResponsesErrorEvent) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ResponsesErrorEvent) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *ResponsesErrorEvent) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *ResponsesErrorEvent) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetSequenceNumber

`func (o *ResponsesErrorEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesErrorEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesErrorEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *ResponsesErrorEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesErrorEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesErrorEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



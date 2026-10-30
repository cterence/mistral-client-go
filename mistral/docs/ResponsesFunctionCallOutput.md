# ResponsesFunctionCallOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | **string** |  | 
**Output** | [**Output1**](Output1.md) |  | 
**Type** | **string** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesFunctionCallOutput

`func NewResponsesFunctionCallOutput(callId string, output Output1, type_ string, ) *ResponsesFunctionCallOutput`

NewResponsesFunctionCallOutput instantiates a new ResponsesFunctionCallOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesFunctionCallOutputWithDefaults

`func NewResponsesFunctionCallOutputWithDefaults() *ResponsesFunctionCallOutput`

NewResponsesFunctionCallOutputWithDefaults instantiates a new ResponsesFunctionCallOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *ResponsesFunctionCallOutput) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *ResponsesFunctionCallOutput) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *ResponsesFunctionCallOutput) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetOutput

`func (o *ResponsesFunctionCallOutput) GetOutput() Output1`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ResponsesFunctionCallOutput) GetOutputOk() (*Output1, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ResponsesFunctionCallOutput) SetOutput(v Output1)`

SetOutput sets Output field to given value.


### GetType

`func (o *ResponsesFunctionCallOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesFunctionCallOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesFunctionCallOutput) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ResponsesFunctionCallOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesFunctionCallOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesFunctionCallOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesFunctionCallOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ResponsesFunctionCallOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResponsesFunctionCallOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCaller

`func (o *ResponsesFunctionCallOutput) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *ResponsesFunctionCallOutput) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *ResponsesFunctionCallOutput) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *ResponsesFunctionCallOutput) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *ResponsesFunctionCallOutput) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *ResponsesFunctionCallOutput) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetStatus

`func (o *ResponsesFunctionCallOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesFunctionCallOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesFunctionCallOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesFunctionCallOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ResponsesFunctionCallOutput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ResponsesFunctionCallOutput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



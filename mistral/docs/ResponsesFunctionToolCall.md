# ResponsesFunctionToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | **string** |  | 
**CallId** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Caller** | Pointer to [**NullableCaller1**](Caller1.md) |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesFunctionToolCall

`func NewResponsesFunctionToolCall(arguments string, callId string, name string, type_ string, ) *ResponsesFunctionToolCall`

NewResponsesFunctionToolCall instantiates a new ResponsesFunctionToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesFunctionToolCallWithDefaults

`func NewResponsesFunctionToolCallWithDefaults() *ResponsesFunctionToolCall`

NewResponsesFunctionToolCallWithDefaults instantiates a new ResponsesFunctionToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *ResponsesFunctionToolCall) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ResponsesFunctionToolCall) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ResponsesFunctionToolCall) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetCallId

`func (o *ResponsesFunctionToolCall) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *ResponsesFunctionToolCall) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *ResponsesFunctionToolCall) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetName

`func (o *ResponsesFunctionToolCall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesFunctionToolCall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesFunctionToolCall) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResponsesFunctionToolCall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesFunctionToolCall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesFunctionToolCall) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ResponsesFunctionToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesFunctionToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesFunctionToolCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesFunctionToolCall) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ResponsesFunctionToolCall) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResponsesFunctionToolCall) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCaller

`func (o *ResponsesFunctionToolCall) GetCaller() Caller1`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *ResponsesFunctionToolCall) GetCallerOk() (*Caller1, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *ResponsesFunctionToolCall) SetCaller(v Caller1)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *ResponsesFunctionToolCall) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *ResponsesFunctionToolCall) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *ResponsesFunctionToolCall) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetNamespace

`func (o *ResponsesFunctionToolCall) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResponsesFunctionToolCall) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResponsesFunctionToolCall) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ResponsesFunctionToolCall) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ResponsesFunctionToolCall) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ResponsesFunctionToolCall) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStatus

`func (o *ResponsesFunctionToolCall) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesFunctionToolCall) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesFunctionToolCall) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesFunctionToolCall) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ResponsesFunctionToolCall) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ResponsesFunctionToolCall) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



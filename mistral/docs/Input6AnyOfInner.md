# Input6AnyOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]ResponsesContent**](ResponsesContent.md) |  | 
**Role** | **string** |  | 
**Phase** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Status** | **NullableString** |  | 
**Id** | **string** |  | 
**Arguments** | **string** |  | 
**CallId** | **string** |  | 
**Name** | **string** |  | 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Output** | [**Output1**](Output1.md) |  | 
**Summary** | [**[]ResponsesSummary**](ResponsesSummary.md) |  | 
**EncryptedContent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInput6AnyOfInner

`func NewInput6AnyOfInner(content []ResponsesContent, role string, type_ string, status NullableString, id string, arguments string, callId string, name string, output Output1, summary []ResponsesSummary, ) *Input6AnyOfInner`

NewInput6AnyOfInner instantiates a new Input6AnyOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInput6AnyOfInnerWithDefaults

`func NewInput6AnyOfInnerWithDefaults() *Input6AnyOfInner`

NewInput6AnyOfInnerWithDefaults instantiates a new Input6AnyOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Input6AnyOfInner) GetContent() []ResponsesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Input6AnyOfInner) GetContentOk() (*[]ResponsesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Input6AnyOfInner) SetContent(v []ResponsesContent)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *Input6AnyOfInner) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *Input6AnyOfInner) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRole

`func (o *Input6AnyOfInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Input6AnyOfInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Input6AnyOfInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetPhase

`func (o *Input6AnyOfInner) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Input6AnyOfInner) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Input6AnyOfInner) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Input6AnyOfInner) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *Input6AnyOfInner) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *Input6AnyOfInner) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetType

`func (o *Input6AnyOfInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Input6AnyOfInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Input6AnyOfInner) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Input6AnyOfInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Input6AnyOfInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Input6AnyOfInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Input6AnyOfInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Input6AnyOfInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetId

`func (o *Input6AnyOfInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Input6AnyOfInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Input6AnyOfInner) SetId(v string)`

SetId sets Id field to given value.


### GetArguments

`func (o *Input6AnyOfInner) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Input6AnyOfInner) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Input6AnyOfInner) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetCallId

`func (o *Input6AnyOfInner) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Input6AnyOfInner) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Input6AnyOfInner) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetName

`func (o *Input6AnyOfInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Input6AnyOfInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Input6AnyOfInner) SetName(v string)`

SetName sets Name field to given value.


### GetCaller

`func (o *Input6AnyOfInner) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Input6AnyOfInner) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Input6AnyOfInner) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Input6AnyOfInner) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Input6AnyOfInner) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Input6AnyOfInner) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetNamespace

`func (o *Input6AnyOfInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Input6AnyOfInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Input6AnyOfInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Input6AnyOfInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Input6AnyOfInner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Input6AnyOfInner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetOutput

`func (o *Input6AnyOfInner) GetOutput() Output1`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Input6AnyOfInner) GetOutputOk() (*Output1, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Input6AnyOfInner) SetOutput(v Output1)`

SetOutput sets Output field to given value.


### GetSummary

`func (o *Input6AnyOfInner) GetSummary() []ResponsesSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Input6AnyOfInner) GetSummaryOk() (*[]ResponsesSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Input6AnyOfInner) SetSummary(v []ResponsesSummary)`

SetSummary sets Summary field to given value.


### GetEncryptedContent

`func (o *Input6AnyOfInner) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *Input6AnyOfInner) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *Input6AnyOfInner) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *Input6AnyOfInner) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### SetEncryptedContentNil

`func (o *Input6AnyOfInner) SetEncryptedContentNil(b bool)`

 SetEncryptedContentNil sets the value for EncryptedContent to be an explicit nil

### UnsetEncryptedContent
`func (o *Input6AnyOfInner) UnsetEncryptedContent()`

UnsetEncryptedContent ensures that no value is present for EncryptedContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Output2Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Content** | [**[]ResponsesContent**](ResponsesContent.md) |  | 
**Role** | **string** |  | 
**Status** | **NullableString** |  | 
**Type** | **string** |  | 
**Phase** | Pointer to **NullableString** |  | [optional] 
**Arguments** | **string** |  | 
**CallId** | **string** |  | 
**Name** | **string** |  | 
**Caller** | Pointer to [**NullableCaller1**](Caller1.md) |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Summary** | [**[]ResponsesSummary**](ResponsesSummary.md) |  | 
**EncryptedContent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOutput2Inner

`func NewOutput2Inner(id string, content []ResponsesContent, role string, status NullableString, type_ string, arguments string, callId string, name string, summary []ResponsesSummary, ) *Output2Inner`

NewOutput2Inner instantiates a new Output2Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutput2InnerWithDefaults

`func NewOutput2InnerWithDefaults() *Output2Inner`

NewOutput2InnerWithDefaults instantiates a new Output2Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Output2Inner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Output2Inner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Output2Inner) SetId(v string)`

SetId sets Id field to given value.


### GetContent

`func (o *Output2Inner) GetContent() []ResponsesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Output2Inner) GetContentOk() (*[]ResponsesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Output2Inner) SetContent(v []ResponsesContent)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *Output2Inner) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *Output2Inner) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRole

`func (o *Output2Inner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Output2Inner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Output2Inner) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *Output2Inner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Output2Inner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Output2Inner) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Output2Inner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Output2Inner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *Output2Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Output2Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Output2Inner) SetType(v string)`

SetType sets Type field to given value.


### GetPhase

`func (o *Output2Inner) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Output2Inner) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Output2Inner) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Output2Inner) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *Output2Inner) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *Output2Inner) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetArguments

`func (o *Output2Inner) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Output2Inner) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Output2Inner) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetCallId

`func (o *Output2Inner) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Output2Inner) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Output2Inner) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetName

`func (o *Output2Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Output2Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Output2Inner) SetName(v string)`

SetName sets Name field to given value.


### GetCaller

`func (o *Output2Inner) GetCaller() Caller1`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Output2Inner) GetCallerOk() (*Caller1, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Output2Inner) SetCaller(v Caller1)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Output2Inner) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Output2Inner) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Output2Inner) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetNamespace

`func (o *Output2Inner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Output2Inner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Output2Inner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Output2Inner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Output2Inner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Output2Inner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetSummary

`func (o *Output2Inner) GetSummary() []ResponsesSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Output2Inner) GetSummaryOk() (*[]ResponsesSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Output2Inner) SetSummary(v []ResponsesSummary)`

SetSummary sets Summary field to given value.


### GetEncryptedContent

`func (o *Output2Inner) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *Output2Inner) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *Output2Inner) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *Output2Inner) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### SetEncryptedContentNil

`func (o *Output2Inner) SetEncryptedContentNil(b bool)`

 SetEncryptedContentNil sets the value for EncryptedContent to be an explicit nil

### UnsetEncryptedContent
`func (o *Output2Inner) UnsetEncryptedContent()`

UnsetEncryptedContent ensures that no value is present for EncryptedContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



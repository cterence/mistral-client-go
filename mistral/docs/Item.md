# Item

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

### NewItem

`func NewItem(id string, content []ResponsesContent, role string, status NullableString, type_ string, arguments string, callId string, name string, summary []ResponsesSummary, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.


### GetContent

`func (o *Item) GetContent() []ResponsesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Item) GetContentOk() (*[]ResponsesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Item) SetContent(v []ResponsesContent)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *Item) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *Item) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRole

`func (o *Item) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Item) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Item) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *Item) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Item) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Item) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Item) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Item) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *Item) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Item) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Item) SetType(v string)`

SetType sets Type field to given value.


### GetPhase

`func (o *Item) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Item) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Item) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Item) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *Item) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *Item) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetArguments

`func (o *Item) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Item) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Item) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetCallId

`func (o *Item) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Item) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Item) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.


### GetCaller

`func (o *Item) GetCaller() Caller1`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Item) GetCallerOk() (*Caller1, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Item) SetCaller(v Caller1)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Item) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Item) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Item) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetNamespace

`func (o *Item) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Item) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Item) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Item) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Item) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Item) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetSummary

`func (o *Item) GetSummary() []ResponsesSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Item) GetSummaryOk() (*[]ResponsesSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Item) SetSummary(v []ResponsesSummary)`

SetSummary sets Summary field to given value.


### GetEncryptedContent

`func (o *Item) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *Item) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *Item) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *Item) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### SetEncryptedContentNil

`func (o *Item) SetEncryptedContentNil(b bool)`

 SetEncryptedContentNil sets the value for EncryptedContent to be an explicit nil

### UnsetEncryptedContent
`func (o *Item) UnsetEncryptedContent()`

UnsetEncryptedContent ensures that no value is present for EncryptedContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



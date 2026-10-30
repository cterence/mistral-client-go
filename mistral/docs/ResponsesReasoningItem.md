# ResponsesReasoningItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Summary** | [**[]ResponsesSummary**](ResponsesSummary.md) |  | 
**Type** | **string** |  | 
**Content** | Pointer to [**[]ResponsesContent**](ResponsesContent.md) |  | [optional] 
**EncryptedContent** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesReasoningItem

`func NewResponsesReasoningItem(id string, summary []ResponsesSummary, type_ string, ) *ResponsesReasoningItem`

NewResponsesReasoningItem instantiates a new ResponsesReasoningItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesReasoningItemWithDefaults

`func NewResponsesReasoningItemWithDefaults() *ResponsesReasoningItem`

NewResponsesReasoningItemWithDefaults instantiates a new ResponsesReasoningItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsesReasoningItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesReasoningItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesReasoningItem) SetId(v string)`

SetId sets Id field to given value.


### GetSummary

`func (o *ResponsesReasoningItem) GetSummary() []ResponsesSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ResponsesReasoningItem) GetSummaryOk() (*[]ResponsesSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ResponsesReasoningItem) SetSummary(v []ResponsesSummary)`

SetSummary sets Summary field to given value.


### GetType

`func (o *ResponsesReasoningItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesReasoningItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesReasoningItem) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *ResponsesReasoningItem) GetContent() []ResponsesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResponsesReasoningItem) GetContentOk() (*[]ResponsesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResponsesReasoningItem) SetContent(v []ResponsesContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *ResponsesReasoningItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ResponsesReasoningItem) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ResponsesReasoningItem) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEncryptedContent

`func (o *ResponsesReasoningItem) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *ResponsesReasoningItem) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *ResponsesReasoningItem) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *ResponsesReasoningItem) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### SetEncryptedContentNil

`func (o *ResponsesReasoningItem) SetEncryptedContentNil(b bool)`

 SetEncryptedContentNil sets the value for EncryptedContent to be an explicit nil

### UnsetEncryptedContent
`func (o *ResponsesReasoningItem) UnsetEncryptedContent()`

UnsetEncryptedContent ensures that no value is present for EncryptedContent, not even an explicit nil
### GetStatus

`func (o *ResponsesReasoningItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesReasoningItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesReasoningItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesReasoningItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ResponsesReasoningItem) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ResponsesReasoningItem) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



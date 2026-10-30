# ResponsesInProgressEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**CreateResponseResponse**](CreateResponseResponse.md) |  | 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesInProgressEvent

`func NewResponsesInProgressEvent(response CreateResponseResponse, sequenceNumber int32, type_ string, ) *ResponsesInProgressEvent`

NewResponsesInProgressEvent instantiates a new ResponsesInProgressEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesInProgressEventWithDefaults

`func NewResponsesInProgressEventWithDefaults() *ResponsesInProgressEvent`

NewResponsesInProgressEventWithDefaults instantiates a new ResponsesInProgressEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ResponsesInProgressEvent) GetResponse() CreateResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ResponsesInProgressEvent) GetResponseOk() (*CreateResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ResponsesInProgressEvent) SetResponse(v CreateResponseResponse)`

SetResponse sets Response field to given value.


### GetSequenceNumber

`func (o *ResponsesInProgressEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesInProgressEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesInProgressEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *ResponsesInProgressEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesInProgressEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesInProgressEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



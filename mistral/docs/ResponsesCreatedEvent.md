# ResponsesCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**CreateResponseResponse**](CreateResponseResponse.md) |  | 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesCreatedEvent

`func NewResponsesCreatedEvent(response CreateResponseResponse, sequenceNumber int32, type_ string, ) *ResponsesCreatedEvent`

NewResponsesCreatedEvent instantiates a new ResponsesCreatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesCreatedEventWithDefaults

`func NewResponsesCreatedEventWithDefaults() *ResponsesCreatedEvent`

NewResponsesCreatedEventWithDefaults instantiates a new ResponsesCreatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ResponsesCreatedEvent) GetResponse() CreateResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ResponsesCreatedEvent) GetResponseOk() (*CreateResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ResponsesCreatedEvent) SetResponse(v CreateResponseResponse)`

SetResponse sets Response field to given value.


### GetSequenceNumber

`func (o *ResponsesCreatedEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesCreatedEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesCreatedEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *ResponsesCreatedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesCreatedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesCreatedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



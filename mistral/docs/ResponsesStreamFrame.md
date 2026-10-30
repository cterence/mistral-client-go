# ResponsesStreamFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Data** | [**Data4**](Data4.md) |  | 

## Methods

### NewResponsesStreamFrame

`func NewResponsesStreamFrame(event string, data Data4, ) *ResponsesStreamFrame`

NewResponsesStreamFrame instantiates a new ResponsesStreamFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesStreamFrameWithDefaults

`func NewResponsesStreamFrameWithDefaults() *ResponsesStreamFrame`

NewResponsesStreamFrameWithDefaults instantiates a new ResponsesStreamFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ResponsesStreamFrame) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ResponsesStreamFrame) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ResponsesStreamFrame) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetData

`func (o *ResponsesStreamFrame) GetData() Data4`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponsesStreamFrame) GetDataOk() (*Data4, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponsesStreamFrame) SetData(v Data4)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



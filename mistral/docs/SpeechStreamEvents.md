# SpeechStreamEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Data** | [**Data1**](Data1.md) |  | 

## Methods

### NewSpeechStreamEvents

`func NewSpeechStreamEvents(event string, data Data1, ) *SpeechStreamEvents`

NewSpeechStreamEvents instantiates a new SpeechStreamEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeechStreamEventsWithDefaults

`func NewSpeechStreamEventsWithDefaults() *SpeechStreamEvents`

NewSpeechStreamEventsWithDefaults instantiates a new SpeechStreamEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *SpeechStreamEvents) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SpeechStreamEvents) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SpeechStreamEvents) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetData

`func (o *SpeechStreamEvents) GetData() Data1`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SpeechStreamEvents) GetDataOk() (*Data1, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SpeechStreamEvents) SetData(v Data1)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



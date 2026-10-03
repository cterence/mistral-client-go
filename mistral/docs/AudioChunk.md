# AudioChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "input_audio"]
**InputAudio** | [**InputAudio**](InputAudio.md) |  | 

## Methods

### NewAudioChunk

`func NewAudioChunk(inputAudio InputAudio, ) *AudioChunk`

NewAudioChunk instantiates a new AudioChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioChunkWithDefaults

`func NewAudioChunkWithDefaults() *AudioChunk`

NewAudioChunkWithDefaults instantiates a new AudioChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AudioChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudioChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudioChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AudioChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInputAudio

`func (o *AudioChunk) GetInputAudio() InputAudio`

GetInputAudio returns the InputAudio field if non-nil, zero value otherwise.

### GetInputAudioOk

`func (o *AudioChunk) GetInputAudioOk() (*InputAudio, bool)`

GetInputAudioOk returns a tuple with the InputAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAudio

`func (o *AudioChunk) SetInputAudio(v InputAudio)`

SetInputAudio sets InputAudio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VibeCodeUsageDataJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completion** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Vibe Code completion (token) usage data. | 
**Ocr** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Vibe Code OCR usage data. | 
**Connectors** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Vibe Code connectors usage data. | 
**Audio** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Vibe Code audio transcription usage data. | 
**AudioCharacters** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Vibe Code text-to-speech usage data. | 

## Methods

### NewVibeCodeUsageDataJSON

`func NewVibeCodeUsageDataJSON(completion BasicModelUsageDataJSON, ocr BasicModelUsageDataJSON, connectors BasicModelUsageDataJSON, audio BasicModelUsageDataJSON, audioCharacters BasicModelUsageDataJSON, ) *VibeCodeUsageDataJSON`

NewVibeCodeUsageDataJSON instantiates a new VibeCodeUsageDataJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVibeCodeUsageDataJSONWithDefaults

`func NewVibeCodeUsageDataJSONWithDefaults() *VibeCodeUsageDataJSON`

NewVibeCodeUsageDataJSONWithDefaults instantiates a new VibeCodeUsageDataJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletion

`func (o *VibeCodeUsageDataJSON) GetCompletion() BasicModelUsageDataJSON`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *VibeCodeUsageDataJSON) GetCompletionOk() (*BasicModelUsageDataJSON, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *VibeCodeUsageDataJSON) SetCompletion(v BasicModelUsageDataJSON)`

SetCompletion sets Completion field to given value.


### GetOcr

`func (o *VibeCodeUsageDataJSON) GetOcr() BasicModelUsageDataJSON`

GetOcr returns the Ocr field if non-nil, zero value otherwise.

### GetOcrOk

`func (o *VibeCodeUsageDataJSON) GetOcrOk() (*BasicModelUsageDataJSON, bool)`

GetOcrOk returns a tuple with the Ocr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcr

`func (o *VibeCodeUsageDataJSON) SetOcr(v BasicModelUsageDataJSON)`

SetOcr sets Ocr field to given value.


### GetConnectors

`func (o *VibeCodeUsageDataJSON) GetConnectors() BasicModelUsageDataJSON`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *VibeCodeUsageDataJSON) GetConnectorsOk() (*BasicModelUsageDataJSON, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *VibeCodeUsageDataJSON) SetConnectors(v BasicModelUsageDataJSON)`

SetConnectors sets Connectors field to given value.


### GetAudio

`func (o *VibeCodeUsageDataJSON) GetAudio() BasicModelUsageDataJSON`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *VibeCodeUsageDataJSON) GetAudioOk() (*BasicModelUsageDataJSON, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *VibeCodeUsageDataJSON) SetAudio(v BasicModelUsageDataJSON)`

SetAudio sets Audio field to given value.


### GetAudioCharacters

`func (o *VibeCodeUsageDataJSON) GetAudioCharacters() BasicModelUsageDataJSON`

GetAudioCharacters returns the AudioCharacters field if non-nil, zero value otherwise.

### GetAudioCharactersOk

`func (o *VibeCodeUsageDataJSON) GetAudioCharactersOk() (*BasicModelUsageDataJSON, bool)`

GetAudioCharactersOk returns a tuple with the AudioCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCharacters

`func (o *VibeCodeUsageDataJSON) SetAudioCharacters(v BasicModelUsageDataJSON)`

SetAudioCharacters sets AudioCharacters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UsageOUTJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completion** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Completion usage data. | 
**Ocr** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | OCR usage data. | 
**Connectors** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Connector usage data. | 
**LibrariesApi** | [**LibrariesAPIUsageDataJSON**](LibrariesAPIUsageDataJSON.md) | Libraries API usage data. | 
**FineTuning** | [**FineTuningDataJSON**](FineTuningDataJSON.md) | Fine-tuning usage data. | 
**Audio** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Audio usage data. | 
**AudioCharacters** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Text-to-speech (audio characters) usage data. | 
**AudioAgents** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Audio Agent session duration usage data. | 
**VibeUsage** | **float32** | Legacy Vibe usage field. Always 0; use the Vibe usage API instead. | 
**VibeCode** | [**VibeCodeUsageDataJSON**](VibeCodeUsageDataJSON.md) | Vibe Code usage data, broken down by sub-usage type. | 
**Chat** | [**BasicModelUsageDataJSON**](BasicModelUsageDataJSON.md) | Chat usage data. | 
**Date** | **time.Time** | Reference date for the usage period. | 
**PreviousMonth** | **NullableString** | Previous month in the usage report. | 
**NextMonth** | **NullableString** | Next month in the usage report. | 
**StartDate** | **time.Time** | Start of the usage period. | 
**EndDate** | **time.Time** | End of the usage period. | 
**Currency** | **NullableString** | Currency used for usage prices. | 
**CurrencySymbol** | **NullableString** | Currency symbol used for usage prices. | 
**Prices** | [**[]PriceData**](PriceData.md) | Prices used to calculate usage amounts. | 

## Methods

### NewUsageOUTJSON

`func NewUsageOUTJSON(completion BasicModelUsageDataJSON, ocr BasicModelUsageDataJSON, connectors BasicModelUsageDataJSON, librariesApi LibrariesAPIUsageDataJSON, fineTuning FineTuningDataJSON, audio BasicModelUsageDataJSON, audioCharacters BasicModelUsageDataJSON, audioAgents BasicModelUsageDataJSON, vibeUsage float32, vibeCode VibeCodeUsageDataJSON, chat BasicModelUsageDataJSON, date time.Time, previousMonth NullableString, nextMonth NullableString, startDate time.Time, endDate time.Time, currency NullableString, currencySymbol NullableString, prices []PriceData, ) *UsageOUTJSON`

NewUsageOUTJSON instantiates a new UsageOUTJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageOUTJSONWithDefaults

`func NewUsageOUTJSONWithDefaults() *UsageOUTJSON`

NewUsageOUTJSONWithDefaults instantiates a new UsageOUTJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletion

`func (o *UsageOUTJSON) GetCompletion() BasicModelUsageDataJSON`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *UsageOUTJSON) GetCompletionOk() (*BasicModelUsageDataJSON, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *UsageOUTJSON) SetCompletion(v BasicModelUsageDataJSON)`

SetCompletion sets Completion field to given value.


### GetOcr

`func (o *UsageOUTJSON) GetOcr() BasicModelUsageDataJSON`

GetOcr returns the Ocr field if non-nil, zero value otherwise.

### GetOcrOk

`func (o *UsageOUTJSON) GetOcrOk() (*BasicModelUsageDataJSON, bool)`

GetOcrOk returns a tuple with the Ocr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcr

`func (o *UsageOUTJSON) SetOcr(v BasicModelUsageDataJSON)`

SetOcr sets Ocr field to given value.


### GetConnectors

`func (o *UsageOUTJSON) GetConnectors() BasicModelUsageDataJSON`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *UsageOUTJSON) GetConnectorsOk() (*BasicModelUsageDataJSON, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *UsageOUTJSON) SetConnectors(v BasicModelUsageDataJSON)`

SetConnectors sets Connectors field to given value.


### GetLibrariesApi

`func (o *UsageOUTJSON) GetLibrariesApi() LibrariesAPIUsageDataJSON`

GetLibrariesApi returns the LibrariesApi field if non-nil, zero value otherwise.

### GetLibrariesApiOk

`func (o *UsageOUTJSON) GetLibrariesApiOk() (*LibrariesAPIUsageDataJSON, bool)`

GetLibrariesApiOk returns a tuple with the LibrariesApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrariesApi

`func (o *UsageOUTJSON) SetLibrariesApi(v LibrariesAPIUsageDataJSON)`

SetLibrariesApi sets LibrariesApi field to given value.


### GetFineTuning

`func (o *UsageOUTJSON) GetFineTuning() FineTuningDataJSON`

GetFineTuning returns the FineTuning field if non-nil, zero value otherwise.

### GetFineTuningOk

`func (o *UsageOUTJSON) GetFineTuningOk() (*FineTuningDataJSON, bool)`

GetFineTuningOk returns a tuple with the FineTuning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTuning

`func (o *UsageOUTJSON) SetFineTuning(v FineTuningDataJSON)`

SetFineTuning sets FineTuning field to given value.


### GetAudio

`func (o *UsageOUTJSON) GetAudio() BasicModelUsageDataJSON`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *UsageOUTJSON) GetAudioOk() (*BasicModelUsageDataJSON, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *UsageOUTJSON) SetAudio(v BasicModelUsageDataJSON)`

SetAudio sets Audio field to given value.


### GetAudioCharacters

`func (o *UsageOUTJSON) GetAudioCharacters() BasicModelUsageDataJSON`

GetAudioCharacters returns the AudioCharacters field if non-nil, zero value otherwise.

### GetAudioCharactersOk

`func (o *UsageOUTJSON) GetAudioCharactersOk() (*BasicModelUsageDataJSON, bool)`

GetAudioCharactersOk returns a tuple with the AudioCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCharacters

`func (o *UsageOUTJSON) SetAudioCharacters(v BasicModelUsageDataJSON)`

SetAudioCharacters sets AudioCharacters field to given value.


### GetAudioAgents

`func (o *UsageOUTJSON) GetAudioAgents() BasicModelUsageDataJSON`

GetAudioAgents returns the AudioAgents field if non-nil, zero value otherwise.

### GetAudioAgentsOk

`func (o *UsageOUTJSON) GetAudioAgentsOk() (*BasicModelUsageDataJSON, bool)`

GetAudioAgentsOk returns a tuple with the AudioAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioAgents

`func (o *UsageOUTJSON) SetAudioAgents(v BasicModelUsageDataJSON)`

SetAudioAgents sets AudioAgents field to given value.


### GetVibeUsage

`func (o *UsageOUTJSON) GetVibeUsage() float32`

GetVibeUsage returns the VibeUsage field if non-nil, zero value otherwise.

### GetVibeUsageOk

`func (o *UsageOUTJSON) GetVibeUsageOk() (*float32, bool)`

GetVibeUsageOk returns a tuple with the VibeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVibeUsage

`func (o *UsageOUTJSON) SetVibeUsage(v float32)`

SetVibeUsage sets VibeUsage field to given value.


### GetVibeCode

`func (o *UsageOUTJSON) GetVibeCode() VibeCodeUsageDataJSON`

GetVibeCode returns the VibeCode field if non-nil, zero value otherwise.

### GetVibeCodeOk

`func (o *UsageOUTJSON) GetVibeCodeOk() (*VibeCodeUsageDataJSON, bool)`

GetVibeCodeOk returns a tuple with the VibeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVibeCode

`func (o *UsageOUTJSON) SetVibeCode(v VibeCodeUsageDataJSON)`

SetVibeCode sets VibeCode field to given value.


### GetChat

`func (o *UsageOUTJSON) GetChat() BasicModelUsageDataJSON`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *UsageOUTJSON) GetChatOk() (*BasicModelUsageDataJSON, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *UsageOUTJSON) SetChat(v BasicModelUsageDataJSON)`

SetChat sets Chat field to given value.


### GetDate

`func (o *UsageOUTJSON) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsageOUTJSON) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsageOUTJSON) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetPreviousMonth

`func (o *UsageOUTJSON) GetPreviousMonth() string`

GetPreviousMonth returns the PreviousMonth field if non-nil, zero value otherwise.

### GetPreviousMonthOk

`func (o *UsageOUTJSON) GetPreviousMonthOk() (*string, bool)`

GetPreviousMonthOk returns a tuple with the PreviousMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMonth

`func (o *UsageOUTJSON) SetPreviousMonth(v string)`

SetPreviousMonth sets PreviousMonth field to given value.


### SetPreviousMonthNil

`func (o *UsageOUTJSON) SetPreviousMonthNil(b bool)`

 SetPreviousMonthNil sets the value for PreviousMonth to be an explicit nil

### UnsetPreviousMonth
`func (o *UsageOUTJSON) UnsetPreviousMonth()`

UnsetPreviousMonth ensures that no value is present for PreviousMonth, not even an explicit nil
### GetNextMonth

`func (o *UsageOUTJSON) GetNextMonth() string`

GetNextMonth returns the NextMonth field if non-nil, zero value otherwise.

### GetNextMonthOk

`func (o *UsageOUTJSON) GetNextMonthOk() (*string, bool)`

GetNextMonthOk returns a tuple with the NextMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMonth

`func (o *UsageOUTJSON) SetNextMonth(v string)`

SetNextMonth sets NextMonth field to given value.


### SetNextMonthNil

`func (o *UsageOUTJSON) SetNextMonthNil(b bool)`

 SetNextMonthNil sets the value for NextMonth to be an explicit nil

### UnsetNextMonth
`func (o *UsageOUTJSON) UnsetNextMonth()`

UnsetNextMonth ensures that no value is present for NextMonth, not even an explicit nil
### GetStartDate

`func (o *UsageOUTJSON) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageOUTJSON) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageOUTJSON) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *UsageOUTJSON) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageOUTJSON) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageOUTJSON) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetCurrency

`func (o *UsageOUTJSON) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UsageOUTJSON) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UsageOUTJSON) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *UsageOUTJSON) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UsageOUTJSON) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCurrencySymbol

`func (o *UsageOUTJSON) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *UsageOUTJSON) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *UsageOUTJSON) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.


### SetCurrencySymbolNil

`func (o *UsageOUTJSON) SetCurrencySymbolNil(b bool)`

 SetCurrencySymbolNil sets the value for CurrencySymbol to be an explicit nil

### UnsetCurrencySymbol
`func (o *UsageOUTJSON) UnsetCurrencySymbol()`

UnsetCurrencySymbol ensures that no value is present for CurrencySymbol, not even an explicit nil
### GetPrices

`func (o *UsageOUTJSON) GetPrices() []PriceData`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *UsageOUTJSON) GetPricesOk() (*[]PriceData, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *UsageOUTJSON) SetPrices(v []PriceData)`

SetPrices sets Prices field to given value.


### SetPricesNil

`func (o *UsageOUTJSON) SetPricesNil(b bool)`

 SetPricesNil sets the value for Prices to be an explicit nil

### UnsetPrices
`func (o *UsageOUTJSON) UnsetPrices()`

UnsetPrices ensures that no value is present for Prices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



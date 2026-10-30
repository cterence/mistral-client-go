# TurbineMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemPromptName** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to [**NullableServerLocale**](ServerLocale.md) |  | [optional] 

## Methods

### NewTurbineMeta

`func NewTurbineMeta() *TurbineMeta`

NewTurbineMeta instantiates a new TurbineMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTurbineMetaWithDefaults

`func NewTurbineMetaWithDefaults() *TurbineMeta`

NewTurbineMetaWithDefaults instantiates a new TurbineMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemPromptName

`func (o *TurbineMeta) GetSystemPromptName() string`

GetSystemPromptName returns the SystemPromptName field if non-nil, zero value otherwise.

### GetSystemPromptNameOk

`func (o *TurbineMeta) GetSystemPromptNameOk() (*string, bool)`

GetSystemPromptNameOk returns a tuple with the SystemPromptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPromptName

`func (o *TurbineMeta) SetSystemPromptName(v string)`

SetSystemPromptName sets SystemPromptName field to given value.

### HasSystemPromptName

`func (o *TurbineMeta) HasSystemPromptName() bool`

HasSystemPromptName returns a boolean if a field has been set.

### SetSystemPromptNameNil

`func (o *TurbineMeta) SetSystemPromptNameNil(b bool)`

 SetSystemPromptNameNil sets the value for SystemPromptName to be an explicit nil

### UnsetSystemPromptName
`func (o *TurbineMeta) UnsetSystemPromptName()`

UnsetSystemPromptName ensures that no value is present for SystemPromptName, not even an explicit nil
### GetLocale

`func (o *TurbineMeta) GetLocale() ServerLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TurbineMeta) GetLocaleOk() (*ServerLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TurbineMeta) SetLocale(v ServerLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TurbineMeta) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *TurbineMeta) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *TurbineMeta) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



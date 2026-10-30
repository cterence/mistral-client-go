# ResponsesTextConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**NullableFormat1**](Format1.md) |  | [optional] 
**Verbosity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesTextConfig

`func NewResponsesTextConfig() *ResponsesTextConfig`

NewResponsesTextConfig instantiates a new ResponsesTextConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTextConfigWithDefaults

`func NewResponsesTextConfigWithDefaults() *ResponsesTextConfig`

NewResponsesTextConfigWithDefaults instantiates a new ResponsesTextConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ResponsesTextConfig) GetFormat() Format1`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ResponsesTextConfig) GetFormatOk() (*Format1, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ResponsesTextConfig) SetFormat(v Format1)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ResponsesTextConfig) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *ResponsesTextConfig) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *ResponsesTextConfig) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetVerbosity

`func (o *ResponsesTextConfig) GetVerbosity() string`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *ResponsesTextConfig) GetVerbosityOk() (*string, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *ResponsesTextConfig) SetVerbosity(v string)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *ResponsesTextConfig) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.

### SetVerbosityNil

`func (o *ResponsesTextConfig) SetVerbosityNil(b bool)`

 SetVerbosityNil sets the value for Verbosity to be an explicit nil

### UnsetVerbosity
`func (o *ResponsesTextConfig) UnsetVerbosity()`

UnsetVerbosity ensures that no value is present for Verbosity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



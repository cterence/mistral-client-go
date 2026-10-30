# ElicitationCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to **map[string]interface{}** | Capability for form mode elicitation. | [optional] 
**Url** | Pointer to **map[string]interface{}** | Capability for URL mode elicitation (2025-11-25+). | [optional] 

## Methods

### NewElicitationCapability

`func NewElicitationCapability() *ElicitationCapability`

NewElicitationCapability instantiates a new ElicitationCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElicitationCapabilityWithDefaults

`func NewElicitationCapabilityWithDefaults() *ElicitationCapability`

NewElicitationCapabilityWithDefaults instantiates a new ElicitationCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *ElicitationCapability) GetForm() map[string]interface{}`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ElicitationCapability) GetFormOk() (*map[string]interface{}, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ElicitationCapability) SetForm(v map[string]interface{})`

SetForm sets Form field to given value.

### HasForm

`func (o *ElicitationCapability) HasForm() bool`

HasForm returns a boolean if a field has been set.

### SetFormNil

`func (o *ElicitationCapability) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *ElicitationCapability) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil
### GetUrl

`func (o *ElicitationCapability) GetUrl() map[string]interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ElicitationCapability) GetUrlOk() (*map[string]interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ElicitationCapability) SetUrl(v map[string]interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ElicitationCapability) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ElicitationCapability) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ElicitationCapability) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



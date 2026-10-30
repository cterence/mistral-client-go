# MCPServerRemoteHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Header name | 
**Description** | **string** | Human-readable description of the header | 
**IsRequired** | Pointer to **NullableBool** |  | [optional] 
**IsSecret** | Pointer to **NullableBool** |  | [optional] 
**Default** | Pointer to **NullableString** |  | [optional] 
**Choices** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMCPServerRemoteHeader

`func NewMCPServerRemoteHeader(name string, description string, ) *MCPServerRemoteHeader`

NewMCPServerRemoteHeader instantiates a new MCPServerRemoteHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPServerRemoteHeaderWithDefaults

`func NewMCPServerRemoteHeaderWithDefaults() *MCPServerRemoteHeader`

NewMCPServerRemoteHeaderWithDefaults instantiates a new MCPServerRemoteHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MCPServerRemoteHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPServerRemoteHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPServerRemoteHeader) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MCPServerRemoteHeader) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MCPServerRemoteHeader) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MCPServerRemoteHeader) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsRequired

`func (o *MCPServerRemoteHeader) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *MCPServerRemoteHeader) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *MCPServerRemoteHeader) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *MCPServerRemoteHeader) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *MCPServerRemoteHeader) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *MCPServerRemoteHeader) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetIsSecret

`func (o *MCPServerRemoteHeader) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *MCPServerRemoteHeader) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *MCPServerRemoteHeader) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *MCPServerRemoteHeader) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### SetIsSecretNil

`func (o *MCPServerRemoteHeader) SetIsSecretNil(b bool)`

 SetIsSecretNil sets the value for IsSecret to be an explicit nil

### UnsetIsSecret
`func (o *MCPServerRemoteHeader) UnsetIsSecret()`

UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
### GetDefault

`func (o *MCPServerRemoteHeader) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *MCPServerRemoteHeader) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *MCPServerRemoteHeader) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *MCPServerRemoteHeader) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *MCPServerRemoteHeader) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *MCPServerRemoteHeader) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetChoices

`func (o *MCPServerRemoteHeader) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *MCPServerRemoteHeader) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *MCPServerRemoteHeader) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *MCPServerRemoteHeader) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *MCPServerRemoteHeader) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *MCPServerRemoteHeader) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



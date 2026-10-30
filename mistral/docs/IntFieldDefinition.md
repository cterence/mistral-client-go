# IntFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexed** | Pointer to **bool** |  | [optional] [default to false]
**Required** | Pointer to **bool** |  | [optional] [default to false]
**System** | Pointer to **bool** |  | [optional] [default to false]
**Default** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "int"]

## Methods

### NewIntFieldDefinition

`func NewIntFieldDefinition() *IntFieldDefinition`

NewIntFieldDefinition instantiates a new IntFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntFieldDefinitionWithDefaults

`func NewIntFieldDefinitionWithDefaults() *IntFieldDefinition`

NewIntFieldDefinitionWithDefaults instantiates a new IntFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexed

`func (o *IntFieldDefinition) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *IntFieldDefinition) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *IntFieldDefinition) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *IntFieldDefinition) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetRequired

`func (o *IntFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IntFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IntFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IntFieldDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSystem

`func (o *IntFieldDefinition) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IntFieldDefinition) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IntFieldDefinition) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IntFieldDefinition) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetDefault

`func (o *IntFieldDefinition) GetDefault() AnyOf`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IntFieldDefinition) GetDefaultOk() (*AnyOf, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IntFieldDefinition) SetDefault(v AnyOf)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IntFieldDefinition) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *IntFieldDefinition) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *IntFieldDefinition) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetType

`func (o *IntFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntFieldDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntFieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



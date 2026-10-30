# JsonFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexed** | Pointer to **bool** |  | [optional] [default to false]
**Required** | Pointer to **bool** |  | [optional] [default to false]
**System** | Pointer to **bool** |  | [optional] [default to false]
**Default** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "json"]

## Methods

### NewJsonFieldDefinition

`func NewJsonFieldDefinition() *JsonFieldDefinition`

NewJsonFieldDefinition instantiates a new JsonFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonFieldDefinitionWithDefaults

`func NewJsonFieldDefinitionWithDefaults() *JsonFieldDefinition`

NewJsonFieldDefinitionWithDefaults instantiates a new JsonFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexed

`func (o *JsonFieldDefinition) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *JsonFieldDefinition) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *JsonFieldDefinition) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *JsonFieldDefinition) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetRequired

`func (o *JsonFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *JsonFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *JsonFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *JsonFieldDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSystem

`func (o *JsonFieldDefinition) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *JsonFieldDefinition) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *JsonFieldDefinition) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *JsonFieldDefinition) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetDefault

`func (o *JsonFieldDefinition) GetDefault() AnyOf`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *JsonFieldDefinition) GetDefaultOk() (*AnyOf, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *JsonFieldDefinition) SetDefault(v AnyOf)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *JsonFieldDefinition) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *JsonFieldDefinition) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *JsonFieldDefinition) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetType

`func (o *JsonFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonFieldDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonFieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



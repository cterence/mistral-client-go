# DocumentFieldsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexed** | Pointer to **bool** |  | [optional] [default to false]
**Required** | Pointer to **bool** |  | [optional] [default to false]
**System** | Pointer to **bool** |  | [optional] [default to false]
**Default** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "text"]

## Methods

### NewDocumentFieldsValue

`func NewDocumentFieldsValue() *DocumentFieldsValue`

NewDocumentFieldsValue instantiates a new DocumentFieldsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFieldsValueWithDefaults

`func NewDocumentFieldsValueWithDefaults() *DocumentFieldsValue`

NewDocumentFieldsValueWithDefaults instantiates a new DocumentFieldsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexed

`func (o *DocumentFieldsValue) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *DocumentFieldsValue) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *DocumentFieldsValue) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *DocumentFieldsValue) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetRequired

`func (o *DocumentFieldsValue) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DocumentFieldsValue) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DocumentFieldsValue) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *DocumentFieldsValue) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSystem

`func (o *DocumentFieldsValue) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *DocumentFieldsValue) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *DocumentFieldsValue) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *DocumentFieldsValue) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetDefault

`func (o *DocumentFieldsValue) GetDefault() AnyOf`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DocumentFieldsValue) GetDefaultOk() (*AnyOf, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DocumentFieldsValue) SetDefault(v AnyOf)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DocumentFieldsValue) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *DocumentFieldsValue) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *DocumentFieldsValue) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetType

`func (o *DocumentFieldsValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentFieldsValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentFieldsValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentFieldsValue) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



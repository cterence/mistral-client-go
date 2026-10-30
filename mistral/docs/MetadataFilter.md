# MetadataFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Operator** | [**SystemParamOperator**](SystemParamOperator.md) |  | 
**Value** | [**Value**](Value.md) |  | 

## Methods

### NewMetadataFilter

`func NewMetadataFilter(key string, operator SystemParamOperator, value Value, ) *MetadataFilter`

NewMetadataFilter instantiates a new MetadataFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFilterWithDefaults

`func NewMetadataFilterWithDefaults() *MetadataFilter`

NewMetadataFilterWithDefaults instantiates a new MetadataFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetadataFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetadataFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetadataFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *MetadataFilter) GetOperator() SystemParamOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetadataFilter) GetOperatorOk() (*SystemParamOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetadataFilter) SetOperator(v SystemParamOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *MetadataFilter) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetadataFilter) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetadataFilter) SetValue(v Value)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



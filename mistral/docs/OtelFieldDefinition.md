# OtelFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Type** | **string** |  | 
**Group** | Pointer to **NullableString** |  | [optional] 
**SourceAttributeKey** | Pointer to **NullableString** | Raw OpenTelemetry attribute key this field is projected from, whether a span attribute (e.g. &#x60;gen_ai.request.model&#x60; for &#x60;request_model&#x60;) or a resource attribute (e.g. &#x60;service.name&#x60; for &#x60;service_name&#x60;). None when the field is a native OTLP field (like &#x60;span_name&#x60; or &#x60;trace_id&#x60;) or a nested attribute map (like &#x60;span_attributes&#x60;). | [optional] 
**SupportedOperators** | **[]string** |  | [readonly] 
**SupportedAggregations** | [**[]MetricAggregation**](MetricAggregation.md) |  | [readonly] 

## Methods

### NewOtelFieldDefinition

`func NewOtelFieldDefinition(name string, label string, type_ string, supportedOperators []string, supportedAggregations []MetricAggregation, ) *OtelFieldDefinition`

NewOtelFieldDefinition instantiates a new OtelFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelFieldDefinitionWithDefaults

`func NewOtelFieldDefinitionWithDefaults() *OtelFieldDefinition`

NewOtelFieldDefinitionWithDefaults instantiates a new OtelFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OtelFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtelFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtelFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *OtelFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OtelFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OtelFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *OtelFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OtelFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OtelFieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetGroup

`func (o *OtelFieldDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OtelFieldDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OtelFieldDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *OtelFieldDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *OtelFieldDefinition) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *OtelFieldDefinition) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSourceAttributeKey

`func (o *OtelFieldDefinition) GetSourceAttributeKey() string`

GetSourceAttributeKey returns the SourceAttributeKey field if non-nil, zero value otherwise.

### GetSourceAttributeKeyOk

`func (o *OtelFieldDefinition) GetSourceAttributeKeyOk() (*string, bool)`

GetSourceAttributeKeyOk returns a tuple with the SourceAttributeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeKey

`func (o *OtelFieldDefinition) SetSourceAttributeKey(v string)`

SetSourceAttributeKey sets SourceAttributeKey field to given value.

### HasSourceAttributeKey

`func (o *OtelFieldDefinition) HasSourceAttributeKey() bool`

HasSourceAttributeKey returns a boolean if a field has been set.

### SetSourceAttributeKeyNil

`func (o *OtelFieldDefinition) SetSourceAttributeKeyNil(b bool)`

 SetSourceAttributeKeyNil sets the value for SourceAttributeKey to be an explicit nil

### UnsetSourceAttributeKey
`func (o *OtelFieldDefinition) UnsetSourceAttributeKey()`

UnsetSourceAttributeKey ensures that no value is present for SourceAttributeKey, not even an explicit nil
### GetSupportedOperators

`func (o *OtelFieldDefinition) GetSupportedOperators() []string`

GetSupportedOperators returns the SupportedOperators field if non-nil, zero value otherwise.

### GetSupportedOperatorsOk

`func (o *OtelFieldDefinition) GetSupportedOperatorsOk() (*[]string, bool)`

GetSupportedOperatorsOk returns a tuple with the SupportedOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperators

`func (o *OtelFieldDefinition) SetSupportedOperators(v []string)`

SetSupportedOperators sets SupportedOperators field to given value.


### GetSupportedAggregations

`func (o *OtelFieldDefinition) GetSupportedAggregations() []MetricAggregation`

GetSupportedAggregations returns the SupportedAggregations field if non-nil, zero value otherwise.

### GetSupportedAggregationsOk

`func (o *OtelFieldDefinition) GetSupportedAggregationsOk() (*[]MetricAggregation, bool)`

GetSupportedAggregationsOk returns a tuple with the SupportedAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAggregations

`func (o *OtelFieldDefinition) SetSupportedAggregations(v []MetricAggregation)`

SetSupportedAggregations sets SupportedAggregations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



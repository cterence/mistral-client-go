# MCPResultMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsError** | Pointer to **bool** |  | [optional] [default to false]
**StructuredContent** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMCPResultMetadata

`func NewMCPResultMetadata() *MCPResultMetadata`

NewMCPResultMetadata instantiates a new MCPResultMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPResultMetadataWithDefaults

`func NewMCPResultMetadataWithDefaults() *MCPResultMetadata`

NewMCPResultMetadataWithDefaults instantiates a new MCPResultMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsError

`func (o *MCPResultMetadata) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *MCPResultMetadata) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *MCPResultMetadata) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *MCPResultMetadata) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetStructuredContent

`func (o *MCPResultMetadata) GetStructuredContent() map[string]interface{}`

GetStructuredContent returns the StructuredContent field if non-nil, zero value otherwise.

### GetStructuredContentOk

`func (o *MCPResultMetadata) GetStructuredContentOk() (*map[string]interface{}, bool)`

GetStructuredContentOk returns a tuple with the StructuredContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredContent

`func (o *MCPResultMetadata) SetStructuredContent(v map[string]interface{})`

SetStructuredContent sets StructuredContent field to given value.

### HasStructuredContent

`func (o *MCPResultMetadata) HasStructuredContent() bool`

HasStructuredContent returns a boolean if a field has been set.

### SetStructuredContentNil

`func (o *MCPResultMetadata) SetStructuredContentNil(b bool)`

 SetStructuredContentNil sets the value for StructuredContent to be an explicit nil

### UnsetStructuredContent
`func (o *MCPResultMetadata) UnsetStructuredContent()`

UnsetStructuredContent ensures that no value is present for StructuredContent, not even an explicit nil
### GetMeta

`func (o *MCPResultMetadata) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MCPResultMetadata) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MCPResultMetadata) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MCPResultMetadata) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MCPResultMetadata) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MCPResultMetadata) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



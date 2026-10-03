# MCPToolCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]Content5Inner**](Content5Inner.md) |  | 
**Metadata** | Pointer to [**NullableMCPToolCallMetadata**](MCPToolCallMetadata.md) |  | [optional] 

## Methods

### NewMCPToolCallResponse

`func NewMCPToolCallResponse(content []Content5Inner, ) *MCPToolCallResponse`

NewMCPToolCallResponse instantiates a new MCPToolCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPToolCallResponseWithDefaults

`func NewMCPToolCallResponseWithDefaults() *MCPToolCallResponse`

NewMCPToolCallResponseWithDefaults instantiates a new MCPToolCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MCPToolCallResponse) GetContent() []Content5Inner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MCPToolCallResponse) GetContentOk() (*[]Content5Inner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MCPToolCallResponse) SetContent(v []Content5Inner)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *MCPToolCallResponse) GetMetadata() MCPToolCallMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MCPToolCallResponse) GetMetadataOk() (*MCPToolCallMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MCPToolCallResponse) SetMetadata(v MCPToolCallMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MCPToolCallResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *MCPToolCallResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *MCPToolCallResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



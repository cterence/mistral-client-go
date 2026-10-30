# ConnectorToolCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]Content6Inner**](Content6Inner.md) |  | 
**Metadata** | Pointer to [**NullableConnectorToolCallMetadata**](ConnectorToolCallMetadata.md) |  | [optional] 

## Methods

### NewConnectorToolCallResponse

`func NewConnectorToolCallResponse(content []Content6Inner, ) *ConnectorToolCallResponse`

NewConnectorToolCallResponse instantiates a new ConnectorToolCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorToolCallResponseWithDefaults

`func NewConnectorToolCallResponseWithDefaults() *ConnectorToolCallResponse`

NewConnectorToolCallResponseWithDefaults instantiates a new ConnectorToolCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ConnectorToolCallResponse) GetContent() []Content6Inner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConnectorToolCallResponse) GetContentOk() (*[]Content6Inner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConnectorToolCallResponse) SetContent(v []Content6Inner)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *ConnectorToolCallResponse) GetMetadata() ConnectorToolCallMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectorToolCallResponse) GetMetadataOk() (*ConnectorToolCallMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectorToolCallResponse) SetMetadata(v ConnectorToolCallMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectorToolCallResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ConnectorToolCallResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ConnectorToolCallResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



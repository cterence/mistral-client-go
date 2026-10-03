# AgentConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name given to the conversation. | [optional] 
**Description** | Pointer to **NullableString** | Description of the what the conversation is about. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata for the conversation. | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "conversation"]
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AgentId** | **string** |  | 
**AgentVersion** | Pointer to [**NullableAgentVersion1**](AgentVersion1.md) |  | [optional] 

## Methods

### NewAgentConversation

`func NewAgentConversation(id string, createdAt time.Time, updatedAt time.Time, agentId string, ) *AgentConversation`

NewAgentConversation instantiates a new AgentConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConversationWithDefaults

`func NewAgentConversationWithDefaults() *AgentConversation`

NewAgentConversationWithDefaults instantiates a new AgentConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentConversation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentConversation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentConversation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentConversation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentConversation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentConversation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AgentConversation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentConversation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentConversation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentConversation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AgentConversation) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AgentConversation) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *AgentConversation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AgentConversation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AgentConversation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AgentConversation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AgentConversation) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AgentConversation) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetObject

`func (o *AgentConversation) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AgentConversation) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AgentConversation) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AgentConversation) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetId

`func (o *AgentConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentConversation) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AgentConversation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentConversation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentConversation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgentConversation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentConversation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentConversation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAgentId

`func (o *AgentConversation) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentConversation) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentConversation) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAgentVersion

`func (o *AgentConversation) GetAgentVersion() AgentVersion1`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AgentConversation) GetAgentVersionOk() (*AgentVersion1, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AgentConversation) SetAgentVersion(v AgentVersion1)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *AgentConversation) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *AgentConversation) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *AgentConversation) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



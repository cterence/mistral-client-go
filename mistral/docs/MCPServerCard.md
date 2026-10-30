# MCPServerCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **NullableString** | URL to the JSON schema definition | [optional] 
**Name** | **string** | Server identifier in reverse-DNS format with exactly one / | 
**Version** | **string** | Server version (semantic versioning preferred) | 
**Capabilities** | Pointer to [**ServerCapabilities**](ServerCapabilities.md) |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to [**NullableMCPServerRepository**](MCPServerRepository.md) |  | [optional] 
**Icons** | Pointer to [**[]MCPServerIcon**](MCPServerIcon.md) |  | [optional] 
**Remotes** | Pointer to [**[]MCPServerRemote**](MCPServerRemote.md) |  | [optional] 
**Requires** | Pointer to [**NullableClientCapabilities**](ClientCapabilities.md) |  | [optional] 
**Resources** | Pointer to [**NullableResources**](Resources.md) |  | [optional] 
**Tools** | Pointer to [**NullableTools**](Tools.md) |  | [optional] 
**Prompts** | Pointer to [**NullablePrompts**](Prompts.md) |  | [optional] 
**Meta** | Pointer to [**NullableMCPServerCardMeta**](MCPServerCardMeta.md) |  | [optional] 

## Methods

### NewMCPServerCard

`func NewMCPServerCard(name string, version string, ) *MCPServerCard`

NewMCPServerCard instantiates a new MCPServerCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPServerCardWithDefaults

`func NewMCPServerCardWithDefaults() *MCPServerCard`

NewMCPServerCardWithDefaults instantiates a new MCPServerCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *MCPServerCard) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *MCPServerCard) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *MCPServerCard) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *MCPServerCard) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *MCPServerCard) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *MCPServerCard) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetName

`func (o *MCPServerCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPServerCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPServerCard) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *MCPServerCard) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MCPServerCard) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MCPServerCard) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCapabilities

`func (o *MCPServerCard) GetCapabilities() ServerCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MCPServerCard) GetCapabilitiesOk() (*ServerCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MCPServerCard) SetCapabilities(v ServerCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MCPServerCard) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetTitle

`func (o *MCPServerCard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MCPServerCard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MCPServerCard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MCPServerCard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MCPServerCard) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MCPServerCard) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *MCPServerCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MCPServerCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MCPServerCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MCPServerCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MCPServerCard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MCPServerCard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsiteUrl

`func (o *MCPServerCard) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *MCPServerCard) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *MCPServerCard) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *MCPServerCard) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *MCPServerCard) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *MCPServerCard) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetRepository

`func (o *MCPServerCard) GetRepository() MCPServerRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MCPServerCard) GetRepositoryOk() (*MCPServerRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MCPServerCard) SetRepository(v MCPServerRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *MCPServerCard) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *MCPServerCard) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *MCPServerCard) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetIcons

`func (o *MCPServerCard) GetIcons() []MCPServerIcon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *MCPServerCard) GetIconsOk() (*[]MCPServerIcon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *MCPServerCard) SetIcons(v []MCPServerIcon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *MCPServerCard) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### SetIconsNil

`func (o *MCPServerCard) SetIconsNil(b bool)`

 SetIconsNil sets the value for Icons to be an explicit nil

### UnsetIcons
`func (o *MCPServerCard) UnsetIcons()`

UnsetIcons ensures that no value is present for Icons, not even an explicit nil
### GetRemotes

`func (o *MCPServerCard) GetRemotes() []MCPServerRemote`

GetRemotes returns the Remotes field if non-nil, zero value otherwise.

### GetRemotesOk

`func (o *MCPServerCard) GetRemotesOk() (*[]MCPServerRemote, bool)`

GetRemotesOk returns a tuple with the Remotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotes

`func (o *MCPServerCard) SetRemotes(v []MCPServerRemote)`

SetRemotes sets Remotes field to given value.

### HasRemotes

`func (o *MCPServerCard) HasRemotes() bool`

HasRemotes returns a boolean if a field has been set.

### SetRemotesNil

`func (o *MCPServerCard) SetRemotesNil(b bool)`

 SetRemotesNil sets the value for Remotes to be an explicit nil

### UnsetRemotes
`func (o *MCPServerCard) UnsetRemotes()`

UnsetRemotes ensures that no value is present for Remotes, not even an explicit nil
### GetRequires

`func (o *MCPServerCard) GetRequires() ClientCapabilities`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *MCPServerCard) GetRequiresOk() (*ClientCapabilities, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *MCPServerCard) SetRequires(v ClientCapabilities)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *MCPServerCard) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### SetRequiresNil

`func (o *MCPServerCard) SetRequiresNil(b bool)`

 SetRequiresNil sets the value for Requires to be an explicit nil

### UnsetRequires
`func (o *MCPServerCard) UnsetRequires()`

UnsetRequires ensures that no value is present for Requires, not even an explicit nil
### GetResources

`func (o *MCPServerCard) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MCPServerCard) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MCPServerCard) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MCPServerCard) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *MCPServerCard) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *MCPServerCard) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetTools

`func (o *MCPServerCard) GetTools() Tools`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *MCPServerCard) GetToolsOk() (*Tools, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *MCPServerCard) SetTools(v Tools)`

SetTools sets Tools field to given value.

### HasTools

`func (o *MCPServerCard) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *MCPServerCard) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *MCPServerCard) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetPrompts

`func (o *MCPServerCard) GetPrompts() Prompts`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *MCPServerCard) GetPromptsOk() (*Prompts, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *MCPServerCard) SetPrompts(v Prompts)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *MCPServerCard) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### SetPromptsNil

`func (o *MCPServerCard) SetPromptsNil(b bool)`

 SetPromptsNil sets the value for Prompts to be an explicit nil

### UnsetPrompts
`func (o *MCPServerCard) UnsetPrompts()`

UnsetPrompts ensures that no value is present for Prompts, not even an explicit nil
### GetMeta

`func (o *MCPServerCard) GetMeta() MCPServerCardMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MCPServerCard) GetMetaOk() (*MCPServerCardMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MCPServerCard) SetMeta(v MCPServerCardMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MCPServerCard) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MCPServerCard) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MCPServerCard) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



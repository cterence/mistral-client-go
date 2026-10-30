# MCPServerRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Repository URL | 
**Source** | **string** | Source identifier (e.g. &#39;github&#39;) | 
**Subfolder** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMCPServerRepository

`func NewMCPServerRepository(url string, source string, ) *MCPServerRepository`

NewMCPServerRepository instantiates a new MCPServerRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPServerRepositoryWithDefaults

`func NewMCPServerRepositoryWithDefaults() *MCPServerRepository`

NewMCPServerRepositoryWithDefaults instantiates a new MCPServerRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *MCPServerRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MCPServerRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MCPServerRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *MCPServerRepository) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MCPServerRepository) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MCPServerRepository) SetSource(v string)`

SetSource sets Source field to given value.


### GetSubfolder

`func (o *MCPServerRepository) GetSubfolder() string`

GetSubfolder returns the Subfolder field if non-nil, zero value otherwise.

### GetSubfolderOk

`func (o *MCPServerRepository) GetSubfolderOk() (*string, bool)`

GetSubfolderOk returns a tuple with the Subfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolder

`func (o *MCPServerRepository) SetSubfolder(v string)`

SetSubfolder sets Subfolder field to given value.

### HasSubfolder

`func (o *MCPServerRepository) HasSubfolder() bool`

HasSubfolder returns a boolean if a field has been set.

### SetSubfolderNil

`func (o *MCPServerRepository) SetSubfolderNil(b bool)`

 SetSubfolderNil sets the value for Subfolder to be an explicit nil

### UnsetSubfolder
`func (o *MCPServerRepository) UnsetSubfolder()`

UnsetSubfolder ensures that no value is present for Subfolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



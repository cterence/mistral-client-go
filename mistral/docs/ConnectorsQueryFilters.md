# ConnectorsQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsMcp** | Pointer to **NullableBool** | Filter connectors by whether they expose an MCP-compatible tool interface. | [optional] 

## Methods

### NewConnectorsQueryFilters

`func NewConnectorsQueryFilters() *ConnectorsQueryFilters`

NewConnectorsQueryFilters instantiates a new ConnectorsQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorsQueryFiltersWithDefaults

`func NewConnectorsQueryFiltersWithDefaults() *ConnectorsQueryFilters`

NewConnectorsQueryFiltersWithDefaults instantiates a new ConnectorsQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsMcp

`func (o *ConnectorsQueryFilters) GetSupportsMcp() bool`

GetSupportsMcp returns the SupportsMcp field if non-nil, zero value otherwise.

### GetSupportsMcpOk

`func (o *ConnectorsQueryFilters) GetSupportsMcpOk() (*bool, bool)`

GetSupportsMcpOk returns a tuple with the SupportsMcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMcp

`func (o *ConnectorsQueryFilters) SetSupportsMcp(v bool)`

SetSupportsMcp sets SupportsMcp field to given value.

### HasSupportsMcp

`func (o *ConnectorsQueryFilters) HasSupportsMcp() bool`

HasSupportsMcp returns a boolean if a field has been set.

### SetSupportsMcpNil

`func (o *ConnectorsQueryFilters) SetSupportsMcpNil(b bool)`

 SetSupportsMcpNil sets the value for SupportsMcp to be an explicit nil

### UnsetSupportsMcp
`func (o *ConnectorsQueryFilters) UnsetSupportsMcp()`

UnsetSupportsMcp ensures that no value is present for SupportsMcp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



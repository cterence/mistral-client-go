# OTLPDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**Endpoint** | **string** |  | 
**Insecure** | Pointer to **bool** |  | [optional] [default to false]
**Headers** | Pointer to [**map[string]PipelineConfigHeader**](PipelineConfigHeader.md) |  | [optional] 

## Methods

### NewOTLPDestination

`func NewOTLPDestination(protocol string, endpoint string, ) *OTLPDestination`

NewOTLPDestination instantiates a new OTLPDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTLPDestinationWithDefaults

`func NewOTLPDestinationWithDefaults() *OTLPDestination`

NewOTLPDestinationWithDefaults instantiates a new OTLPDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *OTLPDestination) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OTLPDestination) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OTLPDestination) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetEndpoint

`func (o *OTLPDestination) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *OTLPDestination) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *OTLPDestination) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetInsecure

`func (o *OTLPDestination) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *OTLPDestination) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *OTLPDestination) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *OTLPDestination) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetHeaders

`func (o *OTLPDestination) GetHeaders() map[string]PipelineConfigHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OTLPDestination) GetHeadersOk() (*map[string]PipelineConfigHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OTLPDestination) SetHeaders(v map[string]PipelineConfigHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OTLPDestination) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



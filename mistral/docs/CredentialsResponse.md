# CredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]AuthenticationConfiguration**](AuthenticationConfiguration.md) |  | 
**ConnectorPresetCredentialsForAuth** | Pointer to [**[]OutboundAuthenticationType**](OutboundAuthenticationType.md) |  | [optional] [default to {}]

## Methods

### NewCredentialsResponse

`func NewCredentialsResponse(credentials []AuthenticationConfiguration, ) *CredentialsResponse`

NewCredentialsResponse instantiates a new CredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResponseWithDefaults

`func NewCredentialsResponseWithDefaults() *CredentialsResponse`

NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CredentialsResponse) GetCredentials() []AuthenticationConfiguration`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialsResponse) GetCredentialsOk() (*[]AuthenticationConfiguration, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialsResponse) SetCredentials(v []AuthenticationConfiguration)`

SetCredentials sets Credentials field to given value.


### GetConnectorPresetCredentialsForAuth

`func (o *CredentialsResponse) GetConnectorPresetCredentialsForAuth() []OutboundAuthenticationType`

GetConnectorPresetCredentialsForAuth returns the ConnectorPresetCredentialsForAuth field if non-nil, zero value otherwise.

### GetConnectorPresetCredentialsForAuthOk

`func (o *CredentialsResponse) GetConnectorPresetCredentialsForAuthOk() (*[]OutboundAuthenticationType, bool)`

GetConnectorPresetCredentialsForAuthOk returns a tuple with the ConnectorPresetCredentialsForAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPresetCredentialsForAuth

`func (o *CredentialsResponse) SetConnectorPresetCredentialsForAuth(v []OutboundAuthenticationType)`

SetConnectorPresetCredentialsForAuth sets ConnectorPresetCredentialsForAuth field to given value.

### HasConnectorPresetCredentialsForAuth

`func (o *CredentialsResponse) HasConnectorPresetCredentialsForAuth() bool`

HasConnectorPresetCredentialsForAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConnectorsQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **NullableBool** | Filter for active connectors for a given user, workspace and organization. | [optional] 
**FetchConnectionSecrets** | Pointer to **bool** | Fetch connection secrets. | [optional] [default to false]

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

### GetActive

`func (o *ConnectorsQueryFilters) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorsQueryFilters) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorsQueryFilters) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorsQueryFilters) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ConnectorsQueryFilters) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ConnectorsQueryFilters) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetFetchConnectionSecrets

`func (o *ConnectorsQueryFilters) GetFetchConnectionSecrets() bool`

GetFetchConnectionSecrets returns the FetchConnectionSecrets field if non-nil, zero value otherwise.

### GetFetchConnectionSecretsOk

`func (o *ConnectorsQueryFilters) GetFetchConnectionSecretsOk() (*bool, bool)`

GetFetchConnectionSecretsOk returns a tuple with the FetchConnectionSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchConnectionSecrets

`func (o *ConnectorsQueryFilters) SetFetchConnectionSecrets(v bool)`

SetFetchConnectionSecrets sets FetchConnectionSecrets field to given value.

### HasFetchConnectionSecrets

`func (o *ConnectorsQueryFilters) HasFetchConnectionSecrets() bool`

HasFetchConnectionSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



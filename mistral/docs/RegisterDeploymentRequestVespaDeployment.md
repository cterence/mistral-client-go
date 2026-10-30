# RegisterDeploymentRequestVespaDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "vespa"]
**VespaVersion** | **string** |  | 
**Indexes** | [**[]RegisterDeploymentRequestVespaIndex**](RegisterDeploymentRequestVespaIndex.md) |  | 
**QueryUrl** | **string** |  | 

## Methods

### NewRegisterDeploymentRequestVespaDeployment

`func NewRegisterDeploymentRequestVespaDeployment(vespaVersion string, indexes []RegisterDeploymentRequestVespaIndex, queryUrl string, ) *RegisterDeploymentRequestVespaDeployment`

NewRegisterDeploymentRequestVespaDeployment instantiates a new RegisterDeploymentRequestVespaDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeploymentRequestVespaDeploymentWithDefaults

`func NewRegisterDeploymentRequestVespaDeploymentWithDefaults() *RegisterDeploymentRequestVespaDeployment`

NewRegisterDeploymentRequestVespaDeploymentWithDefaults instantiates a new RegisterDeploymentRequestVespaDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RegisterDeploymentRequestVespaDeployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterDeploymentRequestVespaDeployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterDeploymentRequestVespaDeployment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegisterDeploymentRequestVespaDeployment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVespaVersion

`func (o *RegisterDeploymentRequestVespaDeployment) GetVespaVersion() string`

GetVespaVersion returns the VespaVersion field if non-nil, zero value otherwise.

### GetVespaVersionOk

`func (o *RegisterDeploymentRequestVespaDeployment) GetVespaVersionOk() (*string, bool)`

GetVespaVersionOk returns a tuple with the VespaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVespaVersion

`func (o *RegisterDeploymentRequestVespaDeployment) SetVespaVersion(v string)`

SetVespaVersion sets VespaVersion field to given value.


### GetIndexes

`func (o *RegisterDeploymentRequestVespaDeployment) GetIndexes() []RegisterDeploymentRequestVespaIndex`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *RegisterDeploymentRequestVespaDeployment) GetIndexesOk() (*[]RegisterDeploymentRequestVespaIndex, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *RegisterDeploymentRequestVespaDeployment) SetIndexes(v []RegisterDeploymentRequestVespaIndex)`

SetIndexes sets Indexes field to given value.


### GetQueryUrl

`func (o *RegisterDeploymentRequestVespaDeployment) GetQueryUrl() string`

GetQueryUrl returns the QueryUrl field if non-nil, zero value otherwise.

### GetQueryUrlOk

`func (o *RegisterDeploymentRequestVespaDeployment) GetQueryUrlOk() (*string, bool)`

GetQueryUrlOk returns a tuple with the QueryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryUrl

`func (o *RegisterDeploymentRequestVespaDeployment) SetQueryUrl(v string)`

SetQueryUrl sets QueryUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



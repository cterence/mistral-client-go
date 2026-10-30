# GetDeploymentSummariesResponseVespaDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "vespa"]
**Indexes** | [**[]GetDeploymentSummariesResponseVespaIndex**](GetDeploymentSummariesResponseVespaIndex.md) |  | 

## Methods

### NewGetDeploymentSummariesResponseVespaDeployment

`func NewGetDeploymentSummariesResponseVespaDeployment(indexes []GetDeploymentSummariesResponseVespaIndex, ) *GetDeploymentSummariesResponseVespaDeployment`

NewGetDeploymentSummariesResponseVespaDeployment instantiates a new GetDeploymentSummariesResponseVespaDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentSummariesResponseVespaDeploymentWithDefaults

`func NewGetDeploymentSummariesResponseVespaDeploymentWithDefaults() *GetDeploymentSummariesResponseVespaDeployment`

NewGetDeploymentSummariesResponseVespaDeploymentWithDefaults instantiates a new GetDeploymentSummariesResponseVespaDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetDeploymentSummariesResponseVespaDeployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeploymentSummariesResponseVespaDeployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeploymentSummariesResponseVespaDeployment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeploymentSummariesResponseVespaDeployment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIndexes

`func (o *GetDeploymentSummariesResponseVespaDeployment) GetIndexes() []GetDeploymentSummariesResponseVespaIndex`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *GetDeploymentSummariesResponseVespaDeployment) GetIndexesOk() (*[]GetDeploymentSummariesResponseVespaIndex, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *GetDeploymentSummariesResponseVespaDeployment) SetIndexes(v []GetDeploymentSummariesResponseVespaIndex)`

SetIndexes sets Indexes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



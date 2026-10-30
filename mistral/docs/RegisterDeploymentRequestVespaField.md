# RegisterDeploymentRequestVespaField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**SchemaFieldDataType**](SchemaFieldDataType.md) |  | 
**Storage** | [**SchemaFieldStorage**](SchemaFieldStorage.md) |  | 
**Ranking** | [**SchemaFieldRankingType**](SchemaFieldRankingType.md) |  | 
**IndexType** | [**NullableSchemaFieldIndex**](SchemaFieldIndex.md) |  | 
**Multidimensional** | **bool** |  | 

## Methods

### NewRegisterDeploymentRequestVespaField

`func NewRegisterDeploymentRequestVespaField(name string, type_ SchemaFieldDataType, storage SchemaFieldStorage, ranking SchemaFieldRankingType, indexType NullableSchemaFieldIndex, multidimensional bool, ) *RegisterDeploymentRequestVespaField`

NewRegisterDeploymentRequestVespaField instantiates a new RegisterDeploymentRequestVespaField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeploymentRequestVespaFieldWithDefaults

`func NewRegisterDeploymentRequestVespaFieldWithDefaults() *RegisterDeploymentRequestVespaField`

NewRegisterDeploymentRequestVespaFieldWithDefaults instantiates a new RegisterDeploymentRequestVespaField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterDeploymentRequestVespaField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterDeploymentRequestVespaField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterDeploymentRequestVespaField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RegisterDeploymentRequestVespaField) GetType() SchemaFieldDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterDeploymentRequestVespaField) GetTypeOk() (*SchemaFieldDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterDeploymentRequestVespaField) SetType(v SchemaFieldDataType)`

SetType sets Type field to given value.


### GetStorage

`func (o *RegisterDeploymentRequestVespaField) GetStorage() SchemaFieldStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *RegisterDeploymentRequestVespaField) GetStorageOk() (*SchemaFieldStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *RegisterDeploymentRequestVespaField) SetStorage(v SchemaFieldStorage)`

SetStorage sets Storage field to given value.


### GetRanking

`func (o *RegisterDeploymentRequestVespaField) GetRanking() SchemaFieldRankingType`

GetRanking returns the Ranking field if non-nil, zero value otherwise.

### GetRankingOk

`func (o *RegisterDeploymentRequestVespaField) GetRankingOk() (*SchemaFieldRankingType, bool)`

GetRankingOk returns a tuple with the Ranking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanking

`func (o *RegisterDeploymentRequestVespaField) SetRanking(v SchemaFieldRankingType)`

SetRanking sets Ranking field to given value.


### GetIndexType

`func (o *RegisterDeploymentRequestVespaField) GetIndexType() SchemaFieldIndex`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *RegisterDeploymentRequestVespaField) GetIndexTypeOk() (*SchemaFieldIndex, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *RegisterDeploymentRequestVespaField) SetIndexType(v SchemaFieldIndex)`

SetIndexType sets IndexType field to given value.


### SetIndexTypeNil

`func (o *RegisterDeploymentRequestVespaField) SetIndexTypeNil(b bool)`

 SetIndexTypeNil sets the value for IndexType to be an explicit nil

### UnsetIndexType
`func (o *RegisterDeploymentRequestVespaField) UnsetIndexType()`

UnsetIndexType ensures that no value is present for IndexType, not even an explicit nil
### GetMultidimensional

`func (o *RegisterDeploymentRequestVespaField) GetMultidimensional() bool`

GetMultidimensional returns the Multidimensional field if non-nil, zero value otherwise.

### GetMultidimensionalOk

`func (o *RegisterDeploymentRequestVespaField) GetMultidimensionalOk() (*bool, bool)`

GetMultidimensionalOk returns a tuple with the Multidimensional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultidimensional

`func (o *RegisterDeploymentRequestVespaField) SetMultidimensional(v bool)`

SetMultidimensional sets Multidimensional field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



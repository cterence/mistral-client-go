# GetDeploymentSummariesResponseDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatorId** | **string** |  | 
**DocumentCount** | **int32** |  | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Deployment** | [**GetDeploymentSummariesResponseVespaDeployment**](GetDeploymentSummariesResponseVespaDeployment.md) |  | 

## Methods

### NewGetDeploymentSummariesResponseDeployment

`func NewGetDeploymentSummariesResponseDeployment(id string, name string, creatorId string, documentCount int32, status string, createdAt time.Time, modifiedAt time.Time, deployment GetDeploymentSummariesResponseVespaDeployment, ) *GetDeploymentSummariesResponseDeployment`

NewGetDeploymentSummariesResponseDeployment instantiates a new GetDeploymentSummariesResponseDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentSummariesResponseDeploymentWithDefaults

`func NewGetDeploymentSummariesResponseDeploymentWithDefaults() *GetDeploymentSummariesResponseDeployment`

NewGetDeploymentSummariesResponseDeploymentWithDefaults instantiates a new GetDeploymentSummariesResponseDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeploymentSummariesResponseDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeploymentSummariesResponseDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeploymentSummariesResponseDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetDeploymentSummariesResponseDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeploymentSummariesResponseDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeploymentSummariesResponseDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetCreatorId

`func (o *GetDeploymentSummariesResponseDeployment) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *GetDeploymentSummariesResponseDeployment) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *GetDeploymentSummariesResponseDeployment) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetDocumentCount

`func (o *GetDeploymentSummariesResponseDeployment) GetDocumentCount() int32`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *GetDeploymentSummariesResponseDeployment) GetDocumentCountOk() (*int32, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *GetDeploymentSummariesResponseDeployment) SetDocumentCount(v int32)`

SetDocumentCount sets DocumentCount field to given value.


### GetStatus

`func (o *GetDeploymentSummariesResponseDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDeploymentSummariesResponseDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDeploymentSummariesResponseDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *GetDeploymentSummariesResponseDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetDeploymentSummariesResponseDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetDeploymentSummariesResponseDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *GetDeploymentSummariesResponseDeployment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GetDeploymentSummariesResponseDeployment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GetDeploymentSummariesResponseDeployment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetDeployment

`func (o *GetDeploymentSummariesResponseDeployment) GetDeployment() GetDeploymentSummariesResponseVespaDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *GetDeploymentSummariesResponseDeployment) GetDeploymentOk() (*GetDeploymentSummariesResponseVespaDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *GetDeploymentSummariesResponseDeployment) SetDeployment(v GetDeploymentSummariesResponseVespaDeployment)`

SetDeployment sets Deployment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



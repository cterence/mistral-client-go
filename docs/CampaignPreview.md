# CampaignPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**Description** | **string** |  | 
**MaxNbEvents** | **int32** |  | 
**SearchParams** | [**FilterPayload**](FilterPayload.md) |  | 
**Judge** | [**JudgePreview**](JudgePreview.md) |  | 

## Methods

### NewCampaignPreview

`func NewCampaignPreview(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, name string, ownerId string, workspaceId string, description string, maxNbEvents int32, searchParams FilterPayload, judge JudgePreview, ) *CampaignPreview`

NewCampaignPreview instantiates a new CampaignPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignPreviewWithDefaults

`func NewCampaignPreviewWithDefaults() *CampaignPreview`

NewCampaignPreviewWithDefaults instantiates a new CampaignPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignPreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignPreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignPreview) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CampaignPreview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignPreview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignPreview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CampaignPreview) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignPreview) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignPreview) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *CampaignPreview) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CampaignPreview) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CampaignPreview) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *CampaignPreview) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *CampaignPreview) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetName

`func (o *CampaignPreview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignPreview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignPreview) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *CampaignPreview) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CampaignPreview) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CampaignPreview) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetWorkspaceId

`func (o *CampaignPreview) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *CampaignPreview) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *CampaignPreview) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetDescription

`func (o *CampaignPreview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignPreview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignPreview) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMaxNbEvents

`func (o *CampaignPreview) GetMaxNbEvents() int32`

GetMaxNbEvents returns the MaxNbEvents field if non-nil, zero value otherwise.

### GetMaxNbEventsOk

`func (o *CampaignPreview) GetMaxNbEventsOk() (*int32, bool)`

GetMaxNbEventsOk returns a tuple with the MaxNbEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNbEvents

`func (o *CampaignPreview) SetMaxNbEvents(v int32)`

SetMaxNbEvents sets MaxNbEvents field to given value.


### GetSearchParams

`func (o *CampaignPreview) GetSearchParams() FilterPayload`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *CampaignPreview) GetSearchParamsOk() (*FilterPayload, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *CampaignPreview) SetSearchParams(v FilterPayload)`

SetSearchParams sets SearchParams field to given value.


### GetJudge

`func (o *CampaignPreview) GetJudge() JudgePreview`

GetJudge returns the Judge field if non-nil, zero value otherwise.

### GetJudgeOk

`func (o *CampaignPreview) GetJudgeOk() (*JudgePreview, bool)`

GetJudgeOk returns a tuple with the Judge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudge

`func (o *CampaignPreview) SetJudge(v JudgePreview)`

SetJudge sets Judge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



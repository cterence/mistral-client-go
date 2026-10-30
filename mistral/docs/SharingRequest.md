# SharingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** |  | [optional] 
**Level** | [**ShareEnum**](ShareEnum.md) |  | 
**ShareWithUuid** | **string** | The id of the entity (user, workspace or organization) to share with | 
**ShareWithType** | [**EntityType**](EntityType.md) |  | 

## Methods

### NewSharingRequest

`func NewSharingRequest(level ShareEnum, shareWithUuid string, shareWithType EntityType, ) *SharingRequest`

NewSharingRequest instantiates a new SharingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingRequestWithDefaults

`func NewSharingRequestWithDefaults() *SharingRequest`

NewSharingRequestWithDefaults instantiates a new SharingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SharingRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SharingRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SharingRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SharingRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *SharingRequest) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *SharingRequest) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetLevel

`func (o *SharingRequest) GetLevel() ShareEnum`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SharingRequest) GetLevelOk() (*ShareEnum, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SharingRequest) SetLevel(v ShareEnum)`

SetLevel sets Level field to given value.


### GetShareWithUuid

`func (o *SharingRequest) GetShareWithUuid() string`

GetShareWithUuid returns the ShareWithUuid field if non-nil, zero value otherwise.

### GetShareWithUuidOk

`func (o *SharingRequest) GetShareWithUuidOk() (*string, bool)`

GetShareWithUuidOk returns a tuple with the ShareWithUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithUuid

`func (o *SharingRequest) SetShareWithUuid(v string)`

SetShareWithUuid sets ShareWithUuid field to given value.


### GetShareWithType

`func (o *SharingRequest) GetShareWithType() EntityType`

GetShareWithType returns the ShareWithType field if non-nil, zero value otherwise.

### GetShareWithTypeOk

`func (o *SharingRequest) GetShareWithTypeOk() (*EntityType, bool)`

GetShareWithTypeOk returns a tuple with the ShareWithType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithType

`func (o *SharingRequest) SetShareWithType(v EntityType)`

SetShareWithType sets ShareWithType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



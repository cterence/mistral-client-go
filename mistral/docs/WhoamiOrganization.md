# WhoamiOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The organization&#39;s unique identifier. | 
**Name** | **string** | The organization&#39;s display name. | 
**OrgTier** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Internal** | Pointer to **bool** | Whether this is a Mistral-internal organization. | [optional] [default to false]
**Kind** | Pointer to **string** | The organization kind (e.g. &#39;S&#39; standard, &#39;C&#39; company internal, &#39;D&#39; dogfooding). | [optional] [default to "S"]

## Methods

### NewWhoamiOrganization

`func NewWhoamiOrganization(id string, name string, ) *WhoamiOrganization`

NewWhoamiOrganization instantiates a new WhoamiOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiOrganizationWithDefaults

`func NewWhoamiOrganizationWithDefaults() *WhoamiOrganization`

NewWhoamiOrganizationWithDefaults instantiates a new WhoamiOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhoamiOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhoamiOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhoamiOrganization) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WhoamiOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhoamiOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhoamiOrganization) SetName(v string)`

SetName sets Name field to given value.


### GetOrgTier

`func (o *WhoamiOrganization) GetOrgTier() string`

GetOrgTier returns the OrgTier field if non-nil, zero value otherwise.

### GetOrgTierOk

`func (o *WhoamiOrganization) GetOrgTierOk() (*string, bool)`

GetOrgTierOk returns a tuple with the OrgTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTier

`func (o *WhoamiOrganization) SetOrgTier(v string)`

SetOrgTier sets OrgTier field to given value.

### HasOrgTier

`func (o *WhoamiOrganization) HasOrgTier() bool`

HasOrgTier returns a boolean if a field has been set.

### SetOrgTierNil

`func (o *WhoamiOrganization) SetOrgTierNil(b bool)`

 SetOrgTierNil sets the value for OrgTier to be an explicit nil

### UnsetOrgTier
`func (o *WhoamiOrganization) UnsetOrgTier()`

UnsetOrgTier ensures that no value is present for OrgTier, not even an explicit nil
### GetInternal

`func (o *WhoamiOrganization) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WhoamiOrganization) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WhoamiOrganization) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WhoamiOrganization) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetKind

`func (o *WhoamiOrganization) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WhoamiOrganization) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WhoamiOrganization) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WhoamiOrganization) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MemberSubscriptionOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PlanType**](PlanType.md) | Type of subscription assigned to the member. | 
**Plan** | Pointer to **NullableString** | Plan assigned to the member. | [optional] 
**Status** | Pointer to [**NullableSubscriptionStatus**](SubscriptionStatus.md) | Current status of the member subscription. | [optional] 
**User** | Pointer to **NullableString** | User ID associated with the subscription. | [optional] 
**SelfService** | Pointer to **NullableBool** | Whether the subscription was created through self-service. | [optional] 
**MemberHasAccess** | Pointer to **NullableBool** | Whether the member currently has access to the plan. | [optional] 

## Methods

### NewMemberSubscriptionOUT

`func NewMemberSubscriptionOUT(type_ PlanType, ) *MemberSubscriptionOUT`

NewMemberSubscriptionOUT instantiates a new MemberSubscriptionOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSubscriptionOUTWithDefaults

`func NewMemberSubscriptionOUTWithDefaults() *MemberSubscriptionOUT`

NewMemberSubscriptionOUTWithDefaults instantiates a new MemberSubscriptionOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MemberSubscriptionOUT) GetType() PlanType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberSubscriptionOUT) GetTypeOk() (*PlanType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberSubscriptionOUT) SetType(v PlanType)`

SetType sets Type field to given value.


### GetPlan

`func (o *MemberSubscriptionOUT) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MemberSubscriptionOUT) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MemberSubscriptionOUT) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MemberSubscriptionOUT) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### SetPlanNil

`func (o *MemberSubscriptionOUT) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *MemberSubscriptionOUT) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetStatus

`func (o *MemberSubscriptionOUT) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemberSubscriptionOUT) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemberSubscriptionOUT) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MemberSubscriptionOUT) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MemberSubscriptionOUT) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MemberSubscriptionOUT) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUser

`func (o *MemberSubscriptionOUT) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MemberSubscriptionOUT) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MemberSubscriptionOUT) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MemberSubscriptionOUT) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MemberSubscriptionOUT) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MemberSubscriptionOUT) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSelfService

`func (o *MemberSubscriptionOUT) GetSelfService() bool`

GetSelfService returns the SelfService field if non-nil, zero value otherwise.

### GetSelfServiceOk

`func (o *MemberSubscriptionOUT) GetSelfServiceOk() (*bool, bool)`

GetSelfServiceOk returns a tuple with the SelfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfService

`func (o *MemberSubscriptionOUT) SetSelfService(v bool)`

SetSelfService sets SelfService field to given value.

### HasSelfService

`func (o *MemberSubscriptionOUT) HasSelfService() bool`

HasSelfService returns a boolean if a field has been set.

### SetSelfServiceNil

`func (o *MemberSubscriptionOUT) SetSelfServiceNil(b bool)`

 SetSelfServiceNil sets the value for SelfService to be an explicit nil

### UnsetSelfService
`func (o *MemberSubscriptionOUT) UnsetSelfService()`

UnsetSelfService ensures that no value is present for SelfService, not even an explicit nil
### GetMemberHasAccess

`func (o *MemberSubscriptionOUT) GetMemberHasAccess() bool`

GetMemberHasAccess returns the MemberHasAccess field if non-nil, zero value otherwise.

### GetMemberHasAccessOk

`func (o *MemberSubscriptionOUT) GetMemberHasAccessOk() (*bool, bool)`

GetMemberHasAccessOk returns a tuple with the MemberHasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHasAccess

`func (o *MemberSubscriptionOUT) SetMemberHasAccess(v bool)`

SetMemberHasAccess sets MemberHasAccess field to given value.

### HasMemberHasAccess

`func (o *MemberSubscriptionOUT) HasMemberHasAccess() bool`

HasMemberHasAccess returns a boolean if a field has been set.

### SetMemberHasAccessNil

`func (o *MemberSubscriptionOUT) SetMemberHasAccessNil(b bool)`

 SetMemberHasAccessNil sets the value for MemberHasAccess to be an explicit nil

### UnsetMemberHasAccess
`func (o *MemberSubscriptionOUT) UnsetMemberHasAccess()`

UnsetMemberHasAccess ensures that no value is present for MemberHasAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



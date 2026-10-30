# UsageQueryRowOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **NullableString** | Filter by agent name | [optional] 
**UsageType** | Pointer to **NullableString** | Filter by agent name | [optional] 
**BillingMetric** | Pointer to **NullableString** | Filter by agent name | [optional] 
**BillingGroup** | Pointer to **NullableString** | Filter by agent name | [optional] 
**WorkspaceId** | Pointer to **NullableString** | Filter by agent name | [optional] 
**OrganizationId** | Pointer to **NullableString** | Filter by agent name | [optional] 
**UserId** | Pointer to **NullableString** | Filter by agent name | [optional] 
**ServiceAccountId** | Pointer to **NullableString** | Filter by agent name | [optional] 
**ApiKeyId** | Pointer to **NullableString** | Filter by agent name | [optional] 
**ApiZone** | Pointer to **NullableString** | Filter by agent name | [optional] 
**ServiceTier** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Timestamp** | Pointer to **NullableTime** | Only return logs at or after this timestamp | [optional] 
**TimeBucket** | Pointer to **NullableTime** | Only return logs at or after this timestamp | [optional] 
**Value** | Pointer to **NullableInt32** | Total usage value for this row. | [optional] 
**Count** | Pointer to **NullableInt32** | Total usage value for this row. | [optional] 
**NbApiCalls** | Pointer to **NullableInt32** | Total usage value for this row. | [optional] 
**BillingDisplayName** | Pointer to **NullableString** | Filter by agent name | [optional] 
**CustomerTier** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Untracked** | Pointer to **NullableBool** | Deprecated: this parameter will be removed in a future version. | [optional] 
**CallSource** | Pointer to **NullableString** | Filter by agent name | [optional] 
**CallType** | Pointer to **NullableString** | Filter by agent name | [optional] 

## Methods

### NewUsageQueryRowOUT

`func NewUsageQueryRowOUT() *UsageQueryRowOUT`

NewUsageQueryRowOUT instantiates a new UsageQueryRowOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageQueryRowOUTWithDefaults

`func NewUsageQueryRowOUTWithDefaults() *UsageQueryRowOUT`

NewUsageQueryRowOUTWithDefaults instantiates a new UsageQueryRowOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *UsageQueryRowOUT) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UsageQueryRowOUT) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UsageQueryRowOUT) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UsageQueryRowOUT) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *UsageQueryRowOUT) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *UsageQueryRowOUT) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetUsageType

`func (o *UsageQueryRowOUT) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *UsageQueryRowOUT) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *UsageQueryRowOUT) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *UsageQueryRowOUT) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *UsageQueryRowOUT) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *UsageQueryRowOUT) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetBillingMetric

`func (o *UsageQueryRowOUT) GetBillingMetric() string`

GetBillingMetric returns the BillingMetric field if non-nil, zero value otherwise.

### GetBillingMetricOk

`func (o *UsageQueryRowOUT) GetBillingMetricOk() (*string, bool)`

GetBillingMetricOk returns a tuple with the BillingMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMetric

`func (o *UsageQueryRowOUT) SetBillingMetric(v string)`

SetBillingMetric sets BillingMetric field to given value.

### HasBillingMetric

`func (o *UsageQueryRowOUT) HasBillingMetric() bool`

HasBillingMetric returns a boolean if a field has been set.

### SetBillingMetricNil

`func (o *UsageQueryRowOUT) SetBillingMetricNil(b bool)`

 SetBillingMetricNil sets the value for BillingMetric to be an explicit nil

### UnsetBillingMetric
`func (o *UsageQueryRowOUT) UnsetBillingMetric()`

UnsetBillingMetric ensures that no value is present for BillingMetric, not even an explicit nil
### GetBillingGroup

`func (o *UsageQueryRowOUT) GetBillingGroup() string`

GetBillingGroup returns the BillingGroup field if non-nil, zero value otherwise.

### GetBillingGroupOk

`func (o *UsageQueryRowOUT) GetBillingGroupOk() (*string, bool)`

GetBillingGroupOk returns a tuple with the BillingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroup

`func (o *UsageQueryRowOUT) SetBillingGroup(v string)`

SetBillingGroup sets BillingGroup field to given value.

### HasBillingGroup

`func (o *UsageQueryRowOUT) HasBillingGroup() bool`

HasBillingGroup returns a boolean if a field has been set.

### SetBillingGroupNil

`func (o *UsageQueryRowOUT) SetBillingGroupNil(b bool)`

 SetBillingGroupNil sets the value for BillingGroup to be an explicit nil

### UnsetBillingGroup
`func (o *UsageQueryRowOUT) UnsetBillingGroup()`

UnsetBillingGroup ensures that no value is present for BillingGroup, not even an explicit nil
### GetWorkspaceId

`func (o *UsageQueryRowOUT) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UsageQueryRowOUT) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UsageQueryRowOUT) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *UsageQueryRowOUT) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *UsageQueryRowOUT) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *UsageQueryRowOUT) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetOrganizationId

`func (o *UsageQueryRowOUT) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UsageQueryRowOUT) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UsageQueryRowOUT) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UsageQueryRowOUT) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *UsageQueryRowOUT) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *UsageQueryRowOUT) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetUserId

`func (o *UsageQueryRowOUT) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageQueryRowOUT) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageQueryRowOUT) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UsageQueryRowOUT) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UsageQueryRowOUT) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UsageQueryRowOUT) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetServiceAccountId

`func (o *UsageQueryRowOUT) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *UsageQueryRowOUT) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *UsageQueryRowOUT) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *UsageQueryRowOUT) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### SetServiceAccountIdNil

`func (o *UsageQueryRowOUT) SetServiceAccountIdNil(b bool)`

 SetServiceAccountIdNil sets the value for ServiceAccountId to be an explicit nil

### UnsetServiceAccountId
`func (o *UsageQueryRowOUT) UnsetServiceAccountId()`

UnsetServiceAccountId ensures that no value is present for ServiceAccountId, not even an explicit nil
### GetApiKeyId

`func (o *UsageQueryRowOUT) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *UsageQueryRowOUT) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *UsageQueryRowOUT) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *UsageQueryRowOUT) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *UsageQueryRowOUT) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *UsageQueryRowOUT) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetApiZone

`func (o *UsageQueryRowOUT) GetApiZone() string`

GetApiZone returns the ApiZone field if non-nil, zero value otherwise.

### GetApiZoneOk

`func (o *UsageQueryRowOUT) GetApiZoneOk() (*string, bool)`

GetApiZoneOk returns a tuple with the ApiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiZone

`func (o *UsageQueryRowOUT) SetApiZone(v string)`

SetApiZone sets ApiZone field to given value.

### HasApiZone

`func (o *UsageQueryRowOUT) HasApiZone() bool`

HasApiZone returns a boolean if a field has been set.

### SetApiZoneNil

`func (o *UsageQueryRowOUT) SetApiZoneNil(b bool)`

 SetApiZoneNil sets the value for ApiZone to be an explicit nil

### UnsetApiZone
`func (o *UsageQueryRowOUT) UnsetApiZone()`

UnsetApiZone ensures that no value is present for ApiZone, not even an explicit nil
### GetServiceTier

`func (o *UsageQueryRowOUT) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *UsageQueryRowOUT) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *UsageQueryRowOUT) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *UsageQueryRowOUT) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *UsageQueryRowOUT) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *UsageQueryRowOUT) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetTimestamp

`func (o *UsageQueryRowOUT) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UsageQueryRowOUT) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UsageQueryRowOUT) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UsageQueryRowOUT) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UsageQueryRowOUT) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UsageQueryRowOUT) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTimeBucket

`func (o *UsageQueryRowOUT) GetTimeBucket() time.Time`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *UsageQueryRowOUT) GetTimeBucketOk() (*time.Time, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *UsageQueryRowOUT) SetTimeBucket(v time.Time)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *UsageQueryRowOUT) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### SetTimeBucketNil

`func (o *UsageQueryRowOUT) SetTimeBucketNil(b bool)`

 SetTimeBucketNil sets the value for TimeBucket to be an explicit nil

### UnsetTimeBucket
`func (o *UsageQueryRowOUT) UnsetTimeBucket()`

UnsetTimeBucket ensures that no value is present for TimeBucket, not even an explicit nil
### GetValue

`func (o *UsageQueryRowOUT) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsageQueryRowOUT) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsageQueryRowOUT) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *UsageQueryRowOUT) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UsageQueryRowOUT) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UsageQueryRowOUT) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCount

`func (o *UsageQueryRowOUT) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UsageQueryRowOUT) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UsageQueryRowOUT) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UsageQueryRowOUT) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *UsageQueryRowOUT) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *UsageQueryRowOUT) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetNbApiCalls

`func (o *UsageQueryRowOUT) GetNbApiCalls() int32`

GetNbApiCalls returns the NbApiCalls field if non-nil, zero value otherwise.

### GetNbApiCallsOk

`func (o *UsageQueryRowOUT) GetNbApiCallsOk() (*int32, bool)`

GetNbApiCallsOk returns a tuple with the NbApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbApiCalls

`func (o *UsageQueryRowOUT) SetNbApiCalls(v int32)`

SetNbApiCalls sets NbApiCalls field to given value.

### HasNbApiCalls

`func (o *UsageQueryRowOUT) HasNbApiCalls() bool`

HasNbApiCalls returns a boolean if a field has been set.

### SetNbApiCallsNil

`func (o *UsageQueryRowOUT) SetNbApiCallsNil(b bool)`

 SetNbApiCallsNil sets the value for NbApiCalls to be an explicit nil

### UnsetNbApiCalls
`func (o *UsageQueryRowOUT) UnsetNbApiCalls()`

UnsetNbApiCalls ensures that no value is present for NbApiCalls, not even an explicit nil
### GetBillingDisplayName

`func (o *UsageQueryRowOUT) GetBillingDisplayName() string`

GetBillingDisplayName returns the BillingDisplayName field if non-nil, zero value otherwise.

### GetBillingDisplayNameOk

`func (o *UsageQueryRowOUT) GetBillingDisplayNameOk() (*string, bool)`

GetBillingDisplayNameOk returns a tuple with the BillingDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDisplayName

`func (o *UsageQueryRowOUT) SetBillingDisplayName(v string)`

SetBillingDisplayName sets BillingDisplayName field to given value.

### HasBillingDisplayName

`func (o *UsageQueryRowOUT) HasBillingDisplayName() bool`

HasBillingDisplayName returns a boolean if a field has been set.

### SetBillingDisplayNameNil

`func (o *UsageQueryRowOUT) SetBillingDisplayNameNil(b bool)`

 SetBillingDisplayNameNil sets the value for BillingDisplayName to be an explicit nil

### UnsetBillingDisplayName
`func (o *UsageQueryRowOUT) UnsetBillingDisplayName()`

UnsetBillingDisplayName ensures that no value is present for BillingDisplayName, not even an explicit nil
### GetCustomerTier

`func (o *UsageQueryRowOUT) GetCustomerTier() string`

GetCustomerTier returns the CustomerTier field if non-nil, zero value otherwise.

### GetCustomerTierOk

`func (o *UsageQueryRowOUT) GetCustomerTierOk() (*string, bool)`

GetCustomerTierOk returns a tuple with the CustomerTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTier

`func (o *UsageQueryRowOUT) SetCustomerTier(v string)`

SetCustomerTier sets CustomerTier field to given value.

### HasCustomerTier

`func (o *UsageQueryRowOUT) HasCustomerTier() bool`

HasCustomerTier returns a boolean if a field has been set.

### SetCustomerTierNil

`func (o *UsageQueryRowOUT) SetCustomerTierNil(b bool)`

 SetCustomerTierNil sets the value for CustomerTier to be an explicit nil

### UnsetCustomerTier
`func (o *UsageQueryRowOUT) UnsetCustomerTier()`

UnsetCustomerTier ensures that no value is present for CustomerTier, not even an explicit nil
### GetUntracked

`func (o *UsageQueryRowOUT) GetUntracked() bool`

GetUntracked returns the Untracked field if non-nil, zero value otherwise.

### GetUntrackedOk

`func (o *UsageQueryRowOUT) GetUntrackedOk() (*bool, bool)`

GetUntrackedOk returns a tuple with the Untracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntracked

`func (o *UsageQueryRowOUT) SetUntracked(v bool)`

SetUntracked sets Untracked field to given value.

### HasUntracked

`func (o *UsageQueryRowOUT) HasUntracked() bool`

HasUntracked returns a boolean if a field has been set.

### SetUntrackedNil

`func (o *UsageQueryRowOUT) SetUntrackedNil(b bool)`

 SetUntrackedNil sets the value for Untracked to be an explicit nil

### UnsetUntracked
`func (o *UsageQueryRowOUT) UnsetUntracked()`

UnsetUntracked ensures that no value is present for Untracked, not even an explicit nil
### GetCallSource

`func (o *UsageQueryRowOUT) GetCallSource() string`

GetCallSource returns the CallSource field if non-nil, zero value otherwise.

### GetCallSourceOk

`func (o *UsageQueryRowOUT) GetCallSourceOk() (*string, bool)`

GetCallSourceOk returns a tuple with the CallSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSource

`func (o *UsageQueryRowOUT) SetCallSource(v string)`

SetCallSource sets CallSource field to given value.

### HasCallSource

`func (o *UsageQueryRowOUT) HasCallSource() bool`

HasCallSource returns a boolean if a field has been set.

### SetCallSourceNil

`func (o *UsageQueryRowOUT) SetCallSourceNil(b bool)`

 SetCallSourceNil sets the value for CallSource to be an explicit nil

### UnsetCallSource
`func (o *UsageQueryRowOUT) UnsetCallSource()`

UnsetCallSource ensures that no value is present for CallSource, not even an explicit nil
### GetCallType

`func (o *UsageQueryRowOUT) GetCallType() string`

GetCallType returns the CallType field if non-nil, zero value otherwise.

### GetCallTypeOk

`func (o *UsageQueryRowOUT) GetCallTypeOk() (*string, bool)`

GetCallTypeOk returns a tuple with the CallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallType

`func (o *UsageQueryRowOUT) SetCallType(v string)`

SetCallType sets CallType field to given value.

### HasCallType

`func (o *UsageQueryRowOUT) HasCallType() bool`

HasCallType returns a boolean if a field has been set.

### SetCallTypeNil

`func (o *UsageQueryRowOUT) SetCallTypeNil(b bool)`

 SetCallTypeNil sets the value for CallType to be an explicit nil

### UnsetCallType
`func (o *UsageQueryRowOUT) UnsetCallType()`

UnsetCallType ensures that no value is present for CallType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



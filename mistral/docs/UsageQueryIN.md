# UsageQueryIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageType** | Pointer to [**[]UsageType**](UsageType.md) | Usage event types to include. | [optional] 
**BillingMetric** | Pointer to **[]string** | Billing metric identifiers to include. | [optional] 
**BillingDisplayName** | Pointer to **[]string** | Customer-facing billing metric names to include. | [optional] 
**BillingGroup** | Pointer to **[]string** | Billing metric components to include. | [optional] 
**WorkspaceIds** | Pointer to **[]string** | Workspace IDs to include. Use null to include usage without a workspace. | [optional] 
**OrganizationId** | Pointer to **[]string** | Organization IDs to include within the authenticated customer. | [optional] 
**UserId** | Pointer to **[]string** | User IDs to include. | [optional] 
**ServiceAccountId** | Pointer to **[]string** | Service account IDs to include. | [optional] 
**ApiKeyId** | Pointer to **[]string** | API key IDs to include. | [optional] 
**ApiZone** | Pointer to **[]string** | Inference regions to include. | [optional] 
**ServiceTier** | Pointer to **[]string** | Service tiers to include. | [optional] 
**Start** | **time.Time** | Inclusive start of the usage interval. | 
**End** | **time.Time** | Exclusive end of the usage interval. | 
**Filters** | Pointer to [**[]AdminUsageQueryFilter**](AdminUsageQueryFilter.md) | Additional inclusion or exclusion filters applied alongside the dimension fields. | [optional] 
**GroupBy** | Pointer to **[]string** | Dimensions used to group the returned usage rows. | [optional] 
**TimeBucket** | Pointer to **NullableString** | Time interval used to bucket grouped rows. | [optional] 
**Untracked** | Pointer to **NullableBool** | Deprecated: this parameter will be removed in a future version. | [optional] 
**Limit** | Pointer to **NullableInt32** | Maximum number of grouped rows to return. | [optional] 
**Offset** | Pointer to **NullableInt32** | Number of grouped rows to skip. | [optional] 
**IncludeAggregates** | Pointer to **bool** | Return aggregates computed over the complete filtered result set. | [optional] [default to false]
**IncludeNbApiCalls** | Pointer to **bool** | Include the number of distinct API calls in grouped rows and aggregates. | [optional] [default to false]

## Methods

### NewUsageQueryIN

`func NewUsageQueryIN(start time.Time, end time.Time, ) *UsageQueryIN`

NewUsageQueryIN instantiates a new UsageQueryIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageQueryINWithDefaults

`func NewUsageQueryINWithDefaults() *UsageQueryIN`

NewUsageQueryINWithDefaults instantiates a new UsageQueryIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageType

`func (o *UsageQueryIN) GetUsageType() []UsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *UsageQueryIN) GetUsageTypeOk() (*[]UsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *UsageQueryIN) SetUsageType(v []UsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *UsageQueryIN) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *UsageQueryIN) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *UsageQueryIN) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetBillingMetric

`func (o *UsageQueryIN) GetBillingMetric() []string`

GetBillingMetric returns the BillingMetric field if non-nil, zero value otherwise.

### GetBillingMetricOk

`func (o *UsageQueryIN) GetBillingMetricOk() (*[]string, bool)`

GetBillingMetricOk returns a tuple with the BillingMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMetric

`func (o *UsageQueryIN) SetBillingMetric(v []string)`

SetBillingMetric sets BillingMetric field to given value.

### HasBillingMetric

`func (o *UsageQueryIN) HasBillingMetric() bool`

HasBillingMetric returns a boolean if a field has been set.

### SetBillingMetricNil

`func (o *UsageQueryIN) SetBillingMetricNil(b bool)`

 SetBillingMetricNil sets the value for BillingMetric to be an explicit nil

### UnsetBillingMetric
`func (o *UsageQueryIN) UnsetBillingMetric()`

UnsetBillingMetric ensures that no value is present for BillingMetric, not even an explicit nil
### GetBillingDisplayName

`func (o *UsageQueryIN) GetBillingDisplayName() []string`

GetBillingDisplayName returns the BillingDisplayName field if non-nil, zero value otherwise.

### GetBillingDisplayNameOk

`func (o *UsageQueryIN) GetBillingDisplayNameOk() (*[]string, bool)`

GetBillingDisplayNameOk returns a tuple with the BillingDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDisplayName

`func (o *UsageQueryIN) SetBillingDisplayName(v []string)`

SetBillingDisplayName sets BillingDisplayName field to given value.

### HasBillingDisplayName

`func (o *UsageQueryIN) HasBillingDisplayName() bool`

HasBillingDisplayName returns a boolean if a field has been set.

### SetBillingDisplayNameNil

`func (o *UsageQueryIN) SetBillingDisplayNameNil(b bool)`

 SetBillingDisplayNameNil sets the value for BillingDisplayName to be an explicit nil

### UnsetBillingDisplayName
`func (o *UsageQueryIN) UnsetBillingDisplayName()`

UnsetBillingDisplayName ensures that no value is present for BillingDisplayName, not even an explicit nil
### GetBillingGroup

`func (o *UsageQueryIN) GetBillingGroup() []string`

GetBillingGroup returns the BillingGroup field if non-nil, zero value otherwise.

### GetBillingGroupOk

`func (o *UsageQueryIN) GetBillingGroupOk() (*[]string, bool)`

GetBillingGroupOk returns a tuple with the BillingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroup

`func (o *UsageQueryIN) SetBillingGroup(v []string)`

SetBillingGroup sets BillingGroup field to given value.

### HasBillingGroup

`func (o *UsageQueryIN) HasBillingGroup() bool`

HasBillingGroup returns a boolean if a field has been set.

### SetBillingGroupNil

`func (o *UsageQueryIN) SetBillingGroupNil(b bool)`

 SetBillingGroupNil sets the value for BillingGroup to be an explicit nil

### UnsetBillingGroup
`func (o *UsageQueryIN) UnsetBillingGroup()`

UnsetBillingGroup ensures that no value is present for BillingGroup, not even an explicit nil
### GetWorkspaceIds

`func (o *UsageQueryIN) GetWorkspaceIds() []*string`

GetWorkspaceIds returns the WorkspaceIds field if non-nil, zero value otherwise.

### GetWorkspaceIdsOk

`func (o *UsageQueryIN) GetWorkspaceIdsOk() (*[]*string, bool)`

GetWorkspaceIdsOk returns a tuple with the WorkspaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceIds

`func (o *UsageQueryIN) SetWorkspaceIds(v []*string)`

SetWorkspaceIds sets WorkspaceIds field to given value.

### HasWorkspaceIds

`func (o *UsageQueryIN) HasWorkspaceIds() bool`

HasWorkspaceIds returns a boolean if a field has been set.

### SetWorkspaceIdsNil

`func (o *UsageQueryIN) SetWorkspaceIdsNil(b bool)`

 SetWorkspaceIdsNil sets the value for WorkspaceIds to be an explicit nil

### UnsetWorkspaceIds
`func (o *UsageQueryIN) UnsetWorkspaceIds()`

UnsetWorkspaceIds ensures that no value is present for WorkspaceIds, not even an explicit nil
### GetOrganizationId

`func (o *UsageQueryIN) GetOrganizationId() []string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UsageQueryIN) GetOrganizationIdOk() (*[]string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UsageQueryIN) SetOrganizationId(v []string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UsageQueryIN) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *UsageQueryIN) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *UsageQueryIN) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetUserId

`func (o *UsageQueryIN) GetUserId() []string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageQueryIN) GetUserIdOk() (*[]string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageQueryIN) SetUserId(v []string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UsageQueryIN) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UsageQueryIN) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UsageQueryIN) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetServiceAccountId

`func (o *UsageQueryIN) GetServiceAccountId() []string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *UsageQueryIN) GetServiceAccountIdOk() (*[]string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *UsageQueryIN) SetServiceAccountId(v []string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *UsageQueryIN) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### SetServiceAccountIdNil

`func (o *UsageQueryIN) SetServiceAccountIdNil(b bool)`

 SetServiceAccountIdNil sets the value for ServiceAccountId to be an explicit nil

### UnsetServiceAccountId
`func (o *UsageQueryIN) UnsetServiceAccountId()`

UnsetServiceAccountId ensures that no value is present for ServiceAccountId, not even an explicit nil
### GetApiKeyId

`func (o *UsageQueryIN) GetApiKeyId() []string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *UsageQueryIN) GetApiKeyIdOk() (*[]string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *UsageQueryIN) SetApiKeyId(v []string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *UsageQueryIN) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *UsageQueryIN) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *UsageQueryIN) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetApiZone

`func (o *UsageQueryIN) GetApiZone() []string`

GetApiZone returns the ApiZone field if non-nil, zero value otherwise.

### GetApiZoneOk

`func (o *UsageQueryIN) GetApiZoneOk() (*[]string, bool)`

GetApiZoneOk returns a tuple with the ApiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiZone

`func (o *UsageQueryIN) SetApiZone(v []string)`

SetApiZone sets ApiZone field to given value.

### HasApiZone

`func (o *UsageQueryIN) HasApiZone() bool`

HasApiZone returns a boolean if a field has been set.

### SetApiZoneNil

`func (o *UsageQueryIN) SetApiZoneNil(b bool)`

 SetApiZoneNil sets the value for ApiZone to be an explicit nil

### UnsetApiZone
`func (o *UsageQueryIN) UnsetApiZone()`

UnsetApiZone ensures that no value is present for ApiZone, not even an explicit nil
### GetServiceTier

`func (o *UsageQueryIN) GetServiceTier() []string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *UsageQueryIN) GetServiceTierOk() (*[]string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *UsageQueryIN) SetServiceTier(v []string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *UsageQueryIN) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *UsageQueryIN) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *UsageQueryIN) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetStart

`func (o *UsageQueryIN) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UsageQueryIN) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UsageQueryIN) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *UsageQueryIN) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UsageQueryIN) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UsageQueryIN) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetFilters

`func (o *UsageQueryIN) GetFilters() []AdminUsageQueryFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UsageQueryIN) GetFiltersOk() (*[]AdminUsageQueryFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UsageQueryIN) SetFilters(v []AdminUsageQueryFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UsageQueryIN) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *UsageQueryIN) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *UsageQueryIN) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetGroupBy

`func (o *UsageQueryIN) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *UsageQueryIN) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *UsageQueryIN) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *UsageQueryIN) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupByNil

`func (o *UsageQueryIN) SetGroupByNil(b bool)`

 SetGroupByNil sets the value for GroupBy to be an explicit nil

### UnsetGroupBy
`func (o *UsageQueryIN) UnsetGroupBy()`

UnsetGroupBy ensures that no value is present for GroupBy, not even an explicit nil
### GetTimeBucket

`func (o *UsageQueryIN) GetTimeBucket() string`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *UsageQueryIN) GetTimeBucketOk() (*string, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *UsageQueryIN) SetTimeBucket(v string)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *UsageQueryIN) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### SetTimeBucketNil

`func (o *UsageQueryIN) SetTimeBucketNil(b bool)`

 SetTimeBucketNil sets the value for TimeBucket to be an explicit nil

### UnsetTimeBucket
`func (o *UsageQueryIN) UnsetTimeBucket()`

UnsetTimeBucket ensures that no value is present for TimeBucket, not even an explicit nil
### GetUntracked

`func (o *UsageQueryIN) GetUntracked() bool`

GetUntracked returns the Untracked field if non-nil, zero value otherwise.

### GetUntrackedOk

`func (o *UsageQueryIN) GetUntrackedOk() (*bool, bool)`

GetUntrackedOk returns a tuple with the Untracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntracked

`func (o *UsageQueryIN) SetUntracked(v bool)`

SetUntracked sets Untracked field to given value.

### HasUntracked

`func (o *UsageQueryIN) HasUntracked() bool`

HasUntracked returns a boolean if a field has been set.

### SetUntrackedNil

`func (o *UsageQueryIN) SetUntrackedNil(b bool)`

 SetUntrackedNil sets the value for Untracked to be an explicit nil

### UnsetUntracked
`func (o *UsageQueryIN) UnsetUntracked()`

UnsetUntracked ensures that no value is present for Untracked, not even an explicit nil
### GetLimit

`func (o *UsageQueryIN) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UsageQueryIN) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UsageQueryIN) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UsageQueryIN) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *UsageQueryIN) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *UsageQueryIN) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *UsageQueryIN) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UsageQueryIN) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UsageQueryIN) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UsageQueryIN) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *UsageQueryIN) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *UsageQueryIN) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetIncludeAggregates

`func (o *UsageQueryIN) GetIncludeAggregates() bool`

GetIncludeAggregates returns the IncludeAggregates field if non-nil, zero value otherwise.

### GetIncludeAggregatesOk

`func (o *UsageQueryIN) GetIncludeAggregatesOk() (*bool, bool)`

GetIncludeAggregatesOk returns a tuple with the IncludeAggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAggregates

`func (o *UsageQueryIN) SetIncludeAggregates(v bool)`

SetIncludeAggregates sets IncludeAggregates field to given value.

### HasIncludeAggregates

`func (o *UsageQueryIN) HasIncludeAggregates() bool`

HasIncludeAggregates returns a boolean if a field has been set.

### GetIncludeNbApiCalls

`func (o *UsageQueryIN) GetIncludeNbApiCalls() bool`

GetIncludeNbApiCalls returns the IncludeNbApiCalls field if non-nil, zero value otherwise.

### GetIncludeNbApiCallsOk

`func (o *UsageQueryIN) GetIncludeNbApiCallsOk() (*bool, bool)`

GetIncludeNbApiCallsOk returns a tuple with the IncludeNbApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNbApiCalls

`func (o *UsageQueryIN) SetIncludeNbApiCalls(v bool)`

SetIncludeNbApiCalls sets IncludeNbApiCalls field to given value.

### HasIncludeNbApiCalls

`func (o *UsageQueryIN) HasIncludeNbApiCalls() bool`

HasIncludeNbApiCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Data1Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | Pointer to **int32** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] [default to "mistralai"]
**Capabilities** | [**ModelCapabilities**](ModelCapabilities.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaxContextLength** | Pointer to **int32** |  | [optional] [default to 32768]
**Aliases** | Pointer to **[]string** |  | [optional] [default to {}]
**Deprecation** | Pointer to **time.Time** |  | [optional] 
**DeprecationReplacementModel** | Pointer to **string** |  | [optional] 
**DefaultModelTemperature** | Pointer to **float32** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] [default to false]
**BillingModelName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "base"]
**Job** | **string** |  | 
**Root** | **string** |  | 
**Archived** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewData1Inner

`func NewData1Inner(id string, capabilities ModelCapabilities, job string, root string, ) *Data1Inner`

NewData1Inner instantiates a new Data1Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewData1InnerWithDefaults

`func NewData1InnerWithDefaults() *Data1Inner`

NewData1InnerWithDefaults instantiates a new Data1Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Data1Inner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Data1Inner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Data1Inner) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Data1Inner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Data1Inner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Data1Inner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Data1Inner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *Data1Inner) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Data1Inner) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Data1Inner) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Data1Inner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOwnedBy

`func (o *Data1Inner) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *Data1Inner) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *Data1Inner) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *Data1Inner) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetCapabilities

`func (o *Data1Inner) GetCapabilities() ModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Data1Inner) GetCapabilitiesOk() (*ModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Data1Inner) SetCapabilities(v ModelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetName

`func (o *Data1Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Data1Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Data1Inner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Data1Inner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Data1Inner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Data1Inner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Data1Inner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Data1Inner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxContextLength

`func (o *Data1Inner) GetMaxContextLength() int32`

GetMaxContextLength returns the MaxContextLength field if non-nil, zero value otherwise.

### GetMaxContextLengthOk

`func (o *Data1Inner) GetMaxContextLengthOk() (*int32, bool)`

GetMaxContextLengthOk returns a tuple with the MaxContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextLength

`func (o *Data1Inner) SetMaxContextLength(v int32)`

SetMaxContextLength sets MaxContextLength field to given value.

### HasMaxContextLength

`func (o *Data1Inner) HasMaxContextLength() bool`

HasMaxContextLength returns a boolean if a field has been set.

### GetAliases

`func (o *Data1Inner) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Data1Inner) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Data1Inner) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Data1Inner) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDeprecation

`func (o *Data1Inner) GetDeprecation() time.Time`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *Data1Inner) GetDeprecationOk() (*time.Time, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *Data1Inner) SetDeprecation(v time.Time)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *Data1Inner) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### GetDeprecationReplacementModel

`func (o *Data1Inner) GetDeprecationReplacementModel() string`

GetDeprecationReplacementModel returns the DeprecationReplacementModel field if non-nil, zero value otherwise.

### GetDeprecationReplacementModelOk

`func (o *Data1Inner) GetDeprecationReplacementModelOk() (*string, bool)`

GetDeprecationReplacementModelOk returns a tuple with the DeprecationReplacementModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationReplacementModel

`func (o *Data1Inner) SetDeprecationReplacementModel(v string)`

SetDeprecationReplacementModel sets DeprecationReplacementModel field to given value.

### HasDeprecationReplacementModel

`func (o *Data1Inner) HasDeprecationReplacementModel() bool`

HasDeprecationReplacementModel returns a boolean if a field has been set.

### GetDefaultModelTemperature

`func (o *Data1Inner) GetDefaultModelTemperature() float32`

GetDefaultModelTemperature returns the DefaultModelTemperature field if non-nil, zero value otherwise.

### GetDefaultModelTemperatureOk

`func (o *Data1Inner) GetDefaultModelTemperatureOk() (*float32, bool)`

GetDefaultModelTemperatureOk returns a tuple with the DefaultModelTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModelTemperature

`func (o *Data1Inner) SetDefaultModelTemperature(v float32)`

SetDefaultModelTemperature sets DefaultModelTemperature field to given value.

### HasDefaultModelTemperature

`func (o *Data1Inner) HasDefaultModelTemperature() bool`

HasDefaultModelTemperature returns a boolean if a field has been set.

### GetInternal

`func (o *Data1Inner) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Data1Inner) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Data1Inner) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Data1Inner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetBillingModelName

`func (o *Data1Inner) GetBillingModelName() string`

GetBillingModelName returns the BillingModelName field if non-nil, zero value otherwise.

### GetBillingModelNameOk

`func (o *Data1Inner) GetBillingModelNameOk() (*string, bool)`

GetBillingModelNameOk returns a tuple with the BillingModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModelName

`func (o *Data1Inner) SetBillingModelName(v string)`

SetBillingModelName sets BillingModelName field to given value.

### HasBillingModelName

`func (o *Data1Inner) HasBillingModelName() bool`

HasBillingModelName returns a boolean if a field has been set.

### GetType

`func (o *Data1Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Data1Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Data1Inner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Data1Inner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJob

`func (o *Data1Inner) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Data1Inner) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Data1Inner) SetJob(v string)`

SetJob sets Job field to given value.


### GetRoot

`func (o *Data1Inner) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Data1Inner) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *Data1Inner) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetArchived

`func (o *Data1Inner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Data1Inner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Data1Inner) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Data1Inner) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ScimSyncUsersSummaryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provision** | [**[]ScimSyncUserActionOut**](ScimSyncUserActionOut.md) |  | 
**Deprovision** | [**[]ScimSyncUserActionOut**](ScimSyncUserActionOut.md) |  | 

## Methods

### NewScimSyncUsersSummaryOut

`func NewScimSyncUsersSummaryOut(provision []ScimSyncUserActionOut, deprovision []ScimSyncUserActionOut, ) *ScimSyncUsersSummaryOut`

NewScimSyncUsersSummaryOut instantiates a new ScimSyncUsersSummaryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncUsersSummaryOutWithDefaults

`func NewScimSyncUsersSummaryOutWithDefaults() *ScimSyncUsersSummaryOut`

NewScimSyncUsersSummaryOutWithDefaults instantiates a new ScimSyncUsersSummaryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvision

`func (o *ScimSyncUsersSummaryOut) GetProvision() []ScimSyncUserActionOut`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *ScimSyncUsersSummaryOut) GetProvisionOk() (*[]ScimSyncUserActionOut, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *ScimSyncUsersSummaryOut) SetProvision(v []ScimSyncUserActionOut)`

SetProvision sets Provision field to given value.


### GetDeprovision

`func (o *ScimSyncUsersSummaryOut) GetDeprovision() []ScimSyncUserActionOut`

GetDeprovision returns the Deprovision field if non-nil, zero value otherwise.

### GetDeprovisionOk

`func (o *ScimSyncUsersSummaryOut) GetDeprovisionOk() (*[]ScimSyncUserActionOut, bool)`

GetDeprovisionOk returns a tuple with the Deprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprovision

`func (o *ScimSyncUsersSummaryOut) SetDeprovision(v []ScimSyncUserActionOut)`

SetDeprovision sets Deprovision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



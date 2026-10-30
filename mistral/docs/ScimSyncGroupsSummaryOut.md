# ScimSyncGroupsSummaryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | [**[]ScimSyncGroupActionOut**](ScimSyncGroupActionOut.md) |  | 
**Update** | [**[]ScimSyncGroupUpdateActionOut**](ScimSyncGroupUpdateActionOut.md) |  | 
**Delete** | [**[]ScimSyncGroupActionOut**](ScimSyncGroupActionOut.md) |  | 

## Methods

### NewScimSyncGroupsSummaryOut

`func NewScimSyncGroupsSummaryOut(create []ScimSyncGroupActionOut, update []ScimSyncGroupUpdateActionOut, delete []ScimSyncGroupActionOut, ) *ScimSyncGroupsSummaryOut`

NewScimSyncGroupsSummaryOut instantiates a new ScimSyncGroupsSummaryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncGroupsSummaryOutWithDefaults

`func NewScimSyncGroupsSummaryOutWithDefaults() *ScimSyncGroupsSummaryOut`

NewScimSyncGroupsSummaryOutWithDefaults instantiates a new ScimSyncGroupsSummaryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *ScimSyncGroupsSummaryOut) GetCreate() []ScimSyncGroupActionOut`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ScimSyncGroupsSummaryOut) GetCreateOk() (*[]ScimSyncGroupActionOut, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ScimSyncGroupsSummaryOut) SetCreate(v []ScimSyncGroupActionOut)`

SetCreate sets Create field to given value.


### GetUpdate

`func (o *ScimSyncGroupsSummaryOut) GetUpdate() []ScimSyncGroupUpdateActionOut`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ScimSyncGroupsSummaryOut) GetUpdateOk() (*[]ScimSyncGroupUpdateActionOut, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ScimSyncGroupsSummaryOut) SetUpdate(v []ScimSyncGroupUpdateActionOut)`

SetUpdate sets Update field to given value.


### GetDelete

`func (o *ScimSyncGroupsSummaryOut) GetDelete() []ScimSyncGroupActionOut`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ScimSyncGroupsSummaryOut) GetDeleteOk() (*[]ScimSyncGroupActionOut, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ScimSyncGroupsSummaryOut) SetDelete(v []ScimSyncGroupActionOut)`

SetDelete sets Delete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



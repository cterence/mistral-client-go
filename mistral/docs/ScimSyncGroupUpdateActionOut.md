# ScimSyncGroupUpdateActionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**ScimSyncGroupRefOut**](ScimSyncGroupRefOut.md) |  | 
**PreviousGroupName** | **string** |  | 

## Methods

### NewScimSyncGroupUpdateActionOut

`func NewScimSyncGroupUpdateActionOut(group ScimSyncGroupRefOut, previousGroupName string, ) *ScimSyncGroupUpdateActionOut`

NewScimSyncGroupUpdateActionOut instantiates a new ScimSyncGroupUpdateActionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncGroupUpdateActionOutWithDefaults

`func NewScimSyncGroupUpdateActionOutWithDefaults() *ScimSyncGroupUpdateActionOut`

NewScimSyncGroupUpdateActionOutWithDefaults instantiates a new ScimSyncGroupUpdateActionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ScimSyncGroupUpdateActionOut) GetGroup() ScimSyncGroupRefOut`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ScimSyncGroupUpdateActionOut) GetGroupOk() (*ScimSyncGroupRefOut, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ScimSyncGroupUpdateActionOut) SetGroup(v ScimSyncGroupRefOut)`

SetGroup sets Group field to given value.


### GetPreviousGroupName

`func (o *ScimSyncGroupUpdateActionOut) GetPreviousGroupName() string`

GetPreviousGroupName returns the PreviousGroupName field if non-nil, zero value otherwise.

### GetPreviousGroupNameOk

`func (o *ScimSyncGroupUpdateActionOut) GetPreviousGroupNameOk() (*string, bool)`

GetPreviousGroupNameOk returns a tuple with the PreviousGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousGroupName

`func (o *ScimSyncGroupUpdateActionOut) SetPreviousGroupName(v string)`

SetPreviousGroupName sets PreviousGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



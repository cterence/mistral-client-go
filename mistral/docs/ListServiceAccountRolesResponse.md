# ListServiceAccountRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ServiceAccountRole**](ServiceAccountRole.md) |  | 

## Methods

### NewListServiceAccountRolesResponse

`func NewListServiceAccountRolesResponse(items []ServiceAccountRole, ) *ListServiceAccountRolesResponse`

NewListServiceAccountRolesResponse instantiates a new ListServiceAccountRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceAccountRolesResponseWithDefaults

`func NewListServiceAccountRolesResponseWithDefaults() *ListServiceAccountRolesResponse`

NewListServiceAccountRolesResponseWithDefaults instantiates a new ListServiceAccountRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListServiceAccountRolesResponse) GetItems() []ServiceAccountRole`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListServiceAccountRolesResponse) GetItemsOk() (*[]ServiceAccountRole, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListServiceAccountRolesResponse) SetItems(v []ServiceAccountRole)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



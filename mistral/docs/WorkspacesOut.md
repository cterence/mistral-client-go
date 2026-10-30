# WorkspacesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of Workspaces that match the request. | 
**Items** | [**[]WorkspaceOUT**](WorkspaceOUT.md) | Workspaces on this page. | 
**Object** | Pointer to [**ApiObjectType**](ApiObjectType.md) | Type of paginated API object. | [optional] [default to APIOBJECTTYPE_LIST]
**Page** | **int32** | Page number returned. | 
**PageSize** | **int32** | Maximum number of results per page. | 

## Methods

### NewWorkspacesOut

`func NewWorkspacesOut(total int32, items []WorkspaceOUT, page int32, pageSize int32, ) *WorkspacesOut`

NewWorkspacesOut instantiates a new WorkspacesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacesOutWithDefaults

`func NewWorkspacesOutWithDefaults() *WorkspacesOut`

NewWorkspacesOutWithDefaults instantiates a new WorkspacesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *WorkspacesOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkspacesOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkspacesOut) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *WorkspacesOut) GetItems() []WorkspaceOUT`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkspacesOut) GetItemsOk() (*[]WorkspaceOUT, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkspacesOut) SetItems(v []WorkspaceOUT)`

SetItems sets Items field to given value.


### GetObject

`func (o *WorkspacesOut) GetObject() ApiObjectType`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WorkspacesOut) GetObjectOk() (*ApiObjectType, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WorkspacesOut) SetObject(v ApiObjectType)`

SetObject sets Object field to given value.

### HasObject

`func (o *WorkspacesOut) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPage

`func (o *WorkspacesOut) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WorkspacesOut) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WorkspacesOut) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *WorkspacesOut) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *WorkspacesOut) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *WorkspacesOut) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



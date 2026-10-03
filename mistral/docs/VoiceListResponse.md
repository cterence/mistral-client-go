# VoiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]VoiceResponse**](VoiceResponse.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalPages** | **int32** |  | 

## Methods

### NewVoiceListResponse

`func NewVoiceListResponse(items []VoiceResponse, total int32, page int32, pageSize int32, totalPages int32, ) *VoiceListResponse`

NewVoiceListResponse instantiates a new VoiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceListResponseWithDefaults

`func NewVoiceListResponseWithDefaults() *VoiceListResponse`

NewVoiceListResponseWithDefaults instantiates a new VoiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *VoiceListResponse) GetItems() []VoiceResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VoiceListResponse) GetItemsOk() (*[]VoiceResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VoiceListResponse) SetItems(v []VoiceResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *VoiceListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VoiceListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VoiceListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *VoiceListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VoiceListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VoiceListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *VoiceListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *VoiceListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *VoiceListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalPages

`func (o *VoiceListResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *VoiceListResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *VoiceListResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



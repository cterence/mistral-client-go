# PaginatedResultCampaignPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]CampaignPreview**](CampaignPreview.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultCampaignPreview

`func NewPaginatedResultCampaignPreview(count int32, ) *PaginatedResultCampaignPreview`

NewPaginatedResultCampaignPreview instantiates a new PaginatedResultCampaignPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultCampaignPreviewWithDefaults

`func NewPaginatedResultCampaignPreviewWithDefaults() *PaginatedResultCampaignPreview`

NewPaginatedResultCampaignPreviewWithDefaults instantiates a new PaginatedResultCampaignPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultCampaignPreview) GetResults() []CampaignPreview`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultCampaignPreview) GetResultsOk() (*[]CampaignPreview, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultCampaignPreview) SetResults(v []CampaignPreview)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultCampaignPreview) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultCampaignPreview) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultCampaignPreview) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultCampaignPreview) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultCampaignPreview) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultCampaignPreview) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultCampaignPreview) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultCampaignPreview) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultCampaignPreview) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultCampaignPreview) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultCampaignPreview) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultCampaignPreview) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultCampaignPreview) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultCampaignPreview) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultCampaignPreview) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultCampaignPreview) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



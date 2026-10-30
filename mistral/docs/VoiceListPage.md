# VoiceListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VoiceResponse**](VoiceResponse.md) |  | 
**NextPageToken** | Pointer to **NullableString** | Cursor for the next page, null on the last page | [optional] 

## Methods

### NewVoiceListPage

`func NewVoiceListPage(data []VoiceResponse, ) *VoiceListPage`

NewVoiceListPage instantiates a new VoiceListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceListPageWithDefaults

`func NewVoiceListPageWithDefaults() *VoiceListPage`

NewVoiceListPageWithDefaults instantiates a new VoiceListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VoiceListPage) GetData() []VoiceResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VoiceListPage) GetDataOk() (*[]VoiceResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VoiceListPage) SetData(v []VoiceResponse)`

SetData sets Data field to given value.


### GetNextPageToken

`func (o *VoiceListPage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *VoiceListPage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *VoiceListPage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *VoiceListPage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *VoiceListPage) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *VoiceListPage) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FileSchema**](FileSchema.md) |  | 
**Object** | **string** |  | 
**Total** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewListFilesResponse

`func NewListFilesResponse(data []FileSchema, object string, ) *ListFilesResponse`

NewListFilesResponse instantiates a new ListFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFilesResponseWithDefaults

`func NewListFilesResponseWithDefaults() *ListFilesResponse`

NewListFilesResponseWithDefaults instantiates a new ListFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFilesResponse) GetData() []FileSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFilesResponse) GetDataOk() (*[]FileSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFilesResponse) SetData(v []FileSchema)`

SetData sets Data field to given value.


### GetObject

`func (o *ListFilesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ListFilesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ListFilesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetTotal

`func (o *ListFilesResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListFilesResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListFilesResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListFilesResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListFilesResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListFilesResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



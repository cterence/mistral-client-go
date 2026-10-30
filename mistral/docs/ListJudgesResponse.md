# ListJudgesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Judges** | [**PaginatedResultJudgePreview**](PaginatedResultJudgePreview.md) |  | 

## Methods

### NewListJudgesResponse

`func NewListJudgesResponse(judges PaginatedResultJudgePreview, ) *ListJudgesResponse`

NewListJudgesResponse instantiates a new ListJudgesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJudgesResponseWithDefaults

`func NewListJudgesResponseWithDefaults() *ListJudgesResponse`

NewListJudgesResponseWithDefaults instantiates a new ListJudgesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJudges

`func (o *ListJudgesResponse) GetJudges() PaginatedResultJudgePreview`

GetJudges returns the Judges field if non-nil, zero value otherwise.

### GetJudgesOk

`func (o *ListJudgesResponse) GetJudgesOk() (*PaginatedResultJudgePreview, bool)`

GetJudgesOk returns a tuple with the Judges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudges

`func (o *ListJudgesResponse) SetJudges(v PaginatedResultJudgePreview)`

SetJudges sets Judges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



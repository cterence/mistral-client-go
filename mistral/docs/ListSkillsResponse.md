# ListSkillsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListSkillsResponse

`func NewListSkillsResponse() *ListSkillsResponse`

NewListSkillsResponse instantiates a new ListSkillsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSkillsResponseWithDefaults

`func NewListSkillsResponseWithDefaults() *ListSkillsResponse`

NewListSkillsResponseWithDefaults instantiates a new ListSkillsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSkillsResponse) GetData() []Skill`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSkillsResponse) GetDataOk() (*[]Skill, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSkillsResponse) SetData(v []Skill)`

SetData sets Data field to given value.

### HasData

`func (o *ListSkillsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListSkillsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListSkillsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListSkillsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListSkillsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



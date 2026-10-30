# GrepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** |  | 
**Pattern** | **string** |  | 
**Mode** | Pointer to [**GrepMode**](GrepMode.md) |  | [optional] [default to GREPMODE_PHRASE]
**TopK** | Pointer to **int32** |  | [optional] [default to 5]
**ContentType** | Pointer to **string** |  | [optional] [default to "content"]

## Methods

### NewGrepRequest

`func NewGrepRequest(sourceId string, pattern string, ) *GrepRequest`

NewGrepRequest instantiates a new GrepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrepRequestWithDefaults

`func NewGrepRequestWithDefaults() *GrepRequest`

NewGrepRequestWithDefaults instantiates a new GrepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *GrepRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GrepRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GrepRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetPattern

`func (o *GrepRequest) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GrepRequest) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GrepRequest) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMode

`func (o *GrepRequest) GetMode() GrepMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GrepRequest) GetModeOk() (*GrepMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GrepRequest) SetMode(v GrepMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GrepRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTopK

`func (o *GrepRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *GrepRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *GrepRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *GrepRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetContentType

`func (o *GrepRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GrepRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GrepRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GrepRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



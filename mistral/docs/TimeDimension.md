# TimeDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Granularity** | Pointer to [**Granularity**](Granularity.md) |  | [optional] [default to GRANULARITY_AUTO]

## Methods

### NewTimeDimension

`func NewTimeDimension() *TimeDimension`

NewTimeDimension instantiates a new TimeDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeDimensionWithDefaults

`func NewTimeDimensionWithDefaults() *TimeDimension`

NewTimeDimensionWithDefaults instantiates a new TimeDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularity

`func (o *TimeDimension) GetGranularity() Granularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TimeDimension) GetGranularityOk() (*Granularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TimeDimension) SetGranularity(v Granularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *TimeDimension) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



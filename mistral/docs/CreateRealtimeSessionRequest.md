# CreateRealtimeSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purpose** | Pointer to **string** |  | [optional] [default to "realtime"]
**Model** | **string** |  | 
**TtlSeconds** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateRealtimeSessionRequest

`func NewCreateRealtimeSessionRequest(model string, ) *CreateRealtimeSessionRequest`

NewCreateRealtimeSessionRequest instantiates a new CreateRealtimeSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRealtimeSessionRequestWithDefaults

`func NewCreateRealtimeSessionRequestWithDefaults() *CreateRealtimeSessionRequest`

NewCreateRealtimeSessionRequestWithDefaults instantiates a new CreateRealtimeSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *CreateRealtimeSessionRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateRealtimeSessionRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateRealtimeSessionRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CreateRealtimeSessionRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetModel

`func (o *CreateRealtimeSessionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateRealtimeSessionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateRealtimeSessionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetTtlSeconds

`func (o *CreateRealtimeSessionRequest) GetTtlSeconds() int32`

GetTtlSeconds returns the TtlSeconds field if non-nil, zero value otherwise.

### GetTtlSecondsOk

`func (o *CreateRealtimeSessionRequest) GetTtlSecondsOk() (*int32, bool)`

GetTtlSecondsOk returns a tuple with the TtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSeconds

`func (o *CreateRealtimeSessionRequest) SetTtlSeconds(v int32)`

SetTtlSeconds sets TtlSeconds field to given value.

### HasTtlSeconds

`func (o *CreateRealtimeSessionRequest) HasTtlSeconds() bool`

HasTtlSeconds returns a boolean if a field has been set.

### SetTtlSecondsNil

`func (o *CreateRealtimeSessionRequest) SetTtlSecondsNil(b bool)`

 SetTtlSecondsNil sets the value for TtlSeconds to be an explicit nil

### UnsetTtlSeconds
`func (o *CreateRealtimeSessionRequest) UnsetTtlSeconds()`

UnsetTtlSeconds ensures that no value is present for TtlSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



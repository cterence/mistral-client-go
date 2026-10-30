# SystemFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableSystemNameFilterGroup**](SystemNameFilterGroup.md) |  | [optional] 
**Params** | Pointer to [**NullableSystemParamFilterGroup**](SystemParamFilterGroup.md) |  | [optional] 

## Methods

### NewSystemFilter

`func NewSystemFilter() *SystemFilter`

NewSystemFilter instantiates a new SystemFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemFilterWithDefaults

`func NewSystemFilterWithDefaults() *SystemFilter`

NewSystemFilterWithDefaults instantiates a new SystemFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemFilter) GetName() SystemNameFilterGroup`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemFilter) GetNameOk() (*SystemNameFilterGroup, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemFilter) SetName(v SystemNameFilterGroup)`

SetName sets Name field to given value.

### HasName

`func (o *SystemFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SystemFilter) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SystemFilter) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParams

`func (o *SystemFilter) GetParams() SystemParamFilterGroup`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SystemFilter) GetParamsOk() (*SystemParamFilterGroup, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SystemFilter) SetParams(v SystemParamFilterGroup)`

SetParams sets Params field to given value.

### HasParams

`func (o *SystemFilter) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *SystemFilter) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *SystemFilter) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



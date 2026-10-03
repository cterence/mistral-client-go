# QueryInvocationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the query to request | 
**Input** | Pointer to [**NullableInput2**](Input2.md) |  | [optional] 

## Methods

### NewQueryInvocationBody

`func NewQueryInvocationBody(name string, ) *QueryInvocationBody`

NewQueryInvocationBody instantiates a new QueryInvocationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInvocationBodyWithDefaults

`func NewQueryInvocationBodyWithDefaults() *QueryInvocationBody`

NewQueryInvocationBodyWithDefaults instantiates a new QueryInvocationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryInvocationBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryInvocationBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryInvocationBody) SetName(v string)`

SetName sets Name field to given value.


### GetInput

`func (o *QueryInvocationBody) GetInput() Input2`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *QueryInvocationBody) GetInputOk() (*Input2, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *QueryInvocationBody) SetInput(v Input2)`

SetInput sets Input field to given value.

### HasInput

`func (o *QueryInvocationBody) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *QueryInvocationBody) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *QueryInvocationBody) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



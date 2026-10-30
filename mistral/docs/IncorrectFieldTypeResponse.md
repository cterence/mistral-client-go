# IncorrectFieldTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "incorrect_field_type"]
**Name** | **string** |  | 
**ExpectedType** | **string** |  | 
**ReceivedType** | **string** |  | 

## Methods

### NewIncorrectFieldTypeResponse

`func NewIncorrectFieldTypeResponse(name string, expectedType string, receivedType string, ) *IncorrectFieldTypeResponse`

NewIncorrectFieldTypeResponse instantiates a new IncorrectFieldTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncorrectFieldTypeResponseWithDefaults

`func NewIncorrectFieldTypeResponseWithDefaults() *IncorrectFieldTypeResponse`

NewIncorrectFieldTypeResponseWithDefaults instantiates a new IncorrectFieldTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncorrectFieldTypeResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncorrectFieldTypeResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncorrectFieldTypeResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncorrectFieldTypeResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IncorrectFieldTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncorrectFieldTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncorrectFieldTypeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetExpectedType

`func (o *IncorrectFieldTypeResponse) GetExpectedType() string`

GetExpectedType returns the ExpectedType field if non-nil, zero value otherwise.

### GetExpectedTypeOk

`func (o *IncorrectFieldTypeResponse) GetExpectedTypeOk() (*string, bool)`

GetExpectedTypeOk returns a tuple with the ExpectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedType

`func (o *IncorrectFieldTypeResponse) SetExpectedType(v string)`

SetExpectedType sets ExpectedType field to given value.


### GetReceivedType

`func (o *IncorrectFieldTypeResponse) GetReceivedType() string`

GetReceivedType returns the ReceivedType field if non-nil, zero value otherwise.

### GetReceivedTypeOk

`func (o *IncorrectFieldTypeResponse) GetReceivedTypeOk() (*string, bool)`

GetReceivedTypeOk returns a tuple with the ReceivedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedType

`func (o *IncorrectFieldTypeResponse) SetReceivedType(v string)`

SetReceivedType sets ReceivedType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



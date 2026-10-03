# ObservabilityErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ErrorCode** | [**NullableObservabilityErrorCode**](ObservabilityErrorCode.md) |  | 

## Methods

### NewObservabilityErrorDetail

`func NewObservabilityErrorDetail(message string, errorCode NullableObservabilityErrorCode, ) *ObservabilityErrorDetail`

NewObservabilityErrorDetail instantiates a new ObservabilityErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityErrorDetailWithDefaults

`func NewObservabilityErrorDetailWithDefaults() *ObservabilityErrorDetail`

NewObservabilityErrorDetailWithDefaults instantiates a new ObservabilityErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ObservabilityErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ObservabilityErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ObservabilityErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorCode

`func (o *ObservabilityErrorDetail) GetErrorCode() ObservabilityErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ObservabilityErrorDetail) GetErrorCodeOk() (*ObservabilityErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ObservabilityErrorDetail) SetErrorCode(v ObservabilityErrorCode)`

SetErrorCode sets ErrorCode field to given value.


### SetErrorCodeNil

`func (o *ObservabilityErrorDetail) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ObservabilityErrorDetail) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



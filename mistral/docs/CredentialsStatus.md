# CredentialsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusType** | [**AuthStatus**](AuthStatus.md) |  | 
**LastCheckedAt** | Pointer to **NullableTime** |  | [optional] 
**ErrorHttpCode** | Pointer to [**NullableHTTPStatus**](HTTPStatus.md) |  | [optional] 
**ErrorMessage** | Pointer to [**NullableCredentialsStatusErrorReason**](CredentialsStatusErrorReason.md) |  | [optional] 

## Methods

### NewCredentialsStatus

`func NewCredentialsStatus(statusType AuthStatus, ) *CredentialsStatus`

NewCredentialsStatus instantiates a new CredentialsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsStatusWithDefaults

`func NewCredentialsStatusWithDefaults() *CredentialsStatus`

NewCredentialsStatusWithDefaults instantiates a new CredentialsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusType

`func (o *CredentialsStatus) GetStatusType() AuthStatus`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *CredentialsStatus) GetStatusTypeOk() (*AuthStatus, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *CredentialsStatus) SetStatusType(v AuthStatus)`

SetStatusType sets StatusType field to given value.


### GetLastCheckedAt

`func (o *CredentialsStatus) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *CredentialsStatus) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *CredentialsStatus) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *CredentialsStatus) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *CredentialsStatus) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *CredentialsStatus) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetErrorHttpCode

`func (o *CredentialsStatus) GetErrorHttpCode() HTTPStatus`

GetErrorHttpCode returns the ErrorHttpCode field if non-nil, zero value otherwise.

### GetErrorHttpCodeOk

`func (o *CredentialsStatus) GetErrorHttpCodeOk() (*HTTPStatus, bool)`

GetErrorHttpCodeOk returns a tuple with the ErrorHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHttpCode

`func (o *CredentialsStatus) SetErrorHttpCode(v HTTPStatus)`

SetErrorHttpCode sets ErrorHttpCode field to given value.

### HasErrorHttpCode

`func (o *CredentialsStatus) HasErrorHttpCode() bool`

HasErrorHttpCode returns a boolean if a field has been set.

### SetErrorHttpCodeNil

`func (o *CredentialsStatus) SetErrorHttpCodeNil(b bool)`

 SetErrorHttpCodeNil sets the value for ErrorHttpCode to be an explicit nil

### UnsetErrorHttpCode
`func (o *CredentialsStatus) UnsetErrorHttpCode()`

UnsetErrorHttpCode ensures that no value is present for ErrorHttpCode, not even an explicit nil
### GetErrorMessage

`func (o *CredentialsStatus) GetErrorMessage() CredentialsStatusErrorReason`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CredentialsStatus) GetErrorMessageOk() (*CredentialsStatusErrorReason, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CredentialsStatus) SetErrorMessage(v CredentialsStatusErrorReason)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CredentialsStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CredentialsStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CredentialsStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



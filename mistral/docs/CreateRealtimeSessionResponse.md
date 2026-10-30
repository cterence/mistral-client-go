# CreateRealtimeSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "client.session"]
**Purpose** | [**ClientSessionPurpose**](ClientSessionPurpose.md) |  | 
**ExpiresAt** | **time.Time** |  | 
**ClientSecret** | [**ClientSecret**](ClientSecret.md) |  | 

## Methods

### NewCreateRealtimeSessionResponse

`func NewCreateRealtimeSessionResponse(purpose ClientSessionPurpose, expiresAt time.Time, clientSecret ClientSecret, ) *CreateRealtimeSessionResponse`

NewCreateRealtimeSessionResponse instantiates a new CreateRealtimeSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRealtimeSessionResponseWithDefaults

`func NewCreateRealtimeSessionResponseWithDefaults() *CreateRealtimeSessionResponse`

NewCreateRealtimeSessionResponseWithDefaults instantiates a new CreateRealtimeSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreateRealtimeSessionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateRealtimeSessionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateRealtimeSessionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreateRealtimeSessionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPurpose

`func (o *CreateRealtimeSessionResponse) GetPurpose() ClientSessionPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateRealtimeSessionResponse) GetPurposeOk() (*ClientSessionPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateRealtimeSessionResponse) SetPurpose(v ClientSessionPurpose)`

SetPurpose sets Purpose field to given value.


### GetExpiresAt

`func (o *CreateRealtimeSessionResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateRealtimeSessionResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateRealtimeSessionResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetClientSecret

`func (o *CreateRealtimeSessionResponse) GetClientSecret() ClientSecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateRealtimeSessionResponse) GetClientSecretOk() (*ClientSecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateRealtimeSessionResponse) SetClientSecret(v ClientSecret)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



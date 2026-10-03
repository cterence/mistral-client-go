# AuthUrlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthUrl** | **string** |  | 
**Ttl** | **int32** |  | 

## Methods

### NewAuthUrlResponse

`func NewAuthUrlResponse(authUrl string, ttl int32, ) *AuthUrlResponse`

NewAuthUrlResponse instantiates a new AuthUrlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUrlResponseWithDefaults

`func NewAuthUrlResponseWithDefaults() *AuthUrlResponse`

NewAuthUrlResponseWithDefaults instantiates a new AuthUrlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthUrl

`func (o *AuthUrlResponse) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *AuthUrlResponse) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *AuthUrlResponse) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetTtl

`func (o *AuthUrlResponse) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AuthUrlResponse) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AuthUrlResponse) SetTtl(v int32)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



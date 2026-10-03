# VoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] [default to {}]
**Gender** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**RetentionNotice** | Pointer to **int32** |  | [optional] [default to 30]
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **NullableString** |  | 

## Methods

### NewVoiceResponse

`func NewVoiceResponse(name string, id string, createdAt time.Time, userId NullableString, ) *VoiceResponse`

NewVoiceResponse instantiates a new VoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceResponseWithDefaults

`func NewVoiceResponseWithDefaults() *VoiceResponse`

NewVoiceResponseWithDefaults instantiates a new VoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VoiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VoiceResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VoiceResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VoiceResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *VoiceResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *VoiceResponse) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *VoiceResponse) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetLanguages

`func (o *VoiceResponse) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VoiceResponse) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VoiceResponse) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VoiceResponse) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetGender

`func (o *VoiceResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *VoiceResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *VoiceResponse) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *VoiceResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *VoiceResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *VoiceResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetAge

`func (o *VoiceResponse) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *VoiceResponse) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *VoiceResponse) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *VoiceResponse) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *VoiceResponse) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *VoiceResponse) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetTags

`func (o *VoiceResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VoiceResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VoiceResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VoiceResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VoiceResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VoiceResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *VoiceResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *VoiceResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *VoiceResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *VoiceResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *VoiceResponse) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *VoiceResponse) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetRetentionNotice

`func (o *VoiceResponse) GetRetentionNotice() int32`

GetRetentionNotice returns the RetentionNotice field if non-nil, zero value otherwise.

### GetRetentionNoticeOk

`func (o *VoiceResponse) GetRetentionNoticeOk() (*int32, bool)`

GetRetentionNoticeOk returns a tuple with the RetentionNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionNotice

`func (o *VoiceResponse) SetRetentionNotice(v int32)`

SetRetentionNotice sets RetentionNotice field to given value.

### HasRetentionNotice

`func (o *VoiceResponse) HasRetentionNotice() bool`

HasRetentionNotice returns a boolean if a field has been set.

### GetId

`func (o *VoiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VoiceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VoiceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VoiceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *VoiceResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VoiceResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VoiceResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *VoiceResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *VoiceResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



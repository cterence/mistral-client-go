# VoiceCreateRequest

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
**SampleAudio** | **string** | Base64-encoded audio file | 
**SampleFilename** | Pointer to **NullableString** | Original filename for extension detection | [optional] 

## Methods

### NewVoiceCreateRequest

`func NewVoiceCreateRequest(name string, sampleAudio string, ) *VoiceCreateRequest`

NewVoiceCreateRequest instantiates a new VoiceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceCreateRequestWithDefaults

`func NewVoiceCreateRequestWithDefaults() *VoiceCreateRequest`

NewVoiceCreateRequestWithDefaults instantiates a new VoiceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VoiceCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoiceCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoiceCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VoiceCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VoiceCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VoiceCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *VoiceCreateRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *VoiceCreateRequest) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *VoiceCreateRequest) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetLanguages

`func (o *VoiceCreateRequest) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VoiceCreateRequest) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VoiceCreateRequest) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VoiceCreateRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetGender

`func (o *VoiceCreateRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *VoiceCreateRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *VoiceCreateRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *VoiceCreateRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *VoiceCreateRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *VoiceCreateRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetAge

`func (o *VoiceCreateRequest) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *VoiceCreateRequest) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *VoiceCreateRequest) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *VoiceCreateRequest) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *VoiceCreateRequest) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *VoiceCreateRequest) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetTags

`func (o *VoiceCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VoiceCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VoiceCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VoiceCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VoiceCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VoiceCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetColor

`func (o *VoiceCreateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *VoiceCreateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *VoiceCreateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *VoiceCreateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *VoiceCreateRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *VoiceCreateRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetRetentionNotice

`func (o *VoiceCreateRequest) GetRetentionNotice() int32`

GetRetentionNotice returns the RetentionNotice field if non-nil, zero value otherwise.

### GetRetentionNoticeOk

`func (o *VoiceCreateRequest) GetRetentionNoticeOk() (*int32, bool)`

GetRetentionNoticeOk returns a tuple with the RetentionNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionNotice

`func (o *VoiceCreateRequest) SetRetentionNotice(v int32)`

SetRetentionNotice sets RetentionNotice field to given value.

### HasRetentionNotice

`func (o *VoiceCreateRequest) HasRetentionNotice() bool`

HasRetentionNotice returns a boolean if a field has been set.

### GetSampleAudio

`func (o *VoiceCreateRequest) GetSampleAudio() string`

GetSampleAudio returns the SampleAudio field if non-nil, zero value otherwise.

### GetSampleAudioOk

`func (o *VoiceCreateRequest) GetSampleAudioOk() (*string, bool)`

GetSampleAudioOk returns a tuple with the SampleAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleAudio

`func (o *VoiceCreateRequest) SetSampleAudio(v string)`

SetSampleAudio sets SampleAudio field to given value.


### GetSampleFilename

`func (o *VoiceCreateRequest) GetSampleFilename() string`

GetSampleFilename returns the SampleFilename field if non-nil, zero value otherwise.

### GetSampleFilenameOk

`func (o *VoiceCreateRequest) GetSampleFilenameOk() (*string, bool)`

GetSampleFilenameOk returns a tuple with the SampleFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFilename

`func (o *VoiceCreateRequest) SetSampleFilename(v string)`

SetSampleFilename sets SampleFilename field to given value.

### HasSampleFilename

`func (o *VoiceCreateRequest) HasSampleFilename() bool`

HasSampleFilename returns a boolean if a field has been set.

### SetSampleFilenameNil

`func (o *VoiceCreateRequest) SetSampleFilenameNil(b bool)`

 SetSampleFilenameNil sets the value for SampleFilename to be an explicit nil

### UnsetSampleFilename
`func (o *VoiceCreateRequest) UnsetSampleFilename()`

UnsetSampleFilename ensures that no value is present for SampleFilename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



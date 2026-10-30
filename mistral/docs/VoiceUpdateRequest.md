# VoiceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**RetentionNotice** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewVoiceUpdateRequest

`func NewVoiceUpdateRequest() *VoiceUpdateRequest`

NewVoiceUpdateRequest instantiates a new VoiceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceUpdateRequestWithDefaults

`func NewVoiceUpdateRequestWithDefaults() *VoiceUpdateRequest`

NewVoiceUpdateRequestWithDefaults instantiates a new VoiceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VoiceUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoiceUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoiceUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoiceUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VoiceUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VoiceUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLanguages

`func (o *VoiceUpdateRequest) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VoiceUpdateRequest) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VoiceUpdateRequest) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VoiceUpdateRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *VoiceUpdateRequest) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *VoiceUpdateRequest) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetGender

`func (o *VoiceUpdateRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *VoiceUpdateRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *VoiceUpdateRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *VoiceUpdateRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *VoiceUpdateRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *VoiceUpdateRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetAge

`func (o *VoiceUpdateRequest) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *VoiceUpdateRequest) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *VoiceUpdateRequest) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *VoiceUpdateRequest) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *VoiceUpdateRequest) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *VoiceUpdateRequest) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetTags

`func (o *VoiceUpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VoiceUpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VoiceUpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VoiceUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VoiceUpdateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VoiceUpdateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *VoiceUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoiceUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoiceUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoiceUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VoiceUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VoiceUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetentionNotice

`func (o *VoiceUpdateRequest) GetRetentionNotice() int32`

GetRetentionNotice returns the RetentionNotice field if non-nil, zero value otherwise.

### GetRetentionNoticeOk

`func (o *VoiceUpdateRequest) GetRetentionNoticeOk() (*int32, bool)`

GetRetentionNoticeOk returns a tuple with the RetentionNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionNotice

`func (o *VoiceUpdateRequest) SetRetentionNotice(v int32)`

SetRetentionNotice sets RetentionNotice field to given value.

### HasRetentionNotice

`func (o *VoiceUpdateRequest) HasRetentionNotice() bool`

HasRetentionNotice returns a boolean if a field has been set.

### SetRetentionNoticeNil

`func (o *VoiceUpdateRequest) SetRetentionNoticeNil(b bool)`

 SetRetentionNoticeNil sets the value for RetentionNotice to be an explicit nil

### UnsetRetentionNotice
`func (o *VoiceUpdateRequest) UnsetRetentionNotice()`

UnsetRetentionNotice ensures that no value is present for RetentionNotice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



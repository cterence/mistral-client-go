# GitCommitMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to [**NullableGitCommitAuthor**](GitCommitAuthor.md) |  | [optional] 
**HtmlUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGitCommitMetadata

`func NewGitCommitMetadata(sha string, ) *GitCommitMetadata`

NewGitCommitMetadata instantiates a new GitCommitMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitMetadataWithDefaults

`func NewGitCommitMetadataWithDefaults() *GitCommitMetadata`

NewGitCommitMetadataWithDefaults instantiates a new GitCommitMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *GitCommitMetadata) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitCommitMetadata) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitCommitMetadata) SetSha(v string)`

SetSha sets Sha field to given value.


### GetMessage

`func (o *GitCommitMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitCommitMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitCommitMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitCommitMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GitCommitMetadata) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GitCommitMetadata) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetAuthor

`func (o *GitCommitMetadata) GetAuthor() GitCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GitCommitMetadata) GetAuthorOk() (*GitCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GitCommitMetadata) SetAuthor(v GitCommitAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GitCommitMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *GitCommitMetadata) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *GitCommitMetadata) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetHtmlUrl

`func (o *GitCommitMetadata) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GitCommitMetadata) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GitCommitMetadata) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GitCommitMetadata) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### SetHtmlUrlNil

`func (o *GitCommitMetadata) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *GitCommitMetadata) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



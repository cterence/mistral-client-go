# VibeWorkByTimeStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagesCount** | **int32** |  | 
**FilesCount** | **int32** |  | 
**ImagesCount** | **int32** |  | 
**SpreadsheetsCount** | **int32** |  | 
**UniqueConversationsCount** | **int32** |  | 
**TimeBucket** | **NullableTime** |  | 
**UniqueUsersCount** | **int32** |  | 
**MessagesToAgentsCount** | **int32** |  | 
**UniqueAgentsCount** | **int32** |  | 

## Methods

### NewVibeWorkByTimeStat

`func NewVibeWorkByTimeStat(messagesCount int32, filesCount int32, imagesCount int32, spreadsheetsCount int32, uniqueConversationsCount int32, timeBucket NullableTime, uniqueUsersCount int32, messagesToAgentsCount int32, uniqueAgentsCount int32, ) *VibeWorkByTimeStat`

NewVibeWorkByTimeStat instantiates a new VibeWorkByTimeStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVibeWorkByTimeStatWithDefaults

`func NewVibeWorkByTimeStatWithDefaults() *VibeWorkByTimeStat`

NewVibeWorkByTimeStatWithDefaults instantiates a new VibeWorkByTimeStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagesCount

`func (o *VibeWorkByTimeStat) GetMessagesCount() int32`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *VibeWorkByTimeStat) GetMessagesCountOk() (*int32, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *VibeWorkByTimeStat) SetMessagesCount(v int32)`

SetMessagesCount sets MessagesCount field to given value.


### GetFilesCount

`func (o *VibeWorkByTimeStat) GetFilesCount() int32`

GetFilesCount returns the FilesCount field if non-nil, zero value otherwise.

### GetFilesCountOk

`func (o *VibeWorkByTimeStat) GetFilesCountOk() (*int32, bool)`

GetFilesCountOk returns a tuple with the FilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesCount

`func (o *VibeWorkByTimeStat) SetFilesCount(v int32)`

SetFilesCount sets FilesCount field to given value.


### GetImagesCount

`func (o *VibeWorkByTimeStat) GetImagesCount() int32`

GetImagesCount returns the ImagesCount field if non-nil, zero value otherwise.

### GetImagesCountOk

`func (o *VibeWorkByTimeStat) GetImagesCountOk() (*int32, bool)`

GetImagesCountOk returns a tuple with the ImagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesCount

`func (o *VibeWorkByTimeStat) SetImagesCount(v int32)`

SetImagesCount sets ImagesCount field to given value.


### GetSpreadsheetsCount

`func (o *VibeWorkByTimeStat) GetSpreadsheetsCount() int32`

GetSpreadsheetsCount returns the SpreadsheetsCount field if non-nil, zero value otherwise.

### GetSpreadsheetsCountOk

`func (o *VibeWorkByTimeStat) GetSpreadsheetsCountOk() (*int32, bool)`

GetSpreadsheetsCountOk returns a tuple with the SpreadsheetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadsheetsCount

`func (o *VibeWorkByTimeStat) SetSpreadsheetsCount(v int32)`

SetSpreadsheetsCount sets SpreadsheetsCount field to given value.


### GetUniqueConversationsCount

`func (o *VibeWorkByTimeStat) GetUniqueConversationsCount() int32`

GetUniqueConversationsCount returns the UniqueConversationsCount field if non-nil, zero value otherwise.

### GetUniqueConversationsCountOk

`func (o *VibeWorkByTimeStat) GetUniqueConversationsCountOk() (*int32, bool)`

GetUniqueConversationsCountOk returns a tuple with the UniqueConversationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueConversationsCount

`func (o *VibeWorkByTimeStat) SetUniqueConversationsCount(v int32)`

SetUniqueConversationsCount sets UniqueConversationsCount field to given value.


### GetTimeBucket

`func (o *VibeWorkByTimeStat) GetTimeBucket() time.Time`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *VibeWorkByTimeStat) GetTimeBucketOk() (*time.Time, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *VibeWorkByTimeStat) SetTimeBucket(v time.Time)`

SetTimeBucket sets TimeBucket field to given value.


### SetTimeBucketNil

`func (o *VibeWorkByTimeStat) SetTimeBucketNil(b bool)`

 SetTimeBucketNil sets the value for TimeBucket to be an explicit nil

### UnsetTimeBucket
`func (o *VibeWorkByTimeStat) UnsetTimeBucket()`

UnsetTimeBucket ensures that no value is present for TimeBucket, not even an explicit nil
### GetUniqueUsersCount

`func (o *VibeWorkByTimeStat) GetUniqueUsersCount() int32`

GetUniqueUsersCount returns the UniqueUsersCount field if non-nil, zero value otherwise.

### GetUniqueUsersCountOk

`func (o *VibeWorkByTimeStat) GetUniqueUsersCountOk() (*int32, bool)`

GetUniqueUsersCountOk returns a tuple with the UniqueUsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUsersCount

`func (o *VibeWorkByTimeStat) SetUniqueUsersCount(v int32)`

SetUniqueUsersCount sets UniqueUsersCount field to given value.


### GetMessagesToAgentsCount

`func (o *VibeWorkByTimeStat) GetMessagesToAgentsCount() int32`

GetMessagesToAgentsCount returns the MessagesToAgentsCount field if non-nil, zero value otherwise.

### GetMessagesToAgentsCountOk

`func (o *VibeWorkByTimeStat) GetMessagesToAgentsCountOk() (*int32, bool)`

GetMessagesToAgentsCountOk returns a tuple with the MessagesToAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesToAgentsCount

`func (o *VibeWorkByTimeStat) SetMessagesToAgentsCount(v int32)`

SetMessagesToAgentsCount sets MessagesToAgentsCount field to given value.


### GetUniqueAgentsCount

`func (o *VibeWorkByTimeStat) GetUniqueAgentsCount() int32`

GetUniqueAgentsCount returns the UniqueAgentsCount field if non-nil, zero value otherwise.

### GetUniqueAgentsCountOk

`func (o *VibeWorkByTimeStat) GetUniqueAgentsCountOk() (*int32, bool)`

GetUniqueAgentsCountOk returns a tuple with the UniqueAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueAgentsCount

`func (o *VibeWorkByTimeStat) SetUniqueAgentsCount(v int32)`

SetUniqueAgentsCount sets UniqueAgentsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



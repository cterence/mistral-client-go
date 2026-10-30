# DeleteDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedDocumentIds** | **[]string** |  | 
**Missing** | **[]string** |  | 

## Methods

### NewDeleteDocumentsResponse

`func NewDeleteDocumentsResponse(deletedDocumentIds []string, missing []string, ) *DeleteDocumentsResponse`

NewDeleteDocumentsResponse instantiates a new DeleteDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentsResponseWithDefaults

`func NewDeleteDocumentsResponseWithDefaults() *DeleteDocumentsResponse`

NewDeleteDocumentsResponseWithDefaults instantiates a new DeleteDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedDocumentIds

`func (o *DeleteDocumentsResponse) GetDeletedDocumentIds() []string`

GetDeletedDocumentIds returns the DeletedDocumentIds field if non-nil, zero value otherwise.

### GetDeletedDocumentIdsOk

`func (o *DeleteDocumentsResponse) GetDeletedDocumentIdsOk() (*[]string, bool)`

GetDeletedDocumentIdsOk returns a tuple with the DeletedDocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDocumentIds

`func (o *DeleteDocumentsResponse) SetDeletedDocumentIds(v []string)`

SetDeletedDocumentIds sets DeletedDocumentIds field to given value.


### GetMissing

`func (o *DeleteDocumentsResponse) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *DeleteDocumentsResponse) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *DeleteDocumentsResponse) SetMissing(v []string)`

SetMissing sets Missing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



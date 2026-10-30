# DeleteManagedIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | [**ManagedIndexStatus**](ManagedIndexStatus.md) |  | 

## Methods

### NewDeleteManagedIndexResponse

`func NewDeleteManagedIndexResponse(id string, name string, status ManagedIndexStatus, ) *DeleteManagedIndexResponse`

NewDeleteManagedIndexResponse instantiates a new DeleteManagedIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteManagedIndexResponseWithDefaults

`func NewDeleteManagedIndexResponseWithDefaults() *DeleteManagedIndexResponse`

NewDeleteManagedIndexResponseWithDefaults instantiates a new DeleteManagedIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteManagedIndexResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteManagedIndexResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteManagedIndexResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeleteManagedIndexResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteManagedIndexResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteManagedIndexResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DeleteManagedIndexResponse) GetStatus() ManagedIndexStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteManagedIndexResponse) GetStatusOk() (*ManagedIndexStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteManagedIndexResponse) SetStatus(v ManagedIndexStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



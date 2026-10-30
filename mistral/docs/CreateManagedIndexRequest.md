# CreateManagedIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Schema** | Pointer to [**ManagedIndexFields**](ManagedIndexFields.md) |  | [optional] 
**Config** | [**ManagedIndexConfig**](ManagedIndexConfig.md) |  | 

## Methods

### NewCreateManagedIndexRequest

`func NewCreateManagedIndexRequest(name string, config ManagedIndexConfig, ) *CreateManagedIndexRequest`

NewCreateManagedIndexRequest instantiates a new CreateManagedIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManagedIndexRequestWithDefaults

`func NewCreateManagedIndexRequestWithDefaults() *CreateManagedIndexRequest`

NewCreateManagedIndexRequestWithDefaults instantiates a new CreateManagedIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateManagedIndexRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateManagedIndexRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateManagedIndexRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *CreateManagedIndexRequest) GetSchema() ManagedIndexFields`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CreateManagedIndexRequest) GetSchemaOk() (*ManagedIndexFields, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CreateManagedIndexRequest) SetSchema(v ManagedIndexFields)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CreateManagedIndexRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetConfig

`func (o *CreateManagedIndexRequest) GetConfig() ManagedIndexConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateManagedIndexRequest) GetConfigOk() (*ManagedIndexConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateManagedIndexRequest) SetConfig(v ManagedIndexConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



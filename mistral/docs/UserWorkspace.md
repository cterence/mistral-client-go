# UserWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The workspace&#39;s unique identifier. | 
**Name** | **string** | The workspace&#39;s display name. | 
**OrganizationId** | **string** | The identifier of the organization this workspace belongs to. | 

## Methods

### NewUserWorkspace

`func NewUserWorkspace(id string, name string, organizationId string, ) *UserWorkspace`

NewUserWorkspace instantiates a new UserWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWorkspaceWithDefaults

`func NewUserWorkspaceWithDefaults() *UserWorkspace`

NewUserWorkspaceWithDefaults instantiates a new UserWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserWorkspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserWorkspace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserWorkspace) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserWorkspace) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *UserWorkspace) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserWorkspace) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserWorkspace) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



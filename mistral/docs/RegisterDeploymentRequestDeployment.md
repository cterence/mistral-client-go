# RegisterDeploymentRequestDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] [default to "offline"]
**Deployment** | [**RegisterDeploymentRequestVespaDeployment**](RegisterDeploymentRequestVespaDeployment.md) |  | 

## Methods

### NewRegisterDeploymentRequestDeployment

`func NewRegisterDeploymentRequestDeployment(name string, deployment RegisterDeploymentRequestVespaDeployment, ) *RegisterDeploymentRequestDeployment`

NewRegisterDeploymentRequestDeployment instantiates a new RegisterDeploymentRequestDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeploymentRequestDeploymentWithDefaults

`func NewRegisterDeploymentRequestDeploymentWithDefaults() *RegisterDeploymentRequestDeployment`

NewRegisterDeploymentRequestDeploymentWithDefaults instantiates a new RegisterDeploymentRequestDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterDeploymentRequestDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterDeploymentRequestDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterDeploymentRequestDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *RegisterDeploymentRequestDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterDeploymentRequestDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterDeploymentRequestDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisterDeploymentRequestDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployment

`func (o *RegisterDeploymentRequestDeployment) GetDeployment() RegisterDeploymentRequestVespaDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *RegisterDeploymentRequestDeployment) GetDeploymentOk() (*RegisterDeploymentRequestVespaDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *RegisterDeploymentRequestDeployment) SetDeployment(v RegisterDeploymentRequestVespaDeployment)`

SetDeployment sets Deployment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



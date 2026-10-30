# DeploymentObservedState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to **NullableString** |  | [optional] 
**AvailableReplicas** | Pointer to **int32** |  | [optional] [default to 0]
**ReadyReplicas** | Pointer to **int32** |  | [optional] [default to 0]
**Endpoint** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**LastSeen** | Pointer to **NullableTime** |  | [optional] 
**DeployedRevision** | Pointer to **NullableString** |  | [optional] 
**Generation** | Pointer to **NullableInt32** |  | [optional] 
**BuildState** | Pointer to [**NullableDeploymentBuildState**](DeploymentBuildState.md) |  | [optional] 

## Methods

### NewDeploymentObservedState

`func NewDeploymentObservedState() *DeploymentObservedState`

NewDeploymentObservedState instantiates a new DeploymentObservedState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentObservedStateWithDefaults

`func NewDeploymentObservedStateWithDefaults() *DeploymentObservedState`

NewDeploymentObservedStateWithDefaults instantiates a new DeploymentObservedState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *DeploymentObservedState) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *DeploymentObservedState) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *DeploymentObservedState) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *DeploymentObservedState) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *DeploymentObservedState) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *DeploymentObservedState) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetAvailableReplicas

`func (o *DeploymentObservedState) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *DeploymentObservedState) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *DeploymentObservedState) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *DeploymentObservedState) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *DeploymentObservedState) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *DeploymentObservedState) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *DeploymentObservedState) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *DeploymentObservedState) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetEndpoint

`func (o *DeploymentObservedState) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DeploymentObservedState) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DeploymentObservedState) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *DeploymentObservedState) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *DeploymentObservedState) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *DeploymentObservedState) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetMessage

`func (o *DeploymentObservedState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentObservedState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentObservedState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeploymentObservedState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DeploymentObservedState) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DeploymentObservedState) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLastSeen

`func (o *DeploymentObservedState) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DeploymentObservedState) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DeploymentObservedState) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DeploymentObservedState) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### SetLastSeenNil

`func (o *DeploymentObservedState) SetLastSeenNil(b bool)`

 SetLastSeenNil sets the value for LastSeen to be an explicit nil

### UnsetLastSeen
`func (o *DeploymentObservedState) UnsetLastSeen()`

UnsetLastSeen ensures that no value is present for LastSeen, not even an explicit nil
### GetDeployedRevision

`func (o *DeploymentObservedState) GetDeployedRevision() string`

GetDeployedRevision returns the DeployedRevision field if non-nil, zero value otherwise.

### GetDeployedRevisionOk

`func (o *DeploymentObservedState) GetDeployedRevisionOk() (*string, bool)`

GetDeployedRevisionOk returns a tuple with the DeployedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedRevision

`func (o *DeploymentObservedState) SetDeployedRevision(v string)`

SetDeployedRevision sets DeployedRevision field to given value.

### HasDeployedRevision

`func (o *DeploymentObservedState) HasDeployedRevision() bool`

HasDeployedRevision returns a boolean if a field has been set.

### SetDeployedRevisionNil

`func (o *DeploymentObservedState) SetDeployedRevisionNil(b bool)`

 SetDeployedRevisionNil sets the value for DeployedRevision to be an explicit nil

### UnsetDeployedRevision
`func (o *DeploymentObservedState) UnsetDeployedRevision()`

UnsetDeployedRevision ensures that no value is present for DeployedRevision, not even an explicit nil
### GetGeneration

`func (o *DeploymentObservedState) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *DeploymentObservedState) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *DeploymentObservedState) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *DeploymentObservedState) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### SetGenerationNil

`func (o *DeploymentObservedState) SetGenerationNil(b bool)`

 SetGenerationNil sets the value for Generation to be an explicit nil

### UnsetGeneration
`func (o *DeploymentObservedState) UnsetGeneration()`

UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
### GetBuildState

`func (o *DeploymentObservedState) GetBuildState() DeploymentBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *DeploymentObservedState) GetBuildStateOk() (*DeploymentBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *DeploymentObservedState) SetBuildState(v DeploymentBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *DeploymentObservedState) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.

### SetBuildStateNil

`func (o *DeploymentObservedState) SetBuildStateNil(b bool)`

 SetBuildStateNil sets the value for BuildState to be an explicit nil

### UnsetBuildState
`func (o *DeploymentObservedState) UnsetBuildState()`

UnsetBuildState ensures that no value is present for BuildState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



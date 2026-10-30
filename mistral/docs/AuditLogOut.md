# AuditLogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Time when the audit log entry was created. | 
**ActorType** | [**ActorType**](ActorType.md) | Type of actor that performed the action. | 
**ActorMetadata** | **map[string]string** | Details about who performed the action. | 
**EventType** | [**AuditLogEventType**](AuditLogEventType.md) | Type of action recorded in the audit log. | 
**EventMetadata** | **map[string]string** | Details about the recorded action. | 
**TargetType** | [**TargetType**](TargetType.md) | Type of resource affected by the action. | 
**TargetMetadata** | **map[string]string** | Details about the affected resource. | 
**LogId** | **int32** | Audit log entry ID. | 
**OrganizationUuid** | Pointer to **NullableString** | Organization ID for the audit log entry. | [optional] 
**WorkspaceUuid** | Pointer to **NullableString** | Workspace ID for the audit log entry. | [optional] 

## Methods

### NewAuditLogOut

`func NewAuditLogOut(createdAt time.Time, actorType ActorType, actorMetadata map[string]string, eventType AuditLogEventType, eventMetadata map[string]string, targetType TargetType, targetMetadata map[string]string, logId int32, ) *AuditLogOut`

NewAuditLogOut instantiates a new AuditLogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogOutWithDefaults

`func NewAuditLogOutWithDefaults() *AuditLogOut`

NewAuditLogOutWithDefaults instantiates a new AuditLogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AuditLogOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditLogOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditLogOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActorType

`func (o *AuditLogOut) GetActorType() ActorType`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *AuditLogOut) GetActorTypeOk() (*ActorType, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *AuditLogOut) SetActorType(v ActorType)`

SetActorType sets ActorType field to given value.


### GetActorMetadata

`func (o *AuditLogOut) GetActorMetadata() map[string]string`

GetActorMetadata returns the ActorMetadata field if non-nil, zero value otherwise.

### GetActorMetadataOk

`func (o *AuditLogOut) GetActorMetadataOk() (*map[string]string, bool)`

GetActorMetadataOk returns a tuple with the ActorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorMetadata

`func (o *AuditLogOut) SetActorMetadata(v map[string]string)`

SetActorMetadata sets ActorMetadata field to given value.


### GetEventType

`func (o *AuditLogOut) GetEventType() AuditLogEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AuditLogOut) GetEventTypeOk() (*AuditLogEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AuditLogOut) SetEventType(v AuditLogEventType)`

SetEventType sets EventType field to given value.


### GetEventMetadata

`func (o *AuditLogOut) GetEventMetadata() map[string]string`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *AuditLogOut) GetEventMetadataOk() (*map[string]string, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *AuditLogOut) SetEventMetadata(v map[string]string)`

SetEventMetadata sets EventMetadata field to given value.


### GetTargetType

`func (o *AuditLogOut) GetTargetType() TargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AuditLogOut) GetTargetTypeOk() (*TargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AuditLogOut) SetTargetType(v TargetType)`

SetTargetType sets TargetType field to given value.


### GetTargetMetadata

`func (o *AuditLogOut) GetTargetMetadata() map[string]string`

GetTargetMetadata returns the TargetMetadata field if non-nil, zero value otherwise.

### GetTargetMetadataOk

`func (o *AuditLogOut) GetTargetMetadataOk() (*map[string]string, bool)`

GetTargetMetadataOk returns a tuple with the TargetMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMetadata

`func (o *AuditLogOut) SetTargetMetadata(v map[string]string)`

SetTargetMetadata sets TargetMetadata field to given value.


### GetLogId

`func (o *AuditLogOut) GetLogId() int32`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *AuditLogOut) GetLogIdOk() (*int32, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *AuditLogOut) SetLogId(v int32)`

SetLogId sets LogId field to given value.


### GetOrganizationUuid

`func (o *AuditLogOut) GetOrganizationUuid() string`

GetOrganizationUuid returns the OrganizationUuid field if non-nil, zero value otherwise.

### GetOrganizationUuidOk

`func (o *AuditLogOut) GetOrganizationUuidOk() (*string, bool)`

GetOrganizationUuidOk returns a tuple with the OrganizationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUuid

`func (o *AuditLogOut) SetOrganizationUuid(v string)`

SetOrganizationUuid sets OrganizationUuid field to given value.

### HasOrganizationUuid

`func (o *AuditLogOut) HasOrganizationUuid() bool`

HasOrganizationUuid returns a boolean if a field has been set.

### SetOrganizationUuidNil

`func (o *AuditLogOut) SetOrganizationUuidNil(b bool)`

 SetOrganizationUuidNil sets the value for OrganizationUuid to be an explicit nil

### UnsetOrganizationUuid
`func (o *AuditLogOut) UnsetOrganizationUuid()`

UnsetOrganizationUuid ensures that no value is present for OrganizationUuid, not even an explicit nil
### GetWorkspaceUuid

`func (o *AuditLogOut) GetWorkspaceUuid() string`

GetWorkspaceUuid returns the WorkspaceUuid field if non-nil, zero value otherwise.

### GetWorkspaceUuidOk

`func (o *AuditLogOut) GetWorkspaceUuidOk() (*string, bool)`

GetWorkspaceUuidOk returns a tuple with the WorkspaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuid

`func (o *AuditLogOut) SetWorkspaceUuid(v string)`

SetWorkspaceUuid sets WorkspaceUuid field to given value.

### HasWorkspaceUuid

`func (o *AuditLogOut) HasWorkspaceUuid() bool`

HasWorkspaceUuid returns a boolean if a field has been set.

### SetWorkspaceUuidNil

`func (o *AuditLogOut) SetWorkspaceUuidNil(b bool)`

 SetWorkspaceUuidNil sets the value for WorkspaceUuid to be an explicit nil

### UnsetWorkspaceUuid
`func (o *AuditLogOut) UnsetWorkspaceUuid()`

UnsetWorkspaceUuid ensures that no value is present for WorkspaceUuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



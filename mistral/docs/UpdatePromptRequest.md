# UpdatePromptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Display title. | [optional] 
**Description** | Pointer to **string** | Display description. | [optional] 
**SharingScope** | Pointer to [**RegistrySharingScope**](RegistrySharingScope.md) | Registry sharing scope. | [optional] 
**WorkspaceRelation** | Pointer to [**ShareRelation**](ShareRelation.md) | Access every workspace member gets when sharing_scope is &#x60;workspace&#x60;. Defaults to &#x60;writer&#x60;, which is what a workspace publish granted before this field existed, so omitting it preserves today&#39;s behaviour exactly. Rejected unless sharing_scope is &#x60;workspace&#x60;. | [optional] 

## Methods

### NewUpdatePromptRequest

`func NewUpdatePromptRequest() *UpdatePromptRequest`

NewUpdatePromptRequest instantiates a new UpdatePromptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePromptRequestWithDefaults

`func NewUpdatePromptRequestWithDefaults() *UpdatePromptRequest`

NewUpdatePromptRequestWithDefaults instantiates a new UpdatePromptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdatePromptRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePromptRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePromptRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdatePromptRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePromptRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePromptRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePromptRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePromptRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSharingScope

`func (o *UpdatePromptRequest) GetSharingScope() RegistrySharingScope`

GetSharingScope returns the SharingScope field if non-nil, zero value otherwise.

### GetSharingScopeOk

`func (o *UpdatePromptRequest) GetSharingScopeOk() (*RegistrySharingScope, bool)`

GetSharingScopeOk returns a tuple with the SharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingScope

`func (o *UpdatePromptRequest) SetSharingScope(v RegistrySharingScope)`

SetSharingScope sets SharingScope field to given value.

### HasSharingScope

`func (o *UpdatePromptRequest) HasSharingScope() bool`

HasSharingScope returns a boolean if a field has been set.

### GetWorkspaceRelation

`func (o *UpdatePromptRequest) GetWorkspaceRelation() ShareRelation`

GetWorkspaceRelation returns the WorkspaceRelation field if non-nil, zero value otherwise.

### GetWorkspaceRelationOk

`func (o *UpdatePromptRequest) GetWorkspaceRelationOk() (*ShareRelation, bool)`

GetWorkspaceRelationOk returns a tuple with the WorkspaceRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRelation

`func (o *UpdatePromptRequest) SetWorkspaceRelation(v ShareRelation)`

SetWorkspaceRelation sets WorkspaceRelation field to given value.

### HasWorkspaceRelation

`func (o *UpdatePromptRequest) HasWorkspaceRelation() bool`

HasWorkspaceRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



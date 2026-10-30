# APIKeyExtendedOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | API key ID. | 
**Key** | **NullableString** | Plaintext API key value. Only returned at creation time. | 
**Name** | **NullableString** | Name of the API key. | 
**HiddenKey** | **string** | Masked API key value for display. | 
**CreatedAt** | **NullableTime** | Time when the API key was created. | 
**ExpirationDate** | **NullableString** | Date when the API key expires. | 
**Actions** | [**APIKeyActions**](APIKeyActions.md) | Per-action availability for this API key (e.g. whether it can be rotated or deleted). | 
**WorkspaceId** | **NullableString** | Workspace ID for the API key. | 
**WorkspaceName** | **NullableString** | Name of the Workspace for the API key. | 
**CreatedBy** | **NullableString** | User or API key that created this API key. | 
**Product** | [**NullableAPIKeyProduct**](APIKeyProduct.md) | Product the API key belongs to. | 
**LastUsed** | **NullableString** | Date when the API key was last used. | 
**CanDelete** | **NullableBool** | Whether you can delete this API key. | 

## Methods

### NewAPIKeyExtendedOUT

`func NewAPIKeyExtendedOUT(keyId string, key NullableString, name NullableString, hiddenKey string, createdAt NullableTime, expirationDate NullableString, actions APIKeyActions, workspaceId NullableString, workspaceName NullableString, createdBy NullableString, product NullableAPIKeyProduct, lastUsed NullableString, canDelete NullableBool, ) *APIKeyExtendedOUT`

NewAPIKeyExtendedOUT instantiates a new APIKeyExtendedOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyExtendedOUTWithDefaults

`func NewAPIKeyExtendedOUTWithDefaults() *APIKeyExtendedOUT`

NewAPIKeyExtendedOUTWithDefaults instantiates a new APIKeyExtendedOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *APIKeyExtendedOUT) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *APIKeyExtendedOUT) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *APIKeyExtendedOUT) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKey

`func (o *APIKeyExtendedOUT) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *APIKeyExtendedOUT) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *APIKeyExtendedOUT) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *APIKeyExtendedOUT) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *APIKeyExtendedOUT) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *APIKeyExtendedOUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIKeyExtendedOUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIKeyExtendedOUT) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *APIKeyExtendedOUT) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *APIKeyExtendedOUT) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHiddenKey

`func (o *APIKeyExtendedOUT) GetHiddenKey() string`

GetHiddenKey returns the HiddenKey field if non-nil, zero value otherwise.

### GetHiddenKeyOk

`func (o *APIKeyExtendedOUT) GetHiddenKeyOk() (*string, bool)`

GetHiddenKeyOk returns a tuple with the HiddenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenKey

`func (o *APIKeyExtendedOUT) SetHiddenKey(v string)`

SetHiddenKey sets HiddenKey field to given value.


### GetCreatedAt

`func (o *APIKeyExtendedOUT) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIKeyExtendedOUT) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIKeyExtendedOUT) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *APIKeyExtendedOUT) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *APIKeyExtendedOUT) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpirationDate

`func (o *APIKeyExtendedOUT) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *APIKeyExtendedOUT) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *APIKeyExtendedOUT) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *APIKeyExtendedOUT) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *APIKeyExtendedOUT) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetActions

`func (o *APIKeyExtendedOUT) GetActions() APIKeyActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *APIKeyExtendedOUT) GetActionsOk() (*APIKeyActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *APIKeyExtendedOUT) SetActions(v APIKeyActions)`

SetActions sets Actions field to given value.


### GetWorkspaceId

`func (o *APIKeyExtendedOUT) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *APIKeyExtendedOUT) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *APIKeyExtendedOUT) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### SetWorkspaceIdNil

`func (o *APIKeyExtendedOUT) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *APIKeyExtendedOUT) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetWorkspaceName

`func (o *APIKeyExtendedOUT) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *APIKeyExtendedOUT) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *APIKeyExtendedOUT) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.


### SetWorkspaceNameNil

`func (o *APIKeyExtendedOUT) SetWorkspaceNameNil(b bool)`

 SetWorkspaceNameNil sets the value for WorkspaceName to be an explicit nil

### UnsetWorkspaceName
`func (o *APIKeyExtendedOUT) UnsetWorkspaceName()`

UnsetWorkspaceName ensures that no value is present for WorkspaceName, not even an explicit nil
### GetCreatedBy

`func (o *APIKeyExtendedOUT) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APIKeyExtendedOUT) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APIKeyExtendedOUT) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *APIKeyExtendedOUT) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *APIKeyExtendedOUT) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetProduct

`func (o *APIKeyExtendedOUT) GetProduct() APIKeyProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *APIKeyExtendedOUT) GetProductOk() (*APIKeyProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *APIKeyExtendedOUT) SetProduct(v APIKeyProduct)`

SetProduct sets Product field to given value.


### SetProductNil

`func (o *APIKeyExtendedOUT) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *APIKeyExtendedOUT) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetLastUsed

`func (o *APIKeyExtendedOUT) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *APIKeyExtendedOUT) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *APIKeyExtendedOUT) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.


### SetLastUsedNil

`func (o *APIKeyExtendedOUT) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *APIKeyExtendedOUT) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetCanDelete

`func (o *APIKeyExtendedOUT) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *APIKeyExtendedOUT) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *APIKeyExtendedOUT) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.


### SetCanDeleteNil

`func (o *APIKeyExtendedOUT) SetCanDeleteNil(b bool)`

 SetCanDeleteNil sets the value for CanDelete to be an explicit nil

### UnsetCanDelete
`func (o *APIKeyExtendedOUT) UnsetCanDelete()`

UnsetCanDelete ensures that no value is present for CanDelete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



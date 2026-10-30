# ExtendedOAuthServerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** |  | 
**AuthorizationEndpoint** | Pointer to **NullableString** |  | [optional] 
**TokenEndpoint** | **string** |  | 
**RegistrationEndpoint** | Pointer to **NullableString** |  | [optional] 
**ScopesSupported** | Pointer to **[]string** |  | [optional] 
**ResponseTypesSupported** | Pointer to **[]string** |  | [optional] [default to {"code"}]
**ResponseModesSupported** | Pointer to **[]string** |  | [optional] 
**GrantTypesSupported** | Pointer to **[]string** |  | [optional] 
**TokenEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 
**TokenEndpointAuthSigningAlgValuesSupported** | Pointer to **[]string** |  | [optional] 
**ServiceDocumentation** | Pointer to **NullableString** |  | [optional] 
**UiLocalesSupported** | Pointer to **[]string** |  | [optional] 
**OpPolicyUri** | Pointer to **NullableString** |  | [optional] 
**OpTosUri** | Pointer to **NullableString** |  | [optional] 
**RevocationEndpoint** | Pointer to **NullableString** |  | [optional] 
**RevocationEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 
**RevocationEndpointAuthSigningAlgValuesSupported** | Pointer to **[]string** |  | [optional] 
**IntrospectionEndpoint** | Pointer to **NullableString** |  | [optional] 
**IntrospectionEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 
**IntrospectionEndpointAuthSigningAlgValuesSupported** | Pointer to **[]string** |  | [optional] 
**CodeChallengeMethodsSupported** | Pointer to **[]string** |  | [optional] 
**ClientIdMetadataDocumentSupported** | Pointer to **NullableBool** |  | [optional] 
**AuthorizationResponseIssParameterSupported** | Pointer to **NullableBool** |  | [optional] 
**AuthorizationGrantProfilesSupported** | Pointer to **[]string** |  | [optional] 
**XSource** | Pointer to [**NullableOAuthMetadataSource**](OAuthMetadataSource.md) |  | [optional] 
**XResourceUrl** | Pointer to **NullableString** |  | [optional] 
**XScope** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExtendedOAuthServerMetadata

`func NewExtendedOAuthServerMetadata(issuer string, tokenEndpoint string, ) *ExtendedOAuthServerMetadata`

NewExtendedOAuthServerMetadata instantiates a new ExtendedOAuthServerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedOAuthServerMetadataWithDefaults

`func NewExtendedOAuthServerMetadataWithDefaults() *ExtendedOAuthServerMetadata`

NewExtendedOAuthServerMetadataWithDefaults instantiates a new ExtendedOAuthServerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *ExtendedOAuthServerMetadata) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ExtendedOAuthServerMetadata) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ExtendedOAuthServerMetadata) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetAuthorizationEndpoint

`func (o *ExtendedOAuthServerMetadata) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *ExtendedOAuthServerMetadata) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *ExtendedOAuthServerMetadata) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *ExtendedOAuthServerMetadata) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### SetAuthorizationEndpointNil

`func (o *ExtendedOAuthServerMetadata) SetAuthorizationEndpointNil(b bool)`

 SetAuthorizationEndpointNil sets the value for AuthorizationEndpoint to be an explicit nil

### UnsetAuthorizationEndpoint
`func (o *ExtendedOAuthServerMetadata) UnsetAuthorizationEndpoint()`

UnsetAuthorizationEndpoint ensures that no value is present for AuthorizationEndpoint, not even an explicit nil
### GetTokenEndpoint

`func (o *ExtendedOAuthServerMetadata) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *ExtendedOAuthServerMetadata) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *ExtendedOAuthServerMetadata) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetRegistrationEndpoint

`func (o *ExtendedOAuthServerMetadata) GetRegistrationEndpoint() string`

GetRegistrationEndpoint returns the RegistrationEndpoint field if non-nil, zero value otherwise.

### GetRegistrationEndpointOk

`func (o *ExtendedOAuthServerMetadata) GetRegistrationEndpointOk() (*string, bool)`

GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEndpoint

`func (o *ExtendedOAuthServerMetadata) SetRegistrationEndpoint(v string)`

SetRegistrationEndpoint sets RegistrationEndpoint field to given value.

### HasRegistrationEndpoint

`func (o *ExtendedOAuthServerMetadata) HasRegistrationEndpoint() bool`

HasRegistrationEndpoint returns a boolean if a field has been set.

### SetRegistrationEndpointNil

`func (o *ExtendedOAuthServerMetadata) SetRegistrationEndpointNil(b bool)`

 SetRegistrationEndpointNil sets the value for RegistrationEndpoint to be an explicit nil

### UnsetRegistrationEndpoint
`func (o *ExtendedOAuthServerMetadata) UnsetRegistrationEndpoint()`

UnsetRegistrationEndpoint ensures that no value is present for RegistrationEndpoint, not even an explicit nil
### GetScopesSupported

`func (o *ExtendedOAuthServerMetadata) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *ExtendedOAuthServerMetadata) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.

### HasScopesSupported

`func (o *ExtendedOAuthServerMetadata) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.

### SetScopesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetScopesSupportedNil(b bool)`

 SetScopesSupportedNil sets the value for ScopesSupported to be an explicit nil

### UnsetScopesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetScopesSupported()`

UnsetScopesSupported ensures that no value is present for ScopesSupported, not even an explicit nil
### GetResponseTypesSupported

`func (o *ExtendedOAuthServerMetadata) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *ExtendedOAuthServerMetadata) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.

### HasResponseTypesSupported

`func (o *ExtendedOAuthServerMetadata) HasResponseTypesSupported() bool`

HasResponseTypesSupported returns a boolean if a field has been set.

### GetResponseModesSupported

`func (o *ExtendedOAuthServerMetadata) GetResponseModesSupported() []string`

GetResponseModesSupported returns the ResponseModesSupported field if non-nil, zero value otherwise.

### GetResponseModesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetResponseModesSupportedOk() (*[]string, bool)`

GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseModesSupported

`func (o *ExtendedOAuthServerMetadata) SetResponseModesSupported(v []string)`

SetResponseModesSupported sets ResponseModesSupported field to given value.

### HasResponseModesSupported

`func (o *ExtendedOAuthServerMetadata) HasResponseModesSupported() bool`

HasResponseModesSupported returns a boolean if a field has been set.

### SetResponseModesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetResponseModesSupportedNil(b bool)`

 SetResponseModesSupportedNil sets the value for ResponseModesSupported to be an explicit nil

### UnsetResponseModesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetResponseModesSupported()`

UnsetResponseModesSupported ensures that no value is present for ResponseModesSupported, not even an explicit nil
### GetGrantTypesSupported

`func (o *ExtendedOAuthServerMetadata) GetGrantTypesSupported() []string`

GetGrantTypesSupported returns the GrantTypesSupported field if non-nil, zero value otherwise.

### GetGrantTypesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetGrantTypesSupportedOk() (*[]string, bool)`

GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypesSupported

`func (o *ExtendedOAuthServerMetadata) SetGrantTypesSupported(v []string)`

SetGrantTypesSupported sets GrantTypesSupported field to given value.

### HasGrantTypesSupported

`func (o *ExtendedOAuthServerMetadata) HasGrantTypesSupported() bool`

HasGrantTypesSupported returns a boolean if a field has been set.

### SetGrantTypesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetGrantTypesSupportedNil(b bool)`

 SetGrantTypesSupportedNil sets the value for GrantTypesSupported to be an explicit nil

### UnsetGrantTypesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetGrantTypesSupported()`

UnsetGrantTypesSupported ensures that no value is present for GrantTypesSupported, not even an explicit nil
### GetTokenEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) GetTokenEndpointAuthMethodsSupported() []string`

GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodsSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetTokenEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) SetTokenEndpointAuthMethodsSupported(v []string)`

SetTokenEndpointAuthMethodsSupported sets TokenEndpointAuthMethodsSupported field to given value.

### HasTokenEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) HasTokenEndpointAuthMethodsSupported() bool`

HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.

### SetTokenEndpointAuthMethodsSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetTokenEndpointAuthMethodsSupportedNil(b bool)`

 SetTokenEndpointAuthMethodsSupportedNil sets the value for TokenEndpointAuthMethodsSupported to be an explicit nil

### UnsetTokenEndpointAuthMethodsSupported
`func (o *ExtendedOAuthServerMetadata) UnsetTokenEndpointAuthMethodsSupported()`

UnsetTokenEndpointAuthMethodsSupported ensures that no value is present for TokenEndpointAuthMethodsSupported, not even an explicit nil
### GetTokenEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) GetTokenEndpointAuthSigningAlgValuesSupported() []string`

GetTokenEndpointAuthSigningAlgValuesSupported returns the TokenEndpointAuthSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthSigningAlgValuesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetTokenEndpointAuthSigningAlgValuesSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthSigningAlgValuesSupportedOk returns a tuple with the TokenEndpointAuthSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) SetTokenEndpointAuthSigningAlgValuesSupported(v []string)`

SetTokenEndpointAuthSigningAlgValuesSupported sets TokenEndpointAuthSigningAlgValuesSupported field to given value.

### HasTokenEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) HasTokenEndpointAuthSigningAlgValuesSupported() bool`

HasTokenEndpointAuthSigningAlgValuesSupported returns a boolean if a field has been set.

### SetTokenEndpointAuthSigningAlgValuesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetTokenEndpointAuthSigningAlgValuesSupportedNil(b bool)`

 SetTokenEndpointAuthSigningAlgValuesSupportedNil sets the value for TokenEndpointAuthSigningAlgValuesSupported to be an explicit nil

### UnsetTokenEndpointAuthSigningAlgValuesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetTokenEndpointAuthSigningAlgValuesSupported()`

UnsetTokenEndpointAuthSigningAlgValuesSupported ensures that no value is present for TokenEndpointAuthSigningAlgValuesSupported, not even an explicit nil
### GetServiceDocumentation

`func (o *ExtendedOAuthServerMetadata) GetServiceDocumentation() string`

GetServiceDocumentation returns the ServiceDocumentation field if non-nil, zero value otherwise.

### GetServiceDocumentationOk

`func (o *ExtendedOAuthServerMetadata) GetServiceDocumentationOk() (*string, bool)`

GetServiceDocumentationOk returns a tuple with the ServiceDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDocumentation

`func (o *ExtendedOAuthServerMetadata) SetServiceDocumentation(v string)`

SetServiceDocumentation sets ServiceDocumentation field to given value.

### HasServiceDocumentation

`func (o *ExtendedOAuthServerMetadata) HasServiceDocumentation() bool`

HasServiceDocumentation returns a boolean if a field has been set.

### SetServiceDocumentationNil

`func (o *ExtendedOAuthServerMetadata) SetServiceDocumentationNil(b bool)`

 SetServiceDocumentationNil sets the value for ServiceDocumentation to be an explicit nil

### UnsetServiceDocumentation
`func (o *ExtendedOAuthServerMetadata) UnsetServiceDocumentation()`

UnsetServiceDocumentation ensures that no value is present for ServiceDocumentation, not even an explicit nil
### GetUiLocalesSupported

`func (o *ExtendedOAuthServerMetadata) GetUiLocalesSupported() []string`

GetUiLocalesSupported returns the UiLocalesSupported field if non-nil, zero value otherwise.

### GetUiLocalesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetUiLocalesSupportedOk() (*[]string, bool)`

GetUiLocalesSupportedOk returns a tuple with the UiLocalesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiLocalesSupported

`func (o *ExtendedOAuthServerMetadata) SetUiLocalesSupported(v []string)`

SetUiLocalesSupported sets UiLocalesSupported field to given value.

### HasUiLocalesSupported

`func (o *ExtendedOAuthServerMetadata) HasUiLocalesSupported() bool`

HasUiLocalesSupported returns a boolean if a field has been set.

### SetUiLocalesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetUiLocalesSupportedNil(b bool)`

 SetUiLocalesSupportedNil sets the value for UiLocalesSupported to be an explicit nil

### UnsetUiLocalesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetUiLocalesSupported()`

UnsetUiLocalesSupported ensures that no value is present for UiLocalesSupported, not even an explicit nil
### GetOpPolicyUri

`func (o *ExtendedOAuthServerMetadata) GetOpPolicyUri() string`

GetOpPolicyUri returns the OpPolicyUri field if non-nil, zero value otherwise.

### GetOpPolicyUriOk

`func (o *ExtendedOAuthServerMetadata) GetOpPolicyUriOk() (*string, bool)`

GetOpPolicyUriOk returns a tuple with the OpPolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpPolicyUri

`func (o *ExtendedOAuthServerMetadata) SetOpPolicyUri(v string)`

SetOpPolicyUri sets OpPolicyUri field to given value.

### HasOpPolicyUri

`func (o *ExtendedOAuthServerMetadata) HasOpPolicyUri() bool`

HasOpPolicyUri returns a boolean if a field has been set.

### SetOpPolicyUriNil

`func (o *ExtendedOAuthServerMetadata) SetOpPolicyUriNil(b bool)`

 SetOpPolicyUriNil sets the value for OpPolicyUri to be an explicit nil

### UnsetOpPolicyUri
`func (o *ExtendedOAuthServerMetadata) UnsetOpPolicyUri()`

UnsetOpPolicyUri ensures that no value is present for OpPolicyUri, not even an explicit nil
### GetOpTosUri

`func (o *ExtendedOAuthServerMetadata) GetOpTosUri() string`

GetOpTosUri returns the OpTosUri field if non-nil, zero value otherwise.

### GetOpTosUriOk

`func (o *ExtendedOAuthServerMetadata) GetOpTosUriOk() (*string, bool)`

GetOpTosUriOk returns a tuple with the OpTosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpTosUri

`func (o *ExtendedOAuthServerMetadata) SetOpTosUri(v string)`

SetOpTosUri sets OpTosUri field to given value.

### HasOpTosUri

`func (o *ExtendedOAuthServerMetadata) HasOpTosUri() bool`

HasOpTosUri returns a boolean if a field has been set.

### SetOpTosUriNil

`func (o *ExtendedOAuthServerMetadata) SetOpTosUriNil(b bool)`

 SetOpTosUriNil sets the value for OpTosUri to be an explicit nil

### UnsetOpTosUri
`func (o *ExtendedOAuthServerMetadata) UnsetOpTosUri()`

UnsetOpTosUri ensures that no value is present for OpTosUri, not even an explicit nil
### GetRevocationEndpoint

`func (o *ExtendedOAuthServerMetadata) GetRevocationEndpoint() string`

GetRevocationEndpoint returns the RevocationEndpoint field if non-nil, zero value otherwise.

### GetRevocationEndpointOk

`func (o *ExtendedOAuthServerMetadata) GetRevocationEndpointOk() (*string, bool)`

GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpoint

`func (o *ExtendedOAuthServerMetadata) SetRevocationEndpoint(v string)`

SetRevocationEndpoint sets RevocationEndpoint field to given value.

### HasRevocationEndpoint

`func (o *ExtendedOAuthServerMetadata) HasRevocationEndpoint() bool`

HasRevocationEndpoint returns a boolean if a field has been set.

### SetRevocationEndpointNil

`func (o *ExtendedOAuthServerMetadata) SetRevocationEndpointNil(b bool)`

 SetRevocationEndpointNil sets the value for RevocationEndpoint to be an explicit nil

### UnsetRevocationEndpoint
`func (o *ExtendedOAuthServerMetadata) UnsetRevocationEndpoint()`

UnsetRevocationEndpoint ensures that no value is present for RevocationEndpoint, not even an explicit nil
### GetRevocationEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) GetRevocationEndpointAuthMethodsSupported() []string`

GetRevocationEndpointAuthMethodsSupported returns the RevocationEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetRevocationEndpointAuthMethodsSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetRevocationEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetRevocationEndpointAuthMethodsSupportedOk returns a tuple with the RevocationEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) SetRevocationEndpointAuthMethodsSupported(v []string)`

SetRevocationEndpointAuthMethodsSupported sets RevocationEndpointAuthMethodsSupported field to given value.

### HasRevocationEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) HasRevocationEndpointAuthMethodsSupported() bool`

HasRevocationEndpointAuthMethodsSupported returns a boolean if a field has been set.

### SetRevocationEndpointAuthMethodsSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetRevocationEndpointAuthMethodsSupportedNil(b bool)`

 SetRevocationEndpointAuthMethodsSupportedNil sets the value for RevocationEndpointAuthMethodsSupported to be an explicit nil

### UnsetRevocationEndpointAuthMethodsSupported
`func (o *ExtendedOAuthServerMetadata) UnsetRevocationEndpointAuthMethodsSupported()`

UnsetRevocationEndpointAuthMethodsSupported ensures that no value is present for RevocationEndpointAuthMethodsSupported, not even an explicit nil
### GetRevocationEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) GetRevocationEndpointAuthSigningAlgValuesSupported() []string`

GetRevocationEndpointAuthSigningAlgValuesSupported returns the RevocationEndpointAuthSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetRevocationEndpointAuthSigningAlgValuesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetRevocationEndpointAuthSigningAlgValuesSupportedOk() (*[]string, bool)`

GetRevocationEndpointAuthSigningAlgValuesSupportedOk returns a tuple with the RevocationEndpointAuthSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) SetRevocationEndpointAuthSigningAlgValuesSupported(v []string)`

SetRevocationEndpointAuthSigningAlgValuesSupported sets RevocationEndpointAuthSigningAlgValuesSupported field to given value.

### HasRevocationEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) HasRevocationEndpointAuthSigningAlgValuesSupported() bool`

HasRevocationEndpointAuthSigningAlgValuesSupported returns a boolean if a field has been set.

### SetRevocationEndpointAuthSigningAlgValuesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetRevocationEndpointAuthSigningAlgValuesSupportedNil(b bool)`

 SetRevocationEndpointAuthSigningAlgValuesSupportedNil sets the value for RevocationEndpointAuthSigningAlgValuesSupported to be an explicit nil

### UnsetRevocationEndpointAuthSigningAlgValuesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetRevocationEndpointAuthSigningAlgValuesSupported()`

UnsetRevocationEndpointAuthSigningAlgValuesSupported ensures that no value is present for RevocationEndpointAuthSigningAlgValuesSupported, not even an explicit nil
### GetIntrospectionEndpoint

`func (o *ExtendedOAuthServerMetadata) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *ExtendedOAuthServerMetadata) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *ExtendedOAuthServerMetadata) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *ExtendedOAuthServerMetadata) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### SetIntrospectionEndpointNil

`func (o *ExtendedOAuthServerMetadata) SetIntrospectionEndpointNil(b bool)`

 SetIntrospectionEndpointNil sets the value for IntrospectionEndpoint to be an explicit nil

### UnsetIntrospectionEndpoint
`func (o *ExtendedOAuthServerMetadata) UnsetIntrospectionEndpoint()`

UnsetIntrospectionEndpoint ensures that no value is present for IntrospectionEndpoint, not even an explicit nil
### GetIntrospectionEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) GetIntrospectionEndpointAuthMethodsSupported() []string`

GetIntrospectionEndpointAuthMethodsSupported returns the IntrospectionEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetIntrospectionEndpointAuthMethodsSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetIntrospectionEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetIntrospectionEndpointAuthMethodsSupportedOk returns a tuple with the IntrospectionEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) SetIntrospectionEndpointAuthMethodsSupported(v []string)`

SetIntrospectionEndpointAuthMethodsSupported sets IntrospectionEndpointAuthMethodsSupported field to given value.

### HasIntrospectionEndpointAuthMethodsSupported

`func (o *ExtendedOAuthServerMetadata) HasIntrospectionEndpointAuthMethodsSupported() bool`

HasIntrospectionEndpointAuthMethodsSupported returns a boolean if a field has been set.

### SetIntrospectionEndpointAuthMethodsSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetIntrospectionEndpointAuthMethodsSupportedNil(b bool)`

 SetIntrospectionEndpointAuthMethodsSupportedNil sets the value for IntrospectionEndpointAuthMethodsSupported to be an explicit nil

### UnsetIntrospectionEndpointAuthMethodsSupported
`func (o *ExtendedOAuthServerMetadata) UnsetIntrospectionEndpointAuthMethodsSupported()`

UnsetIntrospectionEndpointAuthMethodsSupported ensures that no value is present for IntrospectionEndpointAuthMethodsSupported, not even an explicit nil
### GetIntrospectionEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) GetIntrospectionEndpointAuthSigningAlgValuesSupported() []string`

GetIntrospectionEndpointAuthSigningAlgValuesSupported returns the IntrospectionEndpointAuthSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetIntrospectionEndpointAuthSigningAlgValuesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetIntrospectionEndpointAuthSigningAlgValuesSupportedOk() (*[]string, bool)`

GetIntrospectionEndpointAuthSigningAlgValuesSupportedOk returns a tuple with the IntrospectionEndpointAuthSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) SetIntrospectionEndpointAuthSigningAlgValuesSupported(v []string)`

SetIntrospectionEndpointAuthSigningAlgValuesSupported sets IntrospectionEndpointAuthSigningAlgValuesSupported field to given value.

### HasIntrospectionEndpointAuthSigningAlgValuesSupported

`func (o *ExtendedOAuthServerMetadata) HasIntrospectionEndpointAuthSigningAlgValuesSupported() bool`

HasIntrospectionEndpointAuthSigningAlgValuesSupported returns a boolean if a field has been set.

### SetIntrospectionEndpointAuthSigningAlgValuesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetIntrospectionEndpointAuthSigningAlgValuesSupportedNil(b bool)`

 SetIntrospectionEndpointAuthSigningAlgValuesSupportedNil sets the value for IntrospectionEndpointAuthSigningAlgValuesSupported to be an explicit nil

### UnsetIntrospectionEndpointAuthSigningAlgValuesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetIntrospectionEndpointAuthSigningAlgValuesSupported()`

UnsetIntrospectionEndpointAuthSigningAlgValuesSupported ensures that no value is present for IntrospectionEndpointAuthSigningAlgValuesSupported, not even an explicit nil
### GetCodeChallengeMethodsSupported

`func (o *ExtendedOAuthServerMetadata) GetCodeChallengeMethodsSupported() []string`

GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field if non-nil, zero value otherwise.

### GetCodeChallengeMethodsSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetCodeChallengeMethodsSupportedOk() (*[]string, bool)`

GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethodsSupported

`func (o *ExtendedOAuthServerMetadata) SetCodeChallengeMethodsSupported(v []string)`

SetCodeChallengeMethodsSupported sets CodeChallengeMethodsSupported field to given value.

### HasCodeChallengeMethodsSupported

`func (o *ExtendedOAuthServerMetadata) HasCodeChallengeMethodsSupported() bool`

HasCodeChallengeMethodsSupported returns a boolean if a field has been set.

### SetCodeChallengeMethodsSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetCodeChallengeMethodsSupportedNil(b bool)`

 SetCodeChallengeMethodsSupportedNil sets the value for CodeChallengeMethodsSupported to be an explicit nil

### UnsetCodeChallengeMethodsSupported
`func (o *ExtendedOAuthServerMetadata) UnsetCodeChallengeMethodsSupported()`

UnsetCodeChallengeMethodsSupported ensures that no value is present for CodeChallengeMethodsSupported, not even an explicit nil
### GetClientIdMetadataDocumentSupported

`func (o *ExtendedOAuthServerMetadata) GetClientIdMetadataDocumentSupported() bool`

GetClientIdMetadataDocumentSupported returns the ClientIdMetadataDocumentSupported field if non-nil, zero value otherwise.

### GetClientIdMetadataDocumentSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetClientIdMetadataDocumentSupportedOk() (*bool, bool)`

GetClientIdMetadataDocumentSupportedOk returns a tuple with the ClientIdMetadataDocumentSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdMetadataDocumentSupported

`func (o *ExtendedOAuthServerMetadata) SetClientIdMetadataDocumentSupported(v bool)`

SetClientIdMetadataDocumentSupported sets ClientIdMetadataDocumentSupported field to given value.

### HasClientIdMetadataDocumentSupported

`func (o *ExtendedOAuthServerMetadata) HasClientIdMetadataDocumentSupported() bool`

HasClientIdMetadataDocumentSupported returns a boolean if a field has been set.

### SetClientIdMetadataDocumentSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetClientIdMetadataDocumentSupportedNil(b bool)`

 SetClientIdMetadataDocumentSupportedNil sets the value for ClientIdMetadataDocumentSupported to be an explicit nil

### UnsetClientIdMetadataDocumentSupported
`func (o *ExtendedOAuthServerMetadata) UnsetClientIdMetadataDocumentSupported()`

UnsetClientIdMetadataDocumentSupported ensures that no value is present for ClientIdMetadataDocumentSupported, not even an explicit nil
### GetAuthorizationResponseIssParameterSupported

`func (o *ExtendedOAuthServerMetadata) GetAuthorizationResponseIssParameterSupported() bool`

GetAuthorizationResponseIssParameterSupported returns the AuthorizationResponseIssParameterSupported field if non-nil, zero value otherwise.

### GetAuthorizationResponseIssParameterSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetAuthorizationResponseIssParameterSupportedOk() (*bool, bool)`

GetAuthorizationResponseIssParameterSupportedOk returns a tuple with the AuthorizationResponseIssParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationResponseIssParameterSupported

`func (o *ExtendedOAuthServerMetadata) SetAuthorizationResponseIssParameterSupported(v bool)`

SetAuthorizationResponseIssParameterSupported sets AuthorizationResponseIssParameterSupported field to given value.

### HasAuthorizationResponseIssParameterSupported

`func (o *ExtendedOAuthServerMetadata) HasAuthorizationResponseIssParameterSupported() bool`

HasAuthorizationResponseIssParameterSupported returns a boolean if a field has been set.

### SetAuthorizationResponseIssParameterSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetAuthorizationResponseIssParameterSupportedNil(b bool)`

 SetAuthorizationResponseIssParameterSupportedNil sets the value for AuthorizationResponseIssParameterSupported to be an explicit nil

### UnsetAuthorizationResponseIssParameterSupported
`func (o *ExtendedOAuthServerMetadata) UnsetAuthorizationResponseIssParameterSupported()`

UnsetAuthorizationResponseIssParameterSupported ensures that no value is present for AuthorizationResponseIssParameterSupported, not even an explicit nil
### GetAuthorizationGrantProfilesSupported

`func (o *ExtendedOAuthServerMetadata) GetAuthorizationGrantProfilesSupported() []string`

GetAuthorizationGrantProfilesSupported returns the AuthorizationGrantProfilesSupported field if non-nil, zero value otherwise.

### GetAuthorizationGrantProfilesSupportedOk

`func (o *ExtendedOAuthServerMetadata) GetAuthorizationGrantProfilesSupportedOk() (*[]string, bool)`

GetAuthorizationGrantProfilesSupportedOk returns a tuple with the AuthorizationGrantProfilesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGrantProfilesSupported

`func (o *ExtendedOAuthServerMetadata) SetAuthorizationGrantProfilesSupported(v []string)`

SetAuthorizationGrantProfilesSupported sets AuthorizationGrantProfilesSupported field to given value.

### HasAuthorizationGrantProfilesSupported

`func (o *ExtendedOAuthServerMetadata) HasAuthorizationGrantProfilesSupported() bool`

HasAuthorizationGrantProfilesSupported returns a boolean if a field has been set.

### SetAuthorizationGrantProfilesSupportedNil

`func (o *ExtendedOAuthServerMetadata) SetAuthorizationGrantProfilesSupportedNil(b bool)`

 SetAuthorizationGrantProfilesSupportedNil sets the value for AuthorizationGrantProfilesSupported to be an explicit nil

### UnsetAuthorizationGrantProfilesSupported
`func (o *ExtendedOAuthServerMetadata) UnsetAuthorizationGrantProfilesSupported()`

UnsetAuthorizationGrantProfilesSupported ensures that no value is present for AuthorizationGrantProfilesSupported, not even an explicit nil
### GetXSource

`func (o *ExtendedOAuthServerMetadata) GetXSource() OAuthMetadataSource`

GetXSource returns the XSource field if non-nil, zero value otherwise.

### GetXSourceOk

`func (o *ExtendedOAuthServerMetadata) GetXSourceOk() (*OAuthMetadataSource, bool)`

GetXSourceOk returns a tuple with the XSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXSource

`func (o *ExtendedOAuthServerMetadata) SetXSource(v OAuthMetadataSource)`

SetXSource sets XSource field to given value.

### HasXSource

`func (o *ExtendedOAuthServerMetadata) HasXSource() bool`

HasXSource returns a boolean if a field has been set.

### SetXSourceNil

`func (o *ExtendedOAuthServerMetadata) SetXSourceNil(b bool)`

 SetXSourceNil sets the value for XSource to be an explicit nil

### UnsetXSource
`func (o *ExtendedOAuthServerMetadata) UnsetXSource()`

UnsetXSource ensures that no value is present for XSource, not even an explicit nil
### GetXResourceUrl

`func (o *ExtendedOAuthServerMetadata) GetXResourceUrl() string`

GetXResourceUrl returns the XResourceUrl field if non-nil, zero value otherwise.

### GetXResourceUrlOk

`func (o *ExtendedOAuthServerMetadata) GetXResourceUrlOk() (*string, bool)`

GetXResourceUrlOk returns a tuple with the XResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXResourceUrl

`func (o *ExtendedOAuthServerMetadata) SetXResourceUrl(v string)`

SetXResourceUrl sets XResourceUrl field to given value.

### HasXResourceUrl

`func (o *ExtendedOAuthServerMetadata) HasXResourceUrl() bool`

HasXResourceUrl returns a boolean if a field has been set.

### SetXResourceUrlNil

`func (o *ExtendedOAuthServerMetadata) SetXResourceUrlNil(b bool)`

 SetXResourceUrlNil sets the value for XResourceUrl to be an explicit nil

### UnsetXResourceUrl
`func (o *ExtendedOAuthServerMetadata) UnsetXResourceUrl()`

UnsetXResourceUrl ensures that no value is present for XResourceUrl, not even an explicit nil
### GetXScope

`func (o *ExtendedOAuthServerMetadata) GetXScope() string`

GetXScope returns the XScope field if non-nil, zero value otherwise.

### GetXScopeOk

`func (o *ExtendedOAuthServerMetadata) GetXScopeOk() (*string, bool)`

GetXScopeOk returns a tuple with the XScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXScope

`func (o *ExtendedOAuthServerMetadata) SetXScope(v string)`

SetXScope sets XScope field to given value.

### HasXScope

`func (o *ExtendedOAuthServerMetadata) HasXScope() bool`

HasXScope returns a boolean if a field has been set.

### SetXScopeNil

`func (o *ExtendedOAuthServerMetadata) SetXScopeNil(b bool)`

 SetXScopeNil sets the value for XScope to be an explicit nil

### UnsetXScope
`func (o *ExtendedOAuthServerMetadata) UnsetXScope()`

UnsetXScope ensures that no value is present for XScope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



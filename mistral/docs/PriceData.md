# PriceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**NullableLagoEventType**](LagoEventType.md) | Billing event type for the price. | 
**BillingMetric** | **string** | Billing metric this price applies to. | 
**BillingGroup** | **string** | Billing metric group this price applies to. | 
**ApiZone** | [**ApiZone**](ApiZone.md) | API zone this price applies to. | 
**ServiceTier** | [**ServiceTier**](ServiceTier.md) | Service tier this price applies to. | 
**Price** | [**Price**](Price.md) |  | 

## Methods

### NewPriceData

`func NewPriceData(eventType NullableLagoEventType, billingMetric string, billingGroup string, apiZone ApiZone, serviceTier ServiceTier, price Price, ) *PriceData`

NewPriceData instantiates a new PriceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDataWithDefaults

`func NewPriceDataWithDefaults() *PriceData`

NewPriceDataWithDefaults instantiates a new PriceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *PriceData) GetEventType() LagoEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PriceData) GetEventTypeOk() (*LagoEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PriceData) SetEventType(v LagoEventType)`

SetEventType sets EventType field to given value.


### SetEventTypeNil

`func (o *PriceData) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *PriceData) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetBillingMetric

`func (o *PriceData) GetBillingMetric() string`

GetBillingMetric returns the BillingMetric field if non-nil, zero value otherwise.

### GetBillingMetricOk

`func (o *PriceData) GetBillingMetricOk() (*string, bool)`

GetBillingMetricOk returns a tuple with the BillingMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMetric

`func (o *PriceData) SetBillingMetric(v string)`

SetBillingMetric sets BillingMetric field to given value.


### GetBillingGroup

`func (o *PriceData) GetBillingGroup() string`

GetBillingGroup returns the BillingGroup field if non-nil, zero value otherwise.

### GetBillingGroupOk

`func (o *PriceData) GetBillingGroupOk() (*string, bool)`

GetBillingGroupOk returns a tuple with the BillingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroup

`func (o *PriceData) SetBillingGroup(v string)`

SetBillingGroup sets BillingGroup field to given value.


### GetApiZone

`func (o *PriceData) GetApiZone() ApiZone`

GetApiZone returns the ApiZone field if non-nil, zero value otherwise.

### GetApiZoneOk

`func (o *PriceData) GetApiZoneOk() (*ApiZone, bool)`

GetApiZoneOk returns a tuple with the ApiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiZone

`func (o *PriceData) SetApiZone(v ApiZone)`

SetApiZone sets ApiZone field to given value.


### GetServiceTier

`func (o *PriceData) GetServiceTier() ServiceTier`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *PriceData) GetServiceTierOk() (*ServiceTier, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *PriceData) SetServiceTier(v ServiceTier)`

SetServiceTier sets ServiceTier field to given value.


### GetPrice

`func (o *PriceData) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceData) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceData) SetPrice(v Price)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



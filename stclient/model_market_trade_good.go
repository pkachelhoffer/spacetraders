/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stclient

import (
	"encoding/json"
)

// checks if the MarketTradeGood type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketTradeGood{}

// MarketTradeGood struct for MarketTradeGood
type MarketTradeGood struct {
	// The symbol of the trade good.
	Symbol string `json:"symbol"`
	// The typical volume flowing through the market for this type of good. The larger the trade volume, the more stable prices will be.
	TradeVolume int32 `json:"tradeVolume"`
	// A rough estimate of the total supply of this good in the marketplace.
	Supply string `json:"supply"`
	// The price at which this good can be purchased from the market.
	PurchasePrice int32 `json:"purchasePrice"`
	// The price at which this good can be sold to the market.
	SellPrice int32 `json:"sellPrice"`
}

// NewMarketTradeGood instantiates a new MarketTradeGood object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketTradeGood(symbol string, tradeVolume int32, supply string, purchasePrice int32, sellPrice int32) *MarketTradeGood {
	this := MarketTradeGood{}
	this.Symbol = symbol
	this.TradeVolume = tradeVolume
	this.Supply = supply
	this.PurchasePrice = purchasePrice
	this.SellPrice = sellPrice
	return &this
}

// NewMarketTradeGoodWithDefaults instantiates a new MarketTradeGood object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketTradeGoodWithDefaults() *MarketTradeGood {
	this := MarketTradeGood{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *MarketTradeGood) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MarketTradeGood) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MarketTradeGood) SetSymbol(v string) {
	o.Symbol = v
}

// GetTradeVolume returns the TradeVolume field value
func (o *MarketTradeGood) GetTradeVolume() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TradeVolume
}

// GetTradeVolumeOk returns a tuple with the TradeVolume field value
// and a boolean to check if the value has been set.
func (o *MarketTradeGood) GetTradeVolumeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeVolume, true
}

// SetTradeVolume sets field value
func (o *MarketTradeGood) SetTradeVolume(v int32) {
	o.TradeVolume = v
}

// GetSupply returns the Supply field value
func (o *MarketTradeGood) GetSupply() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value
// and a boolean to check if the value has been set.
func (o *MarketTradeGood) GetSupplyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supply, true
}

// SetSupply sets field value
func (o *MarketTradeGood) SetSupply(v string) {
	o.Supply = v
}

// GetPurchasePrice returns the PurchasePrice field value
func (o *MarketTradeGood) GetPurchasePrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PurchasePrice
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value
// and a boolean to check if the value has been set.
func (o *MarketTradeGood) GetPurchasePriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasePrice, true
}

// SetPurchasePrice sets field value
func (o *MarketTradeGood) SetPurchasePrice(v int32) {
	o.PurchasePrice = v
}

// GetSellPrice returns the SellPrice field value
func (o *MarketTradeGood) GetSellPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SellPrice
}

// GetSellPriceOk returns a tuple with the SellPrice field value
// and a boolean to check if the value has been set.
func (o *MarketTradeGood) GetSellPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellPrice, true
}

// SetSellPrice sets field value
func (o *MarketTradeGood) SetSellPrice(v int32) {
	o.SellPrice = v
}

func (o MarketTradeGood) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketTradeGood) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["tradeVolume"] = o.TradeVolume
	toSerialize["supply"] = o.Supply
	toSerialize["purchasePrice"] = o.PurchasePrice
	toSerialize["sellPrice"] = o.SellPrice
	return toSerialize, nil
}

type NullableMarketTradeGood struct {
	value *MarketTradeGood
	isSet bool
}

func (v NullableMarketTradeGood) Get() *MarketTradeGood {
	return v.value
}

func (v *NullableMarketTradeGood) Set(val *MarketTradeGood) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketTradeGood) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketTradeGood) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketTradeGood(val *MarketTradeGood) *NullableMarketTradeGood {
	return &NullableMarketTradeGood{value: val, isSet: true}
}

func (v NullableMarketTradeGood) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketTradeGood) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



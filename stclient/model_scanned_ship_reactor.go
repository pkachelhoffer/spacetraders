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

// checks if the ScannedShipReactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedShipReactor{}

// ScannedShipReactor The reactor of the ship.
type ScannedShipReactor struct {
	Symbol string `json:"symbol"`
}

// NewScannedShipReactor instantiates a new ScannedShipReactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedShipReactor(symbol string) *ScannedShipReactor {
	this := ScannedShipReactor{}
	this.Symbol = symbol
	return &this
}

// NewScannedShipReactorWithDefaults instantiates a new ScannedShipReactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedShipReactorWithDefaults() *ScannedShipReactor {
	this := ScannedShipReactor{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ScannedShipReactor) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ScannedShipReactor) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ScannedShipReactor) SetSymbol(v string) {
	o.Symbol = v
}

func (o ScannedShipReactor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedShipReactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	return toSerialize, nil
}

type NullableScannedShipReactor struct {
	value *ScannedShipReactor
	isSet bool
}

func (v NullableScannedShipReactor) Get() *ScannedShipReactor {
	return v.value
}

func (v *NullableScannedShipReactor) Set(val *ScannedShipReactor) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedShipReactor) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedShipReactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedShipReactor(val *ScannedShipReactor) *NullableScannedShipReactor {
	return &NullableScannedShipReactor{value: val, isSet: true}
}

func (v NullableScannedShipReactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedShipReactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



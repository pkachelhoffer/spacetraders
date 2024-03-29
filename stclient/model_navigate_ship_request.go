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

// checks if the NavigateShipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NavigateShipRequest{}

// NavigateShipRequest struct for NavigateShipRequest
type NavigateShipRequest struct {
	// The target destination.
	WaypointSymbol string `json:"waypointSymbol"`
}

// NewNavigateShipRequest instantiates a new NavigateShipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNavigateShipRequest(waypointSymbol string) *NavigateShipRequest {
	this := NavigateShipRequest{}
	this.WaypointSymbol = waypointSymbol
	return &this
}

// NewNavigateShipRequestWithDefaults instantiates a new NavigateShipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNavigateShipRequestWithDefaults() *NavigateShipRequest {
	this := NavigateShipRequest{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *NavigateShipRequest) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *NavigateShipRequest) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *NavigateShipRequest) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

func (o NavigateShipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NavigateShipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	return toSerialize, nil
}

type NullableNavigateShipRequest struct {
	value *NavigateShipRequest
	isSet bool
}

func (v NullableNavigateShipRequest) Get() *NavigateShipRequest {
	return v.value
}

func (v *NullableNavigateShipRequest) Set(val *NavigateShipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigateShipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigateShipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigateShipRequest(val *NavigateShipRequest) *NullableNavigateShipRequest {
	return &NullableNavigateShipRequest{value: val, isSet: true}
}

func (v NullableNavigateShipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigateShipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



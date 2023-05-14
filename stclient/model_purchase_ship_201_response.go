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

// checks if the PurchaseShip201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShip201Response{}

// PurchaseShip201Response 
type PurchaseShip201Response struct {
	Data PurchaseShip201ResponseData `json:"data"`
}

// NewPurchaseShip201Response instantiates a new PurchaseShip201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShip201Response(data PurchaseShip201ResponseData) *PurchaseShip201Response {
	this := PurchaseShip201Response{}
	this.Data = data
	return &this
}

// NewPurchaseShip201ResponseWithDefaults instantiates a new PurchaseShip201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseShip201ResponseWithDefaults() *PurchaseShip201Response {
	this := PurchaseShip201Response{}
	return &this
}

// GetData returns the Data field value
func (o *PurchaseShip201Response) GetData() PurchaseShip201ResponseData {
	if o == nil {
		var ret PurchaseShip201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PurchaseShip201Response) GetDataOk() (*PurchaseShip201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PurchaseShip201Response) SetData(v PurchaseShip201ResponseData) {
	o.Data = v
}

func (o PurchaseShip201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseShip201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePurchaseShip201Response struct {
	value *PurchaseShip201Response
	isSet bool
}

func (v NullablePurchaseShip201Response) Get() *PurchaseShip201Response {
	return v.value
}

func (v *NullablePurchaseShip201Response) Set(val *PurchaseShip201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseShip201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseShip201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseShip201Response(val *PurchaseShip201Response) *NullablePurchaseShip201Response {
	return &NullablePurchaseShip201Response{value: val, isSet: true}
}

func (v NullablePurchaseShip201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseShip201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



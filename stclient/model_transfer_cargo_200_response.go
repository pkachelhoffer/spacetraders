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

// checks if the TransferCargo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferCargo200Response{}

// TransferCargo200Response 
type TransferCargo200Response struct {
	Data Jettison200ResponseData `json:"data"`
}

// NewTransferCargo200Response instantiates a new TransferCargo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCargo200Response(data Jettison200ResponseData) *TransferCargo200Response {
	this := TransferCargo200Response{}
	this.Data = data
	return &this
}

// NewTransferCargo200ResponseWithDefaults instantiates a new TransferCargo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCargo200ResponseWithDefaults() *TransferCargo200Response {
	this := TransferCargo200Response{}
	return &this
}

// GetData returns the Data field value
func (o *TransferCargo200Response) GetData() Jettison200ResponseData {
	if o == nil {
		var ret Jettison200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransferCargo200Response) GetDataOk() (*Jettison200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransferCargo200Response) SetData(v Jettison200ResponseData) {
	o.Data = v
}

func (o TransferCargo200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferCargo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTransferCargo200Response struct {
	value *TransferCargo200Response
	isSet bool
}

func (v NullableTransferCargo200Response) Get() *TransferCargo200Response {
	return v.value
}

func (v *NullableTransferCargo200Response) Set(val *TransferCargo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCargo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCargo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCargo200Response(val *TransferCargo200Response) *NullableTransferCargo200Response {
	return &NullableTransferCargo200Response{value: val, isSet: true}
}

func (v NullableTransferCargo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCargo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



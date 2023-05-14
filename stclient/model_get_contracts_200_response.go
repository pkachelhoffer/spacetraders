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

// checks if the GetContracts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContracts200Response{}

// GetContracts200Response 
type GetContracts200Response struct {
	Data []Contract `json:"data"`
	Meta Meta `json:"meta"`
}

// NewGetContracts200Response instantiates a new GetContracts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContracts200Response(data []Contract, meta Meta) *GetContracts200Response {
	this := GetContracts200Response{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewGetContracts200ResponseWithDefaults instantiates a new GetContracts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContracts200ResponseWithDefaults() *GetContracts200Response {
	this := GetContracts200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetContracts200Response) GetData() []Contract {
	if o == nil {
		var ret []Contract
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetContracts200Response) GetDataOk() ([]Contract, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetContracts200Response) SetData(v []Contract) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *GetContracts200Response) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetContracts200Response) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *GetContracts200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetContracts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContracts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

type NullableGetContracts200Response struct {
	value *GetContracts200Response
	isSet bool
}

func (v NullableGetContracts200Response) Get() *GetContracts200Response {
	return v.value
}

func (v *NullableGetContracts200Response) Set(val *GetContracts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContracts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContracts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContracts200Response(val *GetContracts200Response) *NullableGetContracts200Response {
	return &NullableGetContracts200Response{value: val, isSet: true}
}

func (v NullableGetContracts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContracts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



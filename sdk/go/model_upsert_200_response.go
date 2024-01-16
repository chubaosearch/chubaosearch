/*
Vearch Database API

API for sending dynamic records to the Vearch database.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vearch_client

import (
	"encoding/json"
)

// checks if the Upsert200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Upsert200Response{}

// Upsert200Response struct for Upsert200Response
type Upsert200Response struct {
	Msg *string `json:"msg,omitempty"`
}

// NewUpsert200Response instantiates a new Upsert200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsert200Response() *Upsert200Response {
	this := Upsert200Response{}
	return &this
}

// NewUpsert200ResponseWithDefaults instantiates a new Upsert200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsert200ResponseWithDefaults() *Upsert200Response {
	this := Upsert200Response{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *Upsert200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upsert200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *Upsert200Response) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *Upsert200Response) SetMsg(v string) {
	o.Msg = &v
}

func (o Upsert200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Upsert200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableUpsert200Response struct {
	value *Upsert200Response
	isSet bool
}

func (v NullableUpsert200Response) Get() *Upsert200Response {
	return v.value
}

func (v *NullableUpsert200Response) Set(val *Upsert200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsert200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsert200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsert200Response(val *Upsert200Response) *NullableUpsert200Response {
	return &NullableUpsert200Response{value: val, isSet: true}
}

func (v NullableUpsert200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsert200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



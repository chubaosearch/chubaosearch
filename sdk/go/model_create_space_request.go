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

// checks if the CreateSpaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSpaceRequest{}

// CreateSpaceRequest struct for CreateSpaceRequest
type CreateSpaceRequest struct {
	Name *string `json:"name,omitempty"`
	PartitionNum *int32 `json:"partition_num,omitempty"`
	ReplicaNum *int32 `json:"replica_num,omitempty"`
	Engine *CreateSpaceRequestEngine `json:"engine,omitempty"`
	Properties *map[string]CreateSpaceRequestPropertiesValue `json:"properties,omitempty"`
}

// NewCreateSpaceRequest instantiates a new CreateSpaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSpaceRequest() *CreateSpaceRequest {
	this := CreateSpaceRequest{}
	return &this
}

// NewCreateSpaceRequestWithDefaults instantiates a new CreateSpaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSpaceRequestWithDefaults() *CreateSpaceRequest {
	this := CreateSpaceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSpaceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpaceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSpaceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSpaceRequest) SetName(v string) {
	o.Name = &v
}

// GetPartitionNum returns the PartitionNum field value if set, zero value otherwise.
func (o *CreateSpaceRequest) GetPartitionNum() int32 {
	if o == nil || IsNil(o.PartitionNum) {
		var ret int32
		return ret
	}
	return *o.PartitionNum
}

// GetPartitionNumOk returns a tuple with the PartitionNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpaceRequest) GetPartitionNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PartitionNum) {
		return nil, false
	}
	return o.PartitionNum, true
}

// HasPartitionNum returns a boolean if a field has been set.
func (o *CreateSpaceRequest) HasPartitionNum() bool {
	if o != nil && !IsNil(o.PartitionNum) {
		return true
	}

	return false
}

// SetPartitionNum gets a reference to the given int32 and assigns it to the PartitionNum field.
func (o *CreateSpaceRequest) SetPartitionNum(v int32) {
	o.PartitionNum = &v
}

// GetReplicaNum returns the ReplicaNum field value if set, zero value otherwise.
func (o *CreateSpaceRequest) GetReplicaNum() int32 {
	if o == nil || IsNil(o.ReplicaNum) {
		var ret int32
		return ret
	}
	return *o.ReplicaNum
}

// GetReplicaNumOk returns a tuple with the ReplicaNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpaceRequest) GetReplicaNumOk() (*int32, bool) {
	if o == nil || IsNil(o.ReplicaNum) {
		return nil, false
	}
	return o.ReplicaNum, true
}

// HasReplicaNum returns a boolean if a field has been set.
func (o *CreateSpaceRequest) HasReplicaNum() bool {
	if o != nil && !IsNil(o.ReplicaNum) {
		return true
	}

	return false
}

// SetReplicaNum gets a reference to the given int32 and assigns it to the ReplicaNum field.
func (o *CreateSpaceRequest) SetReplicaNum(v int32) {
	o.ReplicaNum = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *CreateSpaceRequest) GetEngine() CreateSpaceRequestEngine {
	if o == nil || IsNil(o.Engine) {
		var ret CreateSpaceRequestEngine
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpaceRequest) GetEngineOk() (*CreateSpaceRequestEngine, bool) {
	if o == nil || IsNil(o.Engine) {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *CreateSpaceRequest) HasEngine() bool {
	if o != nil && !IsNil(o.Engine) {
		return true
	}

	return false
}

// SetEngine gets a reference to the given CreateSpaceRequestEngine and assigns it to the Engine field.
func (o *CreateSpaceRequest) SetEngine(v CreateSpaceRequestEngine) {
	o.Engine = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateSpaceRequest) GetProperties() map[string]CreateSpaceRequestPropertiesValue {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]CreateSpaceRequestPropertiesValue
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpaceRequest) GetPropertiesOk() (*map[string]CreateSpaceRequestPropertiesValue, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateSpaceRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]CreateSpaceRequestPropertiesValue and assigns it to the Properties field.
func (o *CreateSpaceRequest) SetProperties(v map[string]CreateSpaceRequestPropertiesValue) {
	o.Properties = &v
}

func (o CreateSpaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSpaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PartitionNum) {
		toSerialize["partition_num"] = o.PartitionNum
	}
	if !IsNil(o.ReplicaNum) {
		toSerialize["replica_num"] = o.ReplicaNum
	}
	if !IsNil(o.Engine) {
		toSerialize["engine"] = o.Engine
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableCreateSpaceRequest struct {
	value *CreateSpaceRequest
	isSet bool
}

func (v NullableCreateSpaceRequest) Get() *CreateSpaceRequest {
	return v.value
}

func (v *NullableCreateSpaceRequest) Set(val *CreateSpaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSpaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSpaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSpaceRequest(val *CreateSpaceRequest) *NullableCreateSpaceRequest {
	return &NullableCreateSpaceRequest{value: val, isSet: true}
}

func (v NullableCreateSpaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSpaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Тестовое задание для отбора на Backend Bootcamp

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DummyLoginGet500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DummyLoginGet500Response{}

// DummyLoginGet500Response struct for DummyLoginGet500Response
type DummyLoginGet500Response struct {
	// Описание ошибки
	Message string `json:"message"`
	// Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId *string `json:"request_id,omitempty"`
	// Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code *int32 `json:"code,omitempty"`
}

type _DummyLoginGet500Response DummyLoginGet500Response

// NewDummyLoginGet500Response instantiates a new DummyLoginGet500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyLoginGet500Response(message string) *DummyLoginGet500Response {
	this := DummyLoginGet500Response{}
	this.Message = message
	return &this
}

// NewDummyLoginGet500ResponseWithDefaults instantiates a new DummyLoginGet500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyLoginGet500ResponseWithDefaults() *DummyLoginGet500Response {
	this := DummyLoginGet500Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *DummyLoginGet500Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DummyLoginGet500Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DummyLoginGet500Response) SetMessage(v string) {
	o.Message = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DummyLoginGet500Response) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyLoginGet500Response) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DummyLoginGet500Response) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DummyLoginGet500Response) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DummyLoginGet500Response) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyLoginGet500Response) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DummyLoginGet500Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *DummyLoginGet500Response) SetCode(v int32) {
	o.Code = &v
}

func (o DummyLoginGet500Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DummyLoginGet500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

func (o *DummyLoginGet500Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDummyLoginGet500Response := _DummyLoginGet500Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDummyLoginGet500Response)

	if err != nil {
		return err
	}

	*o = DummyLoginGet500Response(varDummyLoginGet500Response)

	return err
}

type NullableDummyLoginGet500Response struct {
	value *DummyLoginGet500Response
	isSet bool
}

func (v NullableDummyLoginGet500Response) Get() *DummyLoginGet500Response {
	return v.value
}

func (v *NullableDummyLoginGet500Response) Set(val *DummyLoginGet500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyLoginGet500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyLoginGet500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyLoginGet500Response(val *DummyLoginGet500Response) *NullableDummyLoginGet500Response {
	return &NullableDummyLoginGet500Response{value: val, isSet: true}
}

func (v NullableDummyLoginGet500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyLoginGet500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

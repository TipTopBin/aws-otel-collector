/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Info struct for Info
type Info struct {
	// The API entry point.
	Href *string `json:"href,omitempty"`
	// The API name.
	Name *string `json:"name,omitempty"`
	// The API version.
	Version *string `json:"version,omitempty"`
}

// NewInfo instantiates a new Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfo() *Info {
	this := Info{}

	return &this
}

// NewInfoWithDefaults instantiates a new Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoWithDefaults() *Info {
	this := Info{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Info) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Info) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Info) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Info) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Info) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Info) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *Info) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *Info) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Info) GetVersion() *string {
	if o == nil {
		return nil
	}

	return o.Version

}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Info) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Version, true
}

// SetVersion sets field value
func (o *Info) SetVersion(v string) {

	o.Version = &v

}

// HasVersion returns a boolean if a field has been set.
func (o *Info) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

func (o Info) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInfo struct {
	value *Info
	isSet bool
}

func (v NullableInfo) Get() *Info {
	return v.value
}

func (v *NullableInfo) Set(val *Info) {
	v.value = val
	v.isSet = true
}

func (v NullableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfo(val *Info) *NullableInfo {
	return &NullableInfo{value: val, isSet: true}
}

func (v NullableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

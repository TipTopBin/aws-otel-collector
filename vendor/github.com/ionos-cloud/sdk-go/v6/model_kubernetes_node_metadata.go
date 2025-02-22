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
	"time"
)

// KubernetesNodeMetadata struct for KubernetesNodeMetadata
type KubernetesNodeMetadata struct {
	// The resource entity tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11  Entity tags are also added as 'ETag' response headers to requests that do not use the 'depth' parameter.
	Etag *string `json:"etag,omitempty"`
	// The date the resource was created.
	CreatedDate *IonosTime
	// The date the resource was last modified.
	LastModifiedDate *IonosTime
	// The resource state.
	State *string `json:"state,omitempty"`
	// The date when the software on the node was last updated.
	LastSoftwareUpdatedDate *IonosTime
}

// NewKubernetesNodeMetadata instantiates a new KubernetesNodeMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeMetadata() *KubernetesNodeMetadata {
	this := KubernetesNodeMetadata{}

	return &this
}

// NewKubernetesNodeMetadataWithDefaults instantiates a new KubernetesNodeMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeMetadataWithDefaults() *KubernetesNodeMetadata {
	this := KubernetesNodeMetadata{}
	return &this
}

// GetEtag returns the Etag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodeMetadata) GetEtag() *string {
	if o == nil {
		return nil
	}

	return o.Etag

}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeMetadata) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Etag, true
}

// SetEtag sets field value
func (o *KubernetesNodeMetadata) SetEtag(v string) {

	o.Etag = &v

}

// HasEtag returns a boolean if a field has been set.
func (o *KubernetesNodeMetadata) HasEtag() bool {
	if o != nil && o.Etag != nil {
		return true
	}

	return false
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *KubernetesNodeMetadata) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeMetadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *KubernetesNodeMetadata) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *KubernetesNodeMetadata) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetLastModifiedDate returns the LastModifiedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *KubernetesNodeMetadata) GetLastModifiedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModifiedDate == nil {
		return nil
	}
	return &o.LastModifiedDate.Time

}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeMetadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModifiedDate == nil {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true

}

// SetLastModifiedDate sets field value
func (o *KubernetesNodeMetadata) SetLastModifiedDate(v time.Time) {

	o.LastModifiedDate = &IonosTime{v}

}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *KubernetesNodeMetadata) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodeMetadata) GetState() *string {
	if o == nil {
		return nil
	}

	return o.State

}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeMetadata) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.State, true
}

// SetState sets field value
func (o *KubernetesNodeMetadata) SetState(v string) {

	o.State = &v

}

// HasState returns a boolean if a field has been set.
func (o *KubernetesNodeMetadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// GetLastSoftwareUpdatedDate returns the LastSoftwareUpdatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *KubernetesNodeMetadata) GetLastSoftwareUpdatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastSoftwareUpdatedDate == nil {
		return nil
	}
	return &o.LastSoftwareUpdatedDate.Time

}

// GetLastSoftwareUpdatedDateOk returns a tuple with the LastSoftwareUpdatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeMetadata) GetLastSoftwareUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastSoftwareUpdatedDate == nil {
		return nil, false
	}
	return &o.LastSoftwareUpdatedDate.Time, true

}

// SetLastSoftwareUpdatedDate sets field value
func (o *KubernetesNodeMetadata) SetLastSoftwareUpdatedDate(v time.Time) {

	o.LastSoftwareUpdatedDate = &IonosTime{v}

}

// HasLastSoftwareUpdatedDate returns a boolean if a field has been set.
func (o *KubernetesNodeMetadata) HasLastSoftwareUpdatedDate() bool {
	if o != nil && o.LastSoftwareUpdatedDate != nil {
		return true
	}

	return false
}

func (o KubernetesNodeMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Etag != nil {
		toSerialize["etag"] = o.Etag
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.LastSoftwareUpdatedDate != nil {
		toSerialize["lastSoftwareUpdatedDate"] = o.LastSoftwareUpdatedDate
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodeMetadata struct {
	value *KubernetesNodeMetadata
	isSet bool
}

func (v NullableKubernetesNodeMetadata) Get() *KubernetesNodeMetadata {
	return v.value
}

func (v *NullableKubernetesNodeMetadata) Set(val *KubernetesNodeMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeMetadata(val *KubernetesNodeMetadata) *NullableKubernetesNodeMetadata {
	return &NullableKubernetesNodeMetadata{value: val, isSet: true}
}

func (v NullableKubernetesNodeMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

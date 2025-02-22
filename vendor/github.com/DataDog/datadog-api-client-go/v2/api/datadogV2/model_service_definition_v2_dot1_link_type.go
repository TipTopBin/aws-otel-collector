// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ServiceDefinitionV2Dot1LinkType Link type.
type ServiceDefinitionV2Dot1LinkType string

// List of ServiceDefinitionV2Dot1LinkType.
const (
	SERVICEDEFINITIONV2DOT1LINKTYPE_DOC       ServiceDefinitionV2Dot1LinkType = "doc"
	SERVICEDEFINITIONV2DOT1LINKTYPE_REPO      ServiceDefinitionV2Dot1LinkType = "repo"
	SERVICEDEFINITIONV2DOT1LINKTYPE_RUNBOOK   ServiceDefinitionV2Dot1LinkType = "runbook"
	SERVICEDEFINITIONV2DOT1LINKTYPE_DASHBOARD ServiceDefinitionV2Dot1LinkType = "dashboard"
	SERVICEDEFINITIONV2DOT1LINKTYPE_OTHER     ServiceDefinitionV2Dot1LinkType = "other"
)

var allowedServiceDefinitionV2Dot1LinkTypeEnumValues = []ServiceDefinitionV2Dot1LinkType{
	SERVICEDEFINITIONV2DOT1LINKTYPE_DOC,
	SERVICEDEFINITIONV2DOT1LINKTYPE_REPO,
	SERVICEDEFINITIONV2DOT1LINKTYPE_RUNBOOK,
	SERVICEDEFINITIONV2DOT1LINKTYPE_DASHBOARD,
	SERVICEDEFINITIONV2DOT1LINKTYPE_OTHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot1LinkType) GetAllowedValues() []ServiceDefinitionV2Dot1LinkType {
	return allowedServiceDefinitionV2Dot1LinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot1LinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot1LinkType(value)
	return nil
}

// NewServiceDefinitionV2Dot1LinkTypeFromValue returns a pointer to a valid ServiceDefinitionV2Dot1LinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot1LinkTypeFromValue(v string) (*ServiceDefinitionV2Dot1LinkType, error) {
	ev := ServiceDefinitionV2Dot1LinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot1LinkType: valid values are %v", v, allowedServiceDefinitionV2Dot1LinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot1LinkType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot1LinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot1LinkType value.
func (v ServiceDefinitionV2Dot1LinkType) Ptr() *ServiceDefinitionV2Dot1LinkType {
	return &v
}

// NullableServiceDefinitionV2Dot1LinkType handles when a null is used for ServiceDefinitionV2Dot1LinkType.
type NullableServiceDefinitionV2Dot1LinkType struct {
	value *ServiceDefinitionV2Dot1LinkType
	isSet bool
}

// Get returns the associated value.
func (v NullableServiceDefinitionV2Dot1LinkType) Get() *ServiceDefinitionV2Dot1LinkType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableServiceDefinitionV2Dot1LinkType) Set(val *ServiceDefinitionV2Dot1LinkType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableServiceDefinitionV2Dot1LinkType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableServiceDefinitionV2Dot1LinkType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceDefinitionV2Dot1LinkType initializes the struct as if Set has been called.
func NewNullableServiceDefinitionV2Dot1LinkType(val *ServiceDefinitionV2Dot1LinkType) *NullableServiceDefinitionV2Dot1LinkType {
	return &NullableServiceDefinitionV2Dot1LinkType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableServiceDefinitionV2Dot1LinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableServiceDefinitionV2Dot1LinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

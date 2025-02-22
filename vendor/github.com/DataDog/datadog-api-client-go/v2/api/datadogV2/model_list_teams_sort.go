// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ListTeamsSort Specifies the order of the returned teams
type ListTeamsSort string

// List of ListTeamsSort.
const (
	LISTTEAMSSORT_NAME        ListTeamsSort = "name"
	LISTTEAMSSORT__NAME       ListTeamsSort = "-name"
	LISTTEAMSSORT_USER_COUNT  ListTeamsSort = "user_count"
	LISTTEAMSSORT__USER_COUNT ListTeamsSort = "-user_count"
)

var allowedListTeamsSortEnumValues = []ListTeamsSort{
	LISTTEAMSSORT_NAME,
	LISTTEAMSSORT__NAME,
	LISTTEAMSSORT_USER_COUNT,
	LISTTEAMSSORT__USER_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListTeamsSort) GetAllowedValues() []ListTeamsSort {
	return allowedListTeamsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListTeamsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListTeamsSort(value)
	return nil
}

// NewListTeamsSortFromValue returns a pointer to a valid ListTeamsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListTeamsSortFromValue(v string) (*ListTeamsSort, error) {
	ev := ListTeamsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListTeamsSort: valid values are %v", v, allowedListTeamsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListTeamsSort) IsValid() bool {
	for _, existing := range allowedListTeamsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListTeamsSort value.
func (v ListTeamsSort) Ptr() *ListTeamsSort {
	return &v
}

// NullableListTeamsSort handles when a null is used for ListTeamsSort.
type NullableListTeamsSort struct {
	value *ListTeamsSort
	isSet bool
}

// Get returns the associated value.
func (v NullableListTeamsSort) Get() *ListTeamsSort {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableListTeamsSort) Set(val *ListTeamsSort) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableListTeamsSort) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableListTeamsSort) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableListTeamsSort initializes the struct as if Set has been called.
func NewNullableListTeamsSort(val *ListTeamsSort) *NullableListTeamsSort {
	return &NullableListTeamsSort{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableListTeamsSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableListTeamsSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

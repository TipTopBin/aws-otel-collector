// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageTimeseriesHour The hourly usage of timeseries.
type UsageTimeseriesHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the number of custom metrics that are inputs for aggregations (metric configured is custom).
	NumCustomInputTimeseries *int64 `json:"num_custom_input_timeseries,omitempty"`
	// Contains the number of custom metrics that are outputs for aggregations (metric configured is custom).
	NumCustomOutputTimeseries *int64 `json:"num_custom_output_timeseries,omitempty"`
	// Contains sum of non-aggregation custom metrics and custom metrics that are outputs for aggregations.
	NumCustomTimeseries *int64 `json:"num_custom_timeseries,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageTimeseriesHour instantiates a new UsageTimeseriesHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageTimeseriesHour() *UsageTimeseriesHour {
	this := UsageTimeseriesHour{}
	return &this
}

// NewUsageTimeseriesHourWithDefaults instantiates a new UsageTimeseriesHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageTimeseriesHourWithDefaults() *UsageTimeseriesHour {
	this := UsageTimeseriesHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageTimeseriesHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageTimeseriesHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageTimeseriesHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetNumCustomInputTimeseries returns the NumCustomInputTimeseries field value if set, zero value otherwise.
func (o *UsageTimeseriesHour) GetNumCustomInputTimeseries() int64 {
	if o == nil || o.NumCustomInputTimeseries == nil {
		var ret int64
		return ret
	}
	return *o.NumCustomInputTimeseries
}

// GetNumCustomInputTimeseriesOk returns a tuple with the NumCustomInputTimeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesHour) GetNumCustomInputTimeseriesOk() (*int64, bool) {
	if o == nil || o.NumCustomInputTimeseries == nil {
		return nil, false
	}
	return o.NumCustomInputTimeseries, true
}

// HasNumCustomInputTimeseries returns a boolean if a field has been set.
func (o *UsageTimeseriesHour) HasNumCustomInputTimeseries() bool {
	return o != nil && o.NumCustomInputTimeseries != nil
}

// SetNumCustomInputTimeseries gets a reference to the given int64 and assigns it to the NumCustomInputTimeseries field.
func (o *UsageTimeseriesHour) SetNumCustomInputTimeseries(v int64) {
	o.NumCustomInputTimeseries = &v
}

// GetNumCustomOutputTimeseries returns the NumCustomOutputTimeseries field value if set, zero value otherwise.
func (o *UsageTimeseriesHour) GetNumCustomOutputTimeseries() int64 {
	if o == nil || o.NumCustomOutputTimeseries == nil {
		var ret int64
		return ret
	}
	return *o.NumCustomOutputTimeseries
}

// GetNumCustomOutputTimeseriesOk returns a tuple with the NumCustomOutputTimeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesHour) GetNumCustomOutputTimeseriesOk() (*int64, bool) {
	if o == nil || o.NumCustomOutputTimeseries == nil {
		return nil, false
	}
	return o.NumCustomOutputTimeseries, true
}

// HasNumCustomOutputTimeseries returns a boolean if a field has been set.
func (o *UsageTimeseriesHour) HasNumCustomOutputTimeseries() bool {
	return o != nil && o.NumCustomOutputTimeseries != nil
}

// SetNumCustomOutputTimeseries gets a reference to the given int64 and assigns it to the NumCustomOutputTimeseries field.
func (o *UsageTimeseriesHour) SetNumCustomOutputTimeseries(v int64) {
	o.NumCustomOutputTimeseries = &v
}

// GetNumCustomTimeseries returns the NumCustomTimeseries field value if set, zero value otherwise.
func (o *UsageTimeseriesHour) GetNumCustomTimeseries() int64 {
	if o == nil || o.NumCustomTimeseries == nil {
		var ret int64
		return ret
	}
	return *o.NumCustomTimeseries
}

// GetNumCustomTimeseriesOk returns a tuple with the NumCustomTimeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesHour) GetNumCustomTimeseriesOk() (*int64, bool) {
	if o == nil || o.NumCustomTimeseries == nil {
		return nil, false
	}
	return o.NumCustomTimeseries, true
}

// HasNumCustomTimeseries returns a boolean if a field has been set.
func (o *UsageTimeseriesHour) HasNumCustomTimeseries() bool {
	return o != nil && o.NumCustomTimeseries != nil
}

// SetNumCustomTimeseries gets a reference to the given int64 and assigns it to the NumCustomTimeseries field.
func (o *UsageTimeseriesHour) SetNumCustomTimeseries(v int64) {
	o.NumCustomTimeseries = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageTimeseriesHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageTimeseriesHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageTimeseriesHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageTimeseriesHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTimeseriesHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageTimeseriesHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageTimeseriesHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageTimeseriesHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.NumCustomInputTimeseries != nil {
		toSerialize["num_custom_input_timeseries"] = o.NumCustomInputTimeseries
	}
	if o.NumCustomOutputTimeseries != nil {
		toSerialize["num_custom_output_timeseries"] = o.NumCustomOutputTimeseries
	}
	if o.NumCustomTimeseries != nil {
		toSerialize["num_custom_timeseries"] = o.NumCustomTimeseries
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageTimeseriesHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Hour                      *time.Time `json:"hour,omitempty"`
		NumCustomInputTimeseries  *int64     `json:"num_custom_input_timeseries,omitempty"`
		NumCustomOutputTimeseries *int64     `json:"num_custom_output_timeseries,omitempty"`
		NumCustomTimeseries       *int64     `json:"num_custom_timeseries,omitempty"`
		OrgName                   *string    `json:"org_name,omitempty"`
		PublicId                  *string    `json:"public_id,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hour", "num_custom_input_timeseries", "num_custom_output_timeseries", "num_custom_timeseries", "org_name", "public_id"})
	} else {
		return err
	}
	o.Hour = all.Hour
	o.NumCustomInputTimeseries = all.NumCustomInputTimeseries
	o.NumCustomOutputTimeseries = all.NumCustomOutputTimeseries
	o.NumCustomTimeseries = all.NumCustomTimeseries
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookAuthor Attributes of user object returned by the API.
type NotebookAuthor struct {
	// Creation time of the user.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Whether the user is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Email of the user.
	Email *string `json:"email,omitempty"`
	// Handle of the user.
	Handle *string `json:"handle,omitempty"`
	// URL of the user's icon.
	Icon *string `json:"icon,omitempty"`
	// Name of the user.
	Name datadog.NullableString `json:"name,omitempty"`
	// Status of the user.
	Status *string `json:"status,omitempty"`
	// Title of the user.
	Title datadog.NullableString `json:"title,omitempty"`
	// Whether the user is verified.
	Verified *bool `json:"verified,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewNotebookAuthor instantiates a new NotebookAuthor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookAuthor() *NotebookAuthor {
	this := NotebookAuthor{}
	return &this
}

// NewNotebookAuthorWithDefaults instantiates a new NotebookAuthor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookAuthorWithDefaults() *NotebookAuthor {
	this := NotebookAuthor{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotebookAuthor) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotebookAuthor) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotebookAuthor) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *NotebookAuthor) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *NotebookAuthor) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *NotebookAuthor) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NotebookAuthor) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NotebookAuthor) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NotebookAuthor) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *NotebookAuthor) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *NotebookAuthor) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *NotebookAuthor) SetHandle(v string) {
	o.Handle = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *NotebookAuthor) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *NotebookAuthor) HasIcon() bool {
	return o != nil && o.Icon != nil
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *NotebookAuthor) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookAuthor) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NotebookAuthor) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *NotebookAuthor) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *NotebookAuthor) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *NotebookAuthor) UnsetName() {
	o.Name.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotebookAuthor) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotebookAuthor) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotebookAuthor) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookAuthor) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookAuthor) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *NotebookAuthor) HasTitle() bool {
	return o != nil && o.Title.IsSet()
}

// SetTitle gets a reference to the given datadog.NullableString and assigns it to the Title field.
func (o *NotebookAuthor) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil.
func (o *NotebookAuthor) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil.
func (o *NotebookAuthor) UnsetTitle() {
	o.Title.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *NotebookAuthor) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAuthor) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *NotebookAuthor) HasVerified() bool {
	return o != nil && o.Verified != nil
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *NotebookAuthor) SetVerified(v bool) {
	o.Verified = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookAuthor) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CreatedAt *time.Time             `json:"created_at,omitempty"`
		Disabled  *bool                  `json:"disabled,omitempty"`
		Email     *string                `json:"email,omitempty"`
		Handle    *string                `json:"handle,omitempty"`
		Icon      *string                `json:"icon,omitempty"`
		Name      datadog.NullableString `json:"name,omitempty"`
		Status    *string                `json:"status,omitempty"`
		Title     datadog.NullableString `json:"title,omitempty"`
		Verified  *bool                  `json:"verified,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "disabled", "email", "handle", "icon", "name", "status", "title", "verified"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Disabled = all.Disabled
	o.Email = all.Email
	o.Handle = all.Handle
	o.Icon = all.Icon
	o.Name = all.Name
	o.Status = all.Status
	o.Title = all.Title
	o.Verified = all.Verified
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

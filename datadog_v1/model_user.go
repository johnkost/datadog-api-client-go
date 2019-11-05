/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
)

// User struct for User
type User struct {
	AccessRole *string `json:"access_role,omitempty"`

	Disabled *bool `json:"disabled,omitempty"`

	Email *string `json:"email,omitempty"`

	Handle *string `json:"handle,omitempty"`

	Icon *string `json:"icon,omitempty"`

	Name *string `json:"name,omitempty"`

	Verified *bool `json:"verified,omitempty"`
}

// GetAccessRole returns the AccessRole field if non-nil, zero value otherwise.
func (o *User) GetAccessRole() string {
	if o == nil || o.AccessRole == nil {
		var ret string
		return ret
	}
	return *o.AccessRole
}

// GetAccessRoleOk returns a tuple with the AccessRole field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccessRoleOk() (string, bool) {
	if o == nil || o.AccessRole == nil {
		var ret string
		return ret, false
	}
	return *o.AccessRole, true
}

// HasAccessRole returns a boolean if a field has been set.
func (o *User) HasAccessRole() bool {
	if o != nil && o.AccessRole != nil {
		return true
	}

	return false
}

// SetAccessRole gets a reference to the given string and assigns it to the AccessRole field.
func (o *User) SetAccessRole(v string) {
	o.AccessRole = &v
}

// GetDisabled returns the Disabled field if non-nil, zero value otherwise.
func (o *User) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDisabledOk() (bool, bool) {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret, false
	}
	return *o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *User) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *User) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEmail returns the Email field if non-nil, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field if non-nil, zero value otherwise.
func (o *User) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetHandleOk() (string, bool) {
	if o == nil || o.Handle == nil {
		var ret string
		return ret, false
	}
	return *o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *User) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *User) SetHandle(v string) {
	o.Handle = &v
}

// GetIcon returns the Icon field if non-nil, zero value otherwise.
func (o *User) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIconOk() (string, bool) {
	if o == nil || o.Icon == nil {
		var ret string
		return ret, false
	}
	return *o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *User) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *User) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetVerified returns the Verified field if non-nil, zero value otherwise.
func (o *User) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifiedOk() (bool, bool) {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret, false
	}
	return *o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *User) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *User) SetVerified(v bool) {
	o.Verified = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessRole != nil {
		toSerialize["access_role"] = o.AccessRole
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

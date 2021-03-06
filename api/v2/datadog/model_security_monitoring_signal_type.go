/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringSignalType The type of event.
type SecurityMonitoringSignalType string

// List of SecurityMonitoringSignalType
const (
	SECURITYMONITORINGSIGNALTYPE_SIGNAL SecurityMonitoringSignalType = "signal"
)

func (v *SecurityMonitoringSignalType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityMonitoringSignalType(value)
	for _, existing := range []SecurityMonitoringSignalType{"signal"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityMonitoringSignalType", value)
}

// Ptr returns reference to SecurityMonitoringSignalType value
func (v SecurityMonitoringSignalType) Ptr() *SecurityMonitoringSignalType {
	return &v
}

type NullableSecurityMonitoringSignalType struct {
	value *SecurityMonitoringSignalType
	isSet bool
}

func (v NullableSecurityMonitoringSignalType) Get() *SecurityMonitoringSignalType {
	return v.value
}

func (v *NullableSecurityMonitoringSignalType) Set(val *SecurityMonitoringSignalType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringSignalType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringSignalType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringSignalType(val *SecurityMonitoringSignalType) *NullableSecurityMonitoringSignalType {
	return &NullableSecurityMonitoringSignalType{value: val, isSet: true}
}

func (v NullableSecurityMonitoringSignalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringSignalType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

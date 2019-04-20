// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package timeEvent

// Request holds the entire request specific to Time Events
type Request struct {
	RequestNumber    string      `json:"requestNumber,omitempty"`
	SourceID         string      `json:"sourceId,omitempty"`
	RequestTimeStamp string      `json:"requestTimestamp,omitempty"`
	TimeEvents       []TimeEvent `json:"timeEvents,omitempty"`
}

// Attribute holds the attribute values
type Attribute struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// TimeEvent holds the TimeEvent specific values
type TimeEvent struct {
	DeviceID            string      `json:"deviceId,omitempty"`
	EventDateTime       string      `json:"eventDateTime,omitempty"`
	SupplierDeviceEvent string      `json:"supplierDeviceEvent,omitempty"`
	ReporterID          int         `json:"reporterId,omitempty"`
	ReporterIDType      string      `json:"reporterIdType,omitempty"`
	TimeEventAttributes []Attribute `json:"timeEventAttributes,omitempty"`
}

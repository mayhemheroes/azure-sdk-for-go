//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

const (
	moduleName    = "armdns"
	moduleVersion = "v1.0.0"
)

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCAA   RecordType = "CAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypeNS    RecordType = "NS"
	RecordTypePTR   RecordType = "PTR"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
)

// PossibleRecordTypeValues returns the possible values for the RecordType const type.
func PossibleRecordTypeValues() []RecordType {
	return []RecordType{
		RecordTypeA,
		RecordTypeAAAA,
		RecordTypeCAA,
		RecordTypeCNAME,
		RecordTypeMX,
		RecordTypeNS,
		RecordTypePTR,
		RecordTypeSOA,
		RecordTypeSRV,
		RecordTypeTXT,
	}
}

// ZoneType - The type of this DNS zone (Public or Private).
type ZoneType string

const (
	ZoneTypePublic  ZoneType = "Public"
	ZoneTypePrivate ZoneType = "Private"
)

// PossibleZoneTypeValues returns the possible values for the ZoneType const type.
func PossibleZoneTypeValues() []ZoneType {
	return []ZoneType{
		ZoneTypePublic,
		ZoneTypePrivate,
	}
}

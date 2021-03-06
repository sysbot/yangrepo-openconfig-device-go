/*
Package openconfig is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by /go/src/github.com/openconfig/ygot/genutil/names.go
using the following YANG input files:
	- /work/public/release/models/network-instance/openconfig-network-instance.yang
	- /work/public/release/models/mpls/openconfig-mpls.yang
	- /work/public/release/models/optical-transport/openconfig-optical-amplifier.yang
	- /work/public/release/models/optical-transport/openconfig-terminal-device.yang
	- /work/public/release/models/optical-transport/openconfig-transport-line-protection.yang
	- /work/public/release/models/platform/openconfig-platform.yang
	- /work/public/release/models/policy/openconfig-routing-policy.yang
	- /work/public/release/models/lacp/openconfig-lacp.yang
	- /work/public/release/models/system/openconfig-system.yang
	- /work/public/release/models/lldp/openconfig-lldp.yang
	- /work/public/release/models/stp/openconfig-spanning-tree.yang
	- /work/public/release/models/interfaces/openconfig-interfaces.yang
	- /work/public/release/models/interfaces/openconfig-if-ip.yang
	- /work/public/release/models/interfaces/openconfig-if-aggregate.yang
	- /work/public/release/models/interfaces/openconfig-if-ethernet.yang
	- /work/public/release/models/interfaces/openconfig-if-ip-ext.yang
	- /work/public/release/models/local-routing/openconfig-local-routing.yang
	- /work/public/release/models/vlan/openconfig-vlan.yang
	- /work/public/release/models/vlan/openconfig-vlan-types.yang
Imported modules were sourced from:
	- /work/public/...
	- yang/...
*/
package openconfig

import (
	"github.com/openconfig/ygot/ygot"
)

// Lacp represents the /openconfig-lacp/lacp YANG schema element.
type Lacp struct {
	ΛMetadata       []ygot.Annotation          `path:"@" ygotAnnotation:"true"`
	Interface       map[string]*Lacp_Interface `path:"interfaces/interface" module:"openconfig-lacp"`
	ΛInterface      []ygot.Annotation          `path:"interfaces/@interface" ygotAnnotation:"true"`
	SystemPriority  *uint16                    `path:"config/system-priority" module:"openconfig-lacp"`
	ΛSystemPriority []ygot.Annotation          `path:"config/@system-priority" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lacp implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lacp) IsYANGGoStruct() {}

// Lacp_Interface represents the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_Interface struct {
	ΛMetadata       []ygot.Annotation                 `path:"@" ygotAnnotation:"true"`
	Interval        E_OpenconfigLacp_LacpPeriodType   `path:"config/interval" module:"openconfig-lacp"`
	ΛInterval       []ygot.Annotation                 `path:"config/@interval" ygotAnnotation:"true"`
	LacpMode        E_OpenconfigLacp_LacpActivityType `path:"config/lacp-mode" module:"openconfig-lacp"`
	ΛLacpMode       []ygot.Annotation                 `path:"config/@lacp-mode" ygotAnnotation:"true"`
	Member          map[string]*Lacp_Interface_Member `path:"members/member" module:"openconfig-lacp"`
	ΛMember         []ygot.Annotation                 `path:"members/@member" ygotAnnotation:"true"`
	Name            *string                           `path:"config/name|name" module:"openconfig-lacp"`
	ΛName           []ygot.Annotation                 `path:"config/@name|@name" ygotAnnotation:"true"`
	SystemIdMac     *string                           `path:"config/system-id-mac" module:"openconfig-lacp"`
	ΛSystemIdMac    []ygot.Annotation                 `path:"config/@system-id-mac" ygotAnnotation:"true"`
	SystemPriority  *uint16                           `path:"config/system-priority" module:"openconfig-lacp"`
	ΛSystemPriority []ygot.Annotation                 `path:"config/@system-priority" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lacp_Interface implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lacp_Interface) IsYANGGoStruct() {}

// Lacp_Interface_Member represents the /openconfig-lacp/lacp/interfaces/interface/members/member YANG schema element.
type Lacp_Interface_Member struct {
	ΛMetadata        []ygot.Annotation                        `path:"@" ygotAnnotation:"true"`
	Activity         E_OpenconfigLacp_LacpActivityType        `path:"state/activity" module:"openconfig-lacp"`
	ΛActivity        []ygot.Annotation                        `path:"state/@activity" ygotAnnotation:"true"`
	Aggregatable     *bool                                    `path:"state/aggregatable" module:"openconfig-lacp"`
	ΛAggregatable    []ygot.Annotation                        `path:"state/@aggregatable" ygotAnnotation:"true"`
	Collecting       *bool                                    `path:"state/collecting" module:"openconfig-lacp"`
	ΛCollecting      []ygot.Annotation                        `path:"state/@collecting" ygotAnnotation:"true"`
	Counters         *Lacp_Interface_Member_Counters          `path:"state/counters" module:"openconfig-lacp"`
	ΛCounters        []ygot.Annotation                        `path:"state/@counters" ygotAnnotation:"true"`
	Distributing     *bool                                    `path:"state/distributing" module:"openconfig-lacp"`
	ΛDistributing    []ygot.Annotation                        `path:"state/@distributing" ygotAnnotation:"true"`
	Interface        *string                                  `path:"state/interface|interface" module:"openconfig-lacp"`
	ΛInterface       []ygot.Annotation                        `path:"state/@interface|@interface" ygotAnnotation:"true"`
	OperKey          *uint16                                  `path:"state/oper-key" module:"openconfig-lacp"`
	ΛOperKey         []ygot.Annotation                        `path:"state/@oper-key" ygotAnnotation:"true"`
	PartnerId        *string                                  `path:"state/partner-id" module:"openconfig-lacp"`
	ΛPartnerId       []ygot.Annotation                        `path:"state/@partner-id" ygotAnnotation:"true"`
	PartnerKey       *uint16                                  `path:"state/partner-key" module:"openconfig-lacp"`
	ΛPartnerKey      []ygot.Annotation                        `path:"state/@partner-key" ygotAnnotation:"true"`
	PartnerPortNum   *uint16                                  `path:"state/partner-port-num" module:"openconfig-lacp"`
	ΛPartnerPortNum  []ygot.Annotation                        `path:"state/@partner-port-num" ygotAnnotation:"true"`
	PortNum          *uint16                                  `path:"state/port-num" module:"openconfig-lacp"`
	ΛPortNum         []ygot.Annotation                        `path:"state/@port-num" ygotAnnotation:"true"`
	Synchronization  E_OpenconfigLacp_LacpSynchronizationType `path:"state/synchronization" module:"openconfig-lacp"`
	ΛSynchronization []ygot.Annotation                        `path:"state/@synchronization" ygotAnnotation:"true"`
	SystemId         *string                                  `path:"state/system-id" module:"openconfig-lacp"`
	ΛSystemId        []ygot.Annotation                        `path:"state/@system-id" ygotAnnotation:"true"`
	Timeout          E_OpenconfigLacp_LacpTimeoutType         `path:"state/timeout" module:"openconfig-lacp"`
	ΛTimeout         []ygot.Annotation                        `path:"state/@timeout" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lacp_Interface_Member implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lacp_Interface_Member) IsYANGGoStruct() {}

// Lacp_Interface_Member_Counters represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters YANG schema element.
type Lacp_Interface_Member_Counters struct {
	ΛMetadata          []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	LacpErrors         *uint64           `path:"lacp-errors" module:"openconfig-lacp"`
	ΛLacpErrors        []ygot.Annotation `path:"@lacp-errors" ygotAnnotation:"true"`
	LacpInPkts         *uint64           `path:"lacp-in-pkts" module:"openconfig-lacp"`
	ΛLacpInPkts        []ygot.Annotation `path:"@lacp-in-pkts" ygotAnnotation:"true"`
	LacpOutPkts        *uint64           `path:"lacp-out-pkts" module:"openconfig-lacp"`
	ΛLacpOutPkts       []ygot.Annotation `path:"@lacp-out-pkts" ygotAnnotation:"true"`
	LacpRxErrors       *uint64           `path:"lacp-rx-errors" module:"openconfig-lacp"`
	ΛLacpRxErrors      []ygot.Annotation `path:"@lacp-rx-errors" ygotAnnotation:"true"`
	LacpTxErrors       *uint64           `path:"lacp-tx-errors" module:"openconfig-lacp"`
	ΛLacpTxErrors      []ygot.Annotation `path:"@lacp-tx-errors" ygotAnnotation:"true"`
	LacpUnknownErrors  *uint64           `path:"lacp-unknown-errors" module:"openconfig-lacp"`
	ΛLacpUnknownErrors []ygot.Annotation `path:"@lacp-unknown-errors" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lacp_Interface_Member_Counters implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lacp_Interface_Member_Counters) IsYANGGoStruct() {}

// Lldp represents the /openconfig-lldp/lldp YANG schema element.
type Lldp struct {
	ΛMetadata                 []ygot.Annotation                `path:"@" ygotAnnotation:"true"`
	ChassisId                 *string                          `path:"config/chassis-id" module:"openconfig-lldp"`
	ΛChassisId                []ygot.Annotation                `path:"config/@chassis-id" ygotAnnotation:"true"`
	ChassisIdType             E_OpenconfigLldp_ChassisIdType   `path:"config/chassis-id-type" module:"openconfig-lldp"`
	ΛChassisIdType            []ygot.Annotation                `path:"config/@chassis-id-type" ygotAnnotation:"true"`
	Counters                  *Lldp_Counters                   `path:"state/counters" module:"openconfig-lldp"`
	ΛCounters                 []ygot.Annotation                `path:"state/@counters" ygotAnnotation:"true"`
	Enabled                   *bool                            `path:"config/enabled" module:"openconfig-lldp"`
	ΛEnabled                  []ygot.Annotation                `path:"config/@enabled" ygotAnnotation:"true"`
	HelloTimer                *uint64                          `path:"config/hello-timer" module:"openconfig-lldp"`
	ΛHelloTimer               []ygot.Annotation                `path:"config/@hello-timer" ygotAnnotation:"true"`
	Interface                 map[string]*Lldp_Interface       `path:"interfaces/interface" module:"openconfig-lldp"`
	ΛInterface                []ygot.Annotation                `path:"interfaces/@interface" ygotAnnotation:"true"`
	SuppressTlvAdvertisement  []E_OpenconfigLldpTypes_LLDP_TLV `path:"config/suppress-tlv-advertisement" module:"openconfig-lldp"`
	ΛSuppressTlvAdvertisement []ygot.Annotation                `path:"config/@suppress-tlv-advertisement" ygotAnnotation:"true"`
	SystemDescription         *string                          `path:"config/system-description" module:"openconfig-lldp"`
	ΛSystemDescription        []ygot.Annotation                `path:"config/@system-description" ygotAnnotation:"true"`
	SystemName                *string                          `path:"config/system-name" module:"openconfig-lldp"`
	ΛSystemName               []ygot.Annotation                `path:"config/@system-name" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp) IsYANGGoStruct() {}

// Lldp_Counters represents the /openconfig-lldp/lldp/state/counters YANG schema element.
type Lldp_Counters struct {
	ΛMetadata       []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	EntriesAgedOut  *uint64           `path:"entries-aged-out" module:"openconfig-lldp"`
	ΛEntriesAgedOut []ygot.Annotation `path:"@entries-aged-out" ygotAnnotation:"true"`
	FrameDiscard    *uint64           `path:"frame-discard" module:"openconfig-lldp"`
	ΛFrameDiscard   []ygot.Annotation `path:"@frame-discard" ygotAnnotation:"true"`
	FrameErrorIn    *uint64           `path:"frame-error-in" module:"openconfig-lldp"`
	ΛFrameErrorIn   []ygot.Annotation `path:"@frame-error-in" ygotAnnotation:"true"`
	FrameIn         *uint64           `path:"frame-in" module:"openconfig-lldp"`
	ΛFrameIn        []ygot.Annotation `path:"@frame-in" ygotAnnotation:"true"`
	FrameOut        *uint64           `path:"frame-out" module:"openconfig-lldp"`
	ΛFrameOut       []ygot.Annotation `path:"@frame-out" ygotAnnotation:"true"`
	LastClear       *string           `path:"last-clear" module:"openconfig-lldp"`
	ΛLastClear      []ygot.Annotation `path:"@last-clear" ygotAnnotation:"true"`
	TlvAccepted     *uint64           `path:"tlv-accepted" module:"openconfig-lldp"`
	ΛTlvAccepted    []ygot.Annotation `path:"@tlv-accepted" ygotAnnotation:"true"`
	TlvDiscard      *uint64           `path:"tlv-discard" module:"openconfig-lldp"`
	ΛTlvDiscard     []ygot.Annotation `path:"@tlv-discard" ygotAnnotation:"true"`
	TlvUnknown      *uint64           `path:"tlv-unknown" module:"openconfig-lldp"`
	ΛTlvUnknown     []ygot.Annotation `path:"@tlv-unknown" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp_Counters implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp_Counters) IsYANGGoStruct() {}

// Lldp_Interface represents the /openconfig-lldp/lldp/interfaces/interface YANG schema element.
type Lldp_Interface struct {
	ΛMetadata []ygot.Annotation                   `path:"@" ygotAnnotation:"true"`
	Counters  *Lldp_Interface_Counters            `path:"state/counters" module:"openconfig-lldp"`
	ΛCounters []ygot.Annotation                   `path:"state/@counters" ygotAnnotation:"true"`
	Enabled   *bool                               `path:"config/enabled" module:"openconfig-lldp"`
	ΛEnabled  []ygot.Annotation                   `path:"config/@enabled" ygotAnnotation:"true"`
	Name      *string                             `path:"config/name|name" module:"openconfig-lldp"`
	ΛName     []ygot.Annotation                   `path:"config/@name|@name" ygotAnnotation:"true"`
	Neighbor  map[string]*Lldp_Interface_Neighbor `path:"neighbors/neighbor" module:"openconfig-lldp"`
	ΛNeighbor []ygot.Annotation                   `path:"neighbors/@neighbor" ygotAnnotation:"true"`
	Snooping  *bool                               `path:"config/snooping" module:"openconfig-terminal-device"`
	ΛSnooping []ygot.Annotation                   `path:"config/@snooping" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp_Interface implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp_Interface) IsYANGGoStruct() {}

// Lldp_Interface_Counters represents the /openconfig-lldp/lldp/interfaces/interface/state/counters YANG schema element.
type Lldp_Interface_Counters struct {
	ΛMetadata      []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	FrameDiscard   *uint64           `path:"frame-discard" module:"openconfig-lldp"`
	ΛFrameDiscard  []ygot.Annotation `path:"@frame-discard" ygotAnnotation:"true"`
	FrameErrorIn   *uint64           `path:"frame-error-in" module:"openconfig-lldp"`
	ΛFrameErrorIn  []ygot.Annotation `path:"@frame-error-in" ygotAnnotation:"true"`
	FrameErrorOut  *uint64           `path:"frame-error-out" module:"openconfig-lldp"`
	ΛFrameErrorOut []ygot.Annotation `path:"@frame-error-out" ygotAnnotation:"true"`
	FrameIn        *uint64           `path:"frame-in" module:"openconfig-lldp"`
	ΛFrameIn       []ygot.Annotation `path:"@frame-in" ygotAnnotation:"true"`
	FrameOut       *uint64           `path:"frame-out" module:"openconfig-lldp"`
	ΛFrameOut      []ygot.Annotation `path:"@frame-out" ygotAnnotation:"true"`
	LastClear      *string           `path:"last-clear" module:"openconfig-lldp"`
	ΛLastClear     []ygot.Annotation `path:"@last-clear" ygotAnnotation:"true"`
	TlvDiscard     *uint64           `path:"tlv-discard" module:"openconfig-lldp"`
	ΛTlvDiscard    []ygot.Annotation `path:"@tlv-discard" ygotAnnotation:"true"`
	TlvUnknown     *uint64           `path:"tlv-unknown" module:"openconfig-lldp"`
	ΛTlvUnknown    []ygot.Annotation `path:"@tlv-unknown" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp_Interface_Counters implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp_Interface_Counters) IsYANGGoStruct() {}

// Lldp_Interface_Neighbor represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor YANG schema element.
type Lldp_Interface_Neighbor struct {
	ΛMetadata              []ygot.Annotation                                                                    `path:"@" ygotAnnotation:"true"`
	Age                    *uint64                                                                              `path:"state/age" module:"openconfig-lldp"`
	ΛAge                   []ygot.Annotation                                                                    `path:"state/@age" ygotAnnotation:"true"`
	Capability             map[E_OpenconfigLldpTypes_LLDP_SYSTEM_CAPABILITY]*Lldp_Interface_Neighbor_Capability `path:"capabilities/capability" module:"openconfig-lldp"`
	ΛCapability            []ygot.Annotation                                                                    `path:"capabilities/@capability" ygotAnnotation:"true"`
	ChassisId              *string                                                                              `path:"state/chassis-id" module:"openconfig-lldp"`
	ΛChassisId             []ygot.Annotation                                                                    `path:"state/@chassis-id" ygotAnnotation:"true"`
	ChassisIdType          E_OpenconfigLldp_ChassisIdType                                                       `path:"state/chassis-id-type" module:"openconfig-lldp"`
	ΛChassisIdType         []ygot.Annotation                                                                    `path:"state/@chassis-id-type" ygotAnnotation:"true"`
	Id                     *string                                                                              `path:"state/id|id" module:"openconfig-lldp"`
	ΛId                    []ygot.Annotation                                                                    `path:"state/@id|@id" ygotAnnotation:"true"`
	LastUpdate             *int64                                                                               `path:"state/last-update" module:"openconfig-lldp"`
	ΛLastUpdate            []ygot.Annotation                                                                    `path:"state/@last-update" ygotAnnotation:"true"`
	ManagementAddress      *string                                                                              `path:"state/management-address" module:"openconfig-lldp"`
	ΛManagementAddress     []ygot.Annotation                                                                    `path:"state/@management-address" ygotAnnotation:"true"`
	ManagementAddressType  *string                                                                              `path:"state/management-address-type" module:"openconfig-lldp"`
	ΛManagementAddressType []ygot.Annotation                                                                    `path:"state/@management-address-type" ygotAnnotation:"true"`
	PortDescription        *string                                                                              `path:"state/port-description" module:"openconfig-lldp"`
	ΛPortDescription       []ygot.Annotation                                                                    `path:"state/@port-description" ygotAnnotation:"true"`
	PortId                 *string                                                                              `path:"state/port-id" module:"openconfig-lldp"`
	ΛPortId                []ygot.Annotation                                                                    `path:"state/@port-id" ygotAnnotation:"true"`
	PortIdType             E_OpenconfigLldp_PortIdType                                                          `path:"state/port-id-type" module:"openconfig-lldp"`
	ΛPortIdType            []ygot.Annotation                                                                    `path:"state/@port-id-type" ygotAnnotation:"true"`
	SystemDescription      *string                                                                              `path:"state/system-description" module:"openconfig-lldp"`
	ΛSystemDescription     []ygot.Annotation                                                                    `path:"state/@system-description" ygotAnnotation:"true"`
	SystemName             *string                                                                              `path:"state/system-name" module:"openconfig-lldp"`
	ΛSystemName            []ygot.Annotation                                                                    `path:"state/@system-name" ygotAnnotation:"true"`
	Tlv                    map[Lldp_Interface_Neighbor_Tlv_Key]*Lldp_Interface_Neighbor_Tlv                     `path:"custom-tlvs/tlv" module:"openconfig-lldp"`
	ΛTlv                   []ygot.Annotation                                                                    `path:"custom-tlvs/@tlv" ygotAnnotation:"true"`
	Ttl                    *uint16                                                                              `path:"state/ttl" module:"openconfig-lldp"`
	ΛTtl                   []ygot.Annotation                                                                    `path:"state/@ttl" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp_Interface_Neighbor implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp_Interface_Neighbor) IsYANGGoStruct() {}

// Lldp_Interface_Neighbor_Tlv_Key represents the key for list Tlv of element /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor.
type Lldp_Interface_Neighbor_Tlv_Key struct {
	Type       int32  `path:"type"`
	Oui        string `path:"oui"`
	OuiSubtype string `path:"oui-subtype"`
}

// Lldp_Interface_Neighbor_Capability represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability YANG schema element.
type Lldp_Interface_Neighbor_Capability struct {
	ΛMetadata []ygot.Annotation                            `path:"@" ygotAnnotation:"true"`
	Enabled   *bool                                        `path:"state/enabled" module:"openconfig-lldp"`
	ΛEnabled  []ygot.Annotation                            `path:"state/@enabled" ygotAnnotation:"true"`
	Name      E_OpenconfigLldpTypes_LLDP_SYSTEM_CAPABILITY `path:"state/name|name" module:"openconfig-lldp"`
	ΛName     []ygot.Annotation                            `path:"state/@name|@name" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp_Interface_Neighbor_Capability implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp_Interface_Neighbor_Capability) IsYANGGoStruct() {}

// Lldp_Interface_Neighbor_Tlv represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv YANG schema element.
type Lldp_Interface_Neighbor_Tlv struct {
	ΛMetadata   []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Oui         *string           `path:"state/oui|oui" module:"openconfig-lldp"`
	ΛOui        []ygot.Annotation `path:"state/@oui|@oui" ygotAnnotation:"true"`
	OuiSubtype  *string           `path:"state/oui-subtype|oui-subtype" module:"openconfig-lldp"`
	ΛOuiSubtype []ygot.Annotation `path:"state/@oui-subtype|@oui-subtype" ygotAnnotation:"true"`
	Type        *int32            `path:"state/type|type" module:"openconfig-lldp"`
	ΛType       []ygot.Annotation `path:"state/@type|@type" ygotAnnotation:"true"`
	Value       Binary            `path:"state/value" module:"openconfig-lldp"`
	ΛValue      []ygot.Annotation `path:"state/@value" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Lldp_Interface_Neighbor_Tlv implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Lldp_Interface_Neighbor_Tlv) IsYANGGoStruct() {}

// LocalRoutes represents the /openconfig-local-routing/local-routes YANG schema element.
type LocalRoutes struct {
	ΛMetadata  []ygot.Annotation                 `path:"@" ygotAnnotation:"true"`
	Aggregate  map[string]*LocalRoutes_Aggregate `path:"local-aggregates/aggregate" module:"openconfig-local-routing"`
	ΛAggregate []ygot.Annotation                 `path:"local-aggregates/@aggregate" ygotAnnotation:"true"`
	Static     map[string]*LocalRoutes_Static    `path:"static-routes/static" module:"openconfig-local-routing"`
	ΛStatic    []ygot.Annotation                 `path:"static-routes/@static" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that LocalRoutes implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*LocalRoutes) IsYANGGoStruct() {}

// LocalRoutes_Aggregate represents the /openconfig-local-routing/local-routes/local-aggregates/aggregate YANG schema element.
type LocalRoutes_Aggregate struct {
	ΛMetadata []ygot.Annotation                  `path:"@" ygotAnnotation:"true"`
	Discard   *bool                              `path:"config/discard" module:"openconfig-local-routing"`
	ΛDiscard  []ygot.Annotation                  `path:"config/@discard" ygotAnnotation:"true"`
	Prefix    *string                            `path:"config/prefix|prefix" module:"openconfig-local-routing"`
	ΛPrefix   []ygot.Annotation                  `path:"config/@prefix|@prefix" ygotAnnotation:"true"`
	SetTag    LocalRoutes_Aggregate_SetTag_Union `path:"config/set-tag" module:"openconfig-local-routing"`
	ΛSetTag   []ygot.Annotation                  `path:"config/@set-tag" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that LocalRoutes_Aggregate implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*LocalRoutes_Aggregate) IsYANGGoStruct() {}

// LocalRoutes_Static represents the /openconfig-local-routing/local-routes/static-routes/static YANG schema element.
type LocalRoutes_Static struct {
	ΛMetadata []ygot.Annotation                      `path:"@" ygotAnnotation:"true"`
	NextHop   map[string]*LocalRoutes_Static_NextHop `path:"next-hops/next-hop" module:"openconfig-local-routing"`
	ΛNextHop  []ygot.Annotation                      `path:"next-hops/@next-hop" ygotAnnotation:"true"`
	Prefix    *string                                `path:"config/prefix|prefix" module:"openconfig-local-routing"`
	ΛPrefix   []ygot.Annotation                      `path:"config/@prefix|@prefix" ygotAnnotation:"true"`
	SetTag    LocalRoutes_Static_SetTag_Union        `path:"config/set-tag" module:"openconfig-local-routing"`
	ΛSetTag   []ygot.Annotation                      `path:"config/@set-tag" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that LocalRoutes_Static implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*LocalRoutes_Static) IsYANGGoStruct() {}

// LocalRoutes_Static_NextHop represents the /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop YANG schema element.
type LocalRoutes_Static_NextHop struct {
	ΛMetadata     []ygot.Annotation                        `path:"@" ygotAnnotation:"true"`
	Index         *string                                  `path:"config/index|index" module:"openconfig-local-routing"`
	ΛIndex        []ygot.Annotation                        `path:"config/@index|@index" ygotAnnotation:"true"`
	InterfaceRef  *LocalRoutes_Static_NextHop_InterfaceRef `path:"interface-ref" module:"openconfig-local-routing"`
	ΛInterfaceRef []ygot.Annotation                        `path:"@interface-ref" ygotAnnotation:"true"`
	Metric        *uint32                                  `path:"config/metric" module:"openconfig-local-routing"`
	ΛMetric       []ygot.Annotation                        `path:"config/@metric" ygotAnnotation:"true"`
	NextHop       LocalRoutes_Static_NextHop_NextHop_Union `path:"config/next-hop" module:"openconfig-local-routing"`
	ΛNextHop      []ygot.Annotation                        `path:"config/@next-hop" ygotAnnotation:"true"`
	Recurse       *bool                                    `path:"config/recurse" module:"openconfig-local-routing"`
	ΛRecurse      []ygot.Annotation                        `path:"config/@recurse" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that LocalRoutes_Static_NextHop implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*LocalRoutes_Static_NextHop) IsYANGGoStruct() {}

// LocalRoutes_Static_NextHop_InterfaceRef represents the /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref YANG schema element.
type LocalRoutes_Static_NextHop_InterfaceRef struct {
	ΛMetadata     []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Interface     *string           `path:"config/interface" module:"openconfig-local-routing"`
	ΛInterface    []ygot.Annotation `path:"config/@interface" ygotAnnotation:"true"`
	Subinterface  *uint32           `path:"config/subinterface" module:"openconfig-local-routing"`
	ΛSubinterface []ygot.Annotation `path:"config/@subinterface" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that LocalRoutes_Static_NextHop_InterfaceRef implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*LocalRoutes_Static_NextHop_InterfaceRef) IsYANGGoStruct() {}

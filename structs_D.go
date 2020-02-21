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

// Device represents the /device YANG schema element.
type Device struct {
	ΛMetadata         []ygot.Annotation           `path:"@" ygotAnnotation:"true"`
	Acl               *Acl                        `path:"acl" module:"openconfig-acl"`
	ΛAcl              []ygot.Annotation           `path:"@acl" ygotAnnotation:"true"`
	Aps               *Aps                        `path:"aps" module:"openconfig-transport-line-protection"`
	ΛAps              []ygot.Annotation           `path:"@aps" ygotAnnotation:"true"`
	Bgp               *Bgp                        `path:"bgp" module:"openconfig-bgp"`
	ΛBgp              []ygot.Annotation           `path:"@bgp" ygotAnnotation:"true"`
	Component         map[string]*Component       `path:"components/component" module:"openconfig-platform"`
	ΛComponent        []ygot.Annotation           `path:"components/@component" ygotAnnotation:"true"`
	Interface         map[string]*Interface       `path:"interfaces/interface" module:"openconfig-interfaces"`
	ΛInterface        []ygot.Annotation           `path:"interfaces/@interface" ygotAnnotation:"true"`
	Lacp              *Lacp                       `path:"lacp" module:"openconfig-lacp"`
	ΛLacp             []ygot.Annotation           `path:"@lacp" ygotAnnotation:"true"`
	Lldp              *Lldp                       `path:"lldp" module:"openconfig-lldp"`
	ΛLldp             []ygot.Annotation           `path:"@lldp" ygotAnnotation:"true"`
	LocalRoutes       *LocalRoutes                `path:"local-routes" module:"openconfig-local-routing"`
	ΛLocalRoutes      []ygot.Annotation           `path:"@local-routes" ygotAnnotation:"true"`
	NetworkInstance   map[string]*NetworkInstance `path:"network-instances/network-instance" module:"openconfig-network-instance"`
	ΛNetworkInstance  []ygot.Annotation           `path:"network-instances/@network-instance" ygotAnnotation:"true"`
	OpticalAmplifier  *OpticalAmplifier           `path:"optical-amplifier" module:"openconfig-optical-amplifier"`
	ΛOpticalAmplifier []ygot.Annotation           `path:"@optical-amplifier" ygotAnnotation:"true"`
	RoutingPolicy     *RoutingPolicy              `path:"routing-policy" module:"openconfig-routing-policy"`
	ΛRoutingPolicy    []ygot.Annotation           `path:"@routing-policy" ygotAnnotation:"true"`
	Stp               *Stp                        `path:"stp" module:"openconfig-spanning-tree"`
	ΛStp              []ygot.Annotation           `path:"@stp" ygotAnnotation:"true"`
	System            *System                     `path:"system" module:"openconfig-system"`
	ΛSystem           []ygot.Annotation           `path:"@system" ygotAnnotation:"true"`
	TerminalDevice    *TerminalDevice             `path:"terminal-device" module:"openconfig-terminal-device"`
	ΛTerminalDevice   []ygot.Annotation           `path:"@terminal-device" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

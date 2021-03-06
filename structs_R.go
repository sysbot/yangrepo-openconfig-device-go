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

// RoutingPolicy represents the /openconfig-routing-policy/routing-policy YANG schema element.
type RoutingPolicy struct {
	ΛMetadata         []ygot.Annotation                          `path:"@" ygotAnnotation:"true"`
	DefinedSets       *RoutingPolicy_DefinedSets                 `path:"defined-sets" module:"openconfig-routing-policy"`
	ΛDefinedSets      []ygot.Annotation                          `path:"@defined-sets" ygotAnnotation:"true"`
	PolicyDefinition  map[string]*RoutingPolicy_PolicyDefinition `path:"policy-definitions/policy-definition" module:"openconfig-routing-policy"`
	ΛPolicyDefinition []ygot.Annotation                          `path:"policy-definitions/@policy-definition" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy) IsYANGGoStruct() {}

// RoutingPolicy_DefinedSets represents the /openconfig-routing-policy/routing-policy/defined-sets YANG schema element.
type RoutingPolicy_DefinedSets struct {
	ΛMetadata    []ygot.Annotation                                 `path:"@" ygotAnnotation:"true"`
	NeighborSet  map[string]*RoutingPolicy_DefinedSets_NeighborSet `path:"neighbor-sets/neighbor-set" module:"openconfig-routing-policy"`
	ΛNeighborSet []ygot.Annotation                                 `path:"neighbor-sets/@neighbor-set" ygotAnnotation:"true"`
	PrefixSet    map[string]*RoutingPolicy_DefinedSets_PrefixSet   `path:"prefix-sets/prefix-set" module:"openconfig-routing-policy"`
	ΛPrefixSet   []ygot.Annotation                                 `path:"prefix-sets/@prefix-set" ygotAnnotation:"true"`
	TagSet       map[string]*RoutingPolicy_DefinedSets_TagSet      `path:"tag-sets/tag-set" module:"openconfig-routing-policy"`
	ΛTagSet      []ygot.Annotation                                 `path:"tag-sets/@tag-set" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_DefinedSets implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_DefinedSets) IsYANGGoStruct() {}

// RoutingPolicy_DefinedSets_NeighborSet represents the /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set YANG schema element.
type RoutingPolicy_DefinedSets_NeighborSet struct {
	ΛMetadata []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Address   []string          `path:"config/address" module:"openconfig-routing-policy"`
	ΛAddress  []ygot.Annotation `path:"config/@address" ygotAnnotation:"true"`
	Name      *string           `path:"config/name|name" module:"openconfig-routing-policy"`
	ΛName     []ygot.Annotation `path:"config/@name|@name" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_DefinedSets_NeighborSet implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_DefinedSets_NeighborSet) IsYANGGoStruct() {}

// RoutingPolicy_DefinedSets_PrefixSet represents the /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set YANG schema element.
type RoutingPolicy_DefinedSets_PrefixSet struct {
	ΛMetadata []ygot.Annotation                                                                              `path:"@" ygotAnnotation:"true"`
	Mode      E_OpenconfigRoutingPolicy_PrefixSet_Mode                                                       `path:"config/mode" module:"openconfig-routing-policy"`
	ΛMode     []ygot.Annotation                                                                              `path:"config/@mode" ygotAnnotation:"true"`
	Name      *string                                                                                        `path:"config/name|name" module:"openconfig-routing-policy"`
	ΛName     []ygot.Annotation                                                                              `path:"config/@name|@name" ygotAnnotation:"true"`
	Prefix    map[RoutingPolicy_DefinedSets_PrefixSet_Prefix_Key]*RoutingPolicy_DefinedSets_PrefixSet_Prefix `path:"prefixes/prefix" module:"openconfig-routing-policy"`
	ΛPrefix   []ygot.Annotation                                                                              `path:"prefixes/@prefix" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_DefinedSets_PrefixSet implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_DefinedSets_PrefixSet) IsYANGGoStruct() {}

// RoutingPolicy_DefinedSets_PrefixSet_Prefix_Key represents the key for list Prefix of element /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set.
type RoutingPolicy_DefinedSets_PrefixSet_Prefix_Key struct {
	IpPrefix        string `path:"ip-prefix"`
	MasklengthRange string `path:"masklength-range"`
}

// RoutingPolicy_DefinedSets_PrefixSet_Prefix represents the /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix YANG schema element.
type RoutingPolicy_DefinedSets_PrefixSet_Prefix struct {
	ΛMetadata        []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	IpPrefix         *string           `path:"config/ip-prefix|ip-prefix" module:"openconfig-routing-policy"`
	ΛIpPrefix        []ygot.Annotation `path:"config/@ip-prefix|@ip-prefix" ygotAnnotation:"true"`
	MasklengthRange  *string           `path:"config/masklength-range|masklength-range" module:"openconfig-routing-policy"`
	ΛMasklengthRange []ygot.Annotation `path:"config/@masklength-range|@masklength-range" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_DefinedSets_PrefixSet_Prefix implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_DefinedSets_PrefixSet_Prefix) IsYANGGoStruct() {}

// RoutingPolicy_DefinedSets_TagSet represents the /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set YANG schema element.
type RoutingPolicy_DefinedSets_TagSet struct {
	ΛMetadata []ygot.Annotation                                 `path:"@" ygotAnnotation:"true"`
	Name      *string                                           `path:"config/name|name" module:"openconfig-routing-policy"`
	ΛName     []ygot.Annotation                                 `path:"config/@name|@name" ygotAnnotation:"true"`
	TagValue  []RoutingPolicy_DefinedSets_TagSet_TagValue_Union `path:"config/tag-value" module:"openconfig-routing-policy"`
	ΛTagValue []ygot.Annotation                                 `path:"config/@tag-value" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_DefinedSets_TagSet implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_DefinedSets_TagSet) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition YANG schema element.
type RoutingPolicy_PolicyDefinition struct {
	ΛMetadata  []ygot.Annotation                                    `path:"@" ygotAnnotation:"true"`
	Name       *string                                              `path:"config/name|name" module:"openconfig-routing-policy"`
	ΛName      []ygot.Annotation                                    `path:"config/@name|@name" ygotAnnotation:"true"`
	Statement  map[string]*RoutingPolicy_PolicyDefinition_Statement `path:"statements/statement" module:"openconfig-routing-policy"`
	ΛStatement []ygot.Annotation                                    `path:"statements/@statement" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement struct {
	ΛMetadata   []ygot.Annotation                                    `path:"@" ygotAnnotation:"true"`
	Actions     *RoutingPolicy_PolicyDefinition_Statement_Actions    `path:"actions" module:"openconfig-routing-policy"`
	ΛActions    []ygot.Annotation                                    `path:"@actions" ygotAnnotation:"true"`
	Conditions  *RoutingPolicy_PolicyDefinition_Statement_Conditions `path:"conditions" module:"openconfig-routing-policy"`
	ΛConditions []ygot.Annotation                                    `path:"@conditions" ygotAnnotation:"true"`
	Name        *string                                              `path:"config/name|name" module:"openconfig-routing-policy"`
	ΛName       []ygot.Annotation                                    `path:"config/@name|@name" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement_Actions represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement_Actions struct {
	ΛMetadata     []ygot.Annotation                          `path:"@" ygotAnnotation:"true"`
	PolicyResult  E_OpenconfigRoutingPolicy_PolicyResultType `path:"config/policy-result" module:"openconfig-routing-policy"`
	ΛPolicyResult []ygot.Annotation                          `path:"config/@policy-result" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement_Actions implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement_Actions) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement_Conditions represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement_Conditions struct {
	ΛMetadata          []ygot.Annotation                                                     `path:"@" ygotAnnotation:"true"`
	CallPolicy         *string                                                               `path:"config/call-policy" module:"openconfig-routing-policy"`
	ΛCallPolicy        []ygot.Annotation                                                     `path:"config/@call-policy" ygotAnnotation:"true"`
	InstallProtocolEq  E_OpenconfigPolicyTypes_INSTALL_PROTOCOL_TYPE                         `path:"config/install-protocol-eq" module:"openconfig-routing-policy"`
	ΛInstallProtocolEq []ygot.Annotation                                                     `path:"config/@install-protocol-eq" ygotAnnotation:"true"`
	MatchInterface     *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface   `path:"match-interface" module:"openconfig-routing-policy"`
	ΛMatchInterface    []ygot.Annotation                                                     `path:"@match-interface" ygotAnnotation:"true"`
	MatchNeighborSet   *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet `path:"match-neighbor-set" module:"openconfig-routing-policy"`
	ΛMatchNeighborSet  []ygot.Annotation                                                     `path:"@match-neighbor-set" ygotAnnotation:"true"`
	MatchPrefixSet     *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet   `path:"match-prefix-set" module:"openconfig-routing-policy"`
	ΛMatchPrefixSet    []ygot.Annotation                                                     `path:"@match-prefix-set" ygotAnnotation:"true"`
	MatchTagSet        *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet      `path:"match-tag-set" module:"openconfig-routing-policy"`
	ΛMatchTagSet       []ygot.Annotation                                                     `path:"@match-tag-set" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement_Conditions implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement_Conditions) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface struct {
	ΛMetadata     []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Interface     *string           `path:"config/interface" module:"openconfig-routing-policy"`
	ΛInterface    []ygot.Annotation `path:"config/@interface" ygotAnnotation:"true"`
	Subinterface  *uint32           `path:"config/subinterface" module:"openconfig-routing-policy"`
	ΛSubinterface []ygot.Annotation `path:"config/@subinterface" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet struct {
	ΛMetadata        []ygot.Annotation                                       `path:"@" ygotAnnotation:"true"`
	MatchSetOptions  E_OpenconfigRoutingPolicy_MatchSetOptionsRestrictedType `path:"config/match-set-options" module:"openconfig-routing-policy"`
	ΛMatchSetOptions []ygot.Annotation                                       `path:"config/@match-set-options" ygotAnnotation:"true"`
	NeighborSet      *string                                                 `path:"config/neighbor-set" module:"openconfig-routing-policy"`
	ΛNeighborSet     []ygot.Annotation                                       `path:"config/@neighbor-set" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet struct {
	ΛMetadata        []ygot.Annotation                                       `path:"@" ygotAnnotation:"true"`
	MatchSetOptions  E_OpenconfigRoutingPolicy_MatchSetOptionsRestrictedType `path:"config/match-set-options" module:"openconfig-routing-policy"`
	ΛMatchSetOptions []ygot.Annotation                                       `path:"config/@match-set-options" ygotAnnotation:"true"`
	PrefixSet        *string                                                 `path:"config/prefix-set" module:"openconfig-routing-policy"`
	ΛPrefixSet       []ygot.Annotation                                       `path:"config/@prefix-set" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) IsYANGGoStruct() {}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet represents the /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set YANG schema element.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet struct {
	ΛMetadata        []ygot.Annotation                                       `path:"@" ygotAnnotation:"true"`
	MatchSetOptions  E_OpenconfigRoutingPolicy_MatchSetOptionsRestrictedType `path:"config/match-set-options" module:"openconfig-routing-policy"`
	ΛMatchSetOptions []ygot.Annotation                                       `path:"config/@match-set-options" ygotAnnotation:"true"`
	TagSet           *string                                                 `path:"config/tag-set" module:"openconfig-routing-policy"`
	ΛTagSet          []ygot.Annotation                                       `path:"config/@tag-set" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) IsYANGGoStruct() {}

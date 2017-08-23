// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.
package interfaces

import "reflect"

var Types = map[string]reflect.Type{
	"SwInterfaceAddDelAddress":      reflect.TypeOf((*SwInterfaceAddDelAddress)(nil)).Elem(),
	"SwInterfaceAddDelAddressReply": reflect.TypeOf((*SwInterfaceAddDelAddressReply)(nil)).Elem(),
	"SwInterfaceClearStats":         reflect.TypeOf((*SwInterfaceClearStats)(nil)).Elem(),
	"SwInterfaceClearStatsReply":    reflect.TypeOf((*SwInterfaceClearStatsReply)(nil)).Elem(),
	"SwInterfaceDetails":            reflect.TypeOf((*SwInterfaceDetails)(nil)).Elem(),
	"SwInterfaceDump":               reflect.TypeOf((*SwInterfaceDump)(nil)).Elem(),
	"SwInterfaceGetTable":           reflect.TypeOf((*SwInterfaceGetTable)(nil)).Elem(),
	"SwInterfaceGetTableReply":      reflect.TypeOf((*SwInterfaceGetTableReply)(nil)).Elem(),
	"SwInterfaceSetFlags":           reflect.TypeOf((*SwInterfaceSetFlags)(nil)).Elem(),
	"SwInterfaceSetFlagsReply":      reflect.TypeOf((*SwInterfaceSetFlagsReply)(nil)).Elem(),
	"SwInterfaceSetMacAddress":      reflect.TypeOf((*SwInterfaceSetMacAddress)(nil)).Elem(),
	"SwInterfaceSetMacAddressReply": reflect.TypeOf((*SwInterfaceSetMacAddressReply)(nil)).Elem(),
	"SwInterfaceSetMtu":             reflect.TypeOf((*SwInterfaceSetMtu)(nil)).Elem(),
	"SwInterfaceSetMtuReply":        reflect.TypeOf((*SwInterfaceSetMtuReply)(nil)).Elem(),
	"SwInterfaceSetTable":           reflect.TypeOf((*SwInterfaceSetTable)(nil)).Elem(),
	"SwInterfaceSetTableReply":      reflect.TypeOf((*SwInterfaceSetTableReply)(nil)).Elem(),
	"SwInterfaceSetUnnumbered":      reflect.TypeOf((*SwInterfaceSetUnnumbered)(nil)).Elem(),
	"SwInterfaceSetUnnumberedReply": reflect.TypeOf((*SwInterfaceSetUnnumberedReply)(nil)).Elem(),
	"SwInterfaceTagAddDel":          reflect.TypeOf((*SwInterfaceTagAddDel)(nil)).Elem(),
	"SwInterfaceTagAddDelReply":     reflect.TypeOf((*SwInterfaceTagAddDelReply)(nil)).Elem(),
	"VlibCounter":                   reflect.TypeOf((*VlibCounter)(nil)).Elem(),
	"VnetInterfaceCombinedCounters": reflect.TypeOf((*VnetInterfaceCombinedCounters)(nil)).Elem(),
	"VnetInterfaceSimpleCounters":   reflect.TypeOf((*VnetInterfaceSimpleCounters)(nil)).Elem(),
	"WantInterfaceEvents":           reflect.TypeOf((*WantInterfaceEvents)(nil)).Elem(),
	"WantInterfaceEventsReply":      reflect.TypeOf((*WantInterfaceEventsReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewSwInterfaceAddDelAddress":      reflect.ValueOf(NewSwInterfaceAddDelAddress),
	"NewSwInterfaceAddDelAddressReply": reflect.ValueOf(NewSwInterfaceAddDelAddressReply),
	"NewSwInterfaceClearStats":         reflect.ValueOf(NewSwInterfaceClearStats),
	"NewSwInterfaceClearStatsReply":    reflect.ValueOf(NewSwInterfaceClearStatsReply),
	"NewSwInterfaceDetails":            reflect.ValueOf(NewSwInterfaceDetails),
	"NewSwInterfaceDump":               reflect.ValueOf(NewSwInterfaceDump),
	"NewSwInterfaceGetTable":           reflect.ValueOf(NewSwInterfaceGetTable),
	"NewSwInterfaceGetTableReply":      reflect.ValueOf(NewSwInterfaceGetTableReply),
	"NewSwInterfaceSetFlags":           reflect.ValueOf(NewSwInterfaceSetFlags),
	"NewSwInterfaceSetFlagsReply":      reflect.ValueOf(NewSwInterfaceSetFlagsReply),
	"NewSwInterfaceSetMacAddress":      reflect.ValueOf(NewSwInterfaceSetMacAddress),
	"NewSwInterfaceSetMacAddressReply": reflect.ValueOf(NewSwInterfaceSetMacAddressReply),
	"NewSwInterfaceSetMtu":             reflect.ValueOf(NewSwInterfaceSetMtu),
	"NewSwInterfaceSetMtuReply":        reflect.ValueOf(NewSwInterfaceSetMtuReply),
	"NewSwInterfaceSetTable":           reflect.ValueOf(NewSwInterfaceSetTable),
	"NewSwInterfaceSetTableReply":      reflect.ValueOf(NewSwInterfaceSetTableReply),
	"NewSwInterfaceSetUnnumbered":      reflect.ValueOf(NewSwInterfaceSetUnnumbered),
	"NewSwInterfaceSetUnnumberedReply": reflect.ValueOf(NewSwInterfaceSetUnnumberedReply),
	"NewSwInterfaceTagAddDel":          reflect.ValueOf(NewSwInterfaceTagAddDel),
	"NewSwInterfaceTagAddDelReply":     reflect.ValueOf(NewSwInterfaceTagAddDelReply),
	"NewVnetInterfaceCombinedCounters": reflect.ValueOf(NewVnetInterfaceCombinedCounters),
	"NewVnetInterfaceSimpleCounters":   reflect.ValueOf(NewVnetInterfaceSimpleCounters),
	"NewWantInterfaceEvents":           reflect.ValueOf(NewWantInterfaceEvents),
	"NewWantInterfaceEventsReply":      reflect.ValueOf(NewWantInterfaceEventsReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{
	"VlAPIVersion": reflect.ValueOf(VlAPIVersion),
}

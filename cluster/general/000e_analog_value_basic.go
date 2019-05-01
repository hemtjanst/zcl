// An interface for setting an analog value, typically used as a control system parameter, and accessing various characteristics of that value.
package general

import (
	"neotor.se/zcl"
)

// AnalogValueBasic
const AnalogValueBasicID zcl.ClusterID = 14

var AnalogValueBasicCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

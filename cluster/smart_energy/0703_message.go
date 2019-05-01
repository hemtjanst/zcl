// This cluster provides an interface for passing text messages between ZigBee devices.
package smart_energy

import (
	"neotor.se/zcl"
)

// Message
const MessageID zcl.ClusterID = 1795

var MessageCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// Attributes and commands that provide an interface to the power profile of a device
package general

import (
	"neotor.se/zcl"
)

// PowerProfile
const PowerProfileID zcl.ClusterID = 26

var PowerProfileCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// dfdf
package hvac

import (
	"hemtjan.st/zcl"
)

// DehumidificationControl
const DehumidificationControlID zcl.ClusterID = 515

var DehumidificationControlCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

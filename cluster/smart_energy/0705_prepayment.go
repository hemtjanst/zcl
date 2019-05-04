package smart_energy

import (
	"hemtjan.st/zcl"
)

// Prepayment
const PrepaymentID zcl.ClusterID = 1797

var PrepaymentCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

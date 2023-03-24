package closures

import "hemtjan.st/zcl"

// DoorLock
const DoorLockID zcl.ClusterID = 257

var DoorLockCluster = zcl.Cluster{
	Name:       "Door Lock",
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

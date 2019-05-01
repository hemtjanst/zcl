// This cluster provides an interface for passing text messages between ZigBee devices.
package smart_energy

import (
	"neotor.se/zcl/cluster/zcl"
)

// Message
// This cluster provides an interface for passing text messages between ZigBee devices.

func NewMessageServer(profile zcl.ProfileID) *MessageServer { return &MessageServer{p: profile} }
func NewMessageClient(profile zcl.ProfileID) *MessageClient { return &MessageClient{p: profile} }

const MessageCluster zcl.ClusterID = 1795

type MessageServer struct {
	p zcl.ProfileID
}

type MessageClient struct {
	p zcl.ProfileID
}

/*
var MessageServer = map[zcl.CommandID]func() zcl.Command{
}

var MessageClient = map[zcl.CommandID]func() zcl.Command{
}
*/

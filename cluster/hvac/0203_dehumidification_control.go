// dfdf
package hvac

import (
	"neotor.se/zcl"
)

// DehumidificationControl
// dfdf

func NewDehumidificationControlServer(profile zcl.ProfileID) *DehumidificationControlServer {
	return &DehumidificationControlServer{p: profile}
}
func NewDehumidificationControlClient(profile zcl.ProfileID) *DehumidificationControlClient {
	return &DehumidificationControlClient{p: profile}
}

const DehumidificationControlCluster zcl.ClusterID = 515

type DehumidificationControlServer struct {
	p zcl.ProfileID
}

type DehumidificationControlClient struct {
	p zcl.ProfileID
}

/*
var DehumidificationControlServer = map[zcl.CommandID]func() zcl.Command{
}

var DehumidificationControlClient = map[zcl.CommandID]func() zcl.Command{
}
*/

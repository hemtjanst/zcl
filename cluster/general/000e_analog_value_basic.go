// An interface for setting an analog value, typically used as a control system parameter, and accessing various characteristics of that value.
package general

import (
	"neotor.se/zcl/cluster/zcl"
)

// AnalogValueBasic
// An interface for setting an analog value, typically used as a control system parameter, and accessing various characteristics of that value.

func NewAnalogValueBasicServer(profile zcl.ProfileID) *AnalogValueBasicServer {
	return &AnalogValueBasicServer{p: profile}
}
func NewAnalogValueBasicClient(profile zcl.ProfileID) *AnalogValueBasicClient {
	return &AnalogValueBasicClient{p: profile}
}

const AnalogValueBasicCluster zcl.ClusterID = 14

type AnalogValueBasicServer struct {
	p zcl.ProfileID
}

type AnalogValueBasicClient struct {
	p zcl.ProfileID
}

/*
var AnalogValueBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var AnalogValueBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/

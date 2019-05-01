package smart_energy

import (
	"neotor.se/zcl"
)

// Prepayment

func NewPrepaymentServer(profile zcl.ProfileID) *PrepaymentServer {
	return &PrepaymentServer{p: profile}
}
func NewPrepaymentClient(profile zcl.ProfileID) *PrepaymentClient {
	return &PrepaymentClient{p: profile}
}

const PrepaymentCluster zcl.ClusterID = 1797

type PrepaymentServer struct {
	p zcl.ProfileID
}

type PrepaymentClient struct {
	p zcl.ProfileID
}

/*
var PrepaymentServer = map[zcl.CommandID]func() zcl.Command{
}

var PrepaymentClient = map[zcl.CommandID]func() zcl.Command{
}
*/

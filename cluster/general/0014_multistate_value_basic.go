package general

import (
	"neotor.se/zcl"
)

// MultistateValueBasic

func NewMultistateValueBasicServer(profile zcl.ProfileID) *MultistateValueBasicServer {
	return &MultistateValueBasicServer{p: profile}
}
func NewMultistateValueBasicClient(profile zcl.ProfileID) *MultistateValueBasicClient {
	return &MultistateValueBasicClient{p: profile}
}

const MultistateValueBasicCluster zcl.ClusterID = 20

type MultistateValueBasicServer struct {
	p zcl.ProfileID
}

type MultistateValueBasicClient struct {
	p zcl.ProfileID
}

/*
var MultistateValueBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var MultistateValueBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/

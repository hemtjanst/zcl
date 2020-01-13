package zdo_old

import (
	"errors"
	"hemtjan.st/zcl"
)

type Command interface {
	Cluster() uint16
	Values() []zcl.Val
}

func addCommand(fn func() Command) {
	s := fn()
	typeMap[s.Cluster()] = fn
}

var typeMap = map[uint16]func() Command{}

func Find(cluster uint16) (Command, bool) {
	c, ok := typeMap[cluster]
	if ok && c != nil {
		return c(), ok
	}
	return nil, ok
}

func Unmarshal(cluster uint16, data []byte) (uint8, Command, error) {
	if len(data) < 1 {
		return 0, nil, errors.New("too few bytes")
	}
	seqNo := uint8(data[0])
	data = data[1:]

	cmd, ok := Find(cluster)
	if !ok || cmd == nil {
		return seqNo, cmd, errors.New("command not found")
	}

	var err error

	for _, v := range cmd.Values() {
		if data, err = v.UnmarshalZcl(data); err != nil {
			return seqNo, cmd, err
		}
		switch v.(type) {
		case *ResponseStatus:
			st := v.(*ResponseStatus)
			if *st != Success {
				// Stop parsing if command was not successful
				return seqNo, cmd, nil
			}
		default:
		}
	}
	return seqNo, cmd, nil
}

func init() {
	addCommand(func() Command { return &NWKAddrRequest{} })                //done
	addCommand(func() Command { return &IEEEAddrRequest{} })               //done
	addCommand(func() Command { return &NodeDescRequest{} })               //done
	addCommand(func() Command { return &PowerDescRequest{} })              //done
	addCommand(func() Command { return &SimpleDescRequest{} })             //done
	addCommand(func() Command { return &ActiveEndpointRequest{} })         //done
	addCommand(func() Command { return &MatchDescRequest{} })              //done
	addCommand(func() Command { return &UserDescRequest{} })               //done
	addCommand(func() Command { return &DiscoveryCacheRequest{} })         //done
	addCommand(func() Command { return &DeviceAnnounce{} })                //done
	addCommand(func() Command { return &SetUserDescRequest{} })            //done
	addCommand(func() Command { return &SystemServerDiscoveryRequest{} })  //done
	addCommand(func() Command { return &DiscoveryStoreRequest{} })         //done
	addCommand(func() Command { return &NodeDescStoreRequest{} })          //done
	addCommand(func() Command { return &ActiveEndpointStoreRequest{} })    //done
	addCommand(func() Command { return &SimpleDescStoreRequest{} })        //done
	addCommand(func() Command { return &RemoveNodeCacheRequest{} })        //done
	addCommand(func() Command { return &FindNodeCacheRequest{} })          //done
	addCommand(func() Command { return &ExtendedSimpleDescRequest{} })     //done
	addCommand(func() Command { return &ExtendedActiveEndpointRequest{} }) //done
	addCommand(func() Command { return &EndDeviceBindRequest{} })          //done
	addCommand(func() Command { return &BindRequest{} })
	addCommand(func() Command { return &UnbindRequest{} })
	addCommand(func() Command { return &MgmtLqiRequest{} })
	addCommand(func() Command { return &MgmtRtgRequest{} })
	addCommand(func() Command { return &MgmtLeaveRequest{} })
	addCommand(func() Command { return &MgmtPermitJoiningRequest{} })
	addCommand(func() Command { return &NWKAddrResponse{} })        //done
	addCommand(func() Command { return &IEEEAddrResponse{} })       //done
	addCommand(func() Command { return &NodeDescResponse{} })       //done
	addCommand(func() Command { return &PowerDescResponse{} })      //done
	addCommand(func() Command { return &SimpleDescResponse{} })     //done
	addCommand(func() Command { return &ActiveEndpointResponse{} }) //done
	addCommand(func() Command { return &MatchDescResponse{} })      //done
	addCommand(func() Command { return &UserDescResponse{} })
	addCommand(func() Command { return &DiscoveryCacheResponse{} })
	addCommand(func() Command { return &SetUserDescResponse{} })
	addCommand(func() Command { return &SystemServerDiscoveryResponse{} })
	addCommand(func() Command { return &DiscoveryStoreResponse{} })
	addCommand(func() Command { return &NodeDescStoreResponse{} })
	addCommand(func() Command { return &PowerDescStoreResponse{} })
	addCommand(func() Command { return &ActiveEndpointStoreResponse{} })
	addCommand(func() Command { return &SimpleDescStoreResponse{} })
	addCommand(func() Command { return &RemoveNodeCacheResponse{} })
	addCommand(func() Command { return &FindNodeCacheResponse{} })
	addCommand(func() Command { return &ExtendedSimpleDescResponse{} })
	addCommand(func() Command { return &ExtendedActiveEndpointResponse{} })
	addCommand(func() Command { return &EndDeviceBindResponse{} })
	addCommand(func() Command { return &BindResponse{} })
	addCommand(func() Command { return &UnbindResponse{} })
	addCommand(func() Command { return &MgmtLqiResponse{} })
	addCommand(func() Command { return &MgmtRtgResponse{} })
	addCommand(func() Command { return &MgmtLeaveResponse{} })
	addCommand(func() Command { return &MgmtPermitJoiningResponse{} })

	addCommand(func() Command { return &ComplexDescRequest{} })
	addCommand(func() Command { return &ComplexDescResponse{} })
}

package device

import (
	"errors"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/utils"
	"hemtjan.st/zcl/zdo"

	"log"
)

type Zdo struct {
	dev *Device
}

func (z *Zdo) Request(cmd zcl.ZdoCommand) (interface{}, error) {
	var reqBytes []byte
	for _, v := range cmd.Values() {
		tmp, err := v.MarshalZcl()
		if err != nil {
			return nil, err
		}
		reqBytes = append(reqBytes, tmp...)
	}

	ch := make(chan interface{})
	seqNo, release := z.dev.seq.Next(ch)
	defer release()

	target := utils.NWKAddress(z.dev.nwk)

	switch cmd.ID() {
	case zdo.SystemServerDiscoverRequestCommand, zdo.DeviceAnnounceCommand:
		target = utils.NWKAddress(zcl.BroadcastRxOnWhenIdle)
	}

	packet := utils.Packet(
		0,
		target,
		0,
		0,
		cmd.Cluster(),
		append([]byte{seqNo}, reqBytes...),
	)

	log.Printf("TX: [ZDO -> %s] Cmd[%d] Seq[%d] %#v (%X)",
		packet.DstAddr(),
		packet.Cluster(),
		seqNo,
		cmd,
		reqBytes)

	err := z.dev.tr.Send(packet)

	if err != nil {
		return nil, err
	}

	resp := <-ch
	return resp, nil
}

func (z *Zdo) Respond(seqNo uint8, cmd zcl.ZdoCommand) error {
	var reqBytes []byte
	for _, v := range cmd.Values() {
		tmp, err := v.MarshalZcl()
		if err != nil {
			return err
		}
		reqBytes = append(reqBytes, tmp...)
	}

	packet := utils.Packet(
		0,
		utils.NWKAddress(z.dev.nwk),
		0,
		0,
		cmd.Cluster(),
		append([]byte{seqNo}, reqBytes...),
	)

	log.Printf("TX: [ZDO -> %s] Cmd[%d] Seq[%d] %#v (%X)",
		packet.DstAddr(),
		packet.Cluster(),
		seqNo,
		cmd,
		reqBytes)

	return z.dev.tr.Send(packet)
}

func (z *Zdo) Handle(seqNo uint8, cmd zcl.ZdoCommand) error {
	if cmd.Cluster()&0x8000 == 0x8000 {
		if z.dev.seq.Handle(seqNo, cmd) {
			return nil
		}
		return errors.New("response to unknown request")
	}
	switch cmd.(type) {
	case *zdo.DeviceAnnounce: //*zdo_old.DeviceAnnounce:
		err := z.dev.Init()
		if err != nil {
			log.Printf("Error re-initializing device: %s", err)
		}
		return nil
	case *zdo.NwkAddressRequest: //*zdo_old.NWKAddrRequest:
	case *zdo.IeeeAddressRequest: //*zdo_old.IEEEAddrRequest:
	case *zdo.MatchDescRequest: //*zdo_old.MatchDescRequest:
		md := cmd.(*zdo.MatchDescRequest)
		rsp := &zdo.MatchDescResponse{}

		if md.ProfileId == 260 {
			rsp.EndpointList.AddValues(1)
			ep := z.dev.Endpoint(1)
			ep.profile = uint16(md.ProfileId)
			for _, cl := range md.InClusterList.ArrayValues() {
				for _, m := range ep.outCluster {
					if m == zcl.ClusterID(cl) {
						continue
					}
				}
				ep.outCluster = append(ep.outCluster, zcl.ClusterID(cl))
			}
			for _, cl := range md.OutClusterList.ArrayValues() {
				for _, m := range ep.inCluster {
					if m == zcl.ClusterID(cl) {
						continue
					}
				}
				ep.inCluster = append(ep.inCluster, zcl.ClusterID(cl))
			}
		}

		return z.Respond(seqNo, rsp)
	}

	log.Printf("Unhandled ZDO request: %#v", cmd)
	return nil
}

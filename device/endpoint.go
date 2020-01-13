package device

import (
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster"
	"hemtjan.st/zcl/foundation"
	"hemtjan.st/zcl/utils"
	"hemtjan.st/zcl/zdo"

	"encoding/json"
	"errors"
	"log"
	"sort"
	"sync"
	"time"
)

type AttrInfo struct {
	ID      zcl.AttrID
	Cluster zcl.ClusterID
	Type    zcl.TypeID
}

type Endpoint struct {
	attrLock   sync.RWMutex
	init       bool
	dev        *Device
	ep         uint8
	profile    uint16
	devType    zdo.DeviceType
	devVersion uint8
	inCluster  []zcl.ClusterID
	outCluster []zcl.ClusterID

	inCommand  map[zcl.ClusterID][]func() zcl.Command
	outCommand map[zcl.ClusterID][]func() zcl.Command
	attr       map[zcl.ClusterID][]zcl.Attr
}

func (z *Endpoint) Device() *Device      { return z.dev }
func (z *Endpoint) ID() uint8            { return z.ep }
func (z *Endpoint) Type() zdo.DeviceType { return z.devType }
func (z *Endpoint) Profile() uint16      { return z.profile }
func (z *Endpoint) Version() uint8       { return z.devVersion }
func (z *Endpoint) Attributes() map[zcl.ClusterID][]zcl.Attr {
	return z.attr
}

func (z *Endpoint) MarshalJSON() ([]byte, error) {

	a := map[zcl.ClusterID][]json.RawMessage{}
	if z.attr != nil {
		for i, k := range z.attr {
			a[i] = []json.RawMessage{}
			for _, v := range k {
				bv, err := zcl.JsonAttr(v)
				if err != nil {
					return nil, err
				}
				a[i] = append(a[i], bv)
			}
		}
	}

	inCommand := map[zcl.ClusterID][]int{}
	outCommand := map[zcl.ClusterID][]int{}

	for k, v := range z.inCommand {
		inCommand[k] = []int{}
		for _, cmdFn := range v {
			inCommand[k] = append(inCommand[k], int(cmdFn().ID()))
		}
	}
	for k, v := range z.outCommand {
		outCommand[k] = []int{}
		for _, cmdFn := range v {
			outCommand[k] = append(outCommand[k], int(cmdFn().ID()))
		}
	}

	return json.Marshal(&struct {
		Profile    uint16
		Type       uint16
		Version    uint8
		InCluster  []zcl.ClusterID
		OutCluster []zcl.ClusterID
		InCommand  map[zcl.ClusterID][]int
		OutCommand map[zcl.ClusterID][]int
		Attribute  map[zcl.ClusterID][]json.RawMessage
	}{
		Profile:    z.profile,
		Type:       uint16(z.devType),
		Version:    z.devVersion,
		InCluster:  z.inCluster,
		OutCluster: z.outCluster,
		InCommand:  inCommand,
		OutCommand: outCommand,
		Attribute:  a,
	})
}

func (z *Endpoint) UnmarshalJSON(b []byte) error {
	s := &struct {
		Profile    uint16
		Type       uint16
		Version    uint8
		InCluster  []zcl.ClusterID
		OutCluster []zcl.ClusterID
		InCommand  map[zcl.ClusterID][]zcl.CommandID
		OutCommand map[zcl.ClusterID][]zcl.CommandID
		Attribute  map[zcl.ClusterID][]*zcl.UnknownAttr
	}{}
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}

	z.profile = s.Profile
	z.devType = zdo.DeviceType(s.Type)
	z.devVersion = s.Version
	z.inCluster = s.InCluster
	z.outCluster = s.OutCluster
	z.attr = map[zcl.ClusterID][]zcl.Attr{}
	if s.Attribute != nil {
		for clusterId, attribs := range s.Attribute {
			cl, ok := cluster.Clusters[clusterId]
			z.attr[clusterId] = []zcl.Attr{}
			for _, ua := range attribs {
				ua.ClusterID = clusterId
				if ok {
					if clAttr := cl.Attr(uint16(ua.ID())); clAttr != nil {
						if err := clAttr.SetValue(ua.Value()); err != nil {
							log.Printf("Unmarshalling %#v into %#v: %s", ua, clAttr, err)
						}
						z.attr[clusterId] = append(z.attr[clusterId], clAttr)
						continue
					}
				}
				z.attr[clusterId] = append(z.attr[clusterId], ua)
			}
		}
	}

	z.inCommand = getCommands(s.InCommand)
	z.outCommand = getCommands(s.OutCommand)
	return nil
}

func getCommands(list map[zcl.ClusterID][]zcl.CommandID) map[zcl.ClusterID][]func() zcl.Command {
	out := map[zcl.ClusterID][]func() zcl.Command{}
	for clusterId, cmds := range list {
		if cl, ok := cluster.Clusters[clusterId]; ok {
			out[clusterId] = []func() zcl.Command{}
			for _, cmdId := range cmds {
				if cmd := cl.CmdFn(uint8(cmdId)); cmd != nil {
					out[clusterId] = append(out[clusterId], cmd)
				}
			}
		}
	}
	return out
}

type AttrDesc interface {
	AttrID() zcl.AttrID
	AttrType() zcl.TypeID
}

func FindAttr(clID zcl.ClusterID, d AttrDesc) (attr zcl.Attr, err error) {
	cl, ok := cluster.Clusters[clID]
	if ok {
		attrFn, ok := cl.ServerAttr[d.AttrID()]
		if ok && attrFn != nil {
			attr = attrFn()
		} else {
			attrFn, ok := cl.ClientAttr[d.AttrID()]
			if ok && attrFn != nil {
				attr = attrFn()
			}
		}
	}
	if attr == nil {
		nv := d.AttrType().New()
		if nv == nil {
			return nil, zcl.ErrInvalidType
		}

		return &zcl.UnknownAttr{
			Type:      d.AttrType(),
			AttrID:    d.AttrID(),
			AttrValue: nv,
			ClusterID: clID,
		}, nil
	}
	return
}

func (z *Endpoint) Commands() (c []zcl.Command) {
	l := map[zcl.ClusterID]map[zcl.CommandID]zcl.Command{}
	add := func(cmd zcl.Command) {
		if cmd != nil {
			if _, ok := l[cmd.Cluster()]; !ok {
				l[cmd.Cluster()] = map[zcl.CommandID]zcl.Command{}
			}
			l[cmd.Cluster()][cmd.ID()] = cmd
		}
	}

	for _, cmds := range z.inCommand {
		for _, fn := range cmds {
			add(fn())
		}
	}
	for _, clId := range z.inCluster {
		if cl, ok := cluster.Clusters[clId]; ok {
			for _, fn := range cl.ServerCmd {
				cmd := fn()
				if cmd.Required() {
					add(cmd)
				}
			}
		}
	}

	for _, a := range l {
		for _, b := range a {
			c = append(c, b)
		}
	}

	sort.Slice(c, func(i, j int) bool {
		if c[i].Cluster() < c[j].Cluster() {
			return true
		} else if c[i].Cluster() > c[j].Cluster() {
			return false
		}
		return c[i].ID() < c[j].ID()
	})

	return
}

func (z *Endpoint) Init() error {
	if z.init {
		return nil
	}
	z.init = true
	if z.attr == nil {
		z.attr = map[zcl.ClusterID][]zcl.Attr{}
	}

	if err := z.ScanClusters(); err != nil {
		return err
	}
	if err := z.ScanAttributes(); err != nil {
		log.Printf("Error while scanning attributes on %04X/%d: %s", z.dev.nwk, z.ep, err)
	}

	return nil
}

func (z *Endpoint) Request(cmd zcl.Command) (interface{}, error) {
	ch := make(chan interface{})
	seqNo, release := z.dev.seq.Next(ch)
	defer release()

	fr, err := zcl.NewClusterFrame(seqNo, z.ep, z.dev.NWK(), z.ep, zcl.ProfileID(z.profile), cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("TX: [EP(%d) -> %s/%d] Profile[%d] Cluster[%d] Seq[%d] %#v (%X)",
		fr.SrcEP(),
		fr.DstAddr(),
		fr.DstEP(),
		fr.Profile(),
		fr.Cluster(),
		seqNo,
		cmd,
		fr.Payload(),
	)

	err = z.dev.tr.Send(fr)
	if err != nil {
		return nil, err
	}

	if fr.DisableDefaultResponse() {
		return nil, nil
	}

	select {
	case ret := <-ch:
		return ret, nil
	case <-time.After(10 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}
}

func (z *Endpoint) General(cluster uint16, cmd zcl.General) (interface{}, error) {
	ch := make(chan interface{})
	seqNo, release := z.dev.seq.Next(ch)
	defer release()

	fr, err := zcl.NewProfileFrame(seqNo, z.ep, z.dev.NWK(), z.ep, zcl.ProfileID(z.profile), zcl.ClusterID(cluster), cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("TX: [EP(%d) -> %s/%d] Profile[%d] Cluster[%d] Seq[%d] %#v (%X)",
		fr.SrcEP(),
		fr.DstAddr(),
		fr.DstEP(),
		fr.Profile(),
		fr.Cluster(),
		seqNo,
		cmd,
		fr.Payload(),
	)

	err = z.dev.tr.Send(fr)
	if err != nil {
		return nil, err
	}

	if fr.DisableDefaultResponse() {
		return nil, nil
	}

	select {
	case ret := <-ch:
		return ret, nil
	case <-time.After(10 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}
}

func (z *Endpoint) HandleReportAttributes(fr zcl.ZclFrame, report *foundation.ReportAttributes) error {
	z.attrLock.Lock()
	defer z.attrLock.Unlock()
	if _, ok := z.attr[fr.Cluster()]; !ok {
		z.attr[fr.Cluster()] = []zcl.Attr{}
	}
nextAttr:
	for _, r := range report.Attributes {
		for _, a := range z.attr[fr.Cluster()] {
			if r.AttributeID == a.ID() {
				if err := a.SetValue(r.Value); err != nil {
					log.Printf("While trying to update %s to %s: %v", a, r, err)
				}
				continue nextAttr
			}
		}
		av, err := FindAttr(fr.Cluster(), &r)
		if err != nil {
			log.Printf("While trying to create attr %s: %v", r, err)
			continue nextAttr
		}
		z.attr[fr.Cluster()] = append(z.attr[fr.Cluster()], av)
	}
	return nil
}

func (z *Endpoint) HandleZcl(fr zcl.ReceivedZclFrame) (zcl.ZclFrame, error) {
	switch fr.CommandType() {
	case zcl.ClusterSpecific:
		cmd, err := utils.ParseZclClusterFrame(fr)
		if err != nil {
			return nil, err
		}
		if cmd.Direction() == zcl.ServerToClient {
			if z.dev.seq.Handle(fr.SeqNo(), cmd) {
				return nil, nil
			}
		}
		ret, _, err := cmd.Handle(fr, z)
		if !foundation.IsRequest(cmd.ID()) {
			z.dev.seq.Handle(fr.SeqNo(), cmd)
			return nil, nil
		}
		if ret != nil {
			return fr.Response(z.ep, ret)
		}
		return nil, err
	case zcl.ProfileWide:
		cmd, err := utils.ParseZclProfileFrame(fr)
		if err != nil {
			return nil, err
		}
		ret, _, err := cmd.Handle(fr, z)
		if !foundation.IsRequest(cmd.ID()) {
			z.dev.seq.Handle(fr.SeqNo(), cmd)
			return nil, nil
		}
		if ret != nil {
			return fr.Response(z.ep, ret)
		}
		return nil, err
	}
	return nil, zcl.ErrInvalidPacket
}

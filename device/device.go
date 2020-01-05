package device

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/utils"
	"hemtjan.st/zcl/zdo"
	//"hemtjan.st/zcl/zdo_old"
	"log"
	"net"
	"sort"
	"sync"
)

type Device struct {
	sync.RWMutex
	scanning bool
	tr       zcl.Transport
	seq      utils.Seq
	nwk      uint16
	mac      zcl.Zuid
	zdo      *Zdo
	ep       map[uint8]*Endpoint
	cap      zdo.Capability
}
type deviceJson struct {
	MAC        string
	Capability uint8
	Endpoint   map[uint8]*Endpoint
}

func (d *Device) NWK() zcl.Address {
	return utils.NWKAddress(d.nwk)
}
func (d *Device) MAC() zcl.Address {
	return utils.IEEEAddress(d.mac.HWAddr())
}

func (d *Device) Endpoints() (e []*Endpoint) {
	for _, ep := range d.ep {
		e = append(e, ep)
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].ep < e[j].ep
	})
	return
}

func (d *Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(&deviceJson{
		MAC:        d.mac.String(),
		Capability: d.cap[0],
		Endpoint:   d.ep,
	})
}

func (d *Device) UnmarshalJSON(b []byte) error {
	s := &deviceJson{}
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}

	if mac, err := net.ParseMAC(s.MAC); err == nil {
		d.mac = zcl.Zuid(binary.BigEndian.Uint64(mac))
	}

	d.cap = zdo.Capability{s.Capability}
	d.ep = s.Endpoint

	return nil
}

func New(tr zcl.Transport, nwk uint16, mac zcl.Zuid, cap zdo.Capability, seq utils.Seq) *Device {
	return &Device{
		tr:  tr,
		seq: seq,
		nwk: nwk,
		mac: mac,
		ep:  map[uint8]*Endpoint{},
		cap: cap,
	}
}

func (d *Device) Resume(nwk uint16, tr zcl.Transport, seq utils.Seq) {
	d.Lock()
	defer d.Unlock()
	d.nwk = nwk
	d.tr = tr
	d.seq = seq
	if d.ep != nil {
		for id, ep := range d.ep {
			ep.dev = d
			ep.ep = id
		}
	}
}

func (d *Device) OnAnnounce(mac zcl.Zuid, cap zdo.Capability) {
	d.Lock()
	defer d.Unlock()
	d.mac = mac
	d.cap = cap

	// TODO: Check for new endpoints/attributes/commands
}

func (d *Device) Endpoint(ep uint8) *Endpoint {
	d.Lock()
	defer d.Unlock()
	if e, ok := d.ep[ep]; ok {
		return e
	}
	e := &Endpoint{dev: d, ep: ep}
	d.ep[ep] = e
	return e
}

func (d *Device) Zdo() *Zdo {
	d.Lock()
	defer d.Unlock()
	if d.zdo == nil {
		d.zdo = &Zdo{dev: d}
	}
	return d.zdo
}

func (d *Device) Handle(seq uint8, data interface{}) bool {
	return d.seq.Handle(seq, data)
}

func (d *Device) Init() error {
	rsp, err := d.Zdo().Request(&zdo.ActiveEndpointRequest{NwkAddress: zdo.NwkAddress(d.nwk)})
	if err != nil {
		return err
	}
	if activeEndpoints, ok := rsp.(*zdo.ActiveEndpointResponse); ok {
		for _, ep := range activeEndpoints.EndpointList.ArrayValues() {
			go func(ep uint8) {
				if ep == 1 {
					err := d.Endpoint(ep).Init(260)
					if err != nil {
						log.Printf("Error initializing endpoint %d on %04x: %s", ep, d.nwk, err)
					}
				}
			}(ep)
		}
		return nil
	}
	return fmt.Errorf("invalid for ActiveEndpointRequest: %#v", rsp)
}

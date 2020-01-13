package zcl

import "errors"

type zclFrame struct {
	Packet
	disableDefaultResponse bool
	direction              Direction
	mnfCode                Zu16
	cmdType                CommandType
	seq                    uint8
	commandID              CommandID
	data                   []byte
}

type recvZclFrame struct {
	zclFrame
}

func (f *zclFrame) MarshalZcl() ([]byte, error) {
	data := []byte{uint8(f.cmdType) & 0x03}
	if f.direction {
		data[0] = data[0] | 0x08
	}
	if !f.disableDefaultResponse {
		data[0] = data[0] | 0x10
	}

	if f.mnfCode > 0 {
		data[0] = data[0] | 0x04
		t, _ := f.mnfCode.MarshalZcl()
		data = append(data, t...)
	}
	data = append(data, f.seq, uint8(f.commandID))
	data = append(data, f.data...)
	return data, nil
}

func (f *zclFrame) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if len(b) < 3 {
		return nil, errors.New("too few bytes(frame:0)")
	}
	st := b[0]
	b = b[1:]
	f.cmdType = CommandType(st & 0x03)
	f.direction = st&0x08 == 0x08 // 1 = Server to Client
	f.disableDefaultResponse = st&0x10 == 0x10
	if st&0x04 == 0x04 {
		if b, err = (&f.mnfCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}
	f.disableDefaultResponse = st&0x10 == 1
	f.seq = b[0]
	f.commandID = CommandID(b[1])
	f.data = b[2:]
	return nil, nil
}

func (f *zclFrame) SeqNo() uint8                 { return f.seq }
func (f *zclFrame) Direction() Direction         { return f.direction }
func (f *zclFrame) Manufacturer() uint16         { return uint16(f.mnfCode) }
func (f *zclFrame) DisableDefaultResponse() bool { return f.disableDefaultResponse }
func (f *zclFrame) Command() CommandID           { return f.commandID }
func (f *zclFrame) Payload() []byte              { return f.data }
func (f *zclFrame) CommandType() CommandType     { return f.cmdType }

func (f *recvZclFrame) SrcAddr() Address { return f.Packet.(ReceivedPacket).SrcAddr() }
func (f *recvZclFrame) LQI() uint8       { return f.Packet.(ReceivedPacket).LQI() }
func (f *recvZclFrame) RSSI() int8       { return f.Packet.(ReceivedPacket).RSSI() }
func (f *recvZclFrame) Response(srcEp uint8, cmd General) (frame ZclFrame, err error) {
	nf := &zclFrame{
		disableDefaultResponse: true,
		direction:              !f.direction,
		cmdType:                f.cmdType,
		mnfCode:                f.mnfCode,
		seq:                    f.seq,
		commandID:              cmd.ID(),
	}
	frame = nf
	if c, ok := cmd.(Command); ok {
		if nf.cmdType != 1 {
			return nil, ErrResponseClusterSpecific
		}
		if c.Cluster() != f.Cluster() {
			return nil, ErrResponseWrongCluster
		}
		if c.MnfCode() != uint16(f.mnfCode) {
			return nil, ErrResponseWrongMnfCode
		}
	} else if cmd.ID() != 0x0b { // if not Default Response
		if nf.cmdType != 0 {
			return nil, ErrResponseProfileWide
		}
		if c.MnfCode() != uint16(f.mnfCode) {
			return nil, ErrResponseWrongMnfCode
		}
	} else {
		nf.cmdType = 0
	}

	if nf.data, err = cmd.MarshalZcl(); err != nil {
		return nil, err
	}
	nf.Packet, err = NewPacket(f.DstEP(), f.SrcAddr(), srcEp, f.Profile(), f.Cluster(), nf)
	return
}

func NewProfileFrame(seq uint8, srcEp uint8, dstAddr Address, dstEp uint8, profile ProfileID, cluster ClusterID, cmd General) (frame ZclFrame, err error) {
	if _, ok := cmd.(Command); ok {
		return nil, ErrCommandClusterSpecific
	}
	nf := &zclFrame{
		seq:                    seq,
		disableDefaultResponse: false,
		direction:              cmd.Direction(),
		cmdType:                0,
		mnfCode:                0,
		commandID:              cmd.ID(),
	}
	frame = nf
	if nf.data, err = cmd.MarshalZcl(); err != nil {
		return nil, err
	}
	nf.Packet, err = NewPacket(srcEp, dstAddr, dstEp, profile, cluster, nf)
	return
}

func NewClusterFrame(seq uint8, srcEp uint8, dstAddr Address, dstEp uint8, profile ProfileID, cmd Command) (frame ZclFrame, err error) {
	nf := &zclFrame{
		seq:                    seq,
		disableDefaultResponse: false,
		direction:              cmd.Direction(),
		cmdType:                1,
		mnfCode:                Zu16(cmd.MnfCode()),
		commandID:              cmd.ID(),
	}
	frame = nf
	if nf.data, err = cmd.MarshalZcl(); err != nil {
		return nil, err
	}
	nf.Packet, err = NewPacket(srcEp, dstAddr, dstEp, profile, cmd.Cluster(), nf)
	return
}

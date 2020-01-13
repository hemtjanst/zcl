package zcl

type zdpFrame struct {
	Packet
	seq  uint8
	data []byte
}

type recvZdpFrame struct {
	zdpFrame
}

func (f *zdpFrame) MarshalZcl() ([]byte, error) {
	return append([]byte{f.seq}, f.data...), nil
}

func (f *zdpFrame) UnmarshalZcl(b []byte) ([]byte, error) {
	if len(b) < 1 {
		return b, ErrNotEnoughData
	}
	f.seq = b[0]
	f.data = b[1:]
	return nil, nil
}

func (f *zdpFrame) SeqNo() uint8      { return f.seq }
func (f *zdpFrame) Command() ZdoCmdID { return ZdoCmdID(f.Cluster()) }
func (f *zdpFrame) Payload() []byte   { return f.data }

func (f *recvZdpFrame) SrcAddr() Address { return f.Packet.(ReceivedPacket).SrcAddr() }
func (f *recvZdpFrame) LQI() uint8       { return f.Packet.(ReceivedPacket).LQI() }
func (f *recvZdpFrame) RSSI() int8       { return f.Packet.(ReceivedPacket).RSSI() }
func (f *recvZdpFrame) Response(cmd ZdoCommand) (frame ZdpFrame, err error) {
	nf := &zdpFrame{seq: f.seq}
	frame = nf
	if nf.data, err = cmd.MarshalZcl(); err != nil {
		return nil, err
	}
	nf.Packet, err = NewPacket(f.DstEP(), f.SrcAddr(), 0, 0, cmd.Cluster(), nf)
	return
}

func NewZdpFrame(seq uint8, dstAddr Address, cmd ZdoCommand) (frame ZdpFrame, err error) {
	nf := &zdpFrame{seq: seq}
	if nf.data, err = cmd.MarshalZcl(); err != nil {
		return nil, err
	}
	nf.Packet, err = NewPacket(0, dstAddr, 0, 0, cmd.Cluster(), nf)
	return nf, nil
}

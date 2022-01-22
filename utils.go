package zcl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	StrJoin   = strings.Join
	Sprintf   = fmt.Sprintf
	Errorf    = fmt.Errorf
	ToJson    = json.Marshal
	ParseJson = json.Unmarshal
)

func Handle(handler Handler, packet Packet) (Packet, error) {
	if handler == nil {
		return nil, Errorf("invalid handler")
	}
	if packet == nil {
		return nil, Errorf("invalid packet")
	}
	if packet.SrcEP() == 0 {
		frame := &recvZdpFrame{zdpFrame{Packet: packet}}
		extra, err := frame.UnmarshalZcl(packet.Data())
		if err != nil {
			return nil, err
		}
		if len(extra) > 0 {
			log.Printf("Got extra data after %v: %X", frame, extra)
		}
		return handler.HandleZdp(frame)
	} else {
		frame := &recvZclFrame{zclFrame{Packet: packet}}
		extra, err := frame.UnmarshalZcl(packet.Data())
		if err != nil {
			return nil, err
		}
		if len(extra) > 0 {
			log.Printf("Got extra data after %v: %X", frame, extra)
		}

		rsp, err := handler.HandleZcl(frame)

		// Check if we should generate a default response
		if rsp != nil || frame.DstAddr().Mode() == GroupAddress || frame.DstEP() == 0 || frame.DisableDefaultResponse() {
			return rsp, err
		}

		// Default to Success for no error
		status := Success
		if err != nil {
			// Default to SoftwareFailure for unknown error
			status = SoftwareFailure
			if st, ok := err.(StatusError); ok {
				// If error supports status code, use that
				status = st.Status()
			}
		}

		return frame.Response(frame.DstEP(), &DefaultResponse{
			Status:  status,
			Command: frame.Command(),
		})
	}
}

func Duration(t int, m int) time.Duration {
	d := time.Duration(t) * time.Second
	if m != 0 {
		d = d / time.Duration(m)
	}
	return d
}

func Equal(a, b Val) bool {
	{
		aa, ok1 := a.(AttrValue)
		bb, ok2 := b.(AttrValue)
		if ok1 != ok2 {
			return false
		}
		if ok1 && ok2 {
			return equalAttrVal(aa, bb)
		}
	}
	{
		aa, ok1 := a.(Command)
		bb, ok2 := b.(Command)
		if ok1 != ok2 {
			return false
		}
		if ok1 && ok2 {
			return equalCommand(aa, bb)
		}
	}
	{
		aa, ok1 := a.(ZdoCommand)
		bb, ok2 := b.(ZdoCommand)
		if ok1 != ok2 {
			return false
		}
		if ok1 && ok2 {
			return equalZdoCommand(aa, bb)
		}
	}
	{
		aa, ok1 := a.(General)
		bb, ok2 := b.(General)
		if ok1 != ok2 {
			return false
		}
		if ok1 && ok2 {
			return equalGeneral(aa, bb)
		}
	}
	{
		aa, ok1 := a.(Attr)
		bb, ok2 := b.(Attr)
		if ok1 != ok2 {
			return false
		}
		if ok1 && ok2 {
			return equalAttr(aa, bb)
		}
	}
	{
		aa, ok1 := a.(TypeVal)
		bb, ok2 := b.(TypeVal)
		if ok1 != ok2 {
			return false
		}
		if ok1 && ok2 {
			return equalTypeVal(aa, bb)
		}
	}
	return equalVal(a, b)
}

func equalVal(a, b Val) bool {
	aa, err1 := a.MarshalZcl()
	bb, err2 := b.MarshalZcl()
	return err1 == err2 && equalBytes(aa, bb)
}

func equalTypeVal(a, b TypeVal) bool {
	return a.TypeID() == b.TypeID() &&
		equalVal(a, b)
}

func equalAttr(a, b Attr) bool {
	return a.ID() == b.ID() &&
		equalTypeVal(a, b)
}

func equalAttrVal(a, b AttrValue) bool {
	return a.AttrID() == b.AttrID() &&
		a.AttrType() == b.AttrType() &&
		equalVal(a.AttrValue(), b.AttrValue())
}

func equalBytes(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

func equalGeneral(a, b General) bool {
	return a.ID() == b.ID() &&
		equalVal(a, b)
}

func equalCommand(a, b Command) bool {
	return a.ID() == b.ID() &&
		a.Cluster() == b.Cluster() &&
		equalVal(a, b)
}

func equalZdoCommand(a, b ZdoCommand) bool {
	return a.ID() == b.ID() &&
		a.Cluster() == b.Cluster() &&
		equalVal(a, b)
}

func reverse(s []byte) []byte {
	a := make([]byte, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

package utils

import (
	"fmt"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster"
	"hemtjan.st/zcl/foundation"
	"hemtjan.st/zcl/zdo"
)

func ParseZclProfileFrame(frame zcl.ZclFrame) (cmd zcl.General, err error) {
	if frame.CommandType() != zcl.ProfileWide {
		return nil, zcl.ErrCommandClusterSpecific
	}
	if cmdFn, ok := foundation.Commands[frame.Command()]; ok {
		cmd = cmdFn()
		_, err = cmd.UnmarshalZcl(frame.Payload())
		return
	} else {
		return UnknownZclCommand{frame}, zcl.ErrUnknownCommand
	}
}

func ParseZclClusterFrame(frame zcl.ZclFrame) (cmd zcl.Command, err error) {
	if frame.CommandType() != zcl.ClusterSpecific {
		return nil, zcl.ErrCommandProfileWide
	}
	cl, ok := cluster.Clusters[frame.Cluster()]
	if !ok {
		return UnknownClusterCommand{UnknownZclCommand{frame}}, zcl.ErrUnknownCommand
	}

	if frame.Direction() == zcl.ClientToServer {
		if cmdFn, ok := cl.ServerCmd[frame.Command()]; ok {
			cmd = cmdFn()
		}
	} else {
		if cmdFn, ok := cl.ClientCmd[frame.Command()]; ok {
			cmd = cmdFn()
		}
	}
	if cmd == nil {
		return UnknownClusterCommand{UnknownZclCommand{frame}}, zcl.ErrUnknownCommand
	}
	_, err = cmd.UnmarshalZcl(frame.Payload())
	return
}

func ParseZdpFrame(frame zcl.ZdpFrame) (cmd zcl.ZdoCommand, err error) {
	if cmdFn, ok := zdo.Commands[frame.Command()]; ok {
		cmd = cmdFn()
		_, err = cmd.UnmarshalZcl(frame.Payload())
		return
	} else {
		return UnknownZdpCommand{frame}, zcl.ErrUnknownCommand
	}
}

type UnknownZdpCommand struct {
	zcl.ZdpFrame
}

func (c UnknownZdpCommand) UnmarshalZcl([]byte) ([]byte, error) { return nil, zcl.ErrUnknownCommand }
func (c UnknownZdpCommand) MarshalZcl() ([]byte, error)         { return c.Payload(), nil }
func (c UnknownZdpCommand) Arguments() []zcl.ArgDesc            { return nil }
func (c UnknownZdpCommand) Values() []zcl.Val {
	v := zcl.Zbytes(c.Payload())
	return []zcl.Val{&v}
}
func (c UnknownZdpCommand) ID() zcl.ZdoCmdID { return c.Command() }
func (c UnknownZdpCommand) Name() string {
	return fmt.Sprintf("ZDP(%04X)", c.Command())
}
func (c UnknownZdpCommand) String() string { return zcl.Zbytes(c.Payload()).String() }
func (c UnknownZdpCommand) Handle(zcl.ReceivedZdpFrame, interface{}) (response zcl.ZdoCommand, found bool, err error) {
	return nil, false, zcl.ErrUnknownCommand
}

type UnknownClusterCommand struct {
	UnknownZclCommand
}

type UnknownZclCommand struct {
	zcl.ZclFrame
}

func (c UnknownZclCommand) Arguments() []zcl.ArgDesc { return nil }
func (c UnknownZclCommand) Required() bool           { return false }
func (c UnknownZclCommand) Values() []zcl.Val {
	v := zcl.Zbytes(c.Payload())
	return []zcl.Val{&v}
}
func (c UnknownZclCommand) ID() zcl.CommandID { return c.Command() }
func (c UnknownZclCommand) Name() string {
	return fmt.Sprintf("%s(%04X/%02X)", c.Direction(), c.Cluster(), c.Command())
}
func (c UnknownZclCommand) String() string                      { return zcl.Zbytes(c.Payload()).String() }
func (c UnknownZclCommand) UnmarshalZcl([]byte) ([]byte, error) { return nil, zcl.ErrUnknownCommand }
func (c UnknownZclCommand) MarshalZcl() ([]byte, error)         { return c.Payload(), nil }
func (c UnknownZclCommand) Handle(zcl.ReceivedZclFrame, interface{}) (response zcl.General, found bool, err error) {
	return nil, false, zcl.ErrUnknownCommand
}

func (c UnknownClusterCommand) MnfCode() uint16 { return c.Manufacturer() }

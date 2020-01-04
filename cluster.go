package zcl

type ClusterID Zu16

func (i ClusterID) MarshalZcl() ([]byte, error) { return Zu16(i).MarshalZcl() }
func (i *ClusterID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu16)
	b, err := v.UnmarshalZcl(b)
	*i = ClusterID(*v)
	return b, err
}

type Cluster struct {
	Name       string
	ServerCmd  map[CommandID]func() Command
	ServerAttr map[AttrID]func() Attr
	ClientCmd  map[CommandID]func() Command
	ClientAttr map[AttrID]func() Attr
	SceneAttr  []AttrID
}

func (c *Cluster) Attr(id uint16) Attr {
	if c.ServerAttr != nil {
		if v, ok := c.ServerAttr[AttrID(id)]; ok {
			return v()
		}
	}
	if c.ClientAttr != nil {
		if v, ok := c.ClientAttr[AttrID(id)]; ok {
			return v()
		}
	}
	return nil
}

func (c *Cluster) CmdFn(id uint8) func() Command {
	if c.ServerCmd != nil {
		if v, ok := c.ServerCmd[CommandID(id)]; ok {
			return v
		}
	}
	if c.ClientCmd != nil {
		if v, ok := c.ClientCmd[CommandID(id)]; ok {
			return v
		}
	}
	return nil
}

func (c *Cluster) Cmd(id uint8) Command {
	if fn := c.CmdFn(id); fn != nil {
		return fn()
	}
	return nil
}

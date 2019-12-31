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

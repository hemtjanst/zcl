package zcl

type AttrID Zu16

func (i AttrID) MarshalZcl() ([]byte, error) { return Zu16(i).MarshalZcl() }

type AttrDef interface {
	ID() AttrID
	Value() Val
}

type UnknownAttr struct {
	TypeID    TypeID
	AttrID    AttrID
	ClusterID ClusterID
	AttrValue Val
}

type Attr interface {
	Val
	ID() AttrID
	Name() string
	String() string
	Readable() bool
	Writable() bool
	Reportable() bool
	SceneIndex() int
	SetValue(v Val) error
}

type ClusterAttrImpl struct {
	Attr
	ClusterID ClusterID
}

func (c ClusterAttrImpl) Cluster() ClusterID {
	return c.ClusterID
}

type ClusterAttr interface {
	Attr
	Cluster() ClusterID
}

func (a UnknownAttr) ID() AttrID                             { return a.AttrID }
func (a UnknownAttr) Name() string                           { return Sprintf("Unknown(0x%04X)", a.AttrID) }
func (a UnknownAttr) String() string                         { return Sprintf("%v", a.AttrValue) }
func (a UnknownAttr) Cluster() ClusterID                     { return a.ClusterID }
func (a UnknownAttr) Readable() bool                         { return true }
func (a UnknownAttr) Writable() bool                         { return false }
func (a UnknownAttr) Reportable() bool                       { return false }
func (a UnknownAttr) SceneIndex() int                        { return -1 }
func (a UnknownAttr) SetValue(v Val) error                   { a.AttrValue = v; return nil }
func (a *UnknownAttr) Values() []Val                         { return []Val{a.AttrValue} }
func (a UnknownAttr) MarshalZcl() ([]byte, error)            { return a.AttrValue.MarshalZcl() }
func (a *UnknownAttr) UnmarshalZcl(b []byte) ([]byte, error) { return a.AttrValue.UnmarshalZcl(b) }

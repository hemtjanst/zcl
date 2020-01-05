package zcl

import "encoding/json"

type AttrID Zu16

func (i AttrID) MarshalZcl() ([]byte, error) { return Zu16(i).MarshalZcl() }

type AttrDef interface {
	ID() AttrID
	Value() Val
}

type UnknownAttr struct {
	Type      TypeID
	AttrID    AttrID
	ClusterID ClusterID
	AttrValue Val
}

func (a *UnknownAttr) UnmarshalJSON(b []byte) error {
	s := &struct {
		ID    AttrID
		Type  TypeID
		Value json.RawMessage
	}{}
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	a.AttrID = s.ID
	a.Type = s.Type
	a.AttrValue = NewValue(uint8(s.Type))
	return json.Unmarshal(s.Value, &a.AttrValue)
}

func JsonAttr(a Attr) ([]byte, error) {
	return json.Marshal(struct {
		ID    AttrID
		Type  TypeID
		Value Val
	}{
		ID:    a.ID(),
		Type:  a.TypeID(),
		Value: a.Value(),
	})
}

type Argument interface {
	Val
	TypeID() TypeID
	Name() string
	String() string
	Value() Val
	SetValue(v Val) error
}

type Attr interface {
	Argument
	ID() AttrID
	Readable() bool
	Writable() bool
	Reportable() bool
	SceneIndex() int
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
func (a UnknownAttr) TypeID() TypeID                         { return a.Type }
func (a UnknownAttr) String() string                         { return Sprintf("%v", a.AttrValue) }
func (a UnknownAttr) Cluster() ClusterID                     { return a.ClusterID }
func (a UnknownAttr) Readable() bool                         { return true }
func (a UnknownAttr) Writable() bool                         { return false }
func (a UnknownAttr) Reportable() bool                       { return false }
func (a UnknownAttr) SceneIndex() int                        { return -1 }
func (a UnknownAttr) SetValue(v Val) error                   { a.AttrValue = v; return nil }
func (a *UnknownAttr) Value() Val                            { return a.AttrValue }
func (a UnknownAttr) MarshalZcl() ([]byte, error)            { return a.AttrValue.MarshalZcl() }
func (a *UnknownAttr) UnmarshalZcl(b []byte) ([]byte, error) { return a.AttrValue.UnmarshalZcl(b) }

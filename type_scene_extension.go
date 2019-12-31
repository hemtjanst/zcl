package zcl

type SceneExtension interface {
	ClusterID() uint16
	Val
}

type SceneExtensionSet []SceneExtension

func (s SceneExtensionSet) MarshalZcl() ([]byte, error) {
	var data []byte
	for _, v := range s {
		cl := Zu16(v.ClusterID())
		clDt, _ := cl.MarshalZcl()
		dt, err := v.MarshalZcl()
		if err != nil {
			return nil, err
		}
		data = append(data, append(clDt, uint8(len(dt)))...)
		data = append(data, dt...)
	}
	return data, nil
}

func (s *SceneExtensionSet) UnmarshalZcl(b []byte) ([]byte, error) {
	for len(b) > 0 {
		se := &sceneExtension{}
		var err error
		b, err = (&se.cluster).UnmarshalZcl(b)
		if err != nil {
			return nil, err
		}
		b, err = (&se.data).UnmarshalZcl(b)
		if err != nil {
			return nil, err
		}
		*s = append(*s, se)
	}
	return nil, nil
}

func (s SceneExtensionSet) String() string {
	var str []string
	for _, v := range s {
		str = append(str, Sprintf("%v", v))
	}
	return StrJoin(str, ", ")
}

type sceneExtension struct {
	cluster Zu16
	data    Zbytes
}

func (s sceneExtension) MarshalZcl() ([]byte, error) {
	return []byte(s.data), nil
}
func (s *sceneExtension) UnmarshalZcl(b []byte) ([]byte, error) {
	s.data = Zbytes(b)
	return nil, nil
}
func (s *sceneExtension) ClusterID() uint16 {
	return uint16(s.cluster)
}

func (s *sceneExtension) String() string {
	return Sprintf("Cluster[0x%04X] %s", s.cluster, s.data.String())
}

package foundation

import "hemtjan.st/zcl"

type DefaultResponse struct {
	Command zcl.CommandID
	Status  zcl.Status
}

func (v DefaultResponse) MarshalZcl() ([]byte, error) {
	var tmp []byte
	var data []byte
	var err error
	if tmp, err = v.Command.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	return data, nil
}

func (v *DefaultResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.Command).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (v *DefaultResponse) Values() []zcl.Val { return []zcl.Val{&v.Command, &v.Status} }
func (DefaultResponse) ID() zcl.CommandID    { return DefaultResponseCommand }
func (DefaultResponse) Name() string         { return "DefaultResponse" }
func (v DefaultResponse) String() string {
	return zcl.Sprintf("DefaultResponse{0x%02X: %v}", v.Command, v.Status.String())
}

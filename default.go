package zcl

const DefaultResponseCommand CommandID = 0x0b

type DefaultResponse struct {
	Command CommandID
	Status  Status
}

type DefaultResponseHandler interface {
	HandleDefaultResponse(ZclFrame, *DefaultResponse) error
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

func (v *DefaultResponse) Values() []Val { return []Val{&v.Command, &v.Status} }
func (DefaultResponse) ID() CommandID    { return DefaultResponseCommand }
func (DefaultResponse) Name() string     { return "DefaultResponse" }
func (v DefaultResponse) String() string {
	return Sprintf("DefaultResponse{0x%02X: %v}", v.Command, v.Status.String())
}

func (DefaultResponse) Direction() Direction { return ServerToClient }

func (v *DefaultResponse) Handle(frame ReceivedZclFrame, handler interface{}) (rsp General, found bool, err error) {
	var h DefaultResponseHandler
	if h, found = handler.(DefaultResponseHandler); found {
		err = h.HandleDefaultResponse(frame, v)
	}
	return
}

package otau

import "hemtjan.st/zcl"

// Otau
const OtauID zcl.ClusterID = 25

var OtauCluster = zcl.Cluster{
	Name: "OTAU",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		QueryNextImageCommand:    func() zcl.Command { return new(QueryNextImage) },
		ImageBlockRequestCommand: func() zcl.Command { return new(ImageBlockRequest) },
		UpgradeEndRequestCommand: func() zcl.Command { return new(UpgradeEndRequest) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		ImageNotifyCommand:            func() zcl.Command { return new(ImageNotify) },
		QueryNextImageResponseCommand: func() zcl.Command { return new(QueryNextImageResponse) },
		ImageBlockResponseCommand:     func() zcl.Command { return new(ImageBlockResponse) },
		UpgradeEndResponseCommand:     func() zcl.Command { return new(UpgradeEndResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{
		UpgradeServerAttr:          func() zcl.Attr { return new(UpgradeServer) },
		FileOffsetAttr:             func() zcl.Attr { return new(FileOffset) },
		CurrentFileVersionAttr:     func() zcl.Attr { return new(CurrentFileVersion) },
		CurrentStackVersionAttr:    func() zcl.Attr { return new(CurrentStackVersion) },
		DownloadedFileVersionAttr:  func() zcl.Attr { return new(DownloadedFileVersion) },
		DownloadedStackVersionAttr: func() zcl.Attr { return new(DownloadedStackVersion) },
		ImageUpgradeStatusAttr:     func() zcl.Attr { return new(ImageUpgradeStatus) },
		MinBlockRequestDelayAttr:   func() zcl.Attr { return new(MinBlockRequestDelay) },
	},
	SceneAttr: []zcl.AttrID{},
}

type QueryNextImage struct {
	HardwareVersionPresent HardwareVersionPresent
	ManufacturerCode       ManufacturerCode
	ImageType              ImageType
	FileVersion            FileVersion
	HardwareVersion        HardwareVersion
}

type QueryNextImageHandler interface {
	HandleQueryNextImage(frame Frame, cmd *QueryNextImage) (*QueryNextImageResponse, error)
}

// QueryNextImageCommand is the Command ID of QueryNextImage
const QueryNextImageCommand CommandID = 0x0001

// Values returns all values of QueryNextImage
func (v *QueryNextImage) Values() []zcl.Val {
	return []zcl.Val{
		&v.HardwareVersionPresent,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
		&v.HardwareVersion,
	}
}

// Arguments returns all values of QueryNextImage
func (v *QueryNextImage) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "HardwareVersionPresent", Argument: &v.HardwareVersionPresent},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
		{Name: "HardwareVersion", Argument: &v.HardwareVersion},
	}
}

// Name of the command
func (QueryNextImage) Name() string { return `Query next image` }

// Description of the command
func (QueryNextImage) Description() string { return `` }

// ID of the command
func (QueryNextImage) ID() CommandID { return QueryNextImageCommand }

// Required
func (QueryNextImage) Required() bool { return true }

// Cluster ID of the command
func (QueryNextImage) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (QueryNextImage) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (QueryNextImage) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (QueryNextImage) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *QueryNextImage) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h QueryNextImageHandler
	if h, found = handler.(QueryNextImageHandler); found {
		rsp, err = h.HandleQueryNextImage(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of QueryNextImage
func (v QueryNextImage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.HardwareVersionPresent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if v.HardwareVersionPresent == 0x01 {
		if tmp, err = v.HardwareVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the QueryNextImage struct
func (v *QueryNextImage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.HardwareVersionPresent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if v.HardwareVersionPresent == 0x01 {
		if b, err = (&v.HardwareVersion).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v QueryNextImage) String() string {
	return zcl.Sprintf(
		"QueryNextImage{"+zcl.StrJoin([]string{
			"HardwareVersionPresent(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
			"HardwareVersion(%v)",
		}, " ")+"}",
		v.HardwareVersionPresent,
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
		v.HardwareVersion,
	)
}

type ImageBlockRequest struct {
	BlockRequestOptions BlockRequestOptions
	ManufacturerCode    ManufacturerCode
	ImageType           ImageType
	FileVersion         FileVersion
	FileOffset          FileOffset
	MaxDataSize         DataSize
}

type ImageBlockRequestHandler interface {
	HandleImageBlockRequest(frame Frame, cmd *ImageBlockRequest) (*ImageBlockResponse, error)
}

// ImageBlockRequestCommand is the Command ID of ImageBlockRequest
const ImageBlockRequestCommand CommandID = 0x0003

// Values returns all values of ImageBlockRequest
func (v *ImageBlockRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.BlockRequestOptions,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
		&v.FileOffset,
		&v.MaxDataSize,
	}
}

// Arguments returns all values of ImageBlockRequest
func (v *ImageBlockRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "BlockRequestOptions", Argument: &v.BlockRequestOptions},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
		{Name: "FileOffset", Argument: &v.FileOffset},
		{Name: "MaxDataSize", Argument: &v.MaxDataSize},
	}
}

// Name of the command
func (ImageBlockRequest) Name() string { return `Image Block Request` }

// Description of the command
func (ImageBlockRequest) Description() string { return `` }

// ID of the command
func (ImageBlockRequest) ID() CommandID { return ImageBlockRequestCommand }

// Required
func (ImageBlockRequest) Required() bool { return true }

// Cluster ID of the command
func (ImageBlockRequest) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (ImageBlockRequest) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ImageBlockRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ImageBlockRequest) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *ImageBlockRequest) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ImageBlockRequestHandler
	if h, found = handler.(ImageBlockRequestHandler); found {
		rsp, err = h.HandleImageBlockRequest(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of ImageBlockRequest
func (v ImageBlockRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.BlockRequestOptions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileOffset.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MaxDataSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ImageBlockRequest struct
func (v *ImageBlockRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.BlockRequestOptions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileOffset).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxDataSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ImageBlockRequest) String() string {
	return zcl.Sprintf(
		"ImageBlockRequest{"+zcl.StrJoin([]string{
			"BlockRequestOptions(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
			"FileOffset(%v)",
			"MaxDataSize(%v)",
		}, " ")+"}",
		v.BlockRequestOptions,
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
		v.FileOffset,
		v.MaxDataSize,
	)
}

type UpgradeEndRequest struct {
	Status           Status
	ManufacturerCode ManufacturerCode
	ImageType        ImageType
	FileVersion      FileVersion
}

type UpgradeEndRequestHandler interface {
	HandleUpgradeEndRequest(frame Frame, cmd *UpgradeEndRequest) (*UpgradeEndResponse, error)
}

// UpgradeEndRequestCommand is the Command ID of UpgradeEndRequest
const UpgradeEndRequestCommand CommandID = 0x0006

// Values returns all values of UpgradeEndRequest
func (v *UpgradeEndRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
	}
}

// Arguments returns all values of UpgradeEndRequest
func (v *UpgradeEndRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
	}
}

// Name of the command
func (UpgradeEndRequest) Name() string { return `Upgrade End Request` }

// Description of the command
func (UpgradeEndRequest) Description() string { return `` }

// ID of the command
func (UpgradeEndRequest) ID() CommandID { return UpgradeEndRequestCommand }

// Required
func (UpgradeEndRequest) Required() bool { return true }

// Cluster ID of the command
func (UpgradeEndRequest) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (UpgradeEndRequest) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (UpgradeEndRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UpgradeEndRequest) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

func (v *UpgradeEndRequest) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h UpgradeEndRequestHandler
	if h, found = handler.(UpgradeEndRequestHandler); found {
		rsp, err = h.HandleUpgradeEndRequest(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of UpgradeEndRequest
func (v UpgradeEndRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UpgradeEndRequest struct
func (v *UpgradeEndRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UpgradeEndRequest) String() string {
	return zcl.Sprintf(
		"UpgradeEndRequest{"+zcl.StrJoin([]string{
			"Status(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
		}, " ")+"}",
		v.Status,
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
	)
}

type ImageNotify struct {
	ImageNotifyPayloadType ImageNotifyPayloadType
	QueryJitter            QueryJitter
	ManufacturerCode       ManufacturerCode
	ImageType              ImageType
	FileVersion            FileVersion
}

type ImageNotifyHandler interface {
	HandleImageNotify(frame Frame, cmd *ImageNotify) error
}

// ImageNotifyCommand is the Command ID of ImageNotify
const ImageNotifyCommand CommandID = 0x0000

// Values returns all values of ImageNotify
func (v *ImageNotify) Values() []zcl.Val {
	return []zcl.Val{
		&v.ImageNotifyPayloadType,
		&v.QueryJitter,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
	}
}

// Arguments returns all values of ImageNotify
func (v *ImageNotify) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ImageNotifyPayloadType", Argument: &v.ImageNotifyPayloadType},
		{Name: "QueryJitter", Argument: &v.QueryJitter},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
	}
}

// Name of the command
func (ImageNotify) Name() string { return `Image notify` }

// Description of the command
func (ImageNotify) Description() string { return `` }

// ID of the command
func (ImageNotify) ID() CommandID { return ImageNotifyCommand }

// Required
func (ImageNotify) Required() bool { return false }

// Cluster ID of the command
func (ImageNotify) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (ImageNotify) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (ImageNotify) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ImageNotify) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *ImageNotify) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ImageNotifyHandler
	if h, found = handler.(ImageNotifyHandler); found {
		err = h.HandleImageNotify(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ImageNotify
func (v ImageNotify) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ImageNotifyPayloadType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.QueryJitter.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if v.ImageNotifyPayloadType >= 0x01 {
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if v.ImageNotifyPayloadType >= 0x02 {
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if v.ImageNotifyPayloadType >= 0x03 {
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ImageNotify struct
func (v *ImageNotify) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.ImageNotifyPayloadType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.QueryJitter).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if v.ImageNotifyPayloadType >= 0x01 {
		if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.ImageNotifyPayloadType >= 0x02 {
		if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.ImageNotifyPayloadType >= 0x03 {
		if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ImageNotify) String() string {
	return zcl.Sprintf(
		"ImageNotify{"+zcl.StrJoin([]string{
			"ImageNotifyPayloadType(%v)",
			"QueryJitter(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
		}, " ")+"}",
		v.ImageNotifyPayloadType,
		v.QueryJitter,
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
	)
}

type QueryNextImageResponse struct {
	Status           Status
	ManufacturerCode ManufacturerCode
	ImageType        ImageType
	FileVersion      FileVersion
	ImageSize        ImageSize
}

type QueryNextImageResponseHandler interface {
	HandleQueryNextImageResponse(frame Frame, cmd *QueryNextImageResponse) error
}

// QueryNextImageResponseCommand is the Command ID of QueryNextImageResponse
const QueryNextImageResponseCommand CommandID = 0x0002

// Values returns all values of QueryNextImageResponse
func (v *QueryNextImageResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
		&v.ImageSize,
	}
}

// Arguments returns all values of QueryNextImageResponse
func (v *QueryNextImageResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
		{Name: "ImageSize", Argument: &v.ImageSize},
	}
}

// Name of the command
func (QueryNextImageResponse) Name() string { return `Query next image response` }

// Description of the command
func (QueryNextImageResponse) Description() string { return `` }

// ID of the command
func (QueryNextImageResponse) ID() CommandID { return QueryNextImageResponseCommand }

// Required
func (QueryNextImageResponse) Required() bool { return true }

// Cluster ID of the command
func (QueryNextImageResponse) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (QueryNextImageResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (QueryNextImageResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (QueryNextImageResponse) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *QueryNextImageResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h QueryNextImageResponseHandler
	if h, found = handler.(QueryNextImageResponseHandler); found {
		err = h.HandleQueryNextImageResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of QueryNextImageResponse
func (v QueryNextImageResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ManufacturerCode is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ImageType is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// FileVersion is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ImageSize is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ImageSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the QueryNextImageResponse struct
func (v *QueryNextImageResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// ManufacturerCode is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ImageType is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// FileVersion is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ImageSize is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ImageSize).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v QueryNextImageResponse) String() string {
	return zcl.Sprintf(
		"QueryNextImageResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
			"ImageSize(%v)",
		}, " ")+"}",
		v.Status,
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
		v.ImageSize,
	)
}

type ImageBlockResponse struct {
	Status           Status
	ManufacturerCode ManufacturerCode
	ImageType        ImageType
	FileVersion      FileVersion
	FileOffset       FileOffset
	Payload          Payload
	CurrentTime      CurrentTime
	RequestTime      RequestTime
}

type ImageBlockResponseHandler interface {
	HandleImageBlockResponse(frame Frame, cmd *ImageBlockResponse) error
}

// ImageBlockResponseCommand is the Command ID of ImageBlockResponse
const ImageBlockResponseCommand CommandID = 0x0005

// Values returns all values of ImageBlockResponse
func (v *ImageBlockResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
		&v.FileOffset,
		&v.Payload,
		&v.CurrentTime,
		&v.RequestTime,
	}
}

// Arguments returns all values of ImageBlockResponse
func (v *ImageBlockResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
		{Name: "FileOffset", Argument: &v.FileOffset},
		{Name: "Payload", Argument: &v.Payload},
		{Name: "CurrentTime", Argument: &v.CurrentTime},
		{Name: "RequestTime", Argument: &v.RequestTime},
	}
}

// Name of the command
func (ImageBlockResponse) Name() string { return `Image Block Response` }

// Description of the command
func (ImageBlockResponse) Description() string { return `` }

// ID of the command
func (ImageBlockResponse) ID() CommandID { return ImageBlockResponseCommand }

// Required
func (ImageBlockResponse) Required() bool { return true }

// Cluster ID of the command
func (ImageBlockResponse) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (ImageBlockResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (ImageBlockResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ImageBlockResponse) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

func (v *ImageBlockResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ImageBlockResponseHandler
	if h, found = handler.(ImageBlockResponseHandler); found {
		err = h.HandleImageBlockResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ImageBlockResponse
func (v ImageBlockResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ManufacturerCode is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ImageType is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// FileVersion is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// FileOffset is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.FileOffset.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// Payload is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.Payload.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// CurrentTime is only included if wait for data status
	if v.Status == 0x97 {
		if tmp, err = v.CurrentTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// RequestTime is only included if wait for data status
	if v.Status == 0x97 {
		if tmp, err = v.RequestTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ImageBlockResponse struct
func (v *ImageBlockResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// ManufacturerCode is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ImageType is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// FileVersion is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// FileOffset is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.FileOffset).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// Payload is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.Payload).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// CurrentTime is only included if wait for data status
	if v.Status == 0x97 {
		if b, err = (&v.CurrentTime).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// RequestTime is only included if wait for data status
	if v.Status == 0x97 {
		if b, err = (&v.RequestTime).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ImageBlockResponse) String() string {
	return zcl.Sprintf(
		"ImageBlockResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
			"FileOffset(%v)",
			"Payload(%v)",
			"CurrentTime(%v)",
			"RequestTime(%v)",
		}, " ")+"}",
		v.Status,
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
		v.FileOffset,
		v.Payload,
		v.CurrentTime,
		v.RequestTime,
	)
}

type UpgradeEndResponse struct {
	ManufacturerCode ManufacturerCode
	ImageType        ImageType
	FileVersion      FileVersion
	CurrentTime      CurrentTime
	RequestTime      RequestTime
}

type UpgradeEndResponseHandler interface {
	HandleUpgradeEndResponse(frame Frame, cmd *UpgradeEndResponse) error
}

// UpgradeEndResponseCommand is the Command ID of UpgradeEndResponse
const UpgradeEndResponseCommand CommandID = 0x0007

// Values returns all values of UpgradeEndResponse
func (v *UpgradeEndResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.ManufacturerCode,
		&v.ImageType,
		&v.FileVersion,
		&v.CurrentTime,
		&v.RequestTime,
	}
}

// Arguments returns all values of UpgradeEndResponse
func (v *UpgradeEndResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "FileVersion", Argument: &v.FileVersion},
		{Name: "CurrentTime", Argument: &v.CurrentTime},
		{Name: "RequestTime", Argument: &v.RequestTime},
	}
}

// Name of the command
func (UpgradeEndResponse) Name() string { return `Upgrade End Response` }

// Description of the command
func (UpgradeEndResponse) Description() string { return `` }

// ID of the command
func (UpgradeEndResponse) ID() CommandID { return UpgradeEndResponseCommand }

// Required
func (UpgradeEndResponse) Required() bool { return true }

// Cluster ID of the command
func (UpgradeEndResponse) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (UpgradeEndResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (UpgradeEndResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UpgradeEndResponse) MarshalJSON() ([]byte, error) { return []byte("7"), nil }

func (v *UpgradeEndResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h UpgradeEndResponseHandler
	if h, found = handler.(UpgradeEndResponseHandler); found {
		err = h.HandleUpgradeEndResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of UpgradeEndResponse
func (v UpgradeEndResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.CurrentTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RequestTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UpgradeEndResponse struct
func (v *UpgradeEndResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CurrentTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RequestTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UpgradeEndResponse) String() string {
	return zcl.Sprintf(
		"UpgradeEndResponse{"+zcl.StrJoin([]string{
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"FileVersion(%v)",
			"CurrentTime(%v)",
			"RequestTime(%v)",
		}, " ")+"}",
		v.ManufacturerCode,
		v.ImageType,
		v.FileVersion,
		v.CurrentTime,
		v.RequestTime,
	)
}

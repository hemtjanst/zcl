package otau

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

func (ApplicationBuild) Name() string        { return `Application Build` }
func (ApplicationBuild) Description() string { return `` }

type ApplicationBuild zcl.Zu8

func (v *ApplicationBuild) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ApplicationBuild) Value() zcl.Val     { return v }

func (v ApplicationBuild) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ApplicationBuild) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ApplicationBuild(*nt)
	return br, err
}

func (v ApplicationBuild) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ApplicationBuild) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApplicationBuild(*a)
	return nil
}

func (v *ApplicationBuild) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ApplicationBuild(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApplicationBuild) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (ApplicationRelease) Name() string        { return `Application Release` }
func (ApplicationRelease) Description() string { return `` }

type ApplicationRelease zcl.Zu8

func (v *ApplicationRelease) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ApplicationRelease) Value() zcl.Val     { return v }

func (v ApplicationRelease) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ApplicationRelease) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ApplicationRelease(*nt)
	return br, err
}

func (v ApplicationRelease) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ApplicationRelease) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApplicationRelease(*a)
	return nil
}

func (v *ApplicationRelease) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ApplicationRelease(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApplicationRelease) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (BlockRequestOptions) Name() string        { return `Block Request Options` }
func (BlockRequestOptions) Description() string { return `` }

type BlockRequestOptions zcl.Zbmp8

func (v *BlockRequestOptions) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *BlockRequestOptions) Value() zcl.Val     { return v }

func (v BlockRequestOptions) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *BlockRequestOptions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = BlockRequestOptions(*nt)
	return br, err
}

func (v BlockRequestOptions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *BlockRequestOptions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BlockRequestOptions(*a)
	return nil
}

func (v *BlockRequestOptions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = BlockRequestOptions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BlockRequestOptions) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 1:
			bstr = append(bstr, "Request Node Address")
		case 2:
			bstr = append(bstr, "Minimum Block Period")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v BlockRequestOptions) IsRequestNodeAddress() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v BlockRequestOptions) IsMinimumBlockPeriod() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v *BlockRequestOptions) SetRequestNodeAddress(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *BlockRequestOptions) SetMinimumBlockPeriod(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}

func (BlockRequestOptions) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 1, Name: "Request Node Address"},
		{Value: 2, Name: "Minimum Block Period"},
	}
}

const CurrentFileVersionAttr zcl.AttrID = 2

func (CurrentFileVersion) ID() zcl.AttrID   { return CurrentFileVersionAttr }
func (CurrentFileVersion) Readable() bool   { return true }
func (CurrentFileVersion) Writable() bool   { return false }
func (CurrentFileVersion) Reportable() bool { return false }
func (CurrentFileVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentFileVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentFileVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentFileVersion) AttrValue() zcl.Val  { return v.Value() }

func (CurrentFileVersion) Name() string        { return `Current File Version` }
func (CurrentFileVersion) Description() string { return `` }

type CurrentFileVersion zcl.Zu32

func (v *CurrentFileVersion) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *CurrentFileVersion) Value() zcl.Val     { return v }

func (v CurrentFileVersion) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *CurrentFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentFileVersion(*nt)
	return br, err
}

func (v CurrentFileVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *CurrentFileVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentFileVersion(*a)
	return nil
}

func (v *CurrentFileVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = CurrentFileVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentFileVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const CurrentStackVersionAttr zcl.AttrID = 3

func (CurrentStackVersion) ID() zcl.AttrID   { return CurrentStackVersionAttr }
func (CurrentStackVersion) Readable() bool   { return true }
func (CurrentStackVersion) Writable() bool   { return false }
func (CurrentStackVersion) Reportable() bool { return false }
func (CurrentStackVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentStackVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentStackVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentStackVersion) AttrValue() zcl.Val  { return v.Value() }

func (CurrentStackVersion) Name() string        { return `Current Stack Version` }
func (CurrentStackVersion) Description() string { return `` }

type CurrentStackVersion zcl.Zu16

func (v *CurrentStackVersion) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *CurrentStackVersion) Value() zcl.Val     { return v }

func (v CurrentStackVersion) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *CurrentStackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentStackVersion(*nt)
	return br, err
}

func (v CurrentStackVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *CurrentStackVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentStackVersion(*a)
	return nil
}

func (v *CurrentStackVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = CurrentStackVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentStackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (CurrentTime) Name() string        { return `Current Time` }
func (CurrentTime) Description() string { return `` }

type CurrentTime zcl.Ztime

func (v *CurrentTime) TypeID() zcl.TypeID { return new(zcl.Ztime).TypeID() }
func (v *CurrentTime) Value() zcl.Val     { return v }

func (v CurrentTime) MarshalZcl() ([]byte, error) { return zcl.Ztime(v).MarshalZcl() }

func (v *CurrentTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Ztime)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentTime(*nt)
	return br, err
}

func (v CurrentTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Ztime(v))
}

func (v *CurrentTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Ztime)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentTime(*a)
	return nil
}

func (v *CurrentTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Ztime); ok {
		*v = CurrentTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentTime) String() string {
	return zcl.Sprintf("%v", zcl.Ztime(v))
}

func (DataSize) Name() string        { return `Data Size` }
func (DataSize) Description() string { return `` }

type DataSize zcl.Zu8

func (v *DataSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DataSize) Value() zcl.Val     { return v }

func (v DataSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DataSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DataSize(*nt)
	return br, err
}

func (v DataSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DataSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DataSize(*a)
	return nil
}

func (v *DataSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DataSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DataSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const DownloadedFileVersionAttr zcl.AttrID = 4

func (DownloadedFileVersion) ID() zcl.AttrID   { return DownloadedFileVersionAttr }
func (DownloadedFileVersion) Readable() bool   { return true }
func (DownloadedFileVersion) Writable() bool   { return false }
func (DownloadedFileVersion) Reportable() bool { return false }
func (DownloadedFileVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DownloadedFileVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v DownloadedFileVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DownloadedFileVersion) AttrValue() zcl.Val  { return v.Value() }

func (DownloadedFileVersion) Name() string        { return `Downloaded File Version` }
func (DownloadedFileVersion) Description() string { return `` }

type DownloadedFileVersion zcl.Zu32

func (v *DownloadedFileVersion) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *DownloadedFileVersion) Value() zcl.Val     { return v }

func (v DownloadedFileVersion) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *DownloadedFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = DownloadedFileVersion(*nt)
	return br, err
}

func (v DownloadedFileVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *DownloadedFileVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DownloadedFileVersion(*a)
	return nil
}

func (v *DownloadedFileVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = DownloadedFileVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DownloadedFileVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const DownloadedStackVersionAttr zcl.AttrID = 5

func (DownloadedStackVersion) ID() zcl.AttrID   { return DownloadedStackVersionAttr }
func (DownloadedStackVersion) Readable() bool   { return true }
func (DownloadedStackVersion) Writable() bool   { return false }
func (DownloadedStackVersion) Reportable() bool { return false }
func (DownloadedStackVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DownloadedStackVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v DownloadedStackVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DownloadedStackVersion) AttrValue() zcl.Val  { return v.Value() }

func (DownloadedStackVersion) Name() string        { return `Downloaded Stack Version` }
func (DownloadedStackVersion) Description() string { return `` }

type DownloadedStackVersion zcl.Zu16

func (v *DownloadedStackVersion) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DownloadedStackVersion) Value() zcl.Val     { return v }

func (v DownloadedStackVersion) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DownloadedStackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DownloadedStackVersion(*nt)
	return br, err
}

func (v DownloadedStackVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DownloadedStackVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DownloadedStackVersion(*a)
	return nil
}

func (v *DownloadedStackVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DownloadedStackVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DownloadedStackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const FileOffsetAttr zcl.AttrID = 1

func (FileOffset) ID() zcl.AttrID   { return FileOffsetAttr }
func (FileOffset) Readable() bool   { return true }
func (FileOffset) Writable() bool   { return false }
func (FileOffset) Reportable() bool { return false }
func (FileOffset) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FileOffset) AttrID() zcl.AttrID   { return v.ID() }
func (v FileOffset) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FileOffset) AttrValue() zcl.Val  { return v.Value() }

func (FileOffset) Name() string        { return `File Offset` }
func (FileOffset) Description() string { return `` }

type FileOffset zcl.Zu32

func (v *FileOffset) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *FileOffset) Value() zcl.Val     { return v }

func (v FileOffset) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *FileOffset) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = FileOffset(*nt)
	return br, err
}

func (v FileOffset) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *FileOffset) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FileOffset(*a)
	return nil
}

func (v *FileOffset) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = FileOffset(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FileOffset) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

func (HardwareVersion) Name() string        { return `Hardware Version` }
func (HardwareVersion) Description() string { return `` }

type HardwareVersion zcl.Zu16

func (v *HardwareVersion) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *HardwareVersion) Value() zcl.Val     { return v }

func (v HardwareVersion) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *HardwareVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = HardwareVersion(*nt)
	return br, err
}

func (v HardwareVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *HardwareVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HardwareVersion(*a)
	return nil
}

func (v *HardwareVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = HardwareVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HardwareVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (HardwareVersionPresent) Name() string        { return `Hardware Version Present` }
func (HardwareVersionPresent) Description() string { return `` }

type HardwareVersionPresent zcl.Zbool

func (v *HardwareVersionPresent) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *HardwareVersionPresent) Value() zcl.Val     { return v }

func (v HardwareVersionPresent) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *HardwareVersionPresent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = HardwareVersionPresent(*nt)
	return br, err
}

func (v HardwareVersionPresent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *HardwareVersionPresent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HardwareVersionPresent(*a)
	return nil
}

func (v *HardwareVersionPresent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = HardwareVersionPresent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HardwareVersionPresent) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

func (v HardwareVersionPresent) Scaled() float64 {
	return float64(v)
}

func (ImageNotifyPayloadType) Name() string        { return `Image Notify Payload Type` }
func (ImageNotifyPayloadType) Description() string { return `` }

type ImageNotifyPayloadType zcl.Zenum8

func (v *ImageNotifyPayloadType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ImageNotifyPayloadType) Value() zcl.Val     { return v }

func (v ImageNotifyPayloadType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ImageNotifyPayloadType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ImageNotifyPayloadType(*nt)
	return br, err
}

func (v ImageNotifyPayloadType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ImageNotifyPayloadType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ImageNotifyPayloadType(*a)
	return nil
}

func (v *ImageNotifyPayloadType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ImageNotifyPayloadType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ImageNotifyPayloadType) String() string {
	switch v {
	case 0x00:
		return "Query jitter"
	case 0x01:
		return "Query jitter and manufacturer code"
	case 0x02:
		return "Query jitter, manufacturer code, and image type"
	case 0x03:
		return "Query jitter, manufacturer code, image type, and new file version"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ImageNotifyPayloadType) IsQueryJitter() bool                             { return v == 0x00 }
func (v ImageNotifyPayloadType) IsQueryJitterAndManufacturerCode() bool          { return v == 0x01 }
func (v ImageNotifyPayloadType) IsQueryJitterManufacturerCodeAndImageType() bool { return v == 0x02 }
func (v ImageNotifyPayloadType) IsQueryJitterManufacturerCodeImageTypeAndNewFileVersion() bool {
	return v == 0x03
}
func (v *ImageNotifyPayloadType) SetQueryJitter()                             { *v = 0x00 }
func (v *ImageNotifyPayloadType) SetQueryJitterAndManufacturerCode()          { *v = 0x01 }
func (v *ImageNotifyPayloadType) SetQueryJitterManufacturerCodeAndImageType() { *v = 0x02 }
func (v *ImageNotifyPayloadType) SetQueryJitterManufacturerCodeImageTypeAndNewFileVersion() {
	*v = 0x03
}

func (ImageNotifyPayloadType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Query jitter"},
		{Value: 0x01, Name: "Query jitter and manufacturer code"},
		{Value: 0x02, Name: "Query jitter, manufacturer code, and image type"},
		{Value: 0x03, Name: "Query jitter, manufacturer code, image type, and new file version"},
	}
}

func (ImageSize) Name() string        { return `Image Size` }
func (ImageSize) Description() string { return `` }

type ImageSize zcl.Zu32

func (v *ImageSize) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *ImageSize) Value() zcl.Val     { return v }

func (v ImageSize) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *ImageSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = ImageSize(*nt)
	return br, err
}

func (v ImageSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *ImageSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ImageSize(*a)
	return nil
}

func (v *ImageSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = ImageSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ImageSize) String() string {
	return zcl.Bytes.Format(float64(v))
}

func (ImageType) Name() string        { return `Image type` }
func (ImageType) Description() string { return `` }

type ImageType zcl.Zenum16

func (v *ImageType) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (v *ImageType) Value() zcl.Val     { return v }

func (v ImageType) MarshalZcl() ([]byte, error) { return zcl.Zenum16(v).MarshalZcl() }

func (v *ImageType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*v = ImageType(*nt)
	return br, err
}

func (v ImageType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(v))
}

func (v *ImageType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ImageType(*a)
	return nil
}

func (v *ImageType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum16); ok {
		*v = ImageType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ImageType) String() string {
	switch v {
	case 0x0000:
		return "Specific image"
	case 0xFFC0:
		return "Security credential"
	case 0xFFC1:
		return "Configuration"
	case 0xFFC2:
		return "Log"
	case 0xFFFF:
		return "Wild card"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(v))
}

func (v ImageType) IsSpecificImage() bool      { return v == 0x0000 }
func (v ImageType) IsSecurityCredential() bool { return v == 0xFFC0 }
func (v ImageType) IsConfiguration() bool      { return v == 0xFFC1 }
func (v ImageType) IsLog() bool                { return v == 0xFFC2 }
func (v ImageType) IsWildCard() bool           { return v == 0xFFFF }
func (v *ImageType) SetSpecificImage()         { *v = 0x0000 }
func (v *ImageType) SetSecurityCredential()    { *v = 0xFFC0 }
func (v *ImageType) SetConfiguration()         { *v = 0xFFC1 }
func (v *ImageType) SetLog()                   { *v = 0xFFC2 }
func (v *ImageType) SetWildCard()              { *v = 0xFFFF }

func (ImageType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x0000, Name: "Specific image"},
		{Value: 0xFFC0, Name: "Security credential"},
		{Value: 0xFFC1, Name: "Configuration"},
		{Value: 0xFFC2, Name: "Log"},
		{Value: 0xFFFF, Name: "Wild card"},
	}
}

const ImageUpgradeStatusAttr zcl.AttrID = 6

func (ImageUpgradeStatus) ID() zcl.AttrID   { return ImageUpgradeStatusAttr }
func (ImageUpgradeStatus) Readable() bool   { return true }
func (ImageUpgradeStatus) Writable() bool   { return true }
func (ImageUpgradeStatus) Reportable() bool { return false }
func (ImageUpgradeStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ImageUpgradeStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v ImageUpgradeStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ImageUpgradeStatus) AttrValue() zcl.Val  { return v.Value() }

func (ImageUpgradeStatus) Name() string        { return `Image upgrade status` }
func (ImageUpgradeStatus) Description() string { return `` }

type ImageUpgradeStatus zcl.Zenum8

func (v *ImageUpgradeStatus) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ImageUpgradeStatus) Value() zcl.Val     { return v }

func (v ImageUpgradeStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ImageUpgradeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ImageUpgradeStatus(*nt)
	return br, err
}

func (v ImageUpgradeStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ImageUpgradeStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ImageUpgradeStatus(*a)
	return nil
}

func (v *ImageUpgradeStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ImageUpgradeStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ImageUpgradeStatus) String() string {
	switch v {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Download in progress"
	case 0x02:
		return "Download complete"
	case 0x03:
		return "Waiting to upgrade"
	case 0x04:
		return "Count down"
	case 0x05:
		return "Wait for more"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ImageUpgradeStatus) IsNormal() bool             { return v == 0x00 }
func (v ImageUpgradeStatus) IsDownloadInProgress() bool { return v == 0x01 }
func (v ImageUpgradeStatus) IsDownloadComplete() bool   { return v == 0x02 }
func (v ImageUpgradeStatus) IsWaitingToUpgrade() bool   { return v == 0x03 }
func (v ImageUpgradeStatus) IsCountDown() bool          { return v == 0x04 }
func (v ImageUpgradeStatus) IsWaitForMore() bool        { return v == 0x05 }
func (v *ImageUpgradeStatus) SetNormal()                { *v = 0x00 }
func (v *ImageUpgradeStatus) SetDownloadInProgress()    { *v = 0x01 }
func (v *ImageUpgradeStatus) SetDownloadComplete()      { *v = 0x02 }
func (v *ImageUpgradeStatus) SetWaitingToUpgrade()      { *v = 0x03 }
func (v *ImageUpgradeStatus) SetCountDown()             { *v = 0x04 }
func (v *ImageUpgradeStatus) SetWaitForMore()           { *v = 0x05 }

func (ImageUpgradeStatus) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Normal"},
		{Value: 0x01, Name: "Download in progress"},
		{Value: 0x02, Name: "Download complete"},
		{Value: 0x03, Name: "Waiting to upgrade"},
		{Value: 0x04, Name: "Count down"},
		{Value: 0x05, Name: "Wait for more"},
	}
}

func (ManufacturerCode) Name() string        { return `Manufacturer Code` }
func (ManufacturerCode) Description() string { return `` }

type ManufacturerCode zcl.Zenum16

func (v *ManufacturerCode) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (v *ManufacturerCode) Value() zcl.Val     { return v }

func (v ManufacturerCode) MarshalZcl() ([]byte, error) { return zcl.Zenum16(v).MarshalZcl() }

func (v *ManufacturerCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*v = ManufacturerCode(*nt)
	return br, err
}

func (v ManufacturerCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(v))
}

func (v *ManufacturerCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ManufacturerCode(*a)
	return nil
}

func (v *ManufacturerCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum16); ok {
		*v = ManufacturerCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ManufacturerCode) String() string {
	return zcl.Sprintf("%v", zcl.Zenum16(v))
}

const MinBlockRequestDelayAttr zcl.AttrID = 9

func (MinBlockRequestDelay) ID() zcl.AttrID   { return MinBlockRequestDelayAttr }
func (MinBlockRequestDelay) Readable() bool   { return true }
func (MinBlockRequestDelay) Writable() bool   { return false }
func (MinBlockRequestDelay) Reportable() bool { return false }
func (MinBlockRequestDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinBlockRequestDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v MinBlockRequestDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinBlockRequestDelay) AttrValue() zcl.Val  { return v.Value() }

func (MinBlockRequestDelay) Name() string        { return `Min block request delay` }
func (MinBlockRequestDelay) Description() string { return `` }

type MinBlockRequestDelay zcl.Zu16

func (v *MinBlockRequestDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinBlockRequestDelay) Value() zcl.Val     { return v }

func (v MinBlockRequestDelay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinBlockRequestDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinBlockRequestDelay(*nt)
	return br, err
}

func (v MinBlockRequestDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinBlockRequestDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinBlockRequestDelay(*a)
	return nil
}

func (v *MinBlockRequestDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinBlockRequestDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinBlockRequestDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (NewFileVersion) Name() string        { return `New File Version` }
func (NewFileVersion) Description() string { return `` }

type NewFileVersion zcl.Zu32

func (v *NewFileVersion) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *NewFileVersion) Value() zcl.Val     { return v }

func (v NewFileVersion) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *NewFileVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = NewFileVersion(*nt)
	return br, err
}

func (v NewFileVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *NewFileVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NewFileVersion(*a)
	return nil
}

func (v *NewFileVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = NewFileVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NewFileVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

func (NextImageStatus) Name() string        { return `Next Image Status` }
func (NextImageStatus) Description() string { return `` }

type NextImageStatus zcl.Zenum8

func (v *NextImageStatus) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *NextImageStatus) Value() zcl.Val     { return v }

func (v NextImageStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *NextImageStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = NextImageStatus(*nt)
	return br, err
}

func (v NextImageStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *NextImageStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NextImageStatus(*a)
	return nil
}

func (v *NextImageStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = NextImageStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NextImageStatus) String() string {
	switch v {
	case 0x00:
		return "Success"
	case 0x98:
		return "No updates"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v NextImageStatus) IsSuccess() bool   { return v == 0x00 }
func (v NextImageStatus) IsNoUpdates() bool { return v == 0x98 }
func (v *NextImageStatus) SetSuccess()      { *v = 0x00 }
func (v *NextImageStatus) SetNoUpdates()    { *v = 0x98 }

func (NextImageStatus) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Success"},
		{Value: 0x98, Name: "No updates"},
	}
}

func (Payload) Name() string        { return `Payload` }
func (Payload) Description() string { return `` }

type Payload zcl.Zostring

func (v *Payload) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (v *Payload) Value() zcl.Val     { return v }

func (v Payload) MarshalZcl() ([]byte, error) { return zcl.Zostring(v).MarshalZcl() }

func (v *Payload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*v = Payload(*nt)
	return br, err
}

func (v Payload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(v))
}

func (v *Payload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zostring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Payload(*a)
	return nil
}

func (v *Payload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zostring); ok {
		*v = Payload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Payload) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(v))
}

func (QueryJitter) Name() string        { return `Query jitter` }
func (QueryJitter) Description() string { return `` }

type QueryJitter zcl.Zu8

func (v *QueryJitter) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *QueryJitter) Value() zcl.Val     { return v }

func (v QueryJitter) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *QueryJitter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = QueryJitter(*nt)
	return br, err
}

func (v QueryJitter) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *QueryJitter) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = QueryJitter(*a)
	return nil
}

func (v *QueryJitter) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = QueryJitter(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v QueryJitter) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (RequestTime) Name() string        { return `Request Time` }
func (RequestTime) Description() string { return `` }

type RequestTime zcl.Ztime

func (v *RequestTime) TypeID() zcl.TypeID { return new(zcl.Ztime).TypeID() }
func (v *RequestTime) Value() zcl.Val     { return v }

func (v RequestTime) MarshalZcl() ([]byte, error) { return zcl.Ztime(v).MarshalZcl() }

func (v *RequestTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Ztime)
	br, err := nt.UnmarshalZcl(b)
	*v = RequestTime(*nt)
	return br, err
}

func (v RequestTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Ztime(v))
}

func (v *RequestTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Ztime)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RequestTime(*a)
	return nil
}

func (v *RequestTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Ztime); ok {
		*v = RequestTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RequestTime) String() string {
	return zcl.Sprintf("%v", zcl.Ztime(v))
}

func (StackBuild) Name() string        { return `Stack Build` }
func (StackBuild) Description() string { return `` }

type StackBuild zcl.Zu8

func (v *StackBuild) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *StackBuild) Value() zcl.Val     { return v }

func (v StackBuild) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *StackBuild) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = StackBuild(*nt)
	return br, err
}

func (v StackBuild) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *StackBuild) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StackBuild(*a)
	return nil
}

func (v *StackBuild) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = StackBuild(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StackBuild) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (StackRelease) Name() string        { return `Stack Release` }
func (StackRelease) Description() string { return `` }

type StackRelease zcl.Zu8

func (v *StackRelease) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *StackRelease) Value() zcl.Val     { return v }

func (v StackRelease) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *StackRelease) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = StackRelease(*nt)
	return br, err
}

func (v StackRelease) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *StackRelease) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StackRelease(*a)
	return nil
}

func (v *StackRelease) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = StackRelease(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StackRelease) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (Status) Name() string        { return `Status` }
func (Status) Description() string { return `` }

type Status zcl.Zenum8

func (v *Status) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *Status) Value() zcl.Val     { return v }

func (v Status) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = Status(*nt)
	return br, err
}

func (v Status) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *Status) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Status(*a)
	return nil
}

func (v *Status) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = Status(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Status) String() string {
	switch v {
	case 0x00:
		return "Success"
	case 0x97:
		return "Wait for data"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v Status) IsSuccess() bool     { return v == 0x00 }
func (v Status) IsWaitForData() bool { return v == 0x97 }
func (v *Status) SetSuccess()        { *v = 0x00 }
func (v *Status) SetWaitForData()    { *v = 0x97 }

func (Status) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Success"},
		{Value: 0x97, Name: "Wait for data"},
	}
}

const UpgradeServerAttr zcl.AttrID = 0

func (UpgradeServer) ID() zcl.AttrID   { return UpgradeServerAttr }
func (UpgradeServer) Readable() bool   { return true }
func (UpgradeServer) Writable() bool   { return true }
func (UpgradeServer) Reportable() bool { return false }
func (UpgradeServer) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UpgradeServer) AttrID() zcl.AttrID   { return v.ID() }
func (v UpgradeServer) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UpgradeServer) AttrValue() zcl.Val  { return v.Value() }

func (UpgradeServer) Name() string        { return `Upgrade server` }
func (UpgradeServer) Description() string { return `` }

type UpgradeServer zcl.Zuid

func (v *UpgradeServer) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (v *UpgradeServer) Value() zcl.Val     { return v }

func (v UpgradeServer) MarshalZcl() ([]byte, error) { return zcl.Zuid(v).MarshalZcl() }

func (v *UpgradeServer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*v = UpgradeServer(*nt)
	return br, err
}

func (v UpgradeServer) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(v))
}

func (v *UpgradeServer) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zuid)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UpgradeServer(*a)
	return nil
}

func (v *UpgradeServer) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zuid); ok {
		*v = UpgradeServer(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UpgradeServer) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(v))
}

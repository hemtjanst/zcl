package tests

import (
	"encoding/hex"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/foundation"
	"testing"
)

func TestFromString(t *testing.T) {
	v := zcl.FromString(40, "-20")
	if vv, ok := v.(*zcl.Zs8); !ok {
		t.Errorf("Wrong type (%+v)", v)
	} else if vv == nil {
		t.Errorf("Result is nil")
	} else if *vv != -20 {
		t.Errorf("Wrong value (%d)", *vv)
	}

	if b, e := v.MarshalZcl(); e != nil {
		t.Error(e)
	} else if len(b) != 1 || b[0] != 0xEC {
		t.Errorf("Marshal failed, expected 0xEC, got 0x%02X", b)
	}

	asr := &foundation.ReadAttributeStatusRecord{}
	b, e := hex.DecodeString("10000028EE")
	if e != nil {
		t.Error(e)
		return
	}
	if b, e = asr.UnmarshalZcl(b); e != nil {
		t.Error(e)
	} else if len(b) > 0 {
		t.Errorf("remaining bytes: %x", b)
	} else if asr.AttributeID != 0x0010 {
		t.Errorf("wrong attributeid (0x0010 != 0x%04X", asr.AttributeID)
	} else if asr.DataType != 40 {
		t.Errorf("wrong data type (40 != %d", asr.DataType)
	} else if vv, ok := asr.Value.(*zcl.Zs8); !ok {
		t.Errorf("wrong value type (s8 != %+v)", asr.Value)
	} else if vv == nil {
		t.Errorf("value is nil %+v", asr)
	} else if *vv != -18 {
		t.Errorf("expected -18, got %d", *vv)
	}

}

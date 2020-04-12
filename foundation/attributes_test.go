package foundation

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"hemtjan.st/zcl/utils/testutils"
	"log"
	"testing"
)

func TestReadAttributesResponse_UnmarshalZcl(t *testing.T) {

	// [NWK(4127)/1 -> NWK(0000)/1] Cluster[0000] Profile[0104] Cmd[01] Type[0x00] MnfCode[0x0000] Unable to parse command: not enough data
	// 0000002003010000202002000020620300002001040000420E494B4541206F662053776564656E05000042165452414446524920636F6E74726F6C206F75746C6574060000420700003001
	v := &ReadAttributesResponse{}

	data, _ := hex.DecodeString(
		"0000002003" + // 0x0000 = uint8(3)
			"0100002020" + // 0x0001 = uint8(32)
			"0200002062" + // 0x0002 = uint8(98)
			"0300002001" + // 0x0003 = uint8(1)
			"040000420E494B4541206F662053776564656E" + // 0x0004 = cstring(IKEA of Sweden)
			"05000042165452414446524920636F6E74726F6C206F75746C6574" + // 0x0005 = cstring(TRADFRI control outlet)
			"", // "060000420700003001",
	)

	data, err := v.UnmarshalZcl(data)

	for _, a := range v.Attributes {
		log.Printf("%s", a)
	}

	assert.Nil(t, err)
	assert.Empty(t, data)

}

func TestConfigureReportingResponse_UnmarshalZcl(t *testing.T) {
	v := &ConfigureReportingResponse{}
	testutils.Unmarshal(t, v, "08CA0700")
	log.Printf("%+v", v)
}

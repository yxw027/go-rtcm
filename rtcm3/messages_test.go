package rtcm3_test

import (
	"bufio"
	"fmt"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3/data"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

var (
	messages = map[int]rtcm3.Message{
		1001: data.Message1001, 1002: data.Message1002, 1003: data.Message1003, 1004: data.Message1004,
		1005: data.Message1005, 1006: data.Message1006, 1007: data.Message1007, 1008: data.Message1008,
		1009: data.Message1009, 1010: data.Message1010, 1011: data.Message1011, 1012: data.Message1012,
		1013: data.Message1013, 1019: data.Message1019, 1020: data.Message1020, 1029: data.Message1029,
		1033: data.Message1033, 1071: data.Message1071, 1072: data.Message1072, 1073: data.Message1073,
		1074: data.Message1074, 1075: data.Message1075, 1076: data.Message1076, 1077: data.Message1077,
		1081: data.Message1081, 1082: data.Message1082, 1083: data.Message1083, 1084: data.Message1084,
		1085: data.Message1085, 1086: data.Message1086, 1087: data.Message1087, 1091: data.Message1091,
		1092: data.Message1092, 1093: data.Message1093, 1094: data.Message1094, 1095: data.Message1095,
		1096: data.Message1096, 1097: data.Message1097, 1111: data.Message1111, 1112: data.Message1112,
		1113: data.Message1113, 1114: data.Message1114, 1115: data.Message1115, 1116: data.Message1116,
		1117: data.Message1117, 1121: data.Message1121, 1122: data.Message1122, 1123: data.Message1123,
		1124: data.Message1124, 1125: data.Message1125, 1126: data.Message1126, 1127: data.Message1127,
	}
)

func readPayload(msgNumber uint) (payload []byte) {
	r, _ := os.Open("data/" + fmt.Sprint(msgNumber) + "_frame.bin")
	br := bufio.NewReader(r)
	frame, _ := rtcm3.DeserializeFrame(br)

	return frame.Payload
}

func TestSerializeDeserialize(t *testing.T) {
	for number, message := range messages {
		binary := readPayload(uint(number))
		if !cmp.Equal(message.Serialize(), binary) {
			t.Errorf("Serialization not equal to binary")
		}

		deserializedMessage := rtcm3.DeserializeMessage(message.Serialize())
		if !cmp.Equal(message, deserializedMessage) {
			t.Errorf("Serialization->Deserialization not equal")
		}

		if !cmp.Equal(deserializedMessage.Number(), number) {
			t.Errorf("MessageNumber incorrect")
		}
	}
}

func TestFrame(t *testing.T) {
	r, _ := os.Open("data/1117_frame.bin")
	br := bufio.NewReader(r)

	binary, _ := br.Peek(227)
	deserializedBinary, _ := rtcm3.DeserializeFrame(br)

	frame := rtcm3.Frame{
		Preamble: 211,
		Reserved: 0,
		Length:   121,
		Payload: []byte{
			0x45, 0xd0, 0x0, 0x6a, 0x9c, 0x8a, 0xa0, 0x0, 0x0, 0x71, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x80, 0x0, 0x7f, 0xc1, 0x47,
			0x37, 0xbe, 0x7f, 0xfe, 0x48, 0x14, 0xb9, 0xc1, 0xf0, 0x7c, 0xc4,
			0xa, 0xf8, 0x14, 0xa0, 0x0, 0x12, 0xbc, 0xa1, 0x2c, 0x23, 0x89,
			0x33, 0x28, 0x9c, 0x6c, 0x7e, 0x2, 0xb7, 0xe1, 0x55, 0x77, 0xf5,
			0x8f, 0x81, 0x99, 0x4, 0xdc, 0xda, 0x5, 0x32, 0x48, 0x82, 0x73,
			0x86, 0x2, 0x1, 0x80, 0xff, 0xa8, 0x62, 0xff, 0x9d, 0x49, 0xfd,
			0xfe, 0x21, 0x7d, 0xe0, 0xa4, 0x4c, 0x96, 0x5, 0x81, 0x60, 0x58,
			0x16, 0x5, 0x81, 0x60, 0x0, 0x51, 0x93, 0xa4, 0x61, 0x19, 0xe4,
			0x19, 0xbc, 0xcf, 0x66, 0x86, 0xc2, 0xc, 0x45, 0xfc, 0xf3, 0xf6,
			0x20, 0xc1, 0xf1, 0x80, 0x80, 0x12, 0x40, 0x14, 0x0,
		},
		Crc: 0xfaf141,
	}

	if !cmp.Equal(frame.Serialize(), binary) {
		t.Errorf("Frame serialization and binary not equal")
	}

	if !cmp.Equal(frame, deserializedBinary) {
		t.Errorf("Frame and deserialized not equal")
	}
}

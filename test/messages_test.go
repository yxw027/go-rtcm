package rtcm_test

import (
    "testing"
    "../pkgs/rtcm"
    "github.com/google/go-cmp/cmp"
    "os"
    "log"
    "bufio"
    "fmt"
    "./data"
)

func readFile(name string) (payload []byte){
    file, fileerr := os.Open(name)
    defer file.Close()
    if fileerr != nil {
        log.Fatal(fileerr)
    }
    fileinfo, staterr := file.Stat()
    if staterr != nil {
        log.Fatal(staterr)
    }
    buffer := make([]byte, fileinfo.Size())
    file.Read(buffer)
    return buffer
}

func readPayload(msgNumber uint) (payload []byte){
    r, _ := os.Open("data/" + fmt.Sprint(msgNumber) + "_frame.bin")
    br := bufio.NewReader(r)
    frame, _ := rtcm.DeserializeRtcm3Frame(br)

    return frame.Payload
}

func TestFrame(t *testing.T) {
    r, _ := os.Open("data/1117_frame.bin")
    br := bufio.NewReader(r)

    binary, _ := br.Peek(227)
    deserializedBinary, _ := rtcm.DeserializeRtcm3Frame(br)

    frame := rtcm.Rtcm3Frame{
        Preamble:211,
        Reserved:0,
        Length:121,
        Payload:[]byte{
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
        Crc:0xfaf141,
    }

    if !cmp.Equal(frame.Serialize(), binary) {
        t.Errorf("Frame serialization and binary not equal")
    }

    if !cmp.Equal(frame, deserializedBinary) {
        t.Errorf("Frame and deserialized not equal")
    }
}

func TestMessage1004(t *testing.T) {
    payload := readPayload(1004)

    msg := data.Message1004
    deserializedMsg := rtcm.NewRtcm3Message1004(msg.Serialize())

    if !cmp.Equal(msg.Serialize(), payload) {
        t.Errorf("Serialization not equal to binary")
    }

    if !cmp.Equal(msg, deserializedMsg) {
        t.Errorf("Serialization->Deserialization not equal")
    }
}

func TestMessage1117(t *testing.T) {
    payload := readPayload(1117)
    msg := data.Message1117
    
    if !cmp.Equal(msg.Serialize(), payload) {
        t.Errorf("Serialization not equal to binary")
    }

    deserialized_msg := rtcm.NewRtcm3Message1117(msg.Serialize())

    if !cmp.Equal(msg, deserialized_msg) {
        t.Errorf("Serialization->Deserialization not equal")
    }
}

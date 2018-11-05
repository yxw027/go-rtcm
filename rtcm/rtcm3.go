package rtcm

import (
    "io"
    "bufio"
    "encoding/binary"
//    "fmt"
    "errors"
    "github.com/bamiaux/iobit"
//    "time"
)

type Rtcm3Message interface {
    Serialize() []byte
    Number() uint16
}

var Rtcm3FramePreamble byte = 0xD3

type Rtcm3Frame struct {
    Preamble uint8
    Reserved uint8
    Length uint16
    Payload []byte
    Crc uint32
}

func (f Rtcm3Frame) Number() uint16 {
    return binary.BigEndian.Uint16(f.Payload[0:2]) >> 4
}

func (f Rtcm3Frame) Serialize() []byte {
    data := make([]byte, f.Length + 6)
    w := iobit.NewWriter(data)
    w.PutUint8(8, f.Preamble)
    w.PutUint8(6, f.Reserved)
    w.PutUint16(10, f.Length)
    w.Write(f.Payload)
    w.PutUint32(24, f.Crc)
    w.Flush()
    return data
}

func Deserialize(reader *bufio.Reader) (msg Rtcm3Message, err error) {
    // Only reads first byte from reader if Preamble or CRC are incorrect
    // Unfortunatly can't construct anything that will read bits (like iobit) until we have a byte array
    preamble, err := reader.ReadByte()
    if err != nil { return msg, err }
    if preamble != Rtcm3FramePreamble { return msg, errors.New("Invalid Preamble") }

    header, err := reader.Peek(2)
    if err != nil { return msg, err }

    reserved := uint8(header[0]) & 0xFC
    length := binary.BigEndian.Uint16(header) & 0x3FF
    data, err := reader.Peek(int(length + 5))
    if err != nil { return msg, err }

    data = append([]byte{preamble}, data...)
    crc := binary.BigEndian.Uint32(data[len(data)-4:]) & 0xFFFFFF

    message := Rtcm3Frame{
        Preamble: preamble,
        Reserved: reserved,
        Length: length,
        Payload: data[3:len(data)-3],
        Crc: crc,
    }

    if Crc24q(data[:len(data)-3]) != int(message.Crc) {
        return &message, errors.New("CRC Failed")
    }

    reader.Discard(len(data) - 1)

    switch message.Number() { //TODO: Figure out a better way to manage this switch statement - maybe a map from message number to constructor function
        case uint16(1071), uint16(1081), uint16(1091), uint16(1111), uint16(1121):
            message := NewRtcm3Msm1Message(message)
            return &message, nil

        case uint16(1072), uint16(1082), uint16(1092), uint16(1112), uint16(1122):
            message := NewRtcm3Msm2Message(message)
            return &message, nil

        case uint16(1073), uint16(1083), uint16(1093), uint16(1113), uint16(1123):
            message := NewRtcm3Msm3Message(message)
            return &message, nil

        case uint16(1074), uint16(1084), uint16(1094), uint16(1114), uint16(1124):
            message := NewRtcm3Msm4Message(message)
            return &message, nil

        case uint16(1075), uint16(1085), uint16(1095), uint16(1115), uint16(1125):
            message := NewRtcm3Msm5Message(message)
            return &message, nil

        case uint16(1076), uint16(1086), uint16(1096), uint16(1116), uint16(1126):
            message := NewRtcm3Msm6Message(message)
            return &message, nil

        case uint16(1077), uint16(1087), uint16(1097), uint16(1117), uint16(1127):
            message := NewRtcm3Msm7Message(message)
            return &message, nil
    }

    return &message, nil
}

type Callback func(Rtcm3Message)

func Scan(r io.Reader, callback Callback) (err error) {
    // Not sure if a function of this signature makes sense, or if we should just be writing back to an io object, or even if Deserialize should just be looping like this
    reader := bufio.NewReader(r)
    for {
        message, err := Deserialize(reader)
        if err != nil {
            if err.Error() == "Invalid Preamble" || err.Error() == "CRC Failed" { continue }
            return err
        }

        go callback(message)
    }
}

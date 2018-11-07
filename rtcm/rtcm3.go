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

    switch int(message.Number()) {
        case 1001:
            message := NewRtcm3Message1001(message)
            return &message, nil

        case 1002:
            message := NewRtcm3Message1002(message)
            return &message, nil

        case 1003:
            message := NewRtcm3Message1003(message)
            return &message, nil

        case 1004:
            message := NewRtcm3Message1004(message)
            return &message, nil

        case 1005:
            message := NewRtcm3Message1005(message)
            return &message, nil

        case 1006:
            message := NewRtcm3Message1006(message)
            return &message, nil

        case 1007:
            message := NewRtcm3Message1007(message)
            return &message, nil

        case 1008:
            message := NewRtcm3Message1008(message)
            return &message, nil

        case 1009:
            message := NewRtcm3Message1009(message)
            return &message, nil

        case 1010:
            message := NewRtcm3Message1010(message)
            return &message, nil

        case 1011:
            message := NewRtcm3Message1011(message)
            return &message, nil

        case 1012:
            message := NewRtcm3Message1012(message)
            return &message, nil

        case 1013:
            message := NewRtcm3Message1013(message)
            return &message, nil

        case 1019:
            message := NewRtcm3Message1019(message)
            return &message, nil

        case 1020:
            message := NewRtcm3Message1020(message)
            return &message, nil

        case 1029:
            message := NewRtcm3Message1029(message)
            return &message, nil

        case 1030:
            message := NewRtcm3Message1030(message)
            return &message, nil

        case 1031:
            message := NewRtcm3Message1031(message)
            return &message, nil

        case 1032:
            message := NewRtcm3Message1032(message)
            return &message, nil

        case 1033:
            message := NewRtcm3Message1033(message)
            return &message, nil

        case 1230:
            message := NewRtcm3Message1230(message)
            return &message, nil

        case 1071, 1081, 1091, 1111, 1121:
            message := NewRtcm3MessageMsm1(message)
            return &message, nil

        case 1072, 1082, 1092, 1112, 1122:
            message := NewRtcm3MessageMsm2(message)
            return &message, nil

        case 1073, 1083, 1093, 1113, 1123:
            message := NewRtcm3MessageMsm3(message)
            return &message, nil

        case 1074, 1084, 1094, 1114, 1124:
            message := NewRtcm3MessageMsm4(message)
            return &message, nil

        case 1075, 1085, 1095, 1115, 1125:
            message := NewRtcm3MessageMsm5(message)
            return &message, nil

        case 1076, 1086, 1096, 1116, 1126:
            message := NewRtcm3MessageMsm6(message)
            return &message, nil

        case 1077, 1087, 1097, 1117, 1127:
            message := NewRtcm3MessageMsm7(message)
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

// Sign-Magnitude Ints

func Sint8(r iobit.Reader, length int) int8 {
    n, v := r.Bit(), int8(r.Uint8(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func Sint16(r iobit.Reader, length int) int16 {
    n, v := r.Bit(), int16(r.Uint16(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func Sint32(r iobit.Reader, length int) int32 {
    n, v := r.Bit(), int32(r.Uint32(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

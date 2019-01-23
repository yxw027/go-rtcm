package rtcm

import (
    "io"
    "bufio"
    "encoding/binary"
    "errors"
    "github.com/bamiaux/iobit"
    "time"
)

type Rtcm3Message interface {
    Serialize() []byte
    Number() uint16
}

func NewRtcm3Message(payload []byte) (msg Rtcm3Message) {
    messageNumber := binary.BigEndian.Uint16(payload[0:2]) >> 4

    switch int(messageNumber) {
        case 1001: return NewRtcm3Message1001(payload)
        case 1002: return NewRtcm3Message1002(payload)
        case 1003: return NewRtcm3Message1003(payload)
        case 1004: return NewRtcm3Message1004(payload)
        case 1005: return NewRtcm3Message1005(payload)
        case 1006: return NewRtcm3Message1006(payload)
        case 1007: return NewRtcm3Message1007(payload)
        case 1008: return NewRtcm3Message1008(payload)
        case 1009: return NewRtcm3Message1009(payload)
        case 1010: return NewRtcm3Message1010(payload)
        case 1011: return NewRtcm3Message1011(payload)
        case 1012: return NewRtcm3Message1012(payload)
        case 1013: return NewRtcm3Message1013(payload)
        case 1019: return NewRtcm3Message1019(payload)
        case 1020: return NewRtcm3Message1020(payload)
        case 1029: return NewRtcm3Message1029(payload)
        case 1030: return NewRtcm3Message1030(payload)
        case 1031: return NewRtcm3Message1031(payload)
        case 1032: return NewRtcm3Message1032(payload)
        case 1033: return NewRtcm3Message1033(payload)
        case 1230: return NewRtcm3Message1230(payload)
        case 1071: return NewRtcm3Message1071(payload)
        case 1072: return NewRtcm3Message1072(payload)
        case 1073: return NewRtcm3Message1073(payload)
        case 1074: return NewRtcm3Message1074(payload)
        case 1075: return NewRtcm3Message1075(payload)
        case 1076: return NewRtcm3Message1076(payload)
        case 1077: return NewRtcm3Message1077(payload)
        case 1081: return NewRtcm3Message1081(payload)
        case 1082: return NewRtcm3Message1082(payload)
        case 1083: return NewRtcm3Message1083(payload)
        case 1084: return NewRtcm3Message1084(payload)
        case 1085: return NewRtcm3Message1085(payload)
        case 1086: return NewRtcm3Message1086(payload)
        case 1087: return NewRtcm3Message1087(payload)
        case 1091: return NewRtcm3Message1091(payload)
        case 1092: return NewRtcm3Message1092(payload)
        case 1093: return NewRtcm3Message1093(payload)
        case 1094: return NewRtcm3Message1094(payload)
        case 1095: return NewRtcm3Message1095(payload)
        case 1096: return NewRtcm3Message1096(payload)
        case 1097: return NewRtcm3Message1097(payload)
        case 1111: return NewRtcm3Message1111(payload)
        case 1112: return NewRtcm3Message1112(payload)
        case 1113: return NewRtcm3Message1113(payload)
        case 1114: return NewRtcm3Message1114(payload)
        case 1115: return NewRtcm3Message1115(payload)
        case 1116: return NewRtcm3Message1116(payload)
        case 1117: return NewRtcm3Message1117(payload)
        case 1121: return NewRtcm3Message1121(payload)
        case 1122: return NewRtcm3Message1122(payload)
        case 1123: return NewRtcm3Message1123(payload)
        case 1124: return NewRtcm3Message1124(payload)
        case 1125: return NewRtcm3Message1125(payload)
        case 1126: return NewRtcm3Message1126(payload)
        case 1127: return NewRtcm3Message1127(payload)
        default:
            return Rtcm3MessageUnknown{payload}
    }
}


type Rtcm3MessageUnknown struct {
    Payload []byte
}

func (msg Rtcm3MessageUnknown) Serialize() []byte {
    return msg.Payload
}

func (msg Rtcm3MessageUnknown) Number() (msgNumber uint16) {
    return binary.BigEndian.Uint16(msg.Payload[0:4]) >> 4
}


type Rtcm3Observable interface {
    Time() time.Time
}


var Rtcm3FramePreamble byte = 0xD3

type Rtcm3Frame struct { // Contains Serialized Rtcm3Message - Should not be used as Rtcm3Message
    Preamble uint8
    Reserved uint8
    Length uint16
    Payload []byte
    Crc uint32
}

// Encapsulate Rtcm3Message in Frame
func NewRtcm3Frame(msg Rtcm3Message) (frame Rtcm3Frame) {
    data := msg.Serialize()
    frame = Rtcm3Frame{
        Preamble: Rtcm3FramePreamble,
        Reserved: 0,
        Length: uint16(len(data)),
        Payload: data,
        Crc: uint32(0),
    }
    frame.Crc = Crc24q(frame.Serialize()[:len(data)+3])
    return frame
}

func (frame Rtcm3Frame) Serialize() []byte {
    data := make([]byte, frame.Length + 6)
    w := iobit.NewWriter(data)
    w.PutUint8(8, frame.Preamble)
    w.PutUint8(6, frame.Reserved)
    w.PutUint16(10, frame.Length)
    w.Write(frame.Payload)
    w.PutUint32(24, frame.Crc)
    w.Flush()
    return data
}

func DeserializeRtcm3Frame(reader *bufio.Reader) (frame Rtcm3Frame, err error) {
    // Only reads first byte from reader if Preamble or CRC are incorrect
    // Unfortunatly can't construct anything that will read bits (like iobit) until we have a byte array
    preamble, err := reader.ReadByte()
    if err != nil { return frame, err }
    if preamble != Rtcm3FramePreamble { return frame, errors.New("Invalid Preamble") }

    header, err := reader.Peek(2)
    if err != nil { return frame, err }

    reserved := uint8(header[0]) & 0xFC
    length := binary.BigEndian.Uint16(header) & 0x3FF
    data, err := reader.Peek(int(length + 5))
    if err != nil { return frame, err }

    data = append([]byte{preamble}, data...)
    crc := binary.BigEndian.Uint32(data[len(data)-4:]) & 0xFFFFFF

    frame = Rtcm3Frame{
        Preamble: preamble,
        Reserved: reserved,
        Length: length,
        Payload: data[3:len(data)-3],
        Crc: crc,
    }

    if Crc24q(data[:len(data)-3]) != frame.Crc {
        return frame, errors.New("CRC Failed")
    }

    reader.Discard(len(data) - 1)
    return frame, nil
}


type Scanner struct {
    Reader *bufio.Reader
}

func NewScanner(r io.Reader) Scanner {
    return Scanner{bufio.NewReader(r)}
}

func (scanner Scanner) Next() (message Rtcm3Message, err error) {
    for {
        frame, err := DeserializeRtcm3Frame(scanner.Reader)
        if err != nil {
            if err.Error() == "Invalid Preamble" || err.Error() == "CRC Failed" { continue }
            return nil, err
        }
        return NewRtcm3Message(frame.Payload), err // probably have frame.Message() return err
    }
}

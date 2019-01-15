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

func (frame Rtcm3Frame) Message() (msg Rtcm3Message) {
    messageNumber := binary.BigEndian.Uint16(frame.Payload[0:2]) >> 4

    switch int(messageNumber) {
        case 1001:
            return NewRtcm3Message1001(frame.Payload)

        case 1002:
            return NewRtcm3Message1002(frame.Payload)

        case 1003:
            return NewRtcm3Message1003(frame.Payload)

        case 1004:
            return NewRtcm3Message1004(frame.Payload)

        case 1005:
            return NewRtcm3Message1005(frame.Payload)

        case 1006:
            return NewRtcm3Message1006(frame.Payload)

        case 1007:
            return NewRtcm3Message1007(frame.Payload)

        case 1008:
            return NewRtcm3Message1008(frame.Payload)

        case 1009:
            return NewRtcm3Message1009(frame.Payload)

        case 1010:
            return NewRtcm3Message1010(frame.Payload)

        case 1011:
            return NewRtcm3Message1011(frame.Payload)

        case 1012:
            return NewRtcm3Message1012(frame.Payload)

        case 1013:
            return NewRtcm3Message1013(frame.Payload)

        case 1019:
            return NewRtcm3Message1019(frame.Payload)

        case 1020:
            return NewRtcm3Message1020(frame.Payload)

        case 1029:
            return NewRtcm3Message1029(frame.Payload)

        case 1030:
            return NewRtcm3Message1030(frame.Payload)

        case 1031:
            return NewRtcm3Message1031(frame.Payload)

        case 1032:
            return NewRtcm3Message1032(frame.Payload)

        case 1033:
            return NewRtcm3Message1033(frame.Payload)

        case 1230:
            return NewRtcm3Message1230(frame.Payload)

        case 1071:
            return NewRtcm3Message1071(frame.Payload)

        case 1072:
            return NewRtcm3Message1072(frame.Payload)

        case 1073:
            return NewRtcm3Message1073(frame.Payload)

        case 1074:
            return NewRtcm3Message1074(frame.Payload)

        case 1075:
            return NewRtcm3Message1075(frame.Payload)

        case 1076:
            return NewRtcm3Message1076(frame.Payload)

        case 1077:
            return NewRtcm3Message1077(frame.Payload)

        case 1081:
            return NewRtcm3Message1081(frame.Payload)

        case 1082:
            return NewRtcm3Message1082(frame.Payload)

        case 1083:
            return NewRtcm3Message1083(frame.Payload)

        case 1084:
            return NewRtcm3Message1084(frame.Payload)

        case 1085:
            return NewRtcm3Message1085(frame.Payload)

        case 1086:
            return NewRtcm3Message1086(frame.Payload)

        case 1087:
            return NewRtcm3Message1087(frame.Payload)

        case 1091:
            return NewRtcm3Message1091(frame.Payload)

        case 1092:
            return NewRtcm3Message1092(frame.Payload)

        case 1093:
            return NewRtcm3Message1093(frame.Payload)

        case 1094:
            return NewRtcm3Message1094(frame.Payload)

        case 1095:
            return NewRtcm3Message1095(frame.Payload)

        case 1096:
            return NewRtcm3Message1096(frame.Payload)

        case 1097:
            return NewRtcm3Message1097(frame.Payload)

        case 1111:
            return NewRtcm3Message1111(frame.Payload)

        case 1112:
            return NewRtcm3Message1112(frame.Payload)

        case 1113:
            return NewRtcm3Message1113(frame.Payload)

        case 1114:
            return NewRtcm3Message1114(frame.Payload)

        case 1115:
            return NewRtcm3Message1115(frame.Payload)

        case 1116:
            return NewRtcm3Message1116(frame.Payload)

        case 1117:
            return NewRtcm3Message1117(frame.Payload)

        case 1121:
            return NewRtcm3Message1121(frame.Payload)

        case 1122:
            return NewRtcm3Message1122(frame.Payload)

        case 1123:
            return NewRtcm3Message1123(frame.Payload)

        case 1124:
            return NewRtcm3Message1124(frame.Payload)

        case 1125:
            return NewRtcm3Message1125(frame.Payload)

        case 1126:
            return NewRtcm3Message1126(frame.Payload)

        case 1127:
            return NewRtcm3Message1127(frame.Payload)

        default:
            return Rtcm3MessageUnknown{frame.Payload}
    }
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
        return frame.Message(), err // probably have frame.Message() return err
    }
}

package rtcm3

import (
    "io"
    "bufio"
    "encoding/binary"
    "errors"
    "github.com/bamiaux/iobit"
    "time"
//    "os"
//    "fmt"
)

type Message interface {
    Serialize() []byte
    Number() uint16
}

func NewMessage(payload []byte) (msg Message) {
    messageNumber := binary.BigEndian.Uint16(payload[0:2]) >> 4

    switch int(messageNumber) {
        case 1001: return NewMessage1001(payload)
        case 1002: return NewMessage1002(payload)
        case 1003: return NewMessage1003(payload)
        case 1004: return NewMessage1004(payload)
        case 1005: return NewMessage1005(payload)
        case 1006: return NewMessage1006(payload)
        case 1007: return NewMessage1007(payload)
        case 1008: return NewMessage1008(payload)
        case 1009: return NewMessage1009(payload)
        case 1010: return NewMessage1010(payload)
        case 1011: return NewMessage1011(payload)
        case 1012: return NewMessage1012(payload)
        case 1013: return NewMessage1013(payload)
        case 1019: return NewMessage1019(payload)
        case 1020: return NewMessage1020(payload)
        case 1029: return NewMessage1029(payload)
        case 1030: return NewMessage1030(payload)
        case 1031: return NewMessage1031(payload)
        case 1032: return NewMessage1032(payload)
        case 1033: return NewMessage1033(payload)
        case 1230: return NewMessage1230(payload)
        case 1071: return NewMessage1071(payload)
        case 1072: return NewMessage1072(payload)
        case 1073: return NewMessage1073(payload)
        case 1074: return NewMessage1074(payload)
        case 1075: return NewMessage1075(payload)
        case 1076: return NewMessage1076(payload)
        case 1077: return NewMessage1077(payload)
        case 1081: return NewMessage1081(payload)
        case 1082: return NewMessage1082(payload)
        case 1083: return NewMessage1083(payload)
        case 1084: return NewMessage1084(payload)
        case 1085: return NewMessage1085(payload)
        case 1086: return NewMessage1086(payload)
        case 1087: return NewMessage1087(payload)
        case 1091: return NewMessage1091(payload)
        case 1092: return NewMessage1092(payload)
        case 1093: return NewMessage1093(payload)
        case 1094: return NewMessage1094(payload)
        case 1095: return NewMessage1095(payload)
        case 1096: return NewMessage1096(payload)
        case 1097: return NewMessage1097(payload)
        case 1101: return NewMessage1101(payload)
        case 1102: return NewMessage1102(payload)
        case 1103: return NewMessage1103(payload)
        case 1104: return NewMessage1104(payload)
        case 1105: return NewMessage1105(payload)
        case 1106: return NewMessage1106(payload)
        case 1107: return NewMessage1107(payload)
        case 1111: return NewMessage1111(payload)
        case 1112: return NewMessage1112(payload)
        case 1113: return NewMessage1113(payload)
        case 1114: return NewMessage1114(payload)
        case 1115: return NewMessage1115(payload)
        case 1116: return NewMessage1116(payload)
        case 1117: return NewMessage1117(payload)
        case 1121: return NewMessage1121(payload)
        case 1122: return NewMessage1122(payload)
        case 1123: return NewMessage1123(payload)
        case 1124: return NewMessage1124(payload)
        case 1125: return NewMessage1125(payload)
        case 1126: return NewMessage1126(payload)
        case 1127: return NewMessage1127(payload)
        default:
            return MessageUnknown{payload}
    }
}


type MessageUnknown struct {
    Payload []byte
}

func (msg MessageUnknown) Serialize() []byte {
    return msg.Payload
}

func (msg MessageUnknown) Number() (msgNumber uint16) {
    return binary.BigEndian.Uint16(msg.Payload[0:4]) >> 4
}


type Observable interface {
    Time() time.Time
}


var FramePreamble byte = 0xD3

type Frame struct { // Contains Serialized Message - Should not be used as Message
    Preamble uint8
    Reserved uint8
    Length uint16
    Payload []byte
    Crc uint32
}

// Encapsulate Message in Frame
func NewFrame(msg Message) (frame Frame) {
    data := msg.Serialize()
    frame = Frame{
        Preamble: FramePreamble,
        Reserved: 0,
        Length: uint16(len(data)),
        Payload: data,
        Crc: uint32(0),
    }
    frame.Crc = Crc24q(frame.Serialize()[:len(data)+3])
    return frame
}

func (frame Frame) Serialize() []byte {
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

func DeserializeFrame(reader *bufio.Reader) (frame Frame, err error) {
    // Only reads first byte from reader if Preamble or CRC are incorrect
    // Unfortunatly can't construct anything that will read bits (like iobit) until we have a byte array
    preamble, err := reader.ReadByte()
    if err != nil { return frame, err }
    if preamble != FramePreamble { return frame, errors.New("Invalid Preamble") }

    header, err := reader.Peek(2)
    if err != nil { return frame, err }

    reserved := uint8(header[0]) & 0xFC
    length := binary.BigEndian.Uint16(header) & 0x3FF
    data, err := reader.Peek(int(length + 5))
    if err != nil { return frame, err }

    data = append([]byte{preamble}, data...)
    crc := binary.BigEndian.Uint32(data[len(data)-4:]) & 0xFFFFFF

    frame = Frame{
        Preamble: preamble,
        Reserved: reserved,
        Length: length,
        Payload: data[3:len(data)-3],
        Crc: crc,
    }

    if Crc24q(data[:len(data)-3]) != frame.Crc {
        return frame, errors.New("CRC Failed")
    }

//    msg := NewMessage(frame.Payload)
    //fmt.Println(NewMessage(frame.Payload))
//    file, err := os.Create(fmt.Sprint(msg.Number()) + "_frame.bin")
//    defer file.Close()
//    file.Write(data[0:len(data)])

    reader.Discard(len(data) - 1)
    return frame, nil
}


type Scanner struct {
    Reader *bufio.Reader
}

func NewScanner(r io.Reader) Scanner {
    return Scanner{bufio.NewReader(r)}
}

func (scanner Scanner) Next() (message Message, err error) {
    for {
        frame, err := DeserializeFrame(scanner.Reader)
        if err != nil {
            if err.Error() == "Invalid Preamble" || err.Error() == "CRC Failed" { continue }
            return nil, err
        }
        return NewMessage(frame.Payload), err // probably have frame.Message() return err
    }
}

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

func (f *Rtcm3Frame) Number() uint16 {
    return binary.BigEndian.Uint16(f.Payload[0:2]) >> 4
}

func (f *Rtcm3Frame) Serialize() (data []byte) { // Pretty sure this isn't working
    w := iobit.NewWriter(data)
    w.PutUint8(8, f.Preamble)
    w.PutUint8(6, f.Reserved)
    w.PutUint16(10, f.Length)
    data = append(data, f.Payload...)
    crc := make([]byte, 4)
    binary.BigEndian.PutUint32(crc, f.Crc)
    return append(data, crc[1:]...)
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

    switch message.Number() {
        case uint16(1077):
            msg := NewRtcm31077(message)
            return &msg, nil
    }

    reader.Discard(len(data) - 1)
    return &message, nil
}

type Callback func(Rtcm3Message)

func Scan(r io.Reader, callback Callback) (err error) {
    // Not sure if a function of this signature makes sense, or if we should just be writing back to an io object
    reader := bufio.NewReader(r)
    for {
        message, err := Deserialize(reader)
        if err != nil {
            if err.Error() == "Invalid Preamble" || err.Error() == "CRC Failed" { continue }
            return err
        }

        callback(message)
    }
}

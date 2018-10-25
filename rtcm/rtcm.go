package rtcm

import (
    "io"
    "bufio"
    "encoding/binary"
    "fmt"
)

var Rtcm3FramePreamble byte = 0xD3

type Rtcm3Message struct {
    Reserved uint8
    Length uint16
    Payload []byte
    Crc uint32
}

func (f *Rtcm3Message) Serialize() (data []byte) {
    data = append(data, Rtcm3FramePreamble, byte(f.Reserved | uint8(f.Length>>8)), byte(f.Length & 0xFF))
    data = append(data, f.Payload...)
    crc := make([]byte, 4)
    binary.BigEndian.PutUint32(crc, f.Crc)
    return append(data, crc[1:]...)
}

func Deserialize(input io.Reader) (err error) {
    reader := bufio.NewReader(input)
    for preamble, err := reader.ReadByte(); err == nil; preamble, err = reader.ReadByte(){
        if preamble != Rtcm3FramePreamble { continue }

        header, _ := reader.Peek(2)
        reserved := header[0] & 0x3
        length := binary.BigEndian.Uint16(header) & 0x3FF

        data, _ := reader.Peek(int(length + 5))
        payload := data[2:][:length]
        crc := binary.BigEndian.Uint32(append([]byte{0x0}, data[2:][length:]...))

        if Crc24q(append([]byte{preamble}, data[:length+2]...)) != int(crc) {
            // Don't consume the buffer if the CRC check fails - continue from next byte after preamble
            continue
        }

        message := Rtcm3Message{reserved, length, payload, crc}
        reader.Discard(int(length + 5))

        fmt.Println(binary.BigEndian.Uint16(message.Payload[0:2]) >> 4)
    }
    return err
}

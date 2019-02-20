package rtcm3_test

import (
    "testing"
    "github.com/geoscienceaustralia/go-rtcm/rtcm3"
    "github.com/google/go-cmp/cmp"
    "bufio"
    "bytes"
    "reflect"
    "github.com/geoscienceaustralia/go-rtcm/rtcm3/data"
)

func TestMessageUnknown(t *testing.T) {
    messageInvalidNumber := data.Message1001
    messageInvalidNumber.GpsObservationHeader.MessageNumber = 1018  // doesn't exist, as far as I know

    unknownMessage := rtcm3.DeserializeMessage(messageInvalidNumber.Serialize())

    if unknownMessage.Number() != 1018 {
       t.Errorf("MessageUnknown message number incorrect")
    }

    if reflect.TypeOf(unknownMessage).Name() != "MessageUnknown" {
        t.Errorf("Deserialize incorrectly infers message type - should default to unknown")
    }

    if !cmp.Equal(messageInvalidNumber.Serialize(), unknownMessage.Serialize()) {
        t.Errorf("MessageUnknown serialization invalid")
    }
}

func TestInvalidPreamble(t *testing.T) {
    frameInvalidPreamble := rtcm3.EncapsulateMessage(data.Message1001)
    frameInvalidPreamble.Preamble = 0xd2  // should be 0xd3

    frameInvalidPreambleReader := bufio.NewReader(bytes.NewReader(frameInvalidPreamble.Serialize()))
    _, err := rtcm3.DeserializeFrame(frameInvalidPreambleReader)
    if err == nil || err.Error() != "Invalid Preamble" {
        t.Errorf("Did not catch invalid Preamble")
    }
}

func TestInvalidCrc(t *testing.T) {
    frameInvalidCrc := rtcm3.EncapsulateMessage(data.Message1001)
    frameInvalidCrc.Crc = 0xb3dfc9  // should be 0xb3dfd0

    frameInvalidCrcReader := bufio.NewReader(bytes.NewReader(frameInvalidCrc.Serialize()))
    _, err := rtcm3.DeserializeFrame(frameInvalidCrcReader)
    if err == nil || err.Error() != "CRC Failed" {
        t.Errorf("Did not catch invalid CRC")
    }
}

package rtcm3_test

import (
    "testing"
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
    "github.com/google/go-cmp/cmp"
    "bufio"
    "bytes"
    "reflect"
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3/data"
)

func TestMessageUnknown(t *testing.T) {
    messageInvalidNumber := data.Message1001
    messageInvalidNumber.GpsObservationHeader.MessageNumber = 0

    unknownMessage := rtcm3.DeserializeMessage(messageInvalidNumber.Serialize())    

    if reflect.TypeOf(unknownMessage).Name() != "MessageUnknown" {
        t.Errorf("Deserialize incorrectly infers message type - should default to unknown")
    }

    if !cmp.Equal(messageInvalidNumber.Serialize(), unknownMessage.Serialize()) {
        t.Errorf("MessageUnknown serialization invalid")
    }
}

func TestInvalidCrc(t *testing.T) {
    frameInvalidCrc := rtcm3.EncapsulateMessage(data.Message1001)
    frameInvalidCrc.Crc = 0  // should be 0xb3dfd0

    frameInvalidCrcReader := bufio.NewReader(bytes.NewReader(frameInvalidCrc.Serialize()))
    _, err := rtcm3.DeserializeFrame(frameInvalidCrcReader)
    if err == nil || err.Error() != "CRC Failed" {
        t.Errorf("Did not catch invalid CRC")
    }
}

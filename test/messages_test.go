package rtcm_test

import (
    "testing"
    "os"
    "bufio"
    "reflect"
    "../rtcm"
)

var messages = map[string]interface{}{
    "1004": &rtcm.Rtcm3Message1004{},
    "1006": &rtcm.Rtcm3Message1006{},
    "1012": &rtcm.Rtcm3Message1012{},
    "1033": &rtcm.Rtcm3Message1033{},
    "1077": &rtcm.Rtcm3Message1077{},
    "1087": &rtcm.Rtcm3Message1087{},
    "1117": &rtcm.Rtcm3Message1117{},
    "1127": &rtcm.Rtcm3Message1127{},
    "1230": &rtcm.Rtcm3Message1230{},
}

func TestMessageTypes(t *testing.T) {
    for msgNumber, msgType := range messages {
        r, _ := os.Open("data/" + msgNumber)
        br := bufio.NewReader(r)

        data, err := rtcm.Deserialize(br)
        if err != nil {
            t.Error("Failed to Deserialize " + msgNumber + "data:", err)
        }

        msg := msgType
        msg = data
        if reflect.TypeOf(msg) != reflect.TypeOf(msgType) {
            t.Error(reflect.TypeOf(msg), reflect.TypeOf(msgType))
        }
    }
}

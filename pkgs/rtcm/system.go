package rtcm

import (
    "github.com/bamiaux/iobit"
)

type Rtcm3MessageAnnouncement struct {
    Id uint16
    SyncFlag bool
    TransmissionInterval uint16
}

type Rtcm3Message1013 struct {
    MessageNumber uint16
    ReferenceStationId uint16
    Mjd uint16
    SecondsOfDay uint32
    MessageCount uint8
    LeapSeconds uint8
    Messages []Rtcm3MessageAnnouncement
}

func NewRtcm3Message1013(data []byte) (msg Rtcm3Message1013) {
    r := iobit.NewReader(data)
    msg = Rtcm3Message1013{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        Mjd: r.Uint16(16),
        SecondsOfDay: r.Uint32(17),
        MessageCount: r.Uint8(5),
        LeapSeconds: r.Uint8(8),
    }
    for i := 0; i < int(msg.MessageCount); i++ {
        msg.Messages = append(msg.Messages, Rtcm3MessageAnnouncement{
            Id: r.Uint16(12),
            SyncFlag: r.Bit(),
            TransmissionInterval: r.Uint16(16),
        })
    }
    return msg
}

func (msg Rtcm3Message1013) Serialize() (data []byte) {
    return data
}

type Rtcm3Message1029 struct {
    MessageNumber uint16
    ReferenceStationId uint16
    Mjd uint16
    SecondsOfDay uint32
    Characters uint8
    CodeUnitsLength uint8
    CodeUnits string
}

func NewRtcm3Message1029(data []byte) (msg Rtcm3Message1029) {
    r := iobit.NewReader(data)
    msg = Rtcm3Message1029{
        MessageNumber: r.Uint16(12),
        ReferenceStationId: r.Uint16(12),
        Mjd: r.Uint16(16),
        SecondsOfDay: r.Uint32(17),
        Characters: r.Uint8(7),
        CodeUnitsLength: r.Uint8(8),
    }
    msg.CodeUnits = r.String(8 * int(msg.CodeUnitsLength))
    return msg
}

func (msg Rtcm3Message1029) Serialize() (data []byte) {
    return data
}

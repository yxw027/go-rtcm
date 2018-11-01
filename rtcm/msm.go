package rtcm

import (
    "github.com/bamiaux/iobit"
    "math/bits"
)

type Rtcm3MsmHeader struct {
    MessageNumber uint16
    StationId uint16
    Epoch uint32
    MultipleMessageBit bool
    Iods uint8
    Reserved uint8
    ClockSteeringIndicator uint8
    ExternalClockIndicator uint8
    SmoothingIndicator bool
    SmoothingInterval uint8
    SatelliteMask uint64
    SignalMask uint32
    CellMask uint64
}

func NewRtcm3MsmHeader(r iobit.Reader) (msm Rtcm3MsmHeader) {
    msm = Rtcm3MsmHeader{
        MessageNumber: r.Uint16(12),
        StationId: r.Uint16(12),
        Epoch: r.Uint32(30),
        MultipleMessageBit: r.Bit(),
        Iods: r.Uint8(3),
        Reserved: r.Uint8(7),
        ClockSteeringIndicator: r.Uint8(2),
        ExternalClockIndicator: r.Uint8(2),
        SmoothingIndicator: r.Bit(),
        SmoothingInterval: r.Uint8(3),
        SatelliteMask: r.Uint64(64),
        SignalMask: r.Uint32(32),
    }

    cellMaskLength := bits.OnesCount(uint(msm.SignalMask)) * bits.OnesCount(uint(msm.SatelliteMask))
    msm.CellMask = r.Uint64(uint(cellMaskLength))
    return msm
}

type Rtcm31077 struct {
    Rtcm3Frame
    Header Rtcm3MsmHeader
}

func NewRtcm31077(msg Rtcm3Frame) Rtcm31077 {
    r := iobit.NewReader(msg.Payload)
    return Rtcm31077{
        Rtcm3Frame: msg,
        Header: NewRtcm3MsmHeader(r),
    }
}

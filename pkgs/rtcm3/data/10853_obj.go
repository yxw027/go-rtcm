package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message10853 = rtcm3.Message1085 {
    MessageMsm5: rtcm3.MessageMsm5 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x43d,
            ReferenceStationId: 0x0,
            Epoch: 0x19c8f090,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0xe000830000000000,
            SignalMask: 0x41000000,
            CellMask: 0xfff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x46, 0x40, 0x4b, 0x43, 0x4d, 0x43,
            },
            Extended: [] uint8 {
                0x8, 0x3, 0xc, 0xb, 0xa, 0x9,
            },
            Ranges: [] uint16 {
                0x21c, 0x124, 0x2f6, 0x6c, 0x28, 0x2f,
            },
            PhaseRangeRates: [] int16 {
                605, -245, -776, -399, 578, 175,
            },
        },
        SignalData: rtcm3.SignalDataMsm5 {
            Pseudoranges: [] int16 {
                -3895, -3564, 8045, 8279, 6928, 7227, -943, -730, -6323, -6003, 5361, 5504,
            },
            PhaseRanges: [] int32 {
                -123577, -113359, 267635, 283407, 226345, 234486, -25914, -9245, -200202, -180576, 174525, 178395,
            },
            PhaseRangeLocks: [] uint8 {
                0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint8 {
                0x2f, 0x2a, 0x35, 0x33, 0x2d, 0x2a, 0x35, 0x32, 0x2c, 0x24, 0x2f, 0x25,
            },
            PhaseRangeRates: [] int16 {
                605, 502, 2703, 2658, 4707, 4635, 4797, 4824, -1479, -1718, 3561, 3232,
            },
        },
    },
}

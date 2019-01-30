package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1086 = rtcm3.Message1086 {
    MessageMsm6: rtcm3.MessageMsm6 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x43e,
            ReferenceStationId: 0x0,
            Epoch: 0x29450ed8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x1,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0x7082e00000000000,
            SignalMask: 0x41000000,
            CellMask: 0xffff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm46 {
            RangeMilliseconds: [] uint8 {
                0x4a, 0x44, 0x4a, 0x48, 0x52, 0x4d, 0x44, 0x46,
            },
            Ranges: [] uint16 {
                0x3fe, 0x3ac, 0x19b, 0x29, 0x4b, 0x36c, 0x16b, 0x28,
            },
        },
        SignalData: rtcm3.SignalDataMsm6 {
            Pseudoranges: [] int32 {
                242720, 252760, 180701, 187458, -179354, -170239, -55762, -49799, 128225, 143748, -76620, -66742, 79429, 88214, -212219, -205757,
            },
            PhaseRanges: [] int32 {
                1021443, 1087199, 754404, 798558, -697878, -646240, -215741, -184518, 425889, 422570, -327908, -305766, 327079, 385443, -812875, -735570,
            },
            PhaseRangeLocks: [] uint16 {
                0x264, 0x264, 0x243, 0x243, 0x20e, 0x20e, 0x238, 0x238, 0x284, 0x284, 0x284, 0x284, 0x272, 0x272, 0x261, 0x261,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x2d3, 0x298, 0x331, 0x304, 0x2d7, 0x2c2, 0x2fd, 0x2e2, 0x26f, 0x28f, 0x2cc, 0x289, 0x328, 0x310, 0x2d6, 0x2c1,
            },
        },
    },
}

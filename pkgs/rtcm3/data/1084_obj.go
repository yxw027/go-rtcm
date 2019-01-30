package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1084 = rtcm3.Message1084 {
    MessageMsm4: rtcm3.MessageMsm4 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x43c,
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
        SignalData: rtcm3.SignalDataMsm4 {
            Pseudoranges: [] int16 {
                7585, 7899, 5647, 5858, -5605, -5320, -1743, -1556, 4007, 4492, -2394, -2086, 2482, 2757, -6632, -6430,
            },
            PhaseRanges: [] int32 {
                255361, 271800, 188601, 199639, -174469, -161560, -53935, -46129, 106472, 105642, -81977, -76441, 81770, 96361, -203219, -183892,
            },
            PhaseRangeLocks: [] uint8 {
                0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint8 {
                0x2d, 0x29, 0x33, 0x30, 0x2d, 0x2c, 0x30, 0x2e, 0x27, 0x29, 0x2d, 0x29, 0x32, 0x31, 0x2d, 0x2c,
            },
        },
    },
}

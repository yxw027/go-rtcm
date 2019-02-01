package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1095 = rtcm3.Message1095 {
    MessageMsm5: rtcm3.MessageMsm5 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x447,
            ReferenceStationId: 0x0,
            Epoch: 0x1a6055a8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0x823021c200000000,
            SignalMask: 0x40020200,
            CellMask: 0x7ffffff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x60, 0x5f, 0x56, 0x4a, 0x57, 0x4d, 0x5b, 0x59, 0x53,
            },
            Extended: [] uint8 {
                0xd, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0,
            },
            Ranges: [] uint16 {
                0xca, 0x59, 0x2fc, 0x2c8, 0x2b1, 0xac, 0xa1, 0x3da, 0x31d,
            },
            PhaseRangeRates: [] int16 {
                584, 260, -342, -120, -469, 21, -399, 367, 494,
            },
        },
        SignalData: rtcm3.SignalDataMsm5 {
            Pseudoranges: [] int16 {
                3515, 3696, 3886, -6068, -5997, -5815, 870, 900, 1082, -7734, -7882, -7749, 6735, 6737, 6885, 3951, 4581, 4743, -1888, -1783, -1629, 4986, 5243, 5357, -6840, -6747, -6591,
            },
            PhaseRanges: [] int32 {
                92379, 92069, 96171, -193396, -186780, -180513, 35081, 34849, 41621, -237025, -236395, -232545, 219304, 239578, 242727, 131509, 154660, 160117, -57259, -55851, -49200, 149854, 145781, 148666, -224776, -222853, -217791,
            },
            PhaseRangeLocks: [] uint8 {
                0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint8 {
                0x22, 0x24, 0x24, 0x24, 0x24, 0x23, 0x18, 0x23, 0x27, 0x2e, 0x31, 0x30, 0x25, 0x28, 0x27, 0x2e, 0x31, 0x30, 0x22, 0x27, 0x26, 0x1f, 0x26, 0x26, 0x2c, 0x2f, 0x2e,
            },
            PhaseRangeRates: [] int16 {
                -4304, -4007, -4122, -194, -177, -628, -902, -843, -742, -3752, -3812, -3874, -740, -1643, -1149, 4684, 4668, 4830, 4320, 4364, 4431, -438, 168, -186, -4849, -4669, -4626,
            },
        },
    },
}

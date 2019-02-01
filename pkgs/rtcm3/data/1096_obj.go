package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1096 = rtcm3.Message1096 {
    MessageMsm6: rtcm3.MessageMsm6 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x448,
            ReferenceStationId: 0x0,
            Epoch: 0x1a6055a8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x1,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0x823021c200000000,
            SignalMask: 0x40020200,
            CellMask: 0x7ffffff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm46 {
            RangeMilliseconds: [] uint8 {
                0x5f, 0x5e, 0x56, 0x4a, 0x57, 0x4c, 0x5a, 0x59, 0x53,
            },
            Ranges: [] uint16 {
                0x398, 0x326, 0x1ca, 0x195, 0x17f, 0x37a, 0x36f, 0x2a8, 0x1ea,
            },
        },
        SignalData: rtcm3.SignalDataMsm6 {
            Pseudoranges: [] int32 {
                38364, 44167, 50226, 255829, 258104, 263907, -46767, -45810, -39975, 202317, 197584, 201823, 140842, 140915, 145631, 52021, 72176, 77361, -135060, -131689, -126780, 85307, 93546, 97191, 231253, 234223, 239207,
            },
            PhaseRanges: [] int32 {
                73064, 71825, 88231, 1026421, 1052884, 1077952, -158111, -159039, -131953, 851093, 853611, 869012, 578508, 659603, 672201, 228379, 320984, 342810, -527594, -521961, -495355, 302500, 286208, 297746, 901401, 909094, 929342,
            },
            PhaseRangeLocks: [] uint16 {
                0x2a0, 0x2a0, 0x2a0, 0x25e, 0x25e, 0x25e, 0x228, 0x229, 0x229, 0x275, 0x275, 0x276, 0x21c, 0x246, 0x245, 0x268, 0x268, 0x268, 0x21c, 0x21d, 0x21e, 0x297, 0x297, 0x297, 0x287, 0x287, 0x287,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x21c, 0x247, 0x247, 0x244, 0x23f, 0x22a, 0x17c, 0x22a, 0x269, 0x2e1, 0x309, 0x2fa, 0x24f, 0x282, 0x26d, 0x2e3, 0x30b, 0x301, 0x226, 0x26c, 0x264, 0x1f1, 0x259, 0x25a, 0x2ba, 0x2f3, 0x2e4,
            },
        },
    },
}

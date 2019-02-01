package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1094 = rtcm3.Message1094 {
    MessageMsm4: rtcm3.MessageMsm4 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x446,
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
        SignalData: rtcm3.SignalDataMsm4 {
            Pseudoranges: [] int16 {
                1199, 1380, 1570, 7995, 8066, 8247, -1461, -1432, -1249, 6322, 6174, 6307, 4401, 4404, 4551, 1626, 2256, 2418, -4221, -4115, -3962, 2666, 2923, 3037, 7227, 7319, 7475,
            },
            PhaseRanges: [] int32 {
                18266, 17956, 22058, 256605, 263221, 269488, -39528, -39760, -32988, 212773, 213403, 217253, 144627, 164901, 168050, 57095, 80246, 85703, -131898, -130490, -123839, 75625, 71552, 74437, 225350, 227274, 232336,
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
        },
    },
}

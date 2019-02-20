package data

import (
    "github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1072 = rtcm3.Message1072 {
    MessageMsm2: rtcm3.MessageMsm2 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x430,
            ReferenceStationId: 0x0,
            Epoch: 0x1a6055a8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x1,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0x481a188a00000000,
            SignalMask: 0x40400000,
            CellMask: 0xfffff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm123 {
            Ranges: [] uint16 {
                0x2d8, 0x16d, 0x277, 0x33c, 0xe8, 0x27b, 0x338, 0x2c3, 0x38, 0x3a2,
            },
        },
        SignalData: rtcm3.SignalDataMsm2 {
            PhaseRanges: [] int32 {
                188717, 182008, 139312, 145832, 17706, 17357, -50425, -45119, -237961, -229946, -194001, -189501, -44516, -39821, -142566, -133983, -177392, -173383, 86776, 84624,
            },
            PhaseRangeLocks: [] uint8 {
                0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xe, 0xe, 0xf, 0xf, 0xb, 0xa, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
        },
    },
}

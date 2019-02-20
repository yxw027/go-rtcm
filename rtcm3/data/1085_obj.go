package data

import (
    "github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1085 = rtcm3.Message1085 {
    MessageMsm5: rtcm3.MessageMsm5 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x43d,
            ReferenceStationId: 0x0,
            Epoch: 0x195821a8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0x208070000000000,
            SignalMask: 0x41000000,
            CellMask: 0x3ff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x50, 0x4d, 0x50, 0x44, 0x45,
            },
            Extended: [] uint8 {
                0xc, 0x5, 0x4, 0xa, 0x9,
            },
            Ranges: [] uint16 {
                0x114, 0x30c, 0x244, 0x24e, 0x28b,
            },
            PhaseRangeRates: [] int16 {
                699, 727, 465, 137, -278,
            },
        },
        SignalData: rtcm3.SignalDataMsm5 {
            Pseudoranges: [] int16 {
                -1777, -1407, -137, 303, -912, -524, 1043, 1377, -1169, -934,
            },
            PhaseRanges: [] int32 {
                -54388, -46718, -8291, -6316, -35182, -33694, 45835, 68495, -34815, -30957,
            },
            PhaseRangeLocks: [] uint8 {
                0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint8 {
                0x28, 0x26, 0x25, 0x23, 0x29, 0x26, 0x34, 0x2b, 0x2e, 0x25,
            },
            PhaseRangeRates: [] int16 {
                606, 531, -3917, -4224, -154, -122, -131, -67, 2871, 3119,
            },
        },
    },
}

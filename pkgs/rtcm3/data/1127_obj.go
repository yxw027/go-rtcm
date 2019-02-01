package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1127 = rtcm3.Message1127 {
    MessageMsm7: rtcm3.MessageMsm7 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x467,
            ReferenceStationId: 0x0,
            Epoch: 0x1a601ef8,
            MultipleMessageBit: false,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0xf340000820000000,
            SignalMask: 0x40040000,
            CellMask: 0x37ffa,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x7b, 0x87, 0x80, 0x7c, 0x7b, 0x7c, 0x7c, 0x52, 0x48,
            },
            Extended: [] uint8 {
                0xd, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0,
            },
            Ranges: [] uint16 {
                0x3e7, 0x1b5, 0x18a, 0x208, 0x230, 0x287, 0x277, 0xb3, 0x3f7,
            },
            PhaseRangeRates: [] int16 {
                3, -10, -4, 2, 228, -7, 123, 369, 20,
            },
        },
        SignalData: rtcm3.SignalDataMsm7 {
            Pseudoranges: [] int32 {
                232851, 220709, 221039, 114010, 108333, -114920, -121259, -221714, -229111, 137864, 131778, -228282, -229632, 254220, 70239,
            },
            PhaseRanges: [] int32 {
                923853, 868442, 873326, 437537, 413045, -464954, -485553, -910633, -947564, 563004, 549527, -927962, -930483, 993276, 292096,
            },
            PhaseRangeLocks: [] uint16 {
                0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x2c0, 0x282, 0x282, 0x272, 0x272, 0x282, 0x282, 0x27f, 0x271,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x2df, 0x312, 0x26a, 0x29b, 0x2b2, 0x2d3, 0x2f9, 0x2f4, 0x314, 0x2be, 0x2ef, 0x2db, 0x2f7, 0x2be, 0x326,
            },
            PhaseRangeRates: [] int16 {
                2202, 2213, 3700, 4271, 3958, 1477, 1594, 3887, 4071, -3780, -3560, 3626, 3736, 4499, -3353,
            },
        },
    },
}

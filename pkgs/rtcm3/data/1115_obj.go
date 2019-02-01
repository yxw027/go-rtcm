package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1115 = rtcm3.Message1115 {
    MessageMsm5: rtcm3.MessageMsm5 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x45b,
            ReferenceStationId: 0x0,
            Epoch: 0x1a6055a8,
            MultipleMessageBit: true,
            Iods: 0x0,
            Reserved: 0x0,
            ClockSteeringIndicator: 0x0,
            ExternalClockIndicator: 0x0,
            SmoothingIndicator: false,
            SmoothingInterval: 0x0,
            SatelliteMask: 0xe200000000000000,
            SignalMask: 0x40010000,
            CellMask: 0xff,
        },
        SatelliteData: rtcm3.SatelliteDataMsm57 {
            RangeMilliseconds: [] uint8 {
                0x88, 0x88, 0x6d, 0x7c,
            },
            Extended: [] uint8 {
                0xd, 0x0, 0x0, 0x0,
            },
            Ranges: [] uint16 {
                0x38d, 0x1ad, 0x2c1, 0x3e4,
            },
            PhaseRangeRates: [] int16 {
                -394, 444, 35, 0,
            },
        },
        SignalData: rtcm3.SignalDataMsm5 {
            Pseudoranges: [] int16 {
                1950, 1956, -7766, -7621, 7591, 7597, 7784, 7816,
            },
            PhaseRanges: [] int32 {
                66010, 76263, -237248, -249890, 248197, 246924, 249537, 245832,
            },
            PhaseRangeLocks: [] uint8 {
                0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint8 {
                0x26, 0x24, 0x27, 0x27, 0x32, 0x33, 0x26, 0x2c,
            },
            PhaseRangeRates: [] int16 {
                -2440, -2224, 3395, 3611, 1404, 1331, -3516, -3617,
            },
        },
    },
}

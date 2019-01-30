package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1076 = rtcm3.Message1076 {
    MessageMsm6: rtcm3.MessageMsm6 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x434,
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
        SatelliteData: rtcm3.SatelliteDataMsm46 {
            RangeMilliseconds: [] uint8 {
                0x4d, 0x47, 0x49, 0x54, 0x54, 0x54, 0x4c, 0x45, 0x45, 0x51,
            },
            Ranges: [] uint16 {
                0x2d8, 0x16d, 0x277, 0x33c, 0xe8, 0x27b, 0x338, 0x2c3, 0x38, 0x3a2,
            },
        },
        SignalData: rtcm3.SignalDataMsm6 {
            Pseudoranges: [] int32 {
                187088, 180501, 124513, 122578, 19691, 18010, -53465, -49001, -243643, -241857, -195567, -190183, -53145, -55841, -141781, -139172, -186748, -189327, 89665, 88750,
            },
            PhaseRanges: [] int32 {
                754866, 728031, 557250, 583330, 70824, 69427, -201699, -180477, -951846, -919784, -776005, -758003, -178063, -159284, -570265, -535932, -709569, -693531, 347106, 338498,
            },
            PhaseRangeLocks: [] uint16 {
                0x286, 0x286, 0x26f, 0x26f, 0x286, 0x286, 0x1ca, 0x1c8, 0x209, 0x208, 0x16a, 0x14c, 0x242, 0x242, 0x276, 0x276, 0x26f, 0x26f, 0x24b, 0x24b,
            },
            HalfCycles: [] bool {
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
            },
            Cnrs: [] uint16 {
                0x270, 0x199, 0x31c, 0x29e, 0x2e0, 0x22c, 0x1ed, 0xa8, 0x252, 0x137, 0x1ca, 0xa2, 0x2b3, 0x1c2, 0x30c, 0x2b7, 0x32d, 0x2c9, 0x279, 0x146,
            },
        },
    },
}

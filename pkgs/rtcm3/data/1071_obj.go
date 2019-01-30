package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1071 = rtcm3.Message1071 {
    MessageMsm1: rtcm3.MessageMsm1 {
        Header: rtcm3.MsmHeader {
            MessageNumber: 0x42f,
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
        SignalData: rtcm3.SignalDataMsm1 {
            Pseudoranges: [] int16 {
                5847, 5641, 3891, 3831, 615, 563, -1671, -1531, -7614, -7558, -6111, -5943, -1661, -1745, -4431, -4349, -5836, -5916, 2802, 2773,
            },
        },
    },
}

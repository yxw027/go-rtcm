package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1006 = rtcm3.Message1006 {
    AntennaReferencePoint: rtcm3.AntennaReferencePoint {
        MessageNumber: 0x3ee,
        ReferenceStationId: 0x0,
        ItrfRealizationYear: 0x0,
        GpsIndicator: true,
        GlonassIndicator: true,
        GalileoIndicator: true,
        ReferenceStationIndicator: false,
        ReferencePointX: -44723575368,
        SingleReceiverOscilator: true,
        Reserved: false,
        ReferencePointY: 26704851794,
        QuarterCycleIndicator: 0x2,
        ReferencePointZ: -36693744263,
    },
    AntennaHeight: 0x0,
}

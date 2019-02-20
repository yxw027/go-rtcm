package data

import (
    "github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1008 = rtcm3.Message1008 {
    MessageAntennaDescriptor: rtcm3.MessageAntennaDescriptor {
        MessageNumber: 0x3f0,
        ReferenceStationId: 0x0,
        AntennaDescriptor: "",
        AntennaSetupId: 0x0,
    },
    SerialNumber: "Unknown",
}

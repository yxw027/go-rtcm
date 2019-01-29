package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1007 = rtcm3.Message1007 {
    MessageAntennaDescriptor: rtcm3.MessageAntennaDescriptor {
        MessageNumber: 0x3ef,
        ReferenceStationId: 0x0,
        AntennaDescriptor: "",
        AntennaSetupId: 0x0,
    },
}

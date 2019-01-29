package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1007 = rtcm3.Message1007 {
    AntennaDescriptor: rtcm3.AntennaDescriptor {
        MessageNumber: 0x3ef,
        ReferenceStationId: 0x0,
        DescriptorLength: 0x0,
        Descriptor: "",
        SetupId: 0x0,
    },
}

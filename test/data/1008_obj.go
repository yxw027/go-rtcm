package data

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
)

var Message1008 = rtcm3.Message1008 {
    AntennaDescriptor: rtcm3.AntennaDescriptor {
        MessageNumber: 0x3f0,
        ReferenceStationId: 0x0,
        DescriptorLength: 0x0,
        Descriptor: "",
        SetupId: 0x0,
    },
    SerialNumberLength: 0x7,
    SerialNumber: "Unknown",
}

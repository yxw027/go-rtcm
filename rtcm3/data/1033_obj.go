package data

import (
    "github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1033 = rtcm3.Message1033 {
    MessageAntennaDescriptor: rtcm3.MessageAntennaDescriptor {
        MessageNumber: 0x409,
        ReferenceStationId: 0x0,
        AntennaDescriptor: "",
        AntennaSetupId: 0x0,
    },
    AntennaSerialNumber: "Unknown",
    ReceiverTypeDescriptor: "SEPT POLARX5",
    ReceiverFirmwareVersion: "5.2.0",
    ReceiverSerialNumber: "3025123",
}

package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1013 = rtcm3.Message1013{
	MessageNumber:      0x3f5,
	ReferenceStationId: 0x0,
	Mjd:                0xe48c,
	SecondsOfDay:       0x2907,
	MessageCount:       0x0,
	LeapSeconds:        0x12,
	Messages:           []rtcm3.MessageAnnouncement(nil),
}

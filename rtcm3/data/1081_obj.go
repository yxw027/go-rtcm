package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1081 = rtcm3.Message1081{
	MessageMsm1: rtcm3.MessageMsm1{
		MsmHeader: rtcm3.MsmHeader{
			MessageNumber:          0x439,
			ReferenceStationId:     0x0,
			Epoch:                  0x29450ed8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x1,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0x7082e00000000000,
			SignalMask:             0x41000000,
			CellMask:               0xffff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm123{
			Ranges: []uint16{
				0x3fe, 0x3ac, 0x19b, 0x29, 0x4b, 0x36c, 0x16b, 0x28,
			},
		},
		SignalData: rtcm3.SignalDataMsm1{
			Pseudoranges: []int16{
				7585, 7899, 5647, 5858, -5605, -5320, -1743, -1556, 4007, 4492, -2394, -2086, 2482, 2757, -6632, -6430,
			},
		},
	},
}

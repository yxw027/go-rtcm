package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1093 = rtcm3.Message1093{
	MessageMsm3: rtcm3.MessageMsm3{
		MsmHeader: rtcm3.MsmHeader{
			MessageNumber:          0x445,
			ReferenceStationId:     0x0,
			Epoch:                  0x1a6055a8,
			MultipleMessageBit:     true,
			Iods:                   0x0,
			Reserved:               0x0,
			ClockSteeringIndicator: 0x1,
			ExternalClockIndicator: 0x0,
			SmoothingIndicator:     false,
			SmoothingInterval:      0x0,
			SatelliteMask:          0x823021c200000000,
			SignalMask:             0x40020200,
			CellMask:               0x7ffffff,
		},
		SatelliteData: rtcm3.SatelliteDataMsm123{
			Ranges: []uint16{
				0x398, 0x326, 0x1ca, 0x195, 0x17f, 0x37a, 0x36f, 0x2a8, 0x1ea,
			},
		},
		SignalData: rtcm3.SignalDataMsm3{
			Pseudoranges: []int16{
				1199, 1380, 1570, 7995, 8066, 8247, -1461, -1432, -1249, 6322, 6174, 6307, 4401, 4404, 4551, 1626, 2256, 2418, -4221, -4115, -3962, 2666, 2923, 3037, 7227, 7319, 7475,
			},
			PhaseRanges: []int32{
				18266, 17956, 22058, 256605, 263221, 269488, -39528, -39760, -32988, 212773, 213403, 217253, 144627, 164901, 168050, 57095, 80246, 85703, -131898, -130490, -123839, 75625, 71552, 74437, 225350, 227274, 232336,
			},
			PhaseRangeLocks: []uint8{
				0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf, 0xf,
			},
			HalfCycles: []bool{
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
		},
	},
}

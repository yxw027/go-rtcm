package data

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
)

var Message1001 = rtcm3.Message1001{
	GpsObservationHeader: rtcm3.GpsObservationHeader{
		MessageNumber:      0x3e9,
		ReferenceStationId: 0x0,
		Epoch:              0xb74ffa8,
		SynchronousGnss:    true,
		SignalsProcessed:   0x9,
		SmoothingIndicator: false,
		SmoothingInterval:  0x0,
	},
	SatelliteData: []rtcm3.SatelliteData1001{
		rtcm3.SatelliteData1001{
			SatelliteId:         0xf,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xbc0c3d,
			L1PhaseRange:        20356,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0x1b,
			L1CodeIndicator:     false,
			L1Pseudorange:       0x348296,
			L1PhaseRange:        3165,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0xd,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xa23993,
			L1PhaseRange:        11237,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0x14,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xc9230,
			L1PhaseRange:        22699,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0x1d,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xd65b21,
			L1PhaseRange:        2203,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0x1a,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xce95c6,
			L1PhaseRange:        5352,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0x10,
			L1CodeIndicator:     false,
			L1Pseudorange:       0x66ed0c,
			L1PhaseRange:        5909,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0xa,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xc110e1,
			L1PhaseRange:        26342,
			L1LockTimeIndicator: 0x7f,
		}, rtcm3.SatelliteData1001{
			SatelliteId:         0x15,
			L1CodeIndicator:     false,
			L1Pseudorange:       0xb22010,
			L1PhaseRange:        15104,
			L1LockTimeIndicator: 0x7f,
		},
	},
}

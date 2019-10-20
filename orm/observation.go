package orm

import (
	"github.com/jinzhu/gorm"
)

type Observation struct {
	gorm.Model
	// MessageNumber encodes constellation atm, could put this into SatelliteData
	// and have each constellation nested under the same "Observation" which is
	// be unique for (Epoch and ReferenceStationId) - that could be the PK
	// This would be more RINEX-like
	MessageNumber          uint16
	ReferenceStationId     uint16
	// TODO: normalize constellation epochs with timestamp
	Epoch                  uint32
	IODS                   uint8
	Reserved               uint8
	ClockSteeringIndicator uint8
	ExternalClockIndicator uint8
	SmoothingIndicator     bool
	SmoothingInterval      uint8
	SatelliteData []SatelliteData `gorm:"foreignkey:ObservationID"`
}

type SatelliteData struct {
	gorm.Model
	ObservationID     uint
	SatelliteID       int
	RangeMilliseconds uint8
	Extended          uint8
	Ranges            uint16
	PhaseRangeRates   int16
	SignalData []SignalData `gorm:"foreignkey:SatelliteDataID"`
}

type SignalData struct {
	gorm.Model
	SatelliteDataID uint
	SignalID        int
	Pseudoranges    int32
	PhaseRanges     int32
	PhaseRangeLocks uint16
	HalfCycles      bool
	CNRs            uint16
	PhaseRangeRates int16
}

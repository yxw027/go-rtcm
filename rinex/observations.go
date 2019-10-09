package rinex

import (
	"time"
	"fmt"
)

type EpochRecord struct {
	Epoch time.Time
	Flag int
	Satellites int
	ClockOffset float64
	ObservationRecords []ObservationRecord
}

func (er EpochRecord) String() string {
	return fmt.Sprint("> ", er.Epoch, er.Flag, er.Satellites, er.ClockOffset, "\n", er.ObservationRecords)
}

type ObservationRecord struct {
	Constellation rune
	SatelliteID int
	ObservationData []ObservationData
}

func (or ObservationRecord) String() string {
	return fmt.Sprint(or.Constellation, or.SatelliteID, or.ObservationData, "\n")
}

type ObservationData struct {
	ObservationType ObservationType
	Observation float64
	LLI int
	SSI int
}

func (od ObservationData) String() string {
	return fmt.Sprint(" ", od.Observation, od.LLI, od.SSI)
}

type ObservationType struct {
	Type rune
	Band int
	Attribute rune
}

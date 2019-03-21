package rtcm3_test

import (
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"testing"
	"time"
)

func TestGlonassTimeShort(t *testing.T) {
	beforeUtcZero := time.Date(2019, 2, 7, 23, 41, 40, 0, time.UTC)

	if beforeUtcZero != rtcm3.GlonassTimeShort(9700000, beforeUtcZero) {
		t.Errorf("GlonassTimeShort time incorrect")
	}

	afterUtcZero := time.Date(2019, 2, 8, 00, 44, 44, 0, time.UTC)

	if afterUtcZero != rtcm3.GlonassTimeShort(13484000, afterUtcZero) {
		t.Errorf("GlonassTimeShort time incorrect")
	}
}

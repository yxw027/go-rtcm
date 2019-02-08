package rtcm3_test

import (
    "testing"
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
    "time"
)

func TestGlonassTime(t *testing.T) {
    week := time.Date(2019, 2, 3, 0, 0, 0, 0, time.UTC)
    deltaTime, _ := time.ParseDuration("5h42m18s238ms")
    if rtcm3.GlonassTime(654609150, week) != week.AddDate(0, 0, 5).Add(deltaTime) {
        t.Errorf("GlonassTime time incorrect")
    }
}

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

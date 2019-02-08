package rtcm3_test

import (
    "testing"
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm3"
    "time"
)

func TestGpsTime(t *testing.T) {
    week := time.Date(2019, 2, 3, 0, 0, 0, 0, time.UTC)
    deltaTime, _ := time.ParseDuration("5h42m18s238ms")
    if rtcm3.GlonassTime(654609150, week) != week.AddDate(0, 0, 5).Add(deltaTime) {
        t.Errorf("GlonassTime time incorrect")
    }
}

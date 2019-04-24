package rtcm3

import (
	"time"
)

func GpsTime(e uint32) time.Time {
	now := time.Now().UTC()
	sow := now.Truncate(time.Hour*24).AddDate(0, 0, -int(now.Weekday()))
	tow := time.Duration(e) * time.Millisecond
	return sow.Add(-(18 * time.Second)).Add(tow)
}

func GlonassTime(e uint32) time.Time {
	now := time.Now().UTC()
	sow := now.Truncate(time.Hour*24).AddDate(0, 0, -int(now.Weekday()))
	dow := int((e >> 27) & 0x7)
	tod := time.Duration(e&0x7FFFFFF) * time.Millisecond
	return sow.AddDate(0, 0, dow).Add(tod).Add(-(3 * time.Hour))
}

func GlonassTimeShort(e uint32, now time.Time) time.Time {
	hours := e / 3600000
	moduloGlonassHours := ((int(hours) - 3%24) + 24) % 24
	rest := int(e) - (int(hours) * 3600000)
	tod := time.Duration(rest+(moduloGlonassHours*3600000)) * time.Millisecond
	dow := now.Truncate(time.Hour * 24)
	return dow.Add(tod)
}

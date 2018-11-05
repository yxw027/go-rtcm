package main

import (
    "github.com/geoscienceaustralia/go-rtcm"
    "github.com/umeat/go-ntrip/ntrip"
    "net/url"
    "fmt"
    "os"
    "time"
)

func main() {
    c, _ := url.Parse(os.Args[1])
    r, _ := ntrip.Client(c, os.Args[2], os.Args[3])

    e := rtcm.Scan(r, func(msg rtcm.Rtcm3Message) {
        now := time.Now().UTC()
        sow := now.Truncate(time.Hour * 24).AddDate(0, 0, -int(now.Weekday()))
        if msg.Number() == uint16(1077) || msg.Number() == uint16(1097) || msg.Number() == uint16(1117) || msg.Number() == uint16(1127) {
            tow := time.Duration(msg.(*rtcm.Rtcm3Msm7Message).Header.Epoch) * time.Millisecond
            latency := now.Sub(sow.Add(-(18 * time.Second)).Add(tow))
            if msg.Number() == uint16(1127) { latency = latency - (14 * time.Second) }
            fmt.Println(msg.Number(), latency)
        } else if msg.Number() == uint16(1087) {
            e := msg.(*rtcm.Rtcm3Msm7Message).Header.Epoch
            dow := int((e >> 27) & 0x7)
            tod := time.Duration(e & 0x7FFFFFF) * time.Millisecond
            fmt.Println(msg.Number(), now.Sub(sow.AddDate(0, 0, dow).Add(tod).Add(-(3 * time.Hour))))
        }
    })
    fmt.Println(e)
}

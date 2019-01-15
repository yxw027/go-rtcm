package main

import (
    "github.com/geoscienceaustralia/go-rtcm/pkgs/rtcm"
    "github.com/umeat/go-ntrip/pkgs/ntrip"
    "fmt"
    "os"
    "time"
)

func main() {
    r, _ := ntrip.Client(os.Args[1], os.Args[2], os.Args[3])
    scanner := rtcm.NewScanner(r)
    msg, err := scanner.Next()
    for ; err == nil; msg, err = scanner.Next() {
        switch int(msg.Number()) {
            case 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1097, 1117, 1087, 1127, 1001, 1002, 1003, 1004, 1009, 1010, 1011, 1012:
                fmt.Println(msg.Number(), time.Now().UTC().Sub(msg.(rtcm.Rtcm3Observable).Time()))
        }
    }

    fmt.Println(err)
}

package rtcm

import (
    "github.com/bamiaux/iobit"
)

// Sign-Magnitude Ints

func Sint8(r iobit.Reader, length int) int8 {
    n, v := r.Bit(), int8(r.Uint8(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func Sint16(r iobit.Reader, length int) int16 {
    n, v := r.Bit(), int16(r.Uint16(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func Sint32(r iobit.Reader, length int) int32 {
    n, v := r.Bit(), int32(r.Uint32(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

package rtcm3

import (
    "github.com/bamiaux/iobit"
    "math"
)

// Sign-Magnitude Ints

func Sint8(r *iobit.Reader, length int) int8 {
    n, v := r.Bit(), int8(r.Uint8(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func Sint16(r *iobit.Reader, length int) int16 {
    n, v := r.Bit(), int16(r.Uint16(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func Sint32(r *iobit.Reader, length int) int32 {
    n, v := r.Bit(), int32(r.Uint32(uint(length - 1)))
    if n == true {
        v = -v
    }
    return v
}

func PutSint8(w *iobit.Writer, length int, sint int8) {
    w.PutBit(math.Signbit(float64(sint)))
    if sint < 0 {
        sint = -sint
    }
    w.PutUint8(uint(length - 1), uint8(sint))
}

func PutSint16(w *iobit.Writer, length int, sint int16) {
    w.PutBit(math.Signbit(float64(sint)))
    if sint < 0 {
        sint = -sint
    }
    w.PutUint16(uint(length - 1), uint16(sint))
}

func PutSint32(w *iobit.Writer, length int, sint int32) {
    w.PutBit(math.Signbit(float64(sint)))
    if sint < 0 {
        sint = -sint
    }
    w.PutUint32(uint(length - 1), uint32(sint))
}

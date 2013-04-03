package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strings"
	"time"
)

func GetDig(password string, interval_no int64) []byte {
	var key, _ = base32.StdEncoding.DecodeString(strings.ToUpper(password))
	var h = hmac.New(sha1.New, []byte(key))
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(interval_no))
	h.Write(buf)
	dig := h.Sum(nil)
	return (dig)
}

func Hotp(password string, interval_no int64) uint32 {
	return (HotpDig(GetDig(password, interval_no)))
}

func Totp(password string) uint32 {
	return (Hotp(password, time.Now().Unix()/30))
}

func HotpDig(dig []byte) uint32 {
	offset := dig[19] & 15
	dt := make([]byte, 4)
	copy(dt, dig[offset:offset+4])
	var ui = binary.BigEndian.Uint32(dt)
	otp := (ui & 0x7fffffff) % 1000000
	return (otp)
}

package utils

import (
	"time"

	"github.com/muyo/sno"
)

var meta byte

func init() {
	meta = 88
}

func SetIDMeta(b byte) {
	AppLogger.Infof("Setting random id generator meta value to: %v", b)
	meta = b
}

func GenerateRequestIDByte() []byte {
	return sno.NewWithTime(meta, time.Now()).Bytes()
}

func GenerateRequestID() sno.ID {
	return sno.NewWithTime(meta, time.Now())
}

func GenerateRequestIDString() string {
	return sno.NewWithTime(meta, time.Now()).String()
}

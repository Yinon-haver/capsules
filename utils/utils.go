package utils

import (
	"fmt"
	"time"
)

func GetTimestampString() string {
	t := time.Now()
	res := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return res
}
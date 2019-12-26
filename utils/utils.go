package utils

import (
	"bytes"
	"fmt"
	"time"
)

type Void struct{}

func GetTimestampString() string {
	t := time.Now()
	res := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return res
}

func RunProcess(process func(), workersNumber int) {
	for i := 0; i < workersNumber; i++ {
		go process()
	}
}

func MultiplyQuote(s string) string {
	var buf bytes.Buffer

	for _, i := range s {
		if i == '\'' {
			buf.WriteRune('\'')
		}
		buf.WriteRune(i)
	}

	return buf.String()
}

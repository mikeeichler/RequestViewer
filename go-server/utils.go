package main

import (
	"fmt"
	"time"
)

func timestamp() (timestamp string) {
	currentTime := time.Now()
	yr := fmt.Sprintf("%04d", currentTime.Year())
	mo := fmt.Sprintf("%02d", int(currentTime.Month()))
	dy := fmt.Sprintf("%02d", currentTime.Day())
	hr := fmt.Sprintf("%02d", currentTime.Hour())
	mi := fmt.Sprintf("%02d", currentTime.Minute())
	sc := fmt.Sprintf("%02d", currentTime.Second())
	timestamp = fmt.Sprintf("%s-%s-%s_T_%s:%s:%s", yr, mo, dy, hr, mi, sc)
	return
}

package main

import (
	"fmt"
	"strings"
	"time"
)

func inSlice(v string, s []string) bool {
	for _, vv := range s {
		if strings.EqualFold(v, vv) {
			return true
		}
	}
	return false
}

func clientHints() []string {
	chSlice := []string{"Sec-CH-UA", "Sec-CH-UA-Arch", "Sec-CH-UA-Bitness", "Sec-CH-UA-Full-Version-List", "Sec-CH-UA-Full-Version", "Sec-CH-UA-Mobile", "Sec-CH-UA-Model", "Sec-CH-UA-Platform", "Sec-CH-UA-Platform-Version"}
	return chSlice
}

func dapropsFromCookie(cookie string) string {
	cookies := strings.Split(cookie, ";")
	for _, c := range cookies {
		c = strings.TrimSpace(c)
		if strings.HasPrefix(c, "DAPROPS") {
			return c
		}
	}
	return ""
}

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

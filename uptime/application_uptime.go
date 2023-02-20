//With inspiration from https://go.dev/play/p/by_nkvhzqD

package uptime

import "time"

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func GetUptime() int {
	return int(time.Since(startTime).Seconds())
}

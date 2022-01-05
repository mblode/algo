package common

import (
	"log"
	"runtime"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	log.Printf("%s took %dms (%dns)", name, elapsed.Microseconds(), elapsed.Nanoseconds())
}

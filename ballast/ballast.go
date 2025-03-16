package ballast

import (
	"fmt"
	"github.com/sourceinsight/gox/env"
	"runtime"
)

var ballastMemory []byte

// Ballast is a function to allocate memory as ballast. unit is MB.
func Ballast(sizeInMB int) {
	ballastMemory = make([]byte, sizeInMB*1024*1024)
	ballastMemory[0] = 1
	runtime.KeepAlive(ballastMemory)
	fmt.Printf("Ballast of %d MB allocated.\n", sizeInMB)
}

func BallastFromEnv() {
	ballastSize := env.GetEnv("BALLAST_SIZE", 0)
	if ballastSize > 0 {
		Ballast(ballastSize)
	}
}

package ballast

import (
	"fmt"
	"github.com/sourceinsight/gox/env"
	"runtime"
)

// Ballast is a function to allocate memory as ballast. unit is MB.
func Ballast(sizeInMB int) {
	ballast := make([]byte, sizeInMB*1024*1024)
	ballast[0] = 0
	runtime.KeepAlive(ballast)
	fmt.Printf("Ballast of %d MB allocated.\n", sizeInMB)
}

func BallastFromEnv() {
	ballastSize := env.GetEnv("BALLAST_SIZE", 0)
	if ballastSize > 0 {
		Ballast(ballastSize)
	}
}

package globaltick

import (
	"sync/atomic"
	"time"
)

var (
	nowTimestamp int64 // now timestamp
)

func init() {
	tick := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-tick.C
			atomic.StoreInt64(&nowTimestamp, time.Now().Unix())
		}
	}()
}

func GetTimestamp() int64 {
	return atomic.LoadInt64(&nowTimestamp)
}

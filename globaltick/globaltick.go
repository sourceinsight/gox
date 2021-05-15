package globaltick

import "time"

var (
	Now          time.Time // now
	NowTimestamp int64     // now timestamp
)

func init() {
	tick := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-tick.C
			Now = time.Now()
			NowTimestamp = Now.Unix()
		}
	}()
}

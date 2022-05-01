package channels

import "sync/atomic"

// FanIn сливает несколько каналов в один
func FanIn[T any](chans ...chan T) chan T {
	out := make(chan T)
	workers := int64(len(chans))
	for _, ch := range chans {
		func(ch chan T) {
			go func() {
				for t := range ch {
					out <- t
				}
				atomic.AddInt64(&workers, -1)
				if workers == 0 {
					close(out)
				}
			}()
		}(ch)
	}
	return out
}

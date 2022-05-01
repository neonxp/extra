package channels

// FanOut раскидывает очередь канала на несколько каналов равномерно
func FanOut[T any](in chan T, workers int) []chan T {
	out := make([]chan T, workers)
	for i := 0; i < workers; i++ {
		out[i] = make(chan T)
	}
	go func() {
		i := 0
		for t := range in {
			out[i] <- t
			i++
			if i == workers {
				i = 0
			}
		}
		for i := 0; i < workers; i++ {
			close(out[i])
		}
	}()
	return out
}

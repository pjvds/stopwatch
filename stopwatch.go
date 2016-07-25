package stopwatch

import "time"

type Stopwatch time.Time

func Start() Stopwatch {
	return Stopwatch(time.Now().UTC())
}

func (this Stopwatch) StartedAt() time.Time {
	return time.Time(this)
}

func (this Stopwatch) Elapsed() time.Duration {
	return time.Since(this.StartedAt())
}

func Time(callback func()) time.Duration {
	watch := Start()
	callback()

	return watch.Elapsed()
}

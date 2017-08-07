package hebiten

import "time"

type MaxUpdater struct {
	previousSecond int
	bestValue      float64
	Value          float64
}

func (this *MaxUpdater) Update(value float64) {
	if this.bestValue < value {
		this.bestValue = value
	}
	var currentSecond = time.Now().Second()
	if currentSecond != this.previousSecond {
		this.Value = this.bestValue
		this.bestValue = 0
		this.previousSecond = currentSecond
	}
}

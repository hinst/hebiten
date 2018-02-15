package hebiten

type Pendulum struct {
	Value     float64
	Speed     float64
	Direction bool
	Top       float64
	Bottom    float64
}

func (this *Pendulum) Update(deltaTime float64) {
	var delta = deltaTime * this.Speed
	if this.Direction {
		this.Value += delta
		if this.Top < this.Value {
			this.Value = this.Top
			this.Direction = false
		}
	} else {
		this.Value -= delta
		if this.Value < this.Bottom {
			this.Value = this.Bottom
			this.Direction = true
		}
	}
}

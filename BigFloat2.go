package hebiten

import (
	"hgo"
)

type BigFloat2 struct {
	X, Y float64
}

func (this BigFloat2) Check0() bool {
	return this.X == 0 && this.Y == 0
}

func (this BigFloat2) Add(a BigFloat2) BigFloat2 {
	return BigFloat2{this.X + a.X, this.Y + a.Y}
}

func (this BigFloat2) Substract(a BigFloat2) BigFloat2 {
	return BigFloat2{this.X - a.X, this.Y - a.Y}
}

func (this BigFloat2) Approach(delta, target BigFloat2) (result BigFloat2) {
	result.X = ApproachFloat64(this.X, delta.X, target.X)
	result.Y = ApproachFloat64(this.Y, delta.Y, target.Y)
	return
}

func (this BigFloat2) FromFloat64(a float64) BigFloat2 {
	return BigFloat2{X: a, Y: a}
}

func (this BigFloat2) ToString() string {
	return hgo.Float64ToStrPres(this.X, 1) + " " + hgo.Float64ToStrPres(this.Y, 1)
}

package hebiten

import "hgo"

type TIntRect struct {
	X, Y, W, H int
}

func (this *TIntRect) RunForEach(f func(position Int2)) {
	var right = this.X + this.W
	var bottom = this.Y + this.H
	for x := this.X; x < right; x++ {
		for y := this.Y; y < bottom; y++ {
			f(Int2{x, y})
		}
	}
}

func (this *TIntRect) ToString() string {
	return hgo.IntToStr(this.X) + "," + hgo.IntToStr(this.Y) + ";" + hgo.IntToStr(this.W) + "," + hgo.IntToStr(this.H)
}

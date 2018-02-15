package hebiten

import (
	"hgo"
	"image"
)

type TIntRect struct {
	X, Y, W, H int
}

var GlobalIntRect TIntRect

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

func (this *TIntRect) Grow(delta int) {
	this.X -= delta
	this.Y -= delta
	this.W += delta * 2
	this.H += delta * 2
}

func (this *TIntRect) ToImageRect() (result image.Rectangle) {
	result.Min.X = this.X
	result.Min.Y = this.Y
	result.Max.X = this.X + this.W
	result.Max.Y = this.Y + this.H
	return
}

// Static
func (this *TIntRect) ImageRectToStr(rect image.Rectangle) string {
	return hgo.IntToStr(rect.Min.X) + "," + hgo.IntToStr(rect.Min.Y) + ";" + hgo.IntToStr(rect.Max.X) + "," + hgo.IntToStr(rect.Max.Y)
}

func (this *TIntRect) GetSize() Int2 {
	return Int2{X: this.W, Y: this.H}
}

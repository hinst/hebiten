package hebiten

import (
	"github.com/hajimehoshi/ebiten"
)

type TFloatRect struct {
	X, Y, W, H float64
}

func (this *TFloatRect) GetLT() BigFloat2 {
	return BigFloat2{this.X, this.Y}
}

func (this *TFloatRect) GetRT() BigFloat2 {
	return BigFloat2{this.GetRight(), this.Y}
}

func (this *TFloatRect) GetLB() BigFloat2 {
	return BigFloat2{this.X, this.GetBottom()}
}

func (this *TFloatRect) GetRB() BigFloat2 {
	return BigFloat2{this.GetRight(), this.GetBottom()}
}

func (this *TFloatRect) GetPoints() BigFloat2s {
	return []BigFloat2{this.GetLT(), this.GetRT(), this.GetLB(), this.GetRB()}
}

func (this *TFloatRect) GetRight() float64 {
	return this.X + this.W
}

func (this *TFloatRect) GetBottom() float64 {
	return this.Y + this.H
}

func (this *TFloatRect) CheckContainsPoint(point BigFloat2) bool {
	return this.X < point.X && point.X < this.GetRight() &&
		this.Y < point.Y && point.Y < this.GetBottom()
}

func (this *TFloatRect) LoadImage(image *ebiten.Image) *TFloatRect {
	this.X = 0
	this.Y = 0
	var w, h = image.Size()
	this.W = float64(w)
	this.H = float64(h)
	return this
}

func (this *TFloatRect) GetEnlarged(delta float64) TFloatRect {
	return TFloatRect{
		X: this.X - delta,
		Y: this.Y - delta,
		W: this.W + 2*delta,
		H: this.H + 2*delta,
	}
}

func (this *TFloatRect) GetSize() BigFloat2 {
	return BigFloat2{X: this.W, Y: this.H}
}

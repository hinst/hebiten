package hebiten

import (
	"hgo"

	"github.com/hajimehoshi/ebiten"
)

type TDraw struct {
	Target, Image *ebiten.Image
	ImageSize     BigFloat2
	Position      BigFloat2
	Angle         float64
	Origin        BigFloat2
	Size          BigFloat2
	Scale         BigFloat2
	CenterOrigin  bool
	Mask          TFloatColor
}

var GlobalDrawNumber int

func (this *TDraw) GetImageSize() BigFloat2 {
	var w, h = this.Image.Size()
	return BigFloat2{float64(w), float64(h)}
}

func (this *TDraw) InferScale() {
	this.Scale.X = this.Size.X / this.ImageSize.X
	this.Scale.Y = this.Size.Y / this.ImageSize.Y
}

func (this *TDraw) Draw() {
	hgo.Assert(this.Image != nil)
	this.ImageSize = this.GetImageSize()
	if this.Scale.Check0() {
		this.InferScale()
	}
	var halfW = this.ImageSize.X / 2
	var halfH = this.ImageSize.Y / 2
	if this.CenterOrigin {
		this.Origin.X = halfW
		this.Origin.Y = halfH
	}
	var o = &ebiten.DrawImageOptions{}
	o.GeoM.Translate(-halfW, -halfH)
	o.GeoM.Rotate(this.Angle)
	o.GeoM.Translate(halfW, halfH)
	o.GeoM.Translate(-this.Origin.X, -this.Origin.Y)
	if false == this.Size.Check0() {
		o.GeoM.Scale(this.Scale.X, this.Scale.Y)
	}
	o.GeoM.Translate(this.Position.X, this.Position.Y)
	if false == this.Mask.CheckIfZero() {
		o.ColorM.Scale(this.Mask.R, this.Mask.G, this.Mask.B, this.Mask.A)
	}
	if this.CheckVisibility() {
		this.Target.DrawImage(this.Image, o)
		GlobalDrawNumber++
	}
}

func (this *TDraw) GetRect() TFloatRect {
	return TFloatRect{
		this.Position.X - this.Origin.X*this.Scale.X,
		this.Position.Y - this.Origin.Y*this.Scale.X,
		this.Size.X,
		this.Size.Y}
}

func (this *TDraw) CheckVisibility() (result bool) {
	var imageRect = this.GetRect()
	var destinationRect TFloatRect
	destinationRect.LoadImage(this.Target)
	destinationRect = destinationRect.GetEnlarged(1)
	for _, point := range imageRect.GetPoints() {
		if destinationRect.CheckContainsPoint(point) {
			result = true
			break
		}
	}
	return
}

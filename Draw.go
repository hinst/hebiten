package hebiten

import (
	"hgo"
	"runtime/debug"

	"github.com/hajimehoshi/ebiten"
)

type TDraw struct {
	Target       *ebiten.Image
	Image        *ebiten.Image
	AtlasTexture AtlasTexture
	Position     BigFloat2
	Angle        float64
	Origin       BigFloat2
	Size         BigFloat2
	Scale        BigFloat2
	CenterOrigin bool
	Mask         TFloatColor
}

var GlobalDrawNumber int
var DebugImage *ebiten.Image
var DebugImageModeEnabled bool

func (this *TDraw) GetImageSize() BigFloat2 {
	var w, h = this.Image.Size()
	return BigFloat2{float64(w), float64(h)}
}

func (this *TDraw) InferScale(imageSize BigFloat2) {
	this.Scale.X = this.Size.X / imageSize.X
	this.Scale.Y = this.Size.Y / imageSize.Y
}

func (this *TDraw) Draw() {
	var o = &ebiten.DrawImageOptions{}
	var imageSize BigFloat2
	if this.AtlasTexture.Exists() {
		this.Image = this.AtlasTexture.Image
		imageSize = this.AtlasTexture.Rect.GetSize().ToBigFloat()
		var imageRect = this.AtlasTexture.Rect.ToImageRect()
		o.SourceRect = &imageRect
		if false {
			println(GlobalIntRect.ImageRectToStr(*o.SourceRect))
			var w, h = this.Image.Size()
			o.SourceRect.Min.X = 0
			o.SourceRect.Min.Y = 0
			o.SourceRect.Max.X = w
			o.SourceRect.Max.Y = h
		}
	} else {
		hgo.AssertWT(this.Image != nil, func() string { return "this.Image field is nil" })
		imageSize = this.GetImageSize()
		if false {
			println(string(debug.Stack()))
		}
	}
	if this.Scale.Check0() {
		this.InferScale(imageSize)
	}
	var halfW = imageSize.X / 2
	var halfH = imageSize.Y / 2
	if this.CenterOrigin {
		this.Origin.X = halfW
		this.Origin.Y = halfH
	}
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
		if DebugImageModeEnabled {
			this.Target.DrawImage(DebugImage, o)
		} else {
			this.Target.DrawImage(this.Image, o)
		}
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

func (this *TDraw) SetPosRect(rect TFloatRect) {
	this.Position.X = rect.X
	this.Position.Y = rect.Y
	this.Size.X = rect.W
	this.Size.Y = rect.H
}

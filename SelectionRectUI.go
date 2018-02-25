package hebiten

import (
	"github.com/hajimehoshi/ebiten"
)

type SelectionRectUI struct {
	BackgroundImage AtlasTexture
	Target          *ebiten.Image
	Pendulum        Pendulum
	Color           TFloatColor
}

func (this *SelectionRectUI) Create() *SelectionRectUI {
	this.Pendulum.Bottom = 0.5
	this.Pendulum.Top = 1
	this.Pendulum.Speed = 0.5
	if this.Color.CheckIfZero() {
		this.Color = this.GetDefaultColor()
	}
	return this
}

func (this *SelectionRectUI) GetDefaultColor() TFloatColor {
	return TFloatColor{R: 128, G: 191, B: 255, A: 255}.Divide(255)
}

func (this *SelectionRectUI) Draw(rect TFloatRect) {
	var selectionRect TDraw
	selectionRect.Target = this.Target
	selectionRect.AtlasTexture = this.BackgroundImage
	selectionRect.Position = rect.GetLT()
	selectionRect.Size = rect.GetSize()
	selectionRect.Mask = this.Color.SetAlpha(this.Pendulum.Value)
	selectionRect.Draw()
}

func (this *SelectionRectUI) Update(deltaTime float64) {
	this.Pendulum.Update(deltaTime)
}

func (this *SelectionRectUI) SetTarget(v *ebiten.Image) {
	if this.Target != v {
		this.Target = v
	}
}

func (this *SelectionRectUI) SetBackgroundImage(v AtlasTexture) {
	if false == this.BackgroundImage.CheckEquals(v) {
		this.BackgroundImage = v
	}
}

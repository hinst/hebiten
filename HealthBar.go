package hebiten

import (
	"github.com/hajimehoshi/ebiten"
)

type HealthBar struct {
	BaseHealth    float64
	CurrentHealth float64
	Position      TFloatRect
	Texture       AtlasTexture
	HealthColor   TFloatColor
	DamageColor   TFloatColor
}

func (this *HealthBar) Create() *HealthBar {
	if this.HealthColor.CheckIfZero() {
		this.HealthColor = this.GetDefaultHealthColor()
	}
	if this.DamageColor.CheckIfZero() {
		this.DamageColor = this.GetDefaultDamageColor()
	}
	return this
}

func (this *HealthBar) GetDefaultHealthColor() TFloatColor {
	return TFloatColor{A: 0.9, G: 0.6}
}

func (this *HealthBar) GetDefaultDamageColor() TFloatColor {
	return TFloatColor{A: 0.9, R: 0.6}
}

func (this *HealthBar) Draw(target *ebiten.Image) {
	var healthWidth = this.Position.W * this.CurrentHealth / this.BaseHealth
	var damageWidth = this.Position.W - healthWidth
	{
		var draw = TDraw{Target: target}
		draw.AtlasTexture = this.Texture
		draw.Mask = this.HealthColor
		draw.SetPosRect(this.Position)
		draw.Size.X = healthWidth
		draw.Draw()
	}
	{
		var draw = TDraw{Target: target}
		draw.AtlasTexture = this.Texture
		draw.Mask = this.DamageColor
		draw.SetPosRect(this.Position)
		draw.Position.X = this.Position.X + healthWidth
		draw.Size.X = damageWidth
		draw.Draw()
	}
}

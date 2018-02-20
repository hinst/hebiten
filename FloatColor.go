package hebiten

import (
	"image/color"
	"math"
)

type TFloatColor struct {
	// Range for these fields is 0...1
	R, G, B, A float64
}

var _ color.Color = TFloatColor{}
var MaxUint32_Float64 = float64(math.MaxUint32)
var MaxUint8_Float64 = float64(math.MaxUint8)

func (this TFloatColor) CheckIfZero() bool {
	return this.R == 0 && this.G == 0 && this.B == 0 && this.A == 0
}

func (this TFloatColor) Divide(x float64) TFloatColor {
	return TFloatColor{R: this.R / x, G: this.G / x, B: this.B / x, A: this.A / x}
}

func (this TFloatColor) SetAlpha(x float64) TFloatColor {
	this.A = x
	return this
}

func (this TFloatColor) ValueToUint32(v float64) uint32 {
	if v == 0 {
		return 0
	}
	if v == 1 {
		return math.MaxUint32
	}
	return uint32(v * MaxUint32_Float64)
}

func (this TFloatColor) ValueToUint8(v float64) uint8 {
	if v == 0 {
		return 0
	}
	if v == 1 {
		return math.MaxUint8
	}
	return uint8(v * MaxUint8_Float64)
}

func (this TFloatColor) RGBA() (r, g, b, a uint32) {
	var f = this.ValueToUint32
	r = f(this.R)
	g = f(this.G)
	b = f(this.B)
	a = f(this.A)
	return
}

func (this TFloatColor) To8bit() (result color.RGBA) {
	var f = this.ValueToUint8
	result.R = f(this.R)
	result.G = f(this.G)
	result.B = f(this.B)
	result.A = f(this.A)
	return
}

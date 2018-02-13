package hebiten

type TFloatColor struct {
	R, G, B, A float64
}

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

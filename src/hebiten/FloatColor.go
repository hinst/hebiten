package hebiten

type TFloatColor struct {
	R, G, B, A float64
}

func (this TFloatColor) CheckIfZero() bool {
	return this.R == 0 && this.G == 0 && this.B == 0 && this.A == 0
}

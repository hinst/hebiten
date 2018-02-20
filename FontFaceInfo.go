package hebiten

import "golang.org/x/image/font"

type FontFaceInfo struct {
	Face *font.Face
	Size BigFloat2
}

func (this *FontFaceInfo) GetTextWidth(s string) float64 {
	return float64(len(s)) * this.Size.X
}

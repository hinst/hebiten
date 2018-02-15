package hebiten

import "github.com/hajimehoshi/ebiten"

type AtlasTexture struct {
	// Source position
	Rect TIntRect
	// Source image
	Image *ebiten.Image
}

func (this AtlasTexture) Exists() bool {
	return this.Image != nil
}

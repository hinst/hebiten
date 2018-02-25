package hebiten

import (
	"github.com/hajimehoshi/ebiten"
)

func GetImageWidth(image *ebiten.Image) int {
	var w, _ = image.Size()
	return w
}

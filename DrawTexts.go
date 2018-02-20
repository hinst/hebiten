package hebiten

import (
	"hgo"

	"github.com/hajimehoshi/ebiten"
	ebiten_text "github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type DrawTexts struct {
	Target *ebiten.Image
	Font   font.Face
	List   []DrawText
	count  int
}

func (this *DrawTexts) EnsureLength(length int) {
	if len(this.List) != length {
		this.List = make([]DrawText, length)
	}
}

func (this *DrawTexts) Reset(length int) {
	this.EnsureLength(length)
	this.count = 0
}

func (this *DrawTexts) Add(item DrawText) {
	var length = len(this.List)
	var count = this.count
	hgo.AssertWT(count < length,
		func() string { return "Length exceeded; l=" + hgo.IntToStr(length) + " c=" + hgo.IntToStr(count) })
	this.List[count] = item
	this.count++
}

func (this *DrawTexts) Draw() {
	for i, drawText := range this.List {
		if this.count <= i {
			break
		}
		ebiten_text.Draw(this.Target, drawText.Text, this.Font, drawText.X, drawText.Y, drawText.Color)
		GlobalDrawNumber++
	}
}

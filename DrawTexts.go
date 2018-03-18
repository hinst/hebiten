package hebiten

import (
	"hgo"

	"github.com/hajimehoshi/ebiten"
	ebiten_text "github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type DrawTexts struct {
	Target     *ebiten.Image
	Font       font.Face
	List       []DrawText
	count      int
	CountLimit int
}

func (this *DrawTexts) Create() *DrawTexts {
	if this.CountLimit == 0 {
		this.CountLimit = 1000
	}
	return this
}

func (this *DrawTexts) Reset() {
	this.count = 0
}

func (this *DrawTexts) Add(item DrawText) {
	var count = this.count
	var countLimit = this.CountLimit
	hgo.AssertWT(count < this.CountLimit, func() string {
		return "Count exceeded; count=" + hgo.IntToStr(count) + " countLimit=" + hgo.IntToStr(countLimit)
	})
	var length = len(this.List)
	if count < length {
		this.List[count] = item
	} else {
		this.List = append(this.List, item)
	}
	this.count++
}

func (this *DrawTexts) Draw() {
	for i, drawText := range this.List {
		if this.count <= i {
			break
		}
		hgo.Assert(this.Target != nil)
		ebiten_text.Draw(this.Target, drawText.Text, this.Font, drawText.X, drawText.Y, drawText.Color)
		GlobalDrawNumber++
	}
}

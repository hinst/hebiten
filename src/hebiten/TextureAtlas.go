package hebiten

import (
	"hgo"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type TextureAtlas struct {
	texture          *ebiten.Image
	size             Int2
	currentPos       Int2
	currentRowHeight int
	images           map[string]TIntRect
}

func (this *TextureAtlas) SetSize(size Int2) {
	hgo.Assert(this.size.CheckIf0())
	this.size = size
}

func (this *TextureAtlas) Create() *TextureAtlas {
	if this.size.CheckIf0() {
		this.size = Int2{X: 512, Y: 512}
	}
	var texture, newImageResult = ebiten.NewImage(this.size.X, this.size.Y, ebiten.FilterNearest)
	hgo.AssertResult(newImageResult)
	this.texture = texture
	this.images = make(map[string]TIntRect)
	return this
}

func (this *TextureAtlas) Write(name string, image *ebiten.Image) {
	var imageWidth, imageHeight = image.Size()
	var movePos = func() {
		hgo.Assert(imageWidth <= this.size.X)
		hgo.Assert(imageHeight <= this.size.Y)
		if this.size.X < this.currentPos.X+imageWidth {
			this.currentPos.X = 0
			this.currentPos.Y += this.currentRowHeight
			this.currentRowHeight = 0
		}
		hgo.Assert(this.currentPos.Y+imageHeight <= this.size.Y)
	}
	movePos()
	var drawImage = func() {
		var draw TDraw
		draw.Position = this.currentPos.ToBigFloat()
		draw.Target = this.texture
		draw.Image = image
		draw.Draw()
	}
	drawImage()
	var storePos = func() {
		var rect TIntRect
		rect.X = this.currentPos.X
		rect.Y = this.currentPos.Y
		rect.W = imageWidth
		rect.H = imageHeight
		this.images[name] = rect
		if this.currentRowHeight < imageHeight {
			this.currentRowHeight = imageHeight
		}
	}
	storePos()
	this.currentPos.X += imageWidth
}

func (this *TextureAtlas) LoadFromDir(path string) int {
	var files, readDirResult = ioutil.ReadDir(path)
	hgo.AssertResult(readDirResult)
	for _, file := range files {
		var fileName = file.Name()
		var filePath = path + "/" + fileName
		var image, _, imageResult = ebitenutil.NewImageFromFile(filePath, ebiten.FilterNearest)
		hgo.AssertResult(imageResult)
		this.Write(fileName, image)
	}
	return len(files)
}

func (this *TextureAtlas) Get(name string) (result AtlasTexture) {
	var rect, readResult = this.images[name]
	if readResult {
		result.Image = this.texture
		result.Rect = rect
	}
	return
}

func (this *TextureAtlas) GetCurrentPos() Int2 {
	return this.currentPos
}

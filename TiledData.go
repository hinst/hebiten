package hebiten

import (
	"encoding/json"
	. "hgo"
	"io/ioutil"
	"path/filepath"
)

type TTiledData struct {
	Data     TTiledDocumentRoot
	Ground   TIntArray2d
	Objects  TIntArray2d
	Markers  TIntArray2d
	FilePath string

	TilesFilePath string
}

func (this *TTiledData) Create() *TTiledData {
	return this
}

func (this *TTiledData) Read() (result bool) {
	var data, readFileResult = ioutil.ReadFile(this.FilePath)
	result = readFileResult == nil
	if result {
		this.ReadData(data)
	}
	return
}

func (this *TTiledData) ReadData(data []byte) {
	AssertResult(json.Unmarshal(data, &this.Data))
	this.WriteLog("map size: " + IntToStr(this.Data.Width) + "," + IntToStr(this.Data.Height) +
		"; count of layers: " + IntToStr(len(this.Data.Layers)) +
		"; count of tiles: " + IntToStr(len(this.Data.Tilesets[0].Tiles)))
	this.Ground = this.LoadLayer("ground")
	this.Objects = this.LoadLayer("objects")
	this.Markers = this.LoadLayer("mission")
}

func (this *TTiledData) GetSize() Int2 {
	return Int2{X: this.Data.Width, Y: this.Data.Height}
}

func (this *TTiledData) GetLayer(name string) (result *TTiledDocumentLayer) {
	for _, layer := range this.Data.Layers {
		if layer.Name == name {
			result = &layer
			break
		}
	}
	return
}

func (this *TTiledData) LoadLayer(name string) (result TIntArray2d) {
	var layer = this.GetLayer(name)
	if layer != nil {
		result.SetSize(this.GetSize())
		result.Size.RunForEach(func(position Int2) {
			var value = layer.Data[position.GetLinearIndex(result.Size.X)]
			result.SetTile(position, value)
		})
	}
	return
}

func (this *TTiledData) GetTileImageFullName(id int) (result string) {
	var tile, found = this.Data.Tilesets[0].Tiles[IntToStr(id)]
	if found {
		result = tile.Image
	}
	return
}

func (this *TTiledData) GetTileImageName(id int) string {
	return filepath.Base(this.GetTileImageFullName(id - 1))
}

func (this *TTiledData) WriteLog(s string) {
}

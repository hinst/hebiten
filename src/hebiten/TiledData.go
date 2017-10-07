package hebiten

import (
	"path/filepath"
)

type TTiledData struct {
	Data          TTiledDocumentRoot
	Ground        TIntArray2d
	Objects       TIntArray2d
	Markers       TIntArray2d
	FilePath      string
	TilesFilePath string
	WriteLog      func(s string)
	Tiles         map[int]string
}

func (this *TTiledData) Create() *TTiledData {
	if nil == this.WriteLog {
		this.WriteLog = func(s string) {}
	}
	return this
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

func (this *TTiledData) GetTileImageFullName(id int) (result string) {
	var tile, found = this.Tiles[id]
	if found {
		result = tile
	}
	return
}

func (this *TTiledData) GetTileImageName(id int) string {
	return filepath.Base(this.GetTileImageFullName(id - 1))
}

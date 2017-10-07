package hebiten

import (
	"encoding/json"
	. "hgo"
	"io/ioutil"
)

func (this *TTiledData) Read() (result bool) {
	if this.FilePath != "" {
		var data, readFileResult = ioutil.ReadFile(this.FilePath)
		result = readFileResult == nil
		if result {
			this.ReadData(data)
		}
	}
	if result {
		this.readTiles()
	}
	return
}

func (this *TTiledData) readTiles() {
	if this.TilesFilePath != "" {
		this.WriteLog("Now loading tiles from '" + this.TilesFilePath + "'")
		var data, readFileResult = ioutil.ReadFile(this.TilesFilePath)
		AssertResult(readFileResult)
		var tileObject = &TTiledDocumentTileset{}
		AssertResult(json.Unmarshal(data, tileObject))
		if false {
			this.WriteLog("Number of tile objects: " + IntToStr(len(tileObject.Tiles)))
		}
		this.Tiles = tileObject.ToMap()
		this.WriteLog("Number of tiles: " + IntToStr(len(this.Tiles)))
	}
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

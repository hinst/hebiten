package hebiten

import "hgo"

type TTiledDocumentTileset struct {
	Tiles map[string]TTiledDocumentImageTile `json:"tiles"`
}

func (this *TTiledDocumentTileset) ToMap() (result map[int]string) {
	result = make(map[int]string)
	for key, value := range this.Tiles {
		var keyInt = hgo.StrToInt(key)
		result[keyInt] = value.Image
	}
	return
}

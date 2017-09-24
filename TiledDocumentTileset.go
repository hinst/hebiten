package hebiten

type TTiledDocumentTileset struct {
	Tiles map[string]TTiledDocumentImageTile
}

func (this *TTiledDocumentTileset) ToMap() (result map[int]string) {
	result = make(map[int]string)
	// for _, v := range this.Tiles {
	// }
	return
}

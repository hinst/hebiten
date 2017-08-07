package hebiten

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
)

func LoadFont(filePath string) (result struct {
	Font   *truetype.Font
	Result error
}) {
	var fontData, fileResult = ioutil.ReadFile(filePath)
	if nil == fileResult {
		var font, fontResult = truetype.Parse(fontData)
		if nil == fontResult {
			result.Font = font
		} else {
			result.Result = fontResult
		}
	} else {
		result.Result = fileResult
	}
	return
}

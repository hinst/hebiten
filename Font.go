package hebiten

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
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

func GetFontLineHeight(fontFace font.Face, screenDensity float64) float64 {
	return float64(fontFace.Metrics().Height)/screenDensity + 1
}

package hebiten

import (
	"github.com/hajimehoshi/ebiten"
)

var DirectionKeys = []ebiten.Key{
	ebiten.KeyUp,
	ebiten.KeyDown,
	ebiten.KeyLeft,
	ebiten.KeyRight,
}

var Keys = append(DirectionKeys, []ebiten.Key{ebiten.KeyEnter, ebiten.KeyEscape, ebiten.KeyPageUp, ebiten.KeyPageDown,
	ebiten.KeyI, ebiten.KeyS, ebiten.KeyT}...)

type TInput struct {
	Keys    map[ebiten.Key]KeyStatus
	OnClick func(position Int2)

	PreviousLeftMouseButtonStatus bool
}

func (this *TInput) Create() *TInput {
	this.Keys = make(map[ebiten.Key]KeyStatus)
	for _, key := range Keys {
		this.Keys[key] = KeyStatus{Pressed: false, TimeSincePressed: 0}
	}
	return this
}

func (this *TInput) Update(deltaTime float64) {
	this.UpdateLeftMouseButton()
	this.UpdateKeys(deltaTime)
}

func (this *TInput) UpdateLeftMouseButton() {
	var currentLeftMouseButtonStatus = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if currentLeftMouseButtonStatus && !this.PreviousLeftMouseButtonStatus && this.OnClick != nil {
		var x, y = ebiten.CursorPosition()
		this.OnClick(Int2{X: x, Y: y})
	}
	this.PreviousLeftMouseButtonStatus = currentLeftMouseButtonStatus
}

func (this *TInput) UpdateKeys(deltaTime float64) {
	for key, value := range this.Keys {
		var pressed = ebiten.IsKeyPressed(key)
		if pressed {
			if !value.Pressed {
				value.TimeSincePressed = 0
			} else {
				value.TimeSincePressed += deltaTime
			}
		} else {
			value.TimeSincePressed = 0
		}
		value.Pressed = pressed
		this.Keys[key] = value
	}
}

func (this *TInput) CheckKeyPressed(key ebiten.Key) bool {
	var keyData = this.Keys[key]
	return keyData.Pressed && keyData.TimeSincePressed == 0
}

func GetDirectionFromKey(key ebiten.Key) (result Direction) {
	switch key {
	case ebiten.KeyUp:
		result = DirectionUp
	case ebiten.KeyDown:
		result = DirectionDown
	case ebiten.KeyLeft:
		result = DirectionLeft
	case ebiten.KeyRight:
		result = DirectionRight
	default:
		result = DirectionNone
	}
	return
}

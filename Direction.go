package hebiten

type Direction int

const (
	DirectionNone  Direction = 0
	DirectionUp    Direction = 1
	DirectionDown  Direction = 2
	DirectionLeft  Direction = 3
	DirectionRight Direction = 4
)

var DirectionString = map[Direction]string{
	DirectionNone:  "None",
	DirectionUp:    "Up",
	DirectionDown:  "Down",
	DirectionLeft:  "Left",
	DirectionRight: "Right",
}

func CheckIf4xDirection(direction Direction) bool {
	return direction == DirectionUp ||
		direction == DirectionDown ||
		direction == DirectionLeft ||
		direction == DirectionRight
}

func (a Direction) ToString() string {
	var text, found = DirectionString[a]
	if !found {
		text = "Unknown"
	}
	return text
}

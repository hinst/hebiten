package hebiten

import "hgo"

type Int2 struct {
	X, Y int
}

func (this Int2) GetLinearIndex(width int) int {
	return this.X + this.Y*width
}

func (this Int2) RunForEach(f func(position Int2)) {
	for y := 0; y < this.Y; y++ {
		for x := 0; x < this.X; x++ {
			f(Int2{x, y})
		}
	}
}

func (this Int2) GetProduct() int {
	return this.X * this.Y
}

func (this Int2) ToString() string {
	return hgo.IntToStr(this.X) + "," + hgo.IntToStr(this.Y)
}

func (this Int2) GetProductByInt(multiplier int) BigFloat2 {
	var m = float64(multiplier)
	return BigFloat2{X: float64(this.X) * m, Y: float64(this.Y) * m}
}

func (this Int2) ToBigFloat() BigFloat2 {
	return BigFloat2{
		X: float64(this.X),
		Y: float64(this.Y),
	}
}

func (this Int2) AddInts(x, y int) Int2 {
	return Int2{X: this.X + x, Y: this.Y + y}
}

func (this Int2) GetNearest4() []Int2 {
	return []Int2{
		this.AddInts(-1, 0),
		this.AddInts(1, 0),
		this.AddInts(0, -1),
		this.AddInts(0, 1),
	}
}

func (this Int2) Equals(a Int2) bool {
	return this.X == a.X && this.Y == a.Y
}

func (this Int2) CheckIfNearest4(a Int2) (result bool) {
	for _, item := range this.GetNearest4() {
		if item.Equals(a) {
			result = true
			break
		}
	}
	return
}

func (this Int2) GetFromDirection(a Direction) (result Int2) {
	switch a {
	case DirectionUp:
		result = Int2{X: 0, Y: -1}
	case DirectionDown:
		result = Int2{X: 0, Y: 1}
	case DirectionLeft:
		result = Int2{X: -1, Y: 0}
	case DirectionRight:
		result = Int2{X: 1, Y: 0}
	}
	return
}

func (this Int2) Add(a Int2) Int2 {
	return Int2{X: this.X + a.X, Y: this.Y + a.Y}
}

func (this Int2) MoveByDirection(direction Direction) (result Int2) {
	result = this.Add(Int2{}.GetFromDirection(direction))
	return
}

// Check pos is a valid position in 2D array of size = this
func (this Int2) CheckInRange(pos Int2) bool {
	return 0 <= pos.X && pos.X < this.X &&
		0 <= pos.Y && pos.Y < this.Y
}

package hebiten

type TIntArray2d struct {
	Size Int2
	Data []int
}

func (this *TIntArray2d) SetSize(value Int2) {
	this.Size = value
	this.Data = make([]int, value.GetProduct())
}

func (this *TIntArray2d) Get(position Int2) int {
	return this.Data[position.GetLinearIndex(this.Size.X)]
}

func (this *TIntArray2d) SetTile(position Int2, value int) {
	this.Data[position.GetLinearIndex(this.Size.X)] = value
}

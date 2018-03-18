package hebiten

import (
	"hgo"

	"github.com/hajimehoshi/ebiten"
)

type ListUI struct {
	ItemHeight        float64
	Top               float64
	Height            float64
	Gap               float64
	CurrentPage       int
	PageSize          int
	SelectedItemIndex int
	SelectionRect     SelectionRectUI
	BackgroundImage   AtlasTexture
	pageText          string
	Input             *TInput
	// This function is optional.
	ReceiveDrawEnd func()
	GetItemCount   func() int
	DrawItem       func(target *ebiten.Image, i int, rect TFloatRect)
	PageTextPrefix string
}

func (this *ListUI) Create() *ListUI {
	hgo.Assert(this.ItemHeight > 0)
	hgo.Assert(this.Height > 0)
	this.PageSize = int((this.Height - this.Gap) / (this.ItemHeight + this.Gap))
	this.SelectionRect.Create()
	this.SelectionRect.BackgroundImage = this.BackgroundImage
	if this.ReceiveDrawEnd == nil {
		this.ReceiveDrawEnd = func() {}
	}
	if this.PageTextPrefix == "" {
		this.PageTextPrefix = "page"
	}
	return this
}

func (this *ListUI) Update(deltaTime float64) {
	this.constrainIndexes()
	this.UpdatePage()
	this.UpdateSelection()
	this.SelectionRect.Update(deltaTime)
}

func (this *ListUI) Draw(target *ebiten.Image) {
	this.constrainIndexes()
	var point = BigFloat2{X: this.Gap, Y: this.Top + this.Gap}
	var countOfDisplayedItems = 0
	var count = this.GetItemCount()
	for i := this.CurrentPage * this.PageSize; i < count; i++ {
		if !(countOfDisplayedItems < this.PageSize) {
			break
		}
		var targetRect TFloatRect
		targetRect.SetLeftTopPoint(point)
		targetRect.W = float64(GetImageWidth(target)) - this.Gap*2
		targetRect.H = this.ItemHeight
		if i == this.SelectedItemIndex {
			this.SelectionRect.SetTarget(target)
			this.SelectionRect.Draw(targetRect)
		}
		this.DrawItem(target, i, targetRect)
		point.Y += this.ItemHeight + this.Gap
		countOfDisplayedItems++
	}
	this.ReceiveDrawEnd()
}

func (this *ListUI) GetPager() Pager {
	return Pager{Count: this.GetItemCount(), PageSize: this.PageSize}
}

func (this *ListUI) UpdatePage() {
	var delta = 0
	if this.Input.CheckKeyPressed(ebiten.KeyLeft) {
		delta = -1
	} else if this.Input.CheckKeyPressed(ebiten.KeyRight) {
		delta = 1
	}
	if delta != 0 {
		this.CurrentPage = this.GetPager().FlipPage(this.CurrentPage, delta)
		this.SelectedItemIndex = this.GetPager().FollowPage(this.CurrentPage, this.SelectedItemIndex)
		this.pageText = ""
	}
}

func (this *ListUI) UpdateSelection() {
	var delta = 0
	if this.Input.CheckKeyPressed(ebiten.KeyUp) {
		delta = -1
	} else if this.Input.CheckKeyPressed(ebiten.KeyDown) {
		delta = 1
	}
	if delta != 0 {
		this.SelectedItemIndex = this.GetPager().MoveSelection(this.SelectedItemIndex, delta)
		this.CurrentPage = this.GetPager().FollowItem(this.SelectedItemIndex)
		this.pageText = ""
	}
}

func (this *ListUI) constrainIndexes() {
	this.CurrentPage = this.GetPager().ConstrainPage(this.CurrentPage)
	this.SelectedItemIndex = this.GetPager().ConstrainSelection(this.SelectedItemIndex)
}

func (this *ListUI) GetPageText() string {
	if "" == this.pageText {
		this.pageText = this.PageTextPrefix + " " + hgo.IntToStr(this.CurrentPage+1) + "/" + hgo.IntToStr(this.GetPager().CountOfPages())
	}
	return this.pageText
}

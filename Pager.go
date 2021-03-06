package hebiten

import (
	"hgo"
)

type Pager struct {
	Count    int
	PageSize int
}

func (this Pager) CountOfPages() (result int) {
	result = this.Count / this.PageSize
	if this.Count > 0 && this.Count%this.PageSize == 0 {
		result -= 1
	}
	result++
	return
}

// Returns page index
func (this Pager) FollowItem(itemIndex int) int {
	return itemIndex / this.PageSize
}

// From pageIndex get [lower .. upper) item indexes to be present on this page
func (this Pager) GetPageRange(pageIndex int) (result Int2) {
	result.X = pageIndex * this.PageSize
	result.Y = result.X + this.PageSize
	return
}

// Returns item index
func (this Pager) FollowPage(pageIndex, itemIndex int) int {
	var pageRange = this.GetPageRange(pageIndex)
	var newItemIndex = hgo.LockIntBetween(pageRange.X, itemIndex, pageRange.Y)
	return newItemIndex
}

func (this Pager) FlipPage(pageIndex int, delta int) int {
	return this.ConstrainPage(pageIndex + delta)
}

func (this Pager) ConstrainPage(pageIndex int) int {
	return hgo.LockIntBetween(0, pageIndex, this.CountOfPages())
}

func (this Pager) MoveSelection(selectedItemIndex int, delta int) int {
	return this.ConstrainSelection(selectedItemIndex + delta)
}

func (this Pager) ConstrainSelection(selectedItemIndex int) int {
	return hgo.LockIntBetween(0, selectedItemIndex, this.Count)
}

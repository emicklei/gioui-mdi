package mdi

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
)

// Inset adds space around a widget by decreasing its maximum
// constraints. The minimum constraints will be adjusted to ensure
// they do not exceed the maximum.
type Inset struct {
	Top, Bottom, Left, Right int
}

// Layout a widget.
func (in Inset) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	top := in.Top
	right := in.Right
	bottom := in.Bottom
	left := in.Left
	mcs := gtx.Constraints
	mcs.Max.X -= left + right
	if mcs.Max.X < 0 {
		left = 0
		right = 0
		mcs.Max.X = 0
	}
	if mcs.Min.X > mcs.Max.X {
		mcs.Min.X = mcs.Max.X
	}
	mcs.Max.Y -= top + bottom
	if mcs.Max.Y < 0 {
		bottom = 0
		top = 0
		mcs.Max.Y = 0
	}
	if mcs.Min.Y > mcs.Max.Y {
		mcs.Min.Y = mcs.Max.Y
	}
	gtx.Constraints = mcs
	trans := op.Offset(image.Point{X: left, Y: top}).Push(gtx.Ops)
	dims := w(gtx)
	trans.Pop()
	return layout.Dimensions{
		Size:     dims.Size.Add(image.Point{X: right + left, Y: top + bottom}),
		Baseline: dims.Baseline + bottom,
	}
}

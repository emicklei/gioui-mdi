package mdi

import (
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/widget/material"
)

type MDIWindowStack struct {
	list  []*MDIWindow
	theme *material.Theme
}

func NewWindowStack(theme *material.Theme) *MDIWindowStack {
	s := &MDIWindowStack{theme: theme}

	// TEMP
	test := NewMDIWindow(
		50, 100, 600, 300,
		material.Body1(s.theme, "gio ui MDI test").Layout)
	s.list = append(s.list, test)
	{
		test := NewMDIWindow(
			250, 300, 600, 300,
			material.Body1(s.theme, "gio ui MDI test too").Layout)
		s.list = append(s.list, test)
	}
	// END TEMP
	return s
}

func (s *MDIWindowStack) HandleDragging(ops *op.Ops, q event.Queue) {
	hitIndex := -1
	for i, each := range s.list {
		for _, ev := range q.Events(each.drag) {
			if each.drag.HandleEvent(ev, each) {
				hitIndex = i
			}
		}
		// Confine the area of interest to header of the window.
		area := clip.Rect(each.draggableBounds()).Push(ops)
		// Declare the tag.
		pointer.InputOp{
			Tag:   each.drag,
			Types: pointer.Press | pointer.Release | pointer.Drag,
		}.Add(ops)
		area.Pop()
	}
	if hitIndex != -1 {
		// put this hit on the back
		newlist := []*MDIWindow{}
		for i := 0; i < len(s.list); i++ {
			if hitIndex != i {
				newlist = append(newlist, s.list[i])
			}
		}
		newlist = append(newlist, s.list[hitIndex])
		//swap
		s.list = newlist
	}
}

func (s *MDIWindowStack) Layout(t *material.Theme, gtx layout.Context) layout.Dimensions {
	for _, each := range s.list {
		each.Layout(t, gtx)
	}
	return layout.Dimensions{}
}

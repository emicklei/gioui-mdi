package mdi

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

const (
	windowBorder   = 5.0
	windowHeader   = 35.0
	windowRounding = 20.0
)

type MDIWindow struct {
	// offset
	left int
	top  int
	// size
	width  int
	height int
	// root widget of the content
	content layout.Widget

	// moving
	pressed         bool
	pressedAt       f32.Point
	preDragLocation f32.Point

	drag *Draggable
}

func NewMDIWindow(left, top, width, height int, content layout.Widget) *MDIWindow {
	return &MDIWindow{
		left: left, top: top, width: width, height: height,
		content: content,
		drag:    new(Draggable),
	}
}

func (m *MDIWindow) draggableBounds() image.Rectangle {
	return image.Rect(int(m.left), int(m.top), int(m.left+m.width), int(m.top+windowHeader))
}

func (m *MDIWindow) setLocationTo(location f32.Point) {
	m.left = int(location.X)
	m.top = int(location.Y)
}

// Position is part of Positionable
func (m *MDIWindow) Position() f32.Point {
	return f32.Pt(float32(m.left), float32(m.top))
}

// SetPosition is part of Positionable
func (m *MDIWindow) SetPosition(location f32.Point) {
	m.left = int(location.X)
	m.top = int(location.Y)
}

// Layout lays out the counter and handles input.
func (m *MDIWindow) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	// paint the background of the mdiwindow, with header and icons
	// then layout the content
	m.drawBackground(gtx)

	offset := Inset{
		Left:   int(m.left + windowBorder),
		Top:    int(m.top + windowHeader),
		Right:  int(m.left + m.width - windowBorder),
		Bottom: int(m.top + m.height - windowBorder),
	}
	return offset.Layout(gtx, m.content)
}

func (m *MDIWindow) drawBackground(gtx layout.Context) {
	w := clip.RRect{
		Rect: image.Rectangle{
			Min: image.Point{X: m.left, Y: m.top},
			Max: image.Point{X: m.left + m.width, Y: m.top + m.height}},
		NE: windowRounding * .5, NW: windowRounding * .5, SE: windowRounding * .5, SW: windowRounding * .5,
	}.Op(gtx.Ops)
	paint.FillShape(gtx.Ops, color.NRGBA{A: 0xff, R: 0x55, G: 0x55}, w)

	w = clip.RRect{
		Rect: image.Rectangle{
			Min: image.Point{X: m.left + windowBorder, Y: m.top + windowHeader},
			Max: image.Point{X: m.left + m.width - windowBorder, Y: m.top + m.height - windowBorder}},
		NE: 0.0, NW: 0.0, SE: 0.0, SW: 0.0,
	}.Op(gtx.Ops)
	paint.FillShape(gtx.Ops, color.NRGBA{A: 0xff, R: 0xff, G: 0xff, B: 0xff}, w)

}

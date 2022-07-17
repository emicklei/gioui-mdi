package mdi

import (
	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
)

type Positionable interface {
	Position() f32.Point
	SetPosition(where f32.Point)
}

type Draggable struct {
	pressed                  bool
	pressedInputPosition     f32.Point
	beforeDragObjectPosition f32.Point
}

func (d *Draggable) HandleEvent(ev event.Event, object Positionable) bool {
	hit := false
	if x, ok := ev.(pointer.Event); ok {
		switch x.Type {
		case pointer.Press:
			hit = true
			d.pressed = true
			d.pressedInputPosition = x.Position
			d.beforeDragObjectPosition = object.Position()
		case pointer.Release:
			d.pressed = false
		case pointer.Drag:
			if d.pressed {
				offset := x.Position.Sub(d.pressedInputPosition)
				object.SetPosition(d.beforeDragObjectPosition.Add(offset))
			}
		}
	}
	return hit
}

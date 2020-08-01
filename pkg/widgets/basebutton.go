package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
)

type BaseButton struct {
	BaseWidget
	pressed bool
}

func NewBaseButton(ID string, x float64, y float64, width int, height int, handler func(event events.IEvent)(bool, error)) *BaseButton {
	bb := BaseButton{}
	bb.BaseWidget = *NewBaseWidget(ID, x, y, width, height,handler)
	bb.pressed = false
	bb.stateChangedSinceLastDraw = true
	//bb.eventHandler = handler

	go bb.ListenToIncomingEvents()
	return &bb
}

func (b *BaseButton) Draw(screen *ebiten.Image) error {
	return nil
}

func (b *BaseButton) HandleEvent(event events.IEvent) (bool, error) {


	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			mouseEvent := event.(events.MouseEvent)
			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				log.Debugf("BaseButton::HandleEvent %s", b.ID)
				log.Debugf("BUTTON DOWN!!!")
				b.hasFocus = true
				// if already pressed, then skip it.. .otherwise lots of repeats.
				if !b.pressed {
					b.pressed = true
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	case events.EventTypeButtonUp:
		{
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if b.ContainsCoords(mouseEvent.X, mouseEvent.Y) {
				log.Debugf("BUTTON UP!!!")
				log.Debugf("BaseButton::HandleEvent %s", b.ID)
				b.hasFocus = true
				if b.pressed {
					// do generic button stuff here.
					b.pressed = false
					b.stateChangedSinceLastDraw = true
				}
			}
		}
	}
	return false, nil
}

type IButton interface {
	HandleEvent(event events.IEvent) error
	Draw(screen *ebiten.Image) error
}

package widgets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"image/color"
)

// Panel has a position, width and height.
// Panels contain other widgets (and also other panels).
type Panel struct {
	BaseWidget

	// panel can contain panels.
	panels []Panel

	// panel can also contain widgets
	// widgets []IWidget   figure out what interface will look like.

	buttons []IButton
}

func NewPanel(x float64, y float64, width int, height int) Panel {
	p := Panel{}
	p.BaseWidget = NewBaseWidget(x, y, width, height)
	p.buttons = []IButton{}
	
	return p
}

// AddButton adds a already created button.
func (p *Panel) AddButton(b IButton) error {
	p.buttons = append(p.buttons, b)
	return nil
}

// Draw renders all the widgets inside the panel (and the panel itself.. .if there is anything to it?)
func (p *Panel) Draw(screen *ebiten.Image) error {

	// colour background of panel first, just so we can see it.
	_ = p.rectImage.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	for _, b := range p.buttons {
		b.Draw(p.rectImage)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	_ = screen.DrawImage(p.rectImage, op)
	return nil
}

func (p *Panel) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			p.HandleMouseEvent(event)
		}
	case events.EventTypeButtonUp:
		{
			p.HandleMouseEvent(event)
		}
	}

	return nil
}

func (p *Panel) HandleMouseEvent(event events.IEvent) error {
	inPanel, _ := p.BaseWidget.CheckMouseEventCoords(event)

	if inPanel {
		mouseEvent := event.(events.MouseEvent)
		log.Debugf("HandleMouseEvent panel %f %f", mouseEvent.X, mouseEvent.Y)
		localCoordMouseEvent := p.GenerateLocalCoordMouseEvent(mouseEvent)
		for _, button := range p.buttons {
			button.HandleEvent(localCoordMouseEvent)
		}
	}

	// should propagate to children nodes?
	return nil
}
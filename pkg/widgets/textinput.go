package widgets

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kpfaulkner/goui/pkg/common"
	"github.com/kpfaulkner/goui/pkg/events"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image/color"
	_ "image/png"
)


type TextInput struct {
	BaseWidget

	text             string
	backgroundColour color.RGBA
	fontInfo         *common.Font
	uiFont           font.Face

}

func NewTextInput(ID string, x float64, y float64, width int, height int, backgroundColour *color.RGBA) TextInput {
	t := TextInput{}
	t.BaseWidget = NewBaseWidget(ID, x, y, width, height)
	t.text = ""
  t.stateChangedSinceLastDraw = true

	if backgroundColour != nil {
		t.backgroundColour = *backgroundColour
	} else {
		t.backgroundColour = color.RGBA{0,0xff,0,0xff}
	}

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	t.uiFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return t
}

func (t *TextInput) HandleEvent(event events.IEvent) error {

	eventType := event.EventType()
	switch eventType {
	case events.EventTypeButtonDown:
		{
			mouseEvent := event.(events.MouseEvent)

			// check click is in button boundary.
			if t.ContainsCoords(mouseEvent.X, mouseEvent.Y, true) {
				t.hasFocus = true
				t.stateChangedSinceLastDraw = true
				// then do application specific stuff!!
				if ev, ok := t.eventRegister[event.EventType()]; ok {
					ev(event)
				}
			}
		}
	case events.EventTypeKeyboard:
		{
			// check if has focus....  if so, can potentially add to string?
			if t.hasFocus {
				keyboardEvent := event.(events.KeyboardEvent)
				t.text = t.text + keyboardEvent.Character
				t.stateChangedSinceLastDraw = true
			}
		}
	}
	return nil
}

func (t *TextInput) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.X, t.Y)

	if t.stateChangedSinceLastDraw {
		log.Debugf("textinput text %s", t.text)
		// how often do we update this?
		emptyImage, _ := ebiten.NewImage(t.Width, t.Height, ebiten.FilterDefault)
		_ = emptyImage.Fill(t.backgroundColour)
		t.rectImage = emptyImage
		text.Draw(t.rectImage, t.text, t.uiFont, 10, 10, color.Black)
		t.stateChangedSinceLastDraw = false
	}

	// if state changed since last draw, recreate colour etc.
	_ = screen.DrawImage(t.rectImage, op)

	return nil
}
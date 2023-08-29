package ecascreen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var textViewMsgs *tview.TextView

func GetMsgsHistory() tview.Primitive {
	textViewMsgs = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetTextColor(tcell.ColorNavajoWhite)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(textViewMsgs, 0, 2, true)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		switch key {
		case tcell.KeyUp:
			ScrollTextViewUp(textViewMsgs)
			return nil
		case tcell.KeyDown:
			ScrollTextViewDown(textViewMsgs)
			return nil
		}
		return event
	})
	return flex
}

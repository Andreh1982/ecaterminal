package ecascreen

import "github.com/rivo/tview"

func CenterScreen(widget tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(nil, 0, 1, false).
				AddItem(widget, width, 0, true).
				AddItem(nil, 0, 1, false),
			height, 0, true).
		AddItem(nil, 0, 1, false)
}

func ScrollTextViewUp(caller *tview.TextView) {
	cy, _ := caller.GetScrollOffset()
	cy--
	caller.ScrollTo(cy, 0)
}

func ScrollTextViewDown(caller *tview.TextView) {
	cy, _ := caller.GetScrollOffset()
	cy++
	caller.ScrollTo(cy, 0)
}

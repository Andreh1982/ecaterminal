package ecascreen

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gdamore/tcell/v2"
	"github.com/go-skynet/go-llama.cpp"
	"github.com/rivo/tview"
)

type Task struct {
	name   string
	widget tview.Primitive
}

// var (
// 	mHistory structx.Messages
// )

const bannerName = "Tech Tweakers Solutions"

func Blur() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	tview.Styles.ContrastBackgroundColor = tcell.ColorDarkRed
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorWhiteSmoke
	tview.Styles.BorderColor = tcell.ColorPapayaWhip
	tview.Styles.TitleColor = tcell.ColorFloralWhite
	tview.Styles.GraphicsColor = tcell.ColorRed
	tview.Styles.PrimaryTextColor = tcell.ColorGhostWhite
	tview.Styles.SecondaryTextColor = tcell.ColorRed
	tview.Styles.TertiaryTextColor = tcell.ColorGreen
	tview.Styles.InverseTextColor = tcell.ColorDeepSkyBlue
	tview.Styles.ContrastSecondaryTextColor = tcell.ColorDarkCyan
}

func tasks(l *llama.LLama, app *tview.Application) []Task {
	tasks := []Task{
		{name: "Chat Messages", widget: ChatMessages(l, app)},
		{name: "Messages History", widget: GetMsgsHistory()},
		{name: "Request Configuration", widget: RequestSetup()},
	}
	return tasks
}

func Screen(l *llama.LLama) {
	configStyles()
	app := createApp(l)
	done := make(chan struct{})

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(
			signals,
			syscall.SIGILL, syscall.SIGINT, syscall.SIGTERM,
		)
		<-signals
		app.Stop()
	}()
	<-done
}

func configStyles() {
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorBlack,
		ContrastBackgroundColor:     tcell.ColorDarkRed,
		MoreContrastBackgroundColor: tcell.ColorWhiteSmoke,
		BorderColor:                 tcell.ColorPapayaWhip,
		TitleColor:                  tcell.ColorFloralWhite,
		GraphicsColor:               tcell.ColorRed,
		PrimaryTextColor:            tcell.ColorGhostWhite,
		SecondaryTextColor:          tcell.ColorRed,
		TertiaryTextColor:           tcell.ColorGreen,
		InverseTextColor:            tcell.ColorDeepSkyBlue,
		ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
	}
}

func createApp(l *llama.LLama) *tview.Application {
	app := tview.NewApplication()
	tasks := tasks(l, app)
	pages := createPages(tasks)
	menu := createSidebar(tasks, pages, app) // Pass app as a parameter
	flex := tview.NewFlex()
	flex.AddItem(menu, len(bannerName)+18, 1, false)
	flex.AddItem(pages, 0, 1, true)
	app.SetRoot(flex, true)
	app.EnableMouse(true)
	app.SetFocus(menu)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyTAB {
			if app.GetFocus() == pages {
				// Handle Tab key inside the app
				// Implement your logic here
				// For example, switch focus to the next element in the app
				return nil
			}
		} else if key == tcell.KeyESC {
			if app.GetFocus() == menu {
				// Handle ESC key when menu has focus
				// Switch focus to the app
				app.SetFocus(pages)
				return nil
			} else {
				// Handle ESC key when app has focus
				// Switch focus to the menu
				app.SetFocus(menu)
				return nil
			}
		}
		// For other keys, let the application handle them
		return event
	})
	return app
}

// createPages creates the pages
func createPages(tasks []Task) *tview.Pages {
	pages := tview.NewPages()
	for i, t := range tasks {
		var view tview.Primitive
		if t.widget != nil {
			view = t.widget
		} else {
			view = tview.NewTextView().SetText(t.name)
		}
		pages.AddPage(t.name, view, true, i == 0)
	}
	pages.SetBorder(true)
	return pages
}

func createSidebar(tasks []Task, pages *tview.Pages, app *tview.Application) tview.Primitive {
	menu := createMenu(tasks, pages, app) // Pass app as a parameter
	frame := tview.NewFrame(menu)
	frame.SetBorder(true)
	frame.SetBorders(0, 0, 1, 1, 1, 1)
	frame.AddText("ECATERMINAL - v1.0.0", true, tview.AlignCenter, tcell.ColorWhite)
	frame.AddText(bannerName, true, tview.AlignCenter, tcell.ColorRed)
	divbar := "------------------------------------------"
	frame.AddText(divbar, true, tview.AlignCenter, tcell.ColorWhite)
	frame.AddText("Tab Choose | Enter Select | Esc Menu", true, tview.AlignCenter, tcell.ColorGreen)
	return frame
}

func createMenu(tasks []Task, pages *tview.Pages, app *tview.Application) *tview.List {
	menu := tview.NewList()
	menuWidth := 0
	for i, t := range tasks {
		if len(t.name) > menuWidth {
			menuWidth = len(t.name)
		}
		menu.AddItem(t.name, "", rune(i+'0'), func() {
			pages.SwitchToPage(t.name)
			app.SetFocus(pages)
		})
	}
	menu.ShowSecondaryText(false)
	menu.SetChangedFunc(func(index int, task string, secondaryText string, shortcut rune) {
		pages.SwitchToPage(task)
	})
	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyTAB {
			if app.GetFocus() == pages {
				// Handle Tab key inside the app
				// Implement your logic here
				// For example, switch focus to the next element in the app
				return nil
			}
		} else if key == tcell.KeyESC {
			if app.GetFocus() == menu {
				// Handle ESC key when menu has focus
				// Switch focus to the app
				app.SetFocus(pages)
				return nil
			} else {
				// Handle ESC key when app has focus
				// Switch focus to the menu
				app.SetFocus(menu)
				return nil
			}
		} else if key == tcell.KeyEnter {
			if app.GetFocus() == menu {
				// Handle Enter key when menu has focus
				// Get the currently selected menu item
				index := menu.GetCurrentItem()
				item, _ := menu.GetItemText(index)
				// Switch to the corresponding page
				pages.SwitchToPage(item)
				app.SetFocus(pages)
				return nil
			}
		}
		// For other keys, let the application handle them
		return event
	})
	return menu
}

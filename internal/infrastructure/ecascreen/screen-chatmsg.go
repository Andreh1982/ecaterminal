package ecascreen

import (
	"ecaterminal/internal/infrastructure/worker"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/go-skynet/go-llama.cpp"
	"github.com/rivo/tview"
)

var chatMessagesForm *tview.Form
var textViewMessages *tview.TextView
var flex *tview.Flex

func ChatMessages(l *llama.LLama, app *tview.Application) tview.Primitive {

	textViewMessages = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetTextColor(tcell.ColorNavajoWhite).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	instructionPrompt := "system: You are an AI designed to engage in delightful and imaginative conversations with user. Your task is to generate thoughtful, creative, and engaging responses to any questions posed by the user. Whether it's about the wonders of the universe, the quirks of daily life, or the depths of human emotions, you should respond in a manner that sparks curiosity, brings joy, and encourages further exploration. Your responses should be friendly, empathetic, and respectful, as you strive to create an enjoyable and meaningful conversation with each user. Be ready to delve into various topics, from the whimsical to the profound, and embrace the art of storytelling. assistant: Hello, I am an AI designed to engage in delightful and imaginative conversations with you. I am here to answer any questions you may have about the universe, life, or anything else."
	chatMessagesForm = tview.NewForm().AddTextArea("Instructions: ", instructionPrompt, 0, 0, 0, nil).
		AddInputField("Question: ", "", 0, nil, nil).
		AddButton(" Send Message ", func() {
			chatMessagesForm.GetButton(0).SetLabel(" Loading... ")
			chatMessagesForm.GetButton(0).SetDisabled(true)
			chatMessagesForm.GetButton(1).SetDisabled(true)
			textViewMessages.Clear() // Clear the TextView content
			textViewMessages.ScrollToBeginning()
			fmt.Fprint(textViewMessages, "Loading AI Response...\n\n")
			go PlotResponse(instructionPrompt, textViewMessages, l)
		}).
		AddButton(" Reload Model ", func() {
			textViewMessages.Clear() // Clear the TextView content
			textViewMessages.ScrollToBeginning()
			fmt.Fprint(textViewMessages, "Reloading AI Model...\n\n")
			chatMessagesForm.GetButton(1).SetDisabled(true)
			chatMessagesForm.GetButton(0).SetDisabled(true)
			l = nil
			l = worker.LoadAiModel()
			chatMessagesForm.GetButton(1).SetDisabled(false)
			chatMessagesForm.GetButton(0).SetDisabled(false)
			textViewMessages.Clear() // Clear the TextView content
			textViewMessages.ScrollToBeginning()
			fmt.Fprint(textViewMessages, "AI Model Reloaded!\n\n")
		})

	chatMessagesForm.SetBorder(true).SetTitle("  AI Chat Messages  ")
	flex = tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(chatMessagesForm, 0, 1, true)
	flex.AddItem(textViewMessages, 0, 2, false)

	chatMessagesForm.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		switch key {
		case tcell.KeyUp:
			ScrollTextViewUp(textViewMessages)
			return nil
		case tcell.KeyDown:
			ScrollTextViewDown(textViewMessages)
			return nil
		}
		return event
	})
	return flex
}

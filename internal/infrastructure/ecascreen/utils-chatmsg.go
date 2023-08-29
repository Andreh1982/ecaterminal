package ecascreen

import (
	"ecaterminal/internal/infrastructure/worker"

	"github.com/go-skynet/go-llama.cpp"
	"github.com/rivo/tview"
)

func PlotResponse(instructionPrompt string, textViewMessages *tview.TextView, l *llama.LLama) {
	question := chatMessagesForm.GetFormItemByLabel("Question: ").(*tview.InputField).GetText()

	setupRequest := ResquestSetup()

	// fmt.Fprint(textViewMessages, setupRequest, " Question: "+question)

	worker.WorkerLlama(instructionPrompt, textViewMessages, l, question, setupRequest)

	chatMessagesForm.GetButton(0).SetLabel(" Send Message ")
	chatMessagesForm.GetButton(0).SetDisabled(false)
	chatMessagesForm.GetButton(1).SetDisabled(false)
	chatMessagesForm.SetFocus(1)

}

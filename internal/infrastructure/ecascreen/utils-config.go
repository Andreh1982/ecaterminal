package ecascreen

import (
	"ecaterminal/internal/infrastructure/structx"
	"strconv"

	"github.com/rivo/tview"
)

func ResquestSetup() (makeRequest structx.RequestConfig) {
	_, model := modelsType.GetFormItemByLabel("Available Models: ").(*tview.DropDown).GetCurrentOption()
	temperature, _ := strconv.ParseFloat(modelsType.GetFormItemByLabel("Temperature: ").(*tview.InputField).GetText(), 64)
	maxtokens, _ := strconv.Atoi(modelsType.GetFormItemByLabel("Max Tokens: ").(*tview.InputField).GetText())
	contextSize, _ := strconv.Atoi(modelsType.GetFormItemByLabel("Context Size: ").(*tview.InputField).GetText())
	n, _ := strconv.Atoi(modelsType.GetFormItemByLabel("N: ").(*tview.InputField).GetText())
	f16, _ := strconv.ParseBool(modelsType.GetFormItemByLabel("F16: ").(*tview.InputField).GetText())
	repeatPenalty, _ := strconv.ParseFloat(modelsType.GetFormItemByLabel("Repeat Penalty: ").(*tview.InputField).GetText(), 64)
	frequencyPenalty, _ := strconv.ParseFloat(modelsType.GetFormItemByLabel("Frequency Penalty: ").(*tview.InputField).GetText(), 64)
	presencePenalty, _ := strconv.ParseFloat(modelsType.GetFormItemByLabel("Presence Penalty: ").(*tview.InputField).GetText(), 64)
	seed, _ := strconv.Atoi(modelsType.GetFormItemByLabel("Seed: ").(*tview.InputField).GetText())
	step, _ := strconv.Atoi(modelsType.GetFormItemByLabel("Step: ").(*tview.InputField).GetText())
	topK, _ := strconv.Atoi(modelsType.GetFormItemByLabel("Top K: ").(*tview.InputField).GetText())
	topP, _ := strconv.ParseFloat(modelsType.GetFormItemByLabel("Top P: ").(*tview.InputField).GetText(), 64)
	instruction := modelsType.GetFormItemByLabel("Instruction: ").(*tview.InputField).GetText()
	prompt := modelsType.GetFormItemByLabel("Prompt: ").(*tview.InputField).GetText()
	stop := []string{".\n"}

	makeRequest = structx.RequestConfig{
		Model:            model,
		Temperature:      temperature,
		Maxtokens:        maxtokens,
		ContextSize:      contextSize,
		N:                n,
		F16:              f16,
		RepeatPenalty:    repeatPenalty,
		FrequencyPenalty: frequencyPenalty,
		PresencePenalty:  presencePenalty,
		Seed:             seed,
		Step:             step,
		TopK:             topK,
		TopP:             topP,
		Instruction:      instruction,
		Prompt:           prompt,
		SetStopWords:     stop,
	}

	return makeRequest
}

package worker

import (
	"ecaterminal/internal/infrastructure/structx"
	"flag"
	"fmt"
	"os"
	"runtime"

	llama "github.com/go-skynet/go-llama.cpp"
	"github.com/rivo/tview"
)

var (
	threads   = 4
	tokens    = 128
	gpulayers = 0
)

func LoadAiModel() (l *llama.LLama) {

	var model string

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&model, "m", "./models/llama-2-7b-chat.ggmlv3.q4_0.bin", "path to q4_0.bin model file to load")
	flags.IntVar(&gpulayers, "ngl", 0, "Number of GPU layers to use")
	flags.IntVar(&threads, "t", runtime.NumCPU(), "number of threads to use during computation")
	flags.IntVar(&tokens, "n", 512, "number of tokens to predict")

	err := flags.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Parsing program arguments failed: ", err)
		os.Exit(1)
	}
	l, err = llama.New(model, llama.EnableF16Memory, llama.SetContext(1024), llama.EnableEmbeddings, llama.SetGPULayers(gpulayers))
	if err != nil {
		fmt.Println("Loading the model failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Model loaded successfully.")

	return l
}

func WorkerLlama(instructionPrompt string, textViewMessages *tview.TextView, l *llama.LLama, question string, setupRequest structx.RequestConfig) {

	send2AI := instructionPrompt + " user: " + question
	msg, err := l.Predict(send2AI, llama.SetTokenCallback(func(token string) bool { return true }),
		llama.SetTokens(setupRequest.Maxtokens),
		llama.SetThreads(8),
		llama.SetTopK(setupRequest.TopK),
		llama.SetTopP(setupRequest.TopP),
		llama.SetStopWords(setupRequest.SetStopWords...),
		llama.SetTemperature(setupRequest.Temperature),
		llama.SetNKeep(setupRequest.N),
		llama.SetSeed(setupRequest.Seed),
		llama.SetPresencePenalty(setupRequest.RepeatPenalty),
		llama.SetFrequencyPenalty(setupRequest.FrequencyPenalty),
		llama.EnablePromptCacheAll,
		llama.SetPathPromptCache("./cache"),
		llama.SetStopWords("user:", "User:", "system:", "System:", ". \n"),
	)
	if err != nil {
		panic(err)
	}

	msg = msg + "."

	textViewMessages.Clear()
	fmt.Fprint(textViewMessages, "AI Response:\n---\n", msg)
	fmt.Fprint(textViewMessages, "\n--- \n")
	fmt.Fprint(textViewMessages, "Finished AI Response.")
}

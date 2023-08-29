package structx

type RequestConfig struct {
	Model            string   `json:"model"`
	Temperature      float64  `json:"temperature"`
	Maxtokens        int      `json:"maxtokens"`
	N                int      `json:"n"`
	F16              bool     `json:"f16"`
	RepeatPenalty    float64  `json:"repeat_penalty"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Seed             int      `json:"seed"`
	Step             int      `json:"step"`
	TopK             int      `json:"top_k"`
	TopP             float64  `json:"top_p"`
	SetStopWords     []string `json:"stop"`
	Instruction      string   `json:"instruction"`
	Prompt           string   `json:"prompt"`
	Size             string   `json:"size"`
	ResponseFormat   string   `json:"response_format"`
	Language         string   `json:"language"`
	File             string   `json:"file"`
	ContextSize      int      `json:"context_size"`
}

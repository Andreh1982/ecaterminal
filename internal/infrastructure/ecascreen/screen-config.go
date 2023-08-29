package ecascreen

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var modelsType *tview.Form

func RequestSetup() tview.Primitive {
	modelsType = tview.NewForm().AddDropDown("Available Models: ", []string{"ggml-gpt4all-j-v1.3-groovy", "ggml-mpt-7b-instruct", "guanaco-7B.ggmlv3.q2_K", "open-llama-7b-q4_0", "samantha-7B.ggmlv3.q2_K"}, 0, nil).
		AddInputField("Temperature: ", "0.2", 26, nil, nil).
		AddInputField("Max Tokens: ", "1024", 26, nil, nil).
		AddInputField("Context Size: ", "1024", 26, nil, nil).
		AddInputField("N: ", "1", 26, nil, nil).
		AddInputField("F16: ", "true", 26, nil, nil).
		AddInputField("Repeat Penalty: ", "2", 26, nil, nil).
		AddInputField("Frequency Penalty: ", "0.3", 26, nil, nil).
		AddInputField("Presence Penalty: ", "2", 26, nil, nil).
		AddInputField("Seed: ", "0", 26, nil, nil).
		AddInputField("Step: ", "0", 26, nil, nil).
		AddInputField("Top K: ", "20", 26, nil, nil).
		AddInputField("Top P: ", "0.50", 26, nil, nil).
		AddInputField("Stop: ", "?, !, \n, .\n, . ", 26, nil, nil).
		AddInputField("Instruction: ", "Always be polite. You are aware of everything.", 0, nil, nil).
		AddInputField("Prompt: ", "The prompt below is a question to answer, a task to complete, or a conversation to respond to. Decide which and write an appropriate response.", 0, nil, nil)

	modelsType.SetTitle(" Request Configuration ").SetBorder(true).SetBorderPadding(1, 1, 1, 1)

	helpText := tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetTextColor(tcell.ColorNavajoWhite)
	fmt.Fprint(helpText, "--- REQUEST PARAMETERS HELP --- \n\nMODEL: This parameter specifies the AI model to be used for generating responses.\n\nFILE: If provided, this parameter specifies a file to be used as input or output in the conversation. It can be used to read conversation history from a file or save the generated responses to a file.\n\nLANGUAGE: This parameter sets the language for the conversation. It helps the AI model understand and generate responses in the desired language. For example, en represents English.\n\nRESPONSEFORMAT: This parameter specifies the desired format for the AI model's response. It can be set to json to receive the response in JSON format, making it easier to parse and process the returned data.\n\nSIZE: This parameter determines the size or capacity of the model to be used. It can be set to medium to select a medium-sized model, but other options like small or large might also be available, depending on the specific implementation.\n\nPROMPT: The prompt parameter represents the initial message or prompt given to the AI model. It helps set the context or starting point of the conversation.\n\nINSTRUCTION: This parameter provides additional instructions or guidance to the AI model. You can use it to give specific directives or ask the model to approach the problem in a certain way.\n\nINPUT: The input parameter represents the user's input or message to be processed by the AI model. It contains the text or content that the model should consider when generating a response.\n\nStop: This parameter specifies a condition or message that indicates when the conversation should stop. When the AI model encounters the stop condition or message, it will cease generating further responses.\n\nMESSAGES: The messages parameter represents a collection of messages exchanged between the user and the AI model. Each message has a role (such as user or assistant) and content (the actual text of the message). Messages are typically used to maintain conversational context.\n\nSTREAM: When set to true, this parameter indicates that the conversation should be streamed. Streaming allows for a back-and-forth conversation experience, with immediate responses from the AI model.\n\nECHO: When set to true, the input messages are echoed back in the response. Setting it to false suppresses the echoing of messages, resulting in cleaner responses.\n\nTOPP: This parameter controls the diversity of the AI model's responses using the nucleus sampling method. A higher value, such as 0.8, allows for more diverse and creative responses.\n\nTOPK: This parameter controls the diversity of the AI model's responses by selecting from the top K tokens. A higher value, such as 20, allows for more diverse responses.\n\nTEMPERATURE: The temperature parameter controls the randomness of the AI model's responses. A higher value, such as 0.6, makes the responses more random, while a lower value, such as 0.2, makes them more focused and deterministic.\n\nMAXTOKENS: This parameter specifies the maximum number of tokens in the AI model's response. It helps limit the response length to a certain number of tokens, ensuring more concise or specific answers.\n\nN: This parameter specifies the number of responses to generate. For example, setting it to 5 will request the AI model to generate 5 different responses.\n\nBATCH: When multiple conversations are processed in parallel, this parameter determines the number of conversations to include in each batch.\n\nF16: Setting this parameter to true enables the use of 16-bit floating-point precision (f16) for processing. It can speed up computation in some cases.\n\nIGNOREOS: When set to true, the AI model ignores the end-of-sequence (EOS) token, allowing it to generate responses beyond the usual endpoint.\n\nREPEAT_PENALTY: This parameter controls the likelihood of the AI model repeating similar responses. A higher value penalizes repeated responses, encouraging more diverse answers.\n\nNKEEP: When using beam search decoding, this parameter specifies the number of hypotheses to keep during the decoding process. It helps maintain a diverse set of possible responses.\n\nMirostatEta, MirostatTau, Mirostat: These parameters are related to the mirostat generation method. The specific values and their impact depend on the implementation and may require domain-specific knowledge.\n\nSEED: Setting a specific value for the seed parameter allows for reproducibility in generating random numbers. Using the same seed value will produce the same sequence of random numbers.\n\nMODE: This parameter specifies a particular mode of operation. The specific values and their meanings may vary depending on the implementation and context.\n\nSTEP: The step parameter refers to a specific step number or iteration. Its meaning and usage can vary depending on the specific application or context.")

	flex = tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(modelsType, 0, 3, true)
	flex.AddItem(helpText, 0, 1, false)

	return flex
}

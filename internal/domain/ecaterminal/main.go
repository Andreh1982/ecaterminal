package ecaterminal

type Input struct {
	Repository Repository
}

type ecaterminal struct {
	repository Repository
}

type UseCases interface {
	MainMenu
	Lister
	Sender
}

type MainMenu interface {
	Screen
}

func New(input *Input) UseCases {
	return &ecaterminal{
		repository: input.Repository,
	}
}

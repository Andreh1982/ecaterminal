package ecaterminal

import (
	"ecaterminal/internal/domain/appcontext"
	"ecaterminal/internal/infrastructure/ecascreen"

	"github.com/go-skynet/go-llama.cpp"
)

type Screen interface {
	Screen(ctx appcontext.Context, l *llama.LLama)
}

func (a *ecaterminal) Screen(ctx appcontext.Context, l *llama.LLama) {
	ecascreen.Screen(l)
}

package ecaterminal

import (
	"ecaterminal/internal/domain/appcontext"
)

type Sender interface {
	SendMessage(ctx appcontext.Context, message string) error
}

func (a *ecaterminal) SendMessage(ctx appcontext.Context, message string) error {

	return nil
}

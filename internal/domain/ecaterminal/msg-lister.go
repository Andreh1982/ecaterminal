package ecaterminal

import (
	"ecaterminal/internal/domain/appcontext"
)

type Lister interface {
	ListMessages(ctx appcontext.Context) (*[]ChatPersistence, error)
}

func (l *ecaterminal) ListMessages(ctx appcontext.Context) (*[]ChatPersistence, error) {
	result, err := l.repository.List()
	return result, err
}

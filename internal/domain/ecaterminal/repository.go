package ecaterminal

type RepositoryReader interface {
	Find(entryID string) (*ChatPersistence, error)
	List() (*[]ChatPersistence, error)
}

type RepositoryWriter interface {
	Insert(ChatPersistence ChatPersistence) (*ChatPersistence, error)
	Delete(entryID string) error
	Upsert(ChatPersistence ChatPersistence) (*ChatPersistence, error)
}

type Repository interface {
	RepositoryReader
	RepositoryWriter
}

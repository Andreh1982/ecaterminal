package database

import (
	"ecaterminal/internal/domain/ecaterminal"
	"fmt"
)

func NewMemoryDatabase() ecaterminal.Repository {
	return &memoryDatabase{
		records: make(map[string]*ecaterminal.ChatPersistence),
	}
}

type memoryDatabase struct {
	records map[string]*ecaterminal.ChatPersistence
}

func (m *memoryDatabase) Find(key string) (*ecaterminal.ChatPersistence, error) {
	record := m.records[key]

	if record == nil {
		return nil, nil
	}

	return record, nil
}

func (m *memoryDatabase) Insert(ChatPersistence ecaterminal.ChatPersistence) (*ecaterminal.ChatPersistence, error) {
	m.records[fmt.Sprint(ChatPersistence.ID)] = &ChatPersistence
	return &ChatPersistence, nil
}

func (m *memoryDatabase) Upsert(applicationEntity ecaterminal.ChatPersistence) (*ecaterminal.ChatPersistence, error) {
	return m.Insert(applicationEntity)
}

func (m *memoryDatabase) Delete(key string) error {
	delete(m.records, key)

	return nil
}

func (m *memoryDatabase) List() (*[]ecaterminal.ChatPersistence, error) {
	var records []ecaterminal.ChatPersistence

	for _, record := range m.records {
		records = append(records, *record)
	}

	return &records, nil
}

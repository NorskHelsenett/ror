package oidctest

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
)

// MemoryStorageAdapter is an in-memory implementation of tokenstoragehelper.StorageAdapter.
type MemoryStorageAdapter struct {
	stored *tokenstoragehelper.KeyStorageProvider
}

// NewMemoryStorageAdapter creates an adapter pre-initialized with a generated key.
func NewMemoryStorageAdapter() (*MemoryStorageAdapter, error) {
	key, err := tokenstoragehelper.GenerateKey()
	if err != nil {
		return nil, err
	}

	provider := &tokenstoragehelper.KeyStorageProvider{
		LastRotation:     time.Now(),
		RotationInterval: 24 * time.Hour,
		NumKeys:          3,
		Keys: map[int]tokenstoragehelper.Key{
			1: key,
		},
	}

	return &MemoryStorageAdapter{stored: provider}, nil
}

func (m *MemoryStorageAdapter) Set(ksp *tokenstoragehelper.KeyStorageProvider) error {
	m.stored = ksp
	return nil
}

func (m *MemoryStorageAdapter) Get() (tokenstoragehelper.KeyStorageProvider, error) {
	if m.stored == nil {
		return tokenstoragehelper.KeyStorageProvider{
			NumKeys:          3,
			RotationInterval: 24 * time.Hour,
			Keys:             make(map[int]tokenstoragehelper.Key),
		}, nil
	}
	return *m.stored, nil
}

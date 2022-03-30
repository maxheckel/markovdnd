package chain

import "github.com/maxheckel/markovdnd/internal/domain"

type Store interface {
	LoadChain(name string) error
	GetChains(name string) ([]*domain.Chain, error)
}

package chain

import "github.com/maxheckel/auto-dnd/internal/domain"

type Store interface {
	LoadChain(name string) error
	GetChains(name string) ([]*domain.Chain, error)
}

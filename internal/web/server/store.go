package server

import "github.com/maxheckel/markovdnd/internal/services/store/chain/drivers"

func (s *server) SetStore(){
	// TODO replace this with a better store
	s.Store = &drivers.FilesystemDriver{}
}
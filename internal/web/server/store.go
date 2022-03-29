package server

import "github.com/maxheckel/auto-dnd/internal/services/store/chain/drivers"

func (s *server) SetStore(){
	// TODO replace this with a better store
	s.Store = &drivers.FilesystemDriver{}
}
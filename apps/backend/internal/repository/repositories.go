package repository

import "github.com/sriniously/go-tasker/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}

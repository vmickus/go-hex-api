package database

import (
	"fmt"

	"github.com/vmickus/go-hexagonal-api/internal/domain"
)

func (pg *Postgres) GetPost() (*domain.Post, error) {
	return nil, fmt.Errorf("GetPost is not implemented yet")
}

func (pg *Postgres) CreatePost(domain.Post) error {
	return fmt.Errorf("CreatePost is not implemented yet")
}

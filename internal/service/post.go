package service

import (
	"fmt"

	"github.com/vmickus/go-hexagonal-api/internal/domain"
	"github.com/vmickus/go-hexagonal-api/internal/port"
)

type PostService struct {
	postRepository port.PostRepository
}

func NewPostService(repository port.PostRepository) *PostService {
	return &PostService{
		postRepository: repository,
	}
}

func (ps *PostService) CreatePost(tile string, content string) (*domain.Post, error) {
	return nil, fmt.Errorf("Service CreatePost is not implemented yet")
}

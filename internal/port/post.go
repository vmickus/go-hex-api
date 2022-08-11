package port

import "github.com/vmickus/go-hexagonal-api/internal/domain"

type PostRepository interface {
	GetPost() (*domain.Post, error)
	CreatePost(domain.Post) error
}

type PostService interface {
	CreatePost(tile string, content string) (*domain.Post, error)
}

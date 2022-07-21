package dto

import (
	"encoding/json"
	"io"
)

type CreatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// FromJSONCreatePostRequest function    convert JSON body request into CreatePostRequest struct.
func FromJSONCreatePostRequest(body io.Reader) (*CreatePostRequest, error) {
	createPostRequest := new(CreatePostRequest)
	if err := json.NewDecoder(body).Decode(createPostRequest); err != nil {
		return nil, err
	}
	return createPostRequest, nil
}

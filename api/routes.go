package api

func (s *Server) Routes() {
	groupV1 := s.Router.Group("/v1")

	groupV1.GET("/posts", ListPosts()).Name = "list-posts"
	groupV1.POST("/posts", CreatePost(s.PostService)).Name = "create-post"
}

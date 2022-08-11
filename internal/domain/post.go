package domain

import "time"

type Post struct {
	title      string
	content    string
	created_at time.Time
}

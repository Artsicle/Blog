package models

import (
	"time"
)

type Post struct {
	Id        int64
	UserId    int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (p *Post) LinkTitle() string {
	return "Post Link Title"
}

func (p *Post) Url() string {
	return "http://www.post-link.com"
}

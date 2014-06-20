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

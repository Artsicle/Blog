package models

import (
  "time"
)

type User struct {
  Id        int64
  FirstName string `sql:"size:255"`
  LastName  string `sql:"size:255"`
  Email     string `sql:"type:varchar(100)";`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time

  Posts []Post
}

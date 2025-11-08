package domain

import (
	"time"
)

type Dataset struct {
	ID        string
	Name      string
	Format    string
	Size      int64
	Path      string
	CreatedAt time.Time
}

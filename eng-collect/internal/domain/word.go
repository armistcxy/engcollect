package domain

import (
	"context"
	"time"
)

type Word struct {
	ID           int64     `json:"id,omitempty"`
	Word         string    `json:"word"`
	Definition   string    `json:"definition"`
	Level        string    `json:"level,omitempty"`
	ExampleUsage string    `json:"example_usage,omitempty"`
	InsertDate   time.Time `json:"insert_date,omitempty"`
}

type WordRepository interface {
	GetByID(ctx context.Context, id int64) (*Word, error)
	GetByWord(ctx context.Context, word string) (*Word, error)

	// future extending: GetByLevel => []*Word, GetByInsertDate => []*Word, pagination, filter, ...

	CreateOrUpdate(ctx context.Context, word *Word) error
	Create(ctx context.Context, word *Word) error
	Update(ctx context.Context, word *Word) error
	DeleteByID(ctx context.Context, id int64) error
	DeleteByWord(ctx context.Context, word string) error
}

package repository

import (
	"context"

	"github.com/armistcxy/engcollect/internal/domain"
	"github.com/jmoiron/sqlx"
)

type sqliteWordRepository struct {
	db *sqlx.DB
}

func NewSQLiteWordRepository(db *sqlx.DB) domain.WordRepository {
	return &sqliteWordRepository{db: db}
}

func (r *sqliteWordRepository) GetByID(ctx context.Context, id int64) (*domain.Word, error) {
	// Implementation here
	return nil, nil
}

func (r *sqliteWordRepository) GetByWord(ctx context.Context, word string) (*domain.Word, error) {
	// Implementation here
	return nil, nil
}

func (r *sqliteWordRepository) CreateOrUpdate(ctx context.Context, word *domain.Word) error {
	// Implementation here
	return nil
}

func (r *sqliteWordRepository) Create(ctx context.Context, word *domain.Word) error {
	// Implementation here
	return nil
}

func (r *sqliteWordRepository) Update(ctx context.Context, word *domain.Word) error {
	// Implementation here
	return nil
}

func (r *sqliteWordRepository) DeleteByID(ctx context.Context, id int64) error {
	// Implementation here
	return nil
}

func (r *sqliteWordRepository) DeleteByWord(ctx context.Context, word string) error {
	// Implementation here
	return nil
}

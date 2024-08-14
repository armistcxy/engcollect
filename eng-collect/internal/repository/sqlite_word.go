package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

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
	query := `
		SELECT id, word, level, definition, example_usage, created_at
		FROM words
		WHERE id=? 
	`

	var resp WordResponse

	if err := r.db.GetContext(ctx, &resp, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Info("can't find word with given id in database", "id", id)
			return nil, nil
		}

		slog.Info("error occured when find word with given id in database", "id", id, "error", err.Error())
		return nil, err
	}

	wordResult := &domain.Word{
		ID:           resp.ID,
		Word:         resp.Word,
		Definition:   resp.Definition,
		ExampleUsage: resp.ExampleUsage,
		InsertDate:   resp.InsertDate,
	}
	return wordResult, nil
}

func (r *sqliteWordRepository) GetByWord(ctx context.Context, word string) (*domain.Word, error) {
	query := `
		SELECT id, word, level, definition, example_usage, created_at
		FROM words
		WHERE word=? 
	`

	var resp WordResponse

	if err := r.db.GetContext(ctx, &resp, query, word); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Info("can't find word with given word(name) in database", "word", word)
			return nil, nil
		}

		slog.Info("error occurred when find word with given word(name) in database", "word", word, "error", err.Error())
		return nil, err
	}

	wordResult := &domain.Word{
		ID:           resp.ID,
		Word:         resp.Word,
		Level:        resp.Level,
		Definition:   resp.Definition,
		ExampleUsage: resp.ExampleUsage,
		InsertDate:   resp.InsertDate,
	}
	return wordResult, nil
}

func (r *sqliteWordRepository) CreateOrUpdate(ctx context.Context, word *domain.Word) error {
	query := `
		INSERT INTO words (
			word, level, definition, example_usage
		) VALUES (?,?,?,?)
		ON CONFLICT(word) DO UPDATE_SET
			level = excluded.level, 
			description = excluded.description,
			example_usage = excluded.example_usage, 
			updated_at = date('now');
	`

	result, err := r.db.ExecContext(ctx, query, word.Word, word.Level, word.Definition, word.ExampleUsage)
	if err != nil {
		slog.Info("error occurred when trying to execute 'create or update' function", "error", err.Error())
		return err
	}

	affectRowsAmount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectRowsAmount == 0 {
		return ErrNoRowAffected
	}

	return nil
}

func (r *sqliteWordRepository) Create(ctx context.Context, word *domain.Word) error {
	query := `
		INSERT INTO words(
			word, level, definition, example_usage
		) VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, word.Word, word.Level, word.Definition, word.ExampleUsage)
	if err != nil {
		slog.Info("error occurred when insert word into database", "error", err.Error())
		return err
	}

	affectRowsAmount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectRowsAmount == 0 {
		return ErrNoRowAffected
	}

	return nil
}

func (r *sqliteWordRepository) Update(ctx context.Context, word *domain.Word) error {
	query := `
		UPDATE words 
		SET level=?,
			definition=?,
			example_usage=?
		WHERE word=?;
	`

	result, err := r.db.ExecContext(ctx, query, word.Level, word.Definition, word.ExampleUsage, word.Word)
	if err != nil {
		slog.Info("error occurred when update word in database", "error", err.Error())
		return err
	}

	affectRowsAmount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectRowsAmount == 0 {
		return ErrNoRowAffected
	}

	return nil
}

func (r *sqliteWordRepository) DeleteByID(ctx context.Context, id int64) error {
	query := `
		DELETE FROM words
		WHERE id=?;
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		slog.Info("error occurred when delete word by id", "error", err.Error())
		return err
	}

	return nil
}

func (r *sqliteWordRepository) DeleteByWord(ctx context.Context, word string) error {
	query := `
		DELETE FROM words 
		WHERE word=?;
	`

	_, err := r.db.ExecContext(ctx, query, word)
	if err != nil {
		slog.Info("error occurred when delete word by word(name)", "error", err.Error())
		return err
	}
	return nil
}

type WordResponse struct {
	ID           int64
	Word         string
	Level        string
	Definition   string
	ExampleUsage string    `db:"example_usage"`
	InsertDate   time.Time `db:"created_at"`
}

var ErrNoRowAffected = errors.New("no row affected")

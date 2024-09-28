package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/recommender-system-for-MTUCI/add_CSV_file_to_postgres/cfg"
)

const queryAddFilm = `INSERT INTO film(item_id, title, release_year, genres, countries, age_rating, directors, actors, description) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9)`

func New() (*pgxpool.Pool, error) {
	c, err := pgxpool.ParseConfig(cfg.DNS())
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе db uri: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании пула базы данных: %w", err)
	}
	return pool, nil
}

func Add(df *DTO, pool *pgxpool.Pool, ctx context.Context) error {
	_, err := pool.Exec(ctx, queryAddFilm, df.ItemID, df.Title, df.ReleaseYear, df.Genres, df.Countries, df.AgeRating, df.Directors, df.Actors, df.Description)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении фильма: %w", err)
	}
	return nil
}

package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/recommender-system-for-MTUCI/add_CSV_file_to_postgres/cfg"
)

const queryAddFilm = `INSERT INTO movie(id, title, genres, overview, production_companies, production_countries, release_data, runtime, vote_average, vote_count, actor, keywords, director, weight_rating) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

func New() (*pgxpool.Pool, error) {
	c, err := pgxpool.ParseConfig(cfg.DNS())
	if err != nil {
		return nil, fmt.Errorf("error while parse db uri: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("error while add db: %w", err)
	}
	return pool, nil
}

func Add(df *DTO, pool *pgxpool.Pool, ctx context.Context) error {
	_, err := pool.Exec(ctx, queryAddFilm, df.Id, df.Title, df.Genres, df.Overview, df.ProductionCompanies, df.ProductionCountries, df.ReleaseDate, df.RunTime, df.VoteAverage, df.VoteCount, df.Actor, df.KeyWords, df.Director, df.WeightRating)
	if err != nil {
		return fmt.Errorf("error while add film: %w", err)
	}
	return nil
}

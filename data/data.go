package data

import (
	"context"
	"log"
	"log/slog"

	"github.com/recommender-system-for-MTUCI/add_CSV_file_to_postgres/csv"
	"github.com/recommender-system-for-MTUCI/add_CSV_file_to_postgres/storage"
)

func Data() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgx, err := storage.New()
	if err != nil {
		log.Fatalf("Failed to connect with db: %v", err)
	}
	dataFrame := csv.LoadData()
	lenDataFrame, _ := dataFrame.Dims()
	for i := 0; i < lenDataFrame; i++ {
		id, err := dataFrame.Elem(i, 0).Int()
		if err != nil {
			log.Fatal("failed to convert id", err)
		}
		title := dataFrame.Elem(i, 1).String()
		releaseYear, err := dataFrame.Elem(i, 2).Int()
		if err != nil {
			log.Fatal("failed to convert release year", err)
		}
		genres := dataFrame.Elem(i, 3).String()
		countries := dataFrame.Elem(i, 4).String()
		ageRaiting, err := dataFrame.Elem(i, 5).Int()
		if err != nil {
			log.Fatal("failed to convert age raiting", err)
		}
		directors := dataFrame.Elem(i, 6).String()
		actors := dataFrame.Elem(i, 7).String()
		description := dataFrame.Elem(i, 8).String()
		err = storage.Add(&storage.DTO{
			ItemID:      id,
			Title:       title,
			ReleaseYear: releaseYear,
			Genres:      genres,
			Countries:   countries,
			AgeRating:   ageRaiting,
			Directors:   directors,
			Actors:      actors,
			Description: description,
		}, pgx, ctx)
		if err != nil {
			log.Fatal("failed to add film in db", err)
		}
	}
	slog.Info("all films added")
}

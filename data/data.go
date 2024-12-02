package data

import (
	"context"
	"fmt"
	"github.com/recommender-system-for-MTUCI/add_CSV_file_to_postgres/csv"
	"github.com/recommender-system-for-MTUCI/add_CSV_file_to_postgres/storage"
	"log"
	"strings"
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

	for j := 0; j < lenDataFrame; j++ {
		id, err := dataFrame.Elem(j, 0).Int()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		title := dataFrame.Elem(j, 1).String()
		genres := dataFrame.Elem(j, 2).String()
		cleanedGenres := strings.Trim(genres, "[]")
		cleanedGenres = strings.ReplaceAll(cleanedGenres, "\"", "")
		readyGenres := strings.Split(cleanedGenres, ", ")
		overview := dataFrame.Elem(j, 3).String()
		productionCompanies := dataFrame.Elem(j, 4).String()
		cleanedProductionCompanies := strings.Trim(productionCompanies, "[]")
		cleanedProductionCompanies = strings.ReplaceAll(cleanedProductionCompanies, "\"", "")
		readyProductionCompanies := strings.Split(cleanedProductionCompanies, ", ")
		productionContries := dataFrame.Elem(j, 5).String()
		cleanedProductionContries := strings.Trim(productionContries, "[]")
		cleanedProductionContries = strings.ReplaceAll(cleanedProductionContries, "\"", "")
		readyProductionContries := strings.Split(cleanedProductionContries, ", ")
		releaseDate := dataFrame.Elem(j, 6).String()
		runTime, err := dataFrame.Elem(j, 7).Int()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		voteAverage := dataFrame.Elem(j, 8).Float()
		voteCount, err := dataFrame.Elem(j, 9).Int()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		actor := dataFrame.Elem(j, 10).String()
		cleanedActor := strings.Trim(actor, "[]")
		cleanedActor = strings.ReplaceAll(cleanedActor, "\"", "")
		readyActor := strings.Split(cleanedActor, ", ")
		keyWords := dataFrame.Elem(j, 11).String()
		cleanedKeyWords := strings.Trim(keyWords, "[]")
		cleanedKeyWords = strings.ReplaceAll(cleanedKeyWords, "\"", "")
		readyKeyWords := strings.Split(cleanedKeyWords, ", ")
		director := dataFrame.Elem(j, 12).String()
		weightRating := dataFrame.Elem(j, 13).Float()
		data := &storage.DTO{
			Id:                  id,
			Title:               title,
			Genres:              readyGenres,
			Overview:            overview,
			ProductionCompanies: readyProductionCompanies,
			ProductionCountries: readyProductionContries,
			ReleaseDate:         releaseDate,
			RunTime:             runTime,
			VoteAverage:         voteAverage,
			VoteCount:           voteCount,
			Actor:               readyActor,
			KeyWords:            readyKeyWords,
			Director:            director,
			WeightRating:        weightRating,
		}
		err = storage.Add(data, pgx, ctx)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("all data in db")

}

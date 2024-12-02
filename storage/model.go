package storage

type DTO struct {
	Id                  int
	Title               string
	Genres              []string
	Overview            string
	ProductionCompanies []string
	ProductionCountries []string
	ReleaseDate         string
	RunTime             int
	VoteAverage         float64
	VoteCount           int
	Actor               []string
	KeyWords            []string
	Director            string
	WeightRating        float64
}

package domain

type Movie struct {
	Title       string  `json:"title"`
	Adult       bool    `json:"adult"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
}

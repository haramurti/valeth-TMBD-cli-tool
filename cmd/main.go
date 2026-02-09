package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"tmdb-cli/internal/domain"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//change from using using the struct in the same main to the importing from domains.

	apiKey := os.Getenv("TMDB_API_KEY")
	movieType := flag.String("type", "popular", "Pilih tipe film: playing, popular, top, upcoming")
	flag.Parse()

	var endpoint string
	switch *movieType {
	case "popular":
		endpoint = "popular"
	case "upcoming":
		endpoint = "upcoming"
	case "playing":
		endpoint = "now_playing"
	case "top":
		endpoint = "top_rated"
	default:
		log.Fatal("Tipe tidak valid")
	}

	baseURL := "https://api.themoviedb.org/3/movie"
	url := fmt.Sprintf("%s/%s?api_key=%s&language=en-US&page=1", baseURL, endpoint, apiKey)
	fmt.Println(url)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("🚀 Mengambil data film...")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Gagal! Status Code:", resp.StatusCode)
	}

	// 1. Siapkan Variabel Penampung
	var record domain.TMDBResponse

	// 2. Decode JSON (Unboxing)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&record); err != nil {
		log.Fatal("Gagal decode JSON:", err)
	}

	fmt.Println("\n🎬 DAFTAR FILM:", *movieType)
	fmt.Println("========================================")

	for i, movie := range record.Results {
		//Change from parsing from giving description to giving, adult ratings wether its adult movie or not.
		AdultMovie := movie.Adult

		year := movie.ReleaseDate
		if len(year) >= 4 {
			year = year[:4]
		}

		fmt.Printf("%d. %s (%s) | ⭐ %.1f\n", i+1, movie.Title, year, movie.VoteAverage)
		fmt.Println("adult? :", AdultMovie)
		fmt.Println("----------------------------------------")
	}
}

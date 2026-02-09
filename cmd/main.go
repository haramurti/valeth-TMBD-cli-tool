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

// --- 1. CETAKAN DATA (STRUCTS) ---
// Taruh ini di luar func main()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	fmt.Println("🚀 Mengambil data film...")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // Tutup pintu memori

	if resp.StatusCode != 200 {
		log.Fatal("Gagal! Status Code:", resp.StatusCode)
	}

	// --- 🚀 PHASE 5: UNBOXING & PAMER (BARU) ---

	// 1. Siapkan Variabel Penampung
	var record domain.TMDBResponse

	// 2. Decode JSON (Unboxing)
	// Kita pake Decoder karena lebih hemat memori daripada Unmarshal buat data stream
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&record); err != nil {
		log.Fatal("Gagal decode JSON:", err)
	}

	// 3. Looping & Cantik-in Output
	fmt.Println("\n🎬 DAFTAR FILM:", *movieType)
	fmt.Println("========================================")

	for i, movie := range record.Results {
		// Logika Truncate: Kalau overview kepanjangan, potong biar terminal rapi
		desc := movie.Overview
		if len(desc) > 100 {
			desc = desc[:100] + "..." // Ambil 100 huruf pertama + titik tiga
		}

		// Ambil Tahun aja (Format asli: YYYY-MM-DD)
		year := movie.ReleaseDate
		if len(year) >= 4 {
			year = year[:4] // Ambil 4 huruf pertama (Tahun)
		}

		// Print Format Cantik
		// %d = angka (nomor)
		// %s = string (teks)
		// %.1f = float dengan 1 angka di belakang koma
		fmt.Printf("%d. %s (%s) | ⭐ %.1f\n", i+1, movie.Title, year, movie.VoteAverage)
		fmt.Printf("   📝 %s\n", desc)
		fmt.Println("----------------------------------------")
	}
}

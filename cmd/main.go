package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 0. LOAD ENV DULUAN (Wajib paling atas)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		log.Fatal("API Key ga ketemu di .env! Isi dulu bang.")
	}

	// 1. SETUP FLAG
	movieType := flag.String("type", "popular", "Pilih tipe film: playing, popular, top, upcoming")
	flag.Parse()

	// 2. VALIDASI & MAPPING (Penerjemah)
	var endpoint string

	switch *movieType {
	case "popular":
		endpoint = "popular"
	case "upcoming":
		endpoint = "upcoming"
	case "playing":
		endpoint = "now_playing" // Beda istilah
	case "top":
		endpoint = "top_rated" // Beda istilah
	default:
		fmt.Println("❌ Error: Tipe film tidak dikenal!")
		os.Exit(1)
	}

	// 3. BANGUN URL (Merakit Alamat)
	// Format: https://api.themoviedb.org/3/movie/{endpoint}?api_key={apiKey}
	baseURL := "https://api.themoviedb.org/3/movie"
	url := fmt.Sprintf("%s/%s?api_key=%s&language=en-US&page=1", baseURL, endpoint, apiKey)

	fmt.Println("🚀 Menghubungi Server TMDB...")
	// Debugging: Boleh diprint URL-nya kalau mau ngecek, tapi hati-hati API Key keliatan
	// fmt.Println("Requesting:", url)

	// 4. EKSEKUSI REQUEST (Kurir Berangkat)
	resp, err := http.Get(url)

	// 5. ERROR HANDLING (Cek Ban Bocor)
	if err != nil {
		log.Fatal("Gagal request ke TMDB:", err)
	}

	// 6. DEFER CLOSE (Tutup Pintu)
	// Penting: Kode ini baru jalan PAS fungsi main mau selesai.
	// Tapi WAJIB ditulis tepat abis error checking.
	defer resp.Body.Close()

	// Cek Status Code (Biar yakin sukses 200 OK)
	if resp.StatusCode != 200 {
		log.Fatal("TMDB Marah! Status Code:", resp.StatusCode)
	}

	fmt.Println("✅ Data berhasil diambil! Status:", resp.Status)

	// Nanti Phase 5 (Decoding) lanjut di sini...
}

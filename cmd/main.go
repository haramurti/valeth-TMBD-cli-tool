package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	movieType := flag.String("type", "popular", "Pilih tipe film: playing, popular, top, upcoming")

	flag.Parse()

	switch *movieType {
	case "playing", "popular", "top", "upcoming":
		fmt.Println("Siap! Sedang mengambil data film kategori:", *movieType)
	default:
		fmt.Println("❌ Error: Tipe film tidak dikenal!")
		fmt.Println("Tolong pilih salah satu: playing, popular, top, upcoming")

		os.Exit(1)
	}

}

//comments to fix

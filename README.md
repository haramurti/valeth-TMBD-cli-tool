README.md By Neval Harimurti

```markdown
# 🎬 TMDB CLI Tool: Bioskop Mini di Layar Hitam

Selamat datang di project gabut tapi berfaedah ini.

Pernah nggak sih lo buka aplikasi streaming, niatnya mau nonton, tapi malah abis 2 jam cuma buat *scrolling* katalog film? Akhirnya ketiduran, kuota abis, dan hidup terasa hampa. Gue sering.

Makanya gue bikin **TMDB CLI Tool** ini. Sebuah alat sederhana berbasis **Golang** yang bikin terminal lo jadi pusat informasi film dunia. Tanpa iklan, tanpa loading poster yang berat, dan yang paling penting: bikin lo keliatan kayak *hacker* di depan gebetan (kalau punya).

## 🤔 Ini Apaan Sih?

Ini adalah *Command Line Interface* (CLI) buat ngambil data dari **The Movie Database (TMDB)**. Lo bisa liat:
- Film yang lagi **Tayang** (*Now Playing*) — buat yang mau ngajak jalan tapi bingung nonton apa.
- Film **Populer** (*Popular*) — biar nggak kudet pas nongkrong.
- Film **Rating Tertinggi** (*Top Rated*) — buat sobat *cinephile* yang seleranya tinggi.
- Film **Akan Datang** (*Upcoming*) — buat nabung duit tiket dari sekarang.

## 🛠️ Persiapan (Jangan Di-skip)

Sebelum lo bisa ngerasain sensasi jadi *developer* keren, pastiin laptop lo udah punya ini:

1.  **Golang** (Versi berapa aja asal jangan jaman batu).
2.  **API Key TMDB** (Gratis kok, daftar aja di [themoviedb.org](https://www.themoviedb.org/). Kalau males daftar, pinjem temen, tapi dosa tanggung sendiri).
3.  **Terminal** (Bukan terminal bus ya, Command Prompt atau iTerm).

## 🚀 Cara Install (Sabar, Dikit Lagi)

1.  **Clone atau Download** repo ini. Anggap aja lagi mindahin kenangan mantan ke harddisk eksternal.
    ```bash
    git clone [https://github.com/username-lo/tmdb-cli.git](https://github.com/username-lo/tmdb-cli.git)
    cd tmdb-cli
    ```

2.  **Bikin File Rahasia**
    Di dunia ini ada hal yang nggak boleh dibagi, salah satunya sikat gigi dan API Key.
    - Buat file bernama `.env` di folder utama.
    - Isi file itu dengan mantra berikut:
      ```env
      TMDB_API_KEY=masukkan_kunci_rahasia_lo_disini
      ```
    *(Ingat: Jangan pake spasi, jangan pake tanda kutip. Polosan aja kayak hati lo pas lagi kosong).*

3.  **Rapikan Dependensi**
    Biar Go-nya nggak bingung nyari alat perang.
    ```bash
    go mod tidy
    ```

## 🏃‍♂️ Cara Pakai (Saatnya Pamer)

Lo punya dua cara buat jalanin ini. Pilih sesuai kepribadian lo.

### Cara 1: Si Praktis (Langsung Run)
Buat lo yang nggak mau ribet kompilasi.
```bash
go run cmd/main.go --type popular

```

### Cara 2: Si Dedikasi (Build Dulu)

Buat lo yang pengen punya file aplikasi sendiri (.exe atau binary).

```bash
# Masak dulu kodenya
go build -o bioskop-mini cmd/main.go

# Sajikan selagi hangat
./bioskop-mini --type upcoming

```

### 📜 Menu Perintah

Ganti bagian `--type` dengan salah satu opsi di bawah ini:

* `popular` : Film yang lagi hits. (Default, kalau lo males ngetik argumen).
* `playing` : Film yang lagi tayang di bioskop sekarang.
* `top`     : Film-film legendaris sepanjang masa.
* `upcoming`: Film masa depan (siapa tau ada jodoh lo di sana).

Contoh:

```bash
./bioskop-mini --type top

```

## 🐛 Troubleshooting (Kalau Error)

Hidup itu penuh masalah, begitu juga codingan.

* **Error 401 Unauthorized**: Itu artinya API Key lo salah, atau lo salah copy token yang panjang banget itu. Pake yang **API Key (v3)** ya, yang pendek aja.
* **Error ".env file not found"**: Lo naruh file `.env` di mana? Harus satu folder sama file `main.go` atau sejajar sama tempat lo nge-run perintahnya.
* **Terminal Kosong**: Cek koneksi internet. Siapa tau wifi tetangga ganti password.

## 🤝 Kontribusi

Kalau lo nemu *bug* atau mau nambahin fitur (misalnya fitur nyari pacar via sinopsis film), silakan *Fork* dan bikin *Pull Request*. Gue terima dengan tangan terbuka, asalkan kodenya rapi dan nggak bikin sakit mata.

---

*Dibuat dengan ❤️, sedikit ☕, dan banyak 🐛 saat debugging.*

```

```

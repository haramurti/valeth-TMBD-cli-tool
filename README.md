
# tmbd cli tool by Neval harimurti
# 🎬 TMDB CLI Tool: Bioskop Mini di Layar Hitam

Selamat datang di project hasil kegabutan yang berfaedah.

Pernah nggak sih lo buka aplikasi streaming, niatnya mau nonton, tapi malah abis 2 jam cuma buat *scrolling* katalog film? Akhirnya ketiduran, kuota abis, dan hidup terasa hampa. Gue sering.

Makanya gue bikin **TMDB CLI Tool** ini. Sebuah alat sederhana berbasis **Golang** yang bikin terminal lo jadi pusat informasi film dunia. Tanpa iklan, tanpa loading poster yang berat, dan yang paling penting: bikin lo keliatan kayak *hacker* lagi nge-hack situs NASA di depan gebetan.

## 🤔 Ini Apaan Sih?

Ini adalah *Command Line Interface* (CLI) buat ngambil data dari **The Movie Database (TMDB)**. Lo bisa liat:
- 🎬 **Now Playing** (`playing`) — Buat yang mau ngajak jalan tapi bingung nonton apa.
- 🔥 **Popular** (`popular`) — Biar nggak kudet pas nongkrong sama temen yang sok tau film.
- ⭐ **Top Rated** (`top`) — Buat sobat *cinephile* yang seleranya tinggi dan anti film rating rendah.
- 📅 **Upcoming** (`upcoming`) — Buat nabung duit tiket dari sekarang.

---

## 🛠️ Persiapan (Jangan Di-skip)

Sebelum lo bisa ngerasain sensasi jadi *developer* keren, pastiin laptop lo udah punya ini:

1.  **Golang** (Versi berapa aja asal jangan jaman batu).
2.  **API Key TMDB** (Gratis kok, daftar aja di [themoviedb.org](https://www.themoviedb.org/). Kalau males daftar, pinjem temen, tapi dosa tanggung sendiri).
3.  **Terminal** (Bukan terminal bus ya; Command Prompt, PowerShell, atau iTerm).

---

## 🚀 Cara Install & Jalanin

Ikutin langkah ini urut dari atas ke bawah. Jangan lompat-lompat kayak kangguru.

### 1. Clone & Setup
Anggap aja lagi mindahin kenangan mantan ke harddisk eksternal.

~~~bash
# Download project
git clone [https://github.com/username-lo/tmdb-cli.git](https://github.com/username-lo/tmdb-cli.git)
cd tmdb-cli

# Rapikan dependensi (biar Go gak bingung)
go mod tidy
~~~

### 2. Bikin Kunci Rahasia (.env)
Di dunia ini ada hal yang nggak boleh dibagi, salah satunya sikat gigi dan API Key.
* Buat file bernama `.env` di folder ini (sejajar sama `go.mod`).
* Isi file itu dengan mantra berikut:

~~~env
TMDB_API_KEY=masukkan_kunci_rahasia_lo_disini_tanpa_spasi
~~~
*(Ingat: Jangan pake spasi, jangan pake tanda kutip. Polosan aja).*

### 3. Mulai Nonton (Run Project)
Lo bisa langsung jalanin (Run) atau bungkus jadi aplikasi (Build). Pilih salah satu command di bawah ini:

**Opsi A: Langsung Jalan (Paling Cepet)**
~~~bash
go run cmd/main.go --type popular
~~~

**Opsi B: Masak Jadi Aplikasi (Biar punya file .exe)**
~~~bash
# Build dulu jadi binary
go build -o bioskop-mini cmd/main.go

# Jalanin aplikasinya (Mac/Linux/Windows)
./bioskop-mini --type upcoming
~~~

---

## 📜 Menu Perintah Lengkap

Ganti bagian `--type` dengan salah satu opsi di bawah ini sesuai *mood* lo:

| Perintah | Fungsi |
| :--- | :--- |
| `popular` | Film yang lagi hits. (Default). |
| `playing` | Film yang lagi tayang di bioskop sekarang. |
| `top` | Film-film legendaris sepanjang masa. |
| `upcoming` | Film masa depan. |

Contoh penggunaan lain:
~~~bash
go run cmd/main.go --type top
~~~

---

## 🐛 Troubleshooting (Kalau Error)

Hidup itu penuh masalah, begitu juga codingan. Kalau error, jangan panik, tarik napas.

- **Error 401 Unauthorized**: Itu artinya API Key lo salah, atau lo salah copy token yang panjang banget itu. Pake yang **API Key (v3)** ya, yang pendek aja.
- **Error ".env file not found"**: Lo naruh file `.env` di mana? Harus satu folder sama file `go.mod`.
- **Error "no such file or directory"**: Pastiin lo jalanin perintahnya dari folder utama project, bukan dari dalem folder `cmd`.

---

## 🤝 Kontribusi

Kalau lo nemu *bug* atau mau nambahin fitur (misalnya fitur nyari pacar via sinopsis film), silakan *Fork* dan bikin *Pull Request*. Gue terima dengan tangan terbuka, asalkan kodenya rapi dan nggak bikin sakit mata.

---

*gua bikin sambil ngantuk jadi maaf kalau error *

project was inspired by roadmap.sh beginner project

https://roadmap.sh/projects/tmdb-cli

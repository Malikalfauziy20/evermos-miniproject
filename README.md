# Evermos Mini Project — Backend Service Transaksi

Service backend REST API untuk platform *social commerce* Evermos, dibangun sebagai bagian dari **program magang project-based Rakamin x Evermos Academy** (peran: Backend Developer).

Evermos adalah platform social commerce reseller yang menjual produk-produk Muslim Indonesia, dengan fitur katalog barang, toko online, dan sistem distributor. Project ini fokus membangun service inti untuk transaksi penjualan di dalamnya.

---

## ✨ Fitur

- **Autentikasi JWT** — register & login, dengan toko yang otomatis terbuat saat user mendaftar (dalam satu database transaction, atomic)
- **Manajemen akun** — user hanya bisa melihat & mengubah datanya sendiri
- **Manajemen toko** — update profil toko, upload foto toko
- **Manajemen alamat** — CRUD alamat pengiriman milik user
- **Kategori produk** — khusus dapat dikelola oleh admin
- **Manajemen produk** — CRUD lengkap dengan **pagination**, **filtering**, dan **upload foto produk**
- **Transaksi** — proses checkout dengan:
  - Snapshot data produk ke tabel `log_produk` (audit trail — riwayat transaksi tidak berubah walau produk aslinya diedit/dihapus)
  - Pengurangan stok otomatis
  - Perhitungan total harga otomatis
  - Seluruh proses dibungkus dalam satu database transaction (atomic — kalau satu langkah gagal, semua di-rollback)
- **Isolasi data antar user** — user tidak dapat mengakses atau mengelola data milik user lain (alamat, toko, produk, maupun transaksi)

---

## 🛠️ Tech Stack

| Kategori | Teknologi |
|---|---|
| Bahasa | [Go (Golang)](https://go.dev/) |
| Web Framework | [Fiber](https://gofiber.io/) |
| ORM | [GORM](https://gorm.io/) |
| Database | MySQL |
| Autentikasi | JWT (JSON Web Token) |
| Hashing Password | bcrypt |
| Dokumentasi API | Postman |
| Arsitektur | Clean Architecture |

---

## 🏗️ Arsitektur

Project ini mengikuti prinsip **Clean Architecture**, memisahkan tanggung jawab tiap layer secara jelas:

```
Request → Router → Handler → Usecase → Repository → Database
```


Pemisahan ini membuat business logic (usecase) tidak bergantung pada framework HTTP maupun detail database, sehingga lebih mudah diuji dan dikembangkan.

---

## 🔐 Endpoint Utama

| Method | Endpoint | Keterangan |
|---|---|---|
| POST | `/auth/register` | Registrasi user + toko otomatis terbuat |
| POST | `/auth/login` | Login, mengembalikan JWT token |
| GET/PUT | `/user` | Lihat / update profil sendiri |
| GET/PUT | `/toko/:id` | Lihat toko / update toko (khusus pemilik) |
| CRUD | `/alamat` | Kelola alamat pengiriman (khusus pemilik) |
| GET/POST/PUT/DELETE | `/kategori` | Kelola kategori (create/update/delete khusus admin) |
| CRUD | `/produk` | Kelola produk, dengan pagination & filtering |
| POST | `/produk/:id/foto` | Upload foto produk |
| POST | `/trx` | Checkout transaksi |
| GET | `/trx` , `/trx/:id` | Riwayat transaksi (khusus pemilik) |

Seluruh endpoint (kecuali register & login) dilindungi middleware JWT, dan endpoint kategori (create/update/delete) tambahan dilindungi middleware admin-only.

---

## 🗄️ Skema Database (ringkas)

`users` → `tokos` (1:1) → `produks` (1:N) → `foto_produks` (1:N)
`users` → `alamats` (1:N)
`transaksis` → `detail_transaksis` → `log_produks` (snapshot produk saat transaksi)

Tabel `log_produk` sengaja tidak memakai foreign key ke tabel produk — ia menyimpan **salinan/snapshot** data produk persis pada saat transaksi terjadi, agar riwayat transaksi lama tetap akurat meskipun data produk aslinya berubah di kemudian hari.

---

## 🚀 Cara Menjalankan

1. Clone repository ini
2. Siapkan database MySQL, buat database kosong (misal `evermos_db`)
3. Copy `.env.example` menjadi `.env`, sesuaikan kredensial database kamu
4. Install dependency & jalankan:
```bash
   go mod tidy
   go run cmd/main.go
```
5. Server berjalan di `http://localhost:8080`, tabel database otomatis dibuat lewat migration saat pertama kali dijalankan
6. Import collection Postman untuk mencoba seluruh endpoint

---

## 📌 Catatan

Project ini merupakan hasil pengerjaan mini project dalam program magang project-based **Rakamin Academy x Evermos**, dikerjakan sebagai simulasi tugas backend developer sesungguhnya — mulai dari desain arsitektur, implementasi fitur, hingga pengujian manual endpoint demi endpoint.

---

**Dibuat oleh:** Malik Alfauziy Hermawan — S1 Ilmu Komputer, Informatika, Universitas Bhayangkara Jakarta Raya II
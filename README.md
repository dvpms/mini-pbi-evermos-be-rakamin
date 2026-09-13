# Evermos E-Commerce RESTful API - Mini Project Backend

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.10.0-008080?style=flat&logo=gin)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-v1.25.12-blue?style=flat)](https://gorm.io/)
[![MySQL](https://img.shields.io/badge/MySQL-8.0_Aiven_Cloud-4479A1?style=flat&logo=mysql)](https://aiven.io/)
[![Newman Tests](https://img.shields.io/badge/Newman_Assertions-72%2F72_Passed-success?style=flat&logo=postman)](https://www.postman.com/)
[![Clean Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-brightgreen?style=flat)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Proyek ini merupakan implementasi RESTful API Backend e-commerce untuk **Virtual Internship Evermos (Rakamin Academy)** yang dibangun menggunakan bahasa pemrograman **Go (Golang)**, framework **Gin Gonic**, ORM **GORM**, dan database cloud **MySQL (Aiven Cloud)** dengan menerapkan standar **Clean Architecture**, **SOLID Principles**, serta **Security & User Isolation**.

---

## 📑 Daftar Isi

1. [Fitur Utama](#-fitur-utama)
2. [Arsitektur & Struktur Direktori](#-arsitektur--struktur-direktori)
3. [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
4. [Persyaratan Sistem](#-persyaratan-sistem)
5. [Instalasi & Menjalankan Aplikasi](#-instalasi--menjalankan-aplikasi)
6. [Database Seeder & Akun Default](#-database-seeder--akun-default)
7. [Pengujian Otomatis (Newman & Postman)](#-pengujian-otomatis-newman--postman)
8. [Katalog Endpoint API](#-katalog-endpoint-api)
9. [Standar & Aturan Bisnis (Soal & Ketentuan)](#-standar--aturan-bisnis-soal--ketentuan)

---

## 🚀 Fitur Utama

- 🔐 **Autentikasi & Otorisasi**: Registrasi user, login dengan JWT (JSON Web Token), hashing password menggunakan `bcrypt`, dan middleware proteksi role Admin.
- 🏬 **Manajemen Toko Otomatis**: Toko (*store*) otomatis terbuat secara instan saat user berhasil mendaftar (Ketentuan #6).
- 📍 **Manajemen Alamat & Wilayah Indonesia**: CRUD alamat pengiriman user dengan integrasi API Wilayah Indonesia (EMSIFA API) untuk provinsi dan kota/kabupaten.
- 🗂️ **Kategori Produk (Admin Only)**: Pengelolaan kategori produk yang diproteksi khusus untuk pengguna dengan role `admin` (Ketentuan #8).
- 🛍️ **Katalog Produk & Upload Foto**:
  - Dukungan multiple upload foto produk tersimpan pada static server `/uploads`.
  - Pembuatan URL *slug* SEO-friendly secara otomatis dari judul produk.
  - Pencarian dan filter dinamis (`nama_produk`, `category_id`, `toko_id`, `min_harga`, `max_harga`).
  - Pagination dinamis (`page` dan `limit`).
- 💳 **Checkout Transaksi & Log Produk Snapshot**:
  - Transaksi database atomik (`gorm.Transaction`) dengan *row-level locking* (`FOR UPDATE`) untuk mencegah *race condition* saat stok berkurang.
  - Pencatatan otomatis snapshot data produk ke tabel `log_produks` saat transaksi berhasil dibuat (Ketentuan #16 & #17).
  - Validasi kepemilikan alamat dan isolasi data riwayat transaksi per user.
- 🌱 **Database Seeding Otomatis**: Inisialisasi data kategori default dan akun super admin secara instan saat migrasi dijalankan.

---

## 🏛️ Arsitektur & Struktur Direktori

Proyek ini mengadopsi **Clean Architecture** berlapis untuk memastikan pemisahan tanggung jawab (*Separation of Concerns*) dan kemudahan pengujian (*Testability*):

```
mini-project-pbi/
├── config/             # Konfigurasi aplikasi & JWT helpers
├── database/           # Inisialisasi GORM connection pool, auto-migration & seeder
│   ├── database.go
│   └── seeder.go
├── docs/               # Dokumentasi pengujian, task checklist, dan spesifikasi soal
│   └── testing-guide.md
├── handlers/           # Presentation Layer (HTTP Request/Response Controllers)
│   ├── alamat_handler.go
│   ├── auth_handler.go
│   ├── category_handler.go
│   ├── produk_handler.go
│   ├── provcity_handler.go
│   ├── toko_handler.go
│   ├── transaksi_handler.go
│   └── user_handler.go
├── middleware/         # Gin Middlewares (JWT Authentication & Role Guard)
│   └── auth_middleware.go
├── models/             # Entity Model GORM & DTO Request/Response Contracts
│   ├── alamat.go
│   ├── category.go
│   ├── log_produk.go
│   ├── produk.go
│   ├── response.go
│   ├── toko.go
│   ├── transaksi.go
│   └── user.go
├── repositories/       # Data Access Layer (GORM queries & atomic DB transactions)
│   ├── alamat_repository.go
│   ├── category_repository.go
│   ├── produk_repository.go
│   ├── toko_repository.go
│   ├── transaksi_repository.go
│   └── user_repository.go
├── routes/             # Dependency Injection & Modular Route Registrations
│   ├── auth_routes.go
│   ├── category_routes.go
│   ├── product_routes.go
│   ├── provcity_routes.go
│   ├── routes.go
│   ├── toko_routes.go
│   ├── trx_routes.go
│   └── user_routes.go
├── services/           # Business Logic Layer (Validasi domain & user isolation)
│   ├── alamat_service.go
│   ├── auth_service.go
│   ├── category_service.go
│   ├── produk_service.go
│   ├── provcity_service.go
│   ├── toko_service.go
│   ├── transaksi_service.go
│   └── user_service.go
├── utils/              # Hash (Bcrypt), JWT generator, dan SEO Slug utility
├── uploads/            # Direktori penyimpanan static file upload
├── .env                # Konfigurasi environment database & port
├── main.go             # Entry point aplikasi backend
└── mini_project.postman_collection.json # File koleksi pengujian Postman & Newman
```

---

## 🛠️ Teknologi yang Digunakan

- **Language**: [Go (Golang)](https://golang.org/) 1.20+
- **HTTP Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: [MySQL 8.0](https://www.mysql.com/) / [Aiven Cloud MySQL](https://aiven.io/)
- **Authentication**: JWT ([golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt))
- **Security**: [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) Password Hashing
- **Testing Runner**: [Newman CLI](https://www.npmjs.com/package/newman) & [Postman](https://www.postman.com/)

---

## 📋 Persyaratan Sistem

Sebelum menjalankan proyek ini, pastikan sistem Anda telah terpasang:

1. **Go** (versi 1.20 atau lebih baru)
2. **MySQL Server** (versi 8.0+) atau instance Cloud MySQL
3. **Node.js** & **npm** (opsional, untuk menjalankan Newman CLI)
4. **Git**

---

## ⚙️ Instalasi & Menjalankan Aplikasi

### 1. Clone Repositori

```bash
git clone https://github.com/dvpms/mini-pbi-evermos-be-rakamin.git
cd mini-project-pbi
```

### 2. Konfigurasi Environment (`.env`)

Buat atau sesuaikan file `.env` pada direktori root proyek:

```env
# Database Configuration
DB_USER=root
DB_PASS=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=evermos_mini

# Application Configuration
JWT_SECRET="secret_evermos_mini_project_key_2025"
PORT="8080"
AUTO_MIGRATE=true
```

> **Catatan**: 
> - Set `AUTO_MIGRATE=true` saat pertama kali menjalankan aplikasi agar tabel database dibuat secara otomatis dan data seeder terisi.
> - Aplikasi juga mendukung `DATABASE_URL` (format DSN URI / cloud connection string) sebagai alternatif `DB_*`.

### 3. Unduh Dependencies

```bash
go mod tidy
```

### 4. Jalankan Server

```bash
go run main.go
```

Server akan aktif dan siap menerima request pada `http://localhost:8080`.

---

## 🌱 Database Seeder & Akun Default

Ketika `AUTO_MIGRATE=true` diaktifkan pada `.env`, aplikasi akan secara otomatis menjalankan seeder idempotensial (`database/seeder.go`):

### 1. Akun Default Admin

| Field | Nilai |
| --- | --- |
| **Nama** | Super Admin |
| **Email** | `admin@evermos.com` |
| **Password** | `admin123` |
| **No. Telepon** | `081234567890` |
| **Role** | `admin` *(Dapat mengelola CRUD kategori)* |
| **Toko Default** | Evermos Official Store |

### 2. Kategori Default

1. **Elektronik**
2. **Fashion**
3. **Makanan & Minuman**
4. **Kecantikan**
5. **Perlengkapan Rumah**

---

## 🧪 Pengujian Otomatis (Newman & Postman)

Koleksi Postman [`mini_project.postman_collection.json`](file:///c:/Users/UDevran/Downloads/Documents/BE-Evermos/code/mini-project-pbi/mini_project.postman_collection.json) telah dilengkapi dengan **36 skenario pengujian** dan **72 poin validasi (assertions)** yang mencakup pengujian alur positif (*happy path*), penanganan error/keamanan (*negative test*), isolasi kepemilikan data, hingga transaksi stok.

### Menjalankan via Newman CLI

```bash
# Install Newman jika belum ada
npm install -g newman

# Jalankan seluruh test suite
newman run mini_project.postman_collection.json
```

**Hasil Pengujian:**

```
┌─────────────────────────┬───────────────────┬───────────────────┐
│                         │          executed │            failed │
├─────────────────────────┼───────────────────┼───────────────────┤
│              iterations │                 1 │                 0 │
│                requests │                36 │                 0 │
│            test-scripts │                36 │                 0 │
│      prerequest-scripts │                 1 │                 0 │
│              assertions │                72 │                 0 │
├─────────────────────────┴───────────────────┴───────────────────┤
│ total run duration: 27.4s                                       │
│ average response time: 677ms                                    │
└─────────────────────────────────────────────────────────────────┘
```

### 📋 Ringkasan Skenario Pengujian (Bahasa Non-Teknis)

Pengujian dibagi ke dalam 7 modul fungsional utama:

#### 1. 🔐 Pendaftaran & Keamanan Akun (4 Skenario)
- **Pendaftaran Pengguna Baru**: Calon pengguna berhasil mendaftar dan otomatis dibuatkan toko untuk berjualan.
- **Pencegahan Akun Duplikat**: Sistem menolak pendaftaran jika nomor HP atau email sudah terdaftar sebelumnya.
- **Masuk Akun (Login) Berhasil**: Pengguna dengan data yang benar berhasil masuk dan memperoleh akses digital resmi.
- **Pencegahan Salah Kata Sandi**: Sistem menolak upaya masuk jika kata sandi salah demi keamanan akun.

#### 2. 👤 Profil & Buku Alamat Pengiriman (7 Skenario)
- **Melihat Profil Pribadi**: Pengguna dapat melihat data akun miliknya sendiri.
- **Memperbarui Profil**: Pengguna dapat mengubah nama dan data kontak pribadi.
- **Menambah Alamat Pengiriman**: Pengguna dapat menyimpan alamat baru untuk pengiriman barang belanjaan.
- **Melihat Daftar Alamat**: Menampilkan seluruh alamat yang sudah tersimpan di akun pengguna.
- **Melihat Rincian Alamat Tertentu**: Memeriksa detail alamat spesifik yang dipilih.
- **Mengubah Data Alamat**: Pengguna dapat memperbaiki nama jalan, kode pos, atau informasi alamat.
- **Menghapus Alamat**: Pengguna dapat menghapus alamat yang sudah tidak digunakan.

#### 3. 🏬 Pengelolaan Toko Penjual (4 Skenario)
- **Melihat Toko Pribadi**: Penjual dapat melihat rincian toko resmi miliknya.
- **Melihat Toko Penjual Lain**: Pembeli dapat melihat profil toko lain saat hendak membeli barang.
- **Menjelajahi Daftar Toko**: Menampilkan seluruh toko secara rapi per halaman (*pagination*).
- **Memperbarui Informasi Toko & Logo**: Penjual dapat mengubah nama/deskripsi toko dan mengunggah foto logo toko.

#### 4. 🗂️ Pengelompokan Kategori Barang (3 Skenario)
- **Melihat Seluruh Kategori**: Pembeli dan penjual dapat melihat daftar kategori produk publik.
- **Perlindungan Hak Akses Admin**: Sistem memblokir pengguna biasa yang mencoba menambah kategori (khusus Admin platform).
- **Pengecekan Kategori Tidak Ada**: Sistem menampilkan pesan yang jelas jika kategori yang dicari tidak ditemukan.

#### 5. 🛍️ Katalog Produk & Upload Foto (7 Skenario)
- **Menambah Produk Baru + Upload Foto**: Penjual dapat memasang barang dagangan baru beserta multiple foto produk.
- **Menjelajahi Seluruh Produk**: Katalog barang tampil rapi dan terorganisir per halaman.
- **Pencarian & Filter Cerdas**: Pembeli dapat mencari barang berdasarkan nama serta batas harga minimum/maksimum.
- **Melihat Rincian Lengkap Produk**: Menampilkan detail harga, deskripsi, stok, dan toko pemilik barang.
- **Mengubah Informasi Produk**: Penjual dapat memperbarui harga, deskripsi, dan foto barang dagangannya.
- **Menghapus Produk**: Penjual dapat menghapus produk yang sudah tidak dijual lagi.
- **Validasi Produk Terhapus**: Memastikan produk yang sudah dihapus tidak bisa diakses kembali oleh pembeli.

#### 6. 📍 Data Wilayah Pengiriman Indonesia (4 Skenario)
- **Daftar Seluruh Provinsi**: Menampilkan daftar resmi seluruh provinsi di Indonesia.
- **Daftar Kota/Kabupaten**: Menampilkan daftar kota/kabupaten sesuai provinsi yang dipilih.
- **Rincian Nama Provinsi**: Memastikan data nama provinsi akurat.
- **Rincian Nama Kota/Kabupaten**: Memastikan data nama kota/kabupaten akurat.

#### 7. 💳 Transaksi Belanja & Checkout (7 Skenario)
- **Persiapan Alamat Belanja**: Memastikan alamat pengiriman pembeli valid sebelum belanja.
- **Persiapan Barang Belanja**: Memastikan barang dagangan siap dipesan.
- **Checkout Pembelian Berhasil**: Pembelian sukses, stok barang otomatis terpotong, dan riwayat belanja tersimpan.
- **Pencegahan Pembelian Melebihi Stok (*Zero Over-Selling*)**: Sistem otomatis menolak transaksi jika jumlah pesanan melebihi sisa stok di toko.
- **Melihat Seluruh Riwayat Belanja**: Pengguna dapat melihat daftar transaksi yang pernah dilakukannya.
- **Melihat Rincian Satu Transaksi**: Pengguna dapat melihat rincian barang, total bayar, dan alamat kirim pada pesanan tertentu.
- **Pengecekan Transaksi Tidak Ada**: Sistem menolak akses ke nomor transaksi palsu atau tidak terdaftar.

---

## 📡 Katalog Endpoint API

Format standar response JSON:

```json
{
    "status": true,
    "message": "Succeed to ...",
    "errors": null,
    "data": { ... }
}
```

### 1. Autentikasi (`/auth`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `POST` | `/auth/register` | Mendaftarkan user baru (otomatis membuat toko) | Public |
| `POST` | `/auth/login` | Login user & mendapatkan token JWT | Public |

### 2. User & Alamat Kirim (`/user`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `GET` | `/user` | Mendapatkan profil user yang sedang login | Protected (JWT) |
| `PUT` | `/user` | Mengubah profil user | Protected (JWT) |
| `GET` | `/user/alamat` | Menampilkan seluruh alamat kirim user login | Protected (JWT) |
| `GET` | `/user/alamat/:id` | Mengambil detail alamat kirim | Protected (JWT) |
| `POST` | `/user/alamat` | Menambahkan alamat kirim baru | Protected (JWT) |
| `PUT` | `/user/alamat/:id` | Mengubah data alamat kirim | Protected (JWT) |
| `DELETE` | `/user/alamat/:id` | Menghapus alamat kirim | Protected (JWT) |

### 3. Toko (`/toko`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `GET` | `/toko/my` | Mendapatkan data toko milik user login | Protected (JWT) |
| `GET` | `/toko/:id_toko` | Mendapatkan detail toko berdasarkan ID | Protected (JWT) |
| `GET` | `/toko` | Menampilkan list toko terpaginasi (`page`, `limit`) | Protected (JWT) |
| `PUT` | `/toko/:id_toko` | Mengubah profil toko & upload logo | Protected (JWT) |

### 4. Kategori (`/category`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `GET` | `/category` | Mendapatkan seluruh daftar kategori produk | Public |
| `GET` | `/category/:id` | Mendapatkan detail kategori | Public |
| `POST` | `/category` | Membuat kategori baru | Admin Only |
| `PUT` | `/category/:id` | Mengubah nama kategori | Admin Only |
| `DELETE` | `/category/:id` | Menghapus kategori | Admin Only |

### 5. Produk (`/product`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `GET` | `/product` | Filter & pagination produk (`nama_produk`, `category_id`, `toko_id`, `min_harga`, `max_harga`, `page`, `limit`) | Public |
| `GET` | `/product/:id` | Mendapatkan detail produk lengkap | Public |
| `POST` | `/product` | Menambah produk baru + upload multiple foto | Protected (JWT) |
| `PUT` | `/product/:id` | Mengubah data & foto produk milik toko user login | Protected (JWT) |
| `DELETE` | `/product/:id` | Menghapus produk toko milik user login | Protected (JWT) |

### 6. Transaksi (`/trx`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `POST` | `/trx` | Checkout pembelian (Atomic DB transaction, pemotongan stok, log snapshot) | Protected (JWT) |
| `GET` | `/trx` | Menampilkan riwayat transaksi pengguna login | Protected (JWT) |
| `GET` | `/trx/:id` | Menampilkan detail transaksi pengguna login | Protected (JWT) |

### 7. Wilayah Indonesia (`/provcity`)

| Method | Endpoint | Deskripsi | Akses |
| --- | --- | --- | --- |
| `GET` | `/provcity/listprovincies` | Daftar seluruh provinsi di Indonesia | Public |
| `GET` | `/provcity/listcities/:prov_id` | Daftar kota/kabupaten berdasarkan ID provinsi | Public |
| `GET` | `/provcity/detailprovince/:prov_id` | Detail nama provinsi | Public |
| `GET` | `/provcity/detailcity/:city_id` | Detail nama kota | Public |

---

## 🔒 Standar & Aturan Bisnis (Soal & Ketentuan)

1. **Routing & Payload Compatibility (Ketentuan #1 & #2)**: Struktur route, status code, format response JSON, dan variabel header sesuai dengan acuan Postman Collection Virtual Internship.
2. **Validasi Unik Email & No Telepon (Ketentuan #3)**: Registrasi menolak duplikasi email dan nomor telepon.
3. **Keamanan Autentikasi JWT (Ketentuan #4)**: Header `token: <jwt_token>` divalidasi pada setiap endpoint terproteksi.
4. **Static File Upload (Ketentuan #5)**: File upload (logo toko, foto produk) disimpan di `./uploads` dan dapat diakses publik via `/uploads/*filepath`.
5. **Auto Store Creation (Ketentuan #6)**: Toko dibuat otomatis saat user registrasi.
6. **Proteksi Role Admin (Ketentuan #8)**: Middleware `AdminOnly` membatasi mutasi kategori khusus untuk user dengan role `admin`.
7. **Isolasi Data Pengguna (Ketentuan #11 - #15)**:
   - User tidak dapat melihat atau mengubah profil, alamat, toko, produk, atau transaksi milik user lain (`403 Forbidden` / `404 Not Found`).
8. **Atomic Transaction & Snapshot Log Produk (Ketentuan #16 & #17)**:
   - Checkout menggunakan database transaction (`tx.Begin()`, `tx.Commit()`, `tx.Rollback()`) dengan row-level lock (`FOR UPDATE`) pada stok produk.
   - Snapshot data produk disimpan ke tabel `log_produks` pada setiap transaksi.
9. **Clean Architecture (Ketentuan #18)**: Arsitektur berlapis yang rapi memisahkan Handler, Service, Repository, Model, Route, dan Middleware.

---

## 📄 Lisensi & Kontributor

- **Project**: Evermos Virtual Internship Mini Project (Rakamin Academy)
- **Author**: Devran Perdana Malik
- **Tahun**: 2026

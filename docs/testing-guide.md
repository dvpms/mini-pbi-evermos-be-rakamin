# Panduan Pengujian & Dokumentasi API - Evermos Mini Project

## 1. Ringkasan Proyek
Proyek ini adalah implementasi RESTful API Backend e-commerce untuk **Virtual Internship Evermos (Rakamin Academy)** yang dibangun menggunakan bahasa pemrograman **Go (Golang)** dengan framework **Gin**, ORM **GORM**, dan database cloud **MySQL (Aiven Cloud)** serta menerapkan prinsip **Clean Architecture** dan **SOLID**.

---

## 2. Arsitektur & Struktur Direktori
```
mini-project-pbi/
├── config/             # Konfigurasi aplikasi & JWT
├── database/           # Inisialisasi koneksi GORM & schema auto-migration
├── docs/               # Panduan tugas, spesifikasi soal, dan testing guide
│   ├── code-standar.md
│   ├── soal.md
│   ├── testing-guide.md
│   └── tasks/
├── handlers/           # HTTP Request handlers (Gin controller layer)
│   ├── alamat_handler.go
│   ├── auth_handler.go
│   ├── category_handler.go
│   ├── produk_handler.go
│   ├── provcity_handler.go
│   ├── toko_handler.go
│   ├── transaksi_handler.go
│   └── user_handler.go
├── middleware/         # Gin Middleware (JWT Auth & Role Guard)
├── models/             # GORM database models & DTO Request/Response contracts
├── repositories/       # Data Access Layer & DB atomic transactions
├── routes/             # Dependency injection & centralized route registrations
├── services/           # Business Logic Layer & user isolation guards
├── utils/              # Hash, JWT, and Slug utilities
├── uploads/            # Static file storage untuk foto toko & produk
├── .env                # Environment variables
├── main.go             # Entry point aplikasi
└── mini_project.postman_collection.json # Comprehensive Postman test collection
```

---

## 3. Prasyarat & Instalasi

### Prasyarat
1. **Go** (versi 1.20 ke atas)
2. **Node.js & Newman** (untuk automated CLI testing)
   ```bash
   npm install -g newman
   ```

### Konfigurasi Environment (`.env`)
Pastikan file `.env` di root direktori memiliki konfigurasi berikut:
```env
DATABASE_URL="avnadmin:<password>@tcp(<host>:<port>)/defaultdb?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify"
JWT_SECRET="secret_evermos_mini_project_key_2025"
PORT="8080"
AUTO_MIGRATE=false
```

---

## 4. Menjalankan Server Backend

Jalankan perintah berikut pada terminal:
```bash
go run main.go
```
Server akan berjalan dan siap menerima request pada `http://localhost:8080`.

---

## 5. Menjalankan Automated Testing (Newman CLI)

Untuk menjalankan seluruh rangkaian pengujian end-to-end secara otomatis dari Phase 1 hingga Phase 6:
```bash
newman run mini_project.postman_collection.json
```

**Hasil Pengujian:**
- **36 HTTP Requests** dieksekusi secara berurutan.
- **72 Test Assertions** lulus dengan status `0 failed`.

---

## 6. Panduan Pengujian via Postman UI

1. Buka aplikasi **Postman**.
2. Klik tombol **Import** di kiri atas.
3. Pilih file `mini_project.postman_collection.json`.
4. Koleksi akan otomatis mengatur variabel environment:
   - `{{local}}`: `http://localhost:8080`
   - `{{token}}`: Ditangkap otomatis setelah login / register.
   - `{{test_email}}`, `{{test_phone}}`: Dibuat acak secara dinamis oleh Pre-request Script.
5. Klik **Run Collection** untuk mengeksekusi semua test case sekaligus.

---

## 7. Katalog Endpoint & Kontrak Response

Semua response API mengikuti standar struktur data:
```json
{
    "status": true,
    "message": "Succeed to ...",
    "errors": null,
    "data": { ... }
}
```

### A. Autentikasi (`/auth`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `POST` | `/auth/register` | Mendaftarkan user baru & otomatis membuat toko | Public |
| `POST` | `/auth/login` | Login menggunakan nomor telepon / email & password | Public |

### B. Profil & Alamat Pengiriman (`/user`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `GET` | `/user` | Mengambil profil user yang sedang login | Protected (JWT) |
| `PUT` | `/user` | Mengubah profil user yang sedang login | Protected (JWT) |
| `GET` | `/user/alamat` | Menampilkan seluruh alamat kirim user login | Protected (JWT) |
| `GET` | `/user/alamat/:id` | Mengambil detail alamat kirim | Protected (JWT) |
| `POST` | `/user/alamat` | Menambahkan alamat kirim baru | Protected (JWT) |
| `PUT` | `/user/alamat/:id` | Mengubah alamat kirim milik user login | Protected (JWT) |
| `DELETE` | `/user/alamat/:id` | Menghapus alamat kirim milik user login | Protected (JWT) |

### C. Toko (`/toko`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `GET` | `/toko/my` | Mengambil data toko milik user yang sedang login | Protected (JWT) |
| `GET` | `/toko/:id_toko` | Mengambil data detail toko berdasarkan ID | Protected (JWT) |
| `GET` | `/toko` | Menampilkan list toko dengan pagination (`page`, `limit`, `nama`) | Protected (JWT) |
| `PUT` | `/toko/:id_toko` | Mengubah profil toko & upload foto logo (Multipart) | Protected (JWT) |

### D. Kategori (`/category`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `GET` | `/category` | Mengambil seluruh daftar kategori produk | Public |
| `GET` | `/category/:id` | Mengambil detail kategori berdasarkan ID | Public |
| `POST` | `/category` | Menambah kategori baru | Admin Only |
| `PUT` | `/category/:id` | Mengubah nama kategori | Admin Only |
| `DELETE` | `/category/:id` | Menghapus kategori | Admin Only |

### E. Produk (`/product`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `GET` | `/product` | Pencarian produk dinamis (`nama_produk`, `category_id`, `toko_id`, `min_harga`, `max_harga`, `page`, `limit`) | Public |
| `GET` | `/product/:id` | Mengambil detail produk beserta foto, toko, dan kategori | Public |
| `POST` | `/product` | Menambah produk baru + upload multiple foto (Multipart) | Protected (JWT) |
| `PUT` | `/product/:id` | Mengubah data & foto produk milik toko user login (Multipart) | Protected (JWT) |
| `DELETE` | `/product/:id` | Menghapus produk milik toko user login | Protected (JWT) |

### F. Transaksi (`/trx`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `POST` | `/trx` | Checkout pembelian produk (Atomic DB transaction, validasi stok, pemotongan stok, snapshot log produk) | Protected (JWT) |
| `GET` | `/trx` | Menampilkan riwayat transaksi pengguna login | Protected (JWT) |
| `GET` | `/trx/:id` | Menampilkan detail transaksi pengguna login (404 jika bukan milik user) | Protected (JWT) |

### G. Wilayah Indonesia (`/provcity`)
| Method | Endpoint | Deskripsi | Akses |
|---|---|---|---|
| `GET` | `/provcity/listprovincies` | Mengambil daftar seluruh provinsi di Indonesia | Public |
| `GET` | `/provcity/listcities/:prov_id` | Mengambil daftar kota berdasarkan ID provinsi | Public |
| `GET` | `/provcity/detailprovince/:prov_id` | Mengambil nama detail provinsi | Public |
| `GET` | `/provcity/detailcity/:city_id` | Mengambil nama detail kota | Public |

---

## 8. Verifikasi Kepatuhan Terhadap Ketentuan Soal

1. **Routing Postman Match (Ketentuan #1 & #2)**: Seluruh rute dan payload sesuai 100% dengan spesifikasi Postman Collection Evermos.
2. **Unique Email & No Telepon (Ketentuan #3)**: Registrasi menolak pendaftaran nomor telepon atau email duplikat dengan response `400 Bad Request`.
3. **JWT Authentication (Ketentuan #4)**: Token JWT diverifikasi pada header `token: <jwt_token>` dengan masa berlaku 24 jam.
4. **File Upload (Ketentuan #5)**: Upload foto toko dan multiple foto produk tersimpan di folder `/uploads` dan dapat diakses statis via URL.
5. **Auto Store Creation (Ketentuan #6)**: Toko otomatis dibuat bersamaan saat registrasi user baru.
6. **Alamat Pengiriman (Ketentuan #7)**: Alamat user divalidasi kepemilikannya dan dikaitkan saat checkout transaksi.
7. **Admin Only Kategori (Ketentuan #8)**: Mutasi kategori diproteksi middleware `AdminOnly`, menolak user biasa dengan status `401/403`.
8. **Pagination & Filtering (Ketentuan #9 & #10)**: Query pagination dan filter harga/kategori/nama produk diimplementasikan di layer repository.
9. **User Isolation (Ketentuan #11 - #15)**: User tidak dapat membaca, mengubah, atau menghapus data user, alamat, toko, produk, atau riwayat transaksi milik user lain.
10. **Snapshot Log Produk & Stok (Ketentuan #16 & #17)**: Setiap checkout secara otomatis membuat snapshot produk di tabel `log_produks` dan memotong stok di `produks` secara atomik (`gorm.Transaction` + `FOR UPDATE`).
11. **Clean Architecture (Ketentuan #18)**: Pemisahan kode yang tegas antara Handler, Service, Repository, Model, dan Middleware.

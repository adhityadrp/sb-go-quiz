# sb-go-quiz
Prasyarat
---------

Daftar endpoint (base URL: https://sb-go-quiz-sb-go-quiz.up.railway.app)
---------------------------------------------------------------

Berikut daftar endpoint yang didefinisikan di `routes/routes.go`. Tambahkan prefix base URL di depan setiap path untuk membentuk URL lengkap.

Base URL: https://sb-go-quiz-sb-go-quiz.up.railway.app

Public endpoints:

- POST /api/users/login

Protected endpoints (memerlukan header Authorization: Bearer <token>):

- GET /api/categories
- POST /api/categories
- GET /api/categories/:id
- DELETE /api/categories/:id
- GET /api/categories/:id/books

- GET /api/books
- POST /api/books
- GET /api/books/:id
- DELETE /api/books/:id

Contoh URL lengkap:

- https://sb-go-quiz-sb-go-quiz.up.railway.app/api/users/login
- https://sb-go-quiz-sb-go-quiz.up.railway.app/api/categories

- Go (minimal versi 1.20 direkomendasikan).
- PostgreSQL yang bisa diakses dari aplikasi.
- Variabel environment yang diperlukan (lihat bagian Konfigurasi).
Berikut daftar path file dan folder utama di proyek ini:

- `go.mod`
- `main.go`
- `config/database.go`
- `controllers/book_controller.go`
- `controllers/category_controller.go`
- `controllers/user_controller.go`
- `middlewares/auth_jwt.go`
- `migrations/001_init_tables.sql`
- `migrations/002_seed_admin_user.sql`
- `models/book.go`
- `models/category.go`
- `models/user.go`
- `routes/routes.go`
- `scripts/generate_hash.go`
- `README.md`


- Go (minimal versi 1.20 direkomendasikan).
- PostgreSQL yang bisa diakses dari aplikasi.
- Variabel environment yang diperlukan (lihat bagian Konfigurasi).

Konfigurasi (.env)
------------------

Sebelum menjalankan aplikasi, pastikan variabel environment diatur (bisa lewat file `.env` atau environment system):

- `DB_HOST` — host database (contoh: `localhost`).
- `DB_PORT` — port PostgreSQL (contoh: `5432`).
- `DB_USER` — username database.
- `DB_PASSWORD` — password database.
- `DB_NAME` — nama database.
- `DB_SSLMODE` — sslmode untuk koneksi (contoh: `disable`).
- `APP_PORT` — port aplikasi (default `8080`).

Migrasi & Seeder
-----------------

Project menggunakan folder `migrations/` yang berisi file SQL dengan marker `-- +migrate Up` dan `-- +migrate Down`.

- Saat aplikasi dijalankan, fungsi `config.RunMigrations()` akan mencari file `migrations/*.sql` dan mengeksekusi hanya bagian `Up` dari setiap file (antara `-- +migrate Up` dan `-- +migrate Down`).
- Contoh file:
  - `001_init_tables.sql` — membuat tabel `users`, `categories`, `books`.
  - `002_seed_admin_user.sql` — menambahkan user admin.

Catatan penting: implementasi migrasi pada `config.RunMigrations()` mengeksekusi SQL mentah dari bagian `Up`. Pastikan file migrasi ditulis dengan benar dan urutan nama file (prefix numerik) memastikan urutan eksekusi.

Admin seeder
------------

File `migrations/002_seed_admin_user.sql` memasukkan user dengan `username` = `admin` dan password `admin` sudah di-hash menggunakan bcrypt. Password yang disimpan bukan teks asli, melainkan hash. Jika ingin mengubah password admin, jalankan ulang seeder menggunakan hash baru atau modifikasi file seeder.

Menjalankan aplikasi (PowerShell)
--------------------------------

1. Pastikan environment variables tersedia (mis. buat file `.env` di root):

```powershell
# contoh .env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=sb_go_quiz
DB_SSLMODE=disable
APP_PORT=8080
```

2. Install dependency (jika belum):

```powershell
# jalankan di folder project
go get ./...
```

3. (Opsional) Jika perlu, generate hash bcrypt untuk password baru (ada di `scripts/generate_hash.go`):

```powershell
cd scripts
go run generate_hash.go
# Salin hasil hashed password dan gunakan bila perlu.
```

4. Jalankan aplikasi:

```powershell
# dari root project
go run .
```

Perilaku saat start: aplikasi akan memuat environment, inisialisasi koneksi DB, menjalankan migrasi (hanya Up section), lalu memulai server HTTP.

Verifikasi migrasi berhasil
--------------------------

Setelah aplikasi berjalan, Anda bisa mengecek bahwa tabel dibuat di PostgreSQL. Contoh menggunakan `psql`:

```powershell
psql -h <DB_HOST> -p <DB_PORT> -U <DB_USER> -d <DB_NAME>
# lalu di dalam psql:
# \dt   -- untuk melihat daftar tabel
# select * from users;  -- cek data admin
```

Keamanan & notes
-----------------

- Pastikan file `.env` tidak di-commit ke repository.
- Password admin yang disertakan di seeder adalah hash bcrypt; jangan menyimpan password plaintext.
- Jika ingin migrasi yang lebih kuat (rollback, pencatatan versi), pertimbangkan menggunakan library migrasi seperti `golang-migrate/migrate`.

Troubleshooting
---------------

- Jika tabel tidak muncul:
  - Pastikan koneksi database benar (cek `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`).
  - Lihat log aplikasi, fungsi migrasi mencetak nama file yang dieksekusi.
  - Pastikan SQL di antara marker `-- +migrate Up` valid untuk PostgreSQL.

- Jika server gagal start karena migrasi gagal: perbaiki SQL migrasi, atau ubah isi file migrasi, lalu jalankan ulang.

Lisensi
-------

Lisensi: bebas dipakai untuk pendidikan dan eksperimen

# sb-go-quiz

Deskripsi
---------

`sb-go-quiz` adalah aplikasi backend sederhana yang ditulis menggunakan bahasa Go. Aplikasi ini menyediakan model dasar untuk buku, kategori, dan pengguna; contoh endpoint untuk operasi CRUD; serta mekanisme migrasi SQL dan seeding data awal (admin). Tujuan proyek ini adalah sebagai latihan membuat REST API, koneksi ke PostgreSQL, autentikasi, dan manajemen migrasi sederhana.

Struktur proyek
----------------

Berikut struktur folder utama dan fungsi singkat tiap folder/file:

- `main.go` — entry point aplikasi, menginisialisasi DB, JWT, menjalankan migrasi dan server.
- `config/database.go` — inisialisasi koneksi database dan fungsi `RunMigrations()` untuk menjalankan file SQL di folder `migrations`.
- `controllers/` — berisi controller untuk resource (`book_controller.go`, `category_controller.go`, `user_controller.go`).
- `models/` — model-data (`book.go`, `category.go`, `user.go`). Pada `user.go` disertakan helper hashing password menggunakan bcrypt.
- `routes/routes.go` — pengaturan route/endpoint aplikasi.
- `middlewares/` — middleware seperti JWT (`auth_jwt.go`).
- `migrations/` — file SQL migrasi dan seeder. Contoh:
  - `001_init_tables.sql` — membuat tabel `users`, `categories`, `books`.
  - `002_seed_admin_user.sql` — menambahkan user `admin` (password `admin`, telah di-hash dengan bcrypt).
- `scripts/` — skrip bantu, mis. `generate_hash.go` untuk menghasilkan hash bcrypt dari password.

Prasyarat
---------

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
